package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelbuyrecord/v1"
)

type NovelBuyRecordService struct {
	pb.UnimplementedNovelBuyRecordServer
	uc  *biz.NovelBuyRecordUseCase
	log *log.Helper
}

func NewNovelBuyRecordService(uc *biz.NovelBuyRecordUseCase, logger log.Logger) *NovelBuyRecordService {
	return &NovelBuyRecordService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelBuyRecordService) GetNovelBuyRecordPage(ctx context.Context, req *pb.NovelBuyRecordPageReq) (*pb.NovelBuyRecordPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelBuyRecordPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelBuyRecordData, 0)
	for i := range datas {
		items = append(items, convert.NovelBuyRecordData2Reply(datas[i]))
	}
	reply := &pb.NovelBuyRecordPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *NovelBuyRecordService) GetNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordReq) (*pb.NovelBuyRecordReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelBuyRecord")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelBuyRecordReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelBuyRecordData2Reply(data),
	}
	return reply, nil
}
func (s *NovelBuyRecordService) UpdateNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordUpdateReq) (*pb.NovelBuyRecordUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelBuyRecord")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelBuyRecordUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelBuyRecordData2Reply(data),
	}
	return reply, nil
}
func (s *NovelBuyRecordService) CreateNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordCreateReq) (*pb.NovelBuyRecordCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelBuyRecord")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelBuyRecordCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelBuyRecordData2Reply(data),
	}
	return reply, err
}
func (s *NovelBuyRecordService) DeleteNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordDeleteReq) (*pb.NovelBuyRecordDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelBuyRecord")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelBuyRecordDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *NovelBuyRecordService) BatchDeleteNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordBatchDeleteReq) (*pb.NovelBuyRecordDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelBuyRecord")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelBuyRecordDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
