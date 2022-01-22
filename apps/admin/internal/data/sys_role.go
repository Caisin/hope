package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysrole/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysrole"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type sysRoleRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysRoleRepo .
func NewSysRoleRepo(data *Data, logger log.Logger) biz.SysRoleRepo {
	return &sysRoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysRole 创建
func (r *sysRoleRepo) CreateSysRole(ctx context.Context, req *v1.SysRoleCreateReq) (*ent.SysRole, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.SysRole.Create().
		SetRoleName(req.RoleName).
		SetStatus(req.Status).
		SetRoleKey(req.RoleKey).
		SetRoleSort(req.RoleSort).
		SetFlag(req.Flag).
		SetRemark(req.Remark).
		SetAdmin(req.Admin).
		SetDataScope(req.DataScope).
		SetSysDept(req.SysDept).
		SetSysMenu(req.SysMenu).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteSysRole 删除
func (r *sysRoleRepo) DeleteSysRole(ctx context.Context, req *v1.SysRoleDeleteReq) error {
	return r.data.db.SysRole.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysRole 批量删除
func (r *sysRoleRepo) BatchDeleteSysRole(ctx context.Context, req *v1.SysRoleBatchDeleteReq) (int, error) {
	return r.data.db.SysRole.Delete().Where(sysrole.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysRole 更新
func (r *sysRoleRepo) UpdateSysRole(ctx context.Context, req *v1.SysRoleUpdateReq) (*ent.SysRole, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	data := convert.SysRoleUpdateReq2Data(req)
	data.UpdateBy = claims.UserId
	return r.data.db.SysRole.UpdateOne(data).Save(ctx)
}

// GetSysRole 根据Id查询
func (r *sysRoleRepo) GetSysRole(ctx context.Context, req *v1.SysRoleReq) (*ent.SysRole, error) {
	return r.data.db.SysRole.Get(ctx, req.Id)
}

// PageSysRole 分页查询
func (r *sysRoleRepo) PageSysRole(ctx context.Context, req *v1.SysRolePageReq) ([]*ent.SysRole, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.SysRole.
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
func (r *sysRoleRepo) genCondition(req *v1.SysRoleReq) []predicate.SysRole {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysRole, 0)
	if req.Id > 0 {
		list = append(list, sysrole.ID(req.Id))
	}
	if str.IsBlank(req.RoleName) {
		list = append(list, sysrole.RoleNameContains(req.RoleName))
	}
	if str.IsBlank(req.Status) {
		list = append(list, sysrole.StatusContains(req.Status))
	}
	if str.IsBlank(req.RoleKey) {
		list = append(list, sysrole.RoleKeyContains(req.RoleKey))
	}
	if req.RoleSort > 0 {
		list = append(list, sysrole.RoleSort(req.RoleSort))
	}
	if str.IsBlank(req.Flag) {
		list = append(list, sysrole.FlagContains(req.Flag))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, sysrole.RemarkContains(req.Remark))
	}
	list = append(list, sysrole.Admin(req.Admin))
	if str.IsBlank(req.DataScope) {
		list = append(list, sysrole.DataScopeContains(req.DataScope))
	}
	if str.IsBlank(req.SysDept) {
		list = append(list, sysrole.SysDeptContains(req.SysDept))
	}
	if str.IsBlank(req.SysMenu) {
		list = append(list, sysrole.SysMenuContains(req.SysMenu))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysrole.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysrole.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysrole.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysrole.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysrole.TenantId(req.TenantId))
	}

	return list
}
