package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (s *CustomerNovelConfigService) GetPageCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigPageReq) (*pb.CustomerNovelConfigPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.CustomerNovelConfigReply, 0)
	for i := range datas {
		items = append(items, convert.CustomerNovelConfigData2Reply(datas[i]))
	}
	reply := &pb.CustomerNovelConfigPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *CustomerNovelConfigService) GetCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigReq) (*pb.CustomerNovelConfigReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.CustomerNovelConfigData2Reply(daya), err
}
func (s *CustomerNovelConfigService) UpdateCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigUpdateReq) (*pb.CustomerNovelConfigUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.CustomerNovelConfigData2UpdateReply(data), err
}
func (s *CustomerNovelConfigService) CreateCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigCreateReq) (*pb.CustomerNovelConfigCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.CustomerNovelConfigData2CreateReply(data), err
}
func (s *CustomerNovelConfigService) DeleteCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigDeleteReq) (*pb.CustomerNovelConfigDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.CustomerNovelConfigDeleteReply{Result: err == nil}, err
}
func (s *CustomerNovelConfigService) BatchDeleteCustomerNovelConfig(ctx context.Context, req *pb.CustomerNovelConfigBatchDeleteReq) (*pb.CustomerNovelConfigDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.CustomerNovelConfigDeleteReply{Result: err == nil && num > 0}, err
}
