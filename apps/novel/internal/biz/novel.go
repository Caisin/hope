package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/novel/v1"
	"hope/apps/novel/internal/data/ent"
)

type NovelRepo interface {
	CreateNovel(context.Context, *v1.NovelCreateReq) (*ent.Novel, error)
	DeleteNovel(context.Context, *v1.NovelDeleteReq) error
	BatchDeleteNovel(context.Context, *v1.NovelBatchDeleteReq) (int, error)
	UpdateNovel(context.Context, *v1.NovelUpdateReq) (*ent.Novel, error)
	GetNovel(context.Context, *v1.NovelReq) (*ent.Novel, error)
	PageNovel(context.Context, *v1.NovelPageReq) ([]*ent.Novel, error)
}

type NovelUseCase struct {
	repo NovelRepo
	log  *log.Helper
}

func NewNovelUseCase(repo NovelRepo, logger log.Logger) *NovelUseCase {
	return &NovelUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelUseCase) Create(ctx context.Context, req *v1.NovelCreateReq) (*ent.Novel, error) {
	return uc.repo.CreateNovel(ctx, req)
}
func (uc *NovelUseCase) Delete(ctx context.Context, req *v1.NovelDeleteReq) error {
	return uc.repo.DeleteNovel(ctx, req)
}
func (uc *NovelUseCase) BatchDelete(ctx context.Context, req *v1.NovelBatchDeleteReq) error {
	return uc.repo.BatchDeleteNovel(ctx, req)
}
func (uc *NovelUseCase) Update(ctx context.Context, req *v1.NovelUpdateReq) (*ent.Novel, error) {
	return uc.repo.UpdateNovel(ctx, req)
}
func (uc *NovelUseCase) Get(ctx context.Context, req *v1.NovelReq) (*ent.Novel, error) {
	return uc.repo.GetNovel(ctx, req)
}
func (uc *NovelUseCase) Page(ctx context.Context, req *v1.NovelPageReq) ([]*ent.Novel, error) {
	return uc.repo.PageNovel(ctx, req)
}
