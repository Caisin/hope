package service
import (
	"context"
	"go.opentelemetry.io/otel"
	"github.com/go-kratos/kratos/v2/log"
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

func (s *NovelService) GetPageNovel(ctx context.Context, req *pb.NovelPageReq) (*pb.NovelPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageNovel")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.NovelReply, 0)
	for i := range datas {
		items = append(items, convert.NovelData2Reply(datas[i]))
	}
	reply := &pb.NovelPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelService) GetNovel(ctx context.Context, req *pb.NovelReq) (*pb.NovelReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetNovel")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelData2Reply(daya), err
}
func (s *NovelService) UpdateNovel(ctx context.Context, req *pb.NovelUpdateReq) (*pb.NovelUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateNovel")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelData2UpdateReply(data), err
}
func (s *NovelService) CreateNovel(ctx context.Context, req *pb.NovelCreateReq) (*pb.NovelCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateNovel")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.NovelData2CreateReply(data), err
}
func (s *NovelService) DeleteNovel(ctx context.Context, req *pb.NovelDeleteReq) (*pb.NovelDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteNovel")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelDeleteReply{Result: err == nil}, err
}
func (s *NovelService) BatchDeleteNovel(ctx context.Context, req *pb.NovelBatchDeleteReq) (*pb.NovelDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteNovel")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.NovelDeleteReply{Result: err == nil && num > 0}, err
}
