package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	datas, err := s.uc.Page(ctx, req)
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
	daya, err := s.uc.Get(ctx, req)
	return convert.CustomerNovelsData2Reply(daya), err
}
func (s *CustomerNovelsService) UpdateCustomerNovels(ctx context.Context, req *pb.CustomerNovelsUpdateReq) (*pb.CustomerNovelsUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.CustomerNovelsData2UpdateReply(data), err
}
func (s *CustomerNovelsService) CreateCustomerNovels(ctx context.Context, req *pb.CustomerNovelsCreateReq) (*pb.CustomerNovelsCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.CustomerNovelsData2CreateReply(data), err
}
func (s *CustomerNovelsService) DeleteCustomerNovels(ctx context.Context, req *pb.CustomerNovelsDeleteReq) (*pb.CustomerNovelsDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.CustomerNovelsDeleteReply{Result: err == nil}, err
}
func (s *CustomerNovelsService) BatchDeleteCustomerNovels(ctx context.Context, req *pb.CustomerNovelsBatchDeleteReq) (*pb.CustomerNovelsDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.CustomerNovelsDeleteReply{Result: err == nil && num > 0}, err
}
