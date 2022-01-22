package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelmsg/v1"
)

type NovelMsgService struct {
	pb.UnimplementedNovelMsgServer
	uc  *biz.NovelMsgUseCase
	log *log.Helper
}

func NewNovelMsgService(uc *biz.NovelMsgUseCase, logger log.Logger) *NovelMsgService {
	return &NovelMsgService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelMsgService) GetPageNovelMsg(ctx context.Context, req *pb.NovelMsgPageReq) (*pb.NovelMsgPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageNovelMsg")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelMsgData, 0)
	for i := range datas {
		items = append(items, convert.NovelMsgData2Reply(datas[i]))
	}
	reply := &pb.NovelMsgPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *NovelMsgService) GetNovelMsg(ctx context.Context, req *pb.NovelMsgReq) (*pb.NovelMsgReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelMsg")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelMsgReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelMsgData2Reply(data),
	}
	return reply, nil
}
func (s *NovelMsgService) UpdateNovelMsg(ctx context.Context, req *pb.NovelMsgUpdateReq) (*pb.NovelMsgUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelMsg")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelMsgUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelMsgData2Reply(data),
	}
	return reply, nil
}
func (s *NovelMsgService) CreateNovelMsg(ctx context.Context, req *pb.NovelMsgCreateReq) (*pb.NovelMsgCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelMsg")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelMsgCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelMsgData2Reply(data),
	}
	return reply, err
}
func (s *NovelMsgService) DeleteNovelMsg(ctx context.Context, req *pb.NovelMsgDeleteReq) (*pb.NovelMsgDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelMsg")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelMsgDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *NovelMsgService) BatchDeleteNovelMsg(ctx context.Context, req *pb.NovelMsgBatchDeleteReq) (*pb.NovelMsgDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelMsg")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelMsgDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
