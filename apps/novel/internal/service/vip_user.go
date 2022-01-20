package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/vipuser/v1"
)

type VipUserService struct {
	pb.UnimplementedVipUserServer
	uc  *biz.VipUserUseCase
	log *log.Helper
}

func NewVipUserService(uc *biz.VipUserUseCase, logger log.Logger) *VipUserService {
	return &VipUserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *VipUserService) GetPageVipUser(ctx context.Context, req *pb.VipUserPageReq) (*pb.VipUserPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.VipUserReply, 0)
	for i := range datas {
		items = append(items, convert.VipUserData2Reply(datas[i]))
	}
	reply := &pb.VipUserPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *VipUserService) GetVipUser(ctx context.Context, req *pb.VipUserReq) (*pb.VipUserReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.VipUserData2Reply(daya), err
}
func (s *VipUserService) UpdateVipUser(ctx context.Context, req *pb.VipUserUpdateReq) (*pb.VipUserUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.VipUserData2UpdateReply(data), err
}
func (s *VipUserService) CreateVipUser(ctx context.Context, req *pb.VipUserCreateReq) (*pb.VipUserCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.VipUserData2CreateReply(data), err
}
func (s *VipUserService) DeleteVipUser(ctx context.Context, req *pb.VipUserDeleteReq) (*pb.VipUserDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.VipUserDeleteReply{Result: err == nil}, err
}
func (s *VipUserService) BatchDeleteVipUser(ctx context.Context, req *pb.VipUserBatchDeleteReq) (*pb.VipUserDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.VipUserDeleteReply{Result: err == nil && num > 0}, err
}
