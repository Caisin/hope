package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"

	pb "hope/api/novel/tasklog/v1"
)

type TaskLogService struct {
	pb.UnimplementedTaskLogServer
	uc  *biz.TaskLogUseCase
	log *log.Helper
}

func NewTaskLogService(uc *biz.TaskLogUseCase, logger log.Logger) *TaskLogService {
	return &TaskLogService{uc: uc, log: log.NewHelper(logger)}
}

func (s *TaskLogService) GetPageTaskLog(ctx context.Context, req *pb.TaskLogPageReq) (*pb.TaskLogPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageTaskLog")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.TaskLogReply, 0)
	for i := range datas {
		items = append(items, convert.TaskLogData2Reply(datas[i]))
	}
	reply := &pb.TaskLogPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *TaskLogService) GetTaskLog(ctx context.Context, req *pb.TaskLogReq) (*pb.TaskLogReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetTaskLog")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.TaskLogData2Reply(daya), err
}
func (s *TaskLogService) UpdateTaskLog(ctx context.Context, req *pb.TaskLogUpdateReq) (*pb.TaskLogUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateTaskLog")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.TaskLogData2UpdateReply(data), err
}
func (s *TaskLogService) CreateTaskLog(ctx context.Context, req *pb.TaskLogCreateReq) (*pb.TaskLogCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateTaskLog")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.TaskLogData2CreateReply(data), err
}
func (s *TaskLogService) DeleteTaskLog(ctx context.Context, req *pb.TaskLogDeleteReq) (*pb.TaskLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteTaskLog")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.TaskLogDeleteReply{Result: err == nil}, err
}
func (s *TaskLogService) BatchDeleteTaskLog(ctx context.Context, req *pb.TaskLogBatchDeleteReq) (*pb.TaskLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteTaskLog")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.TaskLogDeleteReply{Result: err == nil && num > 0}, err
}
