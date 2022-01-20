package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/userresource/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/userresource"
	"hope/pkg/util/str"
	"time"
)

type userResourceRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserResourceRepo .
func NewUserResourceRepo(data *Data, logger log.Logger) biz.UserResourceRepo {
	return &userResourceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUserResource 创建
func (r *userResourceRepo) CreateUserResource(ctx context.Context, req *v1.UserResourceCreateReq) (*ent.UserResource, error) {
	now := time.Now()
	return r.data.db.UserResource.Create().
		SetResType(req.ResType).
		SetName(req.Name).
		SetURL(req.Url).
		SetSummary(req.Summary).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteUserResource 删除
func (r *userResourceRepo) DeleteUserResource(ctx context.Context, req *v1.UserResourceDeleteReq) error {
	return r.data.db.UserResource.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteUserResource 批量删除
func (r *userResourceRepo) BatchDeleteUserResource(ctx context.Context, req *v1.UserResourceBatchDeleteReq) (int, error) {
	return r.data.db.UserResource.Delete().Where(userresource.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateUserResource 更新
func (r *userResourceRepo) UpdateUserResource(ctx context.Context, req *v1.UserResourceUpdateReq) (*ent.UserResource, error) {
	return r.data.db.UserResource.UpdateOne(convert.UserResourceUpdateReq2Data(req)).Save(ctx)
}

// GetUserResource 根据Id查询
func (r *userResourceRepo) GetUserResource(ctx context.Context, req *v1.UserResourceReq) (*ent.UserResource, error) {
	return r.data.db.UserResource.Get(ctx, req.Id)
}

// PageUserResource 分页查询
func (r *userResourceRepo) PageUserResource(ctx context.Context, req *v1.UserResourcePageReq) ([]*ent.UserResource, error) {
	pagin := req.Pagin
	query := r.data.db.UserResource.
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
func (r *userResourceRepo) genCondition(req *v1.UserResourceReq) []predicate.UserResource {
	list := make([]predicate.UserResource, 0)
	if req.Id > 0 {
		list = append(list, userresource.ID(req.Id))
	}
	if str.IsBlank(req.ResType) {
		list = append(list, userresource.ResTypeContains(req.ResType))
	}
	if str.IsBlank(req.Name) {
		list = append(list, userresource.NameContains(req.Name))
	}
	if str.IsBlank(req.Url) {
		list = append(list, userresource.URLContains(req.Url))
	}
	if str.IsBlank(req.Summary) {
		list = append(list, userresource.SummaryContains(req.Summary))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, userresource.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, userresource.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, userresource.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, userresource.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, userresource.TenantId(req.TenantId))
	}

	return list
}
