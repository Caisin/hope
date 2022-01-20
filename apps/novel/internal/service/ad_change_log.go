package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/adchangelog/v1"
)

type AdChangeLogService struct {
	pb.UnimplementedAdChangeLogServer
	uc  *biz.AdChangeLogUseCase
	log *log.Helper
}

func NewAdChangeLogService(uc *biz.AdChangeLogUseCase, logger log.Logger) *AdChangeLogService {
	return &AdChangeLogService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AdChangeLogService) GetPageAdChangeLog(ctx context.Context, req *pb.AdChangeLogPageReq) (*pb.AdChangeLogPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.AdChangeLogReply, 0)
	for i := range datas {
		items = append(items, convert.AdChangeLogData2Reply(datas[i]))
	}
	reply := &pb.AdChangeLogPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *AdChangeLogService) GetAdChangeLog(ctx context.Context, req *pb.AdChangeLogReq) (*pb.AdChangeLogReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.AdChangeLogData2Reply(daya), err
}
func (s *AdChangeLogService) UpdateAdChangeLog(ctx context.Context, req *pb.AdChangeLogUpdateReq) (*pb.AdChangeLogUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.AdChangeLogData2UpdateReply(data), err
}
func (s *AdChangeLogService) CreateAdChangeLog(ctx context.Context, req *pb.AdChangeLogCreateReq) (*pb.AdChangeLogCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.AdChangeLogData2CreateReply(data), err
}
func (s *AdChangeLogService) DeleteAdChangeLog(ctx context.Context, req *pb.AdChangeLogDeleteReq) (*pb.AdChangeLogDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.AdChangeLogDeleteReply{Result: err == nil}, err
}
func (s *AdChangeLogService) BatchDeleteAdChangeLog(ctx context.Context, req *pb.AdChangeLogBatchDeleteReq) (*pb.AdChangeLogDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.AdChangeLogDeleteReply{Result: err == nil && num > 0}, err
}
