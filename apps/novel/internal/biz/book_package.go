package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/bookpackage/v1"
	"hope/apps/novel/internal/data/ent"
)

type BookPackageRepo interface {
	CreateBookPackage(context.Context, *v1.BookPackageCreateReq) (*ent.BookPackage, error)
	DeleteBookPackage(context.Context, *v1.BookPackageDeleteReq) error
	BatchDeleteBookPackage(context.Context, *v1.BookPackageBatchDeleteReq) (int, error)
	UpdateBookPackage(context.Context, *v1.BookPackageUpdateReq) (*ent.BookPackage, error)
	GetBookPackage(context.Context, *v1.BookPackageReq) (*ent.BookPackage, error)
	PageBookPackage(context.Context, *v1.BookPackagePageReq) ([]*ent.BookPackage, error)
}

type BookPackageUseCase struct {
	repo BookPackageRepo
	log  *log.Helper
}

func NewBookPackageUseCase(repo BookPackageRepo, logger log.Logger) *BookPackageUseCase {
	return &BookPackageUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BookPackageUseCase) Create(ctx context.Context, req *v1.BookPackageCreateReq) (*ent.BookPackage, error) {
	return uc.repo.CreateBookPackage(ctx, req)
}
func (uc *BookPackageUseCase) Delete(ctx context.Context, req *v1.BookPackageDeleteReq) error {
	return uc.repo.DeleteBookPackage(ctx, req)
}
func (uc *BookPackageUseCase) BatchDelete(ctx context.Context, req *v1.BookPackageBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteBookPackage(ctx, req)
}
func (uc *BookPackageUseCase) Update(ctx context.Context, req *v1.BookPackageUpdateReq) (*ent.BookPackage, error) {
	return uc.repo.UpdateBookPackage(ctx, req)
}
func (uc *BookPackageUseCase) Get(ctx context.Context, req *v1.BookPackageReq) (*ent.BookPackage, error) {
	return uc.repo.GetBookPackage(ctx, req)
}
func (uc *BookPackageUseCase) Page(ctx context.Context, req *v1.BookPackagePageReq) ([]*ent.BookPackage, error) {
	return uc.repo.PageBookPackage(ctx, req)
}
