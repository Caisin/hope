package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysmenu/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysmenu"
	"hope/pkg/util/str"
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
	now := time.Now()
	return r.data.db.SysMenu.Create().
		SetMenuName(req.MenuName).
		SetTitle(req.Title).
		SetIcon(req.Icon).
		SetPath(req.Path).
		SetPaths(req.Paths).
		SetMenuType(req.MenuType).
		SetAction(req.Action).
		SetPermission(req.Permission).
		SetNoCache(req.NoCache).
		SetBreadcrumb(req.Breadcrumb).
		SetComponent(req.Component).
		SetSort(req.Sort).
		SetVisible(req.Visible).
		SetIsFrame(req.IsFrame).
		SetSysApi(req.SysApi).
		SetCreatedAt(now).
		SetUpdatedAt(now).
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
	return r.data.db.SysMenu.UpdateOne(convert.SysMenuUpdateReq2Data(req)).Save(ctx)
}

// GetSysMenu 根据Id查询
func (r *sysMenuRepo) GetSysMenu(ctx context.Context, req *v1.SysMenuReq) (*ent.SysMenu, error) {
	return r.data.db.SysMenu.Get(ctx, req.Id)
}

// PageSysMenu 分页查询
func (r *sysMenuRepo) PageSysMenu(ctx context.Context, req *v1.SysMenuPageReq) ([]*ent.SysMenu, error) {
	pagin := req.Pagin
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
func (r *sysMenuRepo) genCondition(req *v1.SysMenuReq) []predicate.SysMenu {
	list := make([]predicate.SysMenu, 0)
	if req.Id > 0 {
		list = append(list, sysmenu.ID(req.Id))
	}
	if str.IsBlank(req.MenuName) {
		list = append(list, sysmenu.MenuNameContains(req.MenuName))
	}
	if str.IsBlank(req.Title) {
		list = append(list, sysmenu.TitleContains(req.Title))
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
	if str.IsBlank(req.Breadcrumb) {
		list = append(list, sysmenu.BreadcrumbContains(req.Breadcrumb))
	}
	if str.IsBlank(req.Component) {
		list = append(list, sysmenu.ComponentContains(req.Component))
	}
	if req.Sort > 0 {
		list = append(list, sysmenu.Sort(req.Sort))
	}
	if str.IsBlank(req.SysApi) {
		list = append(list, sysmenu.SysApiContains(req.SysApi))
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
