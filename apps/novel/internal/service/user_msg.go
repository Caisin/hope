package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/usermsg/v1"
)

type UserMsgService struct {
	pb.UnimplementedUserMsgServer
	uc  *biz.UserMsgUseCase
	log *log.Helper
}

func NewUserMsgService(uc *biz.UserMsgUseCase, logger log.Logger) *UserMsgService {
	return &UserMsgService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserMsgService) GetPageUserMsg(ctx context.Context, req *pb.UserMsgPageReq) (*pb.UserMsgPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.UserMsgReply, 0)
	for i := range datas {
		items = append(items, convert.UserMsgData2Reply(datas[i]))
	}
	reply := &pb.UserMsgPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *UserMsgService) GetUserMsg(ctx context.Context, req *pb.UserMsgReq) (*pb.UserMsgReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.UserMsgData2Reply(daya), err
}
func (s *UserMsgService) UpdateUserMsg(ctx context.Context, req *pb.UserMsgUpdateReq) (*pb.UserMsgUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.UserMsgData2UpdateReply(data), err
}
func (s *UserMsgService) CreateUserMsg(ctx context.Context, req *pb.UserMsgCreateReq) (*pb.UserMsgCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.UserMsgData2CreateReply(data), err
}
func (s *UserMsgService) DeleteUserMsg(ctx context.Context, req *pb.UserMsgDeleteReq) (*pb.UserMsgDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.UserMsgDeleteReply{Result: err == nil}, err
}
func (s *UserMsgService) BatchDeleteUserMsg(ctx context.Context, req *pb.UserMsgBatchDeleteReq) (*pb.UserMsgDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.UserMsgDeleteReply{Result: err == nil && num > 0}, err
}
