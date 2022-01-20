package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/payorder/v1"
)

type PayOrderService struct {
	pb.UnimplementedPayOrderServer
	uc  *biz.PayOrderUseCase
	log *log.Helper
}

func NewPayOrderService(uc *biz.PayOrderUseCase, logger log.Logger) *PayOrderService {
	return &PayOrderService{uc: uc, log: log.NewHelper(logger)}
}

func (s *PayOrderService) GetPagePayOrder(ctx context.Context, req *pb.PayOrderPageReq) (*pb.PayOrderPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.PayOrderReply, 0)
	for i := range datas {
		items = append(items, convert.PayOrderData2Reply(datas[i]))
	}
	reply := &pb.PayOrderPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *PayOrderService) GetPayOrder(ctx context.Context, req *pb.PayOrderReq) (*pb.PayOrderReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.PayOrderData2Reply(daya), err
}
func (s *PayOrderService) UpdatePayOrder(ctx context.Context, req *pb.PayOrderUpdateReq) (*pb.PayOrderUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.PayOrderData2UpdateReply(data), err
}
func (s *PayOrderService) CreatePayOrder(ctx context.Context, req *pb.PayOrderCreateReq) (*pb.PayOrderCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.PayOrderData2CreateReply(data), err
}
func (s *PayOrderService) DeletePayOrder(ctx context.Context, req *pb.PayOrderDeleteReq) (*pb.PayOrderDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.PayOrderDeleteReply{Result: err == nil}, err
}
func (s *PayOrderService) BatchDeletePayOrder(ctx context.Context, req *pb.PayOrderBatchDeleteReq) (*pb.PayOrderDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.PayOrderDeleteReply{Result: err == nil && num > 0}, err
}
