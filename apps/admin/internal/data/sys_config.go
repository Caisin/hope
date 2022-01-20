package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysconfig/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysconfig"
	"hope/pkg/util/str"
	"time"
)

type sysConfigRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysConfigRepo .
func NewSysConfigRepo(data *Data, logger log.Logger) biz.SysConfigRepo {
	return &sysConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysConfig 创建
func (r *sysConfigRepo) CreateSysConfig(ctx context.Context, req *v1.SysConfigCreateReq) (*ent.SysConfig, error) {
	now := time.Now()
	return r.data.db.SysConfig.Create().
		SetConfigName(req.ConfigName).
		SetConfigKey(req.ConfigKey).
		SetConfigValue(req.ConfigValue).
		SetConfigType(req.ConfigType).
		SetIsFrontend(req.IsFrontend).
		SetState(sysconfig.State(req.State)).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteSysConfig 删除
func (r *sysConfigRepo) DeleteSysConfig(ctx context.Context, req *v1.SysConfigDeleteReq) error {
	return r.data.db.SysConfig.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysConfig 批量删除
func (r *sysConfigRepo) BatchDeleteSysConfig(ctx context.Context, req *v1.SysConfigBatchDeleteReq) (int, error) {
	return r.data.db.SysConfig.Delete().Where(sysconfig.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysConfig 更新
func (r *sysConfigRepo) UpdateSysConfig(ctx context.Context, req *v1.SysConfigUpdateReq) (*ent.SysConfig, error) {
	return r.data.db.SysConfig.UpdateOne(convert.SysConfigUpdateReq2Data(req)).Save(ctx)
}

// GetSysConfig 根据Id查询
func (r *sysConfigRepo) GetSysConfig(ctx context.Context, req *v1.SysConfigReq) (*ent.SysConfig, error) {
	return r.data.db.SysConfig.Get(ctx, req.Id)
}

// PageSysConfig 分页查询
func (r *sysConfigRepo) PageSysConfig(ctx context.Context, req *v1.SysConfigPageReq) ([]*ent.SysConfig, error) {
	pagin := req.Pagin
	query := r.data.db.SysConfig.
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
func (r *sysConfigRepo) genCondition(req *v1.SysConfigReq) []predicate.SysConfig {
	list := make([]predicate.SysConfig, 0)
	if req.Id > 0 {
		list = append(list, sysconfig.ID(req.Id))
	}
	if str.IsBlank(req.ConfigName) {
		list = append(list, sysconfig.ConfigNameContains(req.ConfigName))
	}
	if str.IsBlank(req.ConfigKey) {
		list = append(list, sysconfig.ConfigKeyContains(req.ConfigKey))
	}
	if str.IsBlank(req.ConfigValue) {
		list = append(list, sysconfig.ConfigValueContains(req.ConfigValue))
	}
	if str.IsBlank(req.ConfigType) {
		list = append(list, sysconfig.ConfigTypeContains(req.ConfigType))
	}
	if req.IsFrontend > 0 {
		list = append(list, sysconfig.IsFrontend(req.IsFrontend))
	}
	state := sysconfig.State(req.State)
	if sysconfig.StateValidator(state) == nil {
		list = append(list, sysconfig.StateEQ(state))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, sysconfig.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysconfig.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysconfig.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysconfig.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysconfig.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysconfig.TenantId(req.TenantId))
	}

	return list
}
