package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysjoblog/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysjoblog"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type sysJobLogRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysJobLogRepo .
func NewSysJobLogRepo(data *Data, logger log.Logger) biz.SysJobLogRepo {
	return &sysJobLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysJobLog 创建
func (r *sysJobLogRepo) CreateSysJobLog(ctx context.Context, req *v1.SysJobLogCreateReq) (*ent.SysJobLog, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.SysJobLog.Create().
		SetJobId(req.JobId).
		SetJobName(req.JobName).
		SetEntryId(req.EntryId).
		SetStatus(req.Status).
		SetDuration(req.Duration.AsDuration()).
		SetInfo(req.Info).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteSysJobLog 删除
func (r *sysJobLogRepo) DeleteSysJobLog(ctx context.Context, req *v1.SysJobLogDeleteReq) error {
	return r.data.db.SysJobLog.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysJobLog 批量删除
func (r *sysJobLogRepo) BatchDeleteSysJobLog(ctx context.Context, req *v1.SysJobLogBatchDeleteReq) (int, error) {
	return r.data.db.SysJobLog.Delete().Where(sysjoblog.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysJobLog 更新
func (r *sysJobLogRepo) UpdateSysJobLog(ctx context.Context, req *v1.SysJobLogUpdateReq) (*ent.SysJobLog, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.SysJobLog.UpdateOneID(req.Id).
		SetJobId(req.JobId).
		SetJobName(req.JobName).
		SetEntryId(req.EntryId).
		SetStatus(req.Status).
		SetDuration(req.Duration.AsDuration()).
		SetInfo(req.Info).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetSysJobLog 根据Id查询
func (r *sysJobLogRepo) GetSysJobLog(ctx context.Context, req *v1.SysJobLogReq) (*ent.SysJobLog, error) {
	return r.data.db.SysJobLog.Get(ctx, req.Id)
}

// PageSysJobLog 分页查询
func (r *sysJobLogRepo) PageSysJobLog(ctx context.Context, req *v1.SysJobLogPageReq) ([]*ent.SysJobLog, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.SysJobLog.
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
func (r *sysJobLogRepo) genCondition(req *v1.SysJobLogReq) []predicate.SysJobLog {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysJobLog, 0)
	if req.Id > 0 {
		list = append(list, sysjoblog.ID(req.Id))
	}
	if req.JobId > 0 {
		list = append(list, sysjoblog.JobId(req.JobId))
	}
	if str.IsNotBlank(req.JobName) {
		list = append(list, sysjoblog.JobNameContains(req.JobName))
	}
	if req.EntryId > 0 {
		list = append(list, sysjoblog.EntryId(req.EntryId))
	}
	list = append(list, sysjoblog.Status(req.Status))
	if req.Duration.AsDuration() > 0 {
		list = append(list, sysjoblog.Duration(req.Duration.AsDuration()))
	}
	if str.IsNotBlank(req.Info) {
		list = append(list, sysjoblog.InfoContains(req.Info))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysjoblog.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysjoblog.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysjoblog.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysjoblog.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysjoblog.TenantId(req.TenantId))
	}

	return list
}
