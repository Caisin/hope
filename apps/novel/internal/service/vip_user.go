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

func (s *VipUserService) GetVipUserPage(ctx context.Context, req *pb.VipUserPageReq) (*pb.VipUserPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetVipUserPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.VipUserData, 0)
	for i := range datas {
		items = append(items, convert.VipUserData2Reply(datas[i]))
	}
	reply := &pb.VipUserPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *VipUserService) GetVipUser(ctx context.Context, req *pb.VipUserReq) (*pb.VipUserReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetVipUser")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.VipUserReply{
		Code:    200,
		Message: "success",
		Result:  convert.VipUserData2Reply(data),
	}
	return reply, nil
}
func (s *VipUserService) UpdateVipUser(ctx context.Context, req *pb.VipUserUpdateReq) (*pb.VipUserUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateVipUser")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.VipUserUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.VipUserData2Reply(data),
	}
	return reply, nil
}
func (s *VipUserService) CreateVipUser(ctx context.Context, req *pb.VipUserCreateReq) (*pb.VipUserCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateVipUser")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.VipUserCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.VipUserData2Reply(data),
	}
	return reply, err
}
func (s *VipUserService) DeleteVipUser(ctx context.Context, req *pb.VipUserDeleteReq) (*pb.VipUserDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteVipUser")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.VipUserDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *VipUserService) BatchDeleteVipUser(ctx context.Context, req *pb.VipUserBatchDeleteReq) (*pb.VipUserDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteVipUser")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.VipUserDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
