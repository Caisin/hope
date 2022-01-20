package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/novelmsg/v1"
	"hope/apps/novel/internal/data/ent"
)

type NovelMsgRepo interface {
	CreateNovelMsg(context.Context, *v1.NovelMsgCreateReq) (*ent.NovelMsg, error)
	DeleteNovelMsg(context.Context, *v1.NovelMsgDeleteReq) error
	BatchDeleteNovelMsg(context.Context, *v1.NovelMsgBatchDeleteReq) (int, error)
	UpdateNovelMsg(context.Context, *v1.NovelMsgUpdateReq) (*ent.NovelMsg, error)
	GetNovelMsg(context.Context, *v1.NovelMsgReq) (*ent.NovelMsg, error)
	PageNovelMsg(context.Context, *v1.NovelMsgPageReq) ([]*ent.NovelMsg, error)
}

type NovelMsgUseCase struct {
	repo NovelMsgRepo
	log  *log.Helper
}

func NewNovelMsgUseCase(repo NovelMsgRepo, logger log.Logger) *NovelMsgUseCase {
	return &NovelMsgUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelMsgUseCase) Create(ctx context.Context, req *v1.NovelMsgCreateReq) (*ent.NovelMsg, error) {
	return uc.repo.CreateNovelMsg(ctx, req)
}
func (uc *NovelMsgUseCase) Delete(ctx context.Context, req *v1.NovelMsgDeleteReq) error {
	return uc.repo.DeleteNovelMsg(ctx, req)
}
func (uc *NovelMsgUseCase) BatchDelete(ctx context.Context, req *v1.NovelMsgBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelMsg(ctx, req)
}
func (uc *NovelMsgUseCase) Update(ctx context.Context, req *v1.NovelMsgUpdateReq) (*ent.NovelMsg, error) {
	return uc.repo.UpdateNovelMsg(ctx, req)
}
func (uc *NovelMsgUseCase) Get(ctx context.Context, req *v1.NovelMsgReq) (*ent.NovelMsg, error) {
	return uc.repo.GetNovelMsg(ctx, req)
}
func (uc *NovelMsgUseCase) Page(ctx context.Context, req *v1.NovelMsgPageReq) ([]*ent.NovelMsg, error) {
	return uc.repo.PageNovelMsg(ctx, req)
}
