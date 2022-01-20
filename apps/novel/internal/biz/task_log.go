package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/tasklog/v1"
	"hope/apps/novel/internal/data/ent"
)

type TaskLogRepo interface {
	CreateTaskLog(context.Context, *v1.TaskLogCreateReq) (*ent.TaskLog, error)
	DeleteTaskLog(context.Context, *v1.TaskLogDeleteReq) error
	BatchDeleteTaskLog(context.Context, *v1.TaskLogBatchDeleteReq) (int, error)
	UpdateTaskLog(context.Context, *v1.TaskLogUpdateReq) (*ent.TaskLog, error)
	GetTaskLog(context.Context, *v1.TaskLogReq) (*ent.TaskLog, error)
	PageTaskLog(context.Context, *v1.TaskLogPageReq) ([]*ent.TaskLog, error)
}

type TaskLogUseCase struct {
	repo TaskLogRepo
	log  *log.Helper
}

func NewTaskLogUseCase(repo TaskLogRepo, logger log.Logger) *TaskLogUseCase {
	return &TaskLogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TaskLogUseCase) Create(ctx context.Context, req *v1.TaskLogCreateReq) (*ent.TaskLog, error) {
	return uc.repo.CreateTaskLog(ctx, req)
}
func (uc *TaskLogUseCase) Delete(ctx context.Context, req *v1.TaskLogDeleteReq) error {
	return uc.repo.DeleteTaskLog(ctx, req)
}
func (uc *TaskLogUseCase) BatchDelete(ctx context.Context, req *v1.TaskLogBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteTaskLog(ctx, req)
}
func (uc *TaskLogUseCase) Update(ctx context.Context, req *v1.TaskLogUpdateReq) (*ent.TaskLog, error) {
	return uc.repo.UpdateTaskLog(ctx, req)
}
func (uc *TaskLogUseCase) Get(ctx context.Context, req *v1.TaskLogReq) (*ent.TaskLog, error) {
	return uc.repo.GetTaskLog(ctx, req)
}
func (uc *TaskLogUseCase) Page(ctx context.Context, req *v1.TaskLogPageReq) ([]*ent.TaskLog, error) {
	return uc.repo.PageTaskLog(ctx, req)
}
