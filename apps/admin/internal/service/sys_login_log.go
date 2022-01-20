package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysloginlog/v1"
)

type SysLoginLogService struct {
	pb.UnimplementedSysLoginLogServer
	uc  *biz.SysLoginLogUseCase
	log *log.Helper
}

func NewSysLoginLogService(uc *biz.SysLoginLogUseCase, logger log.Logger) *SysLoginLogService {
	return &SysLoginLogService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysLoginLogService) GetPageSysLoginLog(ctx context.Context, req *pb.SysLoginLogPageReq) (*pb.SysLoginLogPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysLoginLogReply, 0)
	for i := range datas {
		items = append(items, convert.SysLoginLogData2Reply(datas[i]))
	}
	reply := &pb.SysLoginLogPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysLoginLogService) GetSysLoginLog(ctx context.Context, req *pb.SysLoginLogReq) (*pb.SysLoginLogReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysLoginLogData2Reply(daya), err
}
func (s *SysLoginLogService) UpdateSysLoginLog(ctx context.Context, req *pb.SysLoginLogUpdateReq) (*pb.SysLoginLogUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysLoginLogData2UpdateReply(data), err
}
func (s *SysLoginLogService) CreateSysLoginLog(ctx context.Context, req *pb.SysLoginLogCreateReq) (*pb.SysLoginLogCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysLoginLogData2CreateReply(data), err
}
func (s *SysLoginLogService) DeleteSysLoginLog(ctx context.Context, req *pb.SysLoginLogDeleteReq) (*pb.SysLoginLogDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysLoginLogDeleteReply{Result: err == nil}, err
}
func (s *SysLoginLogService) BatchDeleteSysLoginLog(ctx context.Context, req *pb.SysLoginLogBatchDeleteReq) (*pb.SysLoginLogDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysLoginLogDeleteReply{Result: err == nil && num > 0}, err
}
