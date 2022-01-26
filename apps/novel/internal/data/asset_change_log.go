package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/assetchangelog/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/assetchangelog"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type assetChangeLogRepo struct {
	data *Data
	log  *log.Helper
}

// NewAssetChangeLogRepo .
func NewAssetChangeLogRepo(data *Data, logger log.Logger) biz.AssetChangeLogRepo {
	return &assetChangeLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateAssetChangeLog 创建
func (r *assetChangeLogRepo) CreateAssetChangeLog(ctx context.Context, req *v1.AssetChangeLogCreateReq) (*ent.AssetChangeLog, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.AssetChangeLog.Create().
		SetOrderId(req.OrderId).
		SetBalanceId(req.BalanceId).
		SetEventId(req.EventId).
		SetUserId(req.UserId).
		SetAssetItemId(req.AssetItemId).
		SetAmount(req.Amount).
		SetOldBalance(req.OldBalance).
		SetNewBalance(req.NewBalance).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteAssetChangeLog 删除
func (r *assetChangeLogRepo) DeleteAssetChangeLog(ctx context.Context, req *v1.AssetChangeLogDeleteReq) error {
	return r.data.db.AssetChangeLog.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteAssetChangeLog 批量删除
func (r *assetChangeLogRepo) BatchDeleteAssetChangeLog(ctx context.Context, req *v1.AssetChangeLogBatchDeleteReq) (int, error) {
	return r.data.db.AssetChangeLog.Delete().Where(assetchangelog.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateAssetChangeLog 更新
func (r *assetChangeLogRepo) UpdateAssetChangeLog(ctx context.Context, req *v1.AssetChangeLogUpdateReq) (*ent.AssetChangeLog, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.AssetChangeLog.UpdateOneID(req.Id).
		SetOrderId(req.OrderId).
		SetBalanceId(req.BalanceId).
		SetEventId(req.EventId).
		SetUserId(req.UserId).
		SetAssetItemId(req.AssetItemId).
		SetAmount(req.Amount).
		SetOldBalance(req.OldBalance).
		SetNewBalance(req.NewBalance).
		SetRemark(req.Remark).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetAssetChangeLog 根据Id查询
func (r *assetChangeLogRepo) GetAssetChangeLog(ctx context.Context, req *v1.AssetChangeLogReq) (*ent.AssetChangeLog, error) {
	return r.data.db.AssetChangeLog.Get(ctx, req.Id)
}

// PageAssetChangeLog 分页查询
func (r *assetChangeLogRepo) PageAssetChangeLog(ctx context.Context, req *v1.AssetChangeLogPageReq) ([]*ent.AssetChangeLog, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.AssetChangeLog.
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
func (r *assetChangeLogRepo) genCondition(req *v1.AssetChangeLogReq) []predicate.AssetChangeLog {
	if req == nil {
		return nil
	}
	list := make([]predicate.AssetChangeLog, 0)
	if req.Id > 0 {
		list = append(list, assetchangelog.ID(req.Id))
	}
	if str.IsBlank(req.OrderId) {
		list = append(list, assetchangelog.OrderIdContains(req.OrderId))
	}
	if req.BalanceId > 0 {
		list = append(list, assetchangelog.BalanceId(req.BalanceId))
	}
	if req.EventId > 0 {
		list = append(list, assetchangelog.EventId(req.EventId))
	}
	if req.UserId > 0 {
		list = append(list, assetchangelog.UserId(req.UserId))
	}
	if req.AssetItemId > 0 {
		list = append(list, assetchangelog.AssetItemId(req.AssetItemId))
	}
	if req.Amount > 0 {
		list = append(list, assetchangelog.Amount(req.Amount))
	}
	if req.OldBalance > 0 {
		list = append(list, assetchangelog.OldBalance(req.OldBalance))
	}
	if req.NewBalance > 0 {
		list = append(list, assetchangelog.NewBalance(req.NewBalance))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, assetchangelog.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, assetchangelog.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, assetchangelog.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, assetchangelog.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, assetchangelog.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, assetchangelog.TenantId(req.TenantId))
	}

	return list
}
