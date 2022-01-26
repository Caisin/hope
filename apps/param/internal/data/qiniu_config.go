package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/qiniuconfig/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/qiniuconfig"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type qiniuConfigRepo struct {
	data *Data
	log  *log.Helper
}

// NewQiniuConfigRepo .
func NewQiniuConfigRepo(data *Data, logger log.Logger) biz.QiniuConfigRepo {
	return &qiniuConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateQiniuConfig 创建
func (r *qiniuConfigRepo) CreateQiniuConfig(ctx context.Context, req *v1.QiniuConfigCreateReq) (*ent.QiniuConfig, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.QiniuConfig.Create().
		SetAccessKey(req.AccessKey).
		SetBucket(req.Bucket).
		SetHost(req.Host).
		SetSecretKey(req.SecretKey).
		SetType(req.Type).
		SetZone(req.Zone).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteQiniuConfig 删除
func (r *qiniuConfigRepo) DeleteQiniuConfig(ctx context.Context, req *v1.QiniuConfigDeleteReq) error {
	return r.data.db.QiniuConfig.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteQiniuConfig 批量删除
func (r *qiniuConfigRepo) BatchDeleteQiniuConfig(ctx context.Context, req *v1.QiniuConfigBatchDeleteReq) (int, error) {
	return r.data.db.QiniuConfig.Delete().Where(qiniuconfig.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateQiniuConfig 更新
func (r *qiniuConfigRepo) UpdateQiniuConfig(ctx context.Context, req *v1.QiniuConfigUpdateReq) (*ent.QiniuConfig, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.QiniuConfig.UpdateOneID(req.Id).
		SetAccessKey(req.AccessKey).
		SetBucket(req.Bucket).
		SetHost(req.Host).
		SetSecretKey(req.SecretKey).
		SetType(req.Type).
		SetZone(req.Zone).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetQiniuConfig 根据Id查询
func (r *qiniuConfigRepo) GetQiniuConfig(ctx context.Context, req *v1.QiniuConfigReq) (*ent.QiniuConfig, error) {
	return r.data.db.QiniuConfig.Get(ctx, req.Id)
}

// PageQiniuConfig 分页查询
func (r *qiniuConfigRepo) PageQiniuConfig(ctx context.Context, req *v1.QiniuConfigPageReq) ([]*ent.QiniuConfig, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.QiniuConfig.
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
func (r *qiniuConfigRepo) genCondition(req *v1.QiniuConfigReq) []predicate.QiniuConfig {
	if req == nil {
		return nil
	}
	list := make([]predicate.QiniuConfig, 0)
	if req.Id > 0 {
		list = append(list, qiniuconfig.ID(req.Id))
	}
	if str.IsBlank(req.AccessKey) {
		list = append(list, qiniuconfig.AccessKeyContains(req.AccessKey))
	}
	if str.IsBlank(req.Bucket) {
		list = append(list, qiniuconfig.BucketContains(req.Bucket))
	}
	if str.IsBlank(req.Host) {
		list = append(list, qiniuconfig.HostContains(req.Host))
	}
	if str.IsBlank(req.SecretKey) {
		list = append(list, qiniuconfig.SecretKeyContains(req.SecretKey))
	}
	if str.IsBlank(req.Type) {
		list = append(list, qiniuconfig.TypeContains(req.Type))
	}
	if str.IsBlank(req.Zone) {
		list = append(list, qiniuconfig.ZoneContains(req.Zone))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, qiniuconfig.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, qiniuconfig.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, qiniuconfig.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, qiniuconfig.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, qiniuconfig.TenantId(req.TenantId))
	}

	return list
}
