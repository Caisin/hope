package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/systables/v1"
)

type SysTablesService struct {
	pb.UnimplementedSysTablesServer
	uc  *biz.SysTablesUseCase
	log *log.Helper
}

func NewSysTablesService(uc *biz.SysTablesUseCase, logger log.Logger) *SysTablesService {
	return &SysTablesService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysTablesService) GetPageSysTables(ctx context.Context, req *pb.SysTablesPageReq) (*pb.SysTablesPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysTablesReply, 0)
	for i := range datas {
		items = append(items, convert.SysTablesData2Reply(datas[i]))
	}
	reply := &pb.SysTablesPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysTablesService) GetSysTables(ctx context.Context, req *pb.SysTablesReq) (*pb.SysTablesReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysTablesData2Reply(daya), err
}
func (s *SysTablesService) UpdateSysTables(ctx context.Context, req *pb.SysTablesUpdateReq) (*pb.SysTablesUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysTablesData2UpdateReply(data), err
}
func (s *SysTablesService) CreateSysTables(ctx context.Context, req *pb.SysTablesCreateReq) (*pb.SysTablesCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysTablesData2CreateReply(data), err
}
func (s *SysTablesService) DeleteSysTables(ctx context.Context, req *pb.SysTablesDeleteReq) (*pb.SysTablesDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysTablesDeleteReply{Result: err == nil}, err
}
func (s *SysTablesService) BatchDeleteSysTables(ctx context.Context, req *pb.SysTablesBatchDeleteReq) (*pb.SysTablesDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysTablesDeleteReply{Result: err == nil && num > 0}, err
}
