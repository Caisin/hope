package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysjob/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysjob"
	"hope/pkg/util/str"
	"time"
)

type sysJobRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysJobRepo .
func NewSysJobRepo(data *Data, logger log.Logger) biz.SysJobRepo {
	return &sysJobRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysJob 创建
func (r *sysJobRepo) CreateSysJob(ctx context.Context, req *v1.SysJobCreateReq) (*ent.SysJob, error) {
	now := time.Now()
	return r.data.db.SysJob.Create().
		SetJobName(req.JobName).
		SetJobGroup(req.JobGroup).
		SetJobType(req.JobType).
		SetCronExpression(req.CronExpression).
		SetInvokeTarget(req.InvokeTarget).
		SetArgs(req.Args).
		SetExecPolicy(req.ExecPolicy).
		SetConcurrent(req.Concurrent).
		SetState(sysjob.State(req.State)).
		SetEntryId(req.EntryId).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteSysJob 删除
func (r *sysJobRepo) DeleteSysJob(ctx context.Context, req *v1.SysJobDeleteReq) error {
	return r.data.db.SysJob.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysJob 批量删除
func (r *sysJobRepo) BatchDeleteSysJob(ctx context.Context, req *v1.SysJobBatchDeleteReq) (int, error) {
	return r.data.db.SysJob.Delete().Where(sysjob.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysJob 更新
func (r *sysJobRepo) UpdateSysJob(ctx context.Context, req *v1.SysJobUpdateReq) (*ent.SysJob, error) {
	return r.data.db.SysJob.UpdateOne(convert.SysJobUpdateReq2Data(req)).Save(ctx)
}

// GetSysJob 根据Id查询
func (r *sysJobRepo) GetSysJob(ctx context.Context, req *v1.SysJobReq) (*ent.SysJob, error) {
	return r.data.db.SysJob.Get(ctx, req.Id)
}

// PageSysJob 分页查询
func (r *sysJobRepo) PageSysJob(ctx context.Context, req *v1.SysJobPageReq) ([]*ent.SysJob, error) {
	pagin := req.Pagin
	query := r.data.db.SysJob.
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
func (r *sysJobRepo) genCondition(req *v1.SysJobReq) []predicate.SysJob {
	list := make([]predicate.SysJob, 0)
	if req.Id > 0 {
		list = append(list, sysjob.ID(req.Id))
	}
	if str.IsBlank(req.JobName) {
		list = append(list, sysjob.JobNameContains(req.JobName))
	}
	if str.IsBlank(req.JobGroup) {
		list = append(list, sysjob.JobGroupContains(req.JobGroup))
	}
	if req.JobType > 0 {
		list = append(list, sysjob.JobType(req.JobType))
	}
	if str.IsBlank(req.CronExpression) {
		list = append(list, sysjob.CronExpressionContains(req.CronExpression))
	}
	if str.IsBlank(req.InvokeTarget) {
		list = append(list, sysjob.InvokeTargetContains(req.InvokeTarget))
	}
	if str.IsBlank(req.Args) {
		list = append(list, sysjob.ArgsContains(req.Args))
	}
	if req.ExecPolicy > 0 {
		list = append(list, sysjob.ExecPolicy(req.ExecPolicy))
	}
	if req.Concurrent > 0 {
		list = append(list, sysjob.Concurrent(req.Concurrent))
	}
	state := sysjob.State(req.State)
	if sysjob.StateValidator(state) == nil {
		list = append(list, sysjob.StateEQ(state))
	}
	if req.EntryId > 0 {
		list = append(list, sysjob.EntryId(req.EntryId))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysjob.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysjob.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysjob.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysjob.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysjob.TenantId(req.TenantId))
	}

	return list
}
