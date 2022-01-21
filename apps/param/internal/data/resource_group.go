package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/resourcegroup/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/resourcegroup"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type resourceGroupRepo struct {
	data *Data
	log  *log.Helper
}

// NewResourceGroupRepo .
func NewResourceGroupRepo(data *Data, logger log.Logger) biz.ResourceGroupRepo {
	return &resourceGroupRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateResourceGroup 创建
func (r *resourceGroupRepo) CreateResourceGroup(ctx context.Context, req *v1.ResourceGroupCreateReq) (*ent.ResourceGroup, error) {
	now := time.Now()
	return r.data.db.ResourceGroup.Create().
		SetName(req.Name).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteResourceGroup 删除
func (r *resourceGroupRepo) DeleteResourceGroup(ctx context.Context, req *v1.ResourceGroupDeleteReq) error {
	return r.data.db.ResourceGroup.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteResourceGroup 批量删除
func (r *resourceGroupRepo) BatchDeleteResourceGroup(ctx context.Context, req *v1.ResourceGroupBatchDeleteReq) (int, error) {
	return r.data.db.ResourceGroup.Delete().Where(resourcegroup.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateResourceGroup 更新
func (r *resourceGroupRepo) UpdateResourceGroup(ctx context.Context, req *v1.ResourceGroupUpdateReq) (*ent.ResourceGroup, error) {
	return r.data.db.ResourceGroup.UpdateOne(convert.ResourceGroupUpdateReq2Data(req)).Save(ctx)
}

// GetResourceGroup 根据Id查询
func (r *resourceGroupRepo) GetResourceGroup(ctx context.Context, req *v1.ResourceGroupReq) (*ent.ResourceGroup, error) {
	return r.data.db.ResourceGroup.Get(ctx, req.Id)
}

// PageResourceGroup 分页查询
func (r *resourceGroupRepo) PageResourceGroup(ctx context.Context, req *v1.ResourceGroupPageReq) ([]*ent.ResourceGroup, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.ResourceGroup.
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
func (r *resourceGroupRepo) genCondition(req *v1.ResourceGroupReq) []predicate.ResourceGroup {
	if req == nil {
		return nil
	}
	list := make([]predicate.ResourceGroup, 0)
	if req.Id > 0 {
		list = append(list, resourcegroup.ID(req.Id))
	}
	if str.IsBlank(req.Name) {
		list = append(list, resourcegroup.NameContains(req.Name))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, resourcegroup.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, resourcegroup.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, resourcegroup.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, resourcegroup.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, resourcegroup.TenantId(req.TenantId))
	}

	return list
}
