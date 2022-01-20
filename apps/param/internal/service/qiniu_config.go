package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/qiniuconfig/v1"
)

type QiniuConfigService struct {
	pb.UnimplementedQiniuConfigServer
	uc  *biz.QiniuConfigUseCase
	log *log.Helper
}

func NewQiniuConfigService(uc *biz.QiniuConfigUseCase, logger log.Logger) *QiniuConfigService {
	return &QiniuConfigService{uc: uc, log: log.NewHelper(logger)}
}

func (s *QiniuConfigService) GetPageQiniuConfig(ctx context.Context, req *pb.QiniuConfigPageReq) (*pb.QiniuConfigPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.QiniuConfigReply, 0)
	for i := range datas {
		items = append(items, convert.QiniuConfigData2Reply(datas[i]))
	}
	reply := &pb.QiniuConfigPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *QiniuConfigService) GetQiniuConfig(ctx context.Context, req *pb.QiniuConfigReq) (*pb.QiniuConfigReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.QiniuConfigData2Reply(daya), err
}
func (s *QiniuConfigService) UpdateQiniuConfig(ctx context.Context, req *pb.QiniuConfigUpdateReq) (*pb.QiniuConfigUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.QiniuConfigData2UpdateReply(data), err
}
func (s *QiniuConfigService) CreateQiniuConfig(ctx context.Context, req *pb.QiniuConfigCreateReq) (*pb.QiniuConfigCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.QiniuConfigData2CreateReply(data), err
}
func (s *QiniuConfigService) DeleteQiniuConfig(ctx context.Context, req *pb.QiniuConfigDeleteReq) (*pb.QiniuConfigDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.QiniuConfigDeleteReply{Result: err == nil}, err
}
func (s *QiniuConfigService) BatchDeleteQiniuConfig(ctx context.Context, req *pb.QiniuConfigBatchDeleteReq) (*pb.QiniuConfigDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.QiniuConfigDeleteReply{Result: err == nil && num > 0}, err
}
