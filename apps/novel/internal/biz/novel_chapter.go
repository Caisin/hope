package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/novelchapter/v1"
	"hope/apps/novel/internal/data/ent"
)

type NovelChapterRepo interface {
	CreateNovelChapter(context.Context, *v1.NovelChapterCreateReq) (*ent.NovelChapter, error)
	DeleteNovelChapter(context.Context, *v1.NovelChapterDeleteReq) error
	BatchDeleteNovelChapter(context.Context, *v1.NovelChapterBatchDeleteReq) (int, error)
	UpdateNovelChapter(context.Context, *v1.NovelChapterUpdateReq) (*ent.NovelChapter, error)
	GetNovelChapter(context.Context, *v1.NovelChapterReq) (*ent.NovelChapter, error)
	PageNovelChapter(context.Context, *v1.NovelChapterPageReq) ([]*ent.NovelChapter, error)
}

type NovelChapterUseCase struct {
	repo NovelChapterRepo
	log  *log.Helper
}

func NewNovelChapterUseCase(repo NovelChapterRepo, logger log.Logger) *NovelChapterUseCase {
	return &NovelChapterUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelChapterUseCase) Create(ctx context.Context, req *v1.NovelChapterCreateReq) (*ent.NovelChapter, error) {
	return uc.repo.CreateNovelChapter(ctx, req)
}
func (uc *NovelChapterUseCase) Delete(ctx context.Context, req *v1.NovelChapterDeleteReq) error {
	return uc.repo.DeleteNovelChapter(ctx, req)
}
func (uc *NovelChapterUseCase) BatchDelete(ctx context.Context, req *v1.NovelChapterBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelChapter(ctx, req)
}
func (uc *NovelChapterUseCase) Update(ctx context.Context, req *v1.NovelChapterUpdateReq) (*ent.NovelChapter, error) {
	return uc.repo.UpdateNovelChapter(ctx, req)
}
func (uc *NovelChapterUseCase) Get(ctx context.Context, req *v1.NovelChapterReq) (*ent.NovelChapter, error) {
	return uc.repo.GetNovelChapter(ctx, req)
}
func (uc *NovelChapterUseCase) Page(ctx context.Context, req *v1.NovelChapterPageReq) ([]*ent.NovelChapter, error) {
	return uc.repo.PageNovelChapter(ctx, req)
}
