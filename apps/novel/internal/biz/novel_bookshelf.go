package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/novelbookshelf/v1"
	"hope/apps/novel/internal/data/ent"
)

type NovelBookshelfRepo interface {
	CreateNovelBookshelf(context.Context, *v1.NovelBookshelfCreateReq) (*ent.NovelBookshelf, error)
	DeleteNovelBookshelf(context.Context, *v1.NovelBookshelfDeleteReq) error
	BatchDeleteNovelBookshelf(context.Context, *v1.NovelBookshelfBatchDeleteReq) (int, error)
	UpdateNovelBookshelf(context.Context, *v1.NovelBookshelfUpdateReq) (*ent.NovelBookshelf, error)
	GetNovelBookshelf(context.Context, *v1.NovelBookshelfReq) (*ent.NovelBookshelf, error)
	PageNovelBookshelf(context.Context, *v1.NovelBookshelfPageReq) ([]*ent.NovelBookshelf, error)
}

type NovelBookshelfUseCase struct {
	repo NovelBookshelfRepo
	log  *log.Helper
}

func NewNovelBookshelfUseCase(repo NovelBookshelfRepo, logger log.Logger) *NovelBookshelfUseCase {
	return &NovelBookshelfUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelBookshelfUseCase) Create(ctx context.Context, req *v1.NovelBookshelfCreateReq) (*ent.NovelBookshelf, error) {
	return uc.repo.CreateNovelBookshelf(ctx, req)
}
func (uc *NovelBookshelfUseCase) Delete(ctx context.Context, req *v1.NovelBookshelfDeleteReq) error {
	return uc.repo.DeleteNovelBookshelf(ctx, req)
}
func (uc *NovelBookshelfUseCase) BatchDelete(ctx context.Context, req *v1.NovelBookshelfBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelBookshelf(ctx, req)
}
func (uc *NovelBookshelfUseCase) Update(ctx context.Context, req *v1.NovelBookshelfUpdateReq) (*ent.NovelBookshelf, error) {
	return uc.repo.UpdateNovelBookshelf(ctx, req)
}
func (uc *NovelBookshelfUseCase) Get(ctx context.Context, req *v1.NovelBookshelfReq) (*ent.NovelBookshelf, error) {
	return uc.repo.GetNovelBookshelf(ctx, req)
}
func (uc *NovelBookshelfUseCase) Page(ctx context.Context, req *v1.NovelBookshelfPageReq) ([]*ent.NovelBookshelf, error) {
	return uc.repo.PageNovelBookshelf(ctx, req)
}
