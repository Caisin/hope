package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
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

func (s *SysJobLogService) GetSysJobLogPage(ctx context.Context, req *pb.SysJobLogPageReq) (*pb.SysJobLogPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysJobLogPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysJobLogData, 0)
	for i := range datas {
		items = append(items, convert.SysJobLogData2Reply(datas[i]))
	}
	reply := &pb.SysJobLogPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SysJobLogService) GetSysJobLog(ctx context.Context, req *pb.SysJobLogReq) (*pb.SysJobLogReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysJobLog")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysJobLogReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysJobLogData2Reply(data),
	}
	return reply, nil
}
func (s *SysJobLogService) UpdateSysJobLog(ctx context.Context, req *pb.SysJobLogUpdateReq) (*pb.SysJobLogUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysJobLog")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysJobLogUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysJobLogData2Reply(data),
	}
	return reply, nil
}
func (s *SysJobLogService) CreateSysJobLog(ctx context.Context, req *pb.SysJobLogCreateReq) (*pb.SysJobLogCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysJobLog")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysJobLogCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysJobLogData2Reply(data),
	}
	return reply, err
}
func (s *SysJobLogService) DeleteSysJobLog(ctx context.Context, req *pb.SysJobLogDeleteReq) (*pb.SysJobLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysJobLog")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysJobLogDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SysJobLogService) BatchDeleteSysJobLog(ctx context.Context, req *pb.SysJobLogBatchDeleteReq) (*pb.SysJobLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysJobLog")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysJobLogDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
