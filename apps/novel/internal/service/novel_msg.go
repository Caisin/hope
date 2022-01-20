package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelmsg/v1"
)

type NovelMsgService struct {
	pb.UnimplementedNovelMsgServer
	uc  *biz.NovelMsgUseCase
	log *log.Helper
}

func NewNovelMsgService(uc *biz.NovelMsgUseCase, logger log.Logger) *NovelMsgService {
	return &NovelMsgService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelMsgService) GetPageNovelMsg(ctx context.Context, req *pb.NovelMsgPageReq) (*pb.NovelMsgPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.NovelMsgReply, 0)
	for i := range datas {
		items = append(items, convert.NovelMsgData2Reply(datas[i]))
	}
	reply := &pb.NovelMsgPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelMsgService) GetNovelMsg(ctx context.Context, req *pb.NovelMsgReq) (*pb.NovelMsgReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.NovelMsgData2Reply(daya), err
}
func (s *NovelMsgService) UpdateNovelMsg(ctx context.Context, req *pb.NovelMsgUpdateReq) (*pb.NovelMsgUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.NovelMsgData2UpdateReply(data), err
}
func (s *NovelMsgService) CreateNovelMsg(ctx context.Context, req *pb.NovelMsgCreateReq) (*pb.NovelMsgCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.NovelMsgData2CreateReply(data), err
}
func (s *NovelMsgService) DeleteNovelMsg(ctx context.Context, req *pb.NovelMsgDeleteReq) (*pb.NovelMsgDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.NovelMsgDeleteReply{Result: err == nil}, err
}
func (s *NovelMsgService) BatchDeleteNovelMsg(ctx context.Context, req *pb.NovelMsgBatchDeleteReq) (*pb.NovelMsgDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.NovelMsgDeleteReply{Result: err == nil && num > 0}, err
}
