package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/pageconfig/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/pageconfig"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/pkg/util/str"
	"time"
)

type pageConfigRepo struct {
	data *Data
	log  *log.Helper
}

// NewPageConfigRepo .
func NewPageConfigRepo(data *Data, logger log.Logger) biz.PageConfigRepo {
	return &pageConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreatePageConfig 创建
func (r *pageConfigRepo) CreatePageConfig(ctx context.Context, req *v1.PageConfigCreateReq) (*ent.PageConfig, error) {
	now := time.Now()
	return r.data.db.PageConfig.Create().
		SetPageCode(req.PageCode).
		SetPageName(req.PageName).
		SetGroupCodes(req.GroupCodes).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeletePageConfig 删除
func (r *pageConfigRepo) DeletePageConfig(ctx context.Context, req *v1.PageConfigDeleteReq) error {
	return r.data.db.PageConfig.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeletePageConfig 批量删除
func (r *pageConfigRepo) BatchDeletePageConfig(ctx context.Context, req *v1.PageConfigBatchDeleteReq) (int, error) {
	return r.data.db.PageConfig.Delete().Where(pageconfig.IDIn(req.Ids...)).Exec(ctx)
}

// UpdatePageConfig 更新
func (r *pageConfigRepo) UpdatePageConfig(ctx context.Context, req *v1.PageConfigUpdateReq) (*ent.PageConfig, error) {
	return r.data.db.PageConfig.UpdateOne(convert.PageConfigUpdateReq2Data(req)).Save(ctx)
}

// GetPageConfig 根据Id查询
func (r *pageConfigRepo) GetPageConfig(ctx context.Context, req *v1.PageConfigReq) (*ent.PageConfig, error) {
	return r.data.db.PageConfig.Get(ctx, req.Id)
}

// PagePageConfig 分页查询
func (r *pageConfigRepo) PagePageConfig(ctx context.Context, req *v1.PageConfigPageReq) ([]*ent.PageConfig, error) {
	pagin := req.Pagin
	query := r.data.db.PageConfig.
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
	query.Limit(int(pagin.GetPage())).
		Offset(int(pagin.GetOffSet()))
	if pagin.NeedOrder() {
		if pagin.IsDesc() {
			query.Order(ent.Desc(pagin.GetField()))
		} else {
			query.Order(ent.Asc(pagin.GetField()))
		}
	}
	return query.All(ctx)
}

// genCondition 构造查询条件
func (r *pageConfigRepo) genCondition(req *v1.PageConfigReq) []predicate.PageConfig {
	list := make([]predicate.PageConfig, 0)
	if req.Id > 0 {
		list = append(list, pageconfig.ID(req.Id))
	}
	if str.IsBlank(req.PageCode) {
		list = append(list, pageconfig.PageCodeContains(req.PageCode))
	}
	if str.IsBlank(req.PageName) {
		list = append(list, pageconfig.PageNameContains(req.PageName))
	}
	if str.IsBlank(req.GroupCodes) {
		list = append(list, pageconfig.GroupCodesContains(req.GroupCodes))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, pageconfig.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, pageconfig.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, pageconfig.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, pageconfig.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, pageconfig.TenantId(req.TenantId))
	}

	return list
}