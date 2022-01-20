package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/assetitem/v1"
	"hope/apps/novel/internal/data/ent"
)

type AssetItemRepo interface {
	CreateAssetItem(context.Context, *v1.AssetItemCreateReq) (*ent.AssetItem, error)
	DeleteAssetItem(context.Context, *v1.AssetItemDeleteReq) error
	BatchDeleteAssetItem(context.Context, *v1.AssetItemBatchDeleteReq) (int, error)
	UpdateAssetItem(context.Context, *v1.AssetItemUpdateReq) (*ent.AssetItem, error)
	GetAssetItem(context.Context, *v1.AssetItemReq) (*ent.AssetItem, error)
	PageAssetItem(context.Context, *v1.AssetItemPageReq) ([]*ent.AssetItem, error)
}

type AssetItemUseCase struct {
	repo AssetItemRepo
	log  *log.Helper
}

func NewAssetItemUseCase(repo AssetItemRepo, logger log.Logger) *AssetItemUseCase {
	return &AssetItemUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AssetItemUseCase) Create(ctx context.Context, req *v1.AssetItemCreateReq) (*ent.AssetItem, error) {
	return uc.repo.CreateAssetItem(ctx, req)
}
func (uc *AssetItemUseCase) Delete(ctx context.Context, req *v1.AssetItemDeleteReq) error {
	return uc.repo.DeleteAssetItem(ctx, req)
}
func (uc *AssetItemUseCase) BatchDelete(ctx context.Context, req *v1.AssetItemBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteAssetItem(ctx, req)
}
func (uc *AssetItemUseCase) Update(ctx context.Context, req *v1.AssetItemUpdateReq) (*ent.AssetItem, error) {
	return uc.repo.UpdateAssetItem(ctx, req)
}
func (uc *AssetItemUseCase) Get(ctx context.Context, req *v1.AssetItemReq) (*ent.AssetItem, error) {
	return uc.repo.GetAssetItem(ctx, req)
}
func (uc *AssetItemUseCase) Page(ctx context.Context, req *v1.AssetItemPageReq) ([]*ent.AssetItem, error) {
	return uc.repo.PageAssetItem(ctx, req)
}
