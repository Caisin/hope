package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/novelautobuy/v1"
	"hope/apps/novel/internal/data/ent"
)

type NovelAutoBuyRepo interface {
	CreateNovelAutoBuy(context.Context, *v1.NovelAutoBuyCreateReq) (*ent.NovelAutoBuy, error)
	DeleteNovelAutoBuy(context.Context, *v1.NovelAutoBuyDeleteReq) error
	BatchDeleteNovelAutoBuy(context.Context, *v1.NovelAutoBuyBatchDeleteReq) (int, error)
	UpdateNovelAutoBuy(context.Context, *v1.NovelAutoBuyUpdateReq) (*ent.NovelAutoBuy, error)
	GetNovelAutoBuy(context.Context, *v1.NovelAutoBuyReq) (*ent.NovelAutoBuy, error)
	PageNovelAutoBuy(context.Context, *v1.NovelAutoBuyPageReq) ([]*ent.NovelAutoBuy, error)
}

type NovelAutoBuyUseCase struct {
	repo NovelAutoBuyRepo
	log  *log.Helper
}

func NewNovelAutoBuyUseCase(repo NovelAutoBuyRepo, logger log.Logger) *NovelAutoBuyUseCase {
	return &NovelAutoBuyUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelAutoBuyUseCase) Create(ctx context.Context, req *v1.NovelAutoBuyCreateReq) (*ent.NovelAutoBuy, error) {
	return uc.repo.CreateNovelAutoBuy(ctx, req)
}
func (uc *NovelAutoBuyUseCase) Delete(ctx context.Context, req *v1.NovelAutoBuyDeleteReq) error {
	return uc.repo.DeleteNovelAutoBuy(ctx, req)
}
func (uc *NovelAutoBuyUseCase) BatchDelete(ctx context.Context, req *v1.NovelAutoBuyBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelAutoBuy(ctx, req)
}
func (uc *NovelAutoBuyUseCase) Update(ctx context.Context, req *v1.NovelAutoBuyUpdateReq) (*ent.NovelAutoBuy, error) {
	return uc.repo.UpdateNovelAutoBuy(ctx, req)
}
func (uc *NovelAutoBuyUseCase) Get(ctx context.Context, req *v1.NovelAutoBuyReq) (*ent.NovelAutoBuy, error) {
	return uc.repo.GetNovelAutoBuy(ctx, req)
}
func (uc *NovelAutoBuyUseCase) Page(ctx context.Context, req *v1.NovelAutoBuyPageReq) ([]*ent.NovelAutoBuy, error) {
	return uc.repo.PageNovelAutoBuy(ctx, req)
}
