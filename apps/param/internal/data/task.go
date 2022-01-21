package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/task/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/convert"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/task"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type taskRepo struct {
	data *Data
	log  *log.Helper
}

// NewTaskRepo .
func NewTaskRepo(data *Data, logger log.Logger) biz.TaskRepo {
	return &taskRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateTask 创建
func (r *taskRepo) CreateTask(ctx context.Context, req *v1.TaskCreateReq) (*ent.Task, error) {
	now := time.Now()
	return r.data.db.Task.Create().
		SetTaskName(req.TaskName).
		SetTaskGroup(req.TaskGroup).
		SetUnit(req.Unit).
		SetTopic(req.Topic).
		SetFunction(req.Function).
		SetTaskCode(req.TaskCode).
		SetPreTask(req.PreTask).
		SetNovelId(req.NovelId).
		SetCycleType(req.CycleType).
		SetRemark(req.Remark).
		SetAmount(req.Amount).
		SetReward(req.Reward).
		SetAmountItem(req.AmountItem).
		SetRewardItem(req.RewardItem).
		SetTargetNames(req.TargetNames).
		SetTargetAmounts(req.TargetAmounts).
		SetStatus(req.Status).
		SetSortNum(req.SortNum).
		SetActionType(req.ActionType).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteTask 删除
func (r *taskRepo) DeleteTask(ctx context.Context, req *v1.TaskDeleteReq) error {
	return r.data.db.Task.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteTask 批量删除
func (r *taskRepo) BatchDeleteTask(ctx context.Context, req *v1.TaskBatchDeleteReq) (int, error) {
	return r.data.db.Task.Delete().Where(task.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateTask 更新
func (r *taskRepo) UpdateTask(ctx context.Context, req *v1.TaskUpdateReq) (*ent.Task, error) {
	return r.data.db.Task.UpdateOne(convert.TaskUpdateReq2Data(req)).Save(ctx)
}

// GetTask 根据Id查询
func (r *taskRepo) GetTask(ctx context.Context, req *v1.TaskReq) (*ent.Task, error) {
	return r.data.db.Task.Get(ctx, req.Id)
}

// PageTask 分页查询
func (r *taskRepo) PageTask(ctx context.Context, req *v1.TaskPageReq) ([]*ent.Task, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.Task.
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
func (r *taskRepo) genCondition(req *v1.TaskReq) []predicate.Task {
	if req == nil {
		return nil
	}
	list := make([]predicate.Task, 0)
	if req.Id > 0 {
		list = append(list, task.ID(req.Id))
	}
	if str.IsBlank(req.TaskName) {
		list = append(list, task.TaskNameContains(req.TaskName))
	}
	if str.IsBlank(req.TaskGroup) {
		list = append(list, task.TaskGroupContains(req.TaskGroup))
	}
	if str.IsBlank(req.Unit) {
		list = append(list, task.UnitContains(req.Unit))
	}
	if str.IsBlank(req.Topic) {
		list = append(list, task.TopicContains(req.Topic))
	}
	if str.IsBlank(req.Function) {
		list = append(list, task.FunctionContains(req.Function))
	}
	if str.IsBlank(req.TaskCode) {
		list = append(list, task.TaskCodeContains(req.TaskCode))
	}
	if req.PreTask > 0 {
		list = append(list, task.PreTask(req.PreTask))
	}
	if req.NovelId > 0 {
		list = append(list, task.NovelId(req.NovelId))
	}
	if str.IsBlank(req.CycleType) {
		list = append(list, task.CycleTypeContains(req.CycleType))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, task.RemarkContains(req.Remark))
	}
	if req.Amount > 0 {
		list = append(list, task.Amount(req.Amount))
	}
	if req.Reward > 0 {
		list = append(list, task.Reward(req.Reward))
	}
	if req.AmountItem > 0 {
		list = append(list, task.AmountItem(req.AmountItem))
	}
	if req.RewardItem > 0 {
		list = append(list, task.RewardItem(req.RewardItem))
	}
	if str.IsBlank(req.TargetNames) {
		list = append(list, task.TargetNamesContains(req.TargetNames))
	}
	if str.IsBlank(req.TargetAmounts) {
		list = append(list, task.TargetAmountsContains(req.TargetAmounts))
	}
	if req.SortNum > 0 {
		list = append(list, task.SortNum(req.SortNum))
	}
	if str.IsBlank(req.ActionType) {
		list = append(list, task.ActionTypeContains(req.ActionType))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, task.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, task.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, task.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, task.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, task.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, task.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, task.TenantId(req.TenantId))
	}

	return list
}
