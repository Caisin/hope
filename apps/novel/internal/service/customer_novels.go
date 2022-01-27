package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/customernovels/v1"
)

type CustomerNovelsService struct {
	pb.UnimplementedCustomerNovelsServer
	uc  *biz.CustomerNovelsUseCase
	log *log.Helper
}

func NewCustomerNovelsService(uc *biz.CustomerNovelsUseCase, logger log.Logger) *CustomerNovelsService {
	return &CustomerNovelsService{uc: uc, log: log.NewHelper(logger)}
}

func (s *CustomerNovelsService) GetCustomerNovelsPage(ctx context.Context, req *pb.CustomerNovelsPageReq) (*pb.CustomerNovelsPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetCustomerNovelsPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.CustomerNovelsData, 0)
	for i := range datas {
		items = append(items, convert.CustomerNovelsData2Reply(datas[i]))
	}
	reply := &pb.CustomerNovelsPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *CustomerNovelsService) GetCustomerNovels(ctx context.Context, req *pb.CustomerNovelsReq) (*pb.CustomerNovelsReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetCustomerNovels")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.CustomerNovelsReply{
		Code:    200,
		Message: "success",
		Result:  convert.CustomerNovelsData2Reply(data),
	}
	return reply, nil
}
func (s *CustomerNovelsService) UpdateCustomerNovels(ctx context.Context, req *pb.CustomerNovelsUpdateReq) (*pb.CustomerNovelsUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateCustomerNovels")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.CustomerNovelsUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.CustomerNovelsData2Reply(data),
	}
	return reply, nil
}
func (s *CustomerNovelsService) CreateCustomerNovels(ctx context.Context, req *pb.CustomerNovelsCreateReq) (*pb.CustomerNovelsCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateCustomerNovels")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.CustomerNovelsCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.CustomerNovelsData2Reply(data),
	}
	return reply, err
}
func (s *CustomerNovelsService) DeleteCustomerNovels(ctx context.Context, req *pb.CustomerNovelsDeleteReq) (*pb.CustomerNovelsDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteCustomerNovels")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CustomerNovelsDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *CustomerNovelsService) BatchDeleteCustomerNovels(ctx context.Context, req *pb.CustomerNovelsBatchDeleteReq) (*pb.CustomerNovelsDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteCustomerNovels")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CustomerNovelsDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
