// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/sysmenu"
	"hope/apps/admin/internal/data/ent/sysrole"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysMenuCreate is the builder for creating a SysMenu entity.
type SysMenuCreate struct {
	config
	mutation *SysMenuMutation
	hooks    []Hook
}

// SetParentId sets the "parentId" field.
func (smc *SysMenuCreate) SetParentId(i int64) *SysMenuCreate {
	smc.mutation.SetParentId(i)
	return smc
}

// SetNillableParentId sets the "parentId" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableParentId(i *int64) *SysMenuCreate {
	if i != nil {
		smc.SetParentId(*i)
	}
	return smc
}

// SetName sets the "name" field.
func (smc *SysMenuCreate) SetName(s string) *SysMenuCreate {
	smc.mutation.SetName(s)
	return smc
}

// SetTitle sets the "title" field.
func (smc *SysMenuCreate) SetTitle(s string) *SysMenuCreate {
	smc.mutation.SetTitle(s)
	return smc
}

// SetRedirect sets the "redirect" field.
func (smc *SysMenuCreate) SetRedirect(s string) *SysMenuCreate {
	smc.mutation.SetRedirect(s)
	return smc
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableRedirect(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetRedirect(*s)
	}
	return smc
}

// SetIcon sets the "icon" field.
func (smc *SysMenuCreate) SetIcon(s string) *SysMenuCreate {
	smc.mutation.SetIcon(s)
	return smc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableIcon(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetIcon(*s)
	}
	return smc
}

// SetPath sets the "path" field.
func (smc *SysMenuCreate) SetPath(s string) *SysMenuCreate {
	smc.mutation.SetPath(s)
	return smc
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillablePath(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetPath(*s)
	}
	return smc
}

// SetPaths sets the "paths" field.
func (smc *SysMenuCreate) SetPaths(s string) *SysMenuCreate {
	smc.mutation.SetPaths(s)
	return smc
}

// SetNillablePaths sets the "paths" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillablePaths(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetPaths(*s)
	}
	return smc
}

// SetMenuType sets the "menuType" field.
func (smc *SysMenuCreate) SetMenuType(s string) *SysMenuCreate {
	smc.mutation.SetMenuType(s)
	return smc
}

// SetNillableMenuType sets the "menuType" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableMenuType(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetMenuType(*s)
	}
	return smc
}

// SetAction sets the "action" field.
func (smc *SysMenuCreate) SetAction(s string) *SysMenuCreate {
	smc.mutation.SetAction(s)
	return smc
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableAction(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetAction(*s)
	}
	return smc
}

// SetPermission sets the "permission" field.
func (smc *SysMenuCreate) SetPermission(s string) *SysMenuCreate {
	smc.mutation.SetPermission(s)
	return smc
}

// SetNillablePermission sets the "permission" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillablePermission(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetPermission(*s)
	}
	return smc
}

// SetIgnoreKeepAlive sets the "ignoreKeepAlive" field.
func (smc *SysMenuCreate) SetIgnoreKeepAlive(b bool) *SysMenuCreate {
	smc.mutation.SetIgnoreKeepAlive(b)
	return smc
}

// SetNillableIgnoreKeepAlive sets the "ignoreKeepAlive" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableIgnoreKeepAlive(b *bool) *SysMenuCreate {
	if b != nil {
		smc.SetIgnoreKeepAlive(*b)
	}
	return smc
}

// SetHideBreadcrumb sets the "hideBreadcrumb" field.
func (smc *SysMenuCreate) SetHideBreadcrumb(b bool) *SysMenuCreate {
	smc.mutation.SetHideBreadcrumb(b)
	return smc
}

// SetNillableHideBreadcrumb sets the "hideBreadcrumb" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableHideBreadcrumb(b *bool) *SysMenuCreate {
	if b != nil {
		smc.SetHideBreadcrumb(*b)
	}
	return smc
}

// SetHideChildrenInMenu sets the "hideChildrenInMenu" field.
func (smc *SysMenuCreate) SetHideChildrenInMenu(b bool) *SysMenuCreate {
	smc.mutation.SetHideChildrenInMenu(b)
	return smc
}

// SetNillableHideChildrenInMenu sets the "hideChildrenInMenu" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableHideChildrenInMenu(b *bool) *SysMenuCreate {
	if b != nil {
		smc.SetHideChildrenInMenu(*b)
	}
	return smc
}

// SetComponent sets the "component" field.
func (smc *SysMenuCreate) SetComponent(s string) *SysMenuCreate {
	smc.mutation.SetComponent(s)
	return smc
}

// SetNillableComponent sets the "component" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableComponent(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetComponent(*s)
	}
	return smc
}

// SetSort sets the "sort" field.
func (smc *SysMenuCreate) SetSort(i int32) *SysMenuCreate {
	smc.mutation.SetSort(i)
	return smc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableSort(i *int32) *SysMenuCreate {
	if i != nil {
		smc.SetSort(*i)
	}
	return smc
}

// SetHideMenu sets the "hideMenu" field.
func (smc *SysMenuCreate) SetHideMenu(b bool) *SysMenuCreate {
	smc.mutation.SetHideMenu(b)
	return smc
}

// SetNillableHideMenu sets the "hideMenu" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableHideMenu(b *bool) *SysMenuCreate {
	if b != nil {
		smc.SetHideMenu(*b)
	}
	return smc
}

// SetFrameSrc sets the "frameSrc" field.
func (smc *SysMenuCreate) SetFrameSrc(s string) *SysMenuCreate {
	smc.mutation.SetFrameSrc(s)
	return smc
}

// SetNillableFrameSrc sets the "frameSrc" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableFrameSrc(s *string) *SysMenuCreate {
	if s != nil {
		smc.SetFrameSrc(*s)
	}
	return smc
}

// SetState sets the "state" field.
func (smc *SysMenuCreate) SetState(s sysmenu.State) *SysMenuCreate {
	smc.mutation.SetState(s)
	return smc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableState(s *sysmenu.State) *SysMenuCreate {
	if s != nil {
		smc.SetState(*s)
	}
	return smc
}

// SetCheckPermission sets the "checkPermission" field.
func (smc *SysMenuCreate) SetCheckPermission(b bool) *SysMenuCreate {
	smc.mutation.SetCheckPermission(b)
	return smc
}

// SetNillableCheckPermission sets the "checkPermission" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableCheckPermission(b *bool) *SysMenuCreate {
	if b != nil {
		smc.SetCheckPermission(*b)
	}
	return smc
}

// SetOperation sets the "operation" field.
func (smc *SysMenuCreate) SetOperation(s string) *SysMenuCreate {
	smc.mutation.SetOperation(s)
	return smc
}

// SetCreatedAt sets the "createdAt" field.
func (smc *SysMenuCreate) SetCreatedAt(t time.Time) *SysMenuCreate {
	smc.mutation.SetCreatedAt(t)
	return smc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableCreatedAt(t *time.Time) *SysMenuCreate {
	if t != nil {
		smc.SetCreatedAt(*t)
	}
	return smc
}

// SetUpdatedAt sets the "updatedAt" field.
func (smc *SysMenuCreate) SetUpdatedAt(t time.Time) *SysMenuCreate {
	smc.mutation.SetUpdatedAt(t)
	return smc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableUpdatedAt(t *time.Time) *SysMenuCreate {
	if t != nil {
		smc.SetUpdatedAt(*t)
	}
	return smc
}

// SetCreateBy sets the "createBy" field.
func (smc *SysMenuCreate) SetCreateBy(i int64) *SysMenuCreate {
	smc.mutation.SetCreateBy(i)
	return smc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableCreateBy(i *int64) *SysMenuCreate {
	if i != nil {
		smc.SetCreateBy(*i)
	}
	return smc
}

// SetUpdateBy sets the "updateBy" field.
func (smc *SysMenuCreate) SetUpdateBy(i int64) *SysMenuCreate {
	smc.mutation.SetUpdateBy(i)
	return smc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableUpdateBy(i *int64) *SysMenuCreate {
	if i != nil {
		smc.SetUpdateBy(*i)
	}
	return smc
}

// SetTenantId sets the "tenantId" field.
func (smc *SysMenuCreate) SetTenantId(i int64) *SysMenuCreate {
	smc.mutation.SetTenantId(i)
	return smc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (smc *SysMenuCreate) SetNillableTenantId(i *int64) *SysMenuCreate {
	if i != nil {
		smc.SetTenantId(*i)
	}
	return smc
}

// AddRoleIDs adds the "role" edge to the SysRole entity by IDs.
func (smc *SysMenuCreate) AddRoleIDs(ids ...int64) *SysMenuCreate {
	smc.mutation.AddRoleIDs(ids...)
	return smc
}

// AddRole adds the "role" edges to the SysRole entity.
func (smc *SysMenuCreate) AddRole(s ...*SysRole) *SysMenuCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return smc.AddRoleIDs(ids...)
}

// SetParentID sets the "parent" edge to the SysMenu entity by ID.
func (smc *SysMenuCreate) SetParentID(id int64) *SysMenuCreate {
	smc.mutation.SetParentID(id)
	return smc
}

// SetParent sets the "parent" edge to the SysMenu entity.
func (smc *SysMenuCreate) SetParent(s *SysMenu) *SysMenuCreate {
	return smc.SetParentID(s.ID)
}

// AddChildIDs adds the "children" edge to the SysMenu entity by IDs.
func (smc *SysMenuCreate) AddChildIDs(ids ...int64) *SysMenuCreate {
	smc.mutation.AddChildIDs(ids...)
	return smc
}

// AddChildren adds the "children" edges to the SysMenu entity.
func (smc *SysMenuCreate) AddChildren(s ...*SysMenu) *SysMenuCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return smc.AddChildIDs(ids...)
}

// Mutation returns the SysMenuMutation object of the builder.
func (smc *SysMenuCreate) Mutation() *SysMenuMutation {
	return smc.mutation
}

// Save creates the SysMenu in the database.
func (smc *SysMenuCreate) Save(ctx context.Context) (*SysMenu, error) {
	var (
		err  error
		node *SysMenu
	)
	smc.defaults()
	if len(smc.hooks) == 0 {
		if err = smc.check(); err != nil {
			return nil, err
		}
		node, err = smc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smc.check(); err != nil {
				return nil, err
			}
			smc.mutation = mutation
			if node, err = smc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(smc.hooks) - 1; i >= 0; i-- {
			if smc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (smc *SysMenuCreate) SaveX(ctx context.Context) *SysMenu {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smc *SysMenuCreate) Exec(ctx context.Context) error {
	_, err := smc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smc *SysMenuCreate) ExecX(ctx context.Context) {
	if err := smc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smc *SysMenuCreate) defaults() {
	if _, ok := smc.mutation.ParentId(); !ok {
		v := sysmenu.DefaultParentId
		smc.mutation.SetParentId(v)
	}
	if _, ok := smc.mutation.State(); !ok {
		v := sysmenu.DefaultState
		smc.mutation.SetState(v)
	}
	if _, ok := smc.mutation.CheckPermission(); !ok {
		v := sysmenu.DefaultCheckPermission
		smc.mutation.SetCheckPermission(v)
	}
	if _, ok := smc.mutation.CreatedAt(); !ok {
		v := sysmenu.DefaultCreatedAt()
		smc.mutation.SetCreatedAt(v)
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		v := sysmenu.DefaultUpdatedAt()
		smc.mutation.SetUpdatedAt(v)
	}
	if _, ok := smc.mutation.CreateBy(); !ok {
		v := sysmenu.DefaultCreateBy
		smc.mutation.SetCreateBy(v)
	}
	if _, ok := smc.mutation.UpdateBy(); !ok {
		v := sysmenu.DefaultUpdateBy
		smc.mutation.SetUpdateBy(v)
	}
	if _, ok := smc.mutation.TenantId(); !ok {
		v := sysmenu.DefaultTenantId
		smc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *SysMenuCreate) check() error {
	if _, ok := smc.mutation.ParentId(); !ok {
		return &ValidationError{Name: "parentId", err: errors.New(`ent: missing required field "parentId"`)}
	}
	if _, ok := smc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := smc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if _, ok := smc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "state"`)}
	}
	if v, ok := smc.mutation.State(); ok {
		if err := sysmenu.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "state": %w`, err)}
		}
	}
	if _, ok := smc.mutation.CheckPermission(); !ok {
		return &ValidationError{Name: "checkPermission", err: errors.New(`ent: missing required field "checkPermission"`)}
	}
	if _, ok := smc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "operation"`)}
	}
	if _, ok := smc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := smc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := smc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := smc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	if _, ok := smc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent", err: errors.New("ent: missing required edge \"parent\"")}
	}
	return nil
}

func (smc *SysMenuCreate) sqlSave(ctx context.Context) (*SysMenu, error) {
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (smc *SysMenuCreate) createSpec() (*SysMenu, *sqlgraph.CreateSpec) {
	var (
		_node = &SysMenu{config: smc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysmenu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysmenu.FieldID,
			},
		}
	)
	if value, ok := smc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldName,
		})
		_node.Name = value
	}
	if value, ok := smc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := smc.mutation.Redirect(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldRedirect,
		})
		_node.Redirect = value
	}
	if value, ok := smc.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldIcon,
		})
		_node.Icon = value
	}
	if value, ok := smc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldPath,
		})
		_node.Path = value
	}
	if value, ok := smc.mutation.Paths(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldPaths,
		})
		_node.Paths = value
	}
	if value, ok := smc.mutation.MenuType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldMenuType,
		})
		_node.MenuType = value
	}
	if value, ok := smc.mutation.Action(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldAction,
		})
		_node.Action = value
	}
	if value, ok := smc.mutation.Permission(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldPermission,
		})
		_node.Permission = value
	}
	if value, ok := smc.mutation.IgnoreKeepAlive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldIgnoreKeepAlive,
		})
		_node.IgnoreKeepAlive = value
	}
	if value, ok := smc.mutation.HideBreadcrumb(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldHideBreadcrumb,
		})
		_node.HideBreadcrumb = value
	}
	if value, ok := smc.mutation.HideChildrenInMenu(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldHideChildrenInMenu,
		})
		_node.HideChildrenInMenu = value
	}
	if value, ok := smc.mutation.Component(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldComponent,
		})
		_node.Component = value
	}
	if value, ok := smc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenu.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := smc.mutation.HideMenu(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldHideMenu,
		})
		_node.HideMenu = value
	}
	if value, ok := smc.mutation.FrameSrc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldFrameSrc,
		})
		_node.FrameSrc = value
	}
	if value, ok := smc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: sysmenu.FieldState,
		})
		_node.State = value
	}
	if value, ok := smc.mutation.CheckPermission(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldCheckPermission,
		})
		_node.CheckPermission = value
	}
	if value, ok := smc.mutation.Operation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldOperation,
		})
		_node.Operation = value
	}
	if value, ok := smc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := smc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := smc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysmenu.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := smc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysmenu.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := smc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysmenu.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := smc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysmenu.RoleTable,
			Columns: sysmenu.RolePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := smc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysmenu.ParentTable,
			Columns: []string{sysmenu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := smc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysmenu.ChildrenTable,
			Columns: []string{sysmenu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SysMenuCreateBulk is the builder for creating many SysMenu entities in bulk.
type SysMenuCreateBulk struct {
	config
	builders []*SysMenuCreate
}

// Save creates the SysMenu entities in the database.
func (smcb *SysMenuCreateBulk) Save(ctx context.Context) ([]*SysMenu, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*SysMenu, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysMenuMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *SysMenuCreateBulk) SaveX(ctx context.Context) []*SysMenu {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smcb *SysMenuCreateBulk) Exec(ctx context.Context) error {
	_, err := smcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smcb *SysMenuCreateBulk) ExecX(ctx context.Context) {
	if err := smcb.Exec(ctx); err != nil {
		panic(err)
	}
}
