package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/noveltag/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/noveltag"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type novelTagRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelTagRepo .
func NewNovelTagRepo(data *Data, logger log.Logger) biz.NovelTagRepo {
	return &novelTagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelTag 创建
func (r *novelTagRepo) CreateNovelTag(ctx context.Context, req *v1.NovelTagCreateReq) (*ent.NovelTag, error) {
	now := time.Now()
	return r.data.db.NovelTag.Create().
		SetTagId(req.TagId).
		SetTagName(req.TagName).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteNovelTag 删除
func (r *novelTagRepo) DeleteNovelTag(ctx context.Context, req *v1.NovelTagDeleteReq) error {
	return r.data.db.NovelTag.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelTag 批量删除
func (r *novelTagRepo) BatchDeleteNovelTag(ctx context.Context, req *v1.NovelTagBatchDeleteReq) (int, error) {
	return r.data.db.NovelTag.Delete().Where(noveltag.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelTag 更新
func (r *novelTagRepo) UpdateNovelTag(ctx context.Context, req *v1.NovelTagUpdateReq) (*ent.NovelTag, error) {
	return r.data.db.NovelTag.UpdateOne(convert.NovelTagUpdateReq2Data(req)).Save(ctx)
}

// GetNovelTag 根据Id查询
func (r *novelTagRepo) GetNovelTag(ctx context.Context, req *v1.NovelTagReq) (*ent.NovelTag, error) {
	return r.data.db.NovelTag.Get(ctx, req.Id)
}

// PageNovelTag 分页查询
func (r *novelTagRepo) PageNovelTag(ctx context.Context, req *v1.NovelTagPageReq) ([]*ent.NovelTag, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.NovelTag.
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
	query.Limit(int(p.GetPage())).
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
func (r *novelTagRepo) genCondition(req *v1.NovelTagReq) []predicate.NovelTag {
	if req == nil {
		return nil
	}
	list := make([]predicate.NovelTag, 0)
	if req.Id > 0 {
		list = append(list, noveltag.ID(req.Id))
	}
	if req.TagId > 0 {
		list = append(list, noveltag.TagId(req.TagId))
	}
	if str.IsBlank(req.TagName) {
		list = append(list, noveltag.TagNameContains(req.TagName))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, noveltag.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, noveltag.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, noveltag.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, noveltag.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, noveltag.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, noveltag.TenantId(req.TenantId))
	}

	return list
}
