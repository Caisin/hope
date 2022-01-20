package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/customernovelconfig/v1"
	"hope/apps/novel/internal/data/ent"
)

type CustomerNovelConfigRepo interface {
	CreateCustomerNovelConfig(context.Context, *v1.CustomerNovelConfigCreateReq) (*ent.CustomerNovelConfig, error)
	DeleteCustomerNovelConfig(context.Context, *v1.CustomerNovelConfigDeleteReq) error
	BatchDeleteCustomerNovelConfig(context.Context, *v1.CustomerNovelConfigBatchDeleteReq) (int, error)
	UpdateCustomerNovelConfig(context.Context, *v1.CustomerNovelConfigUpdateReq) (*ent.CustomerNovelConfig, error)
	GetCustomerNovelConfig(context.Context, *v1.CustomerNovelConfigReq) (*ent.CustomerNovelConfig, error)
	PageCustomerNovelConfig(context.Context, *v1.CustomerNovelConfigPageReq) ([]*ent.CustomerNovelConfig, error)
}

type CustomerNovelConfigUseCase struct {
	repo CustomerNovelConfigRepo
	log  *log.Helper
}

func NewCustomerNovelConfigUseCase(repo CustomerNovelConfigRepo, logger log.Logger) *CustomerNovelConfigUseCase {
	return &CustomerNovelConfigUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CustomerNovelConfigUseCase) Create(ctx context.Context, req *v1.CustomerNovelConfigCreateReq) (*ent.CustomerNovelConfig, error) {
	return uc.repo.CreateCustomerNovelConfig(ctx, req)
}
func (uc *CustomerNovelConfigUseCase) Delete(ctx context.Context, req *v1.CustomerNovelConfigDeleteReq) error {
	return uc.repo.DeleteCustomerNovelConfig(ctx, req)
}
func (uc *CustomerNovelConfigUseCase) BatchDelete(ctx context.Context, req *v1.CustomerNovelConfigBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteCustomerNovelConfig(ctx, req)
}
func (uc *CustomerNovelConfigUseCase) Update(ctx context.Context, req *v1.CustomerNovelConfigUpdateReq) (*ent.CustomerNovelConfig, error) {
	return uc.repo.UpdateCustomerNovelConfig(ctx, req)
}
func (uc *CustomerNovelConfigUseCase) Get(ctx context.Context, req *v1.CustomerNovelConfigReq) (*ent.CustomerNovelConfig, error) {
	return uc.repo.GetCustomerNovelConfig(ctx, req)
}
func (uc *CustomerNovelConfigUseCase) Page(ctx context.Context, req *v1.CustomerNovelConfigPageReq) ([]*ent.CustomerNovelConfig, error) {
	return uc.repo.PageCustomerNovelConfig(ctx, req)
}
