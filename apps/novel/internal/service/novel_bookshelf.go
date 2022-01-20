package service

import (
	"context"
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
	datas, err := s.uc.Page(ctx, req)
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
	daya, err := s.uc.Get(ctx, req)
	return convert.NovelBookshelfData2Reply(daya), err
}
func (s *NovelBookshelfService) UpdateNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfUpdateReq) (*pb.NovelBookshelfUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.NovelBookshelfData2UpdateReply(data), err
}
func (s *NovelBookshelfService) CreateNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfCreateReq) (*pb.NovelBookshelfCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.NovelBookshelfData2CreateReply(data), err
}
func (s *NovelBookshelfService) DeleteNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfDeleteReq) (*pb.NovelBookshelfDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.NovelBookshelfDeleteReply{Result: err == nil}, err
}
func (s *NovelBookshelfService) BatchDeleteNovelBookshelf(ctx context.Context, req *pb.NovelBookshelfBatchDeleteReq) (*pb.NovelBookshelfDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.NovelBookshelfDeleteReply{Result: err == nil && num > 0}, err
}
