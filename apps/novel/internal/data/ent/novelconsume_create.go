// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelconsume"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelConsumeCreate is the builder for creating a NovelConsume entity.
type NovelConsumeCreate struct {
	config
	mutation *NovelConsumeMutation
	hooks    []Hook
}

// SetNovelId sets the "novelId" field.
func (ncc *NovelConsumeCreate) SetNovelId(i int64) *NovelConsumeCreate {
	ncc.mutation.SetNovelId(i)
	return ncc
}

// SetCoin sets the "coin" field.
func (ncc *NovelConsumeCreate) SetCoin(i int64) *NovelConsumeCreate {
	ncc.mutation.SetCoin(i)
	return ncc
}

// SetNillableCoin sets the "coin" field if the given value is not nil.
func (ncc *NovelConsumeCreate) SetNillableCoin(i *int64) *NovelConsumeCreate {
	if i != nil {
		ncc.SetCoin(*i)
	}
	return ncc
}

// SetCoupon sets the "coupon" field.
func (ncc *NovelConsumeCreate) SetCoupon(i int64) *NovelConsumeCreate {
	ncc.mutation.SetCoupon(i)
	return ncc
}

// SetNillableCoupon sets the "coupon" field if the given value is not nil.
func (ncc *NovelConsumeCreate) SetNillableCoupon(i *int64) *NovelConsumeCreate {
	if i != nil {
		ncc.SetCoupon(*i)
	}
	return ncc
}

// SetDiscount sets the "discount" field.
func (ncc *NovelConsumeCreate) SetDiscount(i int64) *NovelConsumeCreate {
	ncc.mutation.SetDiscount(i)
	return ncc
}

// SetNillableDiscount sets the "discount" field if the given value is not nil.
func (ncc *NovelConsumeCreate) SetNillableDiscount(i *int64) *NovelConsumeCreate {
	if i != nil {
		ncc.SetDiscount(*i)
	}
	return ncc
}

// SetReward sets the "reward" field.
func (ncc *NovelConsumeCreate) SetReward(i int64) *NovelConsumeCreate {
	ncc.mutation.SetReward(i)
	return ncc
}

// SetNillableReward sets the "reward" field if the given value is not nil.
func (ncc *NovelConsumeCreate) SetNillableReward(i *int64) *NovelConsumeCreate {
	if i != nil {
		ncc.SetReward(*i)
	}
	return ncc
}

// SetCreatedAt sets the "createdAt" field.
func (ncc *NovelConsumeCreate) SetCreatedAt(t time.Time) *NovelConsumeCreate {
	ncc.mutation.SetCreatedAt(t)
	return ncc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (ncc *NovelConsumeCreate) SetNillableCreatedAt(t *time.Time) *NovelConsumeCreate {
	if t != nil {
		ncc.SetCreatedAt(*t)
	}
	return ncc
}

// SetUpdatedAt sets the "updatedAt" field.
func (ncc *NovelConsumeCreate) SetUpdatedAt(t time.Time) *NovelConsumeCreate {
	ncc.mutation.SetUpdatedAt(t)
	return ncc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (ncc *NovelConsumeCreate) SetNillableUpdatedAt(t *time.Time) *NovelConsumeCreate {
	if t != nil {
		ncc.SetUpdatedAt(*t)
	}
	return ncc
}

// SetCreateBy sets the "createBy" field.
func (ncc *NovelConsumeCreate) SetCreateBy(i int64) *NovelConsumeCreate {
	ncc.mutation.SetCreateBy(i)
	return ncc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (ncc *NovelConsumeCreate) SetNillableCreateBy(i *int64) *NovelConsumeCreate {
	if i != nil {
		ncc.SetCreateBy(*i)
	}
	return ncc
}

// SetUpdateBy sets the "updateBy" field.
func (ncc *NovelConsumeCreate) SetUpdateBy(i int64) *NovelConsumeCreate {
	ncc.mutation.SetUpdateBy(i)
	return ncc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (ncc *NovelConsumeCreate) SetNillableUpdateBy(i *int64) *NovelConsumeCreate {
	if i != nil {
		ncc.SetUpdateBy(*i)
	}
	return ncc
}

// SetTenantId sets the "tenantId" field.
func (ncc *NovelConsumeCreate) SetTenantId(i int64) *NovelConsumeCreate {
	ncc.mutation.SetTenantId(i)
	return ncc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (ncc *NovelConsumeCreate) SetNillableTenantId(i *int64) *NovelConsumeCreate {
	if i != nil {
		ncc.SetTenantId(*i)
	}
	return ncc
}

// Mutation returns the NovelConsumeMutation object of the builder.
func (ncc *NovelConsumeCreate) Mutation() *NovelConsumeMutation {
	return ncc.mutation
}

// Save creates the NovelConsume in the database.
func (ncc *NovelConsumeCreate) Save(ctx context.Context) (*NovelConsume, error) {
	var (
		err  error
		node *NovelConsume
	)
	ncc.defaults()
	if len(ncc.hooks) == 0 {
		if err = ncc.check(); err != nil {
			return nil, err
		}
		node, err = ncc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelConsumeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ncc.check(); err != nil {
				return nil, err
			}
			ncc.mutation = mutation
			if node, err = ncc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ncc.hooks) - 1; i >= 0; i-- {
			if ncc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ncc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ncc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ncc *NovelConsumeCreate) SaveX(ctx context.Context) *NovelConsume {
	v, err := ncc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncc *NovelConsumeCreate) Exec(ctx context.Context) error {
	_, err := ncc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncc *NovelConsumeCreate) ExecX(ctx context.Context) {
	if err := ncc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ncc *NovelConsumeCreate) defaults() {
	if _, ok := ncc.mutation.CreatedAt(); !ok {
		v := novelconsume.DefaultCreatedAt()
		ncc.mutation.SetCreatedAt(v)
	}
	if _, ok := ncc.mutation.UpdatedAt(); !ok {
		v := novelconsume.DefaultUpdatedAt()
		ncc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ncc.mutation.CreateBy(); !ok {
		v := novelconsume.DefaultCreateBy
		ncc.mutation.SetCreateBy(v)
	}
	if _, ok := ncc.mutation.UpdateBy(); !ok {
		v := novelconsume.DefaultUpdateBy
		ncc.mutation.SetUpdateBy(v)
	}
	if _, ok := ncc.mutation.TenantId(); !ok {
		v := novelconsume.DefaultTenantId
		ncc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ncc *NovelConsumeCreate) check() error {
	if _, ok := ncc.mutation.NovelId(); !ok {
		return &ValidationError{Name: "novelId", err: errors.New(`ent: missing required field "novelId"`)}
	}
	if _, ok := ncc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := ncc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := ncc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := ncc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := ncc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (ncc *NovelConsumeCreate) sqlSave(ctx context.Context) (*NovelConsume, error) {
	_node, _spec := ncc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ncc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (ncc *NovelConsumeCreate) createSpec() (*NovelConsume, *sqlgraph.CreateSpec) {
	var (
		_node = &NovelConsume{config: ncc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: novelconsume.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelconsume.FieldID,
			},
		}
	)
	if value, ok := ncc.mutation.NovelId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldNovelId,
		})
		_node.NovelId = value
	}
	if value, ok := ncc.mutation.Coin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCoin,
		})
		_node.Coin = value
	}
	if value, ok := ncc.mutation.Coupon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCoupon,
		})
		_node.Coupon = value
	}
	if value, ok := ncc.mutation.Discount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldDiscount,
		})
		_node.Discount = value
	}
	if value, ok := ncc.mutation.Reward(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldReward,
		})
		_node.Reward = value
	}
	if value, ok := ncc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelconsume.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ncc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelconsume.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ncc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := ncc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := ncc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldTenantId,
		})
		_node.TenantId = value
	}
	return _node, _spec
}

// NovelConsumeCreateBulk is the builder for creating many NovelConsume entities in bulk.
type NovelConsumeCreateBulk struct {
	config
	builders []*NovelConsumeCreate
}

// Save creates the NovelConsume entities in the database.
func (nccb *NovelConsumeCreateBulk) Save(ctx context.Context) ([]*NovelConsume, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nccb.builders))
	nodes := make([]*NovelConsume, len(nccb.builders))
	mutators := make([]Mutator, len(nccb.builders))
	for i := range nccb.builders {
		func(i int, root context.Context) {
			builder := nccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NovelConsumeMutation)
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
					_, err = mutators[i+1].Mutate(root, nccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, nccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nccb *NovelConsumeCreateBulk) SaveX(ctx context.Context) []*NovelConsume {
	v, err := nccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nccb *NovelConsumeCreateBulk) Exec(ctx context.Context) error {
	_, err := nccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nccb *NovelConsumeCreateBulk) ExecX(ctx context.Context) {
	if err := nccb.Exec(ctx); err != nil {
		panic(err)
	}
}