package service

import (
	"context"
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
	datas, err := s.uc.Page(ctx, req)
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
	daya, err := s.uc.Get(ctx, req)
	return convert.NovelChapterData2Reply(daya), err
}
func (s *NovelChapterService) UpdateNovelChapter(ctx context.Context, req *pb.NovelChapterUpdateReq) (*pb.NovelChapterUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.NovelChapterData2UpdateReply(data), err
}
func (s *NovelChapterService) CreateNovelChapter(ctx context.Context, req *pb.NovelChapterCreateReq) (*pb.NovelChapterCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.NovelChapterData2CreateReply(data), err
}
func (s *NovelChapterService) DeleteNovelChapter(ctx context.Context, req *pb.NovelChapterDeleteReq) (*pb.NovelChapterDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.NovelChapterDeleteReply{Result: err == nil}, err
}
func (s *NovelChapterService) BatchDeleteNovelChapter(ctx context.Context, req *pb.NovelChapterBatchDeleteReq) (*pb.NovelChapterDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.NovelChapterDeleteReply{Result: err == nil && num > 0}, err
}
