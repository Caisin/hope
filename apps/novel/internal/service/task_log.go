package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	datas, err := s.uc.Page(ctx, req)
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
	daya, err := s.uc.Get(ctx, req)
	return convert.TaskLogData2Reply(daya), err
}
func (s *TaskLogService) UpdateTaskLog(ctx context.Context, req *pb.TaskLogUpdateReq) (*pb.TaskLogUpdateReply, error) {
	data, err := s.uc.Update(ctx, req)
	return convert.TaskLogData2UpdateReply(data), err
}
func (s *TaskLogService) CreateTaskLog(ctx context.Context, req *pb.TaskLogCreateReq) (*pb.TaskLogCreateReply, error) {
	data, err := s.uc.Create(ctx, req)
	return convert.TaskLogData2CreateReply(data), err
}
func (s *TaskLogService) DeleteTaskLog(ctx context.Context, req *pb.TaskLogDeleteReq) (*pb.TaskLogDeleteReply, error) {
	err := s.uc.Delete(ctx, req)
	return &pb.TaskLogDeleteReply{Result: err == nil}, err
}
func (s *TaskLogService) BatchDeleteTaskLog(ctx context.Context, req *pb.TaskLogBatchDeleteReq) (*pb.TaskLogDeleteReply, error) {
	num, err := s.uc.BatchDelete(ctx, req)
	return &pb.TaskLogDeleteReply{Result: err == nil && num > 0}, err
}
