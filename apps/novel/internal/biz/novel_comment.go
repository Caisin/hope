package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/novelcomment/v1"
	"hope/apps/novel/internal/data/ent"
)

type NovelCommentRepo interface {
	CreateNovelComment(context.Context, *v1.NovelCommentCreateReq) (*ent.NovelComment, error)
	DeleteNovelComment(context.Context, *v1.NovelCommentDeleteReq) error
	BatchDeleteNovelComment(context.Context, *v1.NovelCommentBatchDeleteReq) (int, error)
	UpdateNovelComment(context.Context, *v1.NovelCommentUpdateReq) (*ent.NovelComment, error)
	GetNovelComment(context.Context, *v1.NovelCommentReq) (*ent.NovelComment, error)
	PageNovelComment(context.Context, *v1.NovelCommentPageReq) ([]*ent.NovelComment, error)
}

type NovelCommentUseCase struct {
	repo NovelCommentRepo
	log  *log.Helper
}

func NewNovelCommentUseCase(repo NovelCommentRepo, logger log.Logger) *NovelCommentUseCase {
	return &NovelCommentUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelCommentUseCase) Create(ctx context.Context, req *v1.NovelCommentCreateReq) (*ent.NovelComment, error) {
	return uc.repo.CreateNovelComment(ctx, req)
}
func (uc *NovelCommentUseCase) Delete(ctx context.Context, req *v1.NovelCommentDeleteReq) error {
	return uc.repo.DeleteNovelComment(ctx, req)
}
func (uc *NovelCommentUseCase) BatchDelete(ctx context.Context, req *v1.NovelCommentBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelComment(ctx, req)
}
func (uc *NovelCommentUseCase) Update(ctx context.Context, req *v1.NovelCommentUpdateReq) (*ent.NovelComment, error) {
	return uc.repo.UpdateNovelComment(ctx, req)
}
func (uc *NovelCommentUseCase) Get(ctx context.Context, req *v1.NovelCommentReq) (*ent.NovelComment, error) {
	return uc.repo.GetNovelComment(ctx, req)
}
func (uc *NovelCommentUseCase) Page(ctx context.Context, req *v1.NovelCommentPageReq) ([]*ent.NovelComment, error) {
	return uc.repo.PageNovelComment(ctx, req)
}
