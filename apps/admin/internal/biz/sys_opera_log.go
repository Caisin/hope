package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysoperalog/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysOperaLogRepo interface {
	CreateSysOperaLog(context.Context, *v1.SysOperaLogCreateReq) (*ent.SysOperaLog, error)
	DeleteSysOperaLog(context.Context, *v1.SysOperaLogDeleteReq) error
	BatchDeleteSysOperaLog(context.Context, *v1.SysOperaLogBatchDeleteReq) (int, error)
	UpdateSysOperaLog(context.Context, *v1.SysOperaLogUpdateReq) (*ent.SysOperaLog, error)
	GetSysOperaLog(context.Context, *v1.SysOperaLogReq) (*ent.SysOperaLog, error)
	PageSysOperaLog(context.Context, *v1.SysOperaLogPageReq) ([]*ent.SysOperaLog, error)
}

type SysOperaLogUseCase struct {
	repo SysOperaLogRepo
	log  *log.Helper
}

func NewSysOperaLogUseCase(repo SysOperaLogRepo, logger log.Logger) *SysOperaLogUseCase {
	return &SysOperaLogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysOperaLogUseCase) Create(ctx context.Context, req *v1.SysOperaLogCreateReq) (*ent.SysOperaLog, error) {
	return uc.repo.CreateSysOperaLog(ctx, req)
}
func (uc *SysOperaLogUseCase) Delete(ctx context.Context, req *v1.SysOperaLogDeleteReq) error {
	return uc.repo.DeleteSysOperaLog(ctx, req)
}
func (uc *SysOperaLogUseCase) BatchDelete(ctx context.Context, req *v1.SysOperaLogBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysOperaLog(ctx, req)
}
func (uc *SysOperaLogUseCase) Update(ctx context.Context, req *v1.SysOperaLogUpdateReq) (*ent.SysOperaLog, error) {
	return uc.repo.UpdateSysOperaLog(ctx, req)
}
func (uc *SysOperaLogUseCase) Get(ctx context.Context, req *v1.SysOperaLogReq) (*ent.SysOperaLog, error) {
	return uc.repo.GetSysOperaLog(ctx, req)
}
func (uc *SysOperaLogUseCase) Page(ctx context.Context, req *v1.SysOperaLogPageReq) ([]*ent.SysOperaLog, error) {
	return uc.repo.PageSysOperaLog(ctx, req)
}
