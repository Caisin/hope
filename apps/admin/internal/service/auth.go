package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	pb "hope/api/admin/auth/v1"
	"hope/apps/admin/internal/biz"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	uc  *biz.AuthUseCase
	log *log.Helper
}

func NewAuthService(uc *biz.AuthUseCase, logger log.Logger) *AuthService {
	return &AuthService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "Login")
	defer span.End()
	return s.uc.Login(ctx, req)
}

func (s *AuthService) Logout(ctx context.Context, req *pb.LogOutReq) (*pb.LogOutReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "Logout")
	defer span.End()
	err := s.uc.Logout(ctx, req)
	return &pb.LogOutReply{Code: 200, Message: "成功", Result: err == nil}, err
}

func (s *AuthService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.LoginReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserInfo")
	defer span.End()
	return s.uc.GetUserInfo(ctx, req)
}
func (s *AuthService) GetPermCode(ctx context.Context, req *pb.GetPermReq) (*pb.GetPermReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPermCode")
	defer span.End()
	return s.uc.GetPermCode(ctx, req)
}
