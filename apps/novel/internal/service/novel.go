package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novel/v1"
)

type NovelService struct {
	pb.UnimplementedNovelServer
	uc  *biz.NovelUseCase
	log *log.Helper
}

func NewNovelService(uc *biz.NovelUseCase, logger log.Logger) *NovelService {
	return &NovelService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelService) GetNovelPage(ctx context.Context, req *pb.NovelPageReq) (*pb.NovelPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelData, 0)
	for i := range datas {
		items = append(items, convert.NovelData2Reply(datas[i]))
	}
	reply := &pb.NovelPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *NovelService) GetNovel(ctx context.Context, req *pb.NovelReq) (*pb.NovelReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovel")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelData2Reply(data),
	}
	return reply, nil
}
func (s *NovelService) UpdateNovel(ctx context.Context, req *pb.NovelUpdateReq) (*pb.NovelUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovel")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelData2Reply(data),
	}
	return reply, nil
}
func (s *NovelService) CreateNovel(ctx context.Context, req *pb.NovelCreateReq) (*pb.NovelCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovel")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelData2Reply(data),
	}
	return reply, err
}
func (s *NovelService) DeleteNovel(ctx context.Context, req *pb.NovelDeleteReq) (*pb.NovelDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovel")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *NovelService) BatchDeleteNovel(ctx context.Context, req *pb.NovelBatchDeleteReq) (*pb.NovelDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovel")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
