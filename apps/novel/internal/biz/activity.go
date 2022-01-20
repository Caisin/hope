package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/activity/v1"
	"hope/apps/novel/internal/data/ent"
)

type ActivityRepo interface {
	CreateActivity(context.Context, *v1.ActivityCreateReq) (*ent.Activity, error)
	DeleteActivity(context.Context, *v1.ActivityDeleteReq) error
	BatchDeleteActivity(context.Context, *v1.ActivityBatchDeleteReq) (int, error)
	UpdateActivity(context.Context, *v1.ActivityUpdateReq) (*ent.Activity, error)
	GetActivity(context.Context, *v1.ActivityReq) (*ent.Activity, error)
	PageActivity(context.Context, *v1.ActivityPageReq) ([]*ent.Activity, error)
}

type ActivityUseCase struct {
	repo ActivityRepo
	log  *log.Helper
}

func NewActivityUseCase(repo ActivityRepo, logger log.Logger) *ActivityUseCase {
	return &ActivityUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ActivityUseCase) Create(ctx context.Context, req *v1.ActivityCreateReq) (*ent.Activity, error) {
	return uc.repo.CreateActivity(ctx, req)
}
func (uc *ActivityUseCase) Delete(ctx context.Context, req *v1.ActivityDeleteReq) error {
	return uc.repo.DeleteActivity(ctx, req)
}
func (uc *ActivityUseCase) BatchDelete(ctx context.Context, req *v1.ActivityBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteActivity(ctx, req)
}
func (uc *ActivityUseCase) Update(ctx context.Context, req *v1.ActivityUpdateReq) (*ent.Activity, error) {
	return uc.repo.UpdateActivity(ctx, req)
}
func (uc *ActivityUseCase) Get(ctx context.Context, req *v1.ActivityReq) (*ent.Activity, error) {
	return uc.repo.GetActivity(ctx, req)
}
func (uc *ActivityUseCase) Page(ctx context.Context, req *v1.ActivityPageReq) ([]*ent.Activity, error) {
	return uc.repo.PageActivity(ctx, req)
}
