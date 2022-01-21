package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novelchapter/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novelchapter"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type novelChapterRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelChapterRepo .
func NewNovelChapterRepo(data *Data, logger log.Logger) biz.NovelChapterRepo {
	return &novelChapterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelChapter 创建
func (r *novelChapterRepo) CreateNovelChapter(ctx context.Context, req *v1.NovelChapterCreateReq) (*ent.NovelChapter, error) {
	now := time.Now()
	return r.data.db.NovelChapter.Create().
		SetNovelId(req.NovelId).
		SetOrderNum(req.OrderNum).
		SetChapterName(req.ChapterName).
		SetContent(req.Content).
		SetMediaKey(req.MediaKey).
		SetDuration(req.Duration).
		SetPublishTime(req.PublishTime.AsTime()).
		SetStatus(req.Status).
		SetIsFree(req.IsFree).
		SetPrice(req.Price).
		SetWordNum(req.WordNum).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteNovelChapter 删除
func (r *novelChapterRepo) DeleteNovelChapter(ctx context.Context, req *v1.NovelChapterDeleteReq) error {
	return r.data.db.NovelChapter.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelChapter 批量删除
func (r *novelChapterRepo) BatchDeleteNovelChapter(ctx context.Context, req *v1.NovelChapterBatchDeleteReq) (int, error) {
	return r.data.db.NovelChapter.Delete().Where(novelchapter.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelChapter 更新
func (r *novelChapterRepo) UpdateNovelChapter(ctx context.Context, req *v1.NovelChapterUpdateReq) (*ent.NovelChapter, error) {
	return r.data.db.NovelChapter.UpdateOne(convert.NovelChapterUpdateReq2Data(req)).Save(ctx)
}

// GetNovelChapter 根据Id查询
func (r *novelChapterRepo) GetNovelChapter(ctx context.Context, req *v1.NovelChapterReq) (*ent.NovelChapter, error) {
	return r.data.db.NovelChapter.Get(ctx, req.Id)
}

// PageNovelChapter 分页查询
func (r *novelChapterRepo) PageNovelChapter(ctx context.Context, req *v1.NovelChapterPageReq) ([]*ent.NovelChapter, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.NovelChapter.
		Query().
		Where(
			//查询条件构造
			r.genCondition(req.Param)...,
		)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	req.GetPagin().Total = int64(count)
	if count == 0 {
		return nil, nil
	}
	query.Limit(int(p.GetPageSize())).
		Offset(int(p.GetOffSet()))
	if p.NeedOrder() {
		if p.IsDesc() {
			query.Order(ent.Desc(p.GetField()))
		} else {
			query.Order(ent.Asc(p.GetField()))
		}
	}
	return query.All(ctx)
}

// genCondition 构造查询条件
func (r *novelChapterRepo) genCondition(req *v1.NovelChapterReq) []predicate.NovelChapter {
	if req == nil {
		return nil
	}
	list := make([]predicate.NovelChapter, 0)
	if req.Id > 0 {
		list = append(list, novelchapter.ID(req.Id))
	}
	if req.NovelId > 0 {
		list = append(list, novelchapter.NovelId(req.NovelId))
	}
	if req.OrderNum > 0 {
		list = append(list, novelchapter.OrderNum(req.OrderNum))
	}
	if str.IsBlank(req.ChapterName) {
		list = append(list, novelchapter.ChapterNameContains(req.ChapterName))
	}
	if str.IsBlank(req.Content) {
		list = append(list, novelchapter.ContentContains(req.Content))
	}
	if str.IsBlank(req.MediaKey) {
		list = append(list, novelchapter.MediaKeyContains(req.MediaKey))
	}
	if str.IsBlank(req.Duration) {
		list = append(list, novelchapter.DurationContains(req.Duration))
	}
	if req.PublishTime.IsValid() && !req.PublishTime.AsTime().IsZero() {
		list = append(list, novelchapter.PublishTimeGTE(req.PublishTime.AsTime()))
	}
	if req.Status > 0 {
		list = append(list, novelchapter.Status(req.Status))
	}
	if req.Price > 0 {
		list = append(list, novelchapter.Price(req.Price))
	}
	if req.WordNum > 0 {
		list = append(list, novelchapter.WordNum(req.WordNum))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, novelchapter.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novelchapter.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novelchapter.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novelchapter.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novelchapter.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novelchapter.TenantId(req.TenantId))
	}

	return list
}
