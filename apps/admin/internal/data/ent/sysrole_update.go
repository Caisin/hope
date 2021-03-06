// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysmenu"
	"hope/apps/admin/internal/data/ent/sysrole"
	"hope/apps/admin/internal/data/ent/sysuser"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysRoleUpdate is the builder for updating SysRole entities.
type SysRoleUpdate struct {
	config
	hooks    []Hook
	mutation *SysRoleMutation
}

// Where appends a list predicates to the SysRoleUpdate builder.
func (sru *SysRoleUpdate) Where(ps ...predicate.SysRole) *SysRoleUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// SetRoleName sets the "roleName" field.
func (sru *SysRoleUpdate) SetRoleName(s string) *SysRoleUpdate {
	sru.mutation.SetRoleName(s)
	return sru
}

// SetNillableRoleName sets the "roleName" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableRoleName(s *string) *SysRoleUpdate {
	if s != nil {
		sru.SetRoleName(*s)
	}
	return sru
}

// ClearRoleName clears the value of the "roleName" field.
func (sru *SysRoleUpdate) ClearRoleName() *SysRoleUpdate {
	sru.mutation.ClearRoleName()
	return sru
}

// SetStatus sets the "status" field.
func (sru *SysRoleUpdate) SetStatus(s string) *SysRoleUpdate {
	sru.mutation.SetStatus(s)
	return sru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableStatus(s *string) *SysRoleUpdate {
	if s != nil {
		sru.SetStatus(*s)
	}
	return sru
}

// ClearStatus clears the value of the "status" field.
func (sru *SysRoleUpdate) ClearStatus() *SysRoleUpdate {
	sru.mutation.ClearStatus()
	return sru
}

// SetRoleKey sets the "roleKey" field.
func (sru *SysRoleUpdate) SetRoleKey(s string) *SysRoleUpdate {
	sru.mutation.SetRoleKey(s)
	return sru
}

// SetNillableRoleKey sets the "roleKey" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableRoleKey(s *string) *SysRoleUpdate {
	if s != nil {
		sru.SetRoleKey(*s)
	}
	return sru
}

// ClearRoleKey clears the value of the "roleKey" field.
func (sru *SysRoleUpdate) ClearRoleKey() *SysRoleUpdate {
	sru.mutation.ClearRoleKey()
	return sru
}

// SetRoleSort sets the "roleSort" field.
func (sru *SysRoleUpdate) SetRoleSort(i int32) *SysRoleUpdate {
	sru.mutation.ResetRoleSort()
	sru.mutation.SetRoleSort(i)
	return sru
}

// SetNillableRoleSort sets the "roleSort" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableRoleSort(i *int32) *SysRoleUpdate {
	if i != nil {
		sru.SetRoleSort(*i)
	}
	return sru
}

// AddRoleSort adds i to the "roleSort" field.
func (sru *SysRoleUpdate) AddRoleSort(i int32) *SysRoleUpdate {
	sru.mutation.AddRoleSort(i)
	return sru
}

// ClearRoleSort clears the value of the "roleSort" field.
func (sru *SysRoleUpdate) ClearRoleSort() *SysRoleUpdate {
	sru.mutation.ClearRoleSort()
	return sru
}

// SetFlag sets the "flag" field.
func (sru *SysRoleUpdate) SetFlag(s string) *SysRoleUpdate {
	sru.mutation.SetFlag(s)
	return sru
}

// SetNillableFlag sets the "flag" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableFlag(s *string) *SysRoleUpdate {
	if s != nil {
		sru.SetFlag(*s)
	}
	return sru
}

// ClearFlag clears the value of the "flag" field.
func (sru *SysRoleUpdate) ClearFlag() *SysRoleUpdate {
	sru.mutation.ClearFlag()
	return sru
}

// SetRemark sets the "remark" field.
func (sru *SysRoleUpdate) SetRemark(s string) *SysRoleUpdate {
	sru.mutation.SetRemark(s)
	return sru
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableRemark(s *string) *SysRoleUpdate {
	if s != nil {
		sru.SetRemark(*s)
	}
	return sru
}

// ClearRemark clears the value of the "remark" field.
func (sru *SysRoleUpdate) ClearRemark() *SysRoleUpdate {
	sru.mutation.ClearRemark()
	return sru
}

// SetAdmin sets the "admin" field.
func (sru *SysRoleUpdate) SetAdmin(b bool) *SysRoleUpdate {
	sru.mutation.SetAdmin(b)
	return sru
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableAdmin(b *bool) *SysRoleUpdate {
	if b != nil {
		sru.SetAdmin(*b)
	}
	return sru
}

// ClearAdmin clears the value of the "admin" field.
func (sru *SysRoleUpdate) ClearAdmin() *SysRoleUpdate {
	sru.mutation.ClearAdmin()
	return sru
}

// SetDataScope sets the "dataScope" field.
func (sru *SysRoleUpdate) SetDataScope(s string) *SysRoleUpdate {
	sru.mutation.SetDataScope(s)
	return sru
}

// SetNillableDataScope sets the "dataScope" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableDataScope(s *string) *SysRoleUpdate {
	if s != nil {
		sru.SetDataScope(*s)
	}
	return sru
}

// ClearDataScope clears the value of the "dataScope" field.
func (sru *SysRoleUpdate) ClearDataScope() *SysRoleUpdate {
	sru.mutation.ClearDataScope()
	return sru
}

// SetUpdatedAt sets the "updatedAt" field.
func (sru *SysRoleUpdate) SetUpdatedAt(t time.Time) *SysRoleUpdate {
	sru.mutation.SetUpdatedAt(t)
	return sru
}

// SetCreateBy sets the "createBy" field.
func (sru *SysRoleUpdate) SetCreateBy(i int64) *SysRoleUpdate {
	sru.mutation.ResetCreateBy()
	sru.mutation.SetCreateBy(i)
	return sru
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableCreateBy(i *int64) *SysRoleUpdate {
	if i != nil {
		sru.SetCreateBy(*i)
	}
	return sru
}

// AddCreateBy adds i to the "createBy" field.
func (sru *SysRoleUpdate) AddCreateBy(i int64) *SysRoleUpdate {
	sru.mutation.AddCreateBy(i)
	return sru
}

// SetUpdateBy sets the "updateBy" field.
func (sru *SysRoleUpdate) SetUpdateBy(i int64) *SysRoleUpdate {
	sru.mutation.ResetUpdateBy()
	sru.mutation.SetUpdateBy(i)
	return sru
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableUpdateBy(i *int64) *SysRoleUpdate {
	if i != nil {
		sru.SetUpdateBy(*i)
	}
	return sru
}

// AddUpdateBy adds i to the "updateBy" field.
func (sru *SysRoleUpdate) AddUpdateBy(i int64) *SysRoleUpdate {
	sru.mutation.AddUpdateBy(i)
	return sru
}

// SetTenantId sets the "tenantId" field.
func (sru *SysRoleUpdate) SetTenantId(i int64) *SysRoleUpdate {
	sru.mutation.ResetTenantId()
	sru.mutation.SetTenantId(i)
	return sru
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableTenantId(i *int64) *SysRoleUpdate {
	if i != nil {
		sru.SetTenantId(*i)
	}
	return sru
}

// AddTenantId adds i to the "tenantId" field.
func (sru *SysRoleUpdate) AddTenantId(i int64) *SysRoleUpdate {
	sru.mutation.AddTenantId(i)
	return sru
}

// AddMenuIDs adds the "menus" edge to the SysMenu entity by IDs.
func (sru *SysRoleUpdate) AddMenuIDs(ids ...int64) *SysRoleUpdate {
	sru.mutation.AddMenuIDs(ids...)
	return sru
}

// AddMenus adds the "menus" edges to the SysMenu entity.
func (sru *SysRoleUpdate) AddMenus(s ...*SysMenu) *SysRoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.AddMenuIDs(ids...)
}

// AddUserIDs adds the "users" edge to the SysUser entity by IDs.
func (sru *SysRoleUpdate) AddUserIDs(ids ...int64) *SysRoleUpdate {
	sru.mutation.AddUserIDs(ids...)
	return sru
}

// AddUsers adds the "users" edges to the SysUser entity.
func (sru *SysRoleUpdate) AddUsers(s ...*SysUser) *SysRoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.AddUserIDs(ids...)
}

// Mutation returns the SysRoleMutation object of the builder.
func (sru *SysRoleUpdate) Mutation() *SysRoleMutation {
	return sru.mutation
}

// ClearMenus clears all "menus" edges to the SysMenu entity.
func (sru *SysRoleUpdate) ClearMenus() *SysRoleUpdate {
	sru.mutation.ClearMenus()
	return sru
}

// RemoveMenuIDs removes the "menus" edge to SysMenu entities by IDs.
func (sru *SysRoleUpdate) RemoveMenuIDs(ids ...int64) *SysRoleUpdate {
	sru.mutation.RemoveMenuIDs(ids...)
	return sru
}

// RemoveMenus removes "menus" edges to SysMenu entities.
func (sru *SysRoleUpdate) RemoveMenus(s ...*SysMenu) *SysRoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.RemoveMenuIDs(ids...)
}

// ClearUsers clears all "users" edges to the SysUser entity.
func (sru *SysRoleUpdate) ClearUsers() *SysRoleUpdate {
	sru.mutation.ClearUsers()
	return sru
}

// RemoveUserIDs removes the "users" edge to SysUser entities by IDs.
func (sru *SysRoleUpdate) RemoveUserIDs(ids ...int64) *SysRoleUpdate {
	sru.mutation.RemoveUserIDs(ids...)
	return sru
}

// RemoveUsers removes "users" edges to SysUser entities.
func (sru *SysRoleUpdate) RemoveUsers(s ...*SysUser) *SysRoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *SysRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sru.defaults()
	if len(sru.hooks) == 0 {
		affected, err = sru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sru.mutation = mutation
			affected, err = sru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sru.hooks) - 1; i >= 0; i-- {
			if sru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sru *SysRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *SysRoleUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *SysRoleUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sru *SysRoleUpdate) defaults() {
	if _, ok := sru.mutation.UpdatedAt(); !ok {
		v := sysrole.UpdateDefaultUpdatedAt()
		sru.mutation.SetUpdatedAt(v)
	}
}

func (sru *SysRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrole.Table,
			Columns: sysrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysrole.FieldID,
			},
		},
	}
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.RoleName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRoleName,
		})
	}
	if sru.mutation.RoleNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldRoleName,
		})
	}
	if value, ok := sru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
	}
	if sru.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldStatus,
		})
	}
	if value, ok := sru.mutation.RoleKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRoleKey,
		})
	}
	if sru.mutation.RoleKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldRoleKey,
		})
	}
	if value, ok := sru.mutation.RoleSort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysrole.FieldRoleSort,
		})
	}
	if value, ok := sru.mutation.AddedRoleSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysrole.FieldRoleSort,
		})
	}
	if sru.mutation.RoleSortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: sysrole.FieldRoleSort,
		})
	}
	if value, ok := sru.mutation.Flag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldFlag,
		})
	}
	if sru.mutation.FlagCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldFlag,
		})
	}
	if value, ok := sru.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRemark,
		})
	}
	if sru.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldRemark,
		})
	}
	if value, ok := sru.mutation.Admin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysrole.FieldAdmin,
		})
	}
	if sru.mutation.AdminCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: sysrole.FieldAdmin,
		})
	}
	if value, ok := sru.mutation.DataScope(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldDataScope,
		})
	}
	if sru.mutation.DataScopeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldDataScope,
		})
	}
	if value, ok := sru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldUpdatedAt,
		})
	}
	if value, ok := sru.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldCreateBy,
		})
	}
	if value, ok := sru.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldCreateBy,
		})
	}
	if value, ok := sru.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldUpdateBy,
		})
	}
	if value, ok := sru.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldUpdateBy,
		})
	}
	if value, ok := sru.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldTenantId,
		})
	}
	if value, ok := sru.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldTenantId,
		})
	}
	if sru.mutation.MenusCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.RemovedMenusIDs(); len(nodes) > 0 && !sru.mutation.MenusCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.MenusIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysrole.UsersTable,
			Columns: []string{sysrole.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.RemovedUsersIDs(); len(nodes) > 0 && !sru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysrole.UsersTable,
			Columns: []string{sysrole.UsersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysrole.UsersTable,
			Columns: []string{sysrole.UsersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SysRoleUpdateOne is the builder for updating a single SysRole entity.
type SysRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysRoleMutation
}

// SetRoleName sets the "roleName" field.
func (sruo *SysRoleUpdateOne) SetRoleName(s string) *SysRoleUpdateOne {
	sruo.mutation.SetRoleName(s)
	return sruo
}

// SetNillableRoleName sets the "roleName" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableRoleName(s *string) *SysRoleUpdateOne {
	if s != nil {
		sruo.SetRoleName(*s)
	}
	return sruo
}

// ClearRoleName clears the value of the "roleName" field.
func (sruo *SysRoleUpdateOne) ClearRoleName() *SysRoleUpdateOne {
	sruo.mutation.ClearRoleName()
	return sruo
}

// SetStatus sets the "status" field.
func (sruo *SysRoleUpdateOne) SetStatus(s string) *SysRoleUpdateOne {
	sruo.mutation.SetStatus(s)
	return sruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableStatus(s *string) *SysRoleUpdateOne {
	if s != nil {
		sruo.SetStatus(*s)
	}
	return sruo
}

// ClearStatus clears the value of the "status" field.
func (sruo *SysRoleUpdateOne) ClearStatus() *SysRoleUpdateOne {
	sruo.mutation.ClearStatus()
	return sruo
}

// SetRoleKey sets the "roleKey" field.
func (sruo *SysRoleUpdateOne) SetRoleKey(s string) *SysRoleUpdateOne {
	sruo.mutation.SetRoleKey(s)
	return sruo
}

// SetNillableRoleKey sets the "roleKey" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableRoleKey(s *string) *SysRoleUpdateOne {
	if s != nil {
		sruo.SetRoleKey(*s)
	}
	return sruo
}

// ClearRoleKey clears the value of the "roleKey" field.
func (sruo *SysRoleUpdateOne) ClearRoleKey() *SysRoleUpdateOne {
	sruo.mutation.ClearRoleKey()
	return sruo
}

// SetRoleSort sets the "roleSort" field.
func (sruo *SysRoleUpdateOne) SetRoleSort(i int32) *SysRoleUpdateOne {
	sruo.mutation.ResetRoleSort()
	sruo.mutation.SetRoleSort(i)
	return sruo
}

// SetNillableRoleSort sets the "roleSort" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableRoleSort(i *int32) *SysRoleUpdateOne {
	if i != nil {
		sruo.SetRoleSort(*i)
	}
	return sruo
}

// AddRoleSort adds i to the "roleSort" field.
func (sruo *SysRoleUpdateOne) AddRoleSort(i int32) *SysRoleUpdateOne {
	sruo.mutation.AddRoleSort(i)
	return sruo
}

// ClearRoleSort clears the value of the "roleSort" field.
func (sruo *SysRoleUpdateOne) ClearRoleSort() *SysRoleUpdateOne {
	sruo.mutation.ClearRoleSort()
	return sruo
}

// SetFlag sets the "flag" field.
func (sruo *SysRoleUpdateOne) SetFlag(s string) *SysRoleUpdateOne {
	sruo.mutation.SetFlag(s)
	return sruo
}

// SetNillableFlag sets the "flag" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableFlag(s *string) *SysRoleUpdateOne {
	if s != nil {
		sruo.SetFlag(*s)
	}
	return sruo
}

// ClearFlag clears the value of the "flag" field.
func (sruo *SysRoleUpdateOne) ClearFlag() *SysRoleUpdateOne {
	sruo.mutation.ClearFlag()
	return sruo
}

// SetRemark sets the "remark" field.
func (sruo *SysRoleUpdateOne) SetRemark(s string) *SysRoleUpdateOne {
	sruo.mutation.SetRemark(s)
	return sruo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableRemark(s *string) *SysRoleUpdateOne {
	if s != nil {
		sruo.SetRemark(*s)
	}
	return sruo
}

// ClearRemark clears the value of the "remark" field.
func (sruo *SysRoleUpdateOne) ClearRemark() *SysRoleUpdateOne {
	sruo.mutation.ClearRemark()
	return sruo
}

// SetAdmin sets the "admin" field.
func (sruo *SysRoleUpdateOne) SetAdmin(b bool) *SysRoleUpdateOne {
	sruo.mutation.SetAdmin(b)
	return sruo
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableAdmin(b *bool) *SysRoleUpdateOne {
	if b != nil {
		sruo.SetAdmin(*b)
	}
	return sruo
}

// ClearAdmin clears the value of the "admin" field.
func (sruo *SysRoleUpdateOne) ClearAdmin() *SysRoleUpdateOne {
	sruo.mutation.ClearAdmin()
	return sruo
}

// SetDataScope sets the "dataScope" field.
func (sruo *SysRoleUpdateOne) SetDataScope(s string) *SysRoleUpdateOne {
	sruo.mutation.SetDataScope(s)
	return sruo
}

// SetNillableDataScope sets the "dataScope" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableDataScope(s *string) *SysRoleUpdateOne {
	if s != nil {
		sruo.SetDataScope(*s)
	}
	return sruo
}

// ClearDataScope clears the value of the "dataScope" field.
func (sruo *SysRoleUpdateOne) ClearDataScope() *SysRoleUpdateOne {
	sruo.mutation.ClearDataScope()
	return sruo
}

// SetUpdatedAt sets the "updatedAt" field.
func (sruo *SysRoleUpdateOne) SetUpdatedAt(t time.Time) *SysRoleUpdateOne {
	sruo.mutation.SetUpdatedAt(t)
	return sruo
}

// SetCreateBy sets the "createBy" field.
func (sruo *SysRoleUpdateOne) SetCreateBy(i int64) *SysRoleUpdateOne {
	sruo.mutation.ResetCreateBy()
	sruo.mutation.SetCreateBy(i)
	return sruo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableCreateBy(i *int64) *SysRoleUpdateOne {
	if i != nil {
		sruo.SetCreateBy(*i)
	}
	return sruo
}

// AddCreateBy adds i to the "createBy" field.
func (sruo *SysRoleUpdateOne) AddCreateBy(i int64) *SysRoleUpdateOne {
	sruo.mutation.AddCreateBy(i)
	return sruo
}

// SetUpdateBy sets the "updateBy" field.
func (sruo *SysRoleUpdateOne) SetUpdateBy(i int64) *SysRoleUpdateOne {
	sruo.mutation.ResetUpdateBy()
	sruo.mutation.SetUpdateBy(i)
	return sruo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableUpdateBy(i *int64) *SysRoleUpdateOne {
	if i != nil {
		sruo.SetUpdateBy(*i)
	}
	return sruo
}

// AddUpdateBy adds i to the "updateBy" field.
func (sruo *SysRoleUpdateOne) AddUpdateBy(i int64) *SysRoleUpdateOne {
	sruo.mutation.AddUpdateBy(i)
	return sruo
}

// SetTenantId sets the "tenantId" field.
func (sruo *SysRoleUpdateOne) SetTenantId(i int64) *SysRoleUpdateOne {
	sruo.mutation.ResetTenantId()
	sruo.mutation.SetTenantId(i)
	return sruo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableTenantId(i *int64) *SysRoleUpdateOne {
	if i != nil {
		sruo.SetTenantId(*i)
	}
	return sruo
}

// AddTenantId adds i to the "tenantId" field.
func (sruo *SysRoleUpdateOne) AddTenantId(i int64) *SysRoleUpdateOne {
	sruo.mutation.AddTenantId(i)
	return sruo
}

// AddMenuIDs adds the "menus" edge to the SysMenu entity by IDs.
func (sruo *SysRoleUpdateOne) AddMenuIDs(ids ...int64) *SysRoleUpdateOne {
	sruo.mutation.AddMenuIDs(ids...)
	return sruo
}

// AddMenus adds the "menus" edges to the SysMenu entity.
func (sruo *SysRoleUpdateOne) AddMenus(s ...*SysMenu) *SysRoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.AddMenuIDs(ids...)
}

// AddUserIDs adds the "users" edge to the SysUser entity by IDs.
func (sruo *SysRoleUpdateOne) AddUserIDs(ids ...int64) *SysRoleUpdateOne {
	sruo.mutation.AddUserIDs(ids...)
	return sruo
}

// AddUsers adds the "users" edges to the SysUser entity.
func (sruo *SysRoleUpdateOne) AddUsers(s ...*SysUser) *SysRoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.AddUserIDs(ids...)
}

// Mutation returns the SysRoleMutation object of the builder.
func (sruo *SysRoleUpdateOne) Mutation() *SysRoleMutation {
	return sruo.mutation
}

// ClearMenus clears all "menus" edges to the SysMenu entity.
func (sruo *SysRoleUpdateOne) ClearMenus() *SysRoleUpdateOne {
	sruo.mutation.ClearMenus()
	return sruo
}

// RemoveMenuIDs removes the "menus" edge to SysMenu entities by IDs.
func (sruo *SysRoleUpdateOne) RemoveMenuIDs(ids ...int64) *SysRoleUpdateOne {
	sruo.mutation.RemoveMenuIDs(ids...)
	return sruo
}

// RemoveMenus removes "menus" edges to SysMenu entities.
func (sruo *SysRoleUpdateOne) RemoveMenus(s ...*SysMenu) *SysRoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.RemoveMenuIDs(ids...)
}

// ClearUsers clears all "users" edges to the SysUser entity.
func (sruo *SysRoleUpdateOne) ClearUsers() *SysRoleUpdateOne {
	sruo.mutation.ClearUsers()
	return sruo
}

// RemoveUserIDs removes the "users" edge to SysUser entities by IDs.
func (sruo *SysRoleUpdateOne) RemoveUserIDs(ids ...int64) *SysRoleUpdateOne {
	sruo.mutation.RemoveUserIDs(ids...)
	return sruo
}

// RemoveUsers removes "users" edges to SysUser entities.
func (sruo *SysRoleUpdateOne) RemoveUsers(s ...*SysUser) *SysRoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *SysRoleUpdateOne) Select(field string, fields ...string) *SysRoleUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated SysRole entity.
func (sruo *SysRoleUpdateOne) Save(ctx context.Context) (*SysRole, error) {
	var (
		err  error
		node *SysRole
	)
	sruo.defaults()
	if len(sruo.hooks) == 0 {
		node, err = sruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sruo.mutation = mutation
			node, err = sruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sruo.hooks) - 1; i >= 0; i-- {
			if sruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *SysRoleUpdateOne) SaveX(ctx context.Context) *SysRole {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *SysRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *SysRoleUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sruo *SysRoleUpdateOne) defaults() {
	if _, ok := sruo.mutation.UpdatedAt(); !ok {
		v := sysrole.UpdateDefaultUpdatedAt()
		sruo.mutation.SetUpdatedAt(v)
	}
}

func (sruo *SysRoleUpdateOne) sqlSave(ctx context.Context) (_node *SysRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrole.Table,
			Columns: sysrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysrole.FieldID,
			},
		},
	}
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrole.FieldID)
		for _, f := range fields {
			if !sysrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sruo.mutation.RoleName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRoleName,
		})
	}
	if sruo.mutation.RoleNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldRoleName,
		})
	}
	if value, ok := sruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldStatus,
		})
	}
	if sruo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldStatus,
		})
	}
	if value, ok := sruo.mutation.RoleKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRoleKey,
		})
	}
	if sruo.mutation.RoleKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldRoleKey,
		})
	}
	if value, ok := sruo.mutation.RoleSort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysrole.FieldRoleSort,
		})
	}
	if value, ok := sruo.mutation.AddedRoleSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysrole.FieldRoleSort,
		})
	}
	if sruo.mutation.RoleSortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: sysrole.FieldRoleSort,
		})
	}
	if value, ok := sruo.mutation.Flag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldFlag,
		})
	}
	if sruo.mutation.FlagCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldFlag,
		})
	}
	if value, ok := sruo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldRemark,
		})
	}
	if sruo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldRemark,
		})
	}
	if value, ok := sruo.mutation.Admin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysrole.FieldAdmin,
		})
	}
	if sruo.mutation.AdminCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: sysrole.FieldAdmin,
		})
	}
	if value, ok := sruo.mutation.DataScope(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysrole.FieldDataScope,
		})
	}
	if sruo.mutation.DataScopeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysrole.FieldDataScope,
		})
	}
	if value, ok := sruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysrole.FieldUpdatedAt,
		})
	}
	if value, ok := sruo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldCreateBy,
		})
	}
	if value, ok := sruo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldCreateBy,
		})
	}
	if value, ok := sruo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldUpdateBy,
		})
	}
	if value, ok := sruo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldUpdateBy,
		})
	}
	if value, ok := sruo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldTenantId,
		})
	}
	if value, ok := sruo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysrole.FieldTenantId,
		})
	}
	if sruo.mutation.MenusCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.RemovedMenusIDs(); len(nodes) > 0 && !sruo.mutation.MenusCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.MenusIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysrole.UsersTable,
			Columns: []string{sysrole.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !sruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysrole.UsersTable,
			Columns: []string{sysrole.UsersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysrole.UsersTable,
			Columns: []string{sysrole.UsersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SysRole{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
