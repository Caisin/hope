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

func (s *UserResourceService) GetPageUserResource(ctx context.Context, req *pb.UserResourcePageReq) (*pb.UserResourcePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageUserResource")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.UserResourceReply, 0)
	for i := range datas {
		items = append(items, convert.UserResourceData2Reply(datas[i]))
	}
	reply := &pb.UserResourcePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *UserResourceService) GetUserResource(ctx context.Context, req *pb.UserResourceReq) (*pb.UserResourceReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetUserResource")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.UserResourceData2Reply(daya), err
}
func (s *UserResourceService) UpdateUserResource(ctx context.Context, req *pb.UserResourceUpdateReq) (*pb.UserResourceUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateUserResource")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.UserResourceData2UpdateReply(data), err
}
func (s *UserResourceService) CreateUserResource(ctx context.Context, req *pb.UserResourceCreateReq) (*pb.UserResourceCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateUserResource")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.UserResourceData2CreateReply(data), err
}
func (s *UserResourceService) DeleteUserResource(ctx context.Context, req *pb.UserResourceDeleteReq) (*pb.UserResourceDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteUserResource")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserResourceDeleteReply{Result: err == nil}, err
}
func (s *UserResourceService) BatchDeleteUserResource(ctx context.Context, req *pb.UserResourceBatchDeleteReq) (*pb.UserResourceDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteUserResource")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserResourceDeleteReply{Result: err == nil && num > 0}, err
}
