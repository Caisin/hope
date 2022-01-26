package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysapi/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysapi"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type sysApiRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysApiRepo .
func NewSysApiRepo(data *Data, logger log.Logger) biz.SysApiRepo {
	return &sysApiRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysApi 创建
func (r *sysApiRepo) CreateSysApi(ctx context.Context, req *v1.SysApiCreateReq) (*ent.SysApi, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.SysApi.Create().
		SetHandle(req.Handle).
		SetTitle(req.Title).
		SetPath(req.Path).
		SetAction(req.Action).
		SetType(req.Type).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteSysApi 删除
func (r *sysApiRepo) DeleteSysApi(ctx context.Context, req *v1.SysApiDeleteReq) error {
	return r.data.db.SysApi.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysApi 批量删除
func (r *sysApiRepo) BatchDeleteSysApi(ctx context.Context, req *v1.SysApiBatchDeleteReq) (int, error) {
	return r.data.db.SysApi.Delete().Where(sysapi.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysApi 更新
func (r *sysApiRepo) UpdateSysApi(ctx context.Context, req *v1.SysApiUpdateReq) (*ent.SysApi, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.SysApi.UpdateOneID(req.Id).
		SetHandle(req.Handle).
		SetTitle(req.Title).
		SetPath(req.Path).
		SetAction(req.Action).
		SetType(req.Type).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetSysApi 根据Id查询
func (r *sysApiRepo) GetSysApi(ctx context.Context, req *v1.SysApiReq) (*ent.SysApi, error) {
	return r.data.db.SysApi.Get(ctx, req.Id)
}

// PageSysApi 分页查询
func (r *sysApiRepo) PageSysApi(ctx context.Context, req *v1.SysApiPageReq) ([]*ent.SysApi, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.SysApi.
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
func (r *sysApiRepo) genCondition(req *v1.SysApiReq) []predicate.SysApi {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysApi, 0)
	if req.Id > 0 {
		list = append(list, sysapi.ID(req.Id))
	}
	if str.IsBlank(req.Handle) {
		list = append(list, sysapi.HandleContains(req.Handle))
	}
	if str.IsBlank(req.Title) {
		list = append(list, sysapi.TitleContains(req.Title))
	}
	if str.IsBlank(req.Path) {
		list = append(list, sysapi.PathContains(req.Path))
	}
	if str.IsBlank(req.Action) {
		list = append(list, sysapi.ActionContains(req.Action))
	}
	if str.IsBlank(req.Type) {
		list = append(list, sysapi.TypeContains(req.Type))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysapi.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysapi.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysapi.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysapi.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysapi.TenantId(req.TenantId))
	}

	return list
}
