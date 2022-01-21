package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/assetchangelog/v1"
)

type AssetChangeLogService struct {
	pb.UnimplementedAssetChangeLogServer
	uc  *biz.AssetChangeLogUseCase
	log *log.Helper
}

func NewAssetChangeLogService(uc *biz.AssetChangeLogUseCase, logger log.Logger) *AssetChangeLogService {
	return &AssetChangeLogService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AssetChangeLogService) GetPageAssetChangeLog(ctx context.Context, req *pb.AssetChangeLogPageReq) (*pb.AssetChangeLogPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageAssetChangeLog")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.AssetChangeLogReply, 0)
	for i := range datas {
		items = append(items, convert.AssetChangeLogData2Reply(datas[i]))
	}
	reply := &pb.AssetChangeLogPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *AssetChangeLogService) GetAssetChangeLog(ctx context.Context, req *pb.AssetChangeLogReq) (*pb.AssetChangeLogReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetAssetChangeLog")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.AssetChangeLogData2Reply(daya), err
}
func (s *AssetChangeLogService) UpdateAssetChangeLog(ctx context.Context, req *pb.AssetChangeLogUpdateReq) (*pb.AssetChangeLogUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateAssetChangeLog")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.AssetChangeLogData2UpdateReply(data), err
}
func (s *AssetChangeLogService) CreateAssetChangeLog(ctx context.Context, req *pb.AssetChangeLogCreateReq) (*pb.AssetChangeLogCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateAssetChangeLog")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.AssetChangeLogData2CreateReply(data), err
}
func (s *AssetChangeLogService) DeleteAssetChangeLog(ctx context.Context, req *pb.AssetChangeLogDeleteReq) (*pb.AssetChangeLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteAssetChangeLog")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AssetChangeLogDeleteReply{Result: err == nil}, err
}
func (s *AssetChangeLogService) BatchDeleteAssetChangeLog(ctx context.Context, req *pb.AssetChangeLogBatchDeleteReq) (*pb.AssetChangeLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteAssetChangeLog")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AssetChangeLogDeleteReply{Result: err == nil && num > 0}, err
}
