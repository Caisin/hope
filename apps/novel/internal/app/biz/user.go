package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/app/user/v1"
)

type UserRepo interface {
	Login(context.Context, *v1.LoginReq) (*v1.LoginReply, error)
	Info(context.Context, *v1.InfoReq) (*v1.InfoReply, error)
}

type UserCases struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserCases(repo UserRepo, logger log.Logger) *UserCases {
	return &UserCases{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (e UserCases) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return e.repo.Login(ctx, req)
}
