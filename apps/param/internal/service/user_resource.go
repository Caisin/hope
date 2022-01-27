package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/userresource/v1"
)

type UserResourceService struct {
	pb.UnimplementedUserResourceServer
	uc  *biz.UserResourceUseCase
	log *log.Helper
}

func NewUserResourceService(uc *biz.UserResourceUseCase, logger log.Logger) *UserResourceService {
	return &UserResourceService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserResourceService) GetUserResourcePage(ctx context.Context, req *pb.UserResourcePageReq) (*pb.UserResourcePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserResourcePage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.UserResourceData, 0)
	for i := range datas {
		items = append(items, convert.UserResourceData2Reply(datas[i]))
	}
	reply := &pb.UserResourcePageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *UserResourceService) GetUserResource(ctx context.Context, req *pb.UserResourceReq) (*pb.UserResourceReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserResource")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserResourceReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserResourceData2Reply(data),
	}
	return reply, nil
}
func (s *UserResourceService) UpdateUserResource(ctx context.Context, req *pb.UserResourceUpdateReq) (*pb.UserResourceUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateUserResource")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserResourceUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserResourceData2Reply(data),
	}
	return reply, nil
}
func (s *UserResourceService) CreateUserResource(ctx context.Context, req *pb.UserResourceCreateReq) (*pb.UserResourceCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateUserResource")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.UserResourceCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.UserResourceData2Reply(data),
	}
	return reply, err
}
func (s *UserResourceService) DeleteUserResource(ctx context.Context, req *pb.UserResourceDeleteReq) (*pb.UserResourceDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteUserResource")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserResourceDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *UserResourceService) BatchDeleteUserResource(ctx context.Context, req *pb.UserResourceBatchDeleteReq) (*pb.UserResourceDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteUserResource")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserResourceDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
