package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysrole/v1"
)

type SysRoleService struct {
	pb.UnimplementedSysRoleServer
	uc  *biz.SysRoleUseCase
	log *log.Helper
}

func NewSysRoleService(uc *biz.SysRoleUseCase, logger log.Logger) *SysRoleService {
	return &SysRoleService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysRoleService) GetPageSysRole(ctx context.Context, req *pb.SysRolePageReq) (*pb.SysRolePageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysRoleReply, 0)
	for i := range datas {
		items = append(items, convert.SysRoleData2Reply(datas[i]))
	}
	reply := &pb.SysRolePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysRoleService) GetSysRole(ctx context.Context, req *pb.SysRoleReq) (*pb.SysRoleReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysRoleData2Reply(daya), err
}
func (s *SysRoleService) UpdateSysRole(ctx context.Context, req *pb.SysRoleUpdateReq) (*pb.SysRoleUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysRoleData2UpdateReply(data), err
}
func (s *SysRoleService) CreateSysRole(ctx context.Context, req *pb.SysRoleCreateReq) (*pb.SysRoleCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysRoleData2CreateReply(data), err
}
func (s *SysRoleService) DeleteSysRole(ctx context.Context, req *pb.SysRoleDeleteReq) (*pb.SysRoleDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysRoleDeleteReply{Result: err == nil}, err
}
func (s *SysRoleService) BatchDeleteSysRole(ctx context.Context, req *pb.SysRoleBatchDeleteReq) (*pb.SysRoleDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysRoleDeleteReply{Result: err == nil && num > 0}, err
}
