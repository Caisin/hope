package service

import (
	"context"
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
	datas, err := s.uc.Page(ctx, req)
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
	daya, err := s.uc.Get(ctx, req)
	return convert.NovelData2Reply(daya), err
}
func (s *NovelService) UpdateNovel(ctx context.Context, req *pb.NovelUpdateReq) (*pb.NovelUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.NovelData2UpdateReply(data), err
}
func (s *NovelService) CreateNovel(ctx context.Context, req *pb.NovelCreateReq) (*pb.NovelCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.NovelData2CreateReply(data), err
}
func (s *NovelService) DeleteNovel(ctx context.Context, req *pb.NovelDeleteReq) (*pb.NovelDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.NovelDeleteReply{Result: err == nil}, err
}
func (s *NovelService) BatchDeleteNovel(ctx context.Context, req *pb.NovelBatchDeleteReq) (*pb.NovelDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.NovelDeleteReply{Result: err == nil && num > 0}, err
}
