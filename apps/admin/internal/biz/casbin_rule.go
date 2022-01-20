package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/casbinrule/v1"
	"hope/apps/admin/internal/data/ent"
)

type CasbinRuleRepo interface {
	CreateCasbinRule(context.Context, *v1.CasbinRuleCreateReq) (*ent.CasbinRule, error)
	DeleteCasbinRule(context.Context, *v1.CasbinRuleDeleteReq) error
	BatchDeleteCasbinRule(context.Context, *v1.CasbinRuleBatchDeleteReq) (int, error)
	UpdateCasbinRule(context.Context, *v1.CasbinRuleUpdateReq) (*ent.CasbinRule, error)
	GetCasbinRule(context.Context, *v1.CasbinRuleReq) (*ent.CasbinRule, error)
	PageCasbinRule(context.Context, *v1.CasbinRulePageReq) ([]*ent.CasbinRule, error)
}

type CasbinRuleUseCase struct {
	repo CasbinRuleRepo
	log  *log.Helper
}

func NewCasbinRuleUseCase(repo CasbinRuleRepo, logger log.Logger) *CasbinRuleUseCase {
	return &CasbinRuleUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CasbinRuleUseCase) Create(ctx context.Context, req *v1.CasbinRuleCreateReq) (*ent.CasbinRule, error) {
	return uc.repo.CreateCasbinRule(ctx, req)
}
func (uc *CasbinRuleUseCase) Delete(ctx context.Context, req *v1.CasbinRuleDeleteReq) error {
	return uc.repo.DeleteCasbinRule(ctx, req)
}
func (uc *CasbinRuleUseCase) BatchDelete(ctx context.Context, req *v1.CasbinRuleBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteCasbinRule(ctx, req)
}
func (uc *CasbinRuleUseCase) Update(ctx context.Context, req *v1.CasbinRuleUpdateReq) (*ent.CasbinRule, error) {
	return uc.repo.UpdateCasbinRule(ctx, req)
}
func (uc *CasbinRuleUseCase) Get(ctx context.Context, req *v1.CasbinRuleReq) (*ent.CasbinRule, error) {
	return uc.repo.GetCasbinRule(ctx, req)
}
func (uc *CasbinRuleUseCase) Page(ctx context.Context, req *v1.CasbinRulePageReq) ([]*ent.CasbinRule, error) {
	return uc.repo.PageCasbinRule(ctx, req)
}
