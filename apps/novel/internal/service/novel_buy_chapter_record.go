package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/novelbuychapterrecord/v1"
)

type NovelBuyChapterRecordService struct {
	pb.UnimplementedNovelBuyChapterRecordServer
	uc  *biz.NovelBuyChapterRecordUseCase
	log *log.Helper
}

func NewNovelBuyChapterRecordService(uc *biz.NovelBuyChapterRecordUseCase, logger log.Logger) *NovelBuyChapterRecordService {
	return &NovelBuyChapterRecordService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelBuyChapterRecordService) GetPageNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordPageReq) (*pb.NovelBuyChapterRecordPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.NovelBuyChapterRecordReply, 0)
	for i := range datas {
		items = append(items, convert.NovelBuyChapterRecordData2Reply(datas[i]))
	}
	reply := &pb.NovelBuyChapterRecordPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelBuyChapterRecordService) GetNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordReq) (*pb.NovelBuyChapterRecordReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.NovelBuyChapterRecordData2Reply(daya), err
}
func (s *NovelBuyChapterRecordService) UpdateNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordUpdateReq) (*pb.NovelBuyChapterRecordUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.NovelBuyChapterRecordData2UpdateReply(data), err
}
func (s *NovelBuyChapterRecordService) CreateNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordCreateReq) (*pb.NovelBuyChapterRecordCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.NovelBuyChapterRecordData2CreateReply(data), err
}
func (s *NovelBuyChapterRecordService) DeleteNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordDeleteReq) (*pb.NovelBuyChapterRecordDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.NovelBuyChapterRecordDeleteReply{Result: err == nil}, err
}
func (s *NovelBuyChapterRecordService) BatchDeleteNovelBuyChapterRecord(ctx context.Context, req *pb.NovelBuyChapterRecordBatchDeleteReq) (*pb.NovelBuyChapterRecordDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.NovelBuyChapterRecordDeleteReply{Result: err == nil && num > 0}, err
}
