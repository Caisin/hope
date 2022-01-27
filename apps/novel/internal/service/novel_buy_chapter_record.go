package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelbuychapterrecord/v1"
)

type NovelBuyChapterRecordService struct {
	pb.UnimplementedNovelBuyChapterRecordServer
	uc  *biz.NovelBuyChapterRecordUseCase
	log *log.Helper
}

func NewNovelBuyChapterRecordService(uc *biz.NovelBuyChapterRecordUseCase, logger log.Logger) *NovelBuyChapterRecordService {
	return &NovelBuyChapterRecordService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelBuyChapterRecordService) GetNovelBuyChapterRecordPage(ctx context.Context, req *pb.NovelBuyChapterRecordPageReq) (*pb.NovelBuyChapterRecordPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelBuyChapterRecordPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelBuyChapterRecordData, 0)
	for i := range datas {
		items = append(items, convert.NovelBuyChapterRecordData2Reply(datas[i]))
	}
	reply := &pb.NovelBuyChapterRecordPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *NovelBuyChapterRecordService) GetNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordReq) (*pb.NovelBuyChapterRecordReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelBuyChapterRecord")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelBuyChapterRecordReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelBuyChapterRecordData2Reply(data),
	}
	return reply, nil
}
func (s *NovelBuyChapterRecordService) UpdateNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordUpdateReq) (*pb.NovelBuyChapterRecordUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelBuyChapterRecord")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelBuyChapterRecordUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelBuyChapterRecordData2Reply(data),
	}
	return reply, nil
}
func (s *NovelBuyChapterRecordService) CreateNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordCreateReq) (*pb.NovelBuyChapterRecordCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelBuyChapterRecord")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.NovelBuyChapterRecordCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.NovelBuyChapterRecordData2Reply(data),
	}
	return reply, err
}
func (s *NovelBuyChapterRecordService) DeleteNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordDeleteReq) (*pb.NovelBuyChapterRecordDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelBuyChapterRecord")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelBuyChapterRecordDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *NovelBuyChapterRecordService) BatchDeleteNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordBatchDeleteReq) (*pb.NovelBuyChapterRecordDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelBuyChapterRecord")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelBuyChapterRecordDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
