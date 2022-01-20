package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	datas, err := s.uc.Page(ctx, req)
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
	daya, err := s.uc.Get(ctx, req)
	return convert.TaskData2Reply(daya), err
}
func (s *TaskService) UpdateTask(ctx context.Context, req *pb.TaskUpdateReq) (*pb.TaskUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.TaskData2UpdateReply(data), err
}
func (s *TaskService) CreateTask(ctx context.Context, req *pb.TaskCreateReq) (*pb.TaskCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.TaskData2CreateReply(data), err
}
func (s *TaskService) DeleteTask(ctx context.Context, req *pb.TaskDeleteReq) (*pb.TaskDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.TaskDeleteReply{Result: err == nil}, err
}
func (s *TaskService) BatchDeleteTask(ctx context.Context, req *pb.TaskBatchDeleteReq) (*pb.TaskDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.TaskDeleteReply{Result: err == nil && num > 0}, err
}
