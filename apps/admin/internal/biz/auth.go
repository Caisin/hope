package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/admin/auth/v1"
)

type AuthRepo interface {
	Login(context.Context, *v1.LoginReq) (*v1.LoginReply, error)
	Logout(context.Context, *v1.LogOutReq) error
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
