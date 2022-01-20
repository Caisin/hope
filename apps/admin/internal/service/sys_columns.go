package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/syscolumns/v1"
)

type SysColumnsService struct {
	pb.UnimplementedSysColumnsServer
	uc  *biz.SysColumnsUseCase
	log *log.Helper
}

func NewSysColumnsService(uc *biz.SysColumnsUseCase, logger log.Logger) *SysColumnsService {
	return &SysColumnsService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysColumnsService) GetPageSysColumns(ctx context.Context, req *pb.SysColumnsPageReq) (*pb.SysColumnsPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysColumnsReply, 0)
	for i := range datas {
		items = append(items, convert.SysColumnsData2Reply(datas[i]))
	}
	reply := &pb.SysColumnsPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysColumnsService) GetSysColumns(ctx context.Context, req *pb.SysColumnsReq) (*pb.SysColumnsReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysColumnsData2Reply(daya), err
}
func (s *SysColumnsService) UpdateSysColumns(ctx context.Context, req *pb.SysColumnsUpdateReq) (*pb.SysColumnsUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysColumnsData2UpdateReply(data), err
}
func (s *SysColumnsService) CreateSysColumns(ctx context.Context, req *pb.SysColumnsCreateReq) (*pb.SysColumnsCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysColumnsData2CreateReply(data), err
}
func (s *SysColumnsService) DeleteSysColumns(ctx context.Context, req *pb.SysColumnsDeleteReq) (*pb.SysColumnsDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysColumnsDeleteReply{Result: err == nil}, err
}
func (s *SysColumnsService) BatchDeleteSysColumns(ctx context.Context, req *pb.SysColumnsBatchDeleteReq) (*pb.SysColumnsDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysColumnsDeleteReply{Result: err == nil && num > 0}, err
}
