package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/syspost/v1"
)

type SysPostService struct {
	pb.UnimplementedSysPostServer
	uc  *biz.SysPostUseCase
	log *log.Helper
}

func NewSysPostService(uc *biz.SysPostUseCase, logger log.Logger) *SysPostService {
	return &SysPostService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysPostService) GetPageSysPost(ctx context.Context, req *pb.SysPostPageReq) (*pb.SysPostPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysPostReply, 0)
	for i := range datas {
		items = append(items, convert.SysPostData2Reply(datas[i]))
	}
	reply := &pb.SysPostPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysPostService) GetSysPost(ctx context.Context, req *pb.SysPostReq) (*pb.SysPostReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysPostData2Reply(daya), err
}
func (s *SysPostService) UpdateSysPost(ctx context.Context, req *pb.SysPostUpdateReq) (*pb.SysPostUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysPostData2UpdateReply(data), err
}
func (s *SysPostService) CreateSysPost(ctx context.Context, req *pb.SysPostCreateReq) (*pb.SysPostCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysPostData2CreateReply(data), err
}
func (s *SysPostService) DeleteSysPost(ctx context.Context, req *pb.SysPostDeleteReq) (*pb.SysPostDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysPostDeleteReply{Result: err == nil}, err
}
func (s *SysPostService) BatchDeleteSysPost(ctx context.Context, req *pb.SysPostBatchDeleteReq) (*pb.SysPostDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysPostDeleteReply{Result: err == nil && num > 0}, err
}
