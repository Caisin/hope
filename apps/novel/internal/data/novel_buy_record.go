package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novelbuyrecord/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novelbuyrecord"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type novelBuyRecordRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelBuyRecordRepo .
func NewNovelBuyRecordRepo(data *Data, logger log.Logger) biz.NovelBuyRecordRepo {
	return &novelBuyRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelBuyRecord 创建
func (r *novelBuyRecordRepo) CreateNovelBuyRecord(ctx context.Context, req *v1.NovelBuyRecordCreateReq) (*ent.NovelBuyRecord, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.NovelBuyRecord.Create().
		SetUserId(req.UserId).
		SetUserName(req.UserName).
		SetNovelId(req.NovelId).
		SetNovelName(req.NovelName).
		SetPackageId(req.PackageId).
		SetCover(req.Cover).
		SetCoin(req.Coin).
		SetCoupon(req.Coupon).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteNovelBuyRecord 删除
func (r *novelBuyRecordRepo) DeleteNovelBuyRecord(ctx context.Context, req *v1.NovelBuyRecordDeleteReq) error {
	return r.data.db.NovelBuyRecord.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelBuyRecord 批量删除
func (r *novelBuyRecordRepo) BatchDeleteNovelBuyRecord(ctx context.Context, req *v1.NovelBuyRecordBatchDeleteReq) (int, error) {
	return r.data.db.NovelBuyRecord.Delete().Where(novelbuyrecord.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelBuyRecord 更新
func (r *novelBuyRecordRepo) UpdateNovelBuyRecord(ctx context.Context, req *v1.NovelBuyRecordUpdateReq) (*ent.NovelBuyRecord, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.NovelBuyRecord.UpdateOneID(req.Id).
		SetUserId(req.UserId).
		SetUserName(req.UserName).
		SetNovelId(req.NovelId).
		SetNovelName(req.NovelName).
		SetPackageId(req.PackageId).
		SetCover(req.Cover).
		SetCoin(req.Coin).
		SetCoupon(req.Coupon).
		SetRemark(req.Remark).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetNovelBuyRecord 根据Id查询
func (r *novelBuyRecordRepo) GetNovelBuyRecord(ctx context.Context, req *v1.NovelBuyRecordReq) (*ent.NovelBuyRecord, error) {
	return r.data.db.NovelBuyRecord.Get(ctx, req.Id)
}

// PageNovelBuyRecord 分页查询
func (r *novelBuyRecordRepo) PageNovelBuyRecord(ctx context.Context, req *v1.NovelBuyRecordPageReq) ([]*ent.NovelBuyRecord, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.NovelBuyRecord.
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
func (r *novelBuyRecordRepo) genCondition(req *v1.NovelBuyRecordReq) []predicate.NovelBuyRecord {
	if req == nil {
		return nil
	}
	list := make([]predicate.NovelBuyRecord, 0)
	if req.Id > 0 {
		list = append(list, novelbuyrecord.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, novelbuyrecord.UserId(req.UserId))
	}
	if str.IsNotBlank(req.UserName) {
		list = append(list, novelbuyrecord.UserNameContains(req.UserName))
	}
	if req.NovelId > 0 {
		list = append(list, novelbuyrecord.NovelId(req.NovelId))
	}
	if str.IsNotBlank(req.NovelName) {
		list = append(list, novelbuyrecord.NovelNameContains(req.NovelName))
	}
	if req.PackageId > 0 {
		list = append(list, novelbuyrecord.PackageId(req.PackageId))
	}
	if str.IsNotBlank(req.Cover) {
		list = append(list, novelbuyrecord.CoverContains(req.Cover))
	}
	if req.Coin > 0 {
		list = append(list, novelbuyrecord.Coin(req.Coin))
	}
	if req.Coupon > 0 {
		list = append(list, novelbuyrecord.Coupon(req.Coupon))
	}
	if str.IsNotBlank(req.Remark) {
		list = append(list, novelbuyrecord.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novelbuyrecord.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novelbuyrecord.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novelbuyrecord.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novelbuyrecord.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novelbuyrecord.TenantId(req.TenantId))
	}

	return list
}
