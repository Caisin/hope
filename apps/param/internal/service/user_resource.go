package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/userresource/v1"
)

type UserResourceService struct {
	pb.UnimplementedUserResourceServer
	uc  *biz.UserResourceUseCase
	log *log.Helper
}

func NewUserResourceService(uc *biz.UserResourceUseCase, logger log.Logger) *UserResourceService {
	return &UserResourceService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserResourceService) GetPageUserResource(ctx context.Context, req *pb.UserResourcePageReq) (*pb.UserResourcePageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.UserResourceReply, 0)
	for i := range datas {
		items = append(items, convert.UserResourceData2Reply(datas[i]))
	}
	reply := &pb.UserResourcePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *UserResourceService) GetUserResource(ctx context.Context, req *pb.UserResourceReq) (*pb.UserResourceReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.UserResourceData2Reply(daya), err
}
func (s *UserResourceService) UpdateUserResource(ctx context.Context, req *pb.UserResourceUpdateReq) (*pb.UserResourceUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.UserResourceData2UpdateReply(data), err
}
func (s *UserResourceService) CreateUserResource(ctx context.Context, req *pb.UserResourceCreateReq) (*pb.UserResourceCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.UserResourceData2CreateReply(data), err
}
func (s *UserResourceService) DeleteUserResource(ctx context.Context, req *pb.UserResourceDeleteReq) (*pb.UserResourceDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.UserResourceDeleteReply{Result: err == nil}, err
}
func (s *UserResourceService) BatchDeleteUserResource(ctx context.Context, req *pb.UserResourceBatchDeleteReq) (*pb.UserResourceDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.UserResourceDeleteReply{Result: err == nil && num > 0}, err
}
