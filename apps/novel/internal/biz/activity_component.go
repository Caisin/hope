package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/activitycomponent/v1"
	"hope/apps/novel/internal/data/ent"
)

type ActivityComponentRepo interface {
	CreateActivityComponent(context.Context, *v1.ActivityComponentCreateReq) (*ent.ActivityComponent, error)
	DeleteActivityComponent(context.Context, *v1.ActivityComponentDeleteReq) error
	BatchDeleteActivityComponent(context.Context, *v1.ActivityComponentBatchDeleteReq) (int, error)
	UpdateActivityComponent(context.Context, *v1.ActivityComponentUpdateReq) (*ent.ActivityComponent, error)
	GetActivityComponent(context.Context, *v1.ActivityComponentReq) (*ent.ActivityComponent, error)
	PageActivityComponent(context.Context, *v1.ActivityComponentPageReq) ([]*ent.ActivityComponent, error)
}

type ActivityComponentUseCase struct {
	repo ActivityComponentRepo
	log  *log.Helper
}

func NewActivityComponentUseCase(repo ActivityComponentRepo, logger log.Logger) *ActivityComponentUseCase {
	return &ActivityComponentUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ActivityComponentUseCase) Create(ctx context.Context, req *v1.ActivityComponentCreateReq) (*ent.ActivityComponent, error) {
	return uc.repo.CreateActivityComponent(ctx, req)
}
func (uc *ActivityComponentUseCase) Delete(ctx context.Context, req *v1.ActivityComponentDeleteReq) error {
	return uc.repo.DeleteActivityComponent(ctx, req)
}
func (uc *ActivityComponentUseCase) BatchDelete(ctx context.Context, req *v1.ActivityComponentBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteActivityComponent(ctx, req)
}
func (uc *ActivityComponentUseCase) Update(ctx context.Context, req *v1.ActivityComponentUpdateReq) (*ent.ActivityComponent, error) {
	return uc.repo.UpdateActivityComponent(ctx, req)
}
func (uc *ActivityComponentUseCase) Get(ctx context.Context, req *v1.ActivityComponentReq) (*ent.ActivityComponent, error) {
	return uc.repo.GetActivityComponent(ctx, req)
}
func (uc *ActivityComponentUseCase) Page(ctx context.Context, req *v1.ActivityComponentPageReq) ([]*ent.ActivityComponent, error) {
	return uc.repo.PageActivityComponent(ctx, req)
}
