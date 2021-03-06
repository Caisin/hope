// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/syspost"
	"hope/apps/admin/internal/data/ent/sysuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysPostCreate is the builder for creating a SysPost entity.
type SysPostCreate struct {
	config
	mutation *SysPostMutation
	hooks    []Hook
}

// SetPostName sets the "postName" field.
func (spc *SysPostCreate) SetPostName(s string) *SysPostCreate {
	spc.mutation.SetPostName(s)
	return spc
}

// SetNillablePostName sets the "postName" field if the given value is not nil.
func (spc *SysPostCreate) SetNillablePostName(s *string) *SysPostCreate {
	if s != nil {
		spc.SetPostName(*s)
	}
	return spc
}

// SetPostCode sets the "postCode" field.
func (spc *SysPostCreate) SetPostCode(s string) *SysPostCreate {
	spc.mutation.SetPostCode(s)
	return spc
}

// SetNillablePostCode sets the "postCode" field if the given value is not nil.
func (spc *SysPostCreate) SetNillablePostCode(s *string) *SysPostCreate {
	if s != nil {
		spc.SetPostCode(*s)
	}
	return spc
}

// SetSort sets the "sort" field.
func (spc *SysPostCreate) SetSort(i int32) *SysPostCreate {
	spc.mutation.SetSort(i)
	return spc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableSort(i *int32) *SysPostCreate {
	if i != nil {
		spc.SetSort(*i)
	}
	return spc
}

// SetStatus sets the "status" field.
func (spc *SysPostCreate) SetStatus(i int32) *SysPostCreate {
	spc.mutation.SetStatus(i)
	return spc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableStatus(i *int32) *SysPostCreate {
	if i != nil {
		spc.SetStatus(*i)
	}
	return spc
}

// SetRemark sets the "remark" field.
func (spc *SysPostCreate) SetRemark(s string) *SysPostCreate {
	spc.mutation.SetRemark(s)
	return spc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableRemark(s *string) *SysPostCreate {
	if s != nil {
		spc.SetRemark(*s)
	}
	return spc
}

// SetCreatedAt sets the "createdAt" field.
func (spc *SysPostCreate) SetCreatedAt(t time.Time) *SysPostCreate {
	spc.mutation.SetCreatedAt(t)
	return spc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableCreatedAt(t *time.Time) *SysPostCreate {
	if t != nil {
		spc.SetCreatedAt(*t)
	}
	return spc
}

// SetUpdatedAt sets the "updatedAt" field.
func (spc *SysPostCreate) SetUpdatedAt(t time.Time) *SysPostCreate {
	spc.mutation.SetUpdatedAt(t)
	return spc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableUpdatedAt(t *time.Time) *SysPostCreate {
	if t != nil {
		spc.SetUpdatedAt(*t)
	}
	return spc
}

// SetCreateBy sets the "createBy" field.
func (spc *SysPostCreate) SetCreateBy(i int64) *SysPostCreate {
	spc.mutation.SetCreateBy(i)
	return spc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableCreateBy(i *int64) *SysPostCreate {
	if i != nil {
		spc.SetCreateBy(*i)
	}
	return spc
}

// SetUpdateBy sets the "updateBy" field.
func (spc *SysPostCreate) SetUpdateBy(i int64) *SysPostCreate {
	spc.mutation.SetUpdateBy(i)
	return spc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableUpdateBy(i *int64) *SysPostCreate {
	if i != nil {
		spc.SetUpdateBy(*i)
	}
	return spc
}

// SetTenantId sets the "tenantId" field.
func (spc *SysPostCreate) SetTenantId(i int64) *SysPostCreate {
	spc.mutation.SetTenantId(i)
	return spc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (spc *SysPostCreate) SetNillableTenantId(i *int64) *SysPostCreate {
	if i != nil {
		spc.SetTenantId(*i)
	}
	return spc
}

// AddUserIDs adds the "users" edge to the SysUser entity by IDs.
func (spc *SysPostCreate) AddUserIDs(ids ...int64) *SysPostCreate {
	spc.mutation.AddUserIDs(ids...)
	return spc
}

// AddUsers adds the "users" edges to the SysUser entity.
func (spc *SysPostCreate) AddUsers(s ...*SysUser) *SysPostCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return spc.AddUserIDs(ids...)
}

// Mutation returns the SysPostMutation object of the builder.
func (spc *SysPostCreate) Mutation() *SysPostMutation {
	return spc.mutation
}

// Save creates the SysPost in the database.
func (spc *SysPostCreate) Save(ctx context.Context) (*SysPost, error) {
	var (
		err  error
		node *SysPost
	)
	spc.defaults()
	if len(spc.hooks) == 0 {
		if err = spc.check(); err != nil {
			return nil, err
		}
		node, err = spc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysPostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = spc.check(); err != nil {
				return nil, err
			}
			spc.mutation = mutation
			if node, err = spc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(spc.hooks) - 1; i >= 0; i-- {
			if spc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (spc *SysPostCreate) SaveX(ctx context.Context) *SysPost {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *SysPostCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *SysPostCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spc *SysPostCreate) defaults() {
	if _, ok := spc.mutation.CreatedAt(); !ok {
		v := syspost.DefaultCreatedAt()
		spc.mutation.SetCreatedAt(v)
	}
	if _, ok := spc.mutation.UpdatedAt(); !ok {
		v := syspost.DefaultUpdatedAt()
		spc.mutation.SetUpdatedAt(v)
	}
	if _, ok := spc.mutation.CreateBy(); !ok {
		v := syspost.DefaultCreateBy
		spc.mutation.SetCreateBy(v)
	}
	if _, ok := spc.mutation.UpdateBy(); !ok {
		v := syspost.DefaultUpdateBy
		spc.mutation.SetUpdateBy(v)
	}
	if _, ok := spc.mutation.TenantId(); !ok {
		v := syspost.DefaultTenantId
		spc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *SysPostCreate) check() error {
	if _, ok := spc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "SysPost.createdAt"`)}
	}
	if _, ok := spc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "SysPost.updatedAt"`)}
	}
	if _, ok := spc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "SysPost.createBy"`)}
	}
	if _, ok := spc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "SysPost.updateBy"`)}
	}
	if _, ok := spc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "SysPost.tenantId"`)}
	}
	return nil
}

func (spc *SysPostCreate) sqlSave(ctx context.Context) (*SysPost, error) {
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (spc *SysPostCreate) createSpec() (*SysPost, *sqlgraph.CreateSpec) {
	var (
		_node = &SysPost{config: spc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: syspost.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: syspost.FieldID,
			},
		}
	)
	if value, ok := spc.mutation.PostName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syspost.FieldPostName,
		})
		_node.PostName = value
	}
	if value, ok := spc.mutation.PostCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syspost.FieldPostCode,
		})
		_node.PostCode = value
	}
	if value, ok := spc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syspost.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := spc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syspost.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := spc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syspost.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := spc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: syspost.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := spc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: syspost.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := spc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := spc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := spc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := spc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   syspost.UsersTable,
			Columns: []string{syspost.UsersColumn},
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

// SysPostCreateBulk is the builder for creating many SysPost entities in bulk.
type SysPostCreateBulk struct {
	config
	builders []*SysPostCreate
}

// Save creates the SysPost entities in the database.
func (spcb *SysPostCreateBulk) Save(ctx context.Context) ([]*SysPost, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*SysPost, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysPostMutation)
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
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *SysPostCreateBulk) SaveX(ctx context.Context) []*SysPost {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *SysPostCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *SysPostCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}
