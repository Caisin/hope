package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelconsume/v1"
)

type NovelConsumeService struct {
	pb.UnimplementedNovelConsumeServer
	uc  *biz.NovelConsumeUseCase
	log *log.Helper
}

func NewNovelConsumeService(uc *biz.NovelConsumeUseCase, logger log.Logger) *NovelConsumeService {
	return &NovelConsumeService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelConsumeService) GetNovelConsumePage(ctx context.Context, req *pb.NovelConsumePageReq) (*pb.NovelConsumePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelConsumePage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelConsumeData, 0)
	for i := range datas {
		items = append(items, convert.NovelConsumeData2Reply(datas[i]))
	}
	reply := &pb.NovelConsumePageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *NovelConsumeService) GetNovelConsume(ctx context.Context, req *pb.NovelConsumeReq) (*pb.NovelConsumeReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelConsume")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelConsumeReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelConsumeData2Reply(data),
	}
	return reply, nil
}
func (s *NovelConsumeService) UpdateNovelConsume(ctx context.Context, req *pb.NovelConsumeUpdateReq) (*pb.NovelConsumeUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelConsume")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelConsumeUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelConsumeData2Reply(data),
	}
	return reply, nil
}
func (s *NovelConsumeService) CreateNovelConsume(ctx context.Context, req *pb.NovelConsumeCreateReq) (*pb.NovelConsumeCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelConsume")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelConsumeCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelConsumeData2Reply(data),
	}
	return reply, err
}
func (s *NovelConsumeService) DeleteNovelConsume(ctx context.Context, req *pb.NovelConsumeDeleteReq) (*pb.NovelConsumeDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelConsume")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelConsumeDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *NovelConsumeService) BatchDeleteNovelConsume(ctx context.Context, req *pb.NovelConsumeBatchDeleteReq) (*pb.NovelConsumeDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelConsume")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelConsumeDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
