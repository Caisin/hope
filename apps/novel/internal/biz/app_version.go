package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/appversion/v1"
	"hope/apps/novel/internal/data/ent"
)

type AppVersionRepo interface {
	CreateAppVersion(context.Context, *v1.AppVersionCreateReq) (*ent.AppVersion, error)
	DeleteAppVersion(context.Context, *v1.AppVersionDeleteReq) error
	BatchDeleteAppVersion(context.Context, *v1.AppVersionBatchDeleteReq) (int, error)
	UpdateAppVersion(context.Context, *v1.AppVersionUpdateReq) (*ent.AppVersion, error)
	GetAppVersion(context.Context, *v1.AppVersionReq) (*ent.AppVersion, error)
	PageAppVersion(context.Context, *v1.AppVersionPageReq) ([]*ent.AppVersion, error)
}

type AppVersionUseCase struct {
	repo AppVersionRepo
	log  *log.Helper
}

func NewAppVersionUseCase(repo AppVersionRepo, logger log.Logger) *AppVersionUseCase {
	return &AppVersionUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AppVersionUseCase) Create(ctx context.Context, req *v1.AppVersionCreateReq) (*ent.AppVersion, error) {
	return uc.repo.CreateAppVersion(ctx, req)
}
func (uc *AppVersionUseCase) Delete(ctx context.Context, req *v1.AppVersionDeleteReq) error {
	return uc.repo.DeleteAppVersion(ctx, req)
}
func (uc *AppVersionUseCase) BatchDelete(ctx context.Context, req *v1.AppVersionBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteAppVersion(ctx, req)
}
func (uc *AppVersionUseCase) Update(ctx context.Context, req *v1.AppVersionUpdateReq) (*ent.AppVersion, error) {
	return uc.repo.UpdateAppVersion(ctx, req)
}
func (uc *AppVersionUseCase) Get(ctx context.Context, req *v1.AppVersionReq) (*ent.AppVersion, error) {
	return uc.repo.GetAppVersion(ctx, req)
}
func (uc *AppVersionUseCase) Page(ctx context.Context, req *v1.AppVersionPageReq) ([]*ent.AppVersion, error) {
	return uc.repo.PageAppVersion(ctx, req)
}
