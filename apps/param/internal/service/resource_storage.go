package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.ResourceStorageReply, 0)
	for i := range datas {
		items = append(items, convert.ResourceStorageData2Reply(datas[i]))
	}
	reply := &pb.ResourceStoragePageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *ResourceStorageService) GetResourceStorage(ctx context.Context, req *pb.ResourceStorageReq) (*pb.ResourceStorageReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.ResourceStorageData2Reply(daya), err
}
func (s *ResourceStorageService) UpdateResourceStorage(ctx context.Context, req *pb.ResourceStorageUpdateReq) (*pb.ResourceStorageUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.ResourceStorageData2UpdateReply(data), err
}
func (s *ResourceStorageService) CreateResourceStorage(ctx context.Context, req *pb.ResourceStorageCreateReq) (*pb.ResourceStorageCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.ResourceStorageData2CreateReply(data), err
}
func (s *ResourceStorageService) DeleteResourceStorage(ctx context.Context, req *pb.ResourceStorageDeleteReq) (*pb.ResourceStorageDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.ResourceStorageDeleteReply{Result: err == nil}, err
}
func (s *ResourceStorageService) BatchDeleteResourceStorage(ctx context.Context, req *pb.ResourceStorageBatchDeleteReq) (*pb.ResourceStorageDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.ResourceStorageDeleteReply{Result: err == nil && num > 0}, err
}
