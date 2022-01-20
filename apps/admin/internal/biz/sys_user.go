package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysuser/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysUserRepo interface {
	CreateSysUser(context.Context, *v1.SysUserCreateReq) (*ent.SysUser, error)
	DeleteSysUser(context.Context, *v1.SysUserDeleteReq) error
	BatchDeleteSysUser(context.Context, *v1.SysUserBatchDeleteReq) (int, error)
	UpdateSysUser(context.Context, *v1.SysUserUpdateReq) (*ent.SysUser, error)
	GetSysUser(context.Context, *v1.SysUserReq) (*ent.SysUser, error)
	PageSysUser(context.Context, *v1.SysUserPageReq) ([]*ent.SysUser, error)
}

type SysUserUseCase struct {
	repo SysUserRepo
	log  *log.Helper
}

func NewSysUserUseCase(repo SysUserRepo, logger log.Logger) *SysUserUseCase {
	return &SysUserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysUserUseCase) Create(ctx context.Context, req *v1.SysUserCreateReq) (*ent.SysUser, error) {
	return uc.repo.CreateSysUser(ctx, req)
}
func (uc *SysUserUseCase) Delete(ctx context.Context, req *v1.SysUserDeleteReq) error {
	return uc.repo.DeleteSysUser(ctx, req)
}
func (uc *SysUserUseCase) BatchDelete(ctx context.Context, req *v1.SysUserBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysUser(ctx, req)
}
func (uc *SysUserUseCase) Update(ctx context.Context, req *v1.SysUserUpdateReq) (*ent.SysUser, error) {
	return uc.repo.UpdateSysUser(ctx, req)
}
func (uc *SysUserUseCase) Get(ctx context.Context, req *v1.SysUserReq) (*ent.SysUser, error) {
	return uc.repo.GetSysUser(ctx, req)
}
func (uc *SysUserUseCase) Page(ctx context.Context, req *v1.SysUserPageReq) ([]*ent.SysUser, error) {
	return uc.repo.PageSysUser(ctx, req)
}
