package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysdept/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysDeptRepo interface {
	CreateSysDept(context.Context, *v1.SysDeptCreateReq) (*ent.SysDept, error)
	DeleteSysDept(context.Context, *v1.SysDeptDeleteReq) error
	BatchDeleteSysDept(context.Context, *v1.SysDeptBatchDeleteReq) (int, error)
	UpdateSysDept(context.Context, *v1.SysDeptUpdateReq) (*ent.SysDept, error)
	GetSysDept(context.Context, *v1.SysDeptReq) (*ent.SysDept, error)
	PageSysDept(context.Context, *v1.SysDeptPageReq) ([]*ent.SysDept, error)
}

type SysDeptUseCase struct {
	repo SysDeptRepo
	log  *log.Helper
}

func NewSysDeptUseCase(repo SysDeptRepo, logger log.Logger) *SysDeptUseCase {
	return &SysDeptUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysDeptUseCase) Create(ctx context.Context, req *v1.SysDeptCreateReq) (*ent.SysDept, error) {
	return uc.repo.CreateSysDept(ctx, req)
}
func (uc *SysDeptUseCase) Delete(ctx context.Context, req *v1.SysDeptDeleteReq) error {
	return uc.repo.DeleteSysDept(ctx, req)
}
func (uc *SysDeptUseCase) BatchDelete(ctx context.Context, req *v1.SysDeptBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysDept(ctx, req)
}
func (uc *SysDeptUseCase) Update(ctx context.Context, req *v1.SysDeptUpdateReq) (*ent.SysDept, error) {
	return uc.repo.UpdateSysDept(ctx, req)
}
func (uc *SysDeptUseCase) Get(ctx context.Context, req *v1.SysDeptReq) (*ent.SysDept, error) {
	return uc.repo.GetSysDept(ctx, req)
}
func (uc *SysDeptUseCase) Page(ctx context.Context, req *v1.SysDeptPageReq) ([]*ent.SysDept, error) {
	return uc.repo.PageSysDept(ctx, req)
}
