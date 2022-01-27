package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelbookshelf/v1"
)

type NovelBookshelfService struct {
	pb.UnimplementedNovelBookshelfServer
	uc  *biz.NovelBookshelfUseCase
	log *log.Helper
}

func NewNovelBookshelfService(uc *biz.NovelBookshelfUseCase, logger log.Logger) *NovelBookshelfService {
	return &NovelBookshelfService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelBookshelfService) GetNovelBookshelfPage(ctx context.Context, req *pb.NovelBookshelfPageReq) (*pb.NovelBookshelfPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelBookshelfPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelBookshelfData, 0)
	for i := range datas {
		items = append(items, convert.NovelBookshelfData2Reply(datas[i]))
	}
	reply := &pb.NovelBookshelfPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *NovelBookshelfService) GetNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfReq) (*pb.NovelBookshelfReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelBookshelf")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelBookshelfReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelBookshelfData2Reply(data),
	}
	return reply, nil
}
func (s *NovelBookshelfService) UpdateNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfUpdateReq) (*pb.NovelBookshelfUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelBookshelf")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelBookshelfUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelBookshelfData2Reply(data),
	}
	return reply, nil
}
func (s *NovelBookshelfService) CreateNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfCreateReq) (*pb.NovelBookshelfCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelBookshelf")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelBookshelfCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelBookshelfData2Reply(data),
	}
	return reply, err
}
func (s *NovelBookshelfService) DeleteNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfDeleteReq) (*pb.NovelBookshelfDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelBookshelf")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelBookshelfDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *NovelBookshelfService) BatchDeleteNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfBatchDeleteReq) (*pb.NovelBookshelfDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelBookshelf")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelBookshelfDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
