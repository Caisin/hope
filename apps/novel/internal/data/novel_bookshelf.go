package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novelbookshelf/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novelbookshelf"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type novelBookshelfRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelBookshelfRepo .
func NewNovelBookshelfRepo(data *Data, logger log.Logger) biz.NovelBookshelfRepo {
	return &novelBookshelfRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelBookshelf 创建
func (r *novelBookshelfRepo) CreateNovelBookshelf(ctx context.Context, req *v1.NovelBookshelfCreateReq) (*ent.NovelBookshelf, error) {
	now := time.Now()
	return r.data.db.NovelBookshelf.Create().
		SetUserId(req.UserId).
		SetUserName(req.UserName).
		SetNovelId(req.NovelId).
		SetLastReadTime(req.LastReadTime.AsTime()).
		SetLastChapterOrder(req.LastChapterOrder).
		SetLastChapterId(req.LastChapterId).
		SetLastChapterName(req.LastChapterName).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteNovelBookshelf 删除
func (r *novelBookshelfRepo) DeleteNovelBookshelf(ctx context.Context, req *v1.NovelBookshelfDeleteReq) error {
	return r.data.db.NovelBookshelf.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelBookshelf 批量删除
func (r *novelBookshelfRepo) BatchDeleteNovelBookshelf(ctx context.Context, req *v1.NovelBookshelfBatchDeleteReq) (int, error) {
	return r.data.db.NovelBookshelf.Delete().Where(novelbookshelf.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelBookshelf 更新
func (r *novelBookshelfRepo) UpdateNovelBookshelf(ctx context.Context, req *v1.NovelBookshelfUpdateReq) (*ent.NovelBookshelf, error) {
	return r.data.db.NovelBookshelf.UpdateOne(convert.NovelBookshelfUpdateReq2Data(req)).Save(ctx)
}

// GetNovelBookshelf 根据Id查询
func (r *novelBookshelfRepo) GetNovelBookshelf(ctx context.Context, req *v1.NovelBookshelfReq) (*ent.NovelBookshelf, error) {
	return r.data.db.NovelBookshelf.Get(ctx, req.Id)
}

// PageNovelBookshelf 分页查询
func (r *novelBookshelfRepo) PageNovelBookshelf(ctx context.Context, req *v1.NovelBookshelfPageReq) ([]*ent.NovelBookshelf, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.NovelBookshelf.
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
func (r *novelBookshelfRepo) genCondition(req *v1.NovelBookshelfReq) []predicate.NovelBookshelf {
	if req == nil {
		return nil
	}
	list := make([]predicate.NovelBookshelf, 0)
	if req.Id > 0 {
		list = append(list, novelbookshelf.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, novelbookshelf.UserId(req.UserId))
	}
	if str.IsBlank(req.UserName) {
		list = append(list, novelbookshelf.UserNameContains(req.UserName))
	}
	if req.NovelId > 0 {
		list = append(list, novelbookshelf.NovelId(req.NovelId))
	}
	if req.LastReadTime.IsValid() && !req.LastReadTime.AsTime().IsZero() {
		list = append(list, novelbookshelf.LastReadTimeGTE(req.LastReadTime.AsTime()))
	}
	if req.LastChapterOrder > 0 {
		list = append(list, novelbookshelf.LastChapterOrder(req.LastChapterOrder))
	}
	if req.LastChapterId > 0 {
		list = append(list, novelbookshelf.LastChapterId(req.LastChapterId))
	}
	if str.IsBlank(req.LastChapterName) {
		list = append(list, novelbookshelf.LastChapterNameContains(req.LastChapterName))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, novelbookshelf.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novelbookshelf.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novelbookshelf.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novelbookshelf.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novelbookshelf.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novelbookshelf.TenantId(req.TenantId))
	}

	return list
}
