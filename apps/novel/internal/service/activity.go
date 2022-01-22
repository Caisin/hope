package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/activity/v1"
)

type ActivityService struct {
	pb.UnimplementedActivityServer
	uc  *biz.ActivityUseCase
	log *log.Helper
}

func NewActivityService(uc *biz.ActivityUseCase, logger log.Logger) *ActivityService {
	return &ActivityService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ActivityService) GetPageActivity(ctx context.Context, req *pb.ActivityPageReq) (*pb.ActivityPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageActivity")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.ActivityData, 0)
	for i := range datas {
		items = append(items, convert.ActivityData2Reply(datas[i]))
	}
	reply := &pb.ActivityPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *ActivityService) GetActivity(ctx context.Context, req *pb.ActivityReq) (*pb.ActivityReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetActivity")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ActivityReply{
		Code:    200,
		Message: "success",
		Result:  convert.ActivityData2Reply(data),
	}
	return reply, nil
}
func (s *ActivityService) UpdateActivity(ctx context.Context, req *pb.ActivityUpdateReq) (*pb.ActivityUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateActivity")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ActivityUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.ActivityData2Reply(data),
	}
	return reply, nil
}
func (s *ActivityService) CreateActivity(ctx context.Context, req *pb.ActivityCreateReq) (*pb.ActivityCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateActivity")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ActivityCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.ActivityData2Reply(data),
	}
	return reply, err
}
func (s *ActivityService) DeleteActivity(ctx context.Context, req *pb.ActivityDeleteReq) (*pb.ActivityDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteActivity")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ActivityDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *ActivityService) BatchDeleteActivity(ctx context.Context, req *pb.ActivityBatchDeleteReq) (*pb.ActivityDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteActivity")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ActivityDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
