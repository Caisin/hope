// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelconsume"
	"hope/apps/novel/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelConsumeUpdate is the builder for updating NovelConsume entities.
type NovelConsumeUpdate struct {
	config
	hooks    []Hook
	mutation *NovelConsumeMutation
}

// Where appends a list predicates to the NovelConsumeUpdate builder.
func (ncu *NovelConsumeUpdate) Where(ps ...predicate.NovelConsume) *NovelConsumeUpdate {
	ncu.mutation.Where(ps...)
	return ncu
}

// SetNovelId sets the "novelId" field.
func (ncu *NovelConsumeUpdate) SetNovelId(i int64) *NovelConsumeUpdate {
	ncu.mutation.ResetNovelId()
	ncu.mutation.SetNovelId(i)
	return ncu
}

// AddNovelId adds i to the "novelId" field.
func (ncu *NovelConsumeUpdate) AddNovelId(i int64) *NovelConsumeUpdate {
	ncu.mutation.AddNovelId(i)
	return ncu
}

// SetCoin sets the "coin" field.
func (ncu *NovelConsumeUpdate) SetCoin(i int64) *NovelConsumeUpdate {
	ncu.mutation.ResetCoin()
	ncu.mutation.SetCoin(i)
	return ncu
}

// SetNillableCoin sets the "coin" field if the given value is not nil.
func (ncu *NovelConsumeUpdate) SetNillableCoin(i *int64) *NovelConsumeUpdate {
	if i != nil {
		ncu.SetCoin(*i)
	}
	return ncu
}

// AddCoin adds i to the "coin" field.
func (ncu *NovelConsumeUpdate) AddCoin(i int64) *NovelConsumeUpdate {
	ncu.mutation.AddCoin(i)
	return ncu
}

// ClearCoin clears the value of the "coin" field.
func (ncu *NovelConsumeUpdate) ClearCoin() *NovelConsumeUpdate {
	ncu.mutation.ClearCoin()
	return ncu
}

// SetCoupon sets the "coupon" field.
func (ncu *NovelConsumeUpdate) SetCoupon(i int64) *NovelConsumeUpdate {
	ncu.mutation.ResetCoupon()
	ncu.mutation.SetCoupon(i)
	return ncu
}

// SetNillableCoupon sets the "coupon" field if the given value is not nil.
func (ncu *NovelConsumeUpdate) SetNillableCoupon(i *int64) *NovelConsumeUpdate {
	if i != nil {
		ncu.SetCoupon(*i)
	}
	return ncu
}

// AddCoupon adds i to the "coupon" field.
func (ncu *NovelConsumeUpdate) AddCoupon(i int64) *NovelConsumeUpdate {
	ncu.mutation.AddCoupon(i)
	return ncu
}

// ClearCoupon clears the value of the "coupon" field.
func (ncu *NovelConsumeUpdate) ClearCoupon() *NovelConsumeUpdate {
	ncu.mutation.ClearCoupon()
	return ncu
}

// SetDiscount sets the "discount" field.
func (ncu *NovelConsumeUpdate) SetDiscount(i int64) *NovelConsumeUpdate {
	ncu.mutation.ResetDiscount()
	ncu.mutation.SetDiscount(i)
	return ncu
}

// SetNillableDiscount sets the "discount" field if the given value is not nil.
func (ncu *NovelConsumeUpdate) SetNillableDiscount(i *int64) *NovelConsumeUpdate {
	if i != nil {
		ncu.SetDiscount(*i)
	}
	return ncu
}

// AddDiscount adds i to the "discount" field.
func (ncu *NovelConsumeUpdate) AddDiscount(i int64) *NovelConsumeUpdate {
	ncu.mutation.AddDiscount(i)
	return ncu
}

// ClearDiscount clears the value of the "discount" field.
func (ncu *NovelConsumeUpdate) ClearDiscount() *NovelConsumeUpdate {
	ncu.mutation.ClearDiscount()
	return ncu
}

// SetReward sets the "reward" field.
func (ncu *NovelConsumeUpdate) SetReward(i int64) *NovelConsumeUpdate {
	ncu.mutation.ResetReward()
	ncu.mutation.SetReward(i)
	return ncu
}

// SetNillableReward sets the "reward" field if the given value is not nil.
func (ncu *NovelConsumeUpdate) SetNillableReward(i *int64) *NovelConsumeUpdate {
	if i != nil {
		ncu.SetReward(*i)
	}
	return ncu
}

// AddReward adds i to the "reward" field.
func (ncu *NovelConsumeUpdate) AddReward(i int64) *NovelConsumeUpdate {
	ncu.mutation.AddReward(i)
	return ncu
}

// ClearReward clears the value of the "reward" field.
func (ncu *NovelConsumeUpdate) ClearReward() *NovelConsumeUpdate {
	ncu.mutation.ClearReward()
	return ncu
}

// SetUpdatedAt sets the "updatedAt" field.
func (ncu *NovelConsumeUpdate) SetUpdatedAt(t time.Time) *NovelConsumeUpdate {
	ncu.mutation.SetUpdatedAt(t)
	return ncu
}

// SetCreateBy sets the "createBy" field.
func (ncu *NovelConsumeUpdate) SetCreateBy(i int64) *NovelConsumeUpdate {
	ncu.mutation.ResetCreateBy()
	ncu.mutation.SetCreateBy(i)
	return ncu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (ncu *NovelConsumeUpdate) SetNillableCreateBy(i *int64) *NovelConsumeUpdate {
	if i != nil {
		ncu.SetCreateBy(*i)
	}
	return ncu
}

// AddCreateBy adds i to the "createBy" field.
func (ncu *NovelConsumeUpdate) AddCreateBy(i int64) *NovelConsumeUpdate {
	ncu.mutation.AddCreateBy(i)
	return ncu
}

// SetUpdateBy sets the "updateBy" field.
func (ncu *NovelConsumeUpdate) SetUpdateBy(i int64) *NovelConsumeUpdate {
	ncu.mutation.ResetUpdateBy()
	ncu.mutation.SetUpdateBy(i)
	return ncu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (ncu *NovelConsumeUpdate) SetNillableUpdateBy(i *int64) *NovelConsumeUpdate {
	if i != nil {
		ncu.SetUpdateBy(*i)
	}
	return ncu
}

// AddUpdateBy adds i to the "updateBy" field.
func (ncu *NovelConsumeUpdate) AddUpdateBy(i int64) *NovelConsumeUpdate {
	ncu.mutation.AddUpdateBy(i)
	return ncu
}

// SetTenantId sets the "tenantId" field.
func (ncu *NovelConsumeUpdate) SetTenantId(i int64) *NovelConsumeUpdate {
	ncu.mutation.ResetTenantId()
	ncu.mutation.SetTenantId(i)
	return ncu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (ncu *NovelConsumeUpdate) SetNillableTenantId(i *int64) *NovelConsumeUpdate {
	if i != nil {
		ncu.SetTenantId(*i)
	}
	return ncu
}

// AddTenantId adds i to the "tenantId" field.
func (ncu *NovelConsumeUpdate) AddTenantId(i int64) *NovelConsumeUpdate {
	ncu.mutation.AddTenantId(i)
	return ncu
}

// Mutation returns the NovelConsumeMutation object of the builder.
func (ncu *NovelConsumeUpdate) Mutation() *NovelConsumeMutation {
	return ncu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ncu *NovelConsumeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ncu.defaults()
	if len(ncu.hooks) == 0 {
		affected, err = ncu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelConsumeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ncu.mutation = mutation
			affected, err = ncu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ncu.hooks) - 1; i >= 0; i-- {
			if ncu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ncu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ncu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ncu *NovelConsumeUpdate) SaveX(ctx context.Context) int {
	affected, err := ncu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ncu *NovelConsumeUpdate) Exec(ctx context.Context) error {
	_, err := ncu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncu *NovelConsumeUpdate) ExecX(ctx context.Context) {
	if err := ncu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ncu *NovelConsumeUpdate) defaults() {
	if _, ok := ncu.mutation.UpdatedAt(); !ok {
		v := novelconsume.UpdateDefaultUpdatedAt()
		ncu.mutation.SetUpdatedAt(v)
	}
}

func (ncu *NovelConsumeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelconsume.Table,
			Columns: novelconsume.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelconsume.FieldID,
			},
		},
	}
	if ps := ncu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncu.mutation.NovelId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldNovelId,
		})
	}
	if value, ok := ncu.mutation.AddedNovelId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldNovelId,
		})
	}
	if value, ok := ncu.mutation.Coin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCoin,
		})
	}
	if value, ok := ncu.mutation.AddedCoin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCoin,
		})
	}
	if ncu.mutation.CoinCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelconsume.FieldCoin,
		})
	}
	if value, ok := ncu.mutation.Coupon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCoupon,
		})
	}
	if value, ok := ncu.mutation.AddedCoupon(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCoupon,
		})
	}
	if ncu.mutation.CouponCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelconsume.FieldCoupon,
		})
	}
	if value, ok := ncu.mutation.Discount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldDiscount,
		})
	}
	if value, ok := ncu.mutation.AddedDiscount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldDiscount,
		})
	}
	if ncu.mutation.DiscountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelconsume.FieldDiscount,
		})
	}
	if value, ok := ncu.mutation.Reward(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldReward,
		})
	}
	if value, ok := ncu.mutation.AddedReward(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldReward,
		})
	}
	if ncu.mutation.RewardCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelconsume.FieldReward,
		})
	}
	if value, ok := ncu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelconsume.FieldUpdatedAt,
		})
	}
	if value, ok := ncu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCreateBy,
		})
	}
	if value, ok := ncu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCreateBy,
		})
	}
	if value, ok := ncu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldUpdateBy,
		})
	}
	if value, ok := ncu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldUpdateBy,
		})
	}
	if value, ok := ncu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldTenantId,
		})
	}
	if value, ok := ncu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldTenantId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ncu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{novelconsume.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NovelConsumeUpdateOne is the builder for updating a single NovelConsume entity.
type NovelConsumeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NovelConsumeMutation
}

// SetNovelId sets the "novelId" field.
func (ncuo *NovelConsumeUpdateOne) SetNovelId(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.ResetNovelId()
	ncuo.mutation.SetNovelId(i)
	return ncuo
}

// AddNovelId adds i to the "novelId" field.
func (ncuo *NovelConsumeUpdateOne) AddNovelId(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.AddNovelId(i)
	return ncuo
}

// SetCoin sets the "coin" field.
func (ncuo *NovelConsumeUpdateOne) SetCoin(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.ResetCoin()
	ncuo.mutation.SetCoin(i)
	return ncuo
}

// SetNillableCoin sets the "coin" field if the given value is not nil.
func (ncuo *NovelConsumeUpdateOne) SetNillableCoin(i *int64) *NovelConsumeUpdateOne {
	if i != nil {
		ncuo.SetCoin(*i)
	}
	return ncuo
}

// AddCoin adds i to the "coin" field.
func (ncuo *NovelConsumeUpdateOne) AddCoin(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.AddCoin(i)
	return ncuo
}

// ClearCoin clears the value of the "coin" field.
func (ncuo *NovelConsumeUpdateOne) ClearCoin() *NovelConsumeUpdateOne {
	ncuo.mutation.ClearCoin()
	return ncuo
}

// SetCoupon sets the "coupon" field.
func (ncuo *NovelConsumeUpdateOne) SetCoupon(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.ResetCoupon()
	ncuo.mutation.SetCoupon(i)
	return ncuo
}

// SetNillableCoupon sets the "coupon" field if the given value is not nil.
func (ncuo *NovelConsumeUpdateOne) SetNillableCoupon(i *int64) *NovelConsumeUpdateOne {
	if i != nil {
		ncuo.SetCoupon(*i)
	}
	return ncuo
}

// AddCoupon adds i to the "coupon" field.
func (ncuo *NovelConsumeUpdateOne) AddCoupon(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.AddCoupon(i)
	return ncuo
}

// ClearCoupon clears the value of the "coupon" field.
func (ncuo *NovelConsumeUpdateOne) ClearCoupon() *NovelConsumeUpdateOne {
	ncuo.mutation.ClearCoupon()
	return ncuo
}

// SetDiscount sets the "discount" field.
func (ncuo *NovelConsumeUpdateOne) SetDiscount(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.ResetDiscount()
	ncuo.mutation.SetDiscount(i)
	return ncuo
}

// SetNillableDiscount sets the "discount" field if the given value is not nil.
func (ncuo *NovelConsumeUpdateOne) SetNillableDiscount(i *int64) *NovelConsumeUpdateOne {
	if i != nil {
		ncuo.SetDiscount(*i)
	}
	return ncuo
}

// AddDiscount adds i to the "discount" field.
func (ncuo *NovelConsumeUpdateOne) AddDiscount(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.AddDiscount(i)
	return ncuo
}

// ClearDiscount clears the value of the "discount" field.
func (ncuo *NovelConsumeUpdateOne) ClearDiscount() *NovelConsumeUpdateOne {
	ncuo.mutation.ClearDiscount()
	return ncuo
}

// SetReward sets the "reward" field.
func (ncuo *NovelConsumeUpdateOne) SetReward(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.ResetReward()
	ncuo.mutation.SetReward(i)
	return ncuo
}

// SetNillableReward sets the "reward" field if the given value is not nil.
func (ncuo *NovelConsumeUpdateOne) SetNillableReward(i *int64) *NovelConsumeUpdateOne {
	if i != nil {
		ncuo.SetReward(*i)
	}
	return ncuo
}

// AddReward adds i to the "reward" field.
func (ncuo *NovelConsumeUpdateOne) AddReward(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.AddReward(i)
	return ncuo
}

// ClearReward clears the value of the "reward" field.
func (ncuo *NovelConsumeUpdateOne) ClearReward() *NovelConsumeUpdateOne {
	ncuo.mutation.ClearReward()
	return ncuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (ncuo *NovelConsumeUpdateOne) SetUpdatedAt(t time.Time) *NovelConsumeUpdateOne {
	ncuo.mutation.SetUpdatedAt(t)
	return ncuo
}

// SetCreateBy sets the "createBy" field.
func (ncuo *NovelConsumeUpdateOne) SetCreateBy(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.ResetCreateBy()
	ncuo.mutation.SetCreateBy(i)
	return ncuo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (ncuo *NovelConsumeUpdateOne) SetNillableCreateBy(i *int64) *NovelConsumeUpdateOne {
	if i != nil {
		ncuo.SetCreateBy(*i)
	}
	return ncuo
}

// AddCreateBy adds i to the "createBy" field.
func (ncuo *NovelConsumeUpdateOne) AddCreateBy(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.AddCreateBy(i)
	return ncuo
}

// SetUpdateBy sets the "updateBy" field.
func (ncuo *NovelConsumeUpdateOne) SetUpdateBy(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.ResetUpdateBy()
	ncuo.mutation.SetUpdateBy(i)
	return ncuo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (ncuo *NovelConsumeUpdateOne) SetNillableUpdateBy(i *int64) *NovelConsumeUpdateOne {
	if i != nil {
		ncuo.SetUpdateBy(*i)
	}
	return ncuo
}

// AddUpdateBy adds i to the "updateBy" field.
func (ncuo *NovelConsumeUpdateOne) AddUpdateBy(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.AddUpdateBy(i)
	return ncuo
}

// SetTenantId sets the "tenantId" field.
func (ncuo *NovelConsumeUpdateOne) SetTenantId(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.ResetTenantId()
	ncuo.mutation.SetTenantId(i)
	return ncuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (ncuo *NovelConsumeUpdateOne) SetNillableTenantId(i *int64) *NovelConsumeUpdateOne {
	if i != nil {
		ncuo.SetTenantId(*i)
	}
	return ncuo
}

// AddTenantId adds i to the "tenantId" field.
func (ncuo *NovelConsumeUpdateOne) AddTenantId(i int64) *NovelConsumeUpdateOne {
	ncuo.mutation.AddTenantId(i)
	return ncuo
}

// Mutation returns the NovelConsumeMutation object of the builder.
func (ncuo *NovelConsumeUpdateOne) Mutation() *NovelConsumeMutation {
	return ncuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ncuo *NovelConsumeUpdateOne) Select(field string, fields ...string) *NovelConsumeUpdateOne {
	ncuo.fields = append([]string{field}, fields...)
	return ncuo
}

// Save executes the query and returns the updated NovelConsume entity.
func (ncuo *NovelConsumeUpdateOne) Save(ctx context.Context) (*NovelConsume, error) {
	var (
		err  error
		node *NovelConsume
	)
	ncuo.defaults()
	if len(ncuo.hooks) == 0 {
		node, err = ncuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelConsumeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ncuo.mutation = mutation
			node, err = ncuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ncuo.hooks) - 1; i >= 0; i-- {
			if ncuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ncuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ncuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ncuo *NovelConsumeUpdateOne) SaveX(ctx context.Context) *NovelConsume {
	node, err := ncuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ncuo *NovelConsumeUpdateOne) Exec(ctx context.Context) error {
	_, err := ncuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncuo *NovelConsumeUpdateOne) ExecX(ctx context.Context) {
	if err := ncuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ncuo *NovelConsumeUpdateOne) defaults() {
	if _, ok := ncuo.mutation.UpdatedAt(); !ok {
		v := novelconsume.UpdateDefaultUpdatedAt()
		ncuo.mutation.SetUpdatedAt(v)
	}
}

func (ncuo *NovelConsumeUpdateOne) sqlSave(ctx context.Context) (_node *NovelConsume, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelconsume.Table,
			Columns: novelconsume.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelconsume.FieldID,
			},
		},
	}
	id, ok := ncuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing NovelConsume.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ncuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novelconsume.FieldID)
		for _, f := range fields {
			if !novelconsume.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != novelconsume.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ncuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncuo.mutation.NovelId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldNovelId,
		})
	}
	if value, ok := ncuo.mutation.AddedNovelId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldNovelId,
		})
	}
	if value, ok := ncuo.mutation.Coin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCoin,
		})
	}
	if value, ok := ncuo.mutation.AddedCoin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCoin,
		})
	}
	if ncuo.mutation.CoinCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelconsume.FieldCoin,
		})
	}
	if value, ok := ncuo.mutation.Coupon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCoupon,
		})
	}
	if value, ok := ncuo.mutation.AddedCoupon(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCoupon,
		})
	}
	if ncuo.mutation.CouponCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelconsume.FieldCoupon,
		})
	}
	if value, ok := ncuo.mutation.Discount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldDiscount,
		})
	}
	if value, ok := ncuo.mutation.AddedDiscount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldDiscount,
		})
	}
	if ncuo.mutation.DiscountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelconsume.FieldDiscount,
		})
	}
	if value, ok := ncuo.mutation.Reward(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldReward,
		})
	}
	if value, ok := ncuo.mutation.AddedReward(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldReward,
		})
	}
	if ncuo.mutation.RewardCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelconsume.FieldReward,
		})
	}
	if value, ok := ncuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelconsume.FieldUpdatedAt,
		})
	}
	if value, ok := ncuo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCreateBy,
		})
	}
	if value, ok := ncuo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldCreateBy,
		})
	}
	if value, ok := ncuo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldUpdateBy,
		})
	}
	if value, ok := ncuo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldUpdateBy,
		})
	}
	if value, ok := ncuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldTenantId,
		})
	}
	if value, ok := ncuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelconsume.FieldTenantId,
		})
	}
	_node = &NovelConsume{config: ncuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ncuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{novelconsume.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}