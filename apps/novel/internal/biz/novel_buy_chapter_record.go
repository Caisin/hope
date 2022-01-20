package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/novelbuychapterrecord/v1"
	"hope/apps/novel/internal/data/ent"
)

type NovelBuyChapterRecordRepo interface {
	CreateNovelBuyChapterRecord(context.Context, *v1.NovelBuyChapterRecordCreateReq) (*ent.NovelBuyChapterRecord, error)
	DeleteNovelBuyChapterRecord(context.Context, *v1.NovelBuyChapterRecordDeleteReq) error
	BatchDeleteNovelBuyChapterRecord(context.Context, *v1.NovelBuyChapterRecordBatchDeleteReq) (int, error)
	UpdateNovelBuyChapterRecord(context.Context, *v1.NovelBuyChapterRecordUpdateReq) (*ent.NovelBuyChapterRecord, error)
	GetNovelBuyChapterRecord(context.Context, *v1.NovelBuyChapterRecordReq) (*ent.NovelBuyChapterRecord, error)
	PageNovelBuyChapterRecord(context.Context, *v1.NovelBuyChapterRecordPageReq) ([]*ent.NovelBuyChapterRecord, error)
}

type NovelBuyChapterRecordUseCase struct {
	repo NovelBuyChapterRecordRepo
	log  *log.Helper
}

func NewNovelBuyChapterRecordUseCase(repo NovelBuyChapterRecordRepo, logger log.Logger) *NovelBuyChapterRecordUseCase {
	return &NovelBuyChapterRecordUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *NovelBuyChapterRecordUseCase) Create(ctx context.Context, req *v1.NovelBuyChapterRecordCreateReq) (*ent.NovelBuyChapterRecord, error) {
	return uc.repo.CreateNovelBuyChapterRecord(ctx, req)
}
func (uc *NovelBuyChapterRecordUseCase) Delete(ctx context.Context, req *v1.NovelBuyChapterRecordDeleteReq) error {
	return uc.repo.DeleteNovelBuyChapterRecord(ctx, req)
}
func (uc *NovelBuyChapterRecordUseCase) BatchDelete(ctx context.Context, req *v1.NovelBuyChapterRecordBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteNovelBuyChapterRecord(ctx, req)
}
func (uc *NovelBuyChapterRecordUseCase) Update(ctx context.Context, req *v1.NovelBuyChapterRecordUpdateReq) (*ent.NovelBuyChapterRecord, error) {
	return uc.repo.UpdateNovelBuyChapterRecord(ctx, req)
}
func (uc *NovelBuyChapterRecordUseCase) Get(ctx context.Context, req *v1.NovelBuyChapterRecordReq) (*ent.NovelBuyChapterRecord, error) {
	return uc.repo.GetNovelBuyChapterRecord(ctx, req)
}
func (uc *NovelBuyChapterRecordUseCase) Page(ctx context.Context, req *v1.NovelBuyChapterRecordPageReq) ([]*ent.NovelBuyChapterRecord, error) {
	return uc.repo.PageNovelBuyChapterRecord(ctx, req)
}
