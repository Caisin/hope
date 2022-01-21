package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageScoreProduct")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetScoreProduct")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.ScoreProductData2Reply(daya), err
}
func (s *ScoreProductService) UpdateScoreProduct(ctx context.Context, req *pb.ScoreProductUpdateReq) (*pb.ScoreProductUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateScoreProduct")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.ScoreProductData2UpdateReply(data), err
}
func (s *ScoreProductService) CreateScoreProduct(ctx context.Context, req *pb.ScoreProductCreateReq) (*pb.ScoreProductCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateScoreProduct")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.ScoreProductData2CreateReply(data), err
}
func (s *ScoreProductService) DeleteScoreProduct(ctx context.Context, req *pb.ScoreProductDeleteReq) (*pb.ScoreProductDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteScoreProduct")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ScoreProductDeleteReply{Result: err == nil}, err
}
func (s *ScoreProductService) BatchDeleteScoreProduct(ctx context.Context, req *pb.ScoreProductBatchDeleteReq) (*pb.ScoreProductDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteScoreProduct")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ScoreProductDeleteReply{Result: err == nil && num > 0}, err
}
