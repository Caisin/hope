package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/clienterror/v1"
)

type ClientErrorService struct {
	pb.UnimplementedClientErrorServer
	uc  *biz.ClientErrorUseCase
	log *log.Helper
}

func NewClientErrorService(uc *biz.ClientErrorUseCase, logger log.Logger) *ClientErrorService {
	return &ClientErrorService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ClientErrorService) GetClientErrorPage(ctx context.Context, req *pb.ClientErrorPageReq) (*pb.ClientErrorPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetClientErrorPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.ClientErrorData, 0)
	for i := range datas {
		items = append(items, convert.ClientErrorData2Reply(datas[i]))
	}
	reply := &pb.ClientErrorPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *ClientErrorService) GetClientError(ctx context.Context, req *pb.ClientErrorReq) (*pb.ClientErrorReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetClientError")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ClientErrorReply{
		Code:    200,
		Message: "success",
		Result:  convert.ClientErrorData2Reply(data),
	}
	return reply, nil
}
func (s *ClientErrorService) UpdateClientError(ctx context.Context, req *pb.ClientErrorUpdateReq) (*pb.ClientErrorUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateClientError")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ClientErrorUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.ClientErrorData2Reply(data),
	}
	return reply, nil
}
func (s *ClientErrorService) CreateClientError(ctx context.Context, req *pb.ClientErrorCreateReq) (*pb.ClientErrorCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateClientError")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ClientErrorCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.ClientErrorData2Reply(data),
	}
	return reply, err
}
func (s *ClientErrorService) DeleteClientError(ctx context.Context, req *pb.ClientErrorDeleteReq) (*pb.ClientErrorDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteClientError")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ClientErrorDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *ClientErrorService) BatchDeleteClientError(ctx context.Context, req *pb.ClientErrorBatchDeleteReq) (*pb.ClientErrorDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteClientError")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ClientErrorDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
