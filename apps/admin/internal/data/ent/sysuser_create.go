// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/sysdept"
	"hope/apps/admin/internal/data/ent/sysloginlog"
	"hope/apps/admin/internal/data/ent/sysoperalog"
	"hope/apps/admin/internal/data/ent/syspost"
	"hope/apps/admin/internal/data/ent/sysuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysUserCreate is the builder for creating a SysUser entity.
type SysUserCreate struct {
	config
	mutation *SysUserMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (suc *SysUserCreate) SetUsername(s string) *SysUserCreate {
	suc.mutation.SetUsername(s)
	return suc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUsername(s *string) *SysUserCreate {
	if s != nil {
		suc.SetUsername(*s)
	}
	return suc
}

// SetNickName sets the "nickName" field.
func (suc *SysUserCreate) SetNickName(s string) *SysUserCreate {
	suc.mutation.SetNickName(s)
	return suc
}

// SetNillableNickName sets the "nickName" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableNickName(s *string) *SysUserCreate {
	if s != nil {
		suc.SetNickName(*s)
	}
	return suc
}

// SetPhone sets the "phone" field.
func (suc *SysUserCreate) SetPhone(s string) *SysUserCreate {
	suc.mutation.SetPhone(s)
	return suc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (suc *SysUserCreate) SetNillablePhone(s *string) *SysUserCreate {
	if s != nil {
		suc.SetPhone(*s)
	}
	return suc
}

// SetRoleId sets the "roleId" field.
func (suc *SysUserCreate) SetRoleId(i int32) *SysUserCreate {
	suc.mutation.SetRoleId(i)
	return suc
}

// SetNillableRoleId sets the "roleId" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableRoleId(i *int32) *SysUserCreate {
	if i != nil {
		suc.SetRoleId(*i)
	}
	return suc
}

// SetAvatar sets the "avatar" field.
func (suc *SysUserCreate) SetAvatar(s string) *SysUserCreate {
	suc.mutation.SetAvatar(s)
	return suc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableAvatar(s *string) *SysUserCreate {
	if s != nil {
		suc.SetAvatar(*s)
	}
	return suc
}

// SetSex sets the "sex" field.
func (suc *SysUserCreate) SetSex(i int32) *SysUserCreate {
	suc.mutation.SetSex(i)
	return suc
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableSex(i *int32) *SysUserCreate {
	if i != nil {
		suc.SetSex(*i)
	}
	return suc
}

// SetEmail sets the "email" field.
func (suc *SysUserCreate) SetEmail(s string) *SysUserCreate {
	suc.mutation.SetEmail(s)
	return suc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableEmail(s *string) *SysUserCreate {
	if s != nil {
		suc.SetEmail(*s)
	}
	return suc
}

// SetRemark sets the "remark" field.
func (suc *SysUserCreate) SetRemark(s string) *SysUserCreate {
	suc.mutation.SetRemark(s)
	return suc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableRemark(s *string) *SysUserCreate {
	if s != nil {
		suc.SetRemark(*s)
	}
	return suc
}

// SetStatus sets the "status" field.
func (suc *SysUserCreate) SetStatus(s string) *SysUserCreate {
	suc.mutation.SetStatus(s)
	return suc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableStatus(s *string) *SysUserCreate {
	if s != nil {
		suc.SetStatus(*s)
	}
	return suc
}

// SetExtInfo sets the "extInfo" field.
func (suc *SysUserCreate) SetExtInfo(s string) *SysUserCreate {
	suc.mutation.SetExtInfo(s)
	return suc
}

// SetNillableExtInfo sets the "extInfo" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableExtInfo(s *string) *SysUserCreate {
	if s != nil {
		suc.SetExtInfo(*s)
	}
	return suc
}

// SetCreatedAt sets the "createdAt" field.
func (suc *SysUserCreate) SetCreatedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetCreatedAt(t)
	return suc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableCreatedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetCreatedAt(*t)
	}
	return suc
}

// SetUpdatedAt sets the "updatedAt" field.
func (suc *SysUserCreate) SetUpdatedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetUpdatedAt(t)
	return suc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUpdatedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetUpdatedAt(*t)
	}
	return suc
}

// SetCreateBy sets the "createBy" field.
func (suc *SysUserCreate) SetCreateBy(i int64) *SysUserCreate {
	suc.mutation.SetCreateBy(i)
	return suc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableCreateBy(i *int64) *SysUserCreate {
	if i != nil {
		suc.SetCreateBy(*i)
	}
	return suc
}

// SetUpdateBy sets the "updateBy" field.
func (suc *SysUserCreate) SetUpdateBy(i int64) *SysUserCreate {
	suc.mutation.SetUpdateBy(i)
	return suc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUpdateBy(i *int64) *SysUserCreate {
	if i != nil {
		suc.SetUpdateBy(*i)
	}
	return suc
}

// SetTenantId sets the "tenantId" field.
func (suc *SysUserCreate) SetTenantId(i int64) *SysUserCreate {
	suc.mutation.SetTenantId(i)
	return suc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableTenantId(i *int64) *SysUserCreate {
	if i != nil {
		suc.SetTenantId(*i)
	}
	return suc
}

// SetDeptID sets the "dept" edge to the SysDept entity by ID.
func (suc *SysUserCreate) SetDeptID(id int64) *SysUserCreate {
	suc.mutation.SetDeptID(id)
	return suc
}

// SetNillableDeptID sets the "dept" edge to the SysDept entity by ID if the given value is not nil.
func (suc *SysUserCreate) SetNillableDeptID(id *int64) *SysUserCreate {
	if id != nil {
		suc = suc.SetDeptID(*id)
	}
	return suc
}

// SetDept sets the "dept" edge to the SysDept entity.
func (suc *SysUserCreate) SetDept(s *SysDept) *SysUserCreate {
	return suc.SetDeptID(s.ID)
}

// SetPostID sets the "post" edge to the SysPost entity by ID.
func (suc *SysUserCreate) SetPostID(id int64) *SysUserCreate {
	suc.mutation.SetPostID(id)
	return suc
}

// SetNillablePostID sets the "post" edge to the SysPost entity by ID if the given value is not nil.
func (suc *SysUserCreate) SetNillablePostID(id *int64) *SysUserCreate {
	if id != nil {
		suc = suc.SetPostID(*id)
	}
	return suc
}

// SetPost sets the "post" edge to the SysPost entity.
func (suc *SysUserCreate) SetPost(s *SysPost) *SysUserCreate {
	return suc.SetPostID(s.ID)
}

// AddLoginLogIDs adds the "loginLogs" edge to the SysLoginLog entity by IDs.
func (suc *SysUserCreate) AddLoginLogIDs(ids ...int64) *SysUserCreate {
	suc.mutation.AddLoginLogIDs(ids...)
	return suc
}

// AddLoginLogs adds the "loginLogs" edges to the SysLoginLog entity.
func (suc *SysUserCreate) AddLoginLogs(s ...*SysLoginLog) *SysUserCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suc.AddLoginLogIDs(ids...)
}

// AddOperaLogIDs adds the "operaLogs" edge to the SysOperaLog entity by IDs.
func (suc *SysUserCreate) AddOperaLogIDs(ids ...int64) *SysUserCreate {
	suc.mutation.AddOperaLogIDs(ids...)
	return suc
}

// AddOperaLogs adds the "operaLogs" edges to the SysOperaLog entity.
func (suc *SysUserCreate) AddOperaLogs(s ...*SysOperaLog) *SysUserCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suc.AddOperaLogIDs(ids...)
}

// Mutation returns the SysUserMutation object of the builder.
func (suc *SysUserCreate) Mutation() *SysUserMutation {
	return suc.mutation
}

// Save creates the SysUser in the database.
func (suc *SysUserCreate) Save(ctx context.Context) (*SysUser, error) {
	var (
		err  error
		node *SysUser
	)
	suc.defaults()
	if len(suc.hooks) == 0 {
		if err = suc.check(); err != nil {
			return nil, err
		}
		node, err = suc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suc.check(); err != nil {
				return nil, err
			}
			suc.mutation = mutation
			if node, err = suc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(suc.hooks) - 1; i >= 0; i-- {
			if suc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (suc *SysUserCreate) SaveX(ctx context.Context) *SysUser {
	v, err := suc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (suc *SysUserCreate) Exec(ctx context.Context) error {
	_, err := suc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suc *SysUserCreate) ExecX(ctx context.Context) {
	if err := suc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suc *SysUserCreate) defaults() {
	if _, ok := suc.mutation.CreatedAt(); !ok {
		v := sysuser.DefaultCreatedAt()
		suc.mutation.SetCreatedAt(v)
	}
	if _, ok := suc.mutation.UpdatedAt(); !ok {
		v := sysuser.DefaultUpdatedAt()
		suc.mutation.SetUpdatedAt(v)
	}
	if _, ok := suc.mutation.CreateBy(); !ok {
		v := sysuser.DefaultCreateBy
		suc.mutation.SetCreateBy(v)
	}
	if _, ok := suc.mutation.UpdateBy(); !ok {
		v := sysuser.DefaultUpdateBy
		suc.mutation.SetUpdateBy(v)
	}
	if _, ok := suc.mutation.TenantId(); !ok {
		v := sysuser.DefaultTenantId
		suc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suc *SysUserCreate) check() error {
	if _, ok := suc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := suc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := suc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := suc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := suc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (suc *SysUserCreate) sqlSave(ctx context.Context) (*SysUser, error) {
	_node, _spec := suc.createSpec()
	if err := sqlgraph.CreateNode(ctx, suc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (suc *SysUserCreate) createSpec() (*SysUser, *sqlgraph.CreateSpec) {
	var (
		_node = &SysUser{config: suc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysuser.FieldID,
			},
		}
	)
	if value, ok := suc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := suc.mutation.NickName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldNickName,
		})
		_node.NickName = value
	}
	if value, ok := suc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := suc.mutation.RoleId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysuser.FieldRoleId,
		})
		_node.RoleId = value
	}
	if value, ok := suc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldAvatar,
		})
		_node.Avatar = value
	}
	if value, ok := suc.mutation.Sex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysuser.FieldSex,
		})
		_node.Sex = value
	}
	if value, ok := suc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := suc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := suc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := suc.mutation.ExtInfo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysuser.FieldExtInfo,
		})
		_node.ExtInfo = value
	}
	if value, ok := suc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := suc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysuser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := suc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysuser.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := suc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysuser.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := suc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysuser.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := suc.mutation.DeptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysuser.DeptTable,
			Columns: []string{sysuser.DeptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysdept.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sys_dept_users = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := suc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysuser.PostTable,
			Columns: []string{sysuser.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: syspost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sys_post_users = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := suc.mutation.LoginLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysuser.LoginLogsTable,
			Columns: []string{sysuser.LoginLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysloginlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := suc.mutation.OperaLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysuser.OperaLogsTable,
			Columns: []string{sysuser.OperaLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysoperalog.FieldID,
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

// SysUserCreateBulk is the builder for creating many SysUser entities in bulk.
type SysUserCreateBulk struct {
	config
	builders []*SysUserCreate
}

// Save creates the SysUser entities in the database.
func (sucb *SysUserCreateBulk) Save(ctx context.Context) ([]*SysUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sucb.builders))
	nodes := make([]*SysUser, len(sucb.builders))
	mutators := make([]Mutator, len(sucb.builders))
	for i := range sucb.builders {
		func(i int, root context.Context) {
			builder := sucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysUserMutation)
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
					_, err = mutators[i+1].Mutate(root, sucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sucb *SysUserCreateBulk) SaveX(ctx context.Context) []*SysUser {
	v, err := sucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sucb *SysUserCreateBulk) Exec(ctx context.Context) error {
	_, err := sucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sucb *SysUserCreateBulk) ExecX(ctx context.Context) {
	if err := sucb.Exec(ctx); err != nil {
		panic(err)
	}
}
