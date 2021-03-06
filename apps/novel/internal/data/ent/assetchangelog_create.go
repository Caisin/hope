// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/assetchangelog"
	"hope/apps/novel/internal/data/ent/socialuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AssetChangeLogCreate is the builder for creating a AssetChangeLog entity.
type AssetChangeLogCreate struct {
	config
	mutation *AssetChangeLogMutation
	hooks    []Hook
}

// SetOrderId sets the "orderId" field.
func (aclc *AssetChangeLogCreate) SetOrderId(s string) *AssetChangeLogCreate {
	aclc.mutation.SetOrderId(s)
	return aclc
}

// SetNillableOrderId sets the "orderId" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableOrderId(s *string) *AssetChangeLogCreate {
	if s != nil {
		aclc.SetOrderId(*s)
	}
	return aclc
}

// SetBalanceId sets the "balanceId" field.
func (aclc *AssetChangeLogCreate) SetBalanceId(i int64) *AssetChangeLogCreate {
	aclc.mutation.SetBalanceId(i)
	return aclc
}

// SetNillableBalanceId sets the "balanceId" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableBalanceId(i *int64) *AssetChangeLogCreate {
	if i != nil {
		aclc.SetBalanceId(*i)
	}
	return aclc
}

// SetEventId sets the "eventId" field.
func (aclc *AssetChangeLogCreate) SetEventId(i int64) *AssetChangeLogCreate {
	aclc.mutation.SetEventId(i)
	return aclc
}

// SetNillableEventId sets the "eventId" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableEventId(i *int64) *AssetChangeLogCreate {
	if i != nil {
		aclc.SetEventId(*i)
	}
	return aclc
}

// SetUserId sets the "userId" field.
func (aclc *AssetChangeLogCreate) SetUserId(i int64) *AssetChangeLogCreate {
	aclc.mutation.SetUserId(i)
	return aclc
}

// SetAssetItemId sets the "assetItemId" field.
func (aclc *AssetChangeLogCreate) SetAssetItemId(i int32) *AssetChangeLogCreate {
	aclc.mutation.SetAssetItemId(i)
	return aclc
}

// SetNillableAssetItemId sets the "assetItemId" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableAssetItemId(i *int32) *AssetChangeLogCreate {
	if i != nil {
		aclc.SetAssetItemId(*i)
	}
	return aclc
}

// SetAmount sets the "amount" field.
func (aclc *AssetChangeLogCreate) SetAmount(i int64) *AssetChangeLogCreate {
	aclc.mutation.SetAmount(i)
	return aclc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableAmount(i *int64) *AssetChangeLogCreate {
	if i != nil {
		aclc.SetAmount(*i)
	}
	return aclc
}

// SetOldBalance sets the "oldBalance" field.
func (aclc *AssetChangeLogCreate) SetOldBalance(i int64) *AssetChangeLogCreate {
	aclc.mutation.SetOldBalance(i)
	return aclc
}

// SetNillableOldBalance sets the "oldBalance" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableOldBalance(i *int64) *AssetChangeLogCreate {
	if i != nil {
		aclc.SetOldBalance(*i)
	}
	return aclc
}

// SetNewBalance sets the "newBalance" field.
func (aclc *AssetChangeLogCreate) SetNewBalance(i int64) *AssetChangeLogCreate {
	aclc.mutation.SetNewBalance(i)
	return aclc
}

// SetNillableNewBalance sets the "newBalance" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableNewBalance(i *int64) *AssetChangeLogCreate {
	if i != nil {
		aclc.SetNewBalance(*i)
	}
	return aclc
}

// SetRemark sets the "remark" field.
func (aclc *AssetChangeLogCreate) SetRemark(s string) *AssetChangeLogCreate {
	aclc.mutation.SetRemark(s)
	return aclc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableRemark(s *string) *AssetChangeLogCreate {
	if s != nil {
		aclc.SetRemark(*s)
	}
	return aclc
}

// SetCreatedAt sets the "createdAt" field.
func (aclc *AssetChangeLogCreate) SetCreatedAt(t time.Time) *AssetChangeLogCreate {
	aclc.mutation.SetCreatedAt(t)
	return aclc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableCreatedAt(t *time.Time) *AssetChangeLogCreate {
	if t != nil {
		aclc.SetCreatedAt(*t)
	}
	return aclc
}

// SetUpdatedAt sets the "updatedAt" field.
func (aclc *AssetChangeLogCreate) SetUpdatedAt(t time.Time) *AssetChangeLogCreate {
	aclc.mutation.SetUpdatedAt(t)
	return aclc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableUpdatedAt(t *time.Time) *AssetChangeLogCreate {
	if t != nil {
		aclc.SetUpdatedAt(*t)
	}
	return aclc
}

// SetCreateBy sets the "createBy" field.
func (aclc *AssetChangeLogCreate) SetCreateBy(i int64) *AssetChangeLogCreate {
	aclc.mutation.SetCreateBy(i)
	return aclc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableCreateBy(i *int64) *AssetChangeLogCreate {
	if i != nil {
		aclc.SetCreateBy(*i)
	}
	return aclc
}

// SetUpdateBy sets the "updateBy" field.
func (aclc *AssetChangeLogCreate) SetUpdateBy(i int64) *AssetChangeLogCreate {
	aclc.mutation.SetUpdateBy(i)
	return aclc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableUpdateBy(i *int64) *AssetChangeLogCreate {
	if i != nil {
		aclc.SetUpdateBy(*i)
	}
	return aclc
}

// SetTenantId sets the "tenantId" field.
func (aclc *AssetChangeLogCreate) SetTenantId(i int64) *AssetChangeLogCreate {
	aclc.mutation.SetTenantId(i)
	return aclc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (aclc *AssetChangeLogCreate) SetNillableTenantId(i *int64) *AssetChangeLogCreate {
	if i != nil {
		aclc.SetTenantId(*i)
	}
	return aclc
}

// SetUserID sets the "user" edge to the SocialUser entity by ID.
func (aclc *AssetChangeLogCreate) SetUserID(id int64) *AssetChangeLogCreate {
	aclc.mutation.SetUserID(id)
	return aclc
}

// SetUser sets the "user" edge to the SocialUser entity.
func (aclc *AssetChangeLogCreate) SetUser(s *SocialUser) *AssetChangeLogCreate {
	return aclc.SetUserID(s.ID)
}

// Mutation returns the AssetChangeLogMutation object of the builder.
func (aclc *AssetChangeLogCreate) Mutation() *AssetChangeLogMutation {
	return aclc.mutation
}

// Save creates the AssetChangeLog in the database.
func (aclc *AssetChangeLogCreate) Save(ctx context.Context) (*AssetChangeLog, error) {
	var (
		err  error
		node *AssetChangeLog
	)
	aclc.defaults()
	if len(aclc.hooks) == 0 {
		if err = aclc.check(); err != nil {
			return nil, err
		}
		node, err = aclc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetChangeLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aclc.check(); err != nil {
				return nil, err
			}
			aclc.mutation = mutation
			if node, err = aclc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aclc.hooks) - 1; i >= 0; i-- {
			if aclc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aclc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aclc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aclc *AssetChangeLogCreate) SaveX(ctx context.Context) *AssetChangeLog {
	v, err := aclc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aclc *AssetChangeLogCreate) Exec(ctx context.Context) error {
	_, err := aclc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aclc *AssetChangeLogCreate) ExecX(ctx context.Context) {
	if err := aclc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aclc *AssetChangeLogCreate) defaults() {
	if _, ok := aclc.mutation.CreatedAt(); !ok {
		v := assetchangelog.DefaultCreatedAt()
		aclc.mutation.SetCreatedAt(v)
	}
	if _, ok := aclc.mutation.UpdatedAt(); !ok {
		v := assetchangelog.DefaultUpdatedAt()
		aclc.mutation.SetUpdatedAt(v)
	}
	if _, ok := aclc.mutation.CreateBy(); !ok {
		v := assetchangelog.DefaultCreateBy
		aclc.mutation.SetCreateBy(v)
	}
	if _, ok := aclc.mutation.UpdateBy(); !ok {
		v := assetchangelog.DefaultUpdateBy
		aclc.mutation.SetUpdateBy(v)
	}
	if _, ok := aclc.mutation.TenantId(); !ok {
		v := assetchangelog.DefaultTenantId
		aclc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aclc *AssetChangeLogCreate) check() error {
	if _, ok := aclc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "AssetChangeLog.userId"`)}
	}
	if _, ok := aclc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "AssetChangeLog.createdAt"`)}
	}
	if _, ok := aclc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "AssetChangeLog.updatedAt"`)}
	}
	if _, ok := aclc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "AssetChangeLog.createBy"`)}
	}
	if _, ok := aclc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "AssetChangeLog.updateBy"`)}
	}
	if _, ok := aclc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "AssetChangeLog.tenantId"`)}
	}
	if _, ok := aclc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "AssetChangeLog.user"`)}
	}
	return nil
}

func (aclc *AssetChangeLogCreate) sqlSave(ctx context.Context) (*AssetChangeLog, error) {
	_node, _spec := aclc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aclc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (aclc *AssetChangeLogCreate) createSpec() (*AssetChangeLog, *sqlgraph.CreateSpec) {
	var (
		_node = &AssetChangeLog{config: aclc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: assetchangelog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: assetchangelog.FieldID,
			},
		}
	)
	if value, ok := aclc.mutation.OrderId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assetchangelog.FieldOrderId,
		})
		_node.OrderId = value
	}
	if value, ok := aclc.mutation.BalanceId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetchangelog.FieldBalanceId,
		})
		_node.BalanceId = value
	}
	if value, ok := aclc.mutation.EventId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetchangelog.FieldEventId,
		})
		_node.EventId = value
	}
	if value, ok := aclc.mutation.AssetItemId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetchangelog.FieldAssetItemId,
		})
		_node.AssetItemId = value
	}
	if value, ok := aclc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetchangelog.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := aclc.mutation.OldBalance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetchangelog.FieldOldBalance,
		})
		_node.OldBalance = value
	}
	if value, ok := aclc.mutation.NewBalance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetchangelog.FieldNewBalance,
		})
		_node.NewBalance = value
	}
	if value, ok := aclc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assetchangelog.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := aclc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetchangelog.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := aclc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetchangelog.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := aclc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetchangelog.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := aclc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetchangelog.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := aclc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetchangelog.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := aclc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   assetchangelog.UserTable,
			Columns: []string{assetchangelog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: socialuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AssetChangeLogCreateBulk is the builder for creating many AssetChangeLog entities in bulk.
type AssetChangeLogCreateBulk struct {
	config
	builders []*AssetChangeLogCreate
}

// Save creates the AssetChangeLog entities in the database.
func (aclcb *AssetChangeLogCreateBulk) Save(ctx context.Context) ([]*AssetChangeLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aclcb.builders))
	nodes := make([]*AssetChangeLog, len(aclcb.builders))
	mutators := make([]Mutator, len(aclcb.builders))
	for i := range aclcb.builders {
		func(i int, root context.Context) {
			builder := aclcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AssetChangeLogMutation)
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
					_, err = mutators[i+1].Mutate(root, aclcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aclcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aclcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aclcb *AssetChangeLogCreateBulk) SaveX(ctx context.Context) []*AssetChangeLog {
	v, err := aclcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aclcb *AssetChangeLogCreateBulk) Exec(ctx context.Context) error {
	_, err := aclcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aclcb *AssetChangeLogCreateBulk) ExecX(ctx context.Context) {
	if err := aclcb.Exec(ctx); err != nil {
		panic(err)
	}
}
