package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/userevent/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/userevent"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type userEventRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserEventRepo .
func NewUserEventRepo(data *Data, logger log.Logger) biz.UserEventRepo {
	return &userEventRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUserEvent 创建
func (r *userEventRepo) CreateUserEvent(ctx context.Context, req *v1.UserEventCreateReq) (*ent.UserEvent, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.UserEvent.Create().
		SetUserId(req.UserId).
		SetEventType(req.EventType).
		SetNovelId(req.NovelId).
		SetChapterId(req.ChapterId).
		SetCoin(req.Coin).
		SetCoupon(req.Coupon).
		SetMoney(req.Money).
		SetKeyword(req.Keyword).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteUserEvent 删除
func (r *userEventRepo) DeleteUserEvent(ctx context.Context, req *v1.UserEventDeleteReq) error {
	return r.data.db.UserEvent.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteUserEvent 批量删除
func (r *userEventRepo) BatchDeleteUserEvent(ctx context.Context, req *v1.UserEventBatchDeleteReq) (int, error) {
	return r.data.db.UserEvent.Delete().Where(userevent.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateUserEvent 更新
func (r *userEventRepo) UpdateUserEvent(ctx context.Context, req *v1.UserEventUpdateReq) (*ent.UserEvent, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	data := convert.UserEventUpdateReq2Data(req)
	data.UpdateBy = claims.UserId
	return r.data.db.UserEvent.UpdateOne(data).Save(ctx)
}

// GetUserEvent 根据Id查询
func (r *userEventRepo) GetUserEvent(ctx context.Context, req *v1.UserEventReq) (*ent.UserEvent, error) {
	return r.data.db.UserEvent.Get(ctx, req.Id)
}

// PageUserEvent 分页查询
func (r *userEventRepo) PageUserEvent(ctx context.Context, req *v1.UserEventPageReq) ([]*ent.UserEvent, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.UserEvent.
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
func (r *userEventRepo) genCondition(req *v1.UserEventReq) []predicate.UserEvent {
	if req == nil {
		return nil
	}
	list := make([]predicate.UserEvent, 0)
	if req.Id > 0 {
		list = append(list, userevent.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, userevent.UserId(req.UserId))
	}
	if str.IsBlank(req.EventType) {
		list = append(list, userevent.EventTypeContains(req.EventType))
	}
	if req.NovelId > 0 {
		list = append(list, userevent.NovelId(req.NovelId))
	}
	if req.ChapterId > 0 {
		list = append(list, userevent.ChapterId(req.ChapterId))
	}
	if req.Coin > 0 {
		list = append(list, userevent.Coin(req.Coin))
	}
	if req.Coupon > 0 {
		list = append(list, userevent.Coupon(req.Coupon))
	}
	if req.Money > 0 {
		list = append(list, userevent.Money(req.Money))
	}
	if str.IsBlank(req.Keyword) {
		list = append(list, userevent.KeywordContains(req.Keyword))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, userevent.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, userevent.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, userevent.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, userevent.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, userevent.TenantId(req.TenantId))
	}

	return list
}
