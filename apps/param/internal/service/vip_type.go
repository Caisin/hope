package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/viptype/v1"
)

type VipTypeService struct {
	pb.UnimplementedVipTypeServer
	uc  *biz.VipTypeUseCase
	log *log.Helper
}

func NewVipTypeService(uc *biz.VipTypeUseCase, logger log.Logger) *VipTypeService {
	return &VipTypeService{uc: uc, log: log.NewHelper(logger)}
}

func (s *VipTypeService) GetPageVipType(ctx context.Context, req *pb.VipTypePageReq) (*pb.VipTypePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageVipType")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.VipTypeReply, 0)
	for i := range datas {
		items = append(items, convert.VipTypeData2Reply(datas[i]))
	}
	reply := &pb.VipTypePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *VipTypeService) GetVipType(ctx context.Context, req *pb.VipTypeReq) (*pb.VipTypeReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetVipType")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.VipTypeData2Reply(daya), err
}
func (s *VipTypeService) UpdateVipType(ctx context.Context, req *pb.VipTypeUpdateReq) (*pb.VipTypeUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateVipType")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.VipTypeData2UpdateReply(data), err
}
func (s *VipTypeService) CreateVipType(ctx context.Context, req *pb.VipTypeCreateReq) (*pb.VipTypeCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateVipType")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.VipTypeData2CreateReply(data), err
}
func (s *VipTypeService) DeleteVipType(ctx context.Context, req *pb.VipTypeDeleteReq) (*pb.VipTypeDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteVipType")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.VipTypeDeleteReply{Result: err == nil}, err
}
func (s *VipTypeService) BatchDeleteVipType(ctx context.Context, req *pb.VipTypeBatchDeleteReq) (*pb.VipTypeDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteVipType")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.VipTypeDeleteReply{Result: err == nil && num > 0}, err
}
