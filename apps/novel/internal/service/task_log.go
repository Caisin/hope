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

func (s *TaskLogService) GetTaskLogPage(ctx context.Context, req *pb.TaskLogPageReq) (*pb.TaskLogPageReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetTaskLogPage")
	defer span.End()
	datas, err := s.uc.Page(ctx, req)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.TaskLogData, 0)
	for i := range datas {
		items = append(items, convert.TaskLogData2Reply(datas[i]))
	}
	reply := &pb.TaskLogPageReply{
		Code:    200,
		Message: "success",
		Total:   req.Pagin.Total,
		Items:   items,
	}
	return reply, nil
}
func (s *TaskLogService) GetTaskLog(ctx context.Context, req *pb.TaskLogReq) (*pb.TaskLogReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetTaskLog")
	defer span.End()
	data, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.TaskLogReply{
		Code:    200,
		Message: "success",
		Result:  convert.TaskLogData2Reply(data),
	}
	return reply, nil
}
func (s *TaskLogService) UpdateTaskLog(ctx context.Context, req *pb.TaskLogUpdateReq) (*pb.TaskLogUpdateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "UpdateTaskLog")
	defer span.End()
	data, err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.TaskLogUpdateReply{
		Code:    200,
		Message: "success",
		Result:  convert.TaskLogData2Reply(data),
	}
	return reply, nil
}
func (s *TaskLogService) CreateTaskLog(ctx context.Context, req *pb.TaskLogCreateReq) (*pb.TaskLogCreateReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "CreateTaskLog")
	defer span.End()
	data, err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	reply := &pb.TaskLogCreateReply{
		Code:    200,
		Message: "success",
		Result:  convert.TaskLogData2Reply(data),
	}
	return reply, err
}
func (s *TaskLogService) DeleteTaskLog(ctx context.Context, req *pb.TaskLogDeleteReq) (*pb.TaskLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "DeleteTaskLog")
	defer span.End()
	err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.TaskLogDeleteReply{Code: 200, Message: "success", Result: err == nil}, err
}
func (s *TaskLogService) BatchDeleteTaskLog(ctx context.Context, req *pb.TaskLogBatchDeleteReq) (*pb.TaskLogDeleteReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "BatchDeleteTaskLog")
	defer span.End()
	num, err := s.uc.BatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.TaskLogDeleteReply{Code: 200, Message: "success", Result: err == nil && num > 0}, err
}
