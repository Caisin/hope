package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSysLoginLog")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysLoginLogData, 0)
	for i := range datas {
		items = append(items, convert.SysLoginLogData2Reply(datas[i]))
	}
	reply := &pb.SysLoginLogPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SysLoginLogService) GetSysLoginLog(ctx context.Context, req *pb.SysLoginLogReq) (*pb.SysLoginLogReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysLoginLog")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysLoginLogReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysLoginLogData2Reply(data),
	}
	return reply, nil
}
func (s *SysLoginLogService) UpdateSysLoginLog(ctx context.Context, req *pb.SysLoginLogUpdateReq) (*pb.SysLoginLogUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysLoginLog")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysLoginLogUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysLoginLogData2Reply(data),
	}
	return reply, nil
}
func (s *SysLoginLogService) CreateSysLoginLog(ctx context.Context, req *pb.SysLoginLogCreateReq) (*pb.SysLoginLogCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysLoginLog")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysLoginLogCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysLoginLogData2Reply(data),
	}
	return reply, err
}
func (s *SysLoginLogService) DeleteSysLoginLog(ctx context.Context, req *pb.SysLoginLogDeleteReq) (*pb.SysLoginLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysLoginLog")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysLoginLogDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SysLoginLogService) BatchDeleteSysLoginLog(ctx context.Context, req *pb.SysLoginLogBatchDeleteReq) (*pb.SysLoginLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysLoginLog")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysLoginLogDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
