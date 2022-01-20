package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/userevent/v1"
	"hope/apps/novel/internal/data/ent"
)

type UserEventRepo interface {
	CreateUserEvent(context.Context, *v1.UserEventCreateReq) (*ent.UserEvent, error)
	DeleteUserEvent(context.Context, *v1.UserEventDeleteReq) error
	BatchDeleteUserEvent(context.Context, *v1.UserEventBatchDeleteReq) (int, error)
	UpdateUserEvent(context.Context, *v1.UserEventUpdateReq) (*ent.UserEvent, error)
	GetUserEvent(context.Context, *v1.UserEventReq) (*ent.UserEvent, error)
	PageUserEvent(context.Context, *v1.UserEventPageReq) ([]*ent.UserEvent, error)
}

type UserEventUseCase struct {
	repo UserEventRepo
	log  *log.Helper
}

func NewUserEventUseCase(repo UserEventRepo, logger log.Logger) *UserEventUseCase {
	return &UserEventUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserEventUseCase) Create(ctx context.Context, req *v1.UserEventCreateReq) (*ent.UserEvent, error) {
	return uc.repo.CreateUserEvent(ctx, req)
}
func (uc *UserEventUseCase) Delete(ctx context.Context, req *v1.UserEventDeleteReq) error {
	return uc.repo.DeleteUserEvent(ctx, req)
}
func (uc *UserEventUseCase) BatchDelete(ctx context.Context, req *v1.UserEventBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteUserEvent(ctx, req)
}
func (uc *UserEventUseCase) Update(ctx context.Context, req *v1.UserEventUpdateReq) (*ent.UserEvent, error) {
	return uc.repo.UpdateUserEvent(ctx, req)
}
func (uc *UserEventUseCase) Get(ctx context.Context, req *v1.UserEventReq) (*ent.UserEvent, error) {
	return uc.repo.GetUserEvent(ctx, req)
}
func (uc *UserEventUseCase) Page(ctx context.Context, req *v1.UserEventPageReq) ([]*ent.UserEvent, error) {
	return uc.repo.PageUserEvent(ctx, req)
}
