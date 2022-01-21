package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/vipuser/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/vipuser"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type vipUserRepo struct {
	data *Data
	log  *log.Helper
}

// NewVipUserRepo .
func NewVipUserRepo(data *Data, logger log.Logger) biz.VipUserRepo {
	return &vipUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateVipUser 创建
func (r *vipUserRepo) CreateVipUser(ctx context.Context, req *v1.VipUserCreateReq) (*ent.VipUser, error) {
	now := time.Now()
	return r.data.db.VipUser.Create().
		SetUserId(req.UserId).
		SetVipType(req.VipType).
		SetSvipType(req.SvipType).
		SetSvipEffectTime(req.SvipEffectTime.AsTime()).
		SetSvipExpiredTime(req.SvipExpiredTime.AsTime()).
		SetRemark(req.Remark).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteVipUser 删除
func (r *vipUserRepo) DeleteVipUser(ctx context.Context, req *v1.VipUserDeleteReq) error {
	return r.data.db.VipUser.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteVipUser 批量删除
func (r *vipUserRepo) BatchDeleteVipUser(ctx context.Context, req *v1.VipUserBatchDeleteReq) (int, error) {
	return r.data.db.VipUser.Delete().Where(vipuser.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateVipUser 更新
func (r *vipUserRepo) UpdateVipUser(ctx context.Context, req *v1.VipUserUpdateReq) (*ent.VipUser, error) {
	return r.data.db.VipUser.UpdateOne(convert.VipUserUpdateReq2Data(req)).Save(ctx)
}

// GetVipUser 根据Id查询
func (r *vipUserRepo) GetVipUser(ctx context.Context, req *v1.VipUserReq) (*ent.VipUser, error) {
	return r.data.db.VipUser.Get(ctx, req.Id)
}

// PageVipUser 分页查询
func (r *vipUserRepo) PageVipUser(ctx context.Context, req *v1.VipUserPageReq) ([]*ent.VipUser, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.VipUser.
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
func (r *vipUserRepo) genCondition(req *v1.VipUserReq) []predicate.VipUser {
	if req == nil {
		return nil
	}
	list := make([]predicate.VipUser, 0)
	if req.Id > 0 {
		list = append(list, vipuser.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, vipuser.UserId(req.UserId))
	}
	if req.VipType > 0 {
		list = append(list, vipuser.VipType(req.VipType))
	}
	if req.SvipType > 0 {
		list = append(list, vipuser.SvipType(req.SvipType))
	}
	if req.SvipEffectTime.IsValid() && !req.SvipEffectTime.AsTime().IsZero() {
		list = append(list, vipuser.SvipEffectTimeGTE(req.SvipEffectTime.AsTime()))
	}
	if req.SvipExpiredTime.IsValid() && !req.SvipExpiredTime.AsTime().IsZero() {
		list = append(list, vipuser.SvipExpiredTimeGTE(req.SvipExpiredTime.AsTime()))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, vipuser.RemarkContains(req.Remark))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, vipuser.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, vipuser.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, vipuser.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, vipuser.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, vipuser.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, vipuser.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, vipuser.TenantId(req.TenantId))
	}

	return list
}
