package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/resourcegroup/v1"
)

type ResourceGroupService struct {
	pb.UnimplementedResourceGroupServer
	uc  *biz.ResourceGroupUseCase
	log *log.Helper
}

func NewResourceGroupService(uc *biz.ResourceGroupUseCase, logger log.Logger) *ResourceGroupService {
	return &ResourceGroupService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ResourceGroupService) GetPageResourceGroup(ctx context.Context, req *pb.ResourceGroupPageReq) (*pb.ResourceGroupPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.ResourceGroupReply, 0)
	for i := range datas {
		items = append(items, convert.ResourceGroupData2Reply(datas[i]))
	}
	reply := &pb.ResourceGroupPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *ResourceGroupService) GetResourceGroup(ctx context.Context, req *pb.ResourceGroupReq) (*pb.ResourceGroupReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.ResourceGroupData2Reply(daya), err
}
func (s *ResourceGroupService) UpdateResourceGroup(ctx context.Context, req *pb.ResourceGroupUpdateReq) (*pb.ResourceGroupUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.ResourceGroupData2UpdateReply(data), err
}
func (s *ResourceGroupService) CreateResourceGroup(ctx context.Context, req *pb.ResourceGroupCreateReq) (*pb.ResourceGroupCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.ResourceGroupData2CreateReply(data), err
}
func (s *ResourceGroupService) DeleteResourceGroup(ctx context.Context, req *pb.ResourceGroupDeleteReq) (*pb.ResourceGroupDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.ResourceGroupDeleteReply{Result: err == nil}, err
}
func (s *ResourceGroupService) BatchDeleteResourceGroup(ctx context.Context, req *pb.ResourceGroupBatchDeleteReq) (*pb.ResourceGroupDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.ResourceGroupDeleteReply{Result: err == nil && num > 0}, err
}
