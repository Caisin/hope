package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/resourcestorage/v1"
)

type ResourceStorageService struct {
	pb.UnimplementedResourceStorageServer
	uc  *biz.ResourceStorageUseCase
	log *log.Helper
}

func NewResourceStorageService(uc *biz.ResourceStorageUseCase, logger log.Logger) *ResourceStorageService {
	return &ResourceStorageService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ResourceStorageService) GetPageResourceStorage(ctx context.Context, req *pb.ResourceStoragePageReq) (*pb.ResourceStoragePageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageResourceStorage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.ResourceStorageData, 0)
	for i := range datas {
		items = append(items, convert.ResourceStorageData2Reply(datas[i]))
	}
	reply := &pb.ResourceStoragePageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *ResourceStorageService) GetResourceStorage(ctx context.Context, req *pb.ResourceStorageReq) (*pb.ResourceStorageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetResourceStorage")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ResourceStorageReply{
		Code:    200,
		Message: "success",
		Result:  convert.ResourceStorageData2Reply(data),
	}
	return reply, nil
}
func (s *ResourceStorageService) UpdateResourceStorage(ctx context.Context, req *pb.ResourceStorageUpdateReq) (*pb.ResourceStorageUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateResourceStorage")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ResourceStorageUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.ResourceStorageData2Reply(data),
	}
	return reply, nil
}
func (s *ResourceStorageService) CreateResourceStorage(ctx context.Context, req *pb.ResourceStorageCreateReq) (*pb.ResourceStorageCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateResourceStorage")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.ResourceStorageCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.ResourceStorageData2Reply(data),
	}
	return reply, err
}
func (s *ResourceStorageService) DeleteResourceStorage(ctx context.Context, req *pb.ResourceStorageDeleteReq) (*pb.ResourceStorageDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteResourceStorage")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ResourceStorageDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *ResourceStorageService) BatchDeleteResourceStorage(ctx context.Context, req *pb.ResourceStorageBatchDeleteReq) (*pb.ResourceStorageDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteResourceStorage")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ResourceStorageDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
