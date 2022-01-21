// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/viptype"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VipTypeUpdate is the builder for updating VipType entities.
type VipTypeUpdate struct {
	config
	hooks    []Hook
	mutation *VipTypeMutation
}

// Where appends a list predicates to the VipTypeUpdate builder.
func (vtu *VipTypeUpdate) Where(ps ...predicate.VipType) *VipTypeUpdate {
	vtu.mutation.Where(ps...)
	return vtu
}

// SetVipName sets the "vipName" field.
func (vtu *VipTypeUpdate) SetVipName(s string) *VipTypeUpdate {
	vtu.mutation.SetVipName(s)
	return vtu
}

// SetNillableVipName sets the "vipName" field if the given value is not nil.
func (vtu *VipTypeUpdate) SetNillableVipName(s *string) *VipTypeUpdate {
	if s != nil {
		vtu.SetVipName(*s)
	}
	return vtu
}

// ClearVipName clears the value of the "vipName" field.
func (vtu *VipTypeUpdate) ClearVipName() *VipTypeUpdate {
	vtu.mutation.ClearVipName()
	return vtu
}

// SetIsSuper sets the "isSuper" field.
func (vtu *VipTypeUpdate) SetIsSuper(b bool) *VipTypeUpdate {
	vtu.mutation.SetIsSuper(b)
	return vtu
}

// SetNillableIsSuper sets the "isSuper" field if the given value is not nil.
func (vtu *VipTypeUpdate) SetNillableIsSuper(b *bool) *VipTypeUpdate {
	if b != nil {
		vtu.SetIsSuper(*b)
	}
	return vtu
}

// ClearIsSuper clears the value of the "isSuper" field.
func (vtu *VipTypeUpdate) ClearIsSuper() *VipTypeUpdate {
	vtu.mutation.ClearIsSuper()
	return vtu
}

// SetValidDays sets the "validDays" field.
func (vtu *VipTypeUpdate) SetValidDays(i int32) *VipTypeUpdate {
	vtu.mutation.ResetValidDays()
	vtu.mutation.SetValidDays(i)
	return vtu
}

// SetNillableValidDays sets the "validDays" field if the given value is not nil.
func (vtu *VipTypeUpdate) SetNillableValidDays(i *int32) *VipTypeUpdate {
	if i != nil {
		vtu.SetValidDays(*i)
	}
	return vtu
}

// AddValidDays adds i to the "validDays" field.
func (vtu *VipTypeUpdate) AddValidDays(i int32) *VipTypeUpdate {
	vtu.mutation.AddValidDays(i)
	return vtu
}

// ClearValidDays clears the value of the "validDays" field.
func (vtu *VipTypeUpdate) ClearValidDays() *VipTypeUpdate {
	vtu.mutation.ClearValidDays()
	return vtu
}

// SetDiscountRate sets the "discountRate" field.
func (vtu *VipTypeUpdate) SetDiscountRate(i int64) *VipTypeUpdate {
	vtu.mutation.ResetDiscountRate()
	vtu.mutation.SetDiscountRate(i)
	return vtu
}

// SetNillableDiscountRate sets the "discountRate" field if the given value is not nil.
func (vtu *VipTypeUpdate) SetNillableDiscountRate(i *int64) *VipTypeUpdate {
	if i != nil {
		vtu.SetDiscountRate(*i)
	}
	return vtu
}

// AddDiscountRate adds i to the "discountRate" field.
func (vtu *VipTypeUpdate) AddDiscountRate(i int64) *VipTypeUpdate {
	vtu.mutation.AddDiscountRate(i)
	return vtu
}

// ClearDiscountRate clears the value of the "discountRate" field.
func (vtu *VipTypeUpdate) ClearDiscountRate() *VipTypeUpdate {
	vtu.mutation.ClearDiscountRate()
	return vtu
}

// SetAvatarId sets the "avatarId" field.
func (vtu *VipTypeUpdate) SetAvatarId(i int64) *VipTypeUpdate {
	vtu.mutation.ResetAvatarId()
	vtu.mutation.SetAvatarId(i)
	return vtu
}

// SetNillableAvatarId sets the "avatarId" field if the given value is not nil.
func (vtu *VipTypeUpdate) SetNillableAvatarId(i *int64) *VipTypeUpdate {
	if i != nil {
		vtu.SetAvatarId(*i)
	}
	return vtu
}

// AddAvatarId adds i to the "avatarId" field.
func (vtu *VipTypeUpdate) AddAvatarId(i int64) *VipTypeUpdate {
	vtu.mutation.AddAvatarId(i)
	return vtu
}

// ClearAvatarId clears the value of the "avatarId" field.
func (vtu *VipTypeUpdate) ClearAvatarId() *VipTypeUpdate {
	vtu.mutation.ClearAvatarId()
	return vtu
}

// SetSummary sets the "summary" field.
func (vtu *VipTypeUpdate) SetSummary(s string) *VipTypeUpdate {
	vtu.mutation.SetSummary(s)
	return vtu
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (vtu *VipTypeUpdate) SetNillableSummary(s *string) *VipTypeUpdate {
	if s != nil {
		vtu.SetSummary(*s)
	}
	return vtu
}

// ClearSummary clears the value of the "summary" field.
func (vtu *VipTypeUpdate) ClearSummary() *VipTypeUpdate {
	vtu.mutation.ClearSummary()
	return vtu
}

// SetUpdatedAt sets the "updatedAt" field.
func (vtu *VipTypeUpdate) SetUpdatedAt(t time.Time) *VipTypeUpdate {
	vtu.mutation.SetUpdatedAt(t)
	return vtu
}

// SetCreateBy sets the "createBy" field.
func (vtu *VipTypeUpdate) SetCreateBy(i int64) *VipTypeUpdate {
	vtu.mutation.ResetCreateBy()
	vtu.mutation.SetCreateBy(i)
	return vtu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (vtu *VipTypeUpdate) SetNillableCreateBy(i *int64) *VipTypeUpdate {
	if i != nil {
		vtu.SetCreateBy(*i)
	}
	return vtu
}

// AddCreateBy adds i to the "createBy" field.
func (vtu *VipTypeUpdate) AddCreateBy(i int64) *VipTypeUpdate {
	vtu.mutation.AddCreateBy(i)
	return vtu
}

// SetUpdateBy sets the "updateBy" field.
func (vtu *VipTypeUpdate) SetUpdateBy(i int64) *VipTypeUpdate {
	vtu.mutation.ResetUpdateBy()
	vtu.mutation.SetUpdateBy(i)
	return vtu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (vtu *VipTypeUpdate) SetNillableUpdateBy(i *int64) *VipTypeUpdate {
	if i != nil {
		vtu.SetUpdateBy(*i)
	}
	return vtu
}

// AddUpdateBy adds i to the "updateBy" field.
func (vtu *VipTypeUpdate) AddUpdateBy(i int64) *VipTypeUpdate {
	vtu.mutation.AddUpdateBy(i)
	return vtu
}

// SetTenantId sets the "tenantId" field.
func (vtu *VipTypeUpdate) SetTenantId(i int64) *VipTypeUpdate {
	vtu.mutation.ResetTenantId()
	vtu.mutation.SetTenantId(i)
	return vtu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (vtu *VipTypeUpdate) SetNillableTenantId(i *int64) *VipTypeUpdate {
	if i != nil {
		vtu.SetTenantId(*i)
	}
	return vtu
}

// AddTenantId adds i to the "tenantId" field.
func (vtu *VipTypeUpdate) AddTenantId(i int64) *VipTypeUpdate {
	vtu.mutation.AddTenantId(i)
	return vtu
}

// Mutation returns the VipTypeMutation object of the builder.
func (vtu *VipTypeUpdate) Mutation() *VipTypeMutation {
	return vtu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vtu *VipTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	vtu.defaults()
	if len(vtu.hooks) == 0 {
		affected, err = vtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VipTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vtu.mutation = mutation
			affected, err = vtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vtu.hooks) - 1; i >= 0; i-- {
			if vtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vtu *VipTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := vtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vtu *VipTypeUpdate) Exec(ctx context.Context) error {
	_, err := vtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtu *VipTypeUpdate) ExecX(ctx context.Context) {
	if err := vtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vtu *VipTypeUpdate) defaults() {
	if _, ok := vtu.mutation.UpdatedAt(); !ok {
		v := viptype.UpdateDefaultUpdatedAt()
		vtu.mutation.SetUpdatedAt(v)
	}
}

func (vtu *VipTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   viptype.Table,
			Columns: viptype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: viptype.FieldID,
			},
		},
	}
	if ps := vtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vtu.mutation.VipName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: viptype.FieldVipName,
		})
	}
	if vtu.mutation.VipNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: viptype.FieldVipName,
		})
	}
	if value, ok := vtu.mutation.IsSuper(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: viptype.FieldIsSuper,
		})
	}
	if vtu.mutation.IsSuperCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: viptype.FieldIsSuper,
		})
	}
	if value, ok := vtu.mutation.ValidDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: viptype.FieldValidDays,
		})
	}
	if value, ok := vtu.mutation.AddedValidDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: viptype.FieldValidDays,
		})
	}
	if vtu.mutation.ValidDaysCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: viptype.FieldValidDays,
		})
	}
	if value, ok := vtu.mutation.DiscountRate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldDiscountRate,
		})
	}
	if value, ok := vtu.mutation.AddedDiscountRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldDiscountRate,
		})
	}
	if vtu.mutation.DiscountRateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: viptype.FieldDiscountRate,
		})
	}
	if value, ok := vtu.mutation.AvatarId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldAvatarId,
		})
	}
	if value, ok := vtu.mutation.AddedAvatarId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldAvatarId,
		})
	}
	if vtu.mutation.AvatarIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: viptype.FieldAvatarId,
		})
	}
	if value, ok := vtu.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: viptype.FieldSummary,
		})
	}
	if vtu.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: viptype.FieldSummary,
		})
	}
	if value, ok := vtu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: viptype.FieldUpdatedAt,
		})
	}
	if value, ok := vtu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldCreateBy,
		})
	}
	if value, ok := vtu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldCreateBy,
		})
	}
	if value, ok := vtu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldUpdateBy,
		})
	}
	if value, ok := vtu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldUpdateBy,
		})
	}
	if value, ok := vtu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldTenantId,
		})
	}
	if value, ok := vtu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldTenantId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{viptype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// VipTypeUpdateOne is the builder for updating a single VipType entity.
type VipTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VipTypeMutation
}

// SetVipName sets the "vipName" field.
func (vtuo *VipTypeUpdateOne) SetVipName(s string) *VipTypeUpdateOne {
	vtuo.mutation.SetVipName(s)
	return vtuo
}

// SetNillableVipName sets the "vipName" field if the given value is not nil.
func (vtuo *VipTypeUpdateOne) SetNillableVipName(s *string) *VipTypeUpdateOne {
	if s != nil {
		vtuo.SetVipName(*s)
	}
	return vtuo
}

// ClearVipName clears the value of the "vipName" field.
func (vtuo *VipTypeUpdateOne) ClearVipName() *VipTypeUpdateOne {
	vtuo.mutation.ClearVipName()
	return vtuo
}

// SetIsSuper sets the "isSuper" field.
func (vtuo *VipTypeUpdateOne) SetIsSuper(b bool) *VipTypeUpdateOne {
	vtuo.mutation.SetIsSuper(b)
	return vtuo
}

// SetNillableIsSuper sets the "isSuper" field if the given value is not nil.
func (vtuo *VipTypeUpdateOne) SetNillableIsSuper(b *bool) *VipTypeUpdateOne {
	if b != nil {
		vtuo.SetIsSuper(*b)
	}
	return vtuo
}

// ClearIsSuper clears the value of the "isSuper" field.
func (vtuo *VipTypeUpdateOne) ClearIsSuper() *VipTypeUpdateOne {
	vtuo.mutation.ClearIsSuper()
	return vtuo
}

// SetValidDays sets the "validDays" field.
func (vtuo *VipTypeUpdateOne) SetValidDays(i int32) *VipTypeUpdateOne {
	vtuo.mutation.ResetValidDays()
	vtuo.mutation.SetValidDays(i)
	return vtuo
}

// SetNillableValidDays sets the "validDays" field if the given value is not nil.
func (vtuo *VipTypeUpdateOne) SetNillableValidDays(i *int32) *VipTypeUpdateOne {
	if i != nil {
		vtuo.SetValidDays(*i)
	}
	return vtuo
}

// AddValidDays adds i to the "validDays" field.
func (vtuo *VipTypeUpdateOne) AddValidDays(i int32) *VipTypeUpdateOne {
	vtuo.mutation.AddValidDays(i)
	return vtuo
}

// ClearValidDays clears the value of the "validDays" field.
func (vtuo *VipTypeUpdateOne) ClearValidDays() *VipTypeUpdateOne {
	vtuo.mutation.ClearValidDays()
	return vtuo
}

// SetDiscountRate sets the "discountRate" field.
func (vtuo *VipTypeUpdateOne) SetDiscountRate(i int64) *VipTypeUpdateOne {
	vtuo.mutation.ResetDiscountRate()
	vtuo.mutation.SetDiscountRate(i)
	return vtuo
}

// SetNillableDiscountRate sets the "discountRate" field if the given value is not nil.
func (vtuo *VipTypeUpdateOne) SetNillableDiscountRate(i *int64) *VipTypeUpdateOne {
	if i != nil {
		vtuo.SetDiscountRate(*i)
	}
	return vtuo
}

// AddDiscountRate adds i to the "discountRate" field.
func (vtuo *VipTypeUpdateOne) AddDiscountRate(i int64) *VipTypeUpdateOne {
	vtuo.mutation.AddDiscountRate(i)
	return vtuo
}

// ClearDiscountRate clears the value of the "discountRate" field.
func (vtuo *VipTypeUpdateOne) ClearDiscountRate() *VipTypeUpdateOne {
	vtuo.mutation.ClearDiscountRate()
	return vtuo
}

// SetAvatarId sets the "avatarId" field.
func (vtuo *VipTypeUpdateOne) SetAvatarId(i int64) *VipTypeUpdateOne {
	vtuo.mutation.ResetAvatarId()
	vtuo.mutation.SetAvatarId(i)
	return vtuo
}

// SetNillableAvatarId sets the "avatarId" field if the given value is not nil.
func (vtuo *VipTypeUpdateOne) SetNillableAvatarId(i *int64) *VipTypeUpdateOne {
	if i != nil {
		vtuo.SetAvatarId(*i)
	}
	return vtuo
}

// AddAvatarId adds i to the "avatarId" field.
func (vtuo *VipTypeUpdateOne) AddAvatarId(i int64) *VipTypeUpdateOne {
	vtuo.mutation.AddAvatarId(i)
	return vtuo
}

// ClearAvatarId clears the value of the "avatarId" field.
func (vtuo *VipTypeUpdateOne) ClearAvatarId() *VipTypeUpdateOne {
	vtuo.mutation.ClearAvatarId()
	return vtuo
}

// SetSummary sets the "summary" field.
func (vtuo *VipTypeUpdateOne) SetSummary(s string) *VipTypeUpdateOne {
	vtuo.mutation.SetSummary(s)
	return vtuo
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (vtuo *VipTypeUpdateOne) SetNillableSummary(s *string) *VipTypeUpdateOne {
	if s != nil {
		vtuo.SetSummary(*s)
	}
	return vtuo
}

// ClearSummary clears the value of the "summary" field.
func (vtuo *VipTypeUpdateOne) ClearSummary() *VipTypeUpdateOne {
	vtuo.mutation.ClearSummary()
	return vtuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (vtuo *VipTypeUpdateOne) SetUpdatedAt(t time.Time) *VipTypeUpdateOne {
	vtuo.mutation.SetUpdatedAt(t)
	return vtuo
}

// SetCreateBy sets the "createBy" field.
func (vtuo *VipTypeUpdateOne) SetCreateBy(i int64) *VipTypeUpdateOne {
	vtuo.mutation.ResetCreateBy()
	vtuo.mutation.SetCreateBy(i)
	return vtuo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (vtuo *VipTypeUpdateOne) SetNillableCreateBy(i *int64) *VipTypeUpdateOne {
	if i != nil {
		vtuo.SetCreateBy(*i)
	}
	return vtuo
}

// AddCreateBy adds i to the "createBy" field.
func (vtuo *VipTypeUpdateOne) AddCreateBy(i int64) *VipTypeUpdateOne {
	vtuo.mutation.AddCreateBy(i)
	return vtuo
}

// SetUpdateBy sets the "updateBy" field.
func (vtuo *VipTypeUpdateOne) SetUpdateBy(i int64) *VipTypeUpdateOne {
	vtuo.mutation.ResetUpdateBy()
	vtuo.mutation.SetUpdateBy(i)
	return vtuo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (vtuo *VipTypeUpdateOne) SetNillableUpdateBy(i *int64) *VipTypeUpdateOne {
	if i != nil {
		vtuo.SetUpdateBy(*i)
	}
	return vtuo
}

// AddUpdateBy adds i to the "updateBy" field.
func (vtuo *VipTypeUpdateOne) AddUpdateBy(i int64) *VipTypeUpdateOne {
	vtuo.mutation.AddUpdateBy(i)
	return vtuo
}

// SetTenantId sets the "tenantId" field.
func (vtuo *VipTypeUpdateOne) SetTenantId(i int64) *VipTypeUpdateOne {
	vtuo.mutation.ResetTenantId()
	vtuo.mutation.SetTenantId(i)
	return vtuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (vtuo *VipTypeUpdateOne) SetNillableTenantId(i *int64) *VipTypeUpdateOne {
	if i != nil {
		vtuo.SetTenantId(*i)
	}
	return vtuo
}

// AddTenantId adds i to the "tenantId" field.
func (vtuo *VipTypeUpdateOne) AddTenantId(i int64) *VipTypeUpdateOne {
	vtuo.mutation.AddTenantId(i)
	return vtuo
}

// Mutation returns the VipTypeMutation object of the builder.
func (vtuo *VipTypeUpdateOne) Mutation() *VipTypeMutation {
	return vtuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vtuo *VipTypeUpdateOne) Select(field string, fields ...string) *VipTypeUpdateOne {
	vtuo.fields = append([]string{field}, fields...)
	return vtuo
}

// Save executes the query and returns the updated VipType entity.
func (vtuo *VipTypeUpdateOne) Save(ctx context.Context) (*VipType, error) {
	var (
		err  error
		node *VipType
	)
	vtuo.defaults()
	if len(vtuo.hooks) == 0 {
		node, err = vtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VipTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vtuo.mutation = mutation
			node, err = vtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vtuo.hooks) - 1; i >= 0; i-- {
			if vtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vtuo *VipTypeUpdateOne) SaveX(ctx context.Context) *VipType {
	node, err := vtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vtuo *VipTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := vtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtuo *VipTypeUpdateOne) ExecX(ctx context.Context) {
	if err := vtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vtuo *VipTypeUpdateOne) defaults() {
	if _, ok := vtuo.mutation.UpdatedAt(); !ok {
		v := viptype.UpdateDefaultUpdatedAt()
		vtuo.mutation.SetUpdatedAt(v)
	}
}

func (vtuo *VipTypeUpdateOne) sqlSave(ctx context.Context) (_node *VipType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   viptype.Table,
			Columns: viptype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: viptype.FieldID,
			},
		},
	}
	id, ok := vtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing VipType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := vtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, viptype.FieldID)
		for _, f := range fields {
			if !viptype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != viptype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vtuo.mutation.VipName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: viptype.FieldVipName,
		})
	}
	if vtuo.mutation.VipNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: viptype.FieldVipName,
		})
	}
	if value, ok := vtuo.mutation.IsSuper(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: viptype.FieldIsSuper,
		})
	}
	if vtuo.mutation.IsSuperCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: viptype.FieldIsSuper,
		})
	}
	if value, ok := vtuo.mutation.ValidDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: viptype.FieldValidDays,
		})
	}
	if value, ok := vtuo.mutation.AddedValidDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: viptype.FieldValidDays,
		})
	}
	if vtuo.mutation.ValidDaysCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: viptype.FieldValidDays,
		})
	}
	if value, ok := vtuo.mutation.DiscountRate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldDiscountRate,
		})
	}
	if value, ok := vtuo.mutation.AddedDiscountRate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldDiscountRate,
		})
	}
	if vtuo.mutation.DiscountRateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: viptype.FieldDiscountRate,
		})
	}
	if value, ok := vtuo.mutation.AvatarId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldAvatarId,
		})
	}
	if value, ok := vtuo.mutation.AddedAvatarId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldAvatarId,
		})
	}
	if vtuo.mutation.AvatarIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: viptype.FieldAvatarId,
		})
	}
	if value, ok := vtuo.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: viptype.FieldSummary,
		})
	}
	if vtuo.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: viptype.FieldSummary,
		})
	}
	if value, ok := vtuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: viptype.FieldUpdatedAt,
		})
	}
	if value, ok := vtuo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldCreateBy,
		})
	}
	if value, ok := vtuo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldCreateBy,
		})
	}
	if value, ok := vtuo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldUpdateBy,
		})
	}
	if value, ok := vtuo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldUpdateBy,
		})
	}
	if value, ok := vtuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldTenantId,
		})
	}
	if value, ok := vtuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: viptype.FieldTenantId,
		})
	}
	_node = &VipType{config: vtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{viptype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}