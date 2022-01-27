package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/pageconfig/v1"
)

type PageConfigService struct {
	pb.UnimplementedPageConfigServer
	uc  *biz.PageConfigUseCase
	log *log.Helper
}

func NewPageConfigService(uc *biz.PageConfigUseCase, logger log.Logger) *PageConfigService {
	return &PageConfigService{uc: uc, log: log.NewHelper(logger)}
}

func (s *PageConfigService) GetPageConfigPage(ctx context.Context, req *pb.PageConfigPageReq) (*pb.PageConfigPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageConfigPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.PageConfigData, 0)
	for i := range datas {
		items = append(items, convert.PageConfigData2Reply(datas[i]))
	}
	reply := &pb.PageConfigPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *PageConfigService) GetPageConfig(ctx context.Context, req *pb.PageConfigReq) (*pb.PageConfigReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageConfig")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.PageConfigReply{
		Code:    200,
		Message: "success",
		Result:  convert.PageConfigData2Reply(data),
	}
	return reply, nil
}
func (s *PageConfigService) UpdatePageConfig(ctx context.Context, req *pb.PageConfigUpdateReq) (*pb.PageConfigUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdatePageConfig")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.PageConfigUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.PageConfigData2Reply(data),
	}
	return reply, nil
}
func (s *PageConfigService) CreatePageConfig(ctx context.Context, req *pb.PageConfigCreateReq) (*pb.PageConfigCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreatePageConfig")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.PageConfigCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.PageConfigData2Reply(data),
	}
	return reply, err
}
func (s *PageConfigService) DeletePageConfig(ctx context.Context, req *pb.PageConfigDeleteReq) (*pb.PageConfigDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeletePageConfig")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.PageConfigDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *PageConfigService) BatchDeletePageConfig(ctx context.Context, req *pb.PageConfigBatchDeleteReq) (*pb.PageConfigDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeletePageConfig")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.PageConfigDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
