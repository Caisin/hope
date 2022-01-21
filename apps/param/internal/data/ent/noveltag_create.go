// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/param/internal/data/ent/noveltag"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelTagCreate is the builder for creating a NovelTag entity.
type NovelTagCreate struct {
	config
	mutation *NovelTagMutation
	hooks    []Hook
}

// SetTagId sets the "tagId" field.
func (ntc *NovelTagCreate) SetTagId(i int64) *NovelTagCreate {
	ntc.mutation.SetTagId(i)
	return ntc
}

// SetTagName sets the "tagName" field.
func (ntc *NovelTagCreate) SetTagName(s string) *NovelTagCreate {
	ntc.mutation.SetTagName(s)
	return ntc
}

// SetNillableTagName sets the "tagName" field if the given value is not nil.
func (ntc *NovelTagCreate) SetNillableTagName(s *string) *NovelTagCreate {
	if s != nil {
		ntc.SetTagName(*s)
	}
	return ntc
}

// SetRemark sets the "remark" field.
func (ntc *NovelTagCreate) SetRemark(s string) *NovelTagCreate {
	ntc.mutation.SetRemark(s)
	return ntc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ntc *NovelTagCreate) SetNillableRemark(s *string) *NovelTagCreate {
	if s != nil {
		ntc.SetRemark(*s)
	}
	return ntc
}

// SetCreatedAt sets the "createdAt" field.
func (ntc *NovelTagCreate) SetCreatedAt(t time.Time) *NovelTagCreate {
	ntc.mutation.SetCreatedAt(t)
	return ntc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (ntc *NovelTagCreate) SetNillableCreatedAt(t *time.Time) *NovelTagCreate {
	if t != nil {
		ntc.SetCreatedAt(*t)
	}
	return ntc
}

// SetUpdatedAt sets the "updatedAt" field.
func (ntc *NovelTagCreate) SetUpdatedAt(t time.Time) *NovelTagCreate {
	ntc.mutation.SetUpdatedAt(t)
	return ntc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (ntc *NovelTagCreate) SetNillableUpdatedAt(t *time.Time) *NovelTagCreate {
	if t != nil {
		ntc.SetUpdatedAt(*t)
	}
	return ntc
}

// SetCreateBy sets the "createBy" field.
func (ntc *NovelTagCreate) SetCreateBy(i int64) *NovelTagCreate {
	ntc.mutation.SetCreateBy(i)
	return ntc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (ntc *NovelTagCreate) SetNillableCreateBy(i *int64) *NovelTagCreate {
	if i != nil {
		ntc.SetCreateBy(*i)
	}
	return ntc
}

// SetUpdateBy sets the "updateBy" field.
func (ntc *NovelTagCreate) SetUpdateBy(i int64) *NovelTagCreate {
	ntc.mutation.SetUpdateBy(i)
	return ntc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (ntc *NovelTagCreate) SetNillableUpdateBy(i *int64) *NovelTagCreate {
	if i != nil {
		ntc.SetUpdateBy(*i)
	}
	return ntc
}

// SetTenantId sets the "tenantId" field.
func (ntc *NovelTagCreate) SetTenantId(i int64) *NovelTagCreate {
	ntc.mutation.SetTenantId(i)
	return ntc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (ntc *NovelTagCreate) SetNillableTenantId(i *int64) *NovelTagCreate {
	if i != nil {
		ntc.SetTenantId(*i)
	}
	return ntc
}

// Mutation returns the NovelTagMutation object of the builder.
func (ntc *NovelTagCreate) Mutation() *NovelTagMutation {
	return ntc.mutation
}

// Save creates the NovelTag in the database.
func (ntc *NovelTagCreate) Save(ctx context.Context) (*NovelTag, error) {
	var (
		err  error
		node *NovelTag
	)
	ntc.defaults()
	if len(ntc.hooks) == 0 {
		if err = ntc.check(); err != nil {
			return nil, err
		}
		node, err = ntc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ntc.check(); err != nil {
				return nil, err
			}
			ntc.mutation = mutation
			if node, err = ntc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ntc.hooks) - 1; i >= 0; i-- {
			if ntc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ntc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ntc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ntc *NovelTagCreate) SaveX(ctx context.Context) *NovelTag {
	v, err := ntc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ntc *NovelTagCreate) Exec(ctx context.Context) error {
	_, err := ntc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntc *NovelTagCreate) ExecX(ctx context.Context) {
	if err := ntc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ntc *NovelTagCreate) defaults() {
	if _, ok := ntc.mutation.CreatedAt(); !ok {
		v := noveltag.DefaultCreatedAt()
		ntc.mutation.SetCreatedAt(v)
	}
	if _, ok := ntc.mutation.UpdatedAt(); !ok {
		v := noveltag.DefaultUpdatedAt()
		ntc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ntc.mutation.CreateBy(); !ok {
		v := noveltag.DefaultCreateBy
		ntc.mutation.SetCreateBy(v)
	}
	if _, ok := ntc.mutation.UpdateBy(); !ok {
		v := noveltag.DefaultUpdateBy
		ntc.mutation.SetUpdateBy(v)
	}
	if _, ok := ntc.mutation.TenantId(); !ok {
		v := noveltag.DefaultTenantId
		ntc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ntc *NovelTagCreate) check() error {
	if _, ok := ntc.mutation.TagId(); !ok {
		return &ValidationError{Name: "tagId", err: errors.New(`ent: missing required field "tagId"`)}
	}
	if _, ok := ntc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := ntc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := ntc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := ntc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := ntc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (ntc *NovelTagCreate) sqlSave(ctx context.Context) (*NovelTag, error) {
	_node, _spec := ntc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ntc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (ntc *NovelTagCreate) createSpec() (*NovelTag, *sqlgraph.CreateSpec) {
	var (
		_node = &NovelTag{config: ntc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: noveltag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: noveltag.FieldID,
			},
		}
	)
	if value, ok := ntc.mutation.TagId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: noveltag.FieldTagId,
		})
		_node.TagId = value
	}
	if value, ok := ntc.mutation.TagName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: noveltag.FieldTagName,
		})
		_node.TagName = value
	}
	if value, ok := ntc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: noveltag.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := ntc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: noveltag.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ntc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: noveltag.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ntc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: noveltag.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := ntc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: noveltag.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := ntc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: noveltag.FieldTenantId,
		})
		_node.TenantId = value
	}
	return _node, _spec
}

// NovelTagCreateBulk is the builder for creating many NovelTag entities in bulk.
type NovelTagCreateBulk struct {
	config
	builders []*NovelTagCreate
}

// Save creates the NovelTag entities in the database.
func (ntcb *NovelTagCreateBulk) Save(ctx context.Context) ([]*NovelTag, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ntcb.builders))
	nodes := make([]*NovelTag, len(ntcb.builders))
	mutators := make([]Mutator, len(ntcb.builders))
	for i := range ntcb.builders {
		func(i int, root context.Context) {
			builder := ntcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NovelTagMutation)
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
					_, err = mutators[i+1].Mutate(root, ntcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ntcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ntcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ntcb *NovelTagCreateBulk) SaveX(ctx context.Context) []*NovelTag {
	v, err := ntcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ntcb *NovelTagCreateBulk) Exec(ctx context.Context) error {
	_, err := ntcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntcb *NovelTagCreateBulk) ExecX(ctx context.Context) {
	if err := ntcb.Exec(ctx); err != nil {
		panic(err)
	}
}