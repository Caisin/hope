package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"

	pb "hope/api/param/task/v1"
)

type TaskService struct {
	pb.UnimplementedTaskServer
	uc  *biz.TaskUseCase
	log *log.Helper
}

func NewTaskService(uc *biz.TaskUseCase, logger log.Logger) *TaskService {
	return &TaskService{uc: uc, log: log.NewHelper(logger)}
}

func (s *TaskService) GetPageTask(ctx context.Context, req *pb.TaskPageReq) (*pb.TaskPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetPageTask")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.TaskReply, 0)
	for i := range datas {
		items = append(items, convert.TaskData2Reply(datas[i]))
	}
	reply := &pb.TaskPageReply{
		Pagin: req.Pagin,
		Items: items,
	}
	return reply, err
}
func (s *TaskService) GetTask(ctx context.Context, req *pb.TaskReq) (*pb.TaskReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetTask")
	defer span.End()
	daya, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.TaskData2Reply(daya), err
}
func (s *TaskService) UpdateTask(ctx context.Context, req *pb.TaskUpdateReq) (*pb.TaskUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateTask")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.TaskData2UpdateReply(data), err
}
func (s *TaskService) CreateTask(ctx context.Context, req *pb.TaskCreateReq) (*pb.TaskCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateTask")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert.TaskData2CreateReply(data), err
}
func (s *TaskService) DeleteTask(ctx context.Context, req *pb.TaskDeleteReq) (*pb.TaskDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteTask")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.TaskDeleteReply{Result: err == nil}, err
}
func (s *TaskService) BatchDeleteTask(ctx context.Context, req *pb.TaskBatchDeleteReq) (*pb.TaskDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteTask")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.TaskDeleteReply{Result: err == nil && num > 0}, err
}
