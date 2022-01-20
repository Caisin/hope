package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/pageconfig/v1"
)

type PageConfigService struct {
	pb.UnimplementedPageConfigServer
	uc  *biz.PageConfigUseCase
	log *log.Helper
}

func NewPageConfigService(uc *biz.PageConfigUseCase, logger log.Logger) *PageConfigService {
	return &PageConfigService{uc: uc, log: log.NewHelper(logger)}
}

func (s *PageConfigService) GetPagePageConfig(ctx context.Context, req *pb.PageConfigPageReq) (*pb.PageConfigPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.PageConfigReply, 0)
	for i := range datas {
		items = append(items, convert.PageConfigData2Reply(datas[i]))
	}
	reply := &pb.PageConfigPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *PageConfigService) GetPageConfig(ctx context.Context, req *pb.PageConfigReq) (*pb.PageConfigReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.PageConfigData2Reply(daya), err
}
func (s *PageConfigService) UpdatePageConfig(ctx context.Context, req *pb.PageConfigUpdateReq) (*pb.PageConfigUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.PageConfigData2UpdateReply(data), err
}
func (s *PageConfigService) CreatePageConfig(ctx context.Context, req *pb.PageConfigCreateReq) (*pb.PageConfigCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.PageConfigData2CreateReply(data), err
}
func (s *PageConfigService) DeletePageConfig(ctx context.Context, req *pb.PageConfigDeleteReq) (*pb.PageConfigDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.PageConfigDeleteReply{Result: err == nil}, err
}
func (s *PageConfigService) BatchDeletePageConfig(ctx context.Context, req *pb.PageConfigBatchDeleteReq) (*pb.PageConfigDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.PageConfigDeleteReply{Result: err == nil && num > 0}, err
}
