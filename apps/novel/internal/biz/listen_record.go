package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/listenrecord/v1"
	"hope/apps/novel/internal/data/ent"
)

type ListenRecordRepo interface {
	CreateListenRecord(context.Context, *v1.ListenRecordCreateReq) (*ent.ListenRecord, error)
	DeleteListenRecord(context.Context, *v1.ListenRecordDeleteReq) error
	BatchDeleteListenRecord(context.Context, *v1.ListenRecordBatchDeleteReq) (int, error)
	UpdateListenRecord(context.Context, *v1.ListenRecordUpdateReq) (*ent.ListenRecord, error)
	GetListenRecord(context.Context, *v1.ListenRecordReq) (*ent.ListenRecord, error)
	PageListenRecord(context.Context, *v1.ListenRecordPageReq) ([]*ent.ListenRecord, error)
}

type ListenRecordUseCase struct {
	repo ListenRecordRepo
	log  *log.Helper
}

func NewListenRecordUseCase(repo ListenRecordRepo, logger log.Logger) *ListenRecordUseCase {
	return &ListenRecordUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ListenRecordUseCase) Create(ctx context.Context, req *v1.ListenRecordCreateReq) (*ent.ListenRecord, error) {
	return uc.repo.CreateListenRecord(ctx, req)
}
func (uc *ListenRecordUseCase) Delete(ctx context.Context, req *v1.ListenRecordDeleteReq) error {
	return uc.repo.DeleteListenRecord(ctx, req)
}
func (uc *ListenRecordUseCase) BatchDelete(ctx context.Context, req *v1.ListenRecordBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteListenRecord(ctx, req)
}
func (uc *ListenRecordUseCase) Update(ctx context.Context, req *v1.ListenRecordUpdateReq) (*ent.ListenRecord, error) {
	return uc.repo.UpdateListenRecord(ctx, req)
}
func (uc *ListenRecordUseCase) Get(ctx context.Context, req *v1.ListenRecordReq) (*ent.ListenRecord, error) {
	return uc.repo.GetListenRecord(ctx, req)
}
func (uc *ListenRecordUseCase) Page(ctx context.Context, req *v1.ListenRecordPageReq) ([]*ent.ListenRecord, error) {
	return uc.repo.PageListenRecord(ctx, req)
}
