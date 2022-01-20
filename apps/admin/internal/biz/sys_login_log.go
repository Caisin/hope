package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysloginlog/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysLoginLogRepo interface {
	CreateSysLoginLog(context.Context, *v1.SysLoginLogCreateReq) (*ent.SysLoginLog, error)
	DeleteSysLoginLog(context.Context, *v1.SysLoginLogDeleteReq) error
	BatchDeleteSysLoginLog(context.Context, *v1.SysLoginLogBatchDeleteReq) (int, error)
	UpdateSysLoginLog(context.Context, *v1.SysLoginLogUpdateReq) (*ent.SysLoginLog, error)
	GetSysLoginLog(context.Context, *v1.SysLoginLogReq) (*ent.SysLoginLog, error)
	PageSysLoginLog(context.Context, *v1.SysLoginLogPageReq) ([]*ent.SysLoginLog, error)
}

type SysLoginLogUseCase struct {
	repo SysLoginLogRepo
	log  *log.Helper
}

func NewSysLoginLogUseCase(repo SysLoginLogRepo, logger log.Logger) *SysLoginLogUseCase {
	return &SysLoginLogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysLoginLogUseCase) Create(ctx context.Context, req *v1.SysLoginLogCreateReq) (*ent.SysLoginLog, error) {
	return uc.repo.CreateSysLoginLog(ctx, req)
}
func (uc *SysLoginLogUseCase) Delete(ctx context.Context, req *v1.SysLoginLogDeleteReq) error {
	return uc.repo.DeleteSysLoginLog(ctx, req)
}
func (uc *SysLoginLogUseCase) BatchDelete(ctx context.Context, req *v1.SysLoginLogBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysLoginLog(ctx, req)
}
func (uc *SysLoginLogUseCase) Update(ctx context.Context, req *v1.SysLoginLogUpdateReq) (*ent.SysLoginLog, error) {
	return uc.repo.UpdateSysLoginLog(ctx, req)
}
func (uc *SysLoginLogUseCase) Get(ctx context.Context, req *v1.SysLoginLogReq) (*ent.SysLoginLog, error) {
	return uc.repo.GetSysLoginLog(ctx, req)
}
func (uc *SysLoginLogUseCase) Page(ctx context.Context, req *v1.SysLoginLogPageReq) ([]*ent.SysLoginLog, error) {
	return uc.repo.PageSysLoginLog(ctx, req)
}
