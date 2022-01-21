package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/assetitem/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/assetitem"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type assetItemRepo struct {
	data *Data
	log  *log.Helper
}

// NewAssetItemRepo .
func NewAssetItemRepo(data *Data, logger log.Logger) biz.AssetItemRepo {
	return &assetItemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateAssetItem 创建
func (r *assetItemRepo) CreateAssetItem(ctx context.Context, req *v1.AssetItemCreateReq) (*ent.AssetItem, error) {
	now := time.Now()
	return r.data.db.AssetItem.Create().
		SetAssetItemId(req.AssetItemId).
		SetAssetName(req.AssetName).
		SetCashTag(req.CashTag).
		SetValidDays(req.ValidDays).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteAssetItem 删除
func (r *assetItemRepo) DeleteAssetItem(ctx context.Context, req *v1.AssetItemDeleteReq) error {
	return r.data.db.AssetItem.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteAssetItem 批量删除
func (r *assetItemRepo) BatchDeleteAssetItem(ctx context.Context, req *v1.AssetItemBatchDeleteReq) (int, error) {
	return r.data.db.AssetItem.Delete().Where(assetitem.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateAssetItem 更新
func (r *assetItemRepo) UpdateAssetItem(ctx context.Context, req *v1.AssetItemUpdateReq) (*ent.AssetItem, error) {
	return r.data.db.AssetItem.UpdateOne(convert.AssetItemUpdateReq2Data(req)).Save(ctx)
}

// GetAssetItem 根据Id查询
func (r *assetItemRepo) GetAssetItem(ctx context.Context, req *v1.AssetItemReq) (*ent.AssetItem, error) {
	return r.data.db.AssetItem.Get(ctx, req.Id)
}

// PageAssetItem 分页查询
func (r *assetItemRepo) PageAssetItem(ctx context.Context, req *v1.AssetItemPageReq) ([]*ent.AssetItem, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.AssetItem.
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
func (r *assetItemRepo) genCondition(req *v1.AssetItemReq) []predicate.AssetItem {
	if req == nil {
		return nil
	}
	list := make([]predicate.AssetItem, 0)
	if req.Id > 0 {
		list = append(list, assetitem.ID(req.Id))
	}
	if req.AssetItemId > 0 {
		list = append(list, assetitem.AssetItemId(req.AssetItemId))
	}
	if str.IsBlank(req.AssetName) {
		list = append(list, assetitem.AssetNameContains(req.AssetName))
	}
	if req.CashTag > 0 {
		list = append(list, assetitem.CashTag(req.CashTag))
	}
	if req.ValidDays > 0 {
		list = append(list, assetitem.ValidDays(req.ValidDays))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, assetitem.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, assetitem.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, assetitem.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, assetitem.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, assetitem.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, assetitem.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, assetitem.TenantId(req.TenantId))
	}

	return list
}
