package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/scoreproduct/v1"
	"hope/apps/param/internal/data/ent"
)

type ScoreProductRepo interface {
	CreateScoreProduct(context.Context, *v1.ScoreProductCreateReq) (*ent.ScoreProduct, error)
	DeleteScoreProduct(context.Context, *v1.ScoreProductDeleteReq) error
	BatchDeleteScoreProduct(context.Context, *v1.ScoreProductBatchDeleteReq) (int, error)
	UpdateScoreProduct(context.Context, *v1.ScoreProductUpdateReq) (*ent.ScoreProduct, error)
	GetScoreProduct(context.Context, *v1.ScoreProductReq) (*ent.ScoreProduct, error)
	PageScoreProduct(context.Context, *v1.ScoreProductPageReq) ([]*ent.ScoreProduct, error)
}

type ScoreProductUseCase struct {
	repo ScoreProductRepo
	log  *log.Helper
}

func NewScoreProductUseCase(repo ScoreProductRepo, logger log.Logger) *ScoreProductUseCase {
	return &ScoreProductUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ScoreProductUseCase) Create(ctx context.Context, req *v1.ScoreProductCreateReq) (*ent.ScoreProduct, error) {
	return uc.repo.CreateScoreProduct(ctx, req)
}
func (uc *ScoreProductUseCase) Delete(ctx context.Context, req *v1.ScoreProductDeleteReq) error {
	return uc.repo.DeleteScoreProduct(ctx, req)
}
func (uc *ScoreProductUseCase) BatchDelete(ctx context.Context, req *v1.ScoreProductBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteScoreProduct(ctx, req)
}
func (uc *ScoreProductUseCase) Update(ctx context.Context, req *v1.ScoreProductUpdateReq) (*ent.ScoreProduct, error) {
	return uc.repo.UpdateScoreProduct(ctx, req)
}
func (uc *ScoreProductUseCase) Get(ctx context.Context, req *v1.ScoreProductReq) (*ent.ScoreProduct, error) {
	return uc.repo.GetScoreProduct(ctx, req)
}
func (uc *ScoreProductUseCase) Page(ctx context.Context, req *v1.ScoreProductPageReq) ([]*ent.ScoreProduct, error) {
	return uc.repo.PageScoreProduct(ctx, req)
}
