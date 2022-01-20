package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/novelconsume/v1"
	"hope/apps/novel/internal/data/ent"
)

type NovelConsumeRepo interface {
	CreateNovelConsume(context.Context, *v1.NovelConsumeCreateReq) (*ent.NovelConsume, error)
	DeleteNovelConsume(context.Context, *v1.NovelConsumeDeleteReq) error
	BatchDeleteNovelConsume(context.Context, *v1.NovelConsumeBatchDeleteReq) (int, error)
	UpdateNovelConsume(context.Context, *v1.NovelConsumeUpdateReq) (*ent.NovelConsume, error)
	GetNovelConsume(context.Context, *v1.NovelConsumeReq) (*ent.NovelConsume, error)
	PageNovelConsume(context.Context, *v1.NovelConsumePageReq) ([]*ent.NovelConsume, error)
}

type NovelConsumeUseCase struct {
	repo NovelConsumeRepo
	log  *log.Helper
}

func NewNovelConsumeUseCase(repo NovelConsumeRepo, logger log.Logger) *NovelConsumeUseCase {
	return &NovelConsumeUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelConsumeUseCase) Create(ctx context.Context, req *v1.NovelConsumeCreateReq) (*ent.NovelConsume, error) {
	return uc.repo.CreateNovelConsume(ctx, req)
}
func (uc *NovelConsumeUseCase) Delete(ctx context.Context, req *v1.NovelConsumeDeleteReq) error {
	return uc.repo.DeleteNovelConsume(ctx, req)
}
func (uc *NovelConsumeUseCase) BatchDelete(ctx context.Context, req *v1.NovelConsumeBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelConsume(ctx, req)
}
func (uc *NovelConsumeUseCase) Update(ctx context.Context, req *v1.NovelConsumeUpdateReq) (*ent.NovelConsume, error) {
	return uc.repo.UpdateNovelConsume(ctx, req)
}
func (uc *NovelConsumeUseCase) Get(ctx context.Context, req *v1.NovelConsumeReq) (*ent.NovelConsume, error) {
	return uc.repo.GetNovelConsume(ctx, req)
}
func (uc *NovelConsumeUseCase) Page(ctx context.Context, req *v1.NovelConsumePageReq) ([]*ent.NovelConsume, error) {
	return uc.repo.PageNovelConsume(ctx, req)
}
