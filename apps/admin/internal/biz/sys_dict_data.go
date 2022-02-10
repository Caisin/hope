package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysdictdata/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysDictDataRepo interface {
	CreateSysDictData(context.Context, *v1.SysDictDataCreateReq) (*ent.SysDictData, error)
	DeleteSysDictData(context.Context, *v1.SysDictDataDeleteReq) error
	BatchDeleteSysDictData(context.Context, *v1.SysDictDataBatchDeleteReq) (int, error)
	UpdateSysDictData(context.Context, *v1.SysDictDataUpdateReq) (*ent.SysDictData, error)
	GetSysDictData(context.Context, *v1.SysDictDataReq) (*ent.SysDictData, error)
	PageSysDictData(context.Context, *v1.SysDictDataPageReq) ([]*ent.SysDictData, error)
	GetSysDictDataByType(context.Context, *v1.GetDataByTypeReq) ([]*ent.SysDictData, error)
}

type SysDictDataUseCase struct {
	repo SysDictDataRepo
	log  *log.Helper
}

func NewSysDictDataUseCase(repo SysDictDataRepo, logger log.Logger) *SysDictDataUseCase {
	return &SysDictDataUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysDictDataUseCase) Create(ctx context.Context, req *v1.SysDictDataCreateReq) (*ent.SysDictData, error) {
	return uc.repo.CreateSysDictData(ctx, req)
}
func (uc *SysDictDataUseCase) Delete(ctx context.Context, req *v1.SysDictDataDeleteReq) error {
	return uc.repo.DeleteSysDictData(ctx, req)
}
func (uc *SysDictDataUseCase) BatchDelete(ctx context.Context, req *v1.SysDictDataBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysDictData(ctx, req)
}
func (uc *SysDictDataUseCase) Update(ctx context.Context, req *v1.SysDictDataUpdateReq) (*ent.SysDictData, error) {
	return uc.repo.UpdateSysDictData(ctx, req)
}
func (uc *SysDictDataUseCase) Get(ctx context.Context, req *v1.SysDictDataReq) (*ent.SysDictData, error) {
	return uc.repo.GetSysDictData(ctx, req)
}
func (uc *SysDictDataUseCase) Page(ctx context.Context, req *v1.SysDictDataPageReq) ([]*ent.SysDictData, error) {
	return uc.repo.PageSysDictData(ctx, req)
}

func (uc *SysDictDataUseCase) GetByType(ctx context.Context, req *v1.GetDataByTypeReq) ([]*ent.SysDictData, error) {
	return uc.repo.GetSysDictDataByType(ctx, req)
}
