package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysuser/v1"
)

type SysUserService struct {
	pb.UnimplementedSysUserServer
	uc  *biz.SysUserUseCase
	log *log.Helper
}

func NewSysUserService(uc *biz.SysUserUseCase, logger log.Logger) *SysUserService {
	return &SysUserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysUserService) GetPageSysUser(ctx context.Context, req *pb.SysUserPageReq) (*pb.SysUserPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.SysUserReply, 0)
	for i := range datas {
		items = append(items, convert.SysUserData2Reply(datas[i]))
	}
	reply := &pb.SysUserPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysUserService) GetSysUser(ctx context.Context, req *pb.SysUserReq) (*pb.SysUserReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.SysUserData2Reply(daya), err
}
func (s *SysUserService) UpdateSysUser(ctx context.Context, req *pb.SysUserUpdateReq) (*pb.SysUserUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.SysUserData2UpdateReply(data), err
}
func (s *SysUserService) CreateSysUser(ctx context.Context, req *pb.SysUserCreateReq) (*pb.SysUserCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.SysUserData2CreateReply(data), err
}
func (s *SysUserService) DeleteSysUser(ctx context.Context, req *pb.SysUserDeleteReq) (*pb.SysUserDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.SysUserDeleteReply{Result: err == nil}, err
}
func (s *SysUserService) BatchDeleteSysUser(ctx context.Context, req *pb.SysUserBatchDeleteReq) (*pb.SysUserDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.SysUserDeleteReply{Result: err == nil && num > 0}, err
}
