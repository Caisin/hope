package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novelconsume/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novelconsume"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/auth"

	"hope/pkg/pagin"
	"time"
)

type novelConsumeRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelConsumeRepo .
func NewNovelConsumeRepo(data *Data, logger log.Logger) biz.NovelConsumeRepo {
	return &novelConsumeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelConsume 创建
func (r *novelConsumeRepo) CreateNovelConsume(ctx context.Context, req *v1.NovelConsumeCreateReq) (*ent.NovelConsume, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.NovelConsume.Create().
		SetNovelId(req.NovelId).
		SetCoin(req.Coin).
		SetCoupon(req.Coupon).
		SetDiscount(req.Discount).
		SetReward(req.Reward).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteNovelConsume 删除
func (r *novelConsumeRepo) DeleteNovelConsume(ctx context.Context, req *v1.NovelConsumeDeleteReq) error {
	return r.data.db.NovelConsume.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelConsume 批量删除
func (r *novelConsumeRepo) BatchDeleteNovelConsume(ctx context.Context, req *v1.NovelConsumeBatchDeleteReq) (int, error) {
	return r.data.db.NovelConsume.Delete().Where(novelconsume.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelConsume 更新
func (r *novelConsumeRepo) UpdateNovelConsume(ctx context.Context, req *v1.NovelConsumeUpdateReq) (*ent.NovelConsume, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.NovelConsume.UpdateOneID(req.Id).
		SetNovelId(req.NovelId).
		SetCoin(req.Coin).
		SetCoupon(req.Coupon).
		SetDiscount(req.Discount).
		SetReward(req.Reward).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetNovelConsume 根据Id查询
func (r *novelConsumeRepo) GetNovelConsume(ctx context.Context, req *v1.NovelConsumeReq) (*ent.NovelConsume, error) {
	return r.data.db.NovelConsume.Get(ctx, req.Id)
}

// PageNovelConsume 分页查询
func (r *novelConsumeRepo) PageNovelConsume(ctx context.Context, req *v1.NovelConsumePageReq) ([]*ent.NovelConsume, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.NovelConsume.
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
func (r *novelConsumeRepo) genCondition(req *v1.NovelConsumeReq) []predicate.NovelConsume {
	if req == nil {
		return nil
	}
	list := make([]predicate.NovelConsume, 0)
	if req.Id > 0 {
		list = append(list, novelconsume.ID(req.Id))
	}
	if req.NovelId > 0 {
		list = append(list, novelconsume.NovelId(req.NovelId))
	}
	if req.Coin > 0 {
		list = append(list, novelconsume.Coin(req.Coin))
	}
	if req.Coupon > 0 {
		list = append(list, novelconsume.Coupon(req.Coupon))
	}
	if req.Discount > 0 {
		list = append(list, novelconsume.Discount(req.Discount))
	}
	if req.Reward > 0 {
		list = append(list, novelconsume.Reward(req.Reward))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novelconsume.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novelconsume.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novelconsume.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novelconsume.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novelconsume.TenantId(req.TenantId))
	}

	return list
}
