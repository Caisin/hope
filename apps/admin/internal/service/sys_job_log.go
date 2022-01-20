package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysjoblog/v1"
)

type SysJobLogService struct {
	pb.UnimplementedSysJobLogServer
	uc  *biz.SysJobLogUseCase
	log *log.Helper
}

func NewSysJobLogService(uc *biz.SysJobLogUseCase, logger log.Logger) *SysJobLogService {
	return &SysJobLogService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysJobLogService) GetPageSysJobLog(ctx context.Context, req *pb.SysJobLogPageReq) (*pb.SysJobLogPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysJobLogReply, 0)
	for i := range datas {
		items = append(items, convert.SysJobLogData2Reply(datas[i]))
	}
	reply := &pb.SysJobLogPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysJobLogService) GetSysJobLog(ctx context.Context, req *pb.SysJobLogReq) (*pb.SysJobLogReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysJobLogData2Reply(daya), err
}
func (s *SysJobLogService) UpdateSysJobLog(ctx context.Context, req *pb.SysJobLogUpdateReq) (*pb.SysJobLogUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysJobLogData2UpdateReply(data), err
}
func (s *SysJobLogService) CreateSysJobLog(ctx context.Context, req *pb.SysJobLogCreateReq) (*pb.SysJobLogCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysJobLogData2CreateReply(data), err
}
func (s *SysJobLogService) DeleteSysJobLog(ctx context.Context, req *pb.SysJobLogDeleteReq) (*pb.SysJobLogDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysJobLogDeleteReply{Result: err == nil}, err
}
func (s *SysJobLogService) BatchDeleteSysJobLog(ctx context.Context, req *pb.SysJobLogBatchDeleteReq) (*pb.SysJobLogDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysJobLogDeleteReply{Result: err == nil && num > 0}, err
}
