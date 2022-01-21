package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/appversion/v1"
)

type AppVersionService struct {
	pb.UnimplementedAppVersionServer
	uc  *biz.AppVersionUseCase
	log *log.Helper
}

func NewAppVersionService(uc *biz.AppVersionUseCase, logger log.Logger) *AppVersionService {
	return &AppVersionService{uc: uc, log: log.NewHelper(logger)}
}

func (s *AppVersionService) GetPageAppVersion(ctx context.Context, req *pb.AppVersionPageReq) (*pb.AppVersionPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageAppVersion")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.AppVersionReply, 0)
	for i := range datas {
		items = append(items, convert.AppVersionData2Reply(datas[i]))
	}
	reply := &pb.AppVersionPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *AppVersionService) GetAppVersion(ctx context.Context, req *pb.AppVersionReq) (*pb.AppVersionReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetAppVersion")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.AppVersionData2Reply(daya), err
}
func (s *AppVersionService) UpdateAppVersion(ctx context.Context, req *pb.AppVersionUpdateReq) (*pb.AppVersionUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateAppVersion")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.AppVersionData2UpdateReply(data), err
}
func (s *AppVersionService) CreateAppVersion(ctx context.Context, req *pb.AppVersionCreateReq) (*pb.AppVersionCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateAppVersion")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.AppVersionData2CreateReply(data), err
}
func (s *AppVersionService) DeleteAppVersion(ctx context.Context, req *pb.AppVersionDeleteReq) (*pb.AppVersionDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteAppVersion")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AppVersionDeleteReply{Result: err == nil}, err
}
func (s *AppVersionService) BatchDeleteAppVersion(ctx context.Context, req *pb.AppVersionBatchDeleteReq) (*pb.AppVersionDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteAppVersion")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.AppVersionDeleteReply{Result: err == nil && num > 0}, err
}
