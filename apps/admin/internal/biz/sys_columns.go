package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/syscolumns/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysColumnsRepo interface {
	CreateSysColumns(context.Context, *v1.SysColumnsCreateReq) (*ent.SysColumns, error)
	DeleteSysColumns(context.Context, *v1.SysColumnsDeleteReq) error
	BatchDeleteSysColumns(context.Context, *v1.SysColumnsBatchDeleteReq) (int, error)
	UpdateSysColumns(context.Context, *v1.SysColumnsUpdateReq) (*ent.SysColumns, error)
	GetSysColumns(context.Context, *v1.SysColumnsReq) (*ent.SysColumns, error)
	PageSysColumns(context.Context, *v1.SysColumnsPageReq) ([]*ent.SysColumns, error)
}

type SysColumnsUseCase struct {
	repo SysColumnsRepo
	log  *log.Helper
}

func NewSysColumnsUseCase(repo SysColumnsRepo, logger log.Logger) *SysColumnsUseCase {
	return &SysColumnsUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysColumnsUseCase) Create(ctx context.Context, req *v1.SysColumnsCreateReq) (*ent.SysColumns, error) {
	return uc.repo.CreateSysColumns(ctx, req)
}
func (uc *SysColumnsUseCase) Delete(ctx context.Context, req *v1.SysColumnsDeleteReq) error {
	return uc.repo.DeleteSysColumns(ctx, req)
}
func (uc *SysColumnsUseCase) BatchDelete(ctx context.Context, req *v1.SysColumnsBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysColumns(ctx, req)
}
func (uc *SysColumnsUseCase) Update(ctx context.Context, req *v1.SysColumnsUpdateReq) (*ent.SysColumns, error) {
	return uc.repo.UpdateSysColumns(ctx, req)
}
func (uc *SysColumnsUseCase) Get(ctx context.Context, req *v1.SysColumnsReq) (*ent.SysColumns, error) {
	return uc.repo.GetSysColumns(ctx, req)
}
func (uc *SysColumnsUseCase) Page(ctx context.Context, req *v1.SysColumnsPageReq) ([]*ent.SysColumns, error) {
	return uc.repo.PageSysColumns(ctx, req)
}
