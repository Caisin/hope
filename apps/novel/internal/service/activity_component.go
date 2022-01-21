package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/activitycomponent/v1"
)

type ActivityComponentService struct {
	pb.UnimplementedActivityComponentServer
	uc  *biz.ActivityComponentUseCase
	log *log.Helper
}

func NewActivityComponentService(uc *biz.ActivityComponentUseCase, logger log.Logger) *ActivityComponentService {
	return &ActivityComponentService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ActivityComponentService) GetPageActivityComponent(ctx context.Context, req *pb.ActivityComponentPageReq) (*pb.ActivityComponentPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageActivityComponent")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.ActivityComponentReply, 0)
	for i := range datas {
		items = append(items, convert.ActivityComponentData2Reply(datas[i]))
	}
	reply := &pb.ActivityComponentPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *ActivityComponentService) GetActivityComponent(ctx context.Context, req *pb.ActivityComponentReq) (*pb.ActivityComponentReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetActivityComponent")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.ActivityComponentData2Reply(daya), err
}
func (s *ActivityComponentService) UpdateActivityComponent(ctx context.Context, req *pb.ActivityComponentUpdateReq) (*pb.ActivityComponentUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateActivityComponent")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.ActivityComponentData2UpdateReply(data), err
}
func (s *ActivityComponentService) CreateActivityComponent(ctx context.Context, req *pb.ActivityComponentCreateReq) (*pb.ActivityComponentCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateActivityComponent")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.ActivityComponentData2CreateReply(data), err
}
func (s *ActivityComponentService) DeleteActivityComponent(ctx context.Context, req *pb.ActivityComponentDeleteReq) (*pb.ActivityComponentDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteActivityComponent")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ActivityComponentDeleteReply{Result: err == nil}, err
}
func (s *ActivityComponentService) BatchDeleteActivityComponent(ctx context.Context, req *pb.ActivityComponentBatchDeleteReq) (*pb.ActivityComponentDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteActivityComponent")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ActivityComponentDeleteReply{Result: err == nil && num > 0}, err
}
