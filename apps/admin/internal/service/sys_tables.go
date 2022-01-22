package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSysTables")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysTablesData, 0)
	for i := range datas {
		items = append(items, convert.SysTablesData2Reply(datas[i]))
	}
	reply := &pb.SysTablesPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SysTablesService) GetSysTables(ctx context.Context, req *pb.SysTablesReq) (*pb.SysTablesReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysTables")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysTablesReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysTablesData2Reply(data),
	}
	return reply, nil
}
func (s *SysTablesService) UpdateSysTables(ctx context.Context, req *pb.SysTablesUpdateReq) (*pb.SysTablesUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysTables")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysTablesUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysTablesData2Reply(data),
	}
	return reply, nil
}
func (s *SysTablesService) CreateSysTables(ctx context.Context, req *pb.SysTablesCreateReq) (*pb.SysTablesCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysTables")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysTablesCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysTablesData2Reply(data),
	}
	return reply, err
}
func (s *SysTablesService) DeleteSysTables(ctx context.Context, req *pb.SysTablesDeleteReq) (*pb.SysTablesDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysTables")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysTablesDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SysTablesService) BatchDeleteSysTables(ctx context.Context, req *pb.SysTablesBatchDeleteReq) (*pb.SysTablesDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysTables")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysTablesDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
