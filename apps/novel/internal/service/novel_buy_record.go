package service
import (
	"context"
	"go.opentelemetry.io/otel"
	"github.com/go-kratos/kratos/v2/log"
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

func (s *NovelBuyRecordService) GetPageNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordPageReq) (*pb.NovelBuyRecordPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageNovelBuyRecord")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelBuyRecordReply, 0)
	for i := range datas {
		items = append(items, convert.NovelBuyRecordData2Reply(datas[i]))
	}
	reply := &pb.NovelBuyRecordPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelBuyRecordService) GetNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordReq) (*pb.NovelBuyRecordReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelBuyRecord")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelBuyRecordData2Reply(daya), err
}
func (s *NovelBuyRecordService) UpdateNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordUpdateReq) (*pb.NovelBuyRecordUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelBuyRecord")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelBuyRecordData2UpdateReply(data), err
}
func (s *NovelBuyRecordService) CreateNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordCreateReq) (*pb.NovelBuyRecordCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelBuyRecord")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelBuyRecordData2CreateReply(data), err
}
func (s *NovelBuyRecordService) DeleteNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordDeleteReq) (*pb.NovelBuyRecordDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelBuyRecord")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelBuyRecordDeleteReply{Result: err == nil}, err
}
func (s *NovelBuyRecordService) BatchDeleteNovelBuyRecord(ctx context.Context, req *pb.NovelBuyRecordBatchDeleteReq) (*pb.NovelBuyRecordDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelBuyRecord")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelBuyRecordDeleteReply{Result: err == nil && num > 0}, err
}
