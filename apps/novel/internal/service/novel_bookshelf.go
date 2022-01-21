package service
import (
	"context"
	"go.opentelemetry.io/otel"
	"github.com/go-kratos/kratos/v2/log"
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

func (s *NovelBookshelfService) GetPageNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfPageReq) (*pb.NovelBookshelfPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageNovelBookshelf")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelBookshelfReply, 0)
	for i := range datas {
		items = append(items, convert.NovelBookshelfData2Reply(datas[i]))
	}
	reply := &pb.NovelBookshelfPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelBookshelfService) GetNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfReq) (*pb.NovelBookshelfReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovelBookshelf")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelBookshelfData2Reply(daya), err
}
func (s *NovelBookshelfService) UpdateNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfUpdateReq) (*pb.NovelBookshelfUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovelBookshelf")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelBookshelfData2UpdateReply(data), err
}
func (s *NovelBookshelfService) CreateNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfCreateReq) (*pb.NovelBookshelfCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovelBookshelf")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelBookshelfData2CreateReply(data), err
}
func (s *NovelBookshelfService) DeleteNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfDeleteReq) (*pb.NovelBookshelfDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovelBookshelf")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelBookshelfDeleteReply{Result: err == nil}, err
}
func (s *NovelBookshelfService) BatchDeleteNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfBatchDeleteReq) (*pb.NovelBookshelfDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovelBookshelf")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelBookshelfDeleteReply{Result: err == nil && num > 0}, err
}
