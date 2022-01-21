package service
import (
	"context"
	"go.opentelemetry.io/otel"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelchapter/v1"
)

type NovelChapterService struct {
	pb.UnimplementedNovelChapterServer
	uc  *biz.NovelChapterUseCase
	log *log.Helper
}

func NewNovelChapterService(uc *biz.NovelChapterUseCase, logger log.Logger) *NovelChapterService {
	return &NovelChapterService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelChapterService) GetPageNovelChapter(ctx context.Context, req *pb.NovelChapterPageReq) (*pb.NovelChapterPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageNovelChapter")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelChapterReply, 0)
	for i := range datas {
		items = append(items, convert.NovelChapterData2Reply(datas[i]))
	}
	reply := &pb.NovelChapterPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelChapterService) GetNovelChapter(ctx context.Context, req *pb.NovelChapterReq) (*pb.NovelChapterReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelChapter")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelChapterData2Reply(daya), err
}
func (s *NovelChapterService) UpdateNovelChapter(ctx context.Context, req *pb.NovelChapterUpdateReq) (*pb.NovelChapterUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelChapter")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelChapterData2UpdateReply(data), err
}
func (s *NovelChapterService) CreateNovelChapter(ctx context.Context, req *pb.NovelChapterCreateReq) (*pb.NovelChapterCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelChapter")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelChapterData2CreateReply(data), err
}
func (s *NovelChapterService) DeleteNovelChapter(ctx context.Context, req *pb.NovelChapterDeleteReq) (*pb.NovelChapterDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelChapter")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelChapterDeleteReply{Result: err == nil}, err
}
func (s *NovelChapterService) BatchDeleteNovelChapter(ctx context.Context, req *pb.NovelChapterBatchDeleteReq) (*pb.NovelChapterDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelChapter")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelChapterDeleteReply{Result: err == nil && num > 0}, err
}
