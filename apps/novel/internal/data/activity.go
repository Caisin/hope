package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/activity/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/activity"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type activityRepo struct {
	data *Data
	log  *log.Helper
}

// NewActivityRepo .
func NewActivityRepo(data *Data, logger log.Logger) biz.ActivityRepo {
	return &activityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateActivity 创建
func (r *activityRepo) CreateActivity(ctx context.Context, req *v1.ActivityCreateReq) (*ent.Activity, error) {
	now := time.Now()
	return r.data.db.Activity.Create().
		SetActivityCode(req.ActivityCode).
		SetActivityName(req.ActivityName).
		SetSummary(req.Summary).
		SetRuleImgSc(req.RuleImgSc).
		SetRuleImgTc(req.RuleImgTc).
		SetPopupImg(req.PopupImg).
		SetRegDays(req.RegDays).
		SetCycleType(req.CycleType).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteActivity 删除
func (r *activityRepo) DeleteActivity(ctx context.Context, req *v1.ActivityDeleteReq) error {
	return r.data.db.Activity.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteActivity 批量删除
func (r *activityRepo) BatchDeleteActivity(ctx context.Context, req *v1.ActivityBatchDeleteReq) (int, error) {
	return r.data.db.Activity.Delete().Where(activity.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateActivity 更新
func (r *activityRepo) UpdateActivity(ctx context.Context, req *v1.ActivityUpdateReq) (*ent.Activity, error) {
	return r.data.db.Activity.UpdateOne(convert.ActivityUpdateReq2Data(req)).Save(ctx)
}

// GetActivity 根据Id查询
func (r *activityRepo) GetActivity(ctx context.Context, req *v1.ActivityReq) (*ent.Activity, error) {
	return r.data.db.Activity.Get(ctx, req.Id)
}

// PageActivity 分页查询
func (r *activityRepo) PageActivity(ctx context.Context, req *v1.ActivityPageReq) ([]*ent.Activity, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.Activity.
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
func (r *activityRepo) genCondition(req *v1.ActivityReq) []predicate.Activity {
	if req == nil {
		return nil
	}
	list := make([]predicate.Activity, 0)
	if req.Id > 0 {
		list = append(list, activity.ID(req.Id))
	}
	if str.IsBlank(req.ActivityCode) {
		list = append(list, activity.ActivityCodeContains(req.ActivityCode))
	}
	if str.IsBlank(req.ActivityName) {
		list = append(list, activity.ActivityNameContains(req.ActivityName))
	}
	if str.IsBlank(req.Summary) {
		list = append(list, activity.SummaryContains(req.Summary))
	}
	if str.IsBlank(req.RuleImgSc) {
		list = append(list, activity.RuleImgScContains(req.RuleImgSc))
	}
	if str.IsBlank(req.RuleImgTc) {
		list = append(list, activity.RuleImgTcContains(req.RuleImgTc))
	}
	if str.IsBlank(req.PopupImg) {
		list = append(list, activity.PopupImgContains(req.PopupImg))
	}
	if req.RegDays > 0 {
		list = append(list, activity.RegDays(req.RegDays))
	}
	if str.IsBlank(req.CycleType) {
		list = append(list, activity.CycleTypeContains(req.CycleType))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, activity.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, activity.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, activity.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, activity.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, activity.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, activity.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, activity.TenantId(req.TenantId))
	}

	return list
}
