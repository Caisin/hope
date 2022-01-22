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

func (s *CustomerNovelsService) GetPageCustomerNovels(ctx context.Context, req *pb.CustomerNovelsPageReq) (*pb.CustomerNovelsPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageCustomerNovels")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.CustomerNovelsReply, 0)
	for i := range datas {
		items = append(items, convert.CustomerNovelsData2Reply(datas[i]))
	}
	reply := &pb.CustomerNovelsPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *CustomerNovelsService) GetCustomerNovels(ctx context.Context, req *pb.CustomerNovelsReq) (*pb.CustomerNovelsReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetCustomerNovels")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.CustomerNovelsData2Reply(daya), err
}
func (s *CustomerNovelsService) UpdateCustomerNovels(ctx context.Context, req *pb.CustomerNovelsUpdateReq) (*pb.CustomerNovelsUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateCustomerNovels")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.CustomerNovelsData2UpdateReply(data), err
}
func (s *CustomerNovelsService) CreateCustomerNovels(ctx context.Context, req *pb.CustomerNovelsCreateReq) (*pb.CustomerNovelsCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateCustomerNovels")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.CustomerNovelsData2CreateReply(data), err
}
func (s *CustomerNovelsService) DeleteCustomerNovels(ctx context.Context, req *pb.CustomerNovelsDeleteReq) (*pb.CustomerNovelsDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteCustomerNovels")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CustomerNovelsDeleteReply{Result: err == nil}, err
}
func (s *CustomerNovelsService) BatchDeleteCustomerNovels(ctx context.Context, req *pb.CustomerNovelsBatchDeleteReq) (*pb.CustomerNovelsDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteCustomerNovels")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CustomerNovelsDeleteReply{Result: err == nil && num > 0}, err
}
