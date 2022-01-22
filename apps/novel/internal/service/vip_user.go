package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageVipUser")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetVipUser")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.VipUserData2Reply(daya), err
}
func (s *VipUserService) UpdateVipUser(ctx context.Context, req *pb.VipUserUpdateReq) (*pb.VipUserUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateVipUser")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.VipUserData2UpdateReply(data), err
}
func (s *VipUserService) CreateVipUser(ctx context.Context, req *pb.VipUserCreateReq) (*pb.VipUserCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateVipUser")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.VipUserData2CreateReply(data), err
}
func (s *VipUserService) DeleteVipUser(ctx context.Context, req *pb.VipUserDeleteReq) (*pb.VipUserDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteVipUser")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.VipUserDeleteReply{Result: err == nil}, err
}
func (s *VipUserService) BatchDeleteVipUser(ctx context.Context, req *pb.VipUserBatchDeleteReq) (*pb.VipUserDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteVipUser")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.VipUserDeleteReply{Result: err == nil && num > 0}, err
}
