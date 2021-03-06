package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/userresourcerecord/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/userresourcerecord"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type userResourceRecordRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserResourceRecordRepo .
func NewUserResourceRecordRepo(data *Data, logger log.Logger) biz.UserResourceRecordRepo {
	return &userResourceRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUserResourceRecord 创建
func (r *userResourceRecordRepo) CreateUserResourceRecord(ctx context.Context, req *v1.UserResourceRecordCreateReq) (*ent.UserResourceRecord, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.UserResourceRecord.Create().
		SetUserId(req.UserId).
		SetResId(req.ResId).
		SetDef(req.Def).
		SetName(req.Name).
		SetURL(req.Url).
		SetResType(req.ResType).
		SetRemark(req.Remark).
		SetState(req.State).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteUserResourceRecord 删除
func (r *userResourceRecordRepo) DeleteUserResourceRecord(ctx context.Context, req *v1.UserResourceRecordDeleteReq) error {
	return r.data.db.UserResourceRecord.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteUserResourceRecord 批量删除
func (r *userResourceRecordRepo) BatchDeleteUserResourceRecord(ctx context.Context, req *v1.UserResourceRecordBatchDeleteReq) (int, error) {
	return r.data.db.UserResourceRecord.Delete().Where(userresourcerecord.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateUserResourceRecord 更新
func (r *userResourceRecordRepo) UpdateUserResourceRecord(ctx context.Context, req *v1.UserResourceRecordUpdateReq) (*ent.UserResourceRecord, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.UserResourceRecord.UpdateOneID(req.Id).
		SetUserId(req.UserId).
		SetResId(req.ResId).
		SetDef(req.Def).
		SetName(req.Name).
		SetURL(req.Url).
		SetResType(req.ResType).
		SetRemark(req.Remark).
		SetState(req.State).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetUserResourceRecord 根据Id查询
func (r *userResourceRecordRepo) GetUserResourceRecord(ctx context.Context, req *v1.UserResourceRecordReq) (*ent.UserResourceRecord, error) {
	return r.data.db.UserResourceRecord.Get(ctx, req.Id)
}

// PageUserResourceRecord 分页查询
func (r *userResourceRecordRepo) PageUserResourceRecord(ctx context.Context, req *v1.UserResourceRecordPageReq) ([]*ent.UserResourceRecord, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.UserResourceRecord.
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
func (r *userResourceRecordRepo) genCondition(req *v1.UserResourceRecordReq) []predicate.UserResourceRecord {
	if req == nil {
		return nil
	}
	list := make([]predicate.UserResourceRecord, 0)
	if req.Id > 0 {
		list = append(list, userresourcerecord.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, userresourcerecord.UserId(req.UserId))
	}
	if req.ResId > 0 {
		list = append(list, userresourcerecord.ResId(req.ResId))
	}
	list = append(list, userresourcerecord.Def(req.Def))
	if str.IsNotBlank(req.Name) {
		list = append(list, userresourcerecord.NameContains(req.Name))
	}
	if str.IsNotBlank(req.Url) {
		list = append(list, userresourcerecord.URLContains(req.Url))
	}
	if str.IsNotBlank(req.ResType) {
		list = append(list, userresourcerecord.ResTypeContains(req.ResType))
	}
	if str.IsNotBlank(req.Remark) {
		list = append(list, userresourcerecord.RemarkContains(req.Remark))
	}
	if req.State > 0 {
		list = append(list, userresourcerecord.State(req.State))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, userresourcerecord.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, userresourcerecord.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, userresourcerecord.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, userresourcerecord.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, userresourcerecord.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, userresourcerecord.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, userresourcerecord.TenantId(req.TenantId))
	}

	return list
}
