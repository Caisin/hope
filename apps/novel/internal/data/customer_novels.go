package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/customernovels/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/customernovels"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type customerNovelsRepo struct {
	data *Data
	log  *log.Helper
}

// NewCustomerNovelsRepo .
func NewCustomerNovelsRepo(data *Data, logger log.Logger) biz.CustomerNovelsRepo {
	return &customerNovelsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateCustomerNovels 创建
func (r *customerNovelsRepo) CreateCustomerNovels(ctx context.Context, req *v1.CustomerNovelsCreateReq) (*ent.CustomerNovels, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.CustomerNovels.Create().
		SetNovelId(req.NovelId).
		SetTypeId(req.TypeId).
		SetTypeCode(req.TypeCode).
		SetGroupCode(req.GroupCode).
		SetFieldName(req.FieldName).
		SetCover(req.Cover).
		SetOrderNum(req.OrderNum).
		SetRemark(req.Remark).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteCustomerNovels 删除
func (r *customerNovelsRepo) DeleteCustomerNovels(ctx context.Context, req *v1.CustomerNovelsDeleteReq) error {
	return r.data.db.CustomerNovels.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteCustomerNovels 批量删除
func (r *customerNovelsRepo) BatchDeleteCustomerNovels(ctx context.Context, req *v1.CustomerNovelsBatchDeleteReq) (int, error) {
	return r.data.db.CustomerNovels.Delete().Where(customernovels.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateCustomerNovels 更新
func (r *customerNovelsRepo) UpdateCustomerNovels(ctx context.Context, req *v1.CustomerNovelsUpdateReq) (*ent.CustomerNovels, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.CustomerNovels.UpdateOneID(req.Id).
		SetNovelId(req.NovelId).
		SetTypeId(req.TypeId).
		SetTypeCode(req.TypeCode).
		SetGroupCode(req.GroupCode).
		SetFieldName(req.FieldName).
		SetCover(req.Cover).
		SetOrderNum(req.OrderNum).
		SetRemark(req.Remark).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetCustomerNovels 根据Id查询
func (r *customerNovelsRepo) GetCustomerNovels(ctx context.Context, req *v1.CustomerNovelsReq) (*ent.CustomerNovels, error) {
	return r.data.db.CustomerNovels.Get(ctx, req.Id)
}

// PageCustomerNovels 分页查询
func (r *customerNovelsRepo) PageCustomerNovels(ctx context.Context, req *v1.CustomerNovelsPageReq) ([]*ent.CustomerNovels, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.CustomerNovels.
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
func (r *customerNovelsRepo) genCondition(req *v1.CustomerNovelsReq) []predicate.CustomerNovels {
	if req == nil {
		return nil
	}
	list := make([]predicate.CustomerNovels, 0)
	if req.Id > 0 {
		list = append(list, customernovels.ID(req.Id))
	}
	if req.NovelId > 0 {
		list = append(list, customernovels.NovelId(req.NovelId))
	}
	if req.TypeId > 0 {
		list = append(list, customernovels.TypeId(req.TypeId))
	}
	if str.IsBlank(req.TypeCode) {
		list = append(list, customernovels.TypeCodeContains(req.TypeCode))
	}
	if str.IsBlank(req.GroupCode) {
		list = append(list, customernovels.GroupCodeContains(req.GroupCode))
	}
	if str.IsBlank(req.FieldName) {
		list = append(list, customernovels.FieldNameContains(req.FieldName))
	}
	if str.IsBlank(req.Cover) {
		list = append(list, customernovels.CoverContains(req.Cover))
	}
	if req.OrderNum > 0 {
		list = append(list, customernovels.OrderNum(req.OrderNum))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, customernovels.RemarkContains(req.Remark))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, customernovels.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, customernovels.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, customernovels.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, customernovels.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, customernovels.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, customernovels.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, customernovels.TenantId(req.TenantId))
	}

	return list
}
