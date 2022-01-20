package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/casbinrule/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/casbinrule"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/pkg/util/str"
	"time"
)

type casbinRuleRepo struct {
	data *Data
	log  *log.Helper
}

// NewCasbinRuleRepo .
func NewCasbinRuleRepo(data *Data, logger log.Logger) biz.CasbinRuleRepo {
	return &casbinRuleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateCasbinRule 创建
func (r *casbinRuleRepo) CreateCasbinRule(ctx context.Context, req *v1.CasbinRuleCreateReq) (*ent.CasbinRule, error) {
	now := time.Now()
	return r.data.db.CasbinRule.Create().
		SetPType(req.PType).
		SetV0(req.V0).
		SetV1(req.V1).
		SetV2(req.V2).
		SetV3(req.V3).
		SetV4(req.V4).
		SetV5(req.V5).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteCasbinRule 删除
func (r *casbinRuleRepo) DeleteCasbinRule(ctx context.Context, req *v1.CasbinRuleDeleteReq) error {
	return r.data.db.CasbinRule.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteCasbinRule 批量删除
func (r *casbinRuleRepo) BatchDeleteCasbinRule(ctx context.Context, req *v1.CasbinRuleBatchDeleteReq) (int, error) {
	return r.data.db.CasbinRule.Delete().Where(casbinrule.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateCasbinRule 更新
func (r *casbinRuleRepo) UpdateCasbinRule(ctx context.Context, req *v1.CasbinRuleUpdateReq) (*ent.CasbinRule, error) {
	return r.data.db.CasbinRule.UpdateOne(convert.CasbinRuleUpdateReq2Data(req)).Save(ctx)
}

// GetCasbinRule 根据Id查询
func (r *casbinRuleRepo) GetCasbinRule(ctx context.Context, req *v1.CasbinRuleReq) (*ent.CasbinRule, error) {
	return r.data.db.CasbinRule.Get(ctx, req.Id)
}

// PageCasbinRule 分页查询
func (r *casbinRuleRepo) PageCasbinRule(ctx context.Context, req *v1.CasbinRulePageReq) ([]*ent.CasbinRule, error) {
	pagin := req.Pagin
	query := r.data.db.CasbinRule.
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
func (r *casbinRuleRepo) genCondition(req *v1.CasbinRuleReq) []predicate.CasbinRule {
	list := make([]predicate.CasbinRule, 0)
	if req.Id > 0 {
		list = append(list, casbinrule.ID(req.Id))
	}
	if str.IsBlank(req.PType) {
		list = append(list, casbinrule.PTypeContains(req.PType))
	}
	if str.IsBlank(req.V0) {
		list = append(list, casbinrule.V0Contains(req.V0))
	}
	if str.IsBlank(req.V1) {
		list = append(list, casbinrule.V1Contains(req.V1))
	}
	if str.IsBlank(req.V2) {
		list = append(list, casbinrule.V2Contains(req.V2))
	}
	if str.IsBlank(req.V3) {
		list = append(list, casbinrule.V3Contains(req.V3))
	}
	if str.IsBlank(req.V4) {
		list = append(list, casbinrule.V4Contains(req.V4))
	}
	if str.IsBlank(req.V5) {
		list = append(list, casbinrule.V5Contains(req.V5))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, casbinrule.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, casbinrule.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, casbinrule.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, casbinrule.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, casbinrule.TenantId(req.TenantId))
	}

	return list
}
