package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysoperalog/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysoperalog"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type sysOperaLogRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysOperaLogRepo .
func NewSysOperaLogRepo(data *Data, logger log.Logger) biz.SysOperaLogRepo {
	return &sysOperaLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysOperaLog 创建
func (r *sysOperaLogRepo) CreateSysOperaLog(ctx context.Context, req *v1.SysOperaLogCreateReq) (*ent.SysOperaLog, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.SysOperaLog.Create().
		SetUserId(req.UserId).
		SetTitle(req.Title).
		SetRequestId(req.RequestId).
		SetBusinessType(req.BusinessType).
		SetBusinessTypes(req.BusinessTypes).
		SetMethod(req.Method).
		SetRequestMethod(req.RequestMethod).
		SetOperatorType(req.OperatorType).
		SetOperName(req.OperName).
		SetDeptName(req.DeptName).
		SetOperUrl(req.OperUrl).
		SetOperIp(req.OperIp).
		SetBrowser(req.Browser).
		SetOs(req.Os).
		SetPlatform(req.Platform).
		SetOperLocation(req.OperLocation).
		SetOperParam(req.OperParam).
		SetStatus(req.Status).
		SetOperTime(req.OperTime.AsTime()).
		SetJsonResult(req.JsonResult).
		SetRemark(req.Remark).
		SetLatencyTime(req.LatencyTime).
		SetUserAgent(req.UserAgent).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteSysOperaLog 删除
func (r *sysOperaLogRepo) DeleteSysOperaLog(ctx context.Context, req *v1.SysOperaLogDeleteReq) error {
	return r.data.db.SysOperaLog.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysOperaLog 批量删除
func (r *sysOperaLogRepo) BatchDeleteSysOperaLog(ctx context.Context, req *v1.SysOperaLogBatchDeleteReq) (int, error) {
	return r.data.db.SysOperaLog.Delete().Where(sysoperalog.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysOperaLog 更新
func (r *sysOperaLogRepo) UpdateSysOperaLog(ctx context.Context, req *v1.SysOperaLogUpdateReq) (*ent.SysOperaLog, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	data := convert.SysOperaLogUpdateReq2Data(req)
	data.UpdateBy = claims.UserId
	return r.data.db.SysOperaLog.UpdateOne(data).Save(ctx)
}

// GetSysOperaLog 根据Id查询
func (r *sysOperaLogRepo) GetSysOperaLog(ctx context.Context, req *v1.SysOperaLogReq) (*ent.SysOperaLog, error) {
	return r.data.db.SysOperaLog.Get(ctx, req.Id)
}

// PageSysOperaLog 分页查询
func (r *sysOperaLogRepo) PageSysOperaLog(ctx context.Context, req *v1.SysOperaLogPageReq) ([]*ent.SysOperaLog, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.SysOperaLog.
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
func (r *sysOperaLogRepo) genCondition(req *v1.SysOperaLogReq) []predicate.SysOperaLog {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysOperaLog, 0)
	if req.Id > 0 {
		list = append(list, sysoperalog.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, sysoperalog.UserId(req.UserId))
	}
	if str.IsBlank(req.Title) {
		list = append(list, sysoperalog.TitleContains(req.Title))
	}
	if str.IsBlank(req.RequestId) {
		list = append(list, sysoperalog.RequestIdContains(req.RequestId))
	}
	if str.IsBlank(req.BusinessType) {
		list = append(list, sysoperalog.BusinessTypeContains(req.BusinessType))
	}
	if str.IsBlank(req.BusinessTypes) {
		list = append(list, sysoperalog.BusinessTypesContains(req.BusinessTypes))
	}
	if str.IsBlank(req.Method) {
		list = append(list, sysoperalog.MethodContains(req.Method))
	}
	if str.IsBlank(req.RequestMethod) {
		list = append(list, sysoperalog.RequestMethodContains(req.RequestMethod))
	}
	if str.IsBlank(req.OperatorType) {
		list = append(list, sysoperalog.OperatorTypeContains(req.OperatorType))
	}
	if str.IsBlank(req.OperName) {
		list = append(list, sysoperalog.OperNameContains(req.OperName))
	}
	if str.IsBlank(req.DeptName) {
		list = append(list, sysoperalog.DeptNameContains(req.DeptName))
	}
	if str.IsBlank(req.OperUrl) {
		list = append(list, sysoperalog.OperUrlContains(req.OperUrl))
	}
	if str.IsBlank(req.OperIp) {
		list = append(list, sysoperalog.OperIpContains(req.OperIp))
	}
	if str.IsBlank(req.Browser) {
		list = append(list, sysoperalog.BrowserContains(req.Browser))
	}
	if str.IsBlank(req.Os) {
		list = append(list, sysoperalog.OsContains(req.Os))
	}
	if str.IsBlank(req.Platform) {
		list = append(list, sysoperalog.PlatformContains(req.Platform))
	}
	if str.IsBlank(req.OperLocation) {
		list = append(list, sysoperalog.OperLocationContains(req.OperLocation))
	}
	if str.IsBlank(req.OperParam) {
		list = append(list, sysoperalog.OperParamContains(req.OperParam))
	}
	if str.IsBlank(req.Status) {
		list = append(list, sysoperalog.StatusContains(req.Status))
	}
	if req.OperTime.IsValid() && !req.OperTime.AsTime().IsZero() {
		list = append(list, sysoperalog.OperTimeGTE(req.OperTime.AsTime()))
	}
	if str.IsBlank(req.JsonResult) {
		list = append(list, sysoperalog.JsonResultContains(req.JsonResult))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, sysoperalog.RemarkContains(req.Remark))
	}
	if str.IsBlank(req.LatencyTime) {
		list = append(list, sysoperalog.LatencyTimeContains(req.LatencyTime))
	}
	if str.IsBlank(req.UserAgent) {
		list = append(list, sysoperalog.UserAgentContains(req.UserAgent))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysoperalog.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysoperalog.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysoperalog.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysoperalog.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysoperalog.TenantId(req.TenantId))
	}

	return list
}
