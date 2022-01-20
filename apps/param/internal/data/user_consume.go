package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/userconsume/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/userconsume"
	"time"
)

type userConsumeRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserConsumeRepo .
func NewUserConsumeRepo(data *Data, logger log.Logger) biz.UserConsumeRepo {
	return &userConsumeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUserConsume 创建
func (r *userConsumeRepo) CreateUserConsume(ctx context.Context, req *v1.UserConsumeCreateReq) (*ent.UserConsume, error) {
	now := time.Now()
	return r.data.db.UserConsume.Create().
		SetNovelId(req.NovelId).
		SetCoin(req.Coin).
		SetCoupon(req.Coupon).
		SetDiscount(req.Discount).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteUserConsume 删除
func (r *userConsumeRepo) DeleteUserConsume(ctx context.Context, req *v1.UserConsumeDeleteReq) error {
	return r.data.db.UserConsume.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteUserConsume 批量删除
func (r *userConsumeRepo) BatchDeleteUserConsume(ctx context.Context, req *v1.UserConsumeBatchDeleteReq) (int, error) {
	return r.data.db.UserConsume.Delete().Where(userconsume.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateUserConsume 更新
func (r *userConsumeRepo) UpdateUserConsume(ctx context.Context, req *v1.UserConsumeUpdateReq) (*ent.UserConsume, error) {
	return r.data.db.UserConsume.UpdateOne(convert.UserConsumeUpdateReq2Data(req)).Save(ctx)
}

// GetUserConsume 根据Id查询
func (r *userConsumeRepo) GetUserConsume(ctx context.Context, req *v1.UserConsumeReq) (*ent.UserConsume, error) {
	return r.data.db.UserConsume.Get(ctx, req.Id)
}

// PageUserConsume 分页查询
func (r *userConsumeRepo) PageUserConsume(ctx context.Context, req *v1.UserConsumePageReq) ([]*ent.UserConsume, error) {
	pagin := req.Pagin
	query := r.data.db.UserConsume.
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
func (r *userConsumeRepo) genCondition(req *v1.UserConsumeReq) []predicate.UserConsume {
	list := make([]predicate.UserConsume, 0)
	if req.Id > 0 {
		list = append(list, userconsume.ID(req.Id))
	}
	if req.NovelId > 0 {
		list = append(list, userconsume.NovelId(req.NovelId))
	}
	if req.Coin > 0 {
		list = append(list, userconsume.Coin(req.Coin))
	}
	if req.Coupon > 0 {
		list = append(list, userconsume.Coupon(req.Coupon))
	}
	if req.Discount > 0 {
		list = append(list, userconsume.Discount(req.Discount))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, userconsume.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, userconsume.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, userconsume.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, userconsume.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, userconsume.TenantId(req.TenantId))
	}

	return list
}
