package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/novelclassify/v1"
	"hope/apps/novel/internal/data/ent"
)

type NovelClassifyRepo interface {
	CreateNovelClassify(context.Context, *v1.NovelClassifyCreateReq) (*ent.NovelClassify, error)
	DeleteNovelClassify(context.Context, *v1.NovelClassifyDeleteReq) error
	BatchDeleteNovelClassify(context.Context, *v1.NovelClassifyBatchDeleteReq) (int, error)
	UpdateNovelClassify(context.Context, *v1.NovelClassifyUpdateReq) (*ent.NovelClassify, error)
	GetNovelClassify(context.Context, *v1.NovelClassifyReq) (*ent.NovelClassify, error)
	PageNovelClassify(context.Context, *v1.NovelClassifyPageReq) ([]*ent.NovelClassify, error)
}

type NovelClassifyUseCase struct {
	repo NovelClassifyRepo
	log  *log.Helper
}

func NewNovelClassifyUseCase(repo NovelClassifyRepo, logger log.Logger) *NovelClassifyUseCase {
	return &NovelClassifyUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelClassifyUseCase) Create(ctx context.Context, req *v1.NovelClassifyCreateReq) (*ent.NovelClassify, error) {
	return uc.repo.CreateNovelClassify(ctx, req)
}
func (uc *NovelClassifyUseCase) Delete(ctx context.Context, req *v1.NovelClassifyDeleteReq) error {
	return uc.repo.DeleteNovelClassify(ctx, req)
}
func (uc *NovelClassifyUseCase) BatchDelete(ctx context.Context, req *v1.NovelClassifyBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelClassify(ctx, req)
}
func (uc *NovelClassifyUseCase) Update(ctx context.Context, req *v1.NovelClassifyUpdateReq) (*ent.NovelClassify, error) {
	return uc.repo.UpdateNovelClassify(ctx, req)
}
func (uc *NovelClassifyUseCase) Get(ctx context.Context, req *v1.NovelClassifyReq) (*ent.NovelClassify, error) {
	return uc.repo.GetNovelClassify(ctx, req)
}
func (uc *NovelClassifyUseCase) Page(ctx context.Context, req *v1.NovelClassifyPageReq) ([]*ent.NovelClassify, error) {
	return uc.repo.PageNovelClassify(ctx, req)
}
