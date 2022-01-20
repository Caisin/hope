package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysjob/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysJobRepo interface {
	CreateSysJob(context.Context, *v1.SysJobCreateReq) (*ent.SysJob, error)
	DeleteSysJob(context.Context, *v1.SysJobDeleteReq) error
	BatchDeleteSysJob(context.Context, *v1.SysJobBatchDeleteReq) (int, error)
	UpdateSysJob(context.Context, *v1.SysJobUpdateReq) (*ent.SysJob, error)
	GetSysJob(context.Context, *v1.SysJobReq) (*ent.SysJob, error)
	PageSysJob(context.Context, *v1.SysJobPageReq) ([]*ent.SysJob, error)
}

type SysJobUseCase struct {
	repo SysJobRepo
	log  *log.Helper
}

func NewSysJobUseCase(repo SysJobRepo, logger log.Logger) *SysJobUseCase {
	return &SysJobUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysJobUseCase) Create(ctx context.Context, req *v1.SysJobCreateReq) (*ent.SysJob, error) {
	return uc.repo.CreateSysJob(ctx, req)
}
func (uc *SysJobUseCase) Delete(ctx context.Context, req *v1.SysJobDeleteReq) error {
	return uc.repo.DeleteSysJob(ctx, req)
}
func (uc *SysJobUseCase) BatchDelete(ctx context.Context, req *v1.SysJobBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysJob(ctx, req)
}
func (uc *SysJobUseCase) Update(ctx context.Context, req *v1.SysJobUpdateReq) (*ent.SysJob, error) {
	return uc.repo.UpdateSysJob(ctx, req)
}
func (uc *SysJobUseCase) Get(ctx context.Context, req *v1.SysJobReq) (*ent.SysJob, error) {
	return uc.repo.GetSysJob(ctx, req)
}
func (uc *SysJobUseCase) Page(ctx context.Context, req *v1.SysJobPageReq) ([]*ent.SysJob, error) {
	return uc.repo.PageSysJob(ctx, req)
}
