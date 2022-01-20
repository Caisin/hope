package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/userevent/v1"
)

type UserEventService struct {
	pb.UnimplementedUserEventServer
	uc  *biz.UserEventUseCase
	log *log.Helper
}

func NewUserEventService(uc *biz.UserEventUseCase, logger log.Logger) *UserEventService {
	return &UserEventService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserEventService) GetPageUserEvent(ctx context.Context, req *pb.UserEventPageReq) (*pb.UserEventPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.UserEventReply, 0)
	for i := range datas {
		items = append(items, convert.UserEventData2Reply(datas[i]))
	}
	reply := &pb.UserEventPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *UserEventService) GetUserEvent(ctx context.Context, req *pb.UserEventReq) (*pb.UserEventReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.UserEventData2Reply(daya), err
}
func (s *UserEventService) UpdateUserEvent(ctx context.Context, req *pb.UserEventUpdateReq) (*pb.UserEventUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.UserEventData2UpdateReply(data), err
}
func (s *UserEventService) CreateUserEvent(ctx context.Context, req *pb.UserEventCreateReq) (*pb.UserEventCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.UserEventData2CreateReply(data), err
}
func (s *UserEventService) DeleteUserEvent(ctx context.Context, req *pb.UserEventDeleteReq) (*pb.UserEventDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.UserEventDeleteReply{Result: err == nil}, err
}
func (s *UserEventService) BatchDeleteUserEvent(ctx context.Context, req *pb.UserEventBatchDeleteReq) (*pb.UserEventDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.UserEventDeleteReply{Result: err == nil && num > 0}, err
}
