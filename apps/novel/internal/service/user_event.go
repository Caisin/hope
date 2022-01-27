package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/userevent/v1"
)

type UserEventService struct {
	pb.UnimplementedUserEventServer
	uc  *biz.UserEventUseCase
	log *log.Helper
}

func NewUserEventService(uc *biz.UserEventUseCase, logger log.Logger) *UserEventService {
	return &UserEventService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserEventService) GetUserEventPage(ctx context.Context, req *pb.UserEventPageReq) (*pb.UserEventPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserEventPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.UserEventData, 0)
	for i := range datas {
		items = append(items, convert.UserEventData2Reply(datas[i]))
	}
	reply := &pb.UserEventPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *UserEventService) GetUserEvent(ctx context.Context, req *pb.UserEventReq) (*pb.UserEventReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserEvent")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserEventReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserEventData2Reply(data),
	}
	return reply, nil
}
func (s *UserEventService) UpdateUserEvent(ctx context.Context, req *pb.UserEventUpdateReq) (*pb.UserEventUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateUserEvent")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserEventUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserEventData2Reply(data),
	}
	return reply, nil
}
func (s *UserEventService) CreateUserEvent(ctx context.Context, req *pb.UserEventCreateReq) (*pb.UserEventCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateUserEvent")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserEventCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserEventData2Reply(data),
	}
	return reply, err
}
func (s *UserEventService) DeleteUserEvent(ctx context.Context, req *pb.UserEventDeleteReq) (*pb.UserEventDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteUserEvent")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserEventDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *UserEventService) BatchDeleteUserEvent(ctx context.Context, req *pb.UserEventBatchDeleteReq) (*pb.UserEventDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteUserEvent")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserEventDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
