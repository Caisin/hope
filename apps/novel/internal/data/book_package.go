package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/bookpackage/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/bookpackage"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type bookPackageRepo struct {
	data *Data
	log  *log.Helper
}

// NewBookPackageRepo .
func NewBookPackageRepo(data *Data, logger log.Logger) biz.BookPackageRepo {
	return &bookPackageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateBookPackage 创建
func (r *bookPackageRepo) CreateBookPackage(ctx context.Context, req *v1.BookPackageCreateReq) (*ent.BookPackage, error) {
	now := time.Now()
	return r.data.db.BookPackage.Create().
		SetActivityCode(req.ActivityCode).
		SetPackageName(req.PackageName).
		SetPrice(req.Price).
		SetDailyPrice(req.DailyPrice).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteBookPackage 删除
func (r *bookPackageRepo) DeleteBookPackage(ctx context.Context, req *v1.BookPackageDeleteReq) error {
	return r.data.db.BookPackage.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteBookPackage 批量删除
func (r *bookPackageRepo) BatchDeleteBookPackage(ctx context.Context, req *v1.BookPackageBatchDeleteReq) (int, error) {
	return r.data.db.BookPackage.Delete().Where(bookpackage.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateBookPackage 更新
func (r *bookPackageRepo) UpdateBookPackage(ctx context.Context, req *v1.BookPackageUpdateReq) (*ent.BookPackage, error) {
	return r.data.db.BookPackage.UpdateOne(convert.BookPackageUpdateReq2Data(req)).Save(ctx)
}

// GetBookPackage 根据Id查询
func (r *bookPackageRepo) GetBookPackage(ctx context.Context, req *v1.BookPackageReq) (*ent.BookPackage, error) {
	return r.data.db.BookPackage.Get(ctx, req.Id)
}

// PageBookPackage 分页查询
func (r *bookPackageRepo) PageBookPackage(ctx context.Context, req *v1.BookPackagePageReq) ([]*ent.BookPackage, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.BookPackage.
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
func (r *bookPackageRepo) genCondition(req *v1.BookPackageReq) []predicate.BookPackage {
	if req == nil {
		return nil
	}
	list := make([]predicate.BookPackage, 0)
	if req.Id > 0 {
		list = append(list, bookpackage.ID(req.Id))
	}
	if str.IsBlank(req.ActivityCode) {
		list = append(list, bookpackage.ActivityCodeContains(req.ActivityCode))
	}
	if str.IsBlank(req.PackageName) {
		list = append(list, bookpackage.PackageNameContains(req.PackageName))
	}
	if req.Price > 0 {
		list = append(list, bookpackage.Price(req.Price))
	}
	if req.DailyPrice > 0 {
		list = append(list, bookpackage.DailyPrice(req.DailyPrice))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, bookpackage.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, bookpackage.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, bookpackage.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, bookpackage.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, bookpackage.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, bookpackage.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, bookpackage.TenantId(req.TenantId))
	}

	return list
}
