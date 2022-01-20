package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/payorder/v1"
	"hope/apps/novel/internal/data/ent"
)

type PayOrderRepo interface {
	CreatePayOrder(context.Context, *v1.PayOrderCreateReq) (*ent.PayOrder, error)
	DeletePayOrder(context.Context, *v1.PayOrderDeleteReq) error
	BatchDeletePayOrder(context.Context, *v1.PayOrderBatchDeleteReq) (int, error)
	UpdatePayOrder(context.Context, *v1.PayOrderUpdateReq) (*ent.PayOrder, error)
	GetPayOrder(context.Context, *v1.PayOrderReq) (*ent.PayOrder, error)
	PagePayOrder(context.Context, *v1.PayOrderPageReq) ([]*ent.PayOrder, error)
}

type PayOrderUseCase struct {
	repo PayOrderRepo
	log  *log.Helper
}

func NewPayOrderUseCase(repo PayOrderRepo, logger log.Logger) *PayOrderUseCase {
	return &PayOrderUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PayOrderUseCase) Create(ctx context.Context, req *v1.PayOrderCreateReq) (*ent.PayOrder, error) {
	return uc.repo.CreatePayOrder(ctx, req)
}
func (uc *PayOrderUseCase) Delete(ctx context.Context, req *v1.PayOrderDeleteReq) error {
	return uc.repo.DeletePayOrder(ctx, req)
}
func (uc *PayOrderUseCase) BatchDelete(ctx context.Context, req *v1.PayOrderBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeletePayOrder(ctx, req)
}
func (uc *PayOrderUseCase) Update(ctx context.Context, req *v1.PayOrderUpdateReq) (*ent.PayOrder, error) {
	return uc.repo.UpdatePayOrder(ctx, req)
}
func (uc *PayOrderUseCase) Get(ctx context.Context, req *v1.PayOrderReq) (*ent.PayOrder, error) {
	return uc.repo.GetPayOrder(ctx, req)
}
func (uc *PayOrderUseCase) Page(ctx context.Context, req *v1.PayOrderPageReq) ([]*ent.PayOrder, error) {
	return uc.repo.PagePayOrder(ctx, req)
}
