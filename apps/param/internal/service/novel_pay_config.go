package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/novelpayconfig/v1"
)

type NovelPayConfigService struct {
	pb.UnimplementedNovelPayConfigServer
	uc  *biz.NovelPayConfigUseCase
	log *log.Helper
}

func NewNovelPayConfigService(uc *biz.NovelPayConfigUseCase, logger log.Logger) *NovelPayConfigService {
	return &NovelPayConfigService{uc: uc, log: log.NewHelper(logger)}
}

func (s *NovelPayConfigService) GetPageNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigPageReq) (*pb.NovelPayConfigPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.NovelPayConfigReply, 0)
	for i := range datas {
		items = append(items, convert.NovelPayConfigData2Reply(datas[i]))
	}
	reply := &pb.NovelPayConfigPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *NovelPayConfigService) GetNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigReq) (*pb.NovelPayConfigReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.NovelPayConfigData2Reply(daya), err
}
func (s *NovelPayConfigService) UpdateNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigUpdateReq) (*pb.NovelPayConfigUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.NovelPayConfigData2UpdateReply(data), err
}
func (s *NovelPayConfigService) CreateNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigCreateReq) (*pb.NovelPayConfigCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.NovelPayConfigData2CreateReply(data), err
}
func (s *NovelPayConfigService) DeleteNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigDeleteReq) (*pb.NovelPayConfigDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.NovelPayConfigDeleteReply{Result: err == nil}, err
}
func (s *NovelPayConfigService) BatchDeleteNovelPayConfig(ctx context.Context, req *pb.NovelPayConfigBatchDeleteReq) (*pb.NovelPayConfigDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.NovelPayConfigDeleteReply{Result: err == nil && num > 0}, err
}
