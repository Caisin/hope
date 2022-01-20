package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/adchangelog/v1"
	"hope/apps/novel/internal/data/ent"
)

type AdChangeLogRepo interface {
	CreateAdChangeLog(context.Context, *v1.AdChangeLogCreateReq) (*ent.AdChangeLog, error)
	DeleteAdChangeLog(context.Context, *v1.AdChangeLogDeleteReq) error
	BatchDeleteAdChangeLog(context.Context, *v1.AdChangeLogBatchDeleteReq) (int, error)
	UpdateAdChangeLog(context.Context, *v1.AdChangeLogUpdateReq) (*ent.AdChangeLog, error)
	GetAdChangeLog(context.Context, *v1.AdChangeLogReq) (*ent.AdChangeLog, error)
	PageAdChangeLog(context.Context, *v1.AdChangeLogPageReq) ([]*ent.AdChangeLog, error)
}

type AdChangeLogUseCase struct {
	repo AdChangeLogRepo
	log  *log.Helper
}

func NewAdChangeLogUseCase(repo AdChangeLogRepo, logger log.Logger) *AdChangeLogUseCase {
	return &AdChangeLogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AdChangeLogUseCase) Create(ctx context.Context, req *v1.AdChangeLogCreateReq) (*ent.AdChangeLog, error) {
	return uc.repo.CreateAdChangeLog(ctx, req)
}
func (uc *AdChangeLogUseCase) Delete(ctx context.Context, req *v1.AdChangeLogDeleteReq) error {
	return uc.repo.DeleteAdChangeLog(ctx, req)
}
func (uc *AdChangeLogUseCase) BatchDelete(ctx context.Context, req *v1.AdChangeLogBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteAdChangeLog(ctx, req)
}
func (uc *AdChangeLogUseCase) Update(ctx context.Context, req *v1.AdChangeLogUpdateReq) (*ent.AdChangeLog, error) {
	return uc.repo.UpdateAdChangeLog(ctx, req)
}
func (uc *AdChangeLogUseCase) Get(ctx context.Context, req *v1.AdChangeLogReq) (*ent.AdChangeLog, error) {
	return uc.repo.GetAdChangeLog(ctx, req)
}
func (uc *AdChangeLogUseCase) Page(ctx context.Context, req *v1.AdChangeLogPageReq) ([]*ent.AdChangeLog, error) {
	return uc.repo.PageAdChangeLog(ctx, req)
}
