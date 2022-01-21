package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysoperalog/v1"
)

type SysOperaLogService struct {
	pb.UnimplementedSysOperaLogServer
	uc  *biz.SysOperaLogUseCase
	log *log.Helper
}

func NewSysOperaLogService(uc *biz.SysOperaLogUseCase, logger log.Logger) *SysOperaLogService {
	return &SysOperaLogService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysOperaLogService) GetPageSysOperaLog(ctx context.Context, req *pb.SysOperaLogPageReq) (*pb.SysOperaLogPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSysOperaLog")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysOperaLogReply, 0)
	for i := range datas {
		items = append(items, convert.SysOperaLogData2Reply(datas[i]))
	}
	reply := &pb.SysOperaLogPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysOperaLogService) GetSysOperaLog(ctx context.Context, req *pb.SysOperaLogReq) (*pb.SysOperaLogReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysOperaLog")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SysOperaLogData2Reply(daya), err
}
func (s *SysOperaLogService) UpdateSysOperaLog(ctx context.Context, req *pb.SysOperaLogUpdateReq) (*pb.SysOperaLogUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysOperaLog")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SysOperaLogData2UpdateReply(data), err
}
func (s *SysOperaLogService) CreateSysOperaLog(ctx context.Context, req *pb.SysOperaLogCreateReq) (*pb.SysOperaLogCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysOperaLog")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SysOperaLogData2CreateReply(data), err
}
func (s *SysOperaLogService) DeleteSysOperaLog(ctx context.Context, req *pb.SysOperaLogDeleteReq) (*pb.SysOperaLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysOperaLog")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysOperaLogDeleteReply{Result: err == nil}, err
}
func (s *SysOperaLogService) BatchDeleteSysOperaLog(ctx context.Context, req *pb.SysOperaLogBatchDeleteReq) (*pb.SysOperaLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysOperaLog")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysOperaLogDeleteReply{Result: err == nil && num > 0}, err
}
