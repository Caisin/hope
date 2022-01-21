package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/qiniuconfig/v1"
)

type QiniuConfigService struct {
	pb.UnimplementedQiniuConfigServer
	uc  *biz.QiniuConfigUseCase
	log *log.Helper
}

func NewQiniuConfigService(uc *biz.QiniuConfigUseCase, logger log.Logger) *QiniuConfigService {
	return &QiniuConfigService{uc: uc, log: log.NewHelper(logger)}
}

func (s *QiniuConfigService) GetPageQiniuConfig(ctx context.Context, req *pb.QiniuConfigPageReq) (*pb.QiniuConfigPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageQiniuConfig")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.QiniuConfigReply, 0)
	for i := range datas {
		items = append(items, convert.QiniuConfigData2Reply(datas[i]))
	}
	reply := &pb.QiniuConfigPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *QiniuConfigService) GetQiniuConfig(ctx context.Context, req *pb.QiniuConfigReq) (*pb.QiniuConfigReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetQiniuConfig")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.QiniuConfigData2Reply(daya), err
}
func (s *QiniuConfigService) UpdateQiniuConfig(ctx context.Context, req *pb.QiniuConfigUpdateReq) (*pb.QiniuConfigUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateQiniuConfig")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.QiniuConfigData2UpdateReply(data), err
}
func (s *QiniuConfigService) CreateQiniuConfig(ctx context.Context, req *pb.QiniuConfigCreateReq) (*pb.QiniuConfigCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateQiniuConfig")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.QiniuConfigData2CreateReply(data), err
}
func (s *QiniuConfigService) DeleteQiniuConfig(ctx context.Context, req *pb.QiniuConfigDeleteReq) (*pb.QiniuConfigDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteQiniuConfig")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.QiniuConfigDeleteReply{Result: err == nil}, err
}
func (s *QiniuConfigService) BatchDeleteQiniuConfig(ctx context.Context, req *pb.QiniuConfigBatchDeleteReq) (*pb.QiniuConfigDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteQiniuConfig")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.QiniuConfigDeleteReply{Result: err == nil && num > 0}, err
}
