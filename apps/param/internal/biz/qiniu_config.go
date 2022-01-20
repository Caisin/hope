package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/qiniuconfig/v1"
	"hope/apps/param/internal/data/ent"
)

type QiniuConfigRepo interface {
	CreateQiniuConfig(context.Context, *v1.QiniuConfigCreateReq) (*ent.QiniuConfig, error)
	DeleteQiniuConfig(context.Context, *v1.QiniuConfigDeleteReq) error
	BatchDeleteQiniuConfig(context.Context, *v1.QiniuConfigBatchDeleteReq) (int, error)
	UpdateQiniuConfig(context.Context, *v1.QiniuConfigUpdateReq) (*ent.QiniuConfig, error)
	GetQiniuConfig(context.Context, *v1.QiniuConfigReq) (*ent.QiniuConfig, error)
	PageQiniuConfig(context.Context, *v1.QiniuConfigPageReq) ([]*ent.QiniuConfig, error)
}

type QiniuConfigUseCase struct {
	repo QiniuConfigRepo
	log  *log.Helper
}

func NewQiniuConfigUseCase(repo QiniuConfigRepo, logger log.Logger) *QiniuConfigUseCase {
	return &QiniuConfigUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *QiniuConfigUseCase) Create(ctx context.Context, req *v1.QiniuConfigCreateReq) (*ent.QiniuConfig, error) {
	return uc.repo.CreateQiniuConfig(ctx, req)
}
func (uc *QiniuConfigUseCase) Delete(ctx context.Context, req *v1.QiniuConfigDeleteReq) error {
	return uc.repo.DeleteQiniuConfig(ctx, req)
}
func (uc *QiniuConfigUseCase) BatchDelete(ctx context.Context, req *v1.QiniuConfigBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteQiniuConfig(ctx, req)
}
func (uc *QiniuConfigUseCase) Update(ctx context.Context, req *v1.QiniuConfigUpdateReq) (*ent.QiniuConfig, error) {
	return uc.repo.UpdateQiniuConfig(ctx, req)
}
func (uc *QiniuConfigUseCase) Get(ctx context.Context, req *v1.QiniuConfigReq) (*ent.QiniuConfig, error) {
	return uc.repo.GetQiniuConfig(ctx, req)
}
func (uc *QiniuConfigUseCase) Page(ctx context.Context, req *v1.QiniuConfigPageReq) ([]*ent.QiniuConfig, error) {
	return uc.repo.PageQiniuConfig(ctx, req)
}
