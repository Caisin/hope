package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageResourceGroup")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
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
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetResourceGroup")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.ResourceGroupData2Reply(daya), err
}
func (s *ResourceGroupService) UpdateResourceGroup(ctx context.Context, req *pb.ResourceGroupUpdateReq) (*pb.ResourceGroupUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateResourceGroup")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.ResourceGroupData2UpdateReply(data), err
}
func (s *ResourceGroupService) CreateResourceGroup(ctx context.Context, req *pb.ResourceGroupCreateReq) (*pb.ResourceGroupCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateResourceGroup")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.ResourceGroupData2CreateReply(data), err
}
func (s *ResourceGroupService) DeleteResourceGroup(ctx context.Context, req *pb.ResourceGroupDeleteReq) (*pb.ResourceGroupDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteResourceGroup")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ResourceGroupDeleteReply{Result: err == nil}, err
}
func (s *ResourceGroupService) BatchDeleteResourceGroup(ctx context.Context, req *pb.ResourceGroupBatchDeleteReq) (*pb.ResourceGroupDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteResourceGroup")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ResourceGroupDeleteReply{Result: err == nil && num > 0}, err
}
