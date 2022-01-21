package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/scoreproduct/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/scoreproduct"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type scoreProductRepo struct {
	data *Data
	log  *log.Helper
}

// NewScoreProductRepo .
func NewScoreProductRepo(data *Data, logger log.Logger) biz.ScoreProductRepo {
	return &scoreProductRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateScoreProduct 创建
func (r *scoreProductRepo) CreateScoreProduct(ctx context.Context, req *v1.ScoreProductCreateReq) (*ent.ScoreProduct, error) {
	now := time.Now()
	return r.data.db.ScoreProduct.Create().
		SetProductName(req.ProductName).
		SetSummary(req.Summary).
		SetCardUrl(req.CardUrl).
		SetScore(req.Score).
		SetVipType(req.VipType).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteScoreProduct 删除
func (r *scoreProductRepo) DeleteScoreProduct(ctx context.Context, req *v1.ScoreProductDeleteReq) error {
	return r.data.db.ScoreProduct.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteScoreProduct 批量删除
func (r *scoreProductRepo) BatchDeleteScoreProduct(ctx context.Context, req *v1.ScoreProductBatchDeleteReq) (int, error) {
	return r.data.db.ScoreProduct.Delete().Where(scoreproduct.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateScoreProduct 更新
func (r *scoreProductRepo) UpdateScoreProduct(ctx context.Context, req *v1.ScoreProductUpdateReq) (*ent.ScoreProduct, error) {
	return r.data.db.ScoreProduct.UpdateOne(convert.ScoreProductUpdateReq2Data(req)).Save(ctx)
}

// GetScoreProduct 根据Id查询
func (r *scoreProductRepo) GetScoreProduct(ctx context.Context, req *v1.ScoreProductReq) (*ent.ScoreProduct, error) {
	return r.data.db.ScoreProduct.Get(ctx, req.Id)
}

// PageScoreProduct 分页查询
func (r *scoreProductRepo) PageScoreProduct(ctx context.Context, req *v1.ScoreProductPageReq) ([]*ent.ScoreProduct, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.ScoreProduct.
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
func (r *scoreProductRepo) genCondition(req *v1.ScoreProductReq) []predicate.ScoreProduct {
	if req == nil {
		return nil
	}
	list := make([]predicate.ScoreProduct, 0)
	if req.Id > 0 {
		list = append(list, scoreproduct.ID(req.Id))
	}
	if str.IsBlank(req.ProductName) {
		list = append(list, scoreproduct.ProductNameContains(req.ProductName))
	}
	if str.IsBlank(req.Summary) {
		list = append(list, scoreproduct.SummaryContains(req.Summary))
	}
	if str.IsBlank(req.CardUrl) {
		list = append(list, scoreproduct.CardUrlContains(req.CardUrl))
	}
	if req.Score > 0 {
		list = append(list, scoreproduct.Score(req.Score))
	}
	if req.VipType > 0 {
		list = append(list, scoreproduct.VipType(req.VipType))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, scoreproduct.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, scoreproduct.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, scoreproduct.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, scoreproduct.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, scoreproduct.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, scoreproduct.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, scoreproduct.TenantId(req.TenantId))
	}

	return list
}
