package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageNovelAutoBuy")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelAutoBuy")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelAutoBuyData2Reply(daya), err
}
func (s *NovelAutoBuyService) UpdateNovelAutoBuy(ctx context.Context, req *pb.NovelAutoBuyUpdateReq) (*pb.NovelAutoBuyUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelAutoBuy")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelAutoBuyData2UpdateReply(data), err
}
func (s *NovelAutoBuyService) CreateNovelAutoBuy(ctx context.Context, req *pb.NovelAutoBuyCreateReq) (*pb.NovelAutoBuyCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelAutoBuy")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelAutoBuyData2CreateReply(data), err
}
func (s *NovelAutoBuyService) DeleteNovelAutoBuy(ctx context.Context, req *pb.NovelAutoBuyDeleteReq) (*pb.NovelAutoBuyDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelAutoBuy")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelAutoBuyDeleteReply{Result: err == nil}, err
}
func (s *NovelAutoBuyService) BatchDeleteNovelAutoBuy(ctx context.Context, req *pb.NovelAutoBuyBatchDeleteReq) (*pb.NovelAutoBuyDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelAutoBuy")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelAutoBuyDeleteReply{Result: err == nil && num > 0}, err
}
