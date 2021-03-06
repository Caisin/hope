// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/assetitem"
	"hope/apps/novel/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AssetItemUpdate is the builder for updating AssetItem entities.
type AssetItemUpdate struct {
	config
	hooks    []Hook
	mutation *AssetItemMutation
}

// Where appends a list predicates to the AssetItemUpdate builder.
func (aiu *AssetItemUpdate) Where(ps ...predicate.AssetItem) *AssetItemUpdate {
	aiu.mutation.Where(ps...)
	return aiu
}

// SetAssetItemId sets the "assetItemId" field.
func (aiu *AssetItemUpdate) SetAssetItemId(i int32) *AssetItemUpdate {
	aiu.mutation.ResetAssetItemId()
	aiu.mutation.SetAssetItemId(i)
	return aiu
}

// AddAssetItemId adds i to the "assetItemId" field.
func (aiu *AssetItemUpdate) AddAssetItemId(i int32) *AssetItemUpdate {
	aiu.mutation.AddAssetItemId(i)
	return aiu
}

// SetAssetName sets the "assetName" field.
func (aiu *AssetItemUpdate) SetAssetName(s string) *AssetItemUpdate {
	aiu.mutation.SetAssetName(s)
	return aiu
}

// SetNillableAssetName sets the "assetName" field if the given value is not nil.
func (aiu *AssetItemUpdate) SetNillableAssetName(s *string) *AssetItemUpdate {
	if s != nil {
		aiu.SetAssetName(*s)
	}
	return aiu
}

// ClearAssetName clears the value of the "assetName" field.
func (aiu *AssetItemUpdate) ClearAssetName() *AssetItemUpdate {
	aiu.mutation.ClearAssetName()
	return aiu
}

// SetCashTag sets the "cashTag" field.
func (aiu *AssetItemUpdate) SetCashTag(i int32) *AssetItemUpdate {
	aiu.mutation.ResetCashTag()
	aiu.mutation.SetCashTag(i)
	return aiu
}

// SetNillableCashTag sets the "cashTag" field if the given value is not nil.
func (aiu *AssetItemUpdate) SetNillableCashTag(i *int32) *AssetItemUpdate {
	if i != nil {
		aiu.SetCashTag(*i)
	}
	return aiu
}

// AddCashTag adds i to the "cashTag" field.
func (aiu *AssetItemUpdate) AddCashTag(i int32) *AssetItemUpdate {
	aiu.mutation.AddCashTag(i)
	return aiu
}

// ClearCashTag clears the value of the "cashTag" field.
func (aiu *AssetItemUpdate) ClearCashTag() *AssetItemUpdate {
	aiu.mutation.ClearCashTag()
	return aiu
}

// SetValidDays sets the "validDays" field.
func (aiu *AssetItemUpdate) SetValidDays(i int32) *AssetItemUpdate {
	aiu.mutation.ResetValidDays()
	aiu.mutation.SetValidDays(i)
	return aiu
}

// SetNillableValidDays sets the "validDays" field if the given value is not nil.
func (aiu *AssetItemUpdate) SetNillableValidDays(i *int32) *AssetItemUpdate {
	if i != nil {
		aiu.SetValidDays(*i)
	}
	return aiu
}

// AddValidDays adds i to the "validDays" field.
func (aiu *AssetItemUpdate) AddValidDays(i int32) *AssetItemUpdate {
	aiu.mutation.AddValidDays(i)
	return aiu
}

// ClearValidDays clears the value of the "validDays" field.
func (aiu *AssetItemUpdate) ClearValidDays() *AssetItemUpdate {
	aiu.mutation.ClearValidDays()
	return aiu
}

// SetEffectTime sets the "effectTime" field.
func (aiu *AssetItemUpdate) SetEffectTime(t time.Time) *AssetItemUpdate {
	aiu.mutation.SetEffectTime(t)
	return aiu
}

// SetNillableEffectTime sets the "effectTime" field if the given value is not nil.
func (aiu *AssetItemUpdate) SetNillableEffectTime(t *time.Time) *AssetItemUpdate {
	if t != nil {
		aiu.SetEffectTime(*t)
	}
	return aiu
}

// SetExpiredTime sets the "expiredTime" field.
func (aiu *AssetItemUpdate) SetExpiredTime(t time.Time) *AssetItemUpdate {
	aiu.mutation.SetExpiredTime(t)
	return aiu
}

// SetNillableExpiredTime sets the "expiredTime" field if the given value is not nil.
func (aiu *AssetItemUpdate) SetNillableExpiredTime(t *time.Time) *AssetItemUpdate {
	if t != nil {
		aiu.SetExpiredTime(*t)
	}
	return aiu
}

// SetUpdatedAt sets the "updatedAt" field.
func (aiu *AssetItemUpdate) SetUpdatedAt(t time.Time) *AssetItemUpdate {
	aiu.mutation.SetUpdatedAt(t)
	return aiu
}

// SetCreateBy sets the "createBy" field.
func (aiu *AssetItemUpdate) SetCreateBy(i int64) *AssetItemUpdate {
	aiu.mutation.ResetCreateBy()
	aiu.mutation.SetCreateBy(i)
	return aiu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (aiu *AssetItemUpdate) SetNillableCreateBy(i *int64) *AssetItemUpdate {
	if i != nil {
		aiu.SetCreateBy(*i)
	}
	return aiu
}

// AddCreateBy adds i to the "createBy" field.
func (aiu *AssetItemUpdate) AddCreateBy(i int64) *AssetItemUpdate {
	aiu.mutation.AddCreateBy(i)
	return aiu
}

// SetUpdateBy sets the "updateBy" field.
func (aiu *AssetItemUpdate) SetUpdateBy(i int64) *AssetItemUpdate {
	aiu.mutation.ResetUpdateBy()
	aiu.mutation.SetUpdateBy(i)
	return aiu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (aiu *AssetItemUpdate) SetNillableUpdateBy(i *int64) *AssetItemUpdate {
	if i != nil {
		aiu.SetUpdateBy(*i)
	}
	return aiu
}

// AddUpdateBy adds i to the "updateBy" field.
func (aiu *AssetItemUpdate) AddUpdateBy(i int64) *AssetItemUpdate {
	aiu.mutation.AddUpdateBy(i)
	return aiu
}

// SetTenantId sets the "tenantId" field.
func (aiu *AssetItemUpdate) SetTenantId(i int64) *AssetItemUpdate {
	aiu.mutation.ResetTenantId()
	aiu.mutation.SetTenantId(i)
	return aiu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (aiu *AssetItemUpdate) SetNillableTenantId(i *int64) *AssetItemUpdate {
	if i != nil {
		aiu.SetTenantId(*i)
	}
	return aiu
}

// AddTenantId adds i to the "tenantId" field.
func (aiu *AssetItemUpdate) AddTenantId(i int64) *AssetItemUpdate {
	aiu.mutation.AddTenantId(i)
	return aiu
}

// Mutation returns the AssetItemMutation object of the builder.
func (aiu *AssetItemUpdate) Mutation() *AssetItemMutation {
	return aiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aiu *AssetItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aiu.defaults()
	if len(aiu.hooks) == 0 {
		affected, err = aiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aiu.mutation = mutation
			affected, err = aiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aiu.hooks) - 1; i >= 0; i-- {
			if aiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aiu *AssetItemUpdate) SaveX(ctx context.Context) int {
	affected, err := aiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aiu *AssetItemUpdate) Exec(ctx context.Context) error {
	_, err := aiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiu *AssetItemUpdate) ExecX(ctx context.Context) {
	if err := aiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aiu *AssetItemUpdate) defaults() {
	if _, ok := aiu.mutation.UpdatedAt(); !ok {
		v := assetitem.UpdateDefaultUpdatedAt()
		aiu.mutation.SetUpdatedAt(v)
	}
}

func (aiu *AssetItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   assetitem.Table,
			Columns: assetitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: assetitem.FieldID,
			},
		},
	}
	if ps := aiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiu.mutation.AssetItemId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldAssetItemId,
		})
	}
	if value, ok := aiu.mutation.AddedAssetItemId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldAssetItemId,
		})
	}
	if value, ok := aiu.mutation.AssetName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assetitem.FieldAssetName,
		})
	}
	if aiu.mutation.AssetNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: assetitem.FieldAssetName,
		})
	}
	if value, ok := aiu.mutation.CashTag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldCashTag,
		})
	}
	if value, ok := aiu.mutation.AddedCashTag(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldCashTag,
		})
	}
	if aiu.mutation.CashTagCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: assetitem.FieldCashTag,
		})
	}
	if value, ok := aiu.mutation.ValidDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldValidDays,
		})
	}
	if value, ok := aiu.mutation.AddedValidDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldValidDays,
		})
	}
	if aiu.mutation.ValidDaysCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: assetitem.FieldValidDays,
		})
	}
	if value, ok := aiu.mutation.EffectTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetitem.FieldEffectTime,
		})
	}
	if value, ok := aiu.mutation.ExpiredTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetitem.FieldExpiredTime,
		})
	}
	if value, ok := aiu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetitem.FieldUpdatedAt,
		})
	}
	if value, ok := aiu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldCreateBy,
		})
	}
	if value, ok := aiu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldCreateBy,
		})
	}
	if value, ok := aiu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldUpdateBy,
		})
	}
	if value, ok := aiu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldUpdateBy,
		})
	}
	if value, ok := aiu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldTenantId,
		})
	}
	if value, ok := aiu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldTenantId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assetitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AssetItemUpdateOne is the builder for updating a single AssetItem entity.
type AssetItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AssetItemMutation
}

// SetAssetItemId sets the "assetItemId" field.
func (aiuo *AssetItemUpdateOne) SetAssetItemId(i int32) *AssetItemUpdateOne {
	aiuo.mutation.ResetAssetItemId()
	aiuo.mutation.SetAssetItemId(i)
	return aiuo
}

// AddAssetItemId adds i to the "assetItemId" field.
func (aiuo *AssetItemUpdateOne) AddAssetItemId(i int32) *AssetItemUpdateOne {
	aiuo.mutation.AddAssetItemId(i)
	return aiuo
}

// SetAssetName sets the "assetName" field.
func (aiuo *AssetItemUpdateOne) SetAssetName(s string) *AssetItemUpdateOne {
	aiuo.mutation.SetAssetName(s)
	return aiuo
}

// SetNillableAssetName sets the "assetName" field if the given value is not nil.
func (aiuo *AssetItemUpdateOne) SetNillableAssetName(s *string) *AssetItemUpdateOne {
	if s != nil {
		aiuo.SetAssetName(*s)
	}
	return aiuo
}

// ClearAssetName clears the value of the "assetName" field.
func (aiuo *AssetItemUpdateOne) ClearAssetName() *AssetItemUpdateOne {
	aiuo.mutation.ClearAssetName()
	return aiuo
}

// SetCashTag sets the "cashTag" field.
func (aiuo *AssetItemUpdateOne) SetCashTag(i int32) *AssetItemUpdateOne {
	aiuo.mutation.ResetCashTag()
	aiuo.mutation.SetCashTag(i)
	return aiuo
}

// SetNillableCashTag sets the "cashTag" field if the given value is not nil.
func (aiuo *AssetItemUpdateOne) SetNillableCashTag(i *int32) *AssetItemUpdateOne {
	if i != nil {
		aiuo.SetCashTag(*i)
	}
	return aiuo
}

// AddCashTag adds i to the "cashTag" field.
func (aiuo *AssetItemUpdateOne) AddCashTag(i int32) *AssetItemUpdateOne {
	aiuo.mutation.AddCashTag(i)
	return aiuo
}

// ClearCashTag clears the value of the "cashTag" field.
func (aiuo *AssetItemUpdateOne) ClearCashTag() *AssetItemUpdateOne {
	aiuo.mutation.ClearCashTag()
	return aiuo
}

// SetValidDays sets the "validDays" field.
func (aiuo *AssetItemUpdateOne) SetValidDays(i int32) *AssetItemUpdateOne {
	aiuo.mutation.ResetValidDays()
	aiuo.mutation.SetValidDays(i)
	return aiuo
}

// SetNillableValidDays sets the "validDays" field if the given value is not nil.
func (aiuo *AssetItemUpdateOne) SetNillableValidDays(i *int32) *AssetItemUpdateOne {
	if i != nil {
		aiuo.SetValidDays(*i)
	}
	return aiuo
}

// AddValidDays adds i to the "validDays" field.
func (aiuo *AssetItemUpdateOne) AddValidDays(i int32) *AssetItemUpdateOne {
	aiuo.mutation.AddValidDays(i)
	return aiuo
}

// ClearValidDays clears the value of the "validDays" field.
func (aiuo *AssetItemUpdateOne) ClearValidDays() *AssetItemUpdateOne {
	aiuo.mutation.ClearValidDays()
	return aiuo
}

// SetEffectTime sets the "effectTime" field.
func (aiuo *AssetItemUpdateOne) SetEffectTime(t time.Time) *AssetItemUpdateOne {
	aiuo.mutation.SetEffectTime(t)
	return aiuo
}

// SetNillableEffectTime sets the "effectTime" field if the given value is not nil.
func (aiuo *AssetItemUpdateOne) SetNillableEffectTime(t *time.Time) *AssetItemUpdateOne {
	if t != nil {
		aiuo.SetEffectTime(*t)
	}
	return aiuo
}

// SetExpiredTime sets the "expiredTime" field.
func (aiuo *AssetItemUpdateOne) SetExpiredTime(t time.Time) *AssetItemUpdateOne {
	aiuo.mutation.SetExpiredTime(t)
	return aiuo
}

// SetNillableExpiredTime sets the "expiredTime" field if the given value is not nil.
func (aiuo *AssetItemUpdateOne) SetNillableExpiredTime(t *time.Time) *AssetItemUpdateOne {
	if t != nil {
		aiuo.SetExpiredTime(*t)
	}
	return aiuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (aiuo *AssetItemUpdateOne) SetUpdatedAt(t time.Time) *AssetItemUpdateOne {
	aiuo.mutation.SetUpdatedAt(t)
	return aiuo
}

// SetCreateBy sets the "createBy" field.
func (aiuo *AssetItemUpdateOne) SetCreateBy(i int64) *AssetItemUpdateOne {
	aiuo.mutation.ResetCreateBy()
	aiuo.mutation.SetCreateBy(i)
	return aiuo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (aiuo *AssetItemUpdateOne) SetNillableCreateBy(i *int64) *AssetItemUpdateOne {
	if i != nil {
		aiuo.SetCreateBy(*i)
	}
	return aiuo
}

// AddCreateBy adds i to the "createBy" field.
func (aiuo *AssetItemUpdateOne) AddCreateBy(i int64) *AssetItemUpdateOne {
	aiuo.mutation.AddCreateBy(i)
	return aiuo
}

// SetUpdateBy sets the "updateBy" field.
func (aiuo *AssetItemUpdateOne) SetUpdateBy(i int64) *AssetItemUpdateOne {
	aiuo.mutation.ResetUpdateBy()
	aiuo.mutation.SetUpdateBy(i)
	return aiuo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (aiuo *AssetItemUpdateOne) SetNillableUpdateBy(i *int64) *AssetItemUpdateOne {
	if i != nil {
		aiuo.SetUpdateBy(*i)
	}
	return aiuo
}

// AddUpdateBy adds i to the "updateBy" field.
func (aiuo *AssetItemUpdateOne) AddUpdateBy(i int64) *AssetItemUpdateOne {
	aiuo.mutation.AddUpdateBy(i)
	return aiuo
}

// SetTenantId sets the "tenantId" field.
func (aiuo *AssetItemUpdateOne) SetTenantId(i int64) *AssetItemUpdateOne {
	aiuo.mutation.ResetTenantId()
	aiuo.mutation.SetTenantId(i)
	return aiuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (aiuo *AssetItemUpdateOne) SetNillableTenantId(i *int64) *AssetItemUpdateOne {
	if i != nil {
		aiuo.SetTenantId(*i)
	}
	return aiuo
}

// AddTenantId adds i to the "tenantId" field.
func (aiuo *AssetItemUpdateOne) AddTenantId(i int64) *AssetItemUpdateOne {
	aiuo.mutation.AddTenantId(i)
	return aiuo
}

// Mutation returns the AssetItemMutation object of the builder.
func (aiuo *AssetItemUpdateOne) Mutation() *AssetItemMutation {
	return aiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aiuo *AssetItemUpdateOne) Select(field string, fields ...string) *AssetItemUpdateOne {
	aiuo.fields = append([]string{field}, fields...)
	return aiuo
}

// Save executes the query and returns the updated AssetItem entity.
func (aiuo *AssetItemUpdateOne) Save(ctx context.Context) (*AssetItem, error) {
	var (
		err  error
		node *AssetItem
	)
	aiuo.defaults()
	if len(aiuo.hooks) == 0 {
		node, err = aiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aiuo.mutation = mutation
			node, err = aiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aiuo.hooks) - 1; i >= 0; i-- {
			if aiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aiuo *AssetItemUpdateOne) SaveX(ctx context.Context) *AssetItem {
	node, err := aiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aiuo *AssetItemUpdateOne) Exec(ctx context.Context) error {
	_, err := aiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiuo *AssetItemUpdateOne) ExecX(ctx context.Context) {
	if err := aiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aiuo *AssetItemUpdateOne) defaults() {
	if _, ok := aiuo.mutation.UpdatedAt(); !ok {
		v := assetitem.UpdateDefaultUpdatedAt()
		aiuo.mutation.SetUpdatedAt(v)
	}
}

func (aiuo *AssetItemUpdateOne) sqlSave(ctx context.Context) (_node *AssetItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   assetitem.Table,
			Columns: assetitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: assetitem.FieldID,
			},
		},
	}
	id, ok := aiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AssetItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, assetitem.FieldID)
		for _, f := range fields {
			if !assetitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != assetitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiuo.mutation.AssetItemId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldAssetItemId,
		})
	}
	if value, ok := aiuo.mutation.AddedAssetItemId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldAssetItemId,
		})
	}
	if value, ok := aiuo.mutation.AssetName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assetitem.FieldAssetName,
		})
	}
	if aiuo.mutation.AssetNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: assetitem.FieldAssetName,
		})
	}
	if value, ok := aiuo.mutation.CashTag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldCashTag,
		})
	}
	if value, ok := aiuo.mutation.AddedCashTag(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldCashTag,
		})
	}
	if aiuo.mutation.CashTagCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: assetitem.FieldCashTag,
		})
	}
	if value, ok := aiuo.mutation.ValidDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldValidDays,
		})
	}
	if value, ok := aiuo.mutation.AddedValidDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: assetitem.FieldValidDays,
		})
	}
	if aiuo.mutation.ValidDaysCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: assetitem.FieldValidDays,
		})
	}
	if value, ok := aiuo.mutation.EffectTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetitem.FieldEffectTime,
		})
	}
	if value, ok := aiuo.mutation.ExpiredTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetitem.FieldExpiredTime,
		})
	}
	if value, ok := aiuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: assetitem.FieldUpdatedAt,
		})
	}
	if value, ok := aiuo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldCreateBy,
		})
	}
	if value, ok := aiuo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldCreateBy,
		})
	}
	if value, ok := aiuo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldUpdateBy,
		})
	}
	if value, ok := aiuo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldUpdateBy,
		})
	}
	if value, ok := aiuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldTenantId,
		})
	}
	if value, ok := aiuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: assetitem.FieldTenantId,
		})
	}
	_node = &AssetItem{config: aiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assetitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
