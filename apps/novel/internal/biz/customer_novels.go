package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/customernovels/v1"
	"hope/apps/novel/internal/data/ent"
)

type CustomerNovelsRepo interface {
	CreateCustomerNovels(context.Context, *v1.CustomerNovelsCreateReq) (*ent.CustomerNovels, error)
	DeleteCustomerNovels(context.Context, *v1.CustomerNovelsDeleteReq) error
	BatchDeleteCustomerNovels(context.Context, *v1.CustomerNovelsBatchDeleteReq) (int, error)
	UpdateCustomerNovels(context.Context, *v1.CustomerNovelsUpdateReq) (*ent.CustomerNovels, error)
	GetCustomerNovels(context.Context, *v1.CustomerNovelsReq) (*ent.CustomerNovels, error)
	PageCustomerNovels(context.Context, *v1.CustomerNovelsPageReq) ([]*ent.CustomerNovels, error)
}

type CustomerNovelsUseCase struct {
	repo CustomerNovelsRepo
	log  *log.Helper
}

func NewCustomerNovelsUseCase(repo CustomerNovelsRepo, logger log.Logger) *CustomerNovelsUseCase {
	return &CustomerNovelsUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CustomerNovelsUseCase) Create(ctx context.Context, req *v1.CustomerNovelsCreateReq) (*ent.CustomerNovels, error) {
	return uc.repo.CreateCustomerNovels(ctx, req)
}
func (uc *CustomerNovelsUseCase) Delete(ctx context.Context, req *v1.CustomerNovelsDeleteReq) error {
	return uc.repo.DeleteCustomerNovels(ctx, req)
}
func (uc *CustomerNovelsUseCase) BatchDelete(ctx context.Context, req *v1.CustomerNovelsBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteCustomerNovels(ctx, req)
}
func (uc *CustomerNovelsUseCase) Update(ctx context.Context, req *v1.CustomerNovelsUpdateReq) (*ent.CustomerNovels, error) {
	return uc.repo.UpdateCustomerNovels(ctx, req)
}
func (uc *CustomerNovelsUseCase) Get(ctx context.Context, req *v1.CustomerNovelsReq) (*ent.CustomerNovels, error) {
	return uc.repo.GetCustomerNovels(ctx, req)
}
func (uc *CustomerNovelsUseCase) Page(ctx context.Context, req *v1.CustomerNovelsPageReq) ([]*ent.CustomerNovels, error) {
	return uc.repo.PageCustomerNovels(ctx, req)
}
