package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/auth/v1"
)

type AuthRepo interface {
	Login(context.Context, *v1.LoginReq) (*v1.LoginReply, error)
	Logout(context.Context, *v1.LogOutReq) error
	GetUserInfo(context.Context, *v1.GetUserInfoReq) (*v1.LoginReply, error)
	GetPermCode(context.Context, *v1.GetPermReq) (*v1.GetPermReply, error)
	GetMenuList(context.Context, *v1.GetMenuReq) (*v1.GetMenuReply, error)
}

type AuthUseCase struct {
	repo AuthRepo
	log  *log.Helper
}

func NewAuthUseCase(repo AuthRepo, logger log.Logger) *AuthUseCase {
	return &AuthUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AuthUseCase) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return uc.repo.Login(ctx, req)
}
func (uc *AuthUseCase) Logout(ctx context.Context, req *v1.LogOutReq) error {
	return uc.repo.Logout(ctx, req)
}

func (uc *AuthUseCase) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (*v1.LoginReply, error) {
	return uc.repo.GetUserInfo(ctx, req)
}

func (uc *AuthUseCase) GetPermCode(ctx context.Context, req *v1.GetPermReq) (*v1.GetPermReply, error) {
	return uc.repo.GetPermCode(ctx, req)
}

func (uc *AuthUseCase) GetMenuList(ctx context.Context, req *v1.GetMenuReq) (*v1.GetMenuReply, error) {
	return uc.repo.GetMenuList(ctx, req)
}
