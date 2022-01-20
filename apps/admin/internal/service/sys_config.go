package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysconfig/v1"
)

type SysConfigService struct {
	pb.UnimplementedSysConfigServer
	uc  *biz.SysConfigUseCase
	log *log.Helper
}

func NewSysConfigService(uc *biz.SysConfigUseCase, logger log.Logger) *SysConfigService {
	return &SysConfigService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysConfigService) GetPageSysConfig(ctx context.Context, req *pb.SysConfigPageReq) (*pb.SysConfigPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysConfigReply, 0)
	for i := range datas {
		items = append(items, convert.SysConfigData2Reply(datas[i]))
	}
	reply := &pb.SysConfigPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysConfigService) GetSysConfig(ctx context.Context, req *pb.SysConfigReq) (*pb.SysConfigReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysConfigData2Reply(daya), err
}
func (s *SysConfigService) UpdateSysConfig(ctx context.Context, req *pb.SysConfigUpdateReq) (*pb.SysConfigUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysConfigData2UpdateReply(data), err
}
func (s *SysConfigService) CreateSysConfig(ctx context.Context, req *pb.SysConfigCreateReq) (*pb.SysConfigCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysConfigData2CreateReply(data), err
}
func (s *SysConfigService) DeleteSysConfig(ctx context.Context, req *pb.SysConfigDeleteReq) (*pb.SysConfigDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysConfigDeleteReply{Result: err == nil}, err
}
func (s *SysConfigService) BatchDeleteSysConfig(ctx context.Context, req *pb.SysConfigBatchDeleteReq) (*pb.SysConfigDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysConfigDeleteReply{Result: err == nil && num > 0}, err
}
