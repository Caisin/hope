package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/userconsume/v1"
	"hope/apps/param/internal/data/ent"
)

type UserConsumeRepo interface {
	CreateUserConsume(context.Context, *v1.UserConsumeCreateReq) (*ent.UserConsume, error)
	DeleteUserConsume(context.Context, *v1.UserConsumeDeleteReq) error
	BatchDeleteUserConsume(context.Context, *v1.UserConsumeBatchDeleteReq) (int, error)
	UpdateUserConsume(context.Context, *v1.UserConsumeUpdateReq) (*ent.UserConsume, error)
	GetUserConsume(context.Context, *v1.UserConsumeReq) (*ent.UserConsume, error)
	PageUserConsume(context.Context, *v1.UserConsumePageReq) ([]*ent.UserConsume, error)
}

type UserConsumeUseCase struct {
	repo UserConsumeRepo
	log  *log.Helper
}

func NewUserConsumeUseCase(repo UserConsumeRepo, logger log.Logger) *UserConsumeUseCase {
	return &UserConsumeUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserConsumeUseCase) Create(ctx context.Context, req *v1.UserConsumeCreateReq) (*ent.UserConsume, error) {
	return uc.repo.CreateUserConsume(ctx, req)
}
func (uc *UserConsumeUseCase) Delete(ctx context.Context, req *v1.UserConsumeDeleteReq) error {
	return uc.repo.DeleteUserConsume(ctx, req)
}
func (uc *UserConsumeUseCase) BatchDelete(ctx context.Context, req *v1.UserConsumeBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteUserConsume(ctx, req)
}
func (uc *UserConsumeUseCase) Update(ctx context.Context, req *v1.UserConsumeUpdateReq) (*ent.UserConsume, error) {
	return uc.repo.UpdateUserConsume(ctx, req)
}
func (uc *UserConsumeUseCase) Get(ctx context.Context, req *v1.UserConsumeReq) (*ent.UserConsume, error) {
	return uc.repo.GetUserConsume(ctx, req)
}
func (uc *UserConsumeUseCase) Page(ctx context.Context, req *v1.UserConsumePageReq) ([]*ent.UserConsume, error) {
	return uc.repo.PageUserConsume(ctx, req)
}
