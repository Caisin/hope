package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysdicttype/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysDictTypeRepo interface {
	CreateSysDictType(context.Context, *v1.SysDictTypeCreateReq) (*ent.SysDictType, error)
	DeleteSysDictType(context.Context, *v1.SysDictTypeDeleteReq) error
	BatchDeleteSysDictType(context.Context, *v1.SysDictTypeBatchDeleteReq) (int, error)
	UpdateSysDictType(context.Context, *v1.SysDictTypeUpdateReq) (*ent.SysDictType, error)
	GetSysDictType(context.Context, *v1.SysDictTypeReq) (*ent.SysDictType, error)
	PageSysDictType(context.Context, *v1.SysDictTypePageReq) ([]*ent.SysDictType, error)
}

type SysDictTypeUseCase struct {
	repo SysDictTypeRepo
	log  *log.Helper
}

func NewSysDictTypeUseCase(repo SysDictTypeRepo, logger log.Logger) *SysDictTypeUseCase {
	return &SysDictTypeUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysDictTypeUseCase) Create(ctx context.Context, req *v1.SysDictTypeCreateReq) (*ent.SysDictType, error) {
	return uc.repo.CreateSysDictType(ctx, req)
}
func (uc *SysDictTypeUseCase) Delete(ctx context.Context, req *v1.SysDictTypeDeleteReq) error {
	return uc.repo.DeleteSysDictType(ctx, req)
}
func (uc *SysDictTypeUseCase) BatchDelete(ctx context.Context, req *v1.SysDictTypeBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysDictType(ctx, req)
}
func (uc *SysDictTypeUseCase) Update(ctx context.Context, req *v1.SysDictTypeUpdateReq) (*ent.SysDictType, error) {
	return uc.repo.UpdateSysDictType(ctx, req)
}
func (uc *SysDictTypeUseCase) Get(ctx context.Context, req *v1.SysDictTypeReq) (*ent.SysDictType, error) {
	return uc.repo.GetSysDictType(ctx, req)
}
func (uc *SysDictTypeUseCase) Page(ctx context.Context, req *v1.SysDictTypePageReq) ([]*ent.SysDictType, error) {
	return uc.repo.PageSysDictType(ctx, req)
}
