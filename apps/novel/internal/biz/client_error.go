package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/clienterror/v1"
	"hope/apps/novel/internal/data/ent"
)

type ClientErrorRepo interface {
	CreateClientError(context.Context, *v1.ClientErrorCreateReq) (*ent.ClientError, error)
	DeleteClientError(context.Context, *v1.ClientErrorDeleteReq) error
	BatchDeleteClientError(context.Context, *v1.ClientErrorBatchDeleteReq) (int, error)
	UpdateClientError(context.Context, *v1.ClientErrorUpdateReq) (*ent.ClientError, error)
	GetClientError(context.Context, *v1.ClientErrorReq) (*ent.ClientError, error)
	PageClientError(context.Context, *v1.ClientErrorPageReq) ([]*ent.ClientError, error)
}

type ClientErrorUseCase struct {
	repo ClientErrorRepo
	log  *log.Helper
}

func NewClientErrorUseCase(repo ClientErrorRepo, logger log.Logger) *ClientErrorUseCase {
	return &ClientErrorUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ClientErrorUseCase) Create(ctx context.Context, req *v1.ClientErrorCreateReq) (*ent.ClientError, error) {
	return uc.repo.CreateClientError(ctx, req)
}
func (uc *ClientErrorUseCase) Delete(ctx context.Context, req *v1.ClientErrorDeleteReq) error {
	return uc.repo.DeleteClientError(ctx, req)
}
func (uc *ClientErrorUseCase) BatchDelete(ctx context.Context, req *v1.ClientErrorBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteClientError(ctx, req)
}
func (uc *ClientErrorUseCase) Update(ctx context.Context, req *v1.ClientErrorUpdateReq) (*ent.ClientError, error) {
	return uc.repo.UpdateClientError(ctx, req)
}
func (uc *ClientErrorUseCase) Get(ctx context.Context, req *v1.ClientErrorReq) (*ent.ClientError, error) {
	return uc.repo.GetClientError(ctx, req)
}
func (uc *ClientErrorUseCase) Page(ctx context.Context, req *v1.ClientErrorPageReq) ([]*ent.ClientError, error) {
	return uc.repo.PageClientError(ctx, req)
}
