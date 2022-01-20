package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/sysapi/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysApiRepo interface {
	CreateSysApi(context.Context, *v1.SysApiCreateReq) (*ent.SysApi, error)
	DeleteSysApi(context.Context, *v1.SysApiDeleteReq) error
	BatchDeleteSysApi(context.Context, *v1.SysApiBatchDeleteReq) (int, error)
	UpdateSysApi(context.Context, *v1.SysApiUpdateReq) (*ent.SysApi, error)
	GetSysApi(context.Context, *v1.SysApiReq) (*ent.SysApi, error)
	PageSysApi(context.Context, *v1.SysApiPageReq) ([]*ent.SysApi, error)
}

type SysApiUseCase struct {
	repo SysApiRepo
	log  *log.Helper
}

func NewSysApiUseCase(repo SysApiRepo, logger log.Logger) *SysApiUseCase {
	return &SysApiUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysApiUseCase) Create(ctx context.Context, req *v1.SysApiCreateReq) (*ent.SysApi, error) {
	return uc.repo.CreateSysApi(ctx, req)
}
func (uc *SysApiUseCase) Delete(ctx context.Context, req *v1.SysApiDeleteReq) error {
	return uc.repo.DeleteSysApi(ctx, req)
}
func (uc *SysApiUseCase) BatchDelete(ctx context.Context, req *v1.SysApiBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysApi(ctx, req)
}
func (uc *SysApiUseCase) Update(ctx context.Context, req *v1.SysApiUpdateReq) (*ent.SysApi, error) {
	return uc.repo.UpdateSysApi(ctx, req)
}
func (uc *SysApiUseCase) Get(ctx context.Context, req *v1.SysApiReq) (*ent.SysApi, error) {
	return uc.repo.GetSysApi(ctx, req)
}
func (uc *SysApiUseCase) Page(ctx context.Context, req *v1.SysApiPageReq) ([]*ent.SysApi, error) {
	return uc.repo.PageSysApi(ctx, req)
}
