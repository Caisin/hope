package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/adchannel/v1"
)

type AdChannelService struct {
	pb.UnimplementedAdChannelServer
	uc  *biz.AdChannelUseCase
	log *log.Helper
}

func NewAdChannelService(uc *biz.AdChannelUseCase, logger log.Logger) *AdChannelService {
	return &AdChannelService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AdChannelService) GetPageAdChannel(ctx context.Context, req *pb.AdChannelPageReq) (*pb.AdChannelPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.AdChannelReply, 0)
	for i := range datas {
		items = append(items, convert.AdChannelData2Reply(datas[i]))
	}
	reply := &pb.AdChannelPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *AdChannelService) GetAdChannel(ctx context.Context, req *pb.AdChannelReq) (*pb.AdChannelReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.AdChannelData2Reply(daya), err
}
func (s *AdChannelService) UpdateAdChannel(ctx context.Context, req *pb.AdChannelUpdateReq) (*pb.AdChannelUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.AdChannelData2UpdateReply(data), err
}
func (s *AdChannelService) CreateAdChannel(ctx context.Context, req *pb.AdChannelCreateReq) (*pb.AdChannelCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.AdChannelData2CreateReply(data), err
}
func (s *AdChannelService) DeleteAdChannel(ctx context.Context, req *pb.AdChannelDeleteReq) (*pb.AdChannelDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.AdChannelDeleteReply{Result: err == nil}, err
}
func (s *AdChannelService) BatchDeleteAdChannel(ctx context.Context, req *pb.AdChannelBatchDeleteReq) (*pb.AdChannelDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.AdChannelDeleteReply{Result: err == nil && num > 0}, err
}
