package biz
import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/syspost/v1"
	"hope/apps/admin/internal/data/ent"
)

type SysPostRepo interface {
	CreateSysPost(context.Context, *v1.SysPostCreateReq) (*ent.SysPost, error)
	DeleteSysPost(context.Context, *v1.SysPostDeleteReq) error
	BatchDeleteSysPost(context.Context, *v1.SysPostBatchDeleteReq) (int, error)
	UpdateSysPost(context.Context, *v1.SysPostUpdateReq) (*ent.SysPost, error)
	GetSysPost(context.Context, *v1.SysPostReq) (*ent.SysPost, error)
	PageSysPost(context.Context, *v1.SysPostPageReq) ([]*ent.SysPost, error)
}

type SysPostUseCase struct {
	repo SysPostRepo
	log  *log.Helper
}

func NewSysPostUseCase(repo SysPostRepo, logger log.Logger) *SysPostUseCase {
	return &SysPostUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysPostUseCase) Create(ctx context.Context, req *v1.SysPostCreateReq) (*ent.SysPost, error) {
	return uc.repo.CreateSysPost(ctx, req)
}
func (uc *SysPostUseCase) Delete(ctx context.Context, req *v1.SysPostDeleteReq) error {
	return uc.repo.DeleteSysPost(ctx, req)
}
func (uc *SysPostUseCase) BatchDelete(ctx context.Context, req *v1.SysPostBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSysPost(ctx, req)
}
func (uc *SysPostUseCase) Update(ctx context.Context, req *v1.SysPostUpdateReq) (*ent.SysPost, error) {
	return uc.repo.UpdateSysPost(ctx, req)
}
func (uc *SysPostUseCase) Get(ctx context.Context, req *v1.SysPostReq) (*ent.SysPost, error) {
	return uc.repo.GetSysPost(ctx, req)
}
func (uc *SysPostUseCase) Page(ctx context.Context, req *v1.SysPostPageReq) ([]*ent.SysPost, error) {
	return uc.repo.PageSysPost(ctx, req)
}
