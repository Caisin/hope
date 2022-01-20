// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/admin/internal/data/ent/casbinrule"
	"hope/apps/admin/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CasbinRuleUpdate is the builder for updating CasbinRule entities.
type CasbinRuleUpdate struct {
	config
	hooks    []Hook
	mutation *CasbinRuleMutation
}

// Where appends a list predicates to the CasbinRuleUpdate builder.
func (cru *CasbinRuleUpdate) Where(ps ...predicate.CasbinRule) *CasbinRuleUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetPType sets the "p_type" field.
func (cru *CasbinRuleUpdate) SetPType(s string) *CasbinRuleUpdate {
	cru.mutation.SetPType(s)
	return cru
}

// SetNillablePType sets the "p_type" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillablePType(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetPType(*s)
	}
	return cru
}

// ClearPType clears the value of the "p_type" field.
func (cru *CasbinRuleUpdate) ClearPType() *CasbinRuleUpdate {
	cru.mutation.ClearPType()
	return cru
}

// SetV0 sets the "v0" field.
func (cru *CasbinRuleUpdate) SetV0(s string) *CasbinRuleUpdate {
	cru.mutation.SetV0(s)
	return cru
}

// SetNillableV0 sets the "v0" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV0(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV0(*s)
	}
	return cru
}

// ClearV0 clears the value of the "v0" field.
func (cru *CasbinRuleUpdate) ClearV0() *CasbinRuleUpdate {
	cru.mutation.ClearV0()
	return cru
}

// SetV1 sets the "v1" field.
func (cru *CasbinRuleUpdate) SetV1(s string) *CasbinRuleUpdate {
	cru.mutation.SetV1(s)
	return cru
}

// SetNillableV1 sets the "v1" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV1(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV1(*s)
	}
	return cru
}

// ClearV1 clears the value of the "v1" field.
func (cru *CasbinRuleUpdate) ClearV1() *CasbinRuleUpdate {
	cru.mutation.ClearV1()
	return cru
}

// SetV2 sets the "v2" field.
func (cru *CasbinRuleUpdate) SetV2(s string) *CasbinRuleUpdate {
	cru.mutation.SetV2(s)
	return cru
}

// SetNillableV2 sets the "v2" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV2(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV2(*s)
	}
	return cru
}

// ClearV2 clears the value of the "v2" field.
func (cru *CasbinRuleUpdate) ClearV2() *CasbinRuleUpdate {
	cru.mutation.ClearV2()
	return cru
}

// SetV3 sets the "v3" field.
func (cru *CasbinRuleUpdate) SetV3(s string) *CasbinRuleUpdate {
	cru.mutation.SetV3(s)
	return cru
}

// SetNillableV3 sets the "v3" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV3(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV3(*s)
	}
	return cru
}

// ClearV3 clears the value of the "v3" field.
func (cru *CasbinRuleUpdate) ClearV3() *CasbinRuleUpdate {
	cru.mutation.ClearV3()
	return cru
}

// SetV4 sets the "v4" field.
func (cru *CasbinRuleUpdate) SetV4(s string) *CasbinRuleUpdate {
	cru.mutation.SetV4(s)
	return cru
}

// SetNillableV4 sets the "v4" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV4(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV4(*s)
	}
	return cru
}

// ClearV4 clears the value of the "v4" field.
func (cru *CasbinRuleUpdate) ClearV4() *CasbinRuleUpdate {
	cru.mutation.ClearV4()
	return cru
}

// SetV5 sets the "v5" field.
func (cru *CasbinRuleUpdate) SetV5(s string) *CasbinRuleUpdate {
	cru.mutation.SetV5(s)
	return cru
}

// SetNillableV5 sets the "v5" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV5(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV5(*s)
	}
	return cru
}

// ClearV5 clears the value of the "v5" field.
func (cru *CasbinRuleUpdate) ClearV5() *CasbinRuleUpdate {
	cru.mutation.ClearV5()
	return cru
}

// SetUpdatedAt sets the "updatedAt" field.
func (cru *CasbinRuleUpdate) SetUpdatedAt(t time.Time) *CasbinRuleUpdate {
	cru.mutation.SetUpdatedAt(t)
	return cru
}

// SetCreateBy sets the "createBy" field.
func (cru *CasbinRuleUpdate) SetCreateBy(i int64) *CasbinRuleUpdate {
	cru.mutation.ResetCreateBy()
	cru.mutation.SetCreateBy(i)
	return cru
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableCreateBy(i *int64) *CasbinRuleUpdate {
	if i != nil {
		cru.SetCreateBy(*i)
	}
	return cru
}

// AddCreateBy adds i to the "createBy" field.
func (cru *CasbinRuleUpdate) AddCreateBy(i int64) *CasbinRuleUpdate {
	cru.mutation.AddCreateBy(i)
	return cru
}

// SetUpdateBy sets the "updateBy" field.
func (cru *CasbinRuleUpdate) SetUpdateBy(i int64) *CasbinRuleUpdate {
	cru.mutation.ResetUpdateBy()
	cru.mutation.SetUpdateBy(i)
	return cru
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableUpdateBy(i *int64) *CasbinRuleUpdate {
	if i != nil {
		cru.SetUpdateBy(*i)
	}
	return cru
}

// AddUpdateBy adds i to the "updateBy" field.
func (cru *CasbinRuleUpdate) AddUpdateBy(i int64) *CasbinRuleUpdate {
	cru.mutation.AddUpdateBy(i)
	return cru
}

// SetTenantId sets the "tenantId" field.
func (cru *CasbinRuleUpdate) SetTenantId(i int64) *CasbinRuleUpdate {
	cru.mutation.ResetTenantId()
	cru.mutation.SetTenantId(i)
	return cru
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableTenantId(i *int64) *CasbinRuleUpdate {
	if i != nil {
		cru.SetTenantId(*i)
	}
	return cru
}

// AddTenantId adds i to the "tenantId" field.
func (cru *CasbinRuleUpdate) AddTenantId(i int64) *CasbinRuleUpdate {
	cru.mutation.AddTenantId(i)
	return cru
}

// Mutation returns the CasbinRuleMutation object of the builder.
func (cru *CasbinRuleUpdate) Mutation() *CasbinRuleMutation {
	return cru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *CasbinRuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cru.defaults()
	if len(cru.hooks) == 0 {
		affected, err = cru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CasbinRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cru.mutation = mutation
			affected, err = cru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cru.hooks) - 1; i >= 0; i-- {
			if cru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CasbinRuleUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CasbinRuleUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CasbinRuleUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cru *CasbinRuleUpdate) defaults() {
	if _, ok := cru.mutation.UpdatedAt(); !ok {
		v := casbinrule.UpdateDefaultUpdatedAt()
		cru.mutation.SetUpdatedAt(v)
	}
}

func (cru *CasbinRuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   casbinrule.Table,
			Columns: casbinrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: casbinrule.FieldID,
			},
		},
	}
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.PType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldPType,
		})
	}
	if cru.mutation.PTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldPType,
		})
	}
	if value, ok := cru.mutation.V0(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV0,
		})
	}
	if cru.mutation.V0Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV0,
		})
	}
	if value, ok := cru.mutation.V1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV1,
		})
	}
	if cru.mutation.V1Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV1,
		})
	}
	if value, ok := cru.mutation.V2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV2,
		})
	}
	if cru.mutation.V2Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV2,
		})
	}
	if value, ok := cru.mutation.V3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV3,
		})
	}
	if cru.mutation.V3Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV3,
		})
	}
	if value, ok := cru.mutation.V4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV4,
		})
	}
	if cru.mutation.V4Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV4,
		})
	}
	if value, ok := cru.mutation.V5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV5,
		})
	}
	if cru.mutation.V5Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV5,
		})
	}
	if value, ok := cru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: casbinrule.FieldUpdatedAt,
		})
	}
	if value, ok := cru.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldCreateBy,
		})
	}
	if value, ok := cru.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldCreateBy,
		})
	}
	if value, ok := cru.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldUpdateBy,
		})
	}
	if value, ok := cru.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldUpdateBy,
		})
	}
	if value, ok := cru.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldTenantId,
		})
	}
	if value, ok := cru.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldTenantId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{casbinrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CasbinRuleUpdateOne is the builder for updating a single CasbinRule entity.
type CasbinRuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CasbinRuleMutation
}

// SetPType sets the "p_type" field.
func (cruo *CasbinRuleUpdateOne) SetPType(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetPType(s)
	return cruo
}

// SetNillablePType sets the "p_type" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillablePType(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetPType(*s)
	}
	return cruo
}

// ClearPType clears the value of the "p_type" field.
func (cruo *CasbinRuleUpdateOne) ClearPType() *CasbinRuleUpdateOne {
	cruo.mutation.ClearPType()
	return cruo
}

// SetV0 sets the "v0" field.
func (cruo *CasbinRuleUpdateOne) SetV0(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV0(s)
	return cruo
}

// SetNillableV0 sets the "v0" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV0(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV0(*s)
	}
	return cruo
}

// ClearV0 clears the value of the "v0" field.
func (cruo *CasbinRuleUpdateOne) ClearV0() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV0()
	return cruo
}

// SetV1 sets the "v1" field.
func (cruo *CasbinRuleUpdateOne) SetV1(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV1(s)
	return cruo
}

// SetNillableV1 sets the "v1" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV1(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV1(*s)
	}
	return cruo
}

// ClearV1 clears the value of the "v1" field.
func (cruo *CasbinRuleUpdateOne) ClearV1() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV1()
	return cruo
}

// SetV2 sets the "v2" field.
func (cruo *CasbinRuleUpdateOne) SetV2(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV2(s)
	return cruo
}

// SetNillableV2 sets the "v2" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV2(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV2(*s)
	}
	return cruo
}

// ClearV2 clears the value of the "v2" field.
func (cruo *CasbinRuleUpdateOne) ClearV2() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV2()
	return cruo
}

// SetV3 sets the "v3" field.
func (cruo *CasbinRuleUpdateOne) SetV3(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV3(s)
	return cruo
}

// SetNillableV3 sets the "v3" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV3(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV3(*s)
	}
	return cruo
}

// ClearV3 clears the value of the "v3" field.
func (cruo *CasbinRuleUpdateOne) ClearV3() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV3()
	return cruo
}

// SetV4 sets the "v4" field.
func (cruo *CasbinRuleUpdateOne) SetV4(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV4(s)
	return cruo
}

// SetNillableV4 sets the "v4" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV4(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV4(*s)
	}
	return cruo
}

// ClearV4 clears the value of the "v4" field.
func (cruo *CasbinRuleUpdateOne) ClearV4() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV4()
	return cruo
}

// SetV5 sets the "v5" field.
func (cruo *CasbinRuleUpdateOne) SetV5(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV5(s)
	return cruo
}

// SetNillableV5 sets the "v5" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV5(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV5(*s)
	}
	return cruo
}

// ClearV5 clears the value of the "v5" field.
func (cruo *CasbinRuleUpdateOne) ClearV5() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV5()
	return cruo
}

// SetUpdatedAt sets the "updatedAt" field.
func (cruo *CasbinRuleUpdateOne) SetUpdatedAt(t time.Time) *CasbinRuleUpdateOne {
	cruo.mutation.SetUpdatedAt(t)
	return cruo
}

// SetCreateBy sets the "createBy" field.
func (cruo *CasbinRuleUpdateOne) SetCreateBy(i int64) *CasbinRuleUpdateOne {
	cruo.mutation.ResetCreateBy()
	cruo.mutation.SetCreateBy(i)
	return cruo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableCreateBy(i *int64) *CasbinRuleUpdateOne {
	if i != nil {
		cruo.SetCreateBy(*i)
	}
	return cruo
}

// AddCreateBy adds i to the "createBy" field.
func (cruo *CasbinRuleUpdateOne) AddCreateBy(i int64) *CasbinRuleUpdateOne {
	cruo.mutation.AddCreateBy(i)
	return cruo
}

// SetUpdateBy sets the "updateBy" field.
func (cruo *CasbinRuleUpdateOne) SetUpdateBy(i int64) *CasbinRuleUpdateOne {
	cruo.mutation.ResetUpdateBy()
	cruo.mutation.SetUpdateBy(i)
	return cruo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableUpdateBy(i *int64) *CasbinRuleUpdateOne {
	if i != nil {
		cruo.SetUpdateBy(*i)
	}
	return cruo
}

// AddUpdateBy adds i to the "updateBy" field.
func (cruo *CasbinRuleUpdateOne) AddUpdateBy(i int64) *CasbinRuleUpdateOne {
	cruo.mutation.AddUpdateBy(i)
	return cruo
}

// SetTenantId sets the "tenantId" field.
func (cruo *CasbinRuleUpdateOne) SetTenantId(i int64) *CasbinRuleUpdateOne {
	cruo.mutation.ResetTenantId()
	cruo.mutation.SetTenantId(i)
	return cruo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableTenantId(i *int64) *CasbinRuleUpdateOne {
	if i != nil {
		cruo.SetTenantId(*i)
	}
	return cruo
}

// AddTenantId adds i to the "tenantId" field.
func (cruo *CasbinRuleUpdateOne) AddTenantId(i int64) *CasbinRuleUpdateOne {
	cruo.mutation.AddTenantId(i)
	return cruo
}

// Mutation returns the CasbinRuleMutation object of the builder.
func (cruo *CasbinRuleUpdateOne) Mutation() *CasbinRuleMutation {
	return cruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *CasbinRuleUpdateOne) Select(field string, fields ...string) *CasbinRuleUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated CasbinRule entity.
func (cruo *CasbinRuleUpdateOne) Save(ctx context.Context) (*CasbinRule, error) {
	var (
		err  error
		node *CasbinRule
	)
	cruo.defaults()
	if len(cruo.hooks) == 0 {
		node, err = cruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CasbinRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cruo.mutation = mutation
			node, err = cruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cruo.hooks) - 1; i >= 0; i-- {
			if cruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CasbinRuleUpdateOne) SaveX(ctx context.Context) *CasbinRule {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *CasbinRuleUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CasbinRuleUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cruo *CasbinRuleUpdateOne) defaults() {
	if _, ok := cruo.mutation.UpdatedAt(); !ok {
		v := casbinrule.UpdateDefaultUpdatedAt()
		cruo.mutation.SetUpdatedAt(v)
	}
}

func (cruo *CasbinRuleUpdateOne) sqlSave(ctx context.Context) (_node *CasbinRule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   casbinrule.Table,
			Columns: casbinrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: casbinrule.FieldID,
			},
		},
	}
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CasbinRule.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, casbinrule.FieldID)
		for _, f := range fields {
			if !casbinrule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != casbinrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.PType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldPType,
		})
	}
	if cruo.mutation.PTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldPType,
		})
	}
	if value, ok := cruo.mutation.V0(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV0,
		})
	}
	if cruo.mutation.V0Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV0,
		})
	}
	if value, ok := cruo.mutation.V1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV1,
		})
	}
	if cruo.mutation.V1Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV1,
		})
	}
	if value, ok := cruo.mutation.V2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV2,
		})
	}
	if cruo.mutation.V2Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV2,
		})
	}
	if value, ok := cruo.mutation.V3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV3,
		})
	}
	if cruo.mutation.V3Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV3,
		})
	}
	if value, ok := cruo.mutation.V4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV4,
		})
	}
	if cruo.mutation.V4Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV4,
		})
	}
	if value, ok := cruo.mutation.V5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV5,
		})
	}
	if cruo.mutation.V5Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV5,
		})
	}
	if value, ok := cruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: casbinrule.FieldUpdatedAt,
		})
	}
	if value, ok := cruo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldCreateBy,
		})
	}
	if value, ok := cruo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldCreateBy,
		})
	}
	if value, ok := cruo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldUpdateBy,
		})
	}
	if value, ok := cruo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldUpdateBy,
		})
	}
	if value, ok := cruo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldTenantId,
		})
	}
	if value, ok := cruo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: casbinrule.FieldTenantId,
		})
	}
	_node = &CasbinRule{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{casbinrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
