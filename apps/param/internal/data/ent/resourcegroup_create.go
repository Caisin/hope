// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/param/internal/data/ent/resourcegroup"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ResourceGroupCreate is the builder for creating a ResourceGroup entity.
type ResourceGroupCreate struct {
	config
	mutation *ResourceGroupMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rgc *ResourceGroupCreate) SetName(s string) *ResourceGroupCreate {
	rgc.mutation.SetName(s)
	return rgc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rgc *ResourceGroupCreate) SetNillableName(s *string) *ResourceGroupCreate {
	if s != nil {
		rgc.SetName(*s)
	}
	return rgc
}

// SetCreatedAt sets the "createdAt" field.
func (rgc *ResourceGroupCreate) SetCreatedAt(t time.Time) *ResourceGroupCreate {
	rgc.mutation.SetCreatedAt(t)
	return rgc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (rgc *ResourceGroupCreate) SetNillableCreatedAt(t *time.Time) *ResourceGroupCreate {
	if t != nil {
		rgc.SetCreatedAt(*t)
	}
	return rgc
}

// SetUpdatedAt sets the "updatedAt" field.
func (rgc *ResourceGroupCreate) SetUpdatedAt(t time.Time) *ResourceGroupCreate {
	rgc.mutation.SetUpdatedAt(t)
	return rgc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (rgc *ResourceGroupCreate) SetNillableUpdatedAt(t *time.Time) *ResourceGroupCreate {
	if t != nil {
		rgc.SetUpdatedAt(*t)
	}
	return rgc
}

// SetCreateBy sets the "createBy" field.
func (rgc *ResourceGroupCreate) SetCreateBy(i int64) *ResourceGroupCreate {
	rgc.mutation.SetCreateBy(i)
	return rgc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (rgc *ResourceGroupCreate) SetNillableCreateBy(i *int64) *ResourceGroupCreate {
	if i != nil {
		rgc.SetCreateBy(*i)
	}
	return rgc
}

// SetUpdateBy sets the "updateBy" field.
func (rgc *ResourceGroupCreate) SetUpdateBy(i int64) *ResourceGroupCreate {
	rgc.mutation.SetUpdateBy(i)
	return rgc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (rgc *ResourceGroupCreate) SetNillableUpdateBy(i *int64) *ResourceGroupCreate {
	if i != nil {
		rgc.SetUpdateBy(*i)
	}
	return rgc
}

// SetTenantId sets the "tenantId" field.
func (rgc *ResourceGroupCreate) SetTenantId(i int64) *ResourceGroupCreate {
	rgc.mutation.SetTenantId(i)
	return rgc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (rgc *ResourceGroupCreate) SetNillableTenantId(i *int64) *ResourceGroupCreate {
	if i != nil {
		rgc.SetTenantId(*i)
	}
	return rgc
}

// Mutation returns the ResourceGroupMutation object of the builder.
func (rgc *ResourceGroupCreate) Mutation() *ResourceGroupMutation {
	return rgc.mutation
}

// Save creates the ResourceGroup in the database.
func (rgc *ResourceGroupCreate) Save(ctx context.Context) (*ResourceGroup, error) {
	var (
		err  error
		node *ResourceGroup
	)
	rgc.defaults()
	if len(rgc.hooks) == 0 {
		if err = rgc.check(); err != nil {
			return nil, err
		}
		node, err = rgc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourceGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rgc.check(); err != nil {
				return nil, err
			}
			rgc.mutation = mutation
			if node, err = rgc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rgc.hooks) - 1; i >= 0; i-- {
			if rgc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rgc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rgc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rgc *ResourceGroupCreate) SaveX(ctx context.Context) *ResourceGroup {
	v, err := rgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rgc *ResourceGroupCreate) Exec(ctx context.Context) error {
	_, err := rgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rgc *ResourceGroupCreate) ExecX(ctx context.Context) {
	if err := rgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rgc *ResourceGroupCreate) defaults() {
	if _, ok := rgc.mutation.CreatedAt(); !ok {
		v := resourcegroup.DefaultCreatedAt()
		rgc.mutation.SetCreatedAt(v)
	}
	if _, ok := rgc.mutation.UpdatedAt(); !ok {
		v := resourcegroup.DefaultUpdatedAt()
		rgc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rgc.mutation.CreateBy(); !ok {
		v := resourcegroup.DefaultCreateBy
		rgc.mutation.SetCreateBy(v)
	}
	if _, ok := rgc.mutation.UpdateBy(); !ok {
		v := resourcegroup.DefaultUpdateBy
		rgc.mutation.SetUpdateBy(v)
	}
	if _, ok := rgc.mutation.TenantId(); !ok {
		v := resourcegroup.DefaultTenantId
		rgc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rgc *ResourceGroupCreate) check() error {
	if _, ok := rgc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := rgc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := rgc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := rgc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := rgc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (rgc *ResourceGroupCreate) sqlSave(ctx context.Context) (*ResourceGroup, error) {
	_node, _spec := rgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rgc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (rgc *ResourceGroupCreate) createSpec() (*ResourceGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &ResourceGroup{config: rgc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: resourcegroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: resourcegroup.FieldID,
			},
		}
	)
	if value, ok := rgc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resourcegroup.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rgc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: resourcegroup.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rgc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: resourcegroup.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := rgc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: resourcegroup.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := rgc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: resourcegroup.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := rgc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: resourcegroup.FieldTenantId,
		})
		_node.TenantId = value
	}
	return _node, _spec
}

// ResourceGroupCreateBulk is the builder for creating many ResourceGroup entities in bulk.
type ResourceGroupCreateBulk struct {
	config
	builders []*ResourceGroupCreate
}

// Save creates the ResourceGroup entities in the database.
func (rgcb *ResourceGroupCreateBulk) Save(ctx context.Context) ([]*ResourceGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rgcb.builders))
	nodes := make([]*ResourceGroup, len(rgcb.builders))
	mutators := make([]Mutator, len(rgcb.builders))
	for i := range rgcb.builders {
		func(i int, root context.Context) {
			builder := rgcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResourceGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, rgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rgcb *ResourceGroupCreateBulk) SaveX(ctx context.Context) []*ResourceGroup {
	v, err := rgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rgcb *ResourceGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := rgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rgcb *ResourceGroupCreateBulk) ExecX(ctx context.Context) {
	if err := rgcb.Exec(ctx); err != nil {
		panic(err)
	}
}