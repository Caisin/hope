// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/sysmenu"
	"hope/apps/admin/internal/data/ent/sysrole"
	"hope/apps/admin/internal/data/ent/sysuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysRoleCreate is the builder for creating a SysRole entity.
type SysRoleCreate struct {
	config
	mutation *SysRoleMutation
	hooks    []Hook
}

// SetRoleName sets the "roleName" field.
func (src *SysRoleCreate) SetRoleName(s string) *SysRoleCreate {
	src.mutation.SetRoleName(s)
	return src
}

// SetNillableRoleName sets the "roleName" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableRoleName(s *string) *SysRoleCreate {
	if s != nil {
		src.SetRoleName(*s)
	}
	return src
}

// SetStatus sets the "status" field.
func (src *SysRoleCreate) SetStatus(s string) *SysRoleCreate {
	src.mutation.SetStatus(s)
	return src
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableStatus(s *string) *SysRoleCreate {
	if s != nil {
		src.SetStatus(*s)
	}
	return src
}

// SetRoleKey sets the "roleKey" field.
func (src *SysRoleCreate) SetRoleKey(s string) *SysRoleCreate {
	src.mutation.SetRoleKey(s)
	return src
}

// SetNillableRoleKey sets the "roleKey" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableRoleKey(s *string) *SysRoleCreate {
	if s != nil {
		src.SetRoleKey(*s)
	}
	return src
}

// SetRoleSort sets the "roleSort" field.
func (src *SysRoleCreate) SetRoleSort(i int32) *SysRoleCreate {
	src.mutation.SetRoleSort(i)
	return src
}

// SetNillableRoleSort sets the "roleSort" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableRoleSort(i *int32) *SysRoleCreate {
	if i != nil {
		src.SetRoleSort(*i)
	}
	return src
}

// SetFlag sets the "flag" field.
func (src *SysRoleCreate) SetFlag(s string) *SysRoleCreate {
	src.mutation.SetFlag(s)
	return src
}

// SetNillableFlag sets the "flag" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableFlag(s *string) *SysRoleCreate {
	if s != nil {
		src.SetFlag(*s)
	}
	return src
}

// SetRemark sets the "remark" field.
func (src *SysRoleCreate) SetRemark(s string) *SysRoleCreate {
	src.mutation.SetRemark(s)
	return src
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableRemark(s *string) *SysRoleCreate {
	if s != nil {
		src.SetRemark(*s)
	}
	return src
}

// SetAdmin sets the "admin" field.
func (src *SysRoleCreate) SetAdmin(b bool) *SysRoleCreate {
	src.mutation.SetAdmin(b)
	return src
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableAdmin(b *bool) *SysRoleCreate {
	if b != nil {
		src.SetAdmin(*b)
	}
	return src
}

// SetDataScope sets the "dataScope" field.
func (src *SysRoleCreate) SetDataScope(s string) *SysRoleCreate {
	src.mutation.SetDataScope(s)
	return src
}

// SetNillableDataScope sets the "dataScope" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableDataScope(s *string) *SysRoleCreate {
	if s != nil {
		src.SetDataScope(*s)
	}
	return src
}

// SetSysDept sets the "sysDept" field.
func (src *SysRoleCreate) SetSysDept(s string) *SysRoleCreate {
	src.mutation.SetSysDept(s)
	return src
}

// SetNillableSysDept sets the "sysDept" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableSysDept(s *string) *SysRoleCreate {
	if s != nil {
		src.SetSysDept(*s)
	}
	return src
}

// SetSysMenu sets the "sysMenu" field.
func (src *SysRoleCreate) SetSysMenu(s string) *SysRoleCreate {
	src.mutation.SetSysMenu(s)
	return src
}

// SetNillableSysMenu sets the "sysMenu" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableSysMenu(s *string) *SysRoleCreate {
	if s != nil {
		src.SetSysMenu(*s)
	}
	return src
}

// SetCreatedAt sets the "createdAt" field.
func (src *SysRoleCreate) SetCreatedAt(t time.Time) *SysRoleCreate {
	src.mutation.SetCreatedAt(t)
	return src
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableCreatedAt(t *time.Time) *SysRoleCreate {
	if t != nil {
		src.SetCreatedAt(*t)
	}
	return src
}

// SetUpdatedAt sets the "updatedAt" field.
func (src *SysRoleCreate) SetUpdatedAt(t time.Time) *SysRoleCreate {
	src.mutation.SetUpdatedAt(t)
	return src
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableUpdatedAt(t *time.Time) *SysRoleCreate {
	if t != nil {
		src.SetUpdatedAt(*t)
	}
	return src
}

// SetCreateBy sets the "createBy" field.
func (src *SysRoleCreate) SetCreateBy(i int64) *SysRoleCreate {
	src.mutation.SetCreateBy(i)
	return src
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableCreateBy(i *int64) *SysRoleCreate {
	if i != nil {
		src.SetCreateBy(*i)
	}
	return src
}

// SetUpdateBy sets the "updateBy" field.
func (src *SysRoleCreate) SetUpdateBy(i int64) *SysRoleCreate {
	src.mutation.SetUpdateBy(i)
	return src
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableUpdateBy(i *int64) *SysRoleCreate {
	if i != nil {
		src.SetUpdateBy(*i)
	}
	return src
}

// SetTenantId sets the "tenantId" field.
func (src *SysRoleCreate) SetTenantId(i int64) *SysRoleCreate {
	src.mutation.SetTenantId(i)
	return src
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (src *SysRoleCreate) SetNillableTenantId(i *int64) *SysRoleCreate {
	if i != nil {
		src.SetTenantId(*i)
	}
	return src
}

// AddMenuIDs adds the "menus" edge to the SysMenu entity by IDs.
func (src *SysRoleCreate) AddMenuIDs(ids ...int64) *SysRoleCreate {
	src.mutation.AddMenuIDs(ids...)
	return src
}

// AddMenus adds the "menus" edges to the SysMenu entity.
func (src *SysRoleCreate) AddMenus(s ...*SysMenu) *SysRoleCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return src.AddMenuIDs(ids...)
}

// AddUserIDs adds the "users" edge to the SysUser entity by IDs.
func (src *SysRoleCreate) AddUserIDs(ids ...int64) *SysRoleCreate {
	src.mutation.AddUserIDs(ids...)
	return src
}

// AddUsers adds the "users" edges to the SysUser entity.
func (src *SysRoleCreate) AddUsers(s ...*SysUser) *SysRoleCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return src.AddUserIDs(ids...)
}

// Mutation returns the SysRoleMutation object of the builder.
func (src *SysRoleCreate) Mutation() *SysRoleMutation {
	return src.mutation
}

// Save creates the SysRole in the database.
func (src *SysRoleCreate) Save(ctx context.Context) (*SysRole, error) {
	var (
		err  error
		node *SysRole
	)
	src.defaults()
	if len(src.hooks) == 0 {
		if err = src.check(); err != nil {
			return nil, err
		}
		node, err = src.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = src.check(); err != nil {
				return nil, err
			}
			src.mutation = mutation
			if node, err = src.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(src.hooks) - 1; i >= 0; i-- {
			if src.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = src.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, src.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (src *SysRoleCreate) SaveX(ctx context.Context) *SysRole {
	v, err := src.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (src *SysRoleCreate) Exec(ctx context.Context) error {
	_, err := src.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (src *SysRoleCreate) ExecX(ctx context.Context) {
	if err := src.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (src *SysRoleCreate) defaults() {
	if _, ok := src.mutation.CreatedAt(); !ok {
		v := sysrole.DefaultCreatedAt()
		src.mutation.SetCreatedAt(v)
	}
	if _, ok := src.mutation.UpdatedAt(); !ok {
		v := sysrole.DefaultUpdatedAt()
		src.mutation.SetUpdatedAt(v)
	}
	if _, ok := src.mutation.CreateBy(); !ok {
		v := sysrole.DefaultCreateBy
		src.mutation.SetCreateBy(v)
	}
	if _, ok := src.mutation.UpdateBy(); !ok {
		v := sysrole.DefaultUpdateBy
		src.mutation.SetUpdateBy(v)
	}
	if _, ok := src.mutation.TenantId(); !ok {
		v := sysrole.DefaultTenantId
		src.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (src *SysRoleCreate) check() error {
	if _, ok := src.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := src.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := src.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := src.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := src.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (src *SysRoleCreate) sqlSave(ctx context.Context) (*SysRole, error) {
	_node, _spec := src.createSpec()
	if err := sqlgraph.CreateNode(ctx, src.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (src *SysRoleCreate) createSpec() (*SysRole, *sqlgraph.CreateSpec) {
	var (
		_node = &SysRole{config: src.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysrole.FieldID,
			},
		}
	)
	if value, ok := src.mutation.RoleName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRoleName,
		})
		_node.RoleName = value
	}
	if value, ok := src.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := src.mutation.RoleKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRoleKey,
		})
		_node.RoleKey = value
	}
	if value, ok := src.mutation.RoleSort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysrole.FieldRoleSort,
		})
		_node.RoleSort = value
	}
	if value, ok := src.mutation.Flag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldFlag,
		})
		_node.Flag = value
	}
	if value, ok := src.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := src.mutation.Admin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysrole.FieldAdmin,
		})
		_node.Admin = value
	}
	if value, ok := src.mutation.DataScope(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldDataScope,
		})
		_node.DataScope = value
	}
	if value, ok := src.mutation.SysDept(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldSysDept,
		})
		_node.SysDept = value
	}
	if value, ok := src.mutation.SysMenu(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldSysMenu,
		})
		_node.SysMenu = value
	}
	if value, ok := src.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := src.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := src.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := src.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := src.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := src.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.MenusTable,
			Columns: sysrole.MenusPrimaryKey,
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
	if nodes := src.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.UsersTable,
			Columns: sysrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysuser.FieldID,
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

// SysRoleCreateBulk is the builder for creating many SysRole entities in bulk.
type SysRoleCreateBulk struct {
	config
	builders []*SysRoleCreate
}

// Save creates the SysRole entities in the database.
func (srcb *SysRoleCreateBulk) Save(ctx context.Context) ([]*SysRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(srcb.builders))
	nodes := make([]*SysRole, len(srcb.builders))
	mutators := make([]Mutator, len(srcb.builders))
	for i := range srcb.builders {
		func(i int, root context.Context) {
			builder := srcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, srcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, srcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, srcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (srcb *SysRoleCreateBulk) SaveX(ctx context.Context) []*SysRole {
	v, err := srcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (srcb *SysRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := srcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srcb *SysRoleCreateBulk) ExecX(ctx context.Context) {
	if err := srcb.Exec(ctx); err != nil {
		panic(err)
	}
}
