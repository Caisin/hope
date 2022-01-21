package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/adchangelog/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/adchangelog"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type adChangeLogRepo struct {
	data *Data
	log  *log.Helper
}

// NewAdChangeLogRepo .
func NewAdChangeLogRepo(data *Data, logger log.Logger) biz.AdChangeLogRepo {
	return &adChangeLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateAdChangeLog 创建
func (r *adChangeLogRepo) CreateAdChangeLog(ctx context.Context, req *v1.AdChangeLogCreateReq) (*ent.AdChangeLog, error) {
	now := time.Now()
	return r.data.db.AdChangeLog.Create().
		SetUserId(req.UserId).
		SetAdId(req.AdId).
		SetChId(req.ChId).
		SetDeviceId(req.DeviceId).
		SetExtInfo(req.ExtInfo).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteAdChangeLog 删除
func (r *adChangeLogRepo) DeleteAdChangeLog(ctx context.Context, req *v1.AdChangeLogDeleteReq) error {
	return r.data.db.AdChangeLog.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteAdChangeLog 批量删除
func (r *adChangeLogRepo) BatchDeleteAdChangeLog(ctx context.Context, req *v1.AdChangeLogBatchDeleteReq) (int, error) {
	return r.data.db.AdChangeLog.Delete().Where(adchangelog.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateAdChangeLog 更新
func (r *adChangeLogRepo) UpdateAdChangeLog(ctx context.Context, req *v1.AdChangeLogUpdateReq) (*ent.AdChangeLog, error) {
	return r.data.db.AdChangeLog.UpdateOne(convert.AdChangeLogUpdateReq2Data(req)).Save(ctx)
}

// GetAdChangeLog 根据Id查询
func (r *adChangeLogRepo) GetAdChangeLog(ctx context.Context, req *v1.AdChangeLogReq) (*ent.AdChangeLog, error) {
	return r.data.db.AdChangeLog.Get(ctx, req.Id)
}

// PageAdChangeLog 分页查询
func (r *adChangeLogRepo) PageAdChangeLog(ctx context.Context, req *v1.AdChangeLogPageReq) ([]*ent.AdChangeLog, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.AdChangeLog.
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
func (r *adChangeLogRepo) genCondition(req *v1.AdChangeLogReq) []predicate.AdChangeLog {
	if req == nil {
		return nil
	}
	list := make([]predicate.AdChangeLog, 0)
	if req.Id > 0 {
		list = append(list, adchangelog.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, adchangelog.UserId(req.UserId))
	}
	if str.IsBlank(req.AdId) {
		list = append(list, adchangelog.AdIdContains(req.AdId))
	}
	if req.ChId > 0 {
		list = append(list, adchangelog.ChId(req.ChId))
	}
	if str.IsBlank(req.DeviceId) {
		list = append(list, adchangelog.DeviceIdContains(req.DeviceId))
	}
	if str.IsBlank(req.ExtInfo) {
		list = append(list, adchangelog.ExtInfoContains(req.ExtInfo))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, adchangelog.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, adchangelog.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, adchangelog.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, adchangelog.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, adchangelog.TenantId(req.TenantId))
	}

	return list
}
