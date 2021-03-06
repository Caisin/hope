package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/userconsume/v1"
)

type UserConsumeService struct {
	pb.UnimplementedUserConsumeServer
	uc  *biz.UserConsumeUseCase
	log *log.Helper
}

func NewUserConsumeService(uc *biz.UserConsumeUseCase, logger log.Logger) *UserConsumeService {
	return &UserConsumeService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserConsumeService) GetUserConsumePage(ctx context.Context, req *pb.UserConsumePageReq) (*pb.UserConsumePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserConsumePage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.UserConsumeData, 0)
	for i := range datas {
		items = append(items, convert.UserConsumeData2Reply(datas[i]))
	}
	reply := &pb.UserConsumePageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *UserConsumeService) GetUserConsume(ctx context.Context, req *pb.UserConsumeReq) (*pb.UserConsumeReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserConsume")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserConsumeReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserConsumeData2Reply(data),
	}
	return reply, nil
}
func (s *UserConsumeService) UpdateUserConsume(ctx context.Context, req *pb.UserConsumeUpdateReq) (*pb.UserConsumeUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateUserConsume")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserConsumeUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserConsumeData2Reply(data),
	}
	return reply, nil
}
func (s *UserConsumeService) CreateUserConsume(ctx context.Context, req *pb.UserConsumeCreateReq) (*pb.UserConsumeCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateUserConsume")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserConsumeCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserConsumeData2Reply(data),
	}
	return reply, err
}
func (s *UserConsumeService) DeleteUserConsume(ctx context.Context, req *pb.UserConsumeDeleteReq) (*pb.UserConsumeDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteUserConsume")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserConsumeDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *UserConsumeService) BatchDeleteUserConsume(ctx context.Context, req *pb.UserConsumeBatchDeleteReq) (*pb.UserConsumeDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteUserConsume")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserConsumeDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
