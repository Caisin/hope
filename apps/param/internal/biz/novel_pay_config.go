package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/novelpayconfig/v1"
	"hope/apps/param/internal/data/ent"
)

type NovelPayConfigRepo interface {
	CreateNovelPayConfig(context.Context, *v1.NovelPayConfigCreateReq) (*ent.NovelPayConfig, error)
	DeleteNovelPayConfig(context.Context, *v1.NovelPayConfigDeleteReq) error
	BatchDeleteNovelPayConfig(context.Context, *v1.NovelPayConfigBatchDeleteReq) (int, error)
	UpdateNovelPayConfig(context.Context, *v1.NovelPayConfigUpdateReq) (*ent.NovelPayConfig, error)
	GetNovelPayConfig(context.Context, *v1.NovelPayConfigReq) (*ent.NovelPayConfig, error)
	PageNovelPayConfig(context.Context, *v1.NovelPayConfigPageReq) ([]*ent.NovelPayConfig, error)
}

type NovelPayConfigUseCase struct {
	repo NovelPayConfigRepo
	log  *log.Helper
}

func NewNovelPayConfigUseCase(repo NovelPayConfigRepo, logger log.Logger) *NovelPayConfigUseCase {
	return &NovelPayConfigUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelPayConfigUseCase) Create(ctx context.Context, req *v1.NovelPayConfigCreateReq) (*ent.NovelPayConfig, error) {
	return uc.repo.CreateNovelPayConfig(ctx, req)
}
func (uc *NovelPayConfigUseCase) Delete(ctx context.Context, req *v1.NovelPayConfigDeleteReq) error {
	return uc.repo.DeleteNovelPayConfig(ctx, req)
}
func (uc *NovelPayConfigUseCase) BatchDelete(ctx context.Context, req *v1.NovelPayConfigBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelPayConfig(ctx, req)
}
func (uc *NovelPayConfigUseCase) Update(ctx context.Context, req *v1.NovelPayConfigUpdateReq) (*ent.NovelPayConfig, error) {
	return uc.repo.UpdateNovelPayConfig(ctx, req)
}
func (uc *NovelPayConfigUseCase) Get(ctx context.Context, req *v1.NovelPayConfigReq) (*ent.NovelPayConfig, error) {
	return uc.repo.GetNovelPayConfig(ctx, req)
}
func (uc *NovelPayConfigUseCase) Page(ctx context.Context, req *v1.NovelPayConfigPageReq) ([]*ent.NovelPayConfig, error) {
	return uc.repo.PageNovelPayConfig(ctx, req)
}
