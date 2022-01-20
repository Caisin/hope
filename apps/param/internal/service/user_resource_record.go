package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/userresourcerecord/v1"
)

type UserResourceRecordService struct {
	pb.UnimplementedUserResourceRecordServer
	uc  *biz.UserResourceRecordUseCase
	log *log.Helper
}

func NewUserResourceRecordService(uc *biz.UserResourceRecordUseCase, logger log.Logger) *UserResourceRecordService {
	return &UserResourceRecordService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserResourceRecordService) GetPageUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordPageReq) (*pb.UserResourceRecordPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.UserResourceRecordReply, 0)
	for i := range datas {
		items = append(items, convert.UserResourceRecordData2Reply(datas[i]))
	}
	reply := &pb.UserResourceRecordPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *UserResourceRecordService) GetUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordReq) (*pb.UserResourceRecordReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.UserResourceRecordData2Reply(daya), err
}
func (s *UserResourceRecordService) UpdateUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordUpdateReq) (*pb.UserResourceRecordUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.UserResourceRecordData2UpdateReply(data), err
}
func (s *UserResourceRecordService) CreateUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordCreateReq) (*pb.UserResourceRecordCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.UserResourceRecordData2CreateReply(data), err
}
func (s *UserResourceRecordService) DeleteUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordDeleteReq) (*pb.UserResourceRecordDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.UserResourceRecordDeleteReply{Result: err == nil}, err
}
func (s *UserResourceRecordService) BatchDeleteUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordBatchDeleteReq) (*pb.UserResourceRecordDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.UserResourceRecordDeleteReply{Result: err == nil && num > 0}, err
}
