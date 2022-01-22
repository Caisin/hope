package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysconfig/v1"
)

type SysConfigService struct {
	pb.UnimplementedSysConfigServer
	uc  *biz.SysConfigUseCase
	log *log.Helper
}

func NewSysConfigService(uc *biz.SysConfigUseCase, logger log.Logger) *SysConfigService {
	return &SysConfigService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysConfigService) GetPageSysConfig(ctx context.Context, req *pb.SysConfigPageReq) (*pb.SysConfigPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSysConfig")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysConfigData, 0)
	for i := range datas {
		items = append(items, convert.SysConfigData2Reply(datas[i]))
	}
	reply := &pb.SysConfigPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SysConfigService) GetSysConfig(ctx context.Context, req *pb.SysConfigReq) (*pb.SysConfigReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysConfig")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysConfigReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysConfigData2Reply(data),
	}
	return reply, nil
}
func (s *SysConfigService) UpdateSysConfig(ctx context.Context, req *pb.SysConfigUpdateReq) (*pb.SysConfigUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysConfig")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysConfigUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysConfigData2Reply(data),
	}
	return reply, nil
}
func (s *SysConfigService) CreateSysConfig(ctx context.Context, req *pb.SysConfigCreateReq) (*pb.SysConfigCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysConfig")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysConfigCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysConfigData2Reply(data),
	}
	return reply, err
}
func (s *SysConfigService) DeleteSysConfig(ctx context.Context, req *pb.SysConfigDeleteReq) (*pb.SysConfigDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysConfig")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysConfigDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SysConfigService) BatchDeleteSysConfig(ctx context.Context, req *pb.SysConfigBatchDeleteReq) (*pb.SysConfigDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysConfig")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysConfigDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
