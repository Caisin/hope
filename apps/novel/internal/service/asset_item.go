package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (s *AssetItemService) GetPageAssetItem(ctx context.Context, req *pb.AssetItemPageReq) (*pb.AssetItemPageReply, error) {
	datas, err := s.uc.Page(ctx, req)
	items := make([]*pb.AssetItemReply, 0)
	for i := range datas {
		items = append(items, convert.AssetItemData2Reply(datas[i]))
	}
	reply := &pb.AssetItemPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *AssetItemService) GetAssetItem(ctx context.Context, req *pb.AssetItemReq) (*pb.AssetItemReply, error) {
	daya, err := s.uc.Get(ctx, req)
	return convert.AssetItemData2Reply(daya), err
}
func (s *AssetItemService) UpdateAssetItem(ctx context.Context, req *pb.AssetItemUpdateReq) (*pb.AssetItemUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.AssetItemData2UpdateReply(data), err
}
func (s *AssetItemService) CreateAssetItem(ctx context.Context, req *pb.AssetItemCreateReq) (*pb.AssetItemCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.AssetItemData2CreateReply(data), err
}
func (s *AssetItemService) DeleteAssetItem(ctx context.Context, req *pb.AssetItemDeleteReq) (*pb.AssetItemDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.AssetItemDeleteReply{Result: err == nil}, err
}
func (s *AssetItemService) BatchDeleteAssetItem(ctx context.Context, req *pb.AssetItemBatchDeleteReq) (*pb.AssetItemDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.AssetItemDeleteReply{Result: err == nil && num > 0}, err
}
