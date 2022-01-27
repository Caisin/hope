package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysmenu/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysmenu"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type sysMenuRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysMenuRepo .
func NewSysMenuRepo(data *Data, logger log.Logger) biz.SysMenuRepo {
	return &sysMenuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysMenu 创建
func (r *sysMenuRepo) CreateSysMenu(ctx context.Context, req *v1.SysMenuCreateReq) (*ent.SysMenu, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.SysMenu.Create().
		SetParentId(req.ParentId).
		SetName(req.Name).
		SetTitle(req.Title).
		SetRedirect(req.Redirect).
		SetIcon(req.Icon).
		SetPath(req.Path).
		SetPaths(req.Paths).
		SetMenuType(req.MenuType).
		SetAction(req.Action).
		SetPermission(req.Permission).
		SetIgnoreKeepAlive(req.IgnoreKeepAlive).
		SetHideBreadcrumb(req.HideBreadcrumb).
		SetHideChildrenInMenu(req.HideChildrenInMenu).
		SetComponent(req.Component).
		SetSort(req.Sort).
		SetHideMenu(req.HideMenu).
		SetFrameSrc(req.FrameSrc).
		SetState(sysmenu.State(req.State)).
		SetCheckPermission(req.CheckPermission).
		SetOperation(req.Operation).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteSysMenu 删除
func (r *sysMenuRepo) DeleteSysMenu(ctx context.Context, req *v1.SysMenuDeleteReq) error {
	return r.data.db.SysMenu.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysMenu 批量删除
func (r *sysMenuRepo) BatchDeleteSysMenu(ctx context.Context, req *v1.SysMenuBatchDeleteReq) (int, error) {
	return r.data.db.SysMenu.Delete().Where(sysmenu.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysMenu 更新
func (r *sysMenuRepo) UpdateSysMenu(ctx context.Context, req *v1.SysMenuUpdateReq) (*ent.SysMenu, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.SysMenu.UpdateOneID(req.Id).
		SetParentId(req.ParentId).
		SetName(req.Name).
		SetTitle(req.Title).
		SetRedirect(req.Redirect).
		SetIcon(req.Icon).
		SetPath(req.Path).
		SetPaths(req.Paths).
		SetMenuType(req.MenuType).
		SetAction(req.Action).
		SetPermission(req.Permission).
		SetIgnoreKeepAlive(req.IgnoreKeepAlive).
		SetHideBreadcrumb(req.HideBreadcrumb).
		SetHideChildrenInMenu(req.HideChildrenInMenu).
		SetComponent(req.Component).
		SetSort(req.Sort).
		SetHideMenu(req.HideMenu).
		SetFrameSrc(req.FrameSrc).
		SetState(sysmenu.State(req.State)).
		SetCheckPermission(req.CheckPermission).
		SetOperation(req.Operation).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetSysMenu 根据Id查询
func (r *sysMenuRepo) GetSysMenu(ctx context.Context, req *v1.SysMenuReq) (*ent.SysMenu, error) {
	return r.data.db.SysMenu.Get(ctx, req.Id)
}

// PageSysMenu 分页查询
func (r *sysMenuRepo) PageSysMenu(ctx context.Context, req *v1.SysMenuPageReq) ([]*ent.SysMenu, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.SysMenu.
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
func (r *sysMenuRepo) genCondition(req *v1.SysMenuReq) []predicate.SysMenu {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysMenu, 0)
	if req.Id > 0 {
		list = append(list, sysmenu.ID(req.Id))
	}
	if req.ParentId > 0 {
		list = append(list, sysmenu.ParentId(req.ParentId))
	}
	if str.IsBlank(req.Name) {
		list = append(list, sysmenu.NameContains(req.Name))
	}
	if str.IsBlank(req.Title) {
		list = append(list, sysmenu.TitleContains(req.Title))
	}
	if str.IsBlank(req.Redirect) {
		list = append(list, sysmenu.RedirectContains(req.Redirect))
	}
	if str.IsBlank(req.Icon) {
		list = append(list, sysmenu.IconContains(req.Icon))
	}
	if str.IsBlank(req.Path) {
		list = append(list, sysmenu.PathContains(req.Path))
	}
	if str.IsBlank(req.Paths) {
		list = append(list, sysmenu.PathsContains(req.Paths))
	}
	if str.IsBlank(req.MenuType) {
		list = append(list, sysmenu.MenuTypeContains(req.MenuType))
	}
	if str.IsBlank(req.Action) {
		list = append(list, sysmenu.ActionContains(req.Action))
	}
	if str.IsBlank(req.Permission) {
		list = append(list, sysmenu.PermissionContains(req.Permission))
	}
	list = append(list, sysmenu.IgnoreKeepAlive(req.IgnoreKeepAlive))
	list = append(list, sysmenu.HideBreadcrumb(req.HideBreadcrumb))
	list = append(list, sysmenu.HideChildrenInMenu(req.HideChildrenInMenu))
	if str.IsBlank(req.Component) {
		list = append(list, sysmenu.ComponentContains(req.Component))
	}
	if req.Sort > 0 {
		list = append(list, sysmenu.Sort(req.Sort))
	}
	list = append(list, sysmenu.HideMenu(req.HideMenu))
	if str.IsBlank(req.FrameSrc) {
		list = append(list, sysmenu.FrameSrcContains(req.FrameSrc))
	}
	state := sysmenu.State(req.State)
	if sysmenu.StateValidator(state) == nil {
		list = append(list, sysmenu.StateEQ(state))
	}
	list = append(list, sysmenu.CheckPermission(req.CheckPermission))
	if str.IsBlank(req.Operation) {
		list = append(list, sysmenu.OperationContains(req.Operation))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysmenu.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysmenu.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysmenu.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysmenu.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysmenu.TenantId(req.TenantId))
	}

	return list
}
