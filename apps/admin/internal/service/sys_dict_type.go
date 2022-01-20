package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysdicttype/v1"
)

type SysDictTypeService struct {
	pb.UnimplementedSysDictTypeServer
	uc  *biz.SysDictTypeUseCase
	log *log.Helper
}

func NewSysDictTypeService(uc *biz.SysDictTypeUseCase, logger log.Logger) *SysDictTypeService {
	return &SysDictTypeService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysDictTypeService) GetPageSysDictType(ctx context.Context, req *pb.SysDictTypePageReq) (*pb.SysDictTypePageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysDictTypeReply, 0)
	for i := range datas {
		items = append(items, convert.SysDictTypeData2Reply(datas[i]))
	}
	reply := &pb.SysDictTypePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysDictTypeService) GetSysDictType(ctx context.Context, req *pb.SysDictTypeReq) (*pb.SysDictTypeReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysDictTypeData2Reply(daya), err
}
func (s *SysDictTypeService) UpdateSysDictType(ctx context.Context, req *pb.SysDictTypeUpdateReq) (*pb.SysDictTypeUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysDictTypeData2UpdateReply(data), err
}
func (s *SysDictTypeService) CreateSysDictType(ctx context.Context, req *pb.SysDictTypeCreateReq) (*pb.SysDictTypeCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysDictTypeData2CreateReply(data), err
}
func (s *SysDictTypeService) DeleteSysDictType(ctx context.Context, req *pb.SysDictTypeDeleteReq) (*pb.SysDictTypeDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysDictTypeDeleteReply{Result: err == nil}, err
}
func (s *SysDictTypeService) BatchDeleteSysDictType(ctx context.Context, req *pb.SysDictTypeBatchDeleteReq) (*pb.SysDictTypeDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysDictTypeDeleteReply{Result: err == nil && num > 0}, err
}
