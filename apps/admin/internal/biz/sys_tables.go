package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/systables/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysTablesRepo interface {
	CreateSysTables(context.Context, *v1.SysTablesCreateReq) (*ent.SysTables, error)
	DeleteSysTables(context.Context, *v1.SysTablesDeleteReq) error
	BatchDeleteSysTables(context.Context, *v1.SysTablesBatchDeleteReq) (int, error)
	UpdateSysTables(context.Context, *v1.SysTablesUpdateReq) (*ent.SysTables, error)
	GetSysTables(context.Context, *v1.SysTablesReq) (*ent.SysTables, error)
	PageSysTables(context.Context, *v1.SysTablesPageReq) ([]*ent.SysTables, error)
}

type SysTablesUseCase struct {
	repo SysTablesRepo
	log  *log.Helper
}

func NewSysTablesUseCase(repo SysTablesRepo, logger log.Logger) *SysTablesUseCase {
	return &SysTablesUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysTablesUseCase) Create(ctx context.Context, req *v1.SysTablesCreateReq) (*ent.SysTables, error) {
	return uc.repo.CreateSysTables(ctx, req)
}
func (uc *SysTablesUseCase) Delete(ctx context.Context, req *v1.SysTablesDeleteReq) error {
	return uc.repo.DeleteSysTables(ctx, req)
}
func (uc *SysTablesUseCase) BatchDelete(ctx context.Context, req *v1.SysTablesBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysTables(ctx, req)
}
func (uc *SysTablesUseCase) Update(ctx context.Context, req *v1.SysTablesUpdateReq) (*ent.SysTables, error) {
	return uc.repo.UpdateSysTables(ctx, req)
}
func (uc *SysTablesUseCase) Get(ctx context.Context, req *v1.SysTablesReq) (*ent.SysTables, error) {
	return uc.repo.GetSysTables(ctx, req)
}
func (uc *SysTablesUseCase) Page(ctx context.Context, req *v1.SysTablesPageReq) ([]*ent.SysTables, error) {
	return uc.repo.PageSysTables(ctx, req)
}
