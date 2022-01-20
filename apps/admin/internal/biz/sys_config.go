package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysconfig/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysConfigRepo interface {
	CreateSysConfig(context.Context, *v1.SysConfigCreateReq) (*ent.SysConfig, error)
	DeleteSysConfig(context.Context, *v1.SysConfigDeleteReq) error
	BatchDeleteSysConfig(context.Context, *v1.SysConfigBatchDeleteReq) (int, error)
	UpdateSysConfig(context.Context, *v1.SysConfigUpdateReq) (*ent.SysConfig, error)
	GetSysConfig(context.Context, *v1.SysConfigReq) (*ent.SysConfig, error)
	PageSysConfig(context.Context, *v1.SysConfigPageReq) ([]*ent.SysConfig, error)
}

type SysConfigUseCase struct {
	repo SysConfigRepo
	log  *log.Helper
}

func NewSysConfigUseCase(repo SysConfigRepo, logger log.Logger) *SysConfigUseCase {
	return &SysConfigUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysConfigUseCase) Create(ctx context.Context, req *v1.SysConfigCreateReq) (*ent.SysConfig, error) {
	return uc.repo.CreateSysConfig(ctx, req)
}
func (uc *SysConfigUseCase) Delete(ctx context.Context, req *v1.SysConfigDeleteReq) error {
	return uc.repo.DeleteSysConfig(ctx, req)
}
func (uc *SysConfigUseCase) BatchDelete(ctx context.Context, req *v1.SysConfigBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysConfig(ctx, req)
}
func (uc *SysConfigUseCase) Update(ctx context.Context, req *v1.SysConfigUpdateReq) (*ent.SysConfig, error) {
	return uc.repo.UpdateSysConfig(ctx, req)
}
func (uc *SysConfigUseCase) Get(ctx context.Context, req *v1.SysConfigReq) (*ent.SysConfig, error) {
	return uc.repo.GetSysConfig(ctx, req)
}
func (uc *SysConfigUseCase) Page(ctx context.Context, req *v1.SysConfigPageReq) ([]*ent.SysConfig, error) {
	return uc.repo.PageSysConfig(ctx, req)
}
