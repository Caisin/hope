package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/assetchangelog/v1"
	"hope/apps/novel/internal/data/ent"
)

type AssetChangeLogRepo interface {
	CreateAssetChangeLog(context.Context, *v1.AssetChangeLogCreateReq) (*ent.AssetChangeLog, error)
	DeleteAssetChangeLog(context.Context, *v1.AssetChangeLogDeleteReq) error
	BatchDeleteAssetChangeLog(context.Context, *v1.AssetChangeLogBatchDeleteReq) (int, error)
	UpdateAssetChangeLog(context.Context, *v1.AssetChangeLogUpdateReq) (*ent.AssetChangeLog, error)
	GetAssetChangeLog(context.Context, *v1.AssetChangeLogReq) (*ent.AssetChangeLog, error)
	PageAssetChangeLog(context.Context, *v1.AssetChangeLogPageReq) ([]*ent.AssetChangeLog, error)
}

type AssetChangeLogUseCase struct {
	repo AssetChangeLogRepo
	log  *log.Helper
}

func NewAssetChangeLogUseCase(repo AssetChangeLogRepo, logger log.Logger) *AssetChangeLogUseCase {
	return &AssetChangeLogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AssetChangeLogUseCase) Create(ctx context.Context, req *v1.AssetChangeLogCreateReq) (*ent.AssetChangeLog, error) {
	return uc.repo.CreateAssetChangeLog(ctx, req)
}
func (uc *AssetChangeLogUseCase) Delete(ctx context.Context, req *v1.AssetChangeLogDeleteReq) error {
	return uc.repo.DeleteAssetChangeLog(ctx, req)
}
func (uc *AssetChangeLogUseCase) BatchDelete(ctx context.Context, req *v1.AssetChangeLogBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteAssetChangeLog(ctx, req)
}
func (uc *AssetChangeLogUseCase) Update(ctx context.Context, req *v1.AssetChangeLogUpdateReq) (*ent.AssetChangeLog, error) {
	return uc.repo.UpdateAssetChangeLog(ctx, req)
}
func (uc *AssetChangeLogUseCase) Get(ctx context.Context, req *v1.AssetChangeLogReq) (*ent.AssetChangeLog, error) {
	return uc.repo.GetAssetChangeLog(ctx, req)
}
func (uc *AssetChangeLogUseCase) Page(ctx context.Context, req *v1.AssetChangeLogPageReq) ([]*ent.AssetChangeLog, error) {
	return uc.repo.PageAssetChangeLog(ctx, req)
}
