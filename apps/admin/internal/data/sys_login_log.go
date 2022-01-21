package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysloginlog/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysloginlog"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type sysLoginLogRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysLoginLogRepo .
func NewSysLoginLogRepo(data *Data, logger log.Logger) biz.SysLoginLogRepo {
	return &sysLoginLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysLoginLog 创建
func (r *sysLoginLogRepo) CreateSysLoginLog(ctx context.Context, req *v1.SysLoginLogCreateReq) (*ent.SysLoginLog, error) {
	now := time.Now()
	return r.data.db.SysLoginLog.Create().
		SetUserId(req.UserId).
		SetStatus(req.Status).
		SetIpaddr(req.Ipaddr).
		SetLoginLocation(req.LoginLocation).
		SetBrowser(req.Browser).
		SetOs(req.Os).
		SetPlatform(req.Platform).
		SetLoginTime(req.LoginTime.AsTime()).
		SetRemark(req.Remark).
		SetMsg(req.Msg).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteSysLoginLog 删除
func (r *sysLoginLogRepo) DeleteSysLoginLog(ctx context.Context, req *v1.SysLoginLogDeleteReq) error {
	return r.data.db.SysLoginLog.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysLoginLog 批量删除
func (r *sysLoginLogRepo) BatchDeleteSysLoginLog(ctx context.Context, req *v1.SysLoginLogBatchDeleteReq) (int, error) {
	return r.data.db.SysLoginLog.Delete().Where(sysloginlog.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysLoginLog 更新
func (r *sysLoginLogRepo) UpdateSysLoginLog(ctx context.Context, req *v1.SysLoginLogUpdateReq) (*ent.SysLoginLog, error) {
	return r.data.db.SysLoginLog.UpdateOne(convert.SysLoginLogUpdateReq2Data(req)).Save(ctx)
}

// GetSysLoginLog 根据Id查询
func (r *sysLoginLogRepo) GetSysLoginLog(ctx context.Context, req *v1.SysLoginLogReq) (*ent.SysLoginLog, error) {
	return r.data.db.SysLoginLog.Get(ctx, req.Id)
}

// PageSysLoginLog 分页查询
func (r *sysLoginLogRepo) PageSysLoginLog(ctx context.Context, req *v1.SysLoginLogPageReq) ([]*ent.SysLoginLog, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.SysLoginLog.
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
func (r *sysLoginLogRepo) genCondition(req *v1.SysLoginLogReq) []predicate.SysLoginLog {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysLoginLog, 0)
	if req.Id > 0 {
		list = append(list, sysloginlog.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, sysloginlog.UserId(req.UserId))
	}
	if str.IsBlank(req.Status) {
		list = append(list, sysloginlog.StatusContains(req.Status))
	}
	if str.IsBlank(req.Ipaddr) {
		list = append(list, sysloginlog.IpaddrContains(req.Ipaddr))
	}
	if str.IsBlank(req.LoginLocation) {
		list = append(list, sysloginlog.LoginLocationContains(req.LoginLocation))
	}
	if str.IsBlank(req.Browser) {
		list = append(list, sysloginlog.BrowserContains(req.Browser))
	}
	if str.IsBlank(req.Os) {
		list = append(list, sysloginlog.OsContains(req.Os))
	}
	if str.IsBlank(req.Platform) {
		list = append(list, sysloginlog.PlatformContains(req.Platform))
	}
	if req.LoginTime.IsValid() && !req.LoginTime.AsTime().IsZero() {
		list = append(list, sysloginlog.LoginTimeGTE(req.LoginTime.AsTime()))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, sysloginlog.RemarkContains(req.Remark))
	}
	if str.IsBlank(req.Msg) {
		list = append(list, sysloginlog.MsgContains(req.Msg))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysloginlog.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysloginlog.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysloginlog.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysloginlog.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysloginlog.TenantId(req.TenantId))
	}

	return list
}
