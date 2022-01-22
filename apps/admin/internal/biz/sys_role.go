package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysrole/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysRoleRepo interface {
	CreateSysRole(context.Context, *v1.SysRoleCreateReq) (*ent.SysRole, error)
	DeleteSysRole(context.Context, *v1.SysRoleDeleteReq) error
	BatchDeleteSysRole(context.Context, *v1.SysRoleBatchDeleteReq) (int, error)
	UpdateSysRole(context.Context, *v1.SysRoleUpdateReq) (*ent.SysRole, error)
	GetSysRole(context.Context, *v1.SysRoleReq) (*ent.SysRole, error)
	PageSysRole(context.Context, *v1.SysRolePageReq) ([]*ent.SysRole, error)
}

type SysRoleUseCase struct {
	repo SysRoleRepo
	log  *log.Helper
}

func NewSysRoleUseCase(repo SysRoleRepo, logger log.Logger) *SysRoleUseCase {
	return &SysRoleUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysRoleUseCase) Create(ctx context.Context, req *v1.SysRoleCreateReq) (*ent.SysRole, error) {
	return uc.repo.CreateSysRole(ctx, req)
}
func (uc *SysRoleUseCase) Delete(ctx context.Context, req *v1.SysRoleDeleteReq) error {
	return uc.repo.DeleteSysRole(ctx, req)
}
func (uc *SysRoleUseCase) BatchDelete(ctx context.Context, req *v1.SysRoleBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysRole(ctx, req)
}
func (uc *SysRoleUseCase) Update(ctx context.Context, req *v1.SysRoleUpdateReq) (*ent.SysRole, error) {
	return uc.repo.UpdateSysRole(ctx, req)
}
func (uc *SysRoleUseCase) Get(ctx context.Context, req *v1.SysRoleReq) (*ent.SysRole, error) {
	return uc.repo.GetSysRole(ctx, req)
}
func (uc *SysRoleUseCase) Page(ctx context.Context, req *v1.SysRolePageReq) ([]*ent.SysRole, error) {
	return uc.repo.PageSysRole(ctx, req)
}
