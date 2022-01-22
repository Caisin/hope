package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novel/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novel"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type novelRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelRepo .
func NewNovelRepo(data *Data, logger log.Logger) biz.NovelRepo {
	return &novelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovel 创建
func (r *novelRepo) CreateNovel(ctx context.Context, req *v1.NovelCreateReq) (*ent.Novel, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.Novel.Create().
		SetClassifyId(req.ClassifyId).
		SetClassifyName(req.ClassifyName).
		SetAuthorId(req.AuthorId).
		SetTitle(req.Title).
		SetSummary(req.Summary).
		SetAuthor(req.Author).
		SetAnchor(req.Anchor).
		SetHits(req.Hits).
		SetKeywords(req.Keywords).
		SetSource(req.Source).
		SetScore(req.Score).
		SetCover(req.Cover).
		SetTagIds(req.TagIds).
		SetWordNum(req.WordNum).
		SetFreeNum(req.FreeNum).
		SetOnlineState(req.OnlineState).
		SetPrice(req.Price).
		SetPublish(req.Publish).
		SetOriginalPrice(req.OriginalPrice).
		SetChapterPrice(req.ChapterPrice).
		SetChapterCount(req.ChapterCount).
		SetSignType(req.SignType).
		SetSignDate(req.SignDate.AsTime()).
		SetLeadingMan(req.LeadingMan).
		SetLeadingLady(req.LeadingLady).
		SetRemark(req.Remark).
		SetMediaKey(req.MediaKey).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteNovel 删除
func (r *novelRepo) DeleteNovel(ctx context.Context, req *v1.NovelDeleteReq) error {
	return r.data.db.Novel.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovel 批量删除
func (r *novelRepo) BatchDeleteNovel(ctx context.Context, req *v1.NovelBatchDeleteReq) (int, error) {
	return r.data.db.Novel.Delete().Where(novel.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovel 更新
func (r *novelRepo) UpdateNovel(ctx context.Context, req *v1.NovelUpdateReq) (*ent.Novel, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	data := convert.NovelUpdateReq2Data(req)
	data.UpdateBy = claims.UserId
	return r.data.db.Novel.UpdateOne(data).Save(ctx)
}

// GetNovel 根据Id查询
func (r *novelRepo) GetNovel(ctx context.Context, req *v1.NovelReq) (*ent.Novel, error) {
	return r.data.db.Novel.Get(ctx, req.Id)
}

// PageNovel 分页查询
func (r *novelRepo) PageNovel(ctx context.Context, req *v1.NovelPageReq) ([]*ent.Novel, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.Novel.
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
func (r *novelRepo) genCondition(req *v1.NovelReq) []predicate.Novel {
	if req == nil {
		return nil
	}
	list := make([]predicate.Novel, 0)
	if req.Id > 0 {
		list = append(list, novel.ID(req.Id))
	}
	if req.ClassifyId > 0 {
		list = append(list, novel.ClassifyId(req.ClassifyId))
	}
	if str.IsBlank(req.ClassifyName) {
		list = append(list, novel.ClassifyNameContains(req.ClassifyName))
	}
	if str.IsBlank(req.AuthorId) {
		list = append(list, novel.AuthorIdContains(req.AuthorId))
	}
	if str.IsBlank(req.Title) {
		list = append(list, novel.TitleContains(req.Title))
	}
	if str.IsBlank(req.Summary) {
		list = append(list, novel.SummaryContains(req.Summary))
	}
	if str.IsBlank(req.Author) {
		list = append(list, novel.AuthorContains(req.Author))
	}
	if str.IsBlank(req.Anchor) {
		list = append(list, novel.AnchorContains(req.Anchor))
	}
	if req.Hits > 0 {
		list = append(list, novel.Hits(req.Hits))
	}
	if str.IsBlank(req.Keywords) {
		list = append(list, novel.KeywordsContains(req.Keywords))
	}
	if str.IsBlank(req.Source) {
		list = append(list, novel.SourceContains(req.Source))
	}
	if req.Score > 0 {
		list = append(list, novel.Score(req.Score))
	}
	if str.IsBlank(req.Cover) {
		list = append(list, novel.CoverContains(req.Cover))
	}
	if str.IsBlank(req.TagIds) {
		list = append(list, novel.TagIdsContains(req.TagIds))
	}
	if req.WordNum > 0 {
		list = append(list, novel.WordNum(req.WordNum))
	}
	if req.FreeNum > 0 {
		list = append(list, novel.FreeNum(req.FreeNum))
	}
	if req.OnlineState > 0 {
		list = append(list, novel.OnlineState(req.OnlineState))
	}
	if req.Price > 0 {
		list = append(list, novel.Price(req.Price))
	}
	if req.Publish > 0 {
		list = append(list, novel.Publish(req.Publish))
	}
	if req.OriginalPrice > 0 {
		list = append(list, novel.OriginalPrice(req.OriginalPrice))
	}
	if req.ChapterPrice > 0 {
		list = append(list, novel.ChapterPrice(req.ChapterPrice))
	}
	if req.ChapterCount > 0 {
		list = append(list, novel.ChapterCount(req.ChapterCount))
	}
	if req.SignType > 0 {
		list = append(list, novel.SignType(req.SignType))
	}
	if req.SignDate.IsValid() && !req.SignDate.AsTime().IsZero() {
		list = append(list, novel.SignDateGTE(req.SignDate.AsTime()))
	}
	if str.IsBlank(req.LeadingMan) {
		list = append(list, novel.LeadingManContains(req.LeadingMan))
	}
	if str.IsBlank(req.LeadingLady) {
		list = append(list, novel.LeadingLadyContains(req.LeadingLady))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, novel.RemarkContains(req.Remark))
	}
	if str.IsBlank(req.MediaKey) {
		list = append(list, novel.MediaKeyContains(req.MediaKey))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novel.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novel.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novel.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novel.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novel.TenantId(req.TenantId))
	}

	return list
}
