package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysdept/v1"
)

type SysDeptService struct {
	pb.UnimplementedSysDeptServer
	uc  *biz.SysDeptUseCase
	log *log.Helper
}

func NewSysDeptService(uc *biz.SysDeptUseCase, logger log.Logger) *SysDeptService {
	return &SysDeptService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysDeptService) GetPageSysDept(ctx context.Context, req *pb.SysDeptPageReq) (*pb.SysDeptPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysDeptReply, 0)
	for i := range datas {
		items = append(items, convert.SysDeptData2Reply(datas[i]))
	}
	reply := &pb.SysDeptPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysDeptService) GetSysDept(ctx context.Context, req *pb.SysDeptReq) (*pb.SysDeptReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysDeptData2Reply(daya), err
}
func (s *SysDeptService) UpdateSysDept(ctx context.Context, req *pb.SysDeptUpdateReq) (*pb.SysDeptUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysDeptData2UpdateReply(data), err
}
func (s *SysDeptService) CreateSysDept(ctx context.Context, req *pb.SysDeptCreateReq) (*pb.SysDeptCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysDeptData2CreateReply(data), err
}
func (s *SysDeptService) DeleteSysDept(ctx context.Context, req *pb.SysDeptDeleteReq) (*pb.SysDeptDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysDeptDeleteReply{Result: err == nil}, err
}
func (s *SysDeptService) BatchDeleteSysDept(ctx context.Context, req *pb.SysDeptBatchDeleteReq) (*pb.SysDeptDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysDeptDeleteReply{Result: err == nil && num > 0}, err
}
