package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/resourcestorage/v1"
	"hope/apps/param/internal/data/ent"
)

type ResourceStorageRepo interface {
	CreateResourceStorage(context.Context, *v1.ResourceStorageCreateReq) (*ent.ResourceStorage, error)
	DeleteResourceStorage(context.Context, *v1.ResourceStorageDeleteReq) error
	BatchDeleteResourceStorage(context.Context, *v1.ResourceStorageBatchDeleteReq) (int, error)
	UpdateResourceStorage(context.Context, *v1.ResourceStorageUpdateReq) (*ent.ResourceStorage, error)
	GetResourceStorage(context.Context, *v1.ResourceStorageReq) (*ent.ResourceStorage, error)
	PageResourceStorage(context.Context, *v1.ResourceStoragePageReq) ([]*ent.ResourceStorage, error)
}

type ResourceStorageUseCase struct {
	repo ResourceStorageRepo
	log  *log.Helper
}

func NewResourceStorageUseCase(repo ResourceStorageRepo, logger log.Logger) *ResourceStorageUseCase {
	return &ResourceStorageUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ResourceStorageUseCase) Create(ctx context.Context, req *v1.ResourceStorageCreateReq) (*ent.ResourceStorage, error) {
	return uc.repo.CreateResourceStorage(ctx, req)
}
func (uc *ResourceStorageUseCase) Delete(ctx context.Context, req *v1.ResourceStorageDeleteReq) error {
	return uc.repo.DeleteResourceStorage(ctx, req)
}
func (uc *ResourceStorageUseCase) BatchDelete(ctx context.Context, req *v1.ResourceStorageBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteResourceStorage(ctx, req)
}
func (uc *ResourceStorageUseCase) Update(ctx context.Context, req *v1.ResourceStorageUpdateReq) (*ent.ResourceStorage, error) {
	return uc.repo.UpdateResourceStorage(ctx, req)
}
func (uc *ResourceStorageUseCase) Get(ctx context.Context, req *v1.ResourceStorageReq) (*ent.ResourceStorage, error) {
	return uc.repo.GetResourceStorage(ctx, req)
}
func (uc *ResourceStorageUseCase) Page(ctx context.Context, req *v1.ResourceStoragePageReq) ([]*ent.ResourceStorage, error) {
	return uc.repo.PageResourceStorage(ctx, req)
}
