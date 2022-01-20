package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/payorder/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/payorder"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/schema"
	"hope/pkg/util/str"
	"time"
)

type payOrderRepo struct {
	data *Data
	log  *log.Helper
}

// NewPayOrderRepo .
func NewPayOrderRepo(data *Data, logger log.Logger) biz.PayOrderRepo {
	return &payOrderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreatePayOrder 创建
func (r *payOrderRepo) CreatePayOrder(ctx context.Context, req *v1.PayOrderCreateReq) (*ent.PayOrder, error) {
	now := time.Now()
	return r.data.db.PayOrder.Create().
		SetOrderId(req.OrderId).
		SetLastRead(req.LastRead).
		SetLastChapter(req.LastChapter).
		SetPaymentName(req.PaymentName).
		SetPaymentId(req.PaymentId).
		SetState(schema.OrderState(req.State)).
		SetPayment(req.Payment).
		SetPaymentTime(req.PaymentTime.AsTime()).
		SetCloseTime(req.CloseTime.AsTime()).
		SetPayType(payorder.PayType(req.PayType)).
		SetCoin(req.Coin).
		SetCoupon(req.Coupon).
		SetVipDays(req.VipDays).
		SetVipType(req.VipType).
		SetVipName(req.VipName).
		SetTimes(req.Times).
		SetOtherOrderId(req.OtherOrderId).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeletePayOrder 删除
func (r *payOrderRepo) DeletePayOrder(ctx context.Context, req *v1.PayOrderDeleteReq) error {
	return r.data.db.PayOrder.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeletePayOrder 批量删除
func (r *payOrderRepo) BatchDeletePayOrder(ctx context.Context, req *v1.PayOrderBatchDeleteReq) (int, error) {
	return r.data.db.PayOrder.Delete().Where(payorder.IDIn(req.Ids...)).Exec(ctx)
}

// UpdatePayOrder 更新
func (r *payOrderRepo) UpdatePayOrder(ctx context.Context, req *v1.PayOrderUpdateReq) (*ent.PayOrder, error) {
	return r.data.db.PayOrder.UpdateOne(convert.PayOrderUpdateReq2Data(req)).Save(ctx)
}

// GetPayOrder 根据Id查询
func (r *payOrderRepo) GetPayOrder(ctx context.Context, req *v1.PayOrderReq) (*ent.PayOrder, error) {
	return r.data.db.PayOrder.Get(ctx, req.Id)
}

// PagePayOrder 分页查询
func (r *payOrderRepo) PagePayOrder(ctx context.Context, req *v1.PayOrderPageReq) ([]*ent.PayOrder, error) {
	pagin := req.Pagin
	query := r.data.db.PayOrder.
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
func (r *payOrderRepo) genCondition(req *v1.PayOrderReq) []predicate.PayOrder {
	list := make([]predicate.PayOrder, 0)
	if req.Id > 0 {
		list = append(list, payorder.ID(req.Id))
	}
	if str.IsBlank(req.OrderId) {
		list = append(list, payorder.OrderIdContains(req.OrderId))
	}
	if str.IsBlank(req.LastRead) {
		list = append(list, payorder.LastReadContains(req.LastRead))
	}
	if str.IsBlank(req.LastChapter) {
		list = append(list, payorder.LastChapterContains(req.LastChapter))
	}
	if str.IsBlank(req.PaymentName) {
		list = append(list, payorder.PaymentNameContains(req.PaymentName))
	}
	if str.IsBlank(req.PaymentId) {
		list = append(list, payorder.PaymentIdContains(req.PaymentId))
	}
	if req.State > 0 {
		list = append(list, payorder.State(schema.OrderState(req.State)))
	}
	if req.Payment > 0 {
		list = append(list, payorder.Payment(req.Payment))
	}
	if req.PaymentTime.IsValid() && !req.PaymentTime.AsTime().IsZero() {
		list = append(list, payorder.PaymentTimeGTE(req.PaymentTime.AsTime()))
	}
	if req.CloseTime.IsValid() && !req.CloseTime.AsTime().IsZero() {
		list = append(list, payorder.CloseTimeGTE(req.CloseTime.AsTime()))
	}
	payType := payorder.PayType(req.PayType)
	if payorder.PayTypeValidator(payType) == nil {
		list = append(list, payorder.PayTypeEQ(payType))
	}
	if req.Coin > 0 {
		list = append(list, payorder.Coin(req.Coin))
	}
	if req.Coupon > 0 {
		list = append(list, payorder.Coupon(req.Coupon))
	}
	if str.IsBlank(req.VipDays) {
		list = append(list, payorder.VipDaysContains(req.VipDays))
	}
	if req.VipType > 0 {
		list = append(list, payorder.VipType(req.VipType))
	}
	if str.IsBlank(req.VipName) {
		list = append(list, payorder.VipNameContains(req.VipName))
	}
	if req.Times > 0 {
		list = append(list, payorder.Times(req.Times))
	}
	if str.IsBlank(req.OtherOrderId) {
		list = append(list, payorder.OtherOrderIdContains(req.OtherOrderId))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, payorder.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, payorder.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, payorder.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, payorder.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, payorder.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, payorder.TenantId(req.TenantId))
	}

	return list
}
