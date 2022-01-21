package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/activitycomponent/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/activitycomponent"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type activityComponentRepo struct {
	data *Data
	log  *log.Helper
}

// NewActivityComponentRepo .
func NewActivityComponentRepo(data *Data, logger log.Logger) biz.ActivityComponentRepo {
	return &activityComponentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateActivityComponent 创建
func (r *activityComponentRepo) CreateActivityComponent(ctx context.Context, req *v1.ActivityComponentCreateReq) (*ent.ActivityComponent, error) {
	now := time.Now()
	return r.data.db.ActivityComponent.Create().
		SetActivityCode(req.ActivityCode).
		SetComponentType(req.ComponentType).
		SetPolicy(req.Policy).
		SetVipDays(req.VipDays).
		SetMinConsume(req.MinConsume).
		SetMaxConsume(req.MaxConsume).
		SetMinPayNum(req.MinPayNum).
		SetPayTimes(req.PayTimes).
		SetPayAmount(req.PayAmount).
		SetRegDays(req.RegDays).
		SetSummary(req.Summary).
		SetAssetItemId(req.AssetItemId).
		SetAmount(req.Amount).
		SetResId(req.ResId).
		SetResDays(req.ResDays).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteActivityComponent 删除
func (r *activityComponentRepo) DeleteActivityComponent(ctx context.Context, req *v1.ActivityComponentDeleteReq) error {
	return r.data.db.ActivityComponent.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteActivityComponent 批量删除
func (r *activityComponentRepo) BatchDeleteActivityComponent(ctx context.Context, req *v1.ActivityComponentBatchDeleteReq) (int, error) {
	return r.data.db.ActivityComponent.Delete().Where(activitycomponent.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateActivityComponent 更新
func (r *activityComponentRepo) UpdateActivityComponent(ctx context.Context, req *v1.ActivityComponentUpdateReq) (*ent.ActivityComponent, error) {
	return r.data.db.ActivityComponent.UpdateOne(convert.ActivityComponentUpdateReq2Data(req)).Save(ctx)
}

// GetActivityComponent 根据Id查询
func (r *activityComponentRepo) GetActivityComponent(ctx context.Context, req *v1.ActivityComponentReq) (*ent.ActivityComponent, error) {
	return r.data.db.ActivityComponent.Get(ctx, req.Id)
}

// PageActivityComponent 分页查询
func (r *activityComponentRepo) PageActivityComponent(ctx context.Context, req *v1.ActivityComponentPageReq) ([]*ent.ActivityComponent, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.ActivityComponent.
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
func (r *activityComponentRepo) genCondition(req *v1.ActivityComponentReq) []predicate.ActivityComponent {
	if req == nil {
		return nil
	}
	list := make([]predicate.ActivityComponent, 0)
	if req.Id > 0 {
		list = append(list, activitycomponent.ID(req.Id))
	}
	if str.IsBlank(req.ActivityCode) {
		list = append(list, activitycomponent.ActivityCodeContains(req.ActivityCode))
	}
	if str.IsBlank(req.ComponentType) {
		list = append(list, activitycomponent.ComponentTypeContains(req.ComponentType))
	}
	if str.IsBlank(req.Policy) {
		list = append(list, activitycomponent.PolicyContains(req.Policy))
	}
	if req.MinConsume > 0 {
		list = append(list, activitycomponent.MinConsume(req.MinConsume))
	}
	if req.MaxConsume > 0 {
		list = append(list, activitycomponent.MaxConsume(req.MaxConsume))
	}
	if req.MinPayNum > 0 {
		list = append(list, activitycomponent.MinPayNum(req.MinPayNum))
	}
	if req.PayTimes > 0 {
		list = append(list, activitycomponent.PayTimes(req.PayTimes))
	}
	if req.PayAmount > 0 {
		list = append(list, activitycomponent.PayAmount(req.PayAmount))
	}
	if req.RegDays > 0 {
		list = append(list, activitycomponent.RegDays(req.RegDays))
	}
	if str.IsBlank(req.Summary) {
		list = append(list, activitycomponent.SummaryContains(req.Summary))
	}
	if req.AssetItemId > 0 {
		list = append(list, activitycomponent.AssetItemId(req.AssetItemId))
	}
	if req.Amount > 0 {
		list = append(list, activitycomponent.Amount(req.Amount))
	}
	if req.ResId > 0 {
		list = append(list, activitycomponent.ResId(req.ResId))
	}
	if req.ResDays > 0 {
		list = append(list, activitycomponent.ResDays(req.ResDays))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, activitycomponent.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, activitycomponent.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, activitycomponent.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, activitycomponent.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, activitycomponent.TenantId(req.TenantId))
	}

	return list
}
