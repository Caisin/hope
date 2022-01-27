package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/ambalance/v1"
)

type AmBalanceService struct {
	pb.UnimplementedAmBalanceServer
	uc  *biz.AmBalanceUseCase
	log *log.Helper
}

func NewAmBalanceService(uc *biz.AmBalanceUseCase, logger log.Logger) *AmBalanceService {
	return &AmBalanceService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AmBalanceService) GetAmBalancePage(ctx context.Context, req *pb.AmBalancePageReq) (*pb.AmBalancePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetAmBalancePage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.AmBalanceData, 0)
	for i := range datas {
		items = append(items, convert.AmBalanceData2Reply(datas[i]))
	}
	reply := &pb.AmBalancePageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *AmBalanceService) GetAmBalance(ctx context.Context, req *pb.AmBalanceReq) (*pb.AmBalanceReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetAmBalance")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AmBalanceReply{
		Code:    200,
		Message: "success",
		Result:  convert.AmBalanceData2Reply(data),
	}
	return reply, nil
}
func (s *AmBalanceService) UpdateAmBalance(ctx context.Context, req *pb.AmBalanceUpdateReq) (*pb.AmBalanceUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateAmBalance")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AmBalanceUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.AmBalanceData2Reply(data),
	}
	return reply, nil
}
func (s *AmBalanceService) CreateAmBalance(ctx context.Context, req *pb.AmBalanceCreateReq) (*pb.AmBalanceCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateAmBalance")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AmBalanceCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.AmBalanceData2Reply(data),
	}
	return reply, err
}
func (s *AmBalanceService) DeleteAmBalance(ctx context.Context, req *pb.AmBalanceDeleteReq) (*pb.AmBalanceDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteAmBalance")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AmBalanceDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *AmBalanceService) BatchDeleteAmBalance(ctx context.Context, req *pb.AmBalanceBatchDeleteReq) (*pb.AmBalanceDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteAmBalance")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AmBalanceDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
