package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/resourcegroup/v1"
	"hope/apps/param/internal/data/ent"
)

type ResourceGroupRepo interface {
	CreateResourceGroup(context.Context, *v1.ResourceGroupCreateReq) (*ent.ResourceGroup, error)
	DeleteResourceGroup(context.Context, *v1.ResourceGroupDeleteReq) error
	BatchDeleteResourceGroup(context.Context, *v1.ResourceGroupBatchDeleteReq) (int, error)
	UpdateResourceGroup(context.Context, *v1.ResourceGroupUpdateReq) (*ent.ResourceGroup, error)
	GetResourceGroup(context.Context, *v1.ResourceGroupReq) (*ent.ResourceGroup, error)
	PageResourceGroup(context.Context, *v1.ResourceGroupPageReq) ([]*ent.ResourceGroup, error)
}

type ResourceGroupUseCase struct {
	repo ResourceGroupRepo
	log  *log.Helper
}

func NewResourceGroupUseCase(repo ResourceGroupRepo, logger log.Logger) *ResourceGroupUseCase {
	return &ResourceGroupUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ResourceGroupUseCase) Create(ctx context.Context, req *v1.ResourceGroupCreateReq) (*ent.ResourceGroup, error) {
	return uc.repo.CreateResourceGroup(ctx, req)
}
func (uc *ResourceGroupUseCase) Delete(ctx context.Context, req *v1.ResourceGroupDeleteReq) error {
	return uc.repo.DeleteResourceGroup(ctx, req)
}
func (uc *ResourceGroupUseCase) BatchDelete(ctx context.Context, req *v1.ResourceGroupBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteResourceGroup(ctx, req)
}
func (uc *ResourceGroupUseCase) Update(ctx context.Context, req *v1.ResourceGroupUpdateReq) (*ent.ResourceGroup, error) {
	return uc.repo.UpdateResourceGroup(ctx, req)
}
func (uc *ResourceGroupUseCase) Get(ctx context.Context, req *v1.ResourceGroupReq) (*ent.ResourceGroup, error) {
	return uc.repo.GetResourceGroup(ctx, req)
}
func (uc *ResourceGroupUseCase) Page(ctx context.Context, req *v1.ResourceGroupPageReq) ([]*ent.ResourceGroup, error) {
	return uc.repo.PageResourceGroup(ctx, req)
}
