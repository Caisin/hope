package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelclassify/v1"
)

type NovelClassifyService struct {
	pb.UnimplementedNovelClassifyServer
	uc  *biz.NovelClassifyUseCase
	log *log.Helper
}

func NewNovelClassifyService(uc *biz.NovelClassifyUseCase, logger log.Logger) *NovelClassifyService {
	return &NovelClassifyService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelClassifyService) GetNovelClassifyPage(ctx context.Context, req *pb.NovelClassifyPageReq) (*pb.NovelClassifyPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelClassifyPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelClassifyData, 0)
	for i := range datas {
		items = append(items, convert.NovelClassifyData2Reply(datas[i]))
	}
	reply := &pb.NovelClassifyPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *NovelClassifyService) GetNovelClassify(ctx context.Context, req *pb.NovelClassifyReq) (*pb.NovelClassifyReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelClassify")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelClassifyReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelClassifyData2Reply(data),
	}
	return reply, nil
}
func (s *NovelClassifyService) UpdateNovelClassify(ctx context.Context, req *pb.NovelClassifyUpdateReq) (*pb.NovelClassifyUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelClassify")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelClassifyUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelClassifyData2Reply(data),
	}
	return reply, nil
}
func (s *NovelClassifyService) CreateNovelClassify(ctx context.Context, req *pb.NovelClassifyCreateReq) (*pb.NovelClassifyCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelClassify")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelClassifyCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelClassifyData2Reply(data),
	}
	return reply, err
}
func (s *NovelClassifyService) DeleteNovelClassify(ctx context.Context, req *pb.NovelClassifyDeleteReq) (*pb.NovelClassifyDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelClassify")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelClassifyDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *NovelClassifyService) BatchDeleteNovelClassify(ctx context.Context, req *pb.NovelClassifyBatchDeleteReq) (*pb.NovelClassifyDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelClassify")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelClassifyDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
