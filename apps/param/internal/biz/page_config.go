package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/pageconfig/v1"
	"hope/apps/param/internal/data/ent"
)

type PageConfigRepo interface {
	CreatePageConfig(context.Context, *v1.PageConfigCreateReq) (*ent.PageConfig, error)
	DeletePageConfig(context.Context, *v1.PageConfigDeleteReq) error
	BatchDeletePageConfig(context.Context, *v1.PageConfigBatchDeleteReq) (int, error)
	UpdatePageConfig(context.Context, *v1.PageConfigUpdateReq) (*ent.PageConfig, error)
	GetPageConfig(context.Context, *v1.PageConfigReq) (*ent.PageConfig, error)
	PagePageConfig(context.Context, *v1.PageConfigPageReq) ([]*ent.PageConfig, error)
}

type PageConfigUseCase struct {
	repo PageConfigRepo
	log  *log.Helper
}

func NewPageConfigUseCase(repo PageConfigRepo, logger log.Logger) *PageConfigUseCase {
	return &PageConfigUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PageConfigUseCase) Create(ctx context.Context, req *v1.PageConfigCreateReq) (*ent.PageConfig, error) {
	return uc.repo.CreatePageConfig(ctx, req)
}
func (uc *PageConfigUseCase) Delete(ctx context.Context, req *v1.PageConfigDeleteReq) error {
	return uc.repo.DeletePageConfig(ctx, req)
}
func (uc *PageConfigUseCase) BatchDelete(ctx context.Context, req *v1.PageConfigBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeletePageConfig(ctx, req)
}
func (uc *PageConfigUseCase) Update(ctx context.Context, req *v1.PageConfigUpdateReq) (*ent.PageConfig, error) {
	return uc.repo.UpdatePageConfig(ctx, req)
}
func (uc *PageConfigUseCase) Get(ctx context.Context, req *v1.PageConfigReq) (*ent.PageConfig, error) {
	return uc.repo.GetPageConfig(ctx, req)
}
func (uc *PageConfigUseCase) Page(ctx context.Context, req *v1.PageConfigPageReq) ([]*ent.PageConfig, error) {
	return uc.repo.PagePageConfig(ctx, req)
}
