package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SocialUserReply, 0)
	for i := range datas {
		items = append(items, convert.SocialUserData2Reply(datas[i]))
	}
	reply := &pb.SocialUserPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SocialUserService) GetSocialUser(ctx context.Context, req *pb.SocialUserReq) (*pb.SocialUserReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SocialUserData2Reply(daya), err
}
func (s *SocialUserService) UpdateSocialUser(ctx context.Context, req *pb.SocialUserUpdateReq) (*pb.SocialUserUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SocialUserData2UpdateReply(data), err
}
func (s *SocialUserService) CreateSocialUser(ctx context.Context, req *pb.SocialUserCreateReq) (*pb.SocialUserCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SocialUserData2CreateReply(data), err
}
func (s *SocialUserService) DeleteSocialUser(ctx context.Context, req *pb.SocialUserDeleteReq) (*pb.SocialUserDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SocialUserDeleteReply{Result: err == nil}, err
}
func (s *SocialUserService) BatchDeleteSocialUser(ctx context.Context, req *pb.SocialUserBatchDeleteReq) (*pb.SocialUserDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SocialUserDeleteReply{Result: err == nil && num > 0}, err
}
