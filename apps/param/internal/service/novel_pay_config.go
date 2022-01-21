package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/novelpayconfig/v1"
)

type NovelPayConfigService struct {
	pb.UnimplementedNovelPayConfigServer
	uc  *biz.NovelPayConfigUseCase
	log *log.Helper
}

func NewNovelPayConfigService(uc *biz.NovelPayConfigUseCase, logger log.Logger) *NovelPayConfigService {
	return &NovelPayConfigService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelPayConfigService) GetPageNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigPageReq) (*pb.NovelPayConfigPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageNovelPayConfig")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelPayConfigReply, 0)
	for i := range datas {
		items = append(items, convert.NovelPayConfigData2Reply(datas[i]))
	}
	reply := &pb.NovelPayConfigPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelPayConfigService) GetNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigReq) (*pb.NovelPayConfigReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelPayConfig")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelPayConfigData2Reply(daya), err
}
func (s *NovelPayConfigService) UpdateNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigUpdateReq) (*pb.NovelPayConfigUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelPayConfig")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelPayConfigData2UpdateReply(data), err
}
func (s *NovelPayConfigService) CreateNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigCreateReq) (*pb.NovelPayConfigCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelPayConfig")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelPayConfigData2CreateReply(data), err
}
func (s *NovelPayConfigService) DeleteNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigDeleteReq) (*pb.NovelPayConfigDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelPayConfig")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelPayConfigDeleteReply{Result: err == nil}, err
}
func (s *NovelPayConfigService) BatchDeleteNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigBatchDeleteReq) (*pb.NovelPayConfigDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelPayConfig")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelPayConfigDeleteReply{Result: err == nil && num > 0}, err
}
