package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/task/v1"
	"hope/apps/param/internal/data/ent"
)

type TaskRepo interface {
	CreateTask(context.Context, *v1.TaskCreateReq) (*ent.Task, error)
	DeleteTask(context.Context, *v1.TaskDeleteReq) error
	BatchDeleteTask(context.Context, *v1.TaskBatchDeleteReq) (int, error)
	UpdateTask(context.Context, *v1.TaskUpdateReq) (*ent.Task, error)
	GetTask(context.Context, *v1.TaskReq) (*ent.Task, error)
	PageTask(context.Context, *v1.TaskPageReq) ([]*ent.Task, error)
}

type TaskUseCase struct {
	repo TaskRepo
	log  *log.Helper
}

func NewTaskUseCase(repo TaskRepo, logger log.Logger) *TaskUseCase {
	return &TaskUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TaskUseCase) Create(ctx context.Context, req *v1.TaskCreateReq) (*ent.Task, error) {
	return uc.repo.CreateTask(ctx, req)
}
func (uc *TaskUseCase) Delete(ctx context.Context, req *v1.TaskDeleteReq) error {
	return uc.repo.DeleteTask(ctx, req)
}
func (uc *TaskUseCase) BatchDelete(ctx context.Context, req *v1.TaskBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteTask(ctx, req)
}
func (uc *TaskUseCase) Update(ctx context.Context, req *v1.TaskUpdateReq) (*ent.Task, error) {
	return uc.repo.UpdateTask(ctx, req)
}
func (uc *TaskUseCase) Get(ctx context.Context, req *v1.TaskReq) (*ent.Task, error) {
	return uc.repo.GetTask(ctx, req)
}
func (uc *TaskUseCase) Page(ctx context.Context, req *v1.TaskPageReq) ([]*ent.Task, error) {
	return uc.repo.PageTask(ctx, req)
}
