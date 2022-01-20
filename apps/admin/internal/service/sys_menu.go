package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysmenu/v1"
)

type SysMenuService struct {
	pb.UnimplementedSysMenuServer
	uc  *biz.SysMenuUseCase
	log *log.Helper
}

func NewSysMenuService(uc *biz.SysMenuUseCase, logger log.Logger) *SysMenuService {
	return &SysMenuService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysMenuService) GetPageSysMenu(ctx context.Context, req *pb.SysMenuPageReq) (*pb.SysMenuPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysMenuReply, 0)
	for i := range datas {
		items = append(items, convert.SysMenuData2Reply(datas[i]))
	}
	reply := &pb.SysMenuPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysMenuService) GetSysMenu(ctx context.Context, req *pb.SysMenuReq) (*pb.SysMenuReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysMenuData2Reply(daya), err
}
func (s *SysMenuService) UpdateSysMenu(ctx context.Context, req *pb.SysMenuUpdateReq) (*pb.SysMenuUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysMenuData2UpdateReply(data), err
}
func (s *SysMenuService) CreateSysMenu(ctx context.Context, req *pb.SysMenuCreateReq) (*pb.SysMenuCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysMenuData2CreateReply(data), err
}
func (s *SysMenuService) DeleteSysMenu(ctx context.Context, req *pb.SysMenuDeleteReq) (*pb.SysMenuDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysMenuDeleteReply{Result: err == nil}, err
}
func (s *SysMenuService) BatchDeleteSysMenu(ctx context.Context, req *pb.SysMenuBatchDeleteReq) (*pb.SysMenuDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysMenuDeleteReply{Result: err == nil && num > 0}, err
}
