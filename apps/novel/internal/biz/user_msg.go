package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/usermsg/v1"
	"hope/apps/novel/internal/data/ent"
)

type UserMsgRepo interface {
	CreateUserMsg(context.Context, *v1.UserMsgCreateReq) (*ent.UserMsg, error)
	DeleteUserMsg(context.Context, *v1.UserMsgDeleteReq) error
	BatchDeleteUserMsg(context.Context, *v1.UserMsgBatchDeleteReq) (int, error)
	UpdateUserMsg(context.Context, *v1.UserMsgUpdateReq) (*ent.UserMsg, error)
	GetUserMsg(context.Context, *v1.UserMsgReq) (*ent.UserMsg, error)
	PageUserMsg(context.Context, *v1.UserMsgPageReq) ([]*ent.UserMsg, error)
}

type UserMsgUseCase struct {
	repo UserMsgRepo
	log  *log.Helper
}

func NewUserMsgUseCase(repo UserMsgRepo, logger log.Logger) *UserMsgUseCase {
	return &UserMsgUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserMsgUseCase) Create(ctx context.Context, req *v1.UserMsgCreateReq) (*ent.UserMsg, error) {
	return uc.repo.CreateUserMsg(ctx, req)
}
func (uc *UserMsgUseCase) Delete(ctx context.Context, req *v1.UserMsgDeleteReq) error {
	return uc.repo.DeleteUserMsg(ctx, req)
}
func (uc *UserMsgUseCase) BatchDelete(ctx context.Context, req *v1.UserMsgBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteUserMsg(ctx, req)
}
func (uc *UserMsgUseCase) Update(ctx context.Context, req *v1.UserMsgUpdateReq) (*ent.UserMsg, error) {
	return uc.repo.UpdateUserMsg(ctx, req)
}
func (uc *UserMsgUseCase) Get(ctx context.Context, req *v1.UserMsgReq) (*ent.UserMsg, error) {
	return uc.repo.GetUserMsg(ctx, req)
}
func (uc *UserMsgUseCase) Page(ctx context.Context, req *v1.UserMsgPageReq) ([]*ent.UserMsg, error) {
	return uc.repo.PageUserMsg(ctx, req)
}
