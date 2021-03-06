package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/appversion/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/appversion"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type appVersionRepo struct {
	data *Data
	log  *log.Helper
}

// NewAppVersionRepo .
func NewAppVersionRepo(data *Data, logger log.Logger) biz.AppVersionRepo {
	return &appVersionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateAppVersion 创建
func (r *appVersionRepo) CreateAppVersion(ctx context.Context, req *v1.AppVersionCreateReq) (*ent.AppVersion, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.AppVersion.Create().
		SetTitle(req.Title).
		SetVersion(req.Version).
		SetUpdateInfo(req.UpdateInfo).
		SetDownloadUrl(req.DownloadUrl).
		SetPlatform(req.Platform).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteAppVersion 删除
func (r *appVersionRepo) DeleteAppVersion(ctx context.Context, req *v1.AppVersionDeleteReq) error {
	return r.data.db.AppVersion.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteAppVersion 批量删除
func (r *appVersionRepo) BatchDeleteAppVersion(ctx context.Context, req *v1.AppVersionBatchDeleteReq) (int, error) {
	return r.data.db.AppVersion.Delete().Where(appversion.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateAppVersion 更新
func (r *appVersionRepo) UpdateAppVersion(ctx context.Context, req *v1.AppVersionUpdateReq) (*ent.AppVersion, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.AppVersion.UpdateOneID(req.Id).
		SetTitle(req.Title).
		SetVersion(req.Version).
		SetUpdateInfo(req.UpdateInfo).
		SetDownloadUrl(req.DownloadUrl).
		SetPlatform(req.Platform).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetAppVersion 根据Id查询
func (r *appVersionRepo) GetAppVersion(ctx context.Context, req *v1.AppVersionReq) (*ent.AppVersion, error) {
	return r.data.db.AppVersion.Get(ctx, req.Id)
}

// PageAppVersion 分页查询
func (r *appVersionRepo) PageAppVersion(ctx context.Context, req *v1.AppVersionPageReq) ([]*ent.AppVersion, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.AppVersion.
		Query().
		Where(
			//查询条件构造
			r.genCondition(req.Param)...,
		)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	req.GetPagin().Total = int64(count)
	if count == 0 {
		return nil, nil
	}
	query.Limit(int(p.GetPageSize())).
		Offset(int(p.GetOffSet()))
	if p.NeedOrder() {
		if p.IsDesc() {
			query.Order(ent.Desc(p.GetField()))
		} else {
			query.Order(ent.Asc(p.GetField()))
		}
	}
	return query.All(ctx)
}

// genCondition 构造查询条件
func (r *appVersionRepo) genCondition(req *v1.AppVersionReq) []predicate.AppVersion {
	if req == nil {
		return nil
	}
	list := make([]predicate.AppVersion, 0)
	if req.Id > 0 {
		list = append(list, appversion.ID(req.Id))
	}
	if str.IsNotBlank(req.Title) {
		list = append(list, appversion.TitleContains(req.Title))
	}
	if req.Version > 0 {
		list = append(list, appversion.Version(req.Version))
	}
	if str.IsNotBlank(req.UpdateInfo) {
		list = append(list, appversion.UpdateInfoContains(req.UpdateInfo))
	}
	if str.IsNotBlank(req.DownloadUrl) {
		list = append(list, appversion.DownloadUrlContains(req.DownloadUrl))
	}
	if str.IsNotBlank(req.Platform) {
		list = append(list, appversion.PlatformContains(req.Platform))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, appversion.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, appversion.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, appversion.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, appversion.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, appversion.TenantId(req.TenantId))
	}

	return list
}
