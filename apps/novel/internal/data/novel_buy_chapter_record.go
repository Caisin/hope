package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novelbuychapterrecord/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novelbuychapterrecord"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"
	"time"
)

type novelBuyChapterRecordRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelBuyChapterRecordRepo .
func NewNovelBuyChapterRecordRepo(data *Data, logger log.Logger) biz.NovelBuyChapterRecordRepo {
	return &novelBuyChapterRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelBuyChapterRecord 创建
func (r *novelBuyChapterRecordRepo) CreateNovelBuyChapterRecord(ctx context.Context, req *v1.NovelBuyChapterRecordCreateReq) (*ent.NovelBuyChapterRecord, error) {
	now := time.Now()
	return r.data.db.NovelBuyChapterRecord.Create().
		SetUserId(req.UserId).
		SetUserName(req.UserName).
		SetChapterId(req.ChapterId).
		SetChapterOrderNum(req.ChapterOrderNum).
		SetNovelId(req.NovelId).
		SetNovelName(req.NovelName).
		SetChapterName(req.ChapterName).
		SetIsSvip(req.IsSvip).
		SetCoin(req.Coin).
		SetCoupon(req.Coupon).
		SetDiscount(req.Discount).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteNovelBuyChapterRecord 删除
func (r *novelBuyChapterRecordRepo) DeleteNovelBuyChapterRecord(ctx context.Context, req *v1.NovelBuyChapterRecordDeleteReq) error {
	return r.data.db.NovelBuyChapterRecord.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelBuyChapterRecord 批量删除
func (r *novelBuyChapterRecordRepo) BatchDeleteNovelBuyChapterRecord(ctx context.Context, req *v1.NovelBuyChapterRecordBatchDeleteReq) (int, error) {
	return r.data.db.NovelBuyChapterRecord.Delete().Where(novelbuychapterrecord.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelBuyChapterRecord 更新
func (r *novelBuyChapterRecordRepo) UpdateNovelBuyChapterRecord(ctx context.Context, req *v1.NovelBuyChapterRecordUpdateReq) (*ent.NovelBuyChapterRecord, error) {
	return r.data.db.NovelBuyChapterRecord.UpdateOne(convert.NovelBuyChapterRecordUpdateReq2Data(req)).Save(ctx)
}

// GetNovelBuyChapterRecord 根据Id查询
func (r *novelBuyChapterRecordRepo) GetNovelBuyChapterRecord(ctx context.Context, req *v1.NovelBuyChapterRecordReq) (*ent.NovelBuyChapterRecord, error) {
	return r.data.db.NovelBuyChapterRecord.Get(ctx, req.Id)
}

// PageNovelBuyChapterRecord 分页查询
func (r *novelBuyChapterRecordRepo) PageNovelBuyChapterRecord(ctx context.Context, req *v1.NovelBuyChapterRecordPageReq) ([]*ent.NovelBuyChapterRecord, error) {
	pagin := req.Pagin
	query := r.data.db.NovelBuyChapterRecord.
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
	query.Limit(int(pagin.GetPage())).
		Offset(int(pagin.GetOffSet()))
	if pagin.NeedOrder() {
		if pagin.IsDesc() {
			query.Order(ent.Desc(pagin.GetField()))
		} else {
			query.Order(ent.Asc(pagin.GetField()))
		}
	}
	return query.All(ctx)
}

// genCondition 构造查询条件
func (r *novelBuyChapterRecordRepo) genCondition(req *v1.NovelBuyChapterRecordReq) []predicate.NovelBuyChapterRecord {
	list := make([]predicate.NovelBuyChapterRecord, 0)
	if req.Id > 0 {
		list = append(list, novelbuychapterrecord.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, novelbuychapterrecord.UserId(req.UserId))
	}
	if str.IsBlank(req.UserName) {
		list = append(list, novelbuychapterrecord.UserNameContains(req.UserName))
	}
	if req.ChapterId > 0 {
		list = append(list, novelbuychapterrecord.ChapterId(req.ChapterId))
	}
	if req.ChapterOrderNum > 0 {
		list = append(list, novelbuychapterrecord.ChapterOrderNum(req.ChapterOrderNum))
	}
	if req.NovelId > 0 {
		list = append(list, novelbuychapterrecord.NovelId(req.NovelId))
	}
	if str.IsBlank(req.NovelName) {
		list = append(list, novelbuychapterrecord.NovelNameContains(req.NovelName))
	}
	if str.IsBlank(req.ChapterName) {
		list = append(list, novelbuychapterrecord.ChapterNameContains(req.ChapterName))
	}
	if req.Coin > 0 {
		list = append(list, novelbuychapterrecord.Coin(req.Coin))
	}
	if req.Coupon > 0 {
		list = append(list, novelbuychapterrecord.Coupon(req.Coupon))
	}
	if req.Discount > 0 {
		list = append(list, novelbuychapterrecord.Discount(req.Discount))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, novelbuychapterrecord.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novelbuychapterrecord.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novelbuychapterrecord.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novelbuychapterrecord.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novelbuychapterrecord.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novelbuychapterrecord.TenantId(req.TenantId))
	}

	return list
}
