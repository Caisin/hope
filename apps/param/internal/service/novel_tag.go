package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/noveltag/v1"
)

type NovelTagService struct {
	pb.UnimplementedNovelTagServer
	uc  *biz.NovelTagUseCase
	log *log.Helper
}

func NewNovelTagService(uc *biz.NovelTagUseCase, logger log.Logger) *NovelTagService {
	return &NovelTagService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelTagService) GetNovelTagPage(ctx context.Context, req *pb.NovelTagPageReq) (*pb.NovelTagPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelTagPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelTagData, 0)
	for i := range datas {
		items = append(items, convert.NovelTagData2Reply(datas[i]))
	}
	reply := &pb.NovelTagPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *NovelTagService) GetNovelTag(ctx context.Context, req *pb.NovelTagReq) (*pb.NovelTagReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelTag")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelTagReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelTagData2Reply(data),
	}
	return reply, nil
}
func (s *NovelTagService) UpdateNovelTag(ctx context.Context, req *pb.NovelTagUpdateReq) (*pb.NovelTagUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelTag")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelTagUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelTagData2Reply(data),
	}
	return reply, nil
}
func (s *NovelTagService) CreateNovelTag(ctx context.Context, req *pb.NovelTagCreateReq) (*pb.NovelTagCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelTag")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelTagCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelTagData2Reply(data),
	}
	return reply, err
}
func (s *NovelTagService) DeleteNovelTag(ctx context.Context, req *pb.NovelTagDeleteReq) (*pb.NovelTagDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelTag")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelTagDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *NovelTagService) BatchDeleteNovelTag(ctx context.Context, req *pb.NovelTagBatchDeleteReq) (*pb.NovelTagDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelTag")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelTagDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
