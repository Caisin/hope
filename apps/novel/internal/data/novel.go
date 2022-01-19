package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novel/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novel"
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

func (r *novelRepo) CreateNovel(ctx context.Context, req *v1.NovelCreateReq) (*ent.Novel, error) {
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
		Save(ctx)

}
func (r *novelRepo) DeleteNovel(ctx context.Context, req *v1.NovelDeleteReq) error {
	return r.data.db.Novel.DeleteOneID(req.Id).Exec(ctx)
}
func (r *novelRepo) BatchDeleteNovel(ctx context.Context, req *v1.NovelBatchDeleteReq) (int, error) {
	return r.data.db.Novel.Delete().Where(novel.IDIn(req.Ids...)).Exec(ctx)
}
func (r *novelRepo) UpdateNovel(ctx context.Context, req *v1.NovelUpdateReq) (*ent.Novel, error) {
	return r.data.db.Novel.UpdateOne(convert.NovelUpdateReq2Data(req)).Save(ctx)
}
func (r *novelRepo) GetNovel(ctx context.Context, req *v1.NovelReq) (*ent.Novel, error) {
	return r.data.db.Novel.Get(ctx, req.Id)
}
func (r *novelRepo) PageNovel(ctx context.Context, req *v1.NovelPageReq) ([]*ent.Novel, error) {
	req.Param.GetTitle()
	pagin := req.Pagin
	query := r.data.db.Novel.
		Query().
		Where(
			//查询条件构造
			novel.TitleContains(req.Param.Title),
		).
		Limit(int(pagin.GetPage())).
		Offset(int(pagin.GetOffSet()))
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	req.GetPagin().Total = int64(count)
	if count == 0 {
		return nil, nil
	}
	if pagin.NeedOrder() {
		if pagin.IsDesc() {
			query.Order(ent.Desc(pagin.GetField()))
		} else {
			query.Order(ent.Asc(pagin.GetField()))
		}
	}
	return query.All(ctx)
}
