package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/adchannel/v1"
)

type AdChannelService struct {
	pb.UnimplementedAdChannelServer
	uc  *biz.AdChannelUseCase
	log *log.Helper
}

func NewAdChannelService(uc *biz.AdChannelUseCase, logger log.Logger) *AdChannelService {
	return &AdChannelService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AdChannelService) GetPageAdChannel(ctx context.Context, req *pb.AdChannelPageReq) (*pb.AdChannelPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageAdChannel")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.AdChannelData, 0)
	for i := range datas {
		items = append(items, convert.AdChannelData2Reply(datas[i]))
	}
	reply := &pb.AdChannelPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *AdChannelService) GetAdChannel(ctx context.Context, req *pb.AdChannelReq) (*pb.AdChannelReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetAdChannel")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AdChannelReply{
		Code:    200,
		Message: "success",
		Result:  convert.AdChannelData2Reply(data),
	}
	return reply, nil
}
func (s *AdChannelService) UpdateAdChannel(ctx context.Context, req *pb.AdChannelUpdateReq) (*pb.AdChannelUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateAdChannel")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AdChannelUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.AdChannelData2Reply(data),
	}
	return reply, nil
}
func (s *AdChannelService) CreateAdChannel(ctx context.Context, req *pb.AdChannelCreateReq) (*pb.AdChannelCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateAdChannel")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AdChannelCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.AdChannelData2Reply(data),
	}
	return reply, err
}
func (s *AdChannelService) DeleteAdChannel(ctx context.Context, req *pb.AdChannelDeleteReq) (*pb.AdChannelDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteAdChannel")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AdChannelDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *AdChannelService) BatchDeleteAdChannel(ctx context.Context, req *pb.AdChannelBatchDeleteReq) (*pb.AdChannelDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteAdChannel")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AdChannelDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
