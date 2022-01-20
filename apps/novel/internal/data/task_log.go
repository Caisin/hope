package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/tasklog/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/tasklog"
	"hope/pkg/util/str"
	"time"
)

type taskLogRepo struct {
	data *Data
	log  *log.Helper
}

// NewTaskLogRepo .
func NewTaskLogRepo(data *Data, logger log.Logger) biz.TaskLogRepo {
	return &taskLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateTaskLog 创建
func (r *taskLogRepo) CreateTaskLog(ctx context.Context, req *v1.TaskLogCreateReq) (*ent.TaskLog, error) {
	now := time.Now()
	return r.data.db.TaskLog.Create().
		SetUserId(req.UserId).
		SetTaskGroup(req.TaskGroup).
		SetTaskCode(req.TaskCode).
		SetTaskId(req.TaskId).
		SetTaskName(req.TaskName).
		SetAmount(req.Amount).
		SetReward(req.Reward).
		SetAmountItem(req.AmountItem).
		SetRewardItem(req.RewardItem).
		SetTargetAmount(req.TargetAmount).
		SetDoneAmount(req.DoneAmount).
		SetState(req.State).
		SetDoneAt(req.DoneAt.AsTime()).
		SetObtainAt(req.ObtainAt.AsTime()).
		SetDoneTimes(req.DoneTimes).
		SetAllTimes(req.AllTimes).
		SetEffectTime(req.EffectTime.AsTime()).
		SetExpiredTime(req.ExpiredTime.AsTime()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteTaskLog 删除
func (r *taskLogRepo) DeleteTaskLog(ctx context.Context, req *v1.TaskLogDeleteReq) error {
	return r.data.db.TaskLog.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteTaskLog 批量删除
func (r *taskLogRepo) BatchDeleteTaskLog(ctx context.Context, req *v1.TaskLogBatchDeleteReq) (int, error) {
	return r.data.db.TaskLog.Delete().Where(tasklog.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateTaskLog 更新
func (r *taskLogRepo) UpdateTaskLog(ctx context.Context, req *v1.TaskLogUpdateReq) (*ent.TaskLog, error) {
	return r.data.db.TaskLog.UpdateOne(convert.TaskLogUpdateReq2Data(req)).Save(ctx)
}

// GetTaskLog 根据Id查询
func (r *taskLogRepo) GetTaskLog(ctx context.Context, req *v1.TaskLogReq) (*ent.TaskLog, error) {
	return r.data.db.TaskLog.Get(ctx, req.Id)
}

// PageTaskLog 分页查询
func (r *taskLogRepo) PageTaskLog(ctx context.Context, req *v1.TaskLogPageReq) ([]*ent.TaskLog, error) {
	pagin := req.Pagin
	query := r.data.db.TaskLog.
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
func (r *taskLogRepo) genCondition(req *v1.TaskLogReq) []predicate.TaskLog {
	list := make([]predicate.TaskLog, 0)
	if req.Id > 0 {
		list = append(list, tasklog.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, tasklog.UserId(req.UserId))
	}
	if str.IsBlank(req.TaskGroup) {
		list = append(list, tasklog.TaskGroupContains(req.TaskGroup))
	}
	if str.IsBlank(req.TaskCode) {
		list = append(list, tasklog.TaskCodeContains(req.TaskCode))
	}
	if req.TaskId > 0 {
		list = append(list, tasklog.TaskId(req.TaskId))
	}
	if str.IsBlank(req.TaskName) {
		list = append(list, tasklog.TaskNameContains(req.TaskName))
	}
	if req.Amount > 0 {
		list = append(list, tasklog.Amount(req.Amount))
	}
	if req.Reward > 0 {
		list = append(list, tasklog.Reward(req.Reward))
	}
	if req.AmountItem > 0 {
		list = append(list, tasklog.AmountItem(req.AmountItem))
	}
	if req.RewardItem > 0 {
		list = append(list, tasklog.RewardItem(req.RewardItem))
	}
	if req.TargetAmount > 0 {
		list = append(list, tasklog.TargetAmount(req.TargetAmount))
	}
	if req.DoneAmount > 0 {
		list = append(list, tasklog.DoneAmount(req.DoneAmount))
	}
	if req.State > 0 {
		list = append(list, tasklog.State(req.State))
	}
	if req.DoneAt.IsValid() && !req.DoneAt.AsTime().IsZero() {
		list = append(list, tasklog.DoneAtGTE(req.DoneAt.AsTime()))
	}
	if req.ObtainAt.IsValid() && !req.ObtainAt.AsTime().IsZero() {
		list = append(list, tasklog.ObtainAtGTE(req.ObtainAt.AsTime()))
	}
	if req.DoneTimes > 0 {
		list = append(list, tasklog.DoneTimes(req.DoneTimes))
	}
	if req.AllTimes > 0 {
		list = append(list, tasklog.AllTimes(req.AllTimes))
	}
	if req.EffectTime.IsValid() && !req.EffectTime.AsTime().IsZero() {
		list = append(list, tasklog.EffectTimeGTE(req.EffectTime.AsTime()))
	}
	if req.ExpiredTime.IsValid() && !req.ExpiredTime.AsTime().IsZero() {
		list = append(list, tasklog.ExpiredTimeGTE(req.ExpiredTime.AsTime()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, tasklog.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, tasklog.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, tasklog.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, tasklog.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, tasklog.TenantId(req.TenantId))
	}

	return list
}
