package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/assetitem/v1"
)

type AssetItemService struct {
	pb.UnimplementedAssetItemServer
	uc  *biz.AssetItemUseCase
	log *log.Helper
}

func NewAssetItemService(uc *biz.AssetItemUseCase, logger log.Logger) *AssetItemService {
	return &AssetItemService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AssetItemService) GetAssetItemPage(ctx context.Context, req *pb.AssetItemPageReq) (*pb.AssetItemPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetAssetItemPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.AssetItemData, 0)
	for i := range datas {
		items = append(items, convert.AssetItemData2Reply(datas[i]))
	}
	reply := &pb.AssetItemPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *AssetItemService) GetAssetItem(ctx context.Context, req *pb.AssetItemReq) (*pb.AssetItemReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetAssetItem")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AssetItemReply{
		Code:    200,
		Message: "success",
		Result:  convert.AssetItemData2Reply(data),
	}
	return reply, nil
}
func (s *AssetItemService) UpdateAssetItem(ctx context.Context, req *pb.AssetItemUpdateReq) (*pb.AssetItemUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateAssetItem")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AssetItemUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.AssetItemData2Reply(data),
	}
	return reply, nil
}
func (s *AssetItemService) CreateAssetItem(ctx context.Context, req *pb.AssetItemCreateReq) (*pb.AssetItemCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateAssetItem")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.AssetItemCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.AssetItemData2Reply(data),
	}
	return reply, err
}
func (s *AssetItemService) DeleteAssetItem(ctx context.Context, req *pb.AssetItemDeleteReq) (*pb.AssetItemDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteAssetItem")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AssetItemDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *AssetItemService) BatchDeleteAssetItem(ctx context.Context, req *pb.AssetItemBatchDeleteReq) (*pb.AssetItemDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteAssetItem")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AssetItemDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
