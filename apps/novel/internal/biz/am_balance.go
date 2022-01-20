package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/ambalance/v1"
	"hope/apps/novel/internal/data/ent"
)

type AmBalanceRepo interface {
	CreateAmBalance(context.Context, *v1.AmBalanceCreateReq) (*ent.AmBalance, error)
	DeleteAmBalance(context.Context, *v1.AmBalanceDeleteReq) error
	BatchDeleteAmBalance(context.Context, *v1.AmBalanceBatchDeleteReq) (int, error)
	UpdateAmBalance(context.Context, *v1.AmBalanceUpdateReq) (*ent.AmBalance, error)
	GetAmBalance(context.Context, *v1.AmBalanceReq) (*ent.AmBalance, error)
	PageAmBalance(context.Context, *v1.AmBalancePageReq) ([]*ent.AmBalance, error)
}

type AmBalanceUseCase struct {
	repo AmBalanceRepo
	log  *log.Helper
}

func NewAmBalanceUseCase(repo AmBalanceRepo, logger log.Logger) *AmBalanceUseCase {
	return &AmBalanceUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AmBalanceUseCase) Create(ctx context.Context, req *v1.AmBalanceCreateReq) (*ent.AmBalance, error) {
	return uc.repo.CreateAmBalance(ctx, req)
}
func (uc *AmBalanceUseCase) Delete(ctx context.Context, req *v1.AmBalanceDeleteReq) error {
	return uc.repo.DeleteAmBalance(ctx, req)
}
func (uc *AmBalanceUseCase) BatchDelete(ctx context.Context, req *v1.AmBalanceBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteAmBalance(ctx, req)
}
func (uc *AmBalanceUseCase) Update(ctx context.Context, req *v1.AmBalanceUpdateReq) (*ent.AmBalance, error) {
	return uc.repo.UpdateAmBalance(ctx, req)
}
func (uc *AmBalanceUseCase) Get(ctx context.Context, req *v1.AmBalanceReq) (*ent.AmBalance, error) {
	return uc.repo.GetAmBalance(ctx, req)
}
func (uc *AmBalanceUseCase) Page(ctx context.Context, req *v1.AmBalancePageReq) ([]*ent.AmBalance, error) {
	return uc.repo.PageAmBalance(ctx, req)
}
