package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysmenu/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysMenuRepo interface {
	CreateSysMenu(context.Context, *v1.SysMenuCreateReq) (*ent.SysMenu, error)
	DeleteSysMenu(context.Context, *v1.SysMenuDeleteReq) error
	BatchDeleteSysMenu(context.Context, *v1.SysMenuBatchDeleteReq) (int, error)
	UpdateSysMenu(context.Context, *v1.SysMenuUpdateReq) (*ent.SysMenu, error)
	GetSysMenu(context.Context, *v1.SysMenuReq) (*ent.SysMenu, error)
	PageSysMenu(context.Context, *v1.SysMenuPageReq) ([]*ent.SysMenu, error)
}

type SysMenuUseCase struct {
	repo SysMenuRepo
	log  *log.Helper
}

func NewSysMenuUseCase(repo SysMenuRepo, logger log.Logger) *SysMenuUseCase {
	return &SysMenuUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysMenuUseCase) Create(ctx context.Context, req *v1.SysMenuCreateReq) (*ent.SysMenu, error) {
	return uc.repo.CreateSysMenu(ctx, req)
}
func (uc *SysMenuUseCase) Delete(ctx context.Context, req *v1.SysMenuDeleteReq) error {
	return uc.repo.DeleteSysMenu(ctx, req)
}
func (uc *SysMenuUseCase) BatchDelete(ctx context.Context, req *v1.SysMenuBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysMenu(ctx, req)
}
func (uc *SysMenuUseCase) Update(ctx context.Context, req *v1.SysMenuUpdateReq) (*ent.SysMenu, error) {
	return uc.repo.UpdateSysMenu(ctx, req)
}
func (uc *SysMenuUseCase) Get(ctx context.Context, req *v1.SysMenuReq) (*ent.SysMenu, error) {
	return uc.repo.GetSysMenu(ctx, req)
}
func (uc *SysMenuUseCase) Page(ctx context.Context, req *v1.SysMenuPageReq) ([]*ent.SysMenu, error) {
	return uc.repo.PageSysMenu(ctx, req)
}
