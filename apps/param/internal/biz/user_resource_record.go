package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/userresourcerecord/v1"
	"hope/apps/param/internal/data/ent"
)

type UserResourceRecordRepo interface {
	CreateUserResourceRecord(context.Context, *v1.UserResourceRecordCreateReq) (*ent.UserResourceRecord, error)
	DeleteUserResourceRecord(context.Context, *v1.UserResourceRecordDeleteReq) error
	BatchDeleteUserResourceRecord(context.Context, *v1.UserResourceRecordBatchDeleteReq) (int, error)
	UpdateUserResourceRecord(context.Context, *v1.UserResourceRecordUpdateReq) (*ent.UserResourceRecord, error)
	GetUserResourceRecord(context.Context, *v1.UserResourceRecordReq) (*ent.UserResourceRecord, error)
	PageUserResourceRecord(context.Context, *v1.UserResourceRecordPageReq) ([]*ent.UserResourceRecord, error)
}

type UserResourceRecordUseCase struct {
	repo UserResourceRecordRepo
	log  *log.Helper
}

func NewUserResourceRecordUseCase(repo UserResourceRecordRepo, logger log.Logger) *UserResourceRecordUseCase {
	return &UserResourceRecordUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserResourceRecordUseCase) Create(ctx context.Context, req *v1.UserResourceRecordCreateReq) (*ent.UserResourceRecord, error) {
	return uc.repo.CreateUserResourceRecord(ctx, req)
}
func (uc *UserResourceRecordUseCase) Delete(ctx context.Context, req *v1.UserResourceRecordDeleteReq) error {
	return uc.repo.DeleteUserResourceRecord(ctx, req)
}
func (uc *UserResourceRecordUseCase) BatchDelete(ctx context.Context, req *v1.UserResourceRecordBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteUserResourceRecord(ctx, req)
}
func (uc *UserResourceRecordUseCase) Update(ctx context.Context, req *v1.UserResourceRecordUpdateReq) (*ent.UserResourceRecord, error) {
	return uc.repo.UpdateUserResourceRecord(ctx, req)
}
func (uc *UserResourceRecordUseCase) Get(ctx context.Context, req *v1.UserResourceRecordReq) (*ent.UserResourceRecord, error) {
	return uc.repo.GetUserResourceRecord(ctx, req)
}
func (uc *UserResourceRecordUseCase) Page(ctx context.Context, req *v1.UserResourceRecordPageReq) ([]*ent.UserResourceRecord, error) {
	return uc.repo.PageUserResourceRecord(ctx, req)
}
