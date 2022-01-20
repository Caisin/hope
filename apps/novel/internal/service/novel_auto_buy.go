package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelautobuy/v1"
)

type NovelAutoBuyService struct {
	pb.UnimplementedNovelAutoBuyServer
	uc  *biz.NovelAutoBuyUseCase
	log *log.Helper
}

func NewNovelAutoBuyService(uc *biz.NovelAutoBuyUseCase, logger log.Logger) *NovelAutoBuyService {
	return &NovelAutoBuyService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelAutoBuyService) GetPageNovelAutoBuy(ctx context.Context, req *pb.NovelAutoBuyPageReq) (*pb.NovelAutoBuyPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.NovelAutoBuyReply, 0)
	for i := range datas {
		items = append(items, convert.NovelAutoBuyData2Reply(datas[i]))
	}
	reply := &pb.NovelAutoBuyPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelAutoBuyService) GetNovelAutoBuy(ctx context.Context, req *pb.NovelAutoBuyReq) (*pb.NovelAutoBuyReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.NovelAutoBuyData2Reply(daya), err
}
func (s *NovelAutoBuyService) UpdateNovelAutoBuy(ctx context.Context, req *pb.NovelAutoBuyUpdateReq) (*pb.NovelAutoBuyUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.NovelAutoBuyData2UpdateReply(data), err
}
func (s *NovelAutoBuyService) CreateNovelAutoBuy(ctx context.Context, req *pb.NovelAutoBuyCreateReq) (*pb.NovelAutoBuyCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.NovelAutoBuyData2CreateReply(data), err
}
func (s *NovelAutoBuyService) DeleteNovelAutoBuy(ctx context.Context, req *pb.NovelAutoBuyDeleteReq) (*pb.NovelAutoBuyDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.NovelAutoBuyDeleteReply{Result: err == nil}, err
}
func (s *NovelAutoBuyService) BatchDeleteNovelAutoBuy(ctx context.Context, req *pb.NovelAutoBuyBatchDeleteReq) (*pb.NovelAutoBuyDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.NovelAutoBuyDeleteReply{Result: err == nil && num > 0}, err
}
