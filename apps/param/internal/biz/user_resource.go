package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/userresource/v1"
	"hope/apps/param/internal/data/ent"
)

type UserResourceRepo interface {
	CreateUserResource(context.Context, *v1.UserResourceCreateReq) (*ent.UserResource, error)
	DeleteUserResource(context.Context, *v1.UserResourceDeleteReq) error
	BatchDeleteUserResource(context.Context, *v1.UserResourceBatchDeleteReq) (int, error)
	UpdateUserResource(context.Context, *v1.UserResourceUpdateReq) (*ent.UserResource, error)
	GetUserResource(context.Context, *v1.UserResourceReq) (*ent.UserResource, error)
	PageUserResource(context.Context, *v1.UserResourcePageReq) ([]*ent.UserResource, error)
}

type UserResourceUseCase struct {
	repo UserResourceRepo
	log  *log.Helper
}

func NewUserResourceUseCase(repo UserResourceRepo, logger log.Logger) *UserResourceUseCase {
	return &UserResourceUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserResourceUseCase) Create(ctx context.Context, req *v1.UserResourceCreateReq) (*ent.UserResource, error) {
	return uc.repo.CreateUserResource(ctx, req)
}
func (uc *UserResourceUseCase) Delete(ctx context.Context, req *v1.UserResourceDeleteReq) error {
	return uc.repo.DeleteUserResource(ctx, req)
}
func (uc *UserResourceUseCase) BatchDelete(ctx context.Context, req *v1.UserResourceBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteUserResource(ctx, req)
}
func (uc *UserResourceUseCase) Update(ctx context.Context, req *v1.UserResourceUpdateReq) (*ent.UserResource, error) {
	return uc.repo.UpdateUserResource(ctx, req)
}
func (uc *UserResourceUseCase) Get(ctx context.Context, req *v1.UserResourceReq) (*ent.UserResource, error) {
	return uc.repo.GetUserResource(ctx, req)
}
func (uc *UserResourceUseCase) Page(ctx context.Context, req *v1.UserResourcePageReq) ([]*ent.UserResource, error) {
	return uc.repo.PageUserResource(ctx, req)
}
