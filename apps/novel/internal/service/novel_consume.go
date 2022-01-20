package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelconsume/v1"
)

type NovelConsumeService struct {
	pb.UnimplementedNovelConsumeServer
	uc  *biz.NovelConsumeUseCase
	log *log.Helper
}

func NewNovelConsumeService(uc *biz.NovelConsumeUseCase, logger log.Logger) *NovelConsumeService {
	return &NovelConsumeService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelConsumeService) GetPageNovelConsume(ctx context.Context, req *pb.NovelConsumePageReq) (*pb.NovelConsumePageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.NovelConsumeReply, 0)
	for i := range datas {
		items = append(items, convert.NovelConsumeData2Reply(datas[i]))
	}
	reply := &pb.NovelConsumePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelConsumeService) GetNovelConsume(ctx context.Context, req *pb.NovelConsumeReq) (*pb.NovelConsumeReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.NovelConsumeData2Reply(daya), err
}
func (s *NovelConsumeService) UpdateNovelConsume(ctx context.Context, req *pb.NovelConsumeUpdateReq) (*pb.NovelConsumeUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.NovelConsumeData2UpdateReply(data), err
}
func (s *NovelConsumeService) CreateNovelConsume(ctx context.Context, req *pb.NovelConsumeCreateReq) (*pb.NovelConsumeCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.NovelConsumeData2CreateReply(data), err
}
func (s *NovelConsumeService) DeleteNovelConsume(ctx context.Context, req *pb.NovelConsumeDeleteReq) (*pb.NovelConsumeDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.NovelConsumeDeleteReply{Result: err == nil}, err
}
func (s *NovelConsumeService) BatchDeleteNovelConsume(ctx context.Context, req *pb.NovelConsumeBatchDeleteReq) (*pb.NovelConsumeDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.NovelConsumeDeleteReply{Result: err == nil && num > 0}, err
}
