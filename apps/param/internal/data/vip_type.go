package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/viptype/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/viptype"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type vipTypeRepo struct {
	data *Data
	log  *log.Helper
}

// NewVipTypeRepo .
func NewVipTypeRepo(data *Data, logger log.Logger) biz.VipTypeRepo {
	return &vipTypeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateVipType 创建
func (r *vipTypeRepo) CreateVipType(ctx context.Context, req *v1.VipTypeCreateReq) (*ent.VipType, error) {
	now := time.Now()
	return r.data.db.VipType.Create().
		SetVipName(req.VipName).
		SetIsSuper(req.IsSuper).
		SetValidDays(req.ValidDays).
		SetDiscountRate(req.DiscountRate).
		SetAvatarId(req.AvatarId).
		SetSummary(req.Summary).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteVipType 删除
func (r *vipTypeRepo) DeleteVipType(ctx context.Context, req *v1.VipTypeDeleteReq) error {
	return r.data.db.VipType.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteVipType 批量删除
func (r *vipTypeRepo) BatchDeleteVipType(ctx context.Context, req *v1.VipTypeBatchDeleteReq) (int, error) {
	return r.data.db.VipType.Delete().Where(viptype.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateVipType 更新
func (r *vipTypeRepo) UpdateVipType(ctx context.Context, req *v1.VipTypeUpdateReq) (*ent.VipType, error) {
	return r.data.db.VipType.UpdateOne(convert.VipTypeUpdateReq2Data(req)).Save(ctx)
}

// GetVipType 根据Id查询
func (r *vipTypeRepo) GetVipType(ctx context.Context, req *v1.VipTypeReq) (*ent.VipType, error) {
	return r.data.db.VipType.Get(ctx, req.Id)
}

// PageVipType 分页查询
func (r *vipTypeRepo) PageVipType(ctx context.Context, req *v1.VipTypePageReq) ([]*ent.VipType, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.VipType.
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
func (r *vipTypeRepo) genCondition(req *v1.VipTypeReq) []predicate.VipType {
	if req == nil {
		return nil
	}
	list := make([]predicate.VipType, 0)
	if req.Id > 0 {
		list = append(list, viptype.ID(req.Id))
	}
	if str.IsBlank(req.VipName) {
		list = append(list, viptype.VipNameContains(req.VipName))
	}
	if req.ValidDays > 0 {
		list = append(list, viptype.ValidDays(req.ValidDays))
	}
	if req.DiscountRate > 0 {
		list = append(list, viptype.DiscountRate(req.DiscountRate))
	}
	if req.AvatarId > 0 {
		list = append(list, viptype.AvatarId(req.AvatarId))
	}
	if str.IsBlank(req.Summary) {
		list = append(list, viptype.SummaryContains(req.Summary))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, viptype.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, viptype.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, viptype.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, viptype.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, viptype.TenantId(req.TenantId))
	}

	return list
}
