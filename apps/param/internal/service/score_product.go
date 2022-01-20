package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/scoreproduct/v1"
)

type ScoreProductService struct {
	pb.UnimplementedScoreProductServer
	uc  *biz.ScoreProductUseCase
	log *log.Helper
}

func NewScoreProductService(uc *biz.ScoreProductUseCase, logger log.Logger) *ScoreProductService {
	return &ScoreProductService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ScoreProductService) GetPageScoreProduct(ctx context.Context, req *pb.ScoreProductPageReq) (*pb.ScoreProductPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.ScoreProductReply, 0)
	for i := range datas {
		items = append(items, convert.ScoreProductData2Reply(datas[i]))
	}
	reply := &pb.ScoreProductPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *ScoreProductService) GetScoreProduct(ctx context.Context, req *pb.ScoreProductReq) (*pb.ScoreProductReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.ScoreProductData2Reply(daya), err
}
func (s *ScoreProductService) UpdateScoreProduct(ctx context.Context, req *pb.ScoreProductUpdateReq) (*pb.ScoreProductUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.ScoreProductData2UpdateReply(data), err
}
func (s *ScoreProductService) CreateScoreProduct(ctx context.Context, req *pb.ScoreProductCreateReq) (*pb.ScoreProductCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.ScoreProductData2CreateReply(data), err
}
func (s *ScoreProductService) DeleteScoreProduct(ctx context.Context, req *pb.ScoreProductDeleteReq) (*pb.ScoreProductDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.ScoreProductDeleteReply{Result: err == nil}, err
}
func (s *ScoreProductService) BatchDeleteScoreProduct(ctx context.Context, req *pb.ScoreProductBatchDeleteReq) (*pb.ScoreProductDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.ScoreProductDeleteReply{Result: err == nil && num > 0}, err
}
