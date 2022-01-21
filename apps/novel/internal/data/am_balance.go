package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/ambalance/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/ambalance"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type amBalanceRepo struct {
	data *Data
	log  *log.Helper
}

// NewAmBalanceRepo .
func NewAmBalanceRepo(data *Data, logger log.Logger) biz.AmBalanceRepo {
	return &amBalanceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateAmBalance 创建
func (r *amBalanceRepo) CreateAmBalance(ctx context.Context, req *v1.AmBalanceCreateReq) (*ent.AmBalance, error) {
	now := time.Now()
	return r.data.db.AmBalance.Create().
		SetUserId(req.UserId).
		SetOrderId(req.OrderId).
		SetEventId(req.EventId).
		SetCashTag(req.CashTag).
		SetAssetItemId(req.AssetItemId).
		SetAmount(req.Amount).
		SetBalance(req.Balance).
		SetRemark(req.Remark).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteAmBalance 删除
func (r *amBalanceRepo) DeleteAmBalance(ctx context.Context, req *v1.AmBalanceDeleteReq) error {
	return r.data.db.AmBalance.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteAmBalance 批量删除
func (r *amBalanceRepo) BatchDeleteAmBalance(ctx context.Context, req *v1.AmBalanceBatchDeleteReq) (int, error) {
	return r.data.db.AmBalance.Delete().Where(ambalance.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateAmBalance 更新
func (r *amBalanceRepo) UpdateAmBalance(ctx context.Context, req *v1.AmBalanceUpdateReq) (*ent.AmBalance, error) {
	return r.data.db.AmBalance.UpdateOne(convert.AmBalanceUpdateReq2Data(req)).Save(ctx)
}

// GetAmBalance 根据Id查询
func (r *amBalanceRepo) GetAmBalance(ctx context.Context, req *v1.AmBalanceReq) (*ent.AmBalance, error) {
	return r.data.db.AmBalance.Get(ctx, req.Id)
}

// PageAmBalance 分页查询
func (r *amBalanceRepo) PageAmBalance(ctx context.Context, req *v1.AmBalancePageReq) ([]*ent.AmBalance, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.AmBalance.
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
	query.Limit(int(p.GetPage())).
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
func (r *amBalanceRepo) genCondition(req *v1.AmBalanceReq) []predicate.AmBalance {
	if req == nil {
		return nil
	}
	list := make([]predicate.AmBalance, 0)
	if req.Id > 0 {
		list = append(list, ambalance.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, ambalance.UserId(req.UserId))
	}
	if str.IsBlank(req.OrderId) {
		list = append(list, ambalance.OrderIdContains(req.OrderId))
	}
	if req.EventId > 0 {
		list = append(list, ambalance.EventId(req.EventId))
	}
	if req.CashTag > 0 {
		list = append(list, ambalance.CashTag(req.CashTag))
	}
	if req.AssetItemId > 0 {
		list = append(list, ambalance.AssetItemId(req.AssetItemId))
	}
	if req.Amount > 0 {
		list = append(list, ambalance.Amount(req.Amount))
	}
	if req.Balance > 0 {
		list = append(list, ambalance.Balance(req.Balance))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, ambalance.RemarkContains(req.Remark))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, ambalance.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, ambalance.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, ambalance.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, ambalance.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, ambalance.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, ambalance.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, ambalance.TenantId(req.TenantId))
	}

	return list
}
