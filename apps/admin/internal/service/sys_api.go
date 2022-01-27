package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysapi/v1"
)

type SysApiService struct {
	pb.UnimplementedSysApiServer
	uc  *biz.SysApiUseCase
	log *log.Helper
}

func NewSysApiService(uc *biz.SysApiUseCase, logger log.Logger) *SysApiService {
	return &SysApiService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysApiService) GetSysApiPage(ctx context.Context, req *pb.SysApiPageReq) (*pb.SysApiPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysApiPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysApiData, 0)
	for i := range datas {
		items = append(items, convert.SysApiData2Reply(datas[i]))
	}
	reply := &pb.SysApiPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *SysApiService) GetSysApi(ctx context.Context, req *pb.SysApiReq) (*pb.SysApiReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysApi")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysApiReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysApiData2Reply(data),
	}
	return reply, nil
}
func (s *SysApiService) UpdateSysApi(ctx context.Context, req *pb.SysApiUpdateReq) (*pb.SysApiUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysApi")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysApiUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysApiData2Reply(data),
	}
	return reply, nil
}
func (s *SysApiService) CreateSysApi(ctx context.Context, req *pb.SysApiCreateReq) (*pb.SysApiCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysApi")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.SysApiCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.SysApiData2Reply(data),
	}
	return reply, err
}
func (s *SysApiService) DeleteSysApi(ctx context.Context, req *pb.SysApiDeleteReq) (*pb.SysApiDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysApi")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysApiDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *SysApiService) BatchDeleteSysApi(ctx context.Context, req *pb.SysApiBatchDeleteReq) (*pb.SysApiDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysApi")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysApiDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
