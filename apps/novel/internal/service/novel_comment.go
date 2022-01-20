package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelcomment/v1"
)

type NovelCommentService struct {
	pb.UnimplementedNovelCommentServer
	uc  *biz.NovelCommentUseCase
	log *log.Helper
}

func NewNovelCommentService(uc *biz.NovelCommentUseCase, logger log.Logger) *NovelCommentService {
	return &NovelCommentService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelCommentService) GetPageNovelComment(ctx context.Context, req *pb.NovelCommentPageReq) (*pb.NovelCommentPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.NovelCommentReply, 0)
	for i := range datas {
		items = append(items, convert.NovelCommentData2Reply(datas[i]))
	}
	reply := &pb.NovelCommentPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelCommentService) GetNovelComment(ctx context.Context, req *pb.NovelCommentReq) (*pb.NovelCommentReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.NovelCommentData2Reply(daya), err
}
func (s *NovelCommentService) UpdateNovelComment(ctx context.Context, req *pb.NovelCommentUpdateReq) (*pb.NovelCommentUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.NovelCommentData2UpdateReply(data), err
}
func (s *NovelCommentService) CreateNovelComment(ctx context.Context, req *pb.NovelCommentCreateReq) (*pb.NovelCommentCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.NovelCommentData2CreateReply(data), err
}
func (s *NovelCommentService) DeleteNovelComment(ctx context.Context, req *pb.NovelCommentDeleteReq) (*pb.NovelCommentDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.NovelCommentDeleteReply{Result: err == nil}, err
}
func (s *NovelCommentService) BatchDeleteNovelComment(ctx context.Context, req *pb.NovelCommentBatchDeleteReq) (*pb.NovelCommentDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.NovelCommentDeleteReply{Result: err == nil && num > 0}, err
}
