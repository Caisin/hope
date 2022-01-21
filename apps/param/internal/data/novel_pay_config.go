package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/novelpayconfig/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/novelpayconfig"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type novelPayConfigRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelPayConfigRepo .
func NewNovelPayConfigRepo(data *Data, logger log.Logger) biz.NovelPayConfigRepo {
	return &novelPayConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelPayConfig 创建
func (r *novelPayConfigRepo) CreateNovelPayConfig(ctx context.Context, req *v1.NovelPayConfigCreateReq) (*ent.NovelPayConfig, error) {
	now := time.Now()
	return r.data.db.NovelPayConfig.Create().
		SetProductId(req.ProductId).
		SetPaymentName(req.PaymentName).
		SetFirstPayment(req.FirstPayment).
		SetPayment(req.Payment).
		SetOriginalPrice(req.OriginalPrice).
		SetCfgType(req.CfgType).
		SetCoin(req.Coin).
		SetCurrency(req.Currency).
		SetCoupon(req.Coupon).
		SetCoinItem(req.CoinItem).
		SetCouponItem(req.CouponItem).
		SetSort(req.Sort).
		SetState(req.State).
		SetIsSend(req.IsSend).
		SetPayType(req.PayType).
		SetVipType(req.VipType).
		SetIsHot(req.IsHot).
		SetCycleDay(req.CycleDay).
		SetSummary(req.Summary).
		SetRemark(req.Remark).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteNovelPayConfig 删除
func (r *novelPayConfigRepo) DeleteNovelPayConfig(ctx context.Context, req *v1.NovelPayConfigDeleteReq) error {
	return r.data.db.NovelPayConfig.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelPayConfig 批量删除
func (r *novelPayConfigRepo) BatchDeleteNovelPayConfig(ctx context.Context, req *v1.NovelPayConfigBatchDeleteReq) (int, error) {
	return r.data.db.NovelPayConfig.Delete().Where(novelpayconfig.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelPayConfig 更新
func (r *novelPayConfigRepo) UpdateNovelPayConfig(ctx context.Context, req *v1.NovelPayConfigUpdateReq) (*ent.NovelPayConfig, error) {
	return r.data.db.NovelPayConfig.UpdateOne(convert.NovelPayConfigUpdateReq2Data(req)).Save(ctx)
}

// GetNovelPayConfig 根据Id查询
func (r *novelPayConfigRepo) GetNovelPayConfig(ctx context.Context, req *v1.NovelPayConfigReq) (*ent.NovelPayConfig, error) {
	return r.data.db.NovelPayConfig.Get(ctx, req.Id)
}

// PageNovelPayConfig 分页查询
func (r *novelPayConfigRepo) PageNovelPayConfig(ctx context.Context, req *v1.NovelPayConfigPageReq) ([]*ent.NovelPayConfig, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.NovelPayConfig.
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
func (r *novelPayConfigRepo) genCondition(req *v1.NovelPayConfigReq) []predicate.NovelPayConfig {
	if req == nil {
		return nil
	}
	list := make([]predicate.NovelPayConfig, 0)
	if req.Id > 0 {
		list = append(list, novelpayconfig.ID(req.Id))
	}
	if str.IsBlank(req.ProductId) {
		list = append(list, novelpayconfig.ProductIdContains(req.ProductId))
	}
	if str.IsBlank(req.PaymentName) {
		list = append(list, novelpayconfig.PaymentNameContains(req.PaymentName))
	}
	if req.FirstPayment > 0 {
		list = append(list, novelpayconfig.FirstPayment(req.FirstPayment))
	}
	if req.Payment > 0 {
		list = append(list, novelpayconfig.Payment(req.Payment))
	}
	if req.OriginalPrice > 0 {
		list = append(list, novelpayconfig.OriginalPrice(req.OriginalPrice))
	}
	if str.IsBlank(req.CfgType) {
		list = append(list, novelpayconfig.CfgTypeContains(req.CfgType))
	}
	if req.Coin > 0 {
		list = append(list, novelpayconfig.Coin(req.Coin))
	}
	if str.IsBlank(req.Currency) {
		list = append(list, novelpayconfig.CurrencyContains(req.Currency))
	}
	if req.Coupon > 0 {
		list = append(list, novelpayconfig.Coupon(req.Coupon))
	}
	if req.CoinItem > 0 {
		list = append(list, novelpayconfig.CoinItem(req.CoinItem))
	}
	if req.CouponItem > 0 {
		list = append(list, novelpayconfig.CouponItem(req.CouponItem))
	}
	if req.Sort > 0 {
		list = append(list, novelpayconfig.Sort(req.Sort))
	}
	if req.IsSend > 0 {
		list = append(list, novelpayconfig.IsSend(req.IsSend))
	}
	if req.PayType > 0 {
		list = append(list, novelpayconfig.PayType(req.PayType))
	}
	if req.VipType > 0 {
		list = append(list, novelpayconfig.VipType(req.VipType))
	}
	if req.CycleDay > 0 {
		list = append(list, novelpayconfig.CycleDay(req.CycleDay))
	}
	if str.IsBlank(req.Summary) {
		list = append(list, novelpayconfig.SummaryContains(req.Summary))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, novelpayconfig.RemarkContains(req.Remark))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, novelpayconfig.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, novelpayconfig.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novelpayconfig.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novelpayconfig.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novelpayconfig.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novelpayconfig.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novelpayconfig.TenantId(req.TenantId))
	}

	return list
}
