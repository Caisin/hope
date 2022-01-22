package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSysColumns")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysColumnsData, 0)
	for i := range datas {
		items = append(items, convert.SysColumnsData2Reply(datas[i]))
	}
	reply := &pb.SysColumnsPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SysColumnsService) GetSysColumns(ctx context.Context, req *pb.SysColumnsReq) (*pb.SysColumnsReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysColumns")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysColumnsReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysColumnsData2Reply(data),
	}
	return reply, nil
}
func (s *SysColumnsService) UpdateSysColumns(ctx context.Context, req *pb.SysColumnsUpdateReq) (*pb.SysColumnsUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysColumns")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysColumnsUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysColumnsData2Reply(data),
	}
	return reply, nil
}
func (s *SysColumnsService) CreateSysColumns(ctx context.Context, req *pb.SysColumnsCreateReq) (*pb.SysColumnsCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysColumns")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysColumnsCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysColumnsData2Reply(data),
	}
	return reply, err
}
func (s *SysColumnsService) DeleteSysColumns(ctx context.Context, req *pb.SysColumnsDeleteReq) (*pb.SysColumnsDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysColumns")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysColumnsDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SysColumnsService) BatchDeleteSysColumns(ctx context.Context, req *pb.SysColumnsBatchDeleteReq) (*pb.SysColumnsDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysColumns")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysColumnsDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
