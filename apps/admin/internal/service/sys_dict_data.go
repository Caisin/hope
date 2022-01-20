package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysdictdata/v1"
)

type SysDictDataService struct {
	pb.UnimplementedSysDictDataServer
	uc  *biz.SysDictDataUseCase
	log *log.Helper
}

func NewSysDictDataService(uc *biz.SysDictDataUseCase, logger log.Logger) *SysDictDataService {
	return &SysDictDataService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysDictDataService) GetPageSysDictData(ctx context.Context, req *pb.SysDictDataPageReq) (*pb.SysDictDataPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysDictDataReply, 0)
	for i := range datas {
		items = append(items, convert.SysDictDataData2Reply(datas[i]))
	}
	reply := &pb.SysDictDataPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysDictDataService) GetSysDictData(ctx context.Context, req *pb.SysDictDataReq) (*pb.SysDictDataReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysDictDataData2Reply(daya), err
}
func (s *SysDictDataService) UpdateSysDictData(ctx context.Context, req *pb.SysDictDataUpdateReq) (*pb.SysDictDataUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysDictDataData2UpdateReply(data), err
}
func (s *SysDictDataService) CreateSysDictData(ctx context.Context, req *pb.SysDictDataCreateReq) (*pb.SysDictDataCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysDictDataData2CreateReply(data), err
}
func (s *SysDictDataService) DeleteSysDictData(ctx context.Context, req *pb.SysDictDataDeleteReq) (*pb.SysDictDataDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysDictDataDeleteReply{Result: err == nil}, err
}
func (s *SysDictDataService) BatchDeleteSysDictData(ctx context.Context, req *pb.SysDictDataBatchDeleteReq) (*pb.SysDictDataDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysDictDataDeleteReply{Result: err == nil && num > 0}, err
}
