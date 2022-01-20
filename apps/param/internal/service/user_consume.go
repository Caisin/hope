package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/userconsume/v1"
)

type UserConsumeService struct {
	pb.UnimplementedUserConsumeServer
	uc  *biz.UserConsumeUseCase
	log *log.Helper
}

func NewUserConsumeService(uc *biz.UserConsumeUseCase, logger log.Logger) *UserConsumeService {
	return &UserConsumeService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserConsumeService) GetPageUserConsume(ctx context.Context, req *pb.UserConsumePageReq) (*pb.UserConsumePageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.UserConsumeReply, 0)
	for i := range datas {
		items = append(items, convert.UserConsumeData2Reply(datas[i]))
	}
	reply := &pb.UserConsumePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *UserConsumeService) GetUserConsume(ctx context.Context, req *pb.UserConsumeReq) (*pb.UserConsumeReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.UserConsumeData2Reply(daya), err
}
func (s *UserConsumeService) UpdateUserConsume(ctx context.Context, req *pb.UserConsumeUpdateReq) (*pb.UserConsumeUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.UserConsumeData2UpdateReply(data), err
}
func (s *UserConsumeService) CreateUserConsume(ctx context.Context, req *pb.UserConsumeCreateReq) (*pb.UserConsumeCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.UserConsumeData2CreateReply(data), err
}
func (s *UserConsumeService) DeleteUserConsume(ctx context.Context, req *pb.UserConsumeDeleteReq) (*pb.UserConsumeDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.UserConsumeDeleteReply{Result: err == nil}, err
}
func (s *UserConsumeService) BatchDeleteUserConsume(ctx context.Context, req *pb.UserConsumeBatchDeleteReq) (*pb.UserConsumeDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.UserConsumeDeleteReply{Result: err == nil && num > 0}, err
}
