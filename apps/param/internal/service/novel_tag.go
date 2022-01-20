package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (s *NovelTagService) GetPageNovelTag(ctx context.Context, req *pb.NovelTagPageReq) (*pb.NovelTagPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.NovelTagReply, 0)
	for i := range datas {
		items = append(items, convert.NovelTagData2Reply(datas[i]))
	}
	reply := &pb.NovelTagPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelTagService) GetNovelTag(ctx context.Context, req *pb.NovelTagReq) (*pb.NovelTagReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.NovelTagData2Reply(daya), err
}
func (s *NovelTagService) UpdateNovelTag(ctx context.Context, req *pb.NovelTagUpdateReq) (*pb.NovelTagUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.NovelTagData2UpdateReply(data), err
}
func (s *NovelTagService) CreateNovelTag(ctx context.Context, req *pb.NovelTagCreateReq) (*pb.NovelTagCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.NovelTagData2CreateReply(data), err
}
func (s *NovelTagService) DeleteNovelTag(ctx context.Context, req *pb.NovelTagDeleteReq) (*pb.NovelTagDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.NovelTagDeleteReply{Result: err == nil}, err
}
func (s *NovelTagService) BatchDeleteNovelTag(ctx context.Context, req *pb.NovelTagBatchDeleteReq) (*pb.NovelTagDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.NovelTagDeleteReply{Result: err == nil && num > 0}, err
}
