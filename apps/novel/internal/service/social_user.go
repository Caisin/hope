package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/socialuser/v1"
)

type SocialUserService struct {
	pb.UnimplementedSocialUserServer
	uc  *biz.SocialUserUseCase
	log *log.Helper
}

func NewSocialUserService(uc *biz.SocialUserUseCase, logger log.Logger) *SocialUserService {
	return &SocialUserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SocialUserService) GetPageSocialUser(ctx context.Context, req *pb.SocialUserPageReq) (*pb.SocialUserPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSocialUser")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SocialUserData, 0)
	for i := range datas {
		items = append(items, convert.SocialUserData2Reply(datas[i]))
	}
	reply := &pb.SocialUserPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SocialUserService) GetSocialUser(ctx context.Context, req *pb.SocialUserReq) (*pb.SocialUserReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSocialUser")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SocialUserReply{
		Code:    200,
		Message: "success",
		Result:  convert.SocialUserData2Reply(data),
	}
	return reply, nil
}
func (s *SocialUserService) UpdateSocialUser(ctx context.Context, req *pb.SocialUserUpdateReq) (*pb.SocialUserUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSocialUser")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SocialUserUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SocialUserData2Reply(data),
	}
	return reply, nil
}
func (s *SocialUserService) CreateSocialUser(ctx context.Context, req *pb.SocialUserCreateReq) (*pb.SocialUserCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSocialUser")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SocialUserCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SocialUserData2Reply(data),
	}
	return reply, err
}
func (s *SocialUserService) DeleteSocialUser(ctx context.Context, req *pb.SocialUserDeleteReq) (*pb.SocialUserDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSocialUser")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SocialUserDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SocialUserService) BatchDeleteSocialUser(ctx context.Context, req *pb.SocialUserBatchDeleteReq) (*pb.SocialUserDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSocialUser")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SocialUserDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
