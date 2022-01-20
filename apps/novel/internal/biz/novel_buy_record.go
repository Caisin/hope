package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/novelbuyrecord/v1"
	"hope/apps/novel/internal/data/ent"
)

type NovelBuyRecordRepo interface {
	CreateNovelBuyRecord(context.Context, *v1.NovelBuyRecordCreateReq) (*ent.NovelBuyRecord, error)
	DeleteNovelBuyRecord(context.Context, *v1.NovelBuyRecordDeleteReq) error
	BatchDeleteNovelBuyRecord(context.Context, *v1.NovelBuyRecordBatchDeleteReq) (int, error)
	UpdateNovelBuyRecord(context.Context, *v1.NovelBuyRecordUpdateReq) (*ent.NovelBuyRecord, error)
	GetNovelBuyRecord(context.Context, *v1.NovelBuyRecordReq) (*ent.NovelBuyRecord, error)
	PageNovelBuyRecord(context.Context, *v1.NovelBuyRecordPageReq) ([]*ent.NovelBuyRecord, error)
}

type NovelBuyRecordUseCase struct {
	repo NovelBuyRecordRepo
	log  *log.Helper
}

func NewNovelBuyRecordUseCase(repo NovelBuyRecordRepo, logger log.Logger) *NovelBuyRecordUseCase {
	return &NovelBuyRecordUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelBuyRecordUseCase) Create(ctx context.Context, req *v1.NovelBuyRecordCreateReq) (*ent.NovelBuyRecord, error) {
	return uc.repo.CreateNovelBuyRecord(ctx, req)
}
func (uc *NovelBuyRecordUseCase) Delete(ctx context.Context, req *v1.NovelBuyRecordDeleteReq) error {
	return uc.repo.DeleteNovelBuyRecord(ctx, req)
}
func (uc *NovelBuyRecordUseCase) BatchDelete(ctx context.Context, req *v1.NovelBuyRecordBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelBuyRecord(ctx, req)
}
func (uc *NovelBuyRecordUseCase) Update(ctx context.Context, req *v1.NovelBuyRecordUpdateReq) (*ent.NovelBuyRecord, error) {
	return uc.repo.UpdateNovelBuyRecord(ctx, req)
}
func (uc *NovelBuyRecordUseCase) Get(ctx context.Context, req *v1.NovelBuyRecordReq) (*ent.NovelBuyRecord, error) {
	return uc.repo.GetNovelBuyRecord(ctx, req)
}
func (uc *NovelBuyRecordUseCase) Page(ctx context.Context, req *v1.NovelBuyRecordPageReq) ([]*ent.NovelBuyRecord, error) {
	return uc.repo.PageNovelBuyRecord(ctx, req)
}
