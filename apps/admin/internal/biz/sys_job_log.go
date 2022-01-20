package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysjoblog/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysJobLogRepo interface {
	CreateSysJobLog(context.Context, *v1.SysJobLogCreateReq) (*ent.SysJobLog, error)
	DeleteSysJobLog(context.Context, *v1.SysJobLogDeleteReq) error
	BatchDeleteSysJobLog(context.Context, *v1.SysJobLogBatchDeleteReq) (int, error)
	UpdateSysJobLog(context.Context, *v1.SysJobLogUpdateReq) (*ent.SysJobLog, error)
	GetSysJobLog(context.Context, *v1.SysJobLogReq) (*ent.SysJobLog, error)
	PageSysJobLog(context.Context, *v1.SysJobLogPageReq) ([]*ent.SysJobLog, error)
}

type SysJobLogUseCase struct {
	repo SysJobLogRepo
	log  *log.Helper
}

func NewSysJobLogUseCase(repo SysJobLogRepo, logger log.Logger) *SysJobLogUseCase {
	return &SysJobLogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysJobLogUseCase) Create(ctx context.Context, req *v1.SysJobLogCreateReq) (*ent.SysJobLog, error) {
	return uc.repo.CreateSysJobLog(ctx, req)
}
func (uc *SysJobLogUseCase) Delete(ctx context.Context, req *v1.SysJobLogDeleteReq) error {
	return uc.repo.DeleteSysJobLog(ctx, req)
}
func (uc *SysJobLogUseCase) BatchDelete(ctx context.Context, req *v1.SysJobLogBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysJobLog(ctx, req)
}
func (uc *SysJobLogUseCase) Update(ctx context.Context, req *v1.SysJobLogUpdateReq) (*ent.SysJobLog, error) {
	return uc.repo.UpdateSysJobLog(ctx, req)
}
func (uc *SysJobLogUseCase) Get(ctx context.Context, req *v1.SysJobLogReq) (*ent.SysJobLog, error) {
	return uc.repo.GetSysJobLog(ctx, req)
}
func (uc *SysJobLogUseCase) Page(ctx context.Context, req *v1.SysJobLogPageReq) ([]*ent.SysJobLog, error) {
	return uc.repo.PageSysJobLog(ctx, req)
}
