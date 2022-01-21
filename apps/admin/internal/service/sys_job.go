package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"

	pb "hope/api/admin/sysjob/v1"
)

type SysJobService struct {
	pb.UnimplementedSysJobServer
	uc  *biz.SysJobUseCase
	log *log.Helper
}

func NewSysJobService(uc *biz.SysJobUseCase, logger log.Logger) *SysJobService {
	return &SysJobService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysJobService) GetPageSysJob(ctx context.Context, req *pb.SysJobPageReq) (*pb.SysJobPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageSysJob")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.SysJobReply, 0)
	for i := range datas {
		items = append(items, convert.SysJobData2Reply(datas[i]))
	}
	reply := &pb.SysJobPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *SysJobService) GetSysJob(ctx context.Context, req *pb.SysJobReq) (*pb.SysJobReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetSysJob")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SysJobData2Reply(daya), err
}
func (s *SysJobService) UpdateSysJob(ctx context.Context, req *pb.SysJobUpdateReq) (*pb.SysJobUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateSysJob")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SysJobData2UpdateReply(data), err
}
func (s *SysJobService) CreateSysJob(ctx context.Context, req *pb.SysJobCreateReq) (*pb.SysJobCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateSysJob")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.SysJobData2CreateReply(data), err
}
func (s *SysJobService) DeleteSysJob(ctx context.Context, req *pb.SysJobDeleteReq) (*pb.SysJobDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteSysJob")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysJobDeleteReply{Result: err == nil}, err
}
func (s *SysJobService) BatchDeleteSysJob(ctx context.Context, req *pb.SysJobBatchDeleteReq) (*pb.SysJobDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteSysJob")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.SysJobDeleteReply{Result: err == nil && num > 0}, err
}
