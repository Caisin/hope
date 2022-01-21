// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/param/internal/data/ent/viptype"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VipTypeCreate is the builder for creating a VipType entity.
type VipTypeCreate struct {
	config
	mutation *VipTypeMutation
	hooks    []Hook
}

// SetVipName sets the "vipName" field.
func (vtc *VipTypeCreate) SetVipName(s string) *VipTypeCreate {
	vtc.mutation.SetVipName(s)
	return vtc
}

// SetNillableVipName sets the "vipName" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableVipName(s *string) *VipTypeCreate {
	if s != nil {
		vtc.SetVipName(*s)
	}
	return vtc
}

// SetIsSuper sets the "isSuper" field.
func (vtc *VipTypeCreate) SetIsSuper(b bool) *VipTypeCreate {
	vtc.mutation.SetIsSuper(b)
	return vtc
}

// SetNillableIsSuper sets the "isSuper" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableIsSuper(b *bool) *VipTypeCreate {
	if b != nil {
		vtc.SetIsSuper(*b)
	}
	return vtc
}

// SetValidDays sets the "validDays" field.
func (vtc *VipTypeCreate) SetValidDays(i int32) *VipTypeCreate {
	vtc.mutation.SetValidDays(i)
	return vtc
}

// SetNillableValidDays sets the "validDays" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableValidDays(i *int32) *VipTypeCreate {
	if i != nil {
		vtc.SetValidDays(*i)
	}
	return vtc
}

// SetDiscountRate sets the "discountRate" field.
func (vtc *VipTypeCreate) SetDiscountRate(i int64) *VipTypeCreate {
	vtc.mutation.SetDiscountRate(i)
	return vtc
}

// SetNillableDiscountRate sets the "discountRate" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableDiscountRate(i *int64) *VipTypeCreate {
	if i != nil {
		vtc.SetDiscountRate(*i)
	}
	return vtc
}

// SetAvatarId sets the "avatarId" field.
func (vtc *VipTypeCreate) SetAvatarId(i int64) *VipTypeCreate {
	vtc.mutation.SetAvatarId(i)
	return vtc
}

// SetNillableAvatarId sets the "avatarId" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableAvatarId(i *int64) *VipTypeCreate {
	if i != nil {
		vtc.SetAvatarId(*i)
	}
	return vtc
}

// SetSummary sets the "summary" field.
func (vtc *VipTypeCreate) SetSummary(s string) *VipTypeCreate {
	vtc.mutation.SetSummary(s)
	return vtc
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableSummary(s *string) *VipTypeCreate {
	if s != nil {
		vtc.SetSummary(*s)
	}
	return vtc
}

// SetCreatedAt sets the "createdAt" field.
func (vtc *VipTypeCreate) SetCreatedAt(t time.Time) *VipTypeCreate {
	vtc.mutation.SetCreatedAt(t)
	return vtc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableCreatedAt(t *time.Time) *VipTypeCreate {
	if t != nil {
		vtc.SetCreatedAt(*t)
	}
	return vtc
}

// SetUpdatedAt sets the "updatedAt" field.
func (vtc *VipTypeCreate) SetUpdatedAt(t time.Time) *VipTypeCreate {
	vtc.mutation.SetUpdatedAt(t)
	return vtc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableUpdatedAt(t *time.Time) *VipTypeCreate {
	if t != nil {
		vtc.SetUpdatedAt(*t)
	}
	return vtc
}

// SetCreateBy sets the "createBy" field.
func (vtc *VipTypeCreate) SetCreateBy(i int64) *VipTypeCreate {
	vtc.mutation.SetCreateBy(i)
	return vtc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableCreateBy(i *int64) *VipTypeCreate {
	if i != nil {
		vtc.SetCreateBy(*i)
	}
	return vtc
}

// SetUpdateBy sets the "updateBy" field.
func (vtc *VipTypeCreate) SetUpdateBy(i int64) *VipTypeCreate {
	vtc.mutation.SetUpdateBy(i)
	return vtc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableUpdateBy(i *int64) *VipTypeCreate {
	if i != nil {
		vtc.SetUpdateBy(*i)
	}
	return vtc
}

// SetTenantId sets the "tenantId" field.
func (vtc *VipTypeCreate) SetTenantId(i int64) *VipTypeCreate {
	vtc.mutation.SetTenantId(i)
	return vtc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (vtc *VipTypeCreate) SetNillableTenantId(i *int64) *VipTypeCreate {
	if i != nil {
		vtc.SetTenantId(*i)
	}
	return vtc
}

// Mutation returns the VipTypeMutation object of the builder.
func (vtc *VipTypeCreate) Mutation() *VipTypeMutation {
	return vtc.mutation
}

// Save creates the VipType in the database.
func (vtc *VipTypeCreate) Save(ctx context.Context) (*VipType, error) {
	var (
		err  error
		node *VipType
	)
	vtc.defaults()
	if len(vtc.hooks) == 0 {
		if err = vtc.check(); err != nil {
			return nil, err
		}
		node, err = vtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VipTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vtc.check(); err != nil {
				return nil, err
			}
			vtc.mutation = mutation
			if node, err = vtc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vtc.hooks) - 1; i >= 0; i-- {
			if vtc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vtc *VipTypeCreate) SaveX(ctx context.Context) *VipType {
	v, err := vtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vtc *VipTypeCreate) Exec(ctx context.Context) error {
	_, err := vtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtc *VipTypeCreate) ExecX(ctx context.Context) {
	if err := vtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vtc *VipTypeCreate) defaults() {
	if _, ok := vtc.mutation.CreatedAt(); !ok {
		v := viptype.DefaultCreatedAt()
		vtc.mutation.SetCreatedAt(v)
	}
	if _, ok := vtc.mutation.UpdatedAt(); !ok {
		v := viptype.DefaultUpdatedAt()
		vtc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vtc.mutation.CreateBy(); !ok {
		v := viptype.DefaultCreateBy
		vtc.mutation.SetCreateBy(v)
	}
	if _, ok := vtc.mutation.UpdateBy(); !ok {
		v := viptype.DefaultUpdateBy
		vtc.mutation.SetUpdateBy(v)
	}
	if _, ok := vtc.mutation.TenantId(); !ok {
		v := viptype.DefaultTenantId
		vtc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vtc *VipTypeCreate) check() error {
	if _, ok := vtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := vtc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := vtc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := vtc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := vtc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (vtc *VipTypeCreate) sqlSave(ctx context.Context) (*VipType, error) {
	_node, _spec := vtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (vtc *VipTypeCreate) createSpec() (*VipType, *sqlgraph.CreateSpec) {
	var (
		_node = &VipType{config: vtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: viptype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: viptype.FieldID,
			},
		}
	)
	if value, ok := vtc.mutation.VipName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: viptype.FieldVipName,
		})
		_node.VipName = value
	}
	if value, ok := vtc.mutation.IsSuper(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: viptype.FieldIsSuper,
		})
		_node.IsSuper = value
	}
	if value, ok := vtc.mutation.ValidDays(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: viptype.FieldValidDays,
		})
		_node.ValidDays = value
	}
	if value, ok := vtc.mutation.DiscountRate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldDiscountRate,
		})
		_node.DiscountRate = value
	}
	if value, ok := vtc.mutation.AvatarId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldAvatarId,
		})
		_node.AvatarId = value
	}
	if value, ok := vtc.mutation.Summary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: viptype.FieldSummary,
		})
		_node.Summary = value
	}
	if value, ok := vtc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: viptype.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vtc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: viptype.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := vtc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := vtc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := vtc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldTenantId,
		})
		_node.TenantId = value
	}
	return _node, _spec
}

// VipTypeCreateBulk is the builder for creating many VipType entities in bulk.
type VipTypeCreateBulk struct {
	config
	builders []*VipTypeCreate
}

// Save creates the VipType entities in the database.
func (vtcb *VipTypeCreateBulk) Save(ctx context.Context) ([]*VipType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vtcb.builders))
	nodes := make([]*VipType, len(vtcb.builders))
	mutators := make([]Mutator, len(vtcb.builders))
	for i := range vtcb.builders {
		func(i int, root context.Context) {
			builder := vtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VipTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, vtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vtcb *VipTypeCreateBulk) SaveX(ctx context.Context) []*VipType {
	v, err := vtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vtcb *VipTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := vtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtcb *VipTypeCreateBulk) ExecX(ctx context.Context) {
	if err := vtcb.Exec(ctx); err != nil {
		panic(err)
	}
}