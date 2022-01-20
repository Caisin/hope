package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/noveltag/v1"
	"hope/apps/param/internal/data/ent"
)

type NovelTagRepo interface {
	CreateNovelTag(context.Context, *v1.NovelTagCreateReq) (*ent.NovelTag, error)
	DeleteNovelTag(context.Context, *v1.NovelTagDeleteReq) error
	BatchDeleteNovelTag(context.Context, *v1.NovelTagBatchDeleteReq) (int, error)
	UpdateNovelTag(context.Context, *v1.NovelTagUpdateReq) (*ent.NovelTag, error)
	GetNovelTag(context.Context, *v1.NovelTagReq) (*ent.NovelTag, error)
	PageNovelTag(context.Context, *v1.NovelTagPageReq) ([]*ent.NovelTag, error)
}

type NovelTagUseCase struct {
	repo NovelTagRepo
	log  *log.Helper
}

func NewNovelTagUseCase(repo NovelTagRepo, logger log.Logger) *NovelTagUseCase {
	return &NovelTagUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelTagUseCase) Create(ctx context.Context, req *v1.NovelTagCreateReq) (*ent.NovelTag, error) {
	return uc.repo.CreateNovelTag(ctx, req)
}
func (uc *NovelTagUseCase) Delete(ctx context.Context, req *v1.NovelTagDeleteReq) error {
	return uc.repo.DeleteNovelTag(ctx, req)
}
func (uc *NovelTagUseCase) BatchDelete(ctx context.Context, req *v1.NovelTagBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelTag(ctx, req)
}
func (uc *NovelTagUseCase) Update(ctx context.Context, req *v1.NovelTagUpdateReq) (*ent.NovelTag, error) {
	return uc.repo.UpdateNovelTag(ctx, req)
}
func (uc *NovelTagUseCase) Get(ctx context.Context, req *v1.NovelTagReq) (*ent.NovelTag, error) {
	return uc.repo.GetNovelTag(ctx, req)
}
func (uc *NovelTagUseCase) Page(ctx context.Context, req *v1.NovelTagPageReq) ([]*ent.NovelTag, error) {
	return uc.repo.PageNovelTag(ctx, req)
}
