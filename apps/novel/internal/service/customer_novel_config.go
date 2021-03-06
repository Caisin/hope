package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/customernovelconfig/v1"
)

type CustomerNovelConfigService struct {
	pb.UnimplementedCustomerNovelConfigServer
	uc  *biz.CustomerNovelConfigUseCase
	log *log.Helper
}

func NewCustomerNovelConfigService(uc *biz.CustomerNovelConfigUseCase, logger log.Logger) *CustomerNovelConfigService {
	return &CustomerNovelConfigService{uc: uc, log: log.NewHelper(logger)}
}

func (s *CustomerNovelConfigService) GetCustomerNovelConfigPage(ctx context.Context, req *pb.CustomerNovelConfigPageReq) (*pb.CustomerNovelConfigPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetCustomerNovelConfigPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.CustomerNovelConfigData, 0)
	for i := range datas {
		items = append(items, convert.CustomerNovelConfigData2Reply(datas[i]))
	}
	reply := &pb.CustomerNovelConfigPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *CustomerNovelConfigService) GetCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigReq) (*pb.CustomerNovelConfigReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetCustomerNovelConfig")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.CustomerNovelConfigReply{
		Code:    200,
		Message: "success",
		Result:  convert.CustomerNovelConfigData2Reply(data),
	}
	return reply, nil
}
func (s *CustomerNovelConfigService) UpdateCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigUpdateReq) (*pb.CustomerNovelConfigUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateCustomerNovelConfig")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.CustomerNovelConfigUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.CustomerNovelConfigData2Reply(data),
	}
	return reply, nil
}
func (s *CustomerNovelConfigService) CreateCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigCreateReq) (*pb.CustomerNovelConfigCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateCustomerNovelConfig")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.CustomerNovelConfigCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.CustomerNovelConfigData2Reply(data),
	}
	return reply, err
}
func (s *CustomerNovelConfigService) DeleteCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigDeleteReq) (*pb.CustomerNovelConfigDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteCustomerNovelConfig")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CustomerNovelConfigDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *CustomerNovelConfigService) BatchDeleteCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigBatchDeleteReq) (*pb.CustomerNovelConfigDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteCustomerNovelConfig")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CustomerNovelConfigDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
