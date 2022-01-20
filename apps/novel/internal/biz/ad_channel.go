package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/adchannel/v1"
	"hope/apps/novel/internal/data/ent"
)

type AdChannelRepo interface {
	CreateAdChannel(context.Context, *v1.AdChannelCreateReq) (*ent.AdChannel, error)
	DeleteAdChannel(context.Context, *v1.AdChannelDeleteReq) error
	BatchDeleteAdChannel(context.Context, *v1.AdChannelBatchDeleteReq) (int, error)
	UpdateAdChannel(context.Context, *v1.AdChannelUpdateReq) (*ent.AdChannel, error)
	GetAdChannel(context.Context, *v1.AdChannelReq) (*ent.AdChannel, error)
	PageAdChannel(context.Context, *v1.AdChannelPageReq) ([]*ent.AdChannel, error)
}

type AdChannelUseCase struct {
	repo AdChannelRepo
	log  *log.Helper
}

func NewAdChannelUseCase(repo AdChannelRepo, logger log.Logger) *AdChannelUseCase {
	return &AdChannelUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AdChannelUseCase) Create(ctx context.Context, req *v1.AdChannelCreateReq) (*ent.AdChannel, error) {
	return uc.repo.CreateAdChannel(ctx, req)
}
func (uc *AdChannelUseCase) Delete(ctx context.Context, req *v1.AdChannelDeleteReq) error {
	return uc.repo.DeleteAdChannel(ctx, req)
}
func (uc *AdChannelUseCase) BatchDelete(ctx context.Context, req *v1.AdChannelBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteAdChannel(ctx, req)
}
func (uc *AdChannelUseCase) Update(ctx context.Context, req *v1.AdChannelUpdateReq) (*ent.AdChannel, error) {
	return uc.repo.UpdateAdChannel(ctx, req)
}
func (uc *AdChannelUseCase) Get(ctx context.Context, req *v1.AdChannelReq) (*ent.AdChannel, error) {
	return uc.repo.GetAdChannel(ctx, req)
}
func (uc *AdChannelUseCase) Page(ctx context.Context, req *v1.AdChannelPageReq) ([]*ent.AdChannel, error) {
	return uc.repo.PageAdChannel(ctx, req)
}
