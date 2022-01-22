package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/userresourcerecord/v1"
)

type UserResourceRecordService struct {
	pb.UnimplementedUserResourceRecordServer
	uc  *biz.UserResourceRecordUseCase
	log *log.Helper
}

func NewUserResourceRecordService(uc *biz.UserResourceRecordUseCase, logger log.Logger) *UserResourceRecordService {
	return &UserResourceRecordService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserResourceRecordService) GetPageUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordPageReq) (*pb.UserResourceRecordPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageUserResourceRecord")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.UserResourceRecordData, 0)
	for i := range datas {
		items = append(items, convert.UserResourceRecordData2Reply(datas[i]))
	}
	reply := &pb.UserResourceRecordPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *UserResourceRecordService) GetUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordReq) (*pb.UserResourceRecordReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserResourceRecord")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserResourceRecordReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserResourceRecordData2Reply(data),
	}
	return reply, nil
}
func (s *UserResourceRecordService) UpdateUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordUpdateReq) (*pb.UserResourceRecordUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateUserResourceRecord")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserResourceRecordUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserResourceRecordData2Reply(data),
	}
	return reply, nil
}
func (s *UserResourceRecordService) CreateUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordCreateReq) (*pb.UserResourceRecordCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateUserResourceRecord")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserResourceRecordCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserResourceRecordData2Reply(data),
	}
	return reply, err
}
func (s *UserResourceRecordService) DeleteUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordDeleteReq) (*pb.UserResourceRecordDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteUserResourceRecord")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserResourceRecordDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *UserResourceRecordService) BatchDeleteUserResourceRecord(ctx context.Context, req *pb.UserResourceRecordBatchDeleteReq) (*pb.UserResourceRecordDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteUserResourceRecord")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserResourceRecordDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
