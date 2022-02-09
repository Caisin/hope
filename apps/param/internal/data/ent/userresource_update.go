// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/userresource"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserResourceUpdate is the builder for updating UserResource entities.
type UserResourceUpdate struct {
	config
	hooks    []Hook
	mutation *UserResourceMutation
}

// Where appends a list predicates to the UserResourceUpdate builder.
func (uru *UserResourceUpdate) Where(ps ...predicate.UserResource) *UserResourceUpdate {
	uru.mutation.Where(ps...)
	return uru
}

// SetResType sets the "resType" field.
func (uru *UserResourceUpdate) SetResType(s string) *UserResourceUpdate {
	uru.mutation.SetResType(s)
	return uru
}

// SetNillableResType sets the "resType" field if the given value is not nil.
func (uru *UserResourceUpdate) SetNillableResType(s *string) *UserResourceUpdate {
	if s != nil {
		uru.SetResType(*s)
	}
	return uru
}

// ClearResType clears the value of the "resType" field.
func (uru *UserResourceUpdate) ClearResType() *UserResourceUpdate {
	uru.mutation.ClearResType()
	return uru
}

// SetName sets the "Name" field.
func (uru *UserResourceUpdate) SetName(s string) *UserResourceUpdate {
	uru.mutation.SetName(s)
	return uru
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (uru *UserResourceUpdate) SetNillableName(s *string) *UserResourceUpdate {
	if s != nil {
		uru.SetName(*s)
	}
	return uru
}

// ClearName clears the value of the "Name" field.
func (uru *UserResourceUpdate) ClearName() *UserResourceUpdate {
	uru.mutation.ClearName()
	return uru
}

// SetURL sets the "url" field.
func (uru *UserResourceUpdate) SetURL(s string) *UserResourceUpdate {
	uru.mutation.SetURL(s)
	return uru
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (uru *UserResourceUpdate) SetNillableURL(s *string) *UserResourceUpdate {
	if s != nil {
		uru.SetURL(*s)
	}
	return uru
}

// ClearURL clears the value of the "url" field.
func (uru *UserResourceUpdate) ClearURL() *UserResourceUpdate {
	uru.mutation.ClearURL()
	return uru
}

// SetSummary sets the "summary" field.
func (uru *UserResourceUpdate) SetSummary(s string) *UserResourceUpdate {
	uru.mutation.SetSummary(s)
	return uru
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (uru *UserResourceUpdate) SetNillableSummary(s *string) *UserResourceUpdate {
	if s != nil {
		uru.SetSummary(*s)
	}
	return uru
}

// ClearSummary clears the value of the "summary" field.
func (uru *UserResourceUpdate) ClearSummary() *UserResourceUpdate {
	uru.mutation.ClearSummary()
	return uru
}

// SetUpdatedAt sets the "updatedAt" field.
func (uru *UserResourceUpdate) SetUpdatedAt(t time.Time) *UserResourceUpdate {
	uru.mutation.SetUpdatedAt(t)
	return uru
}

// SetCreateBy sets the "createBy" field.
func (uru *UserResourceUpdate) SetCreateBy(i int64) *UserResourceUpdate {
	uru.mutation.ResetCreateBy()
	uru.mutation.SetCreateBy(i)
	return uru
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (uru *UserResourceUpdate) SetNillableCreateBy(i *int64) *UserResourceUpdate {
	if i != nil {
		uru.SetCreateBy(*i)
	}
	return uru
}

// AddCreateBy adds i to the "createBy" field.
func (uru *UserResourceUpdate) AddCreateBy(i int64) *UserResourceUpdate {
	uru.mutation.AddCreateBy(i)
	return uru
}

// SetUpdateBy sets the "updateBy" field.
func (uru *UserResourceUpdate) SetUpdateBy(i int64) *UserResourceUpdate {
	uru.mutation.ResetUpdateBy()
	uru.mutation.SetUpdateBy(i)
	return uru
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (uru *UserResourceUpdate) SetNillableUpdateBy(i *int64) *UserResourceUpdate {
	if i != nil {
		uru.SetUpdateBy(*i)
	}
	return uru
}

// AddUpdateBy adds i to the "updateBy" field.
func (uru *UserResourceUpdate) AddUpdateBy(i int64) *UserResourceUpdate {
	uru.mutation.AddUpdateBy(i)
	return uru
}

// SetTenantId sets the "tenantId" field.
func (uru *UserResourceUpdate) SetTenantId(i int64) *UserResourceUpdate {
	uru.mutation.ResetTenantId()
	uru.mutation.SetTenantId(i)
	return uru
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (uru *UserResourceUpdate) SetNillableTenantId(i *int64) *UserResourceUpdate {
	if i != nil {
		uru.SetTenantId(*i)
	}
	return uru
}

// AddTenantId adds i to the "tenantId" field.
func (uru *UserResourceUpdate) AddTenantId(i int64) *UserResourceUpdate {
	uru.mutation.AddTenantId(i)
	return uru
}

// Mutation returns the UserResourceMutation object of the builder.
func (uru *UserResourceUpdate) Mutation() *UserResourceMutation {
	return uru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uru *UserResourceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uru.defaults()
	if len(uru.hooks) == 0 {
		affected, err = uru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uru.mutation = mutation
			affected, err = uru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uru.hooks) - 1; i >= 0; i-- {
			if uru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uru *UserResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := uru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uru *UserResourceUpdate) Exec(ctx context.Context) error {
	_, err := uru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uru *UserResourceUpdate) ExecX(ctx context.Context) {
	if err := uru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uru *UserResourceUpdate) defaults() {
	if _, ok := uru.mutation.UpdatedAt(); !ok {
		v := userresource.UpdateDefaultUpdatedAt()
		uru.mutation.SetUpdatedAt(v)
	}
}

func (uru *UserResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userresource.Table,
			Columns: userresource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userresource.FieldID,
			},
		},
	}
	if ps := uru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uru.mutation.ResType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userresource.FieldResType,
		})
	}
	if uru.mutation.ResTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userresource.FieldResType,
		})
	}
	if value, ok := uru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userresource.FieldName,
		})
	}
	if uru.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userresource.FieldName,
		})
	}
	if value, ok := uru.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userresource.FieldURL,
		})
	}
	if uru.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userresource.FieldURL,
		})
	}
	if value, ok := uru.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userresource.FieldSummary,
		})
	}
	if uru.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userresource.FieldSummary,
		})
	}
	if value, ok := uru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userresource.FieldUpdatedAt,
		})
	}
	if value, ok := uru.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldCreateBy,
		})
	}
	if value, ok := uru.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldCreateBy,
		})
	}
	if value, ok := uru.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldUpdateBy,
		})
	}
	if value, ok := uru.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldUpdateBy,
		})
	}
	if value, ok := uru.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldTenantId,
		})
	}
	if value, ok := uru.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldTenantId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserResourceUpdateOne is the builder for updating a single UserResource entity.
type UserResourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserResourceMutation
}

// SetResType sets the "resType" field.
func (uruo *UserResourceUpdateOne) SetResType(s string) *UserResourceUpdateOne {
	uruo.mutation.SetResType(s)
	return uruo
}

// SetNillableResType sets the "resType" field if the given value is not nil.
func (uruo *UserResourceUpdateOne) SetNillableResType(s *string) *UserResourceUpdateOne {
	if s != nil {
		uruo.SetResType(*s)
	}
	return uruo
}

// ClearResType clears the value of the "resType" field.
func (uruo *UserResourceUpdateOne) ClearResType() *UserResourceUpdateOne {
	uruo.mutation.ClearResType()
	return uruo
}

// SetName sets the "Name" field.
func (uruo *UserResourceUpdateOne) SetName(s string) *UserResourceUpdateOne {
	uruo.mutation.SetName(s)
	return uruo
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (uruo *UserResourceUpdateOne) SetNillableName(s *string) *UserResourceUpdateOne {
	if s != nil {
		uruo.SetName(*s)
	}
	return uruo
}

// ClearName clears the value of the "Name" field.
func (uruo *UserResourceUpdateOne) ClearName() *UserResourceUpdateOne {
	uruo.mutation.ClearName()
	return uruo
}

// SetURL sets the "url" field.
func (uruo *UserResourceUpdateOne) SetURL(s string) *UserResourceUpdateOne {
	uruo.mutation.SetURL(s)
	return uruo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (uruo *UserResourceUpdateOne) SetNillableURL(s *string) *UserResourceUpdateOne {
	if s != nil {
		uruo.SetURL(*s)
	}
	return uruo
}

// ClearURL clears the value of the "url" field.
func (uruo *UserResourceUpdateOne) ClearURL() *UserResourceUpdateOne {
	uruo.mutation.ClearURL()
	return uruo
}

// SetSummary sets the "summary" field.
func (uruo *UserResourceUpdateOne) SetSummary(s string) *UserResourceUpdateOne {
	uruo.mutation.SetSummary(s)
	return uruo
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (uruo *UserResourceUpdateOne) SetNillableSummary(s *string) *UserResourceUpdateOne {
	if s != nil {
		uruo.SetSummary(*s)
	}
	return uruo
}

// ClearSummary clears the value of the "summary" field.
func (uruo *UserResourceUpdateOne) ClearSummary() *UserResourceUpdateOne {
	uruo.mutation.ClearSummary()
	return uruo
}

// SetUpdatedAt sets the "updatedAt" field.
func (uruo *UserResourceUpdateOne) SetUpdatedAt(t time.Time) *UserResourceUpdateOne {
	uruo.mutation.SetUpdatedAt(t)
	return uruo
}

// SetCreateBy sets the "createBy" field.
func (uruo *UserResourceUpdateOne) SetCreateBy(i int64) *UserResourceUpdateOne {
	uruo.mutation.ResetCreateBy()
	uruo.mutation.SetCreateBy(i)
	return uruo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (uruo *UserResourceUpdateOne) SetNillableCreateBy(i *int64) *UserResourceUpdateOne {
	if i != nil {
		uruo.SetCreateBy(*i)
	}
	return uruo
}

// AddCreateBy adds i to the "createBy" field.
func (uruo *UserResourceUpdateOne) AddCreateBy(i int64) *UserResourceUpdateOne {
	uruo.mutation.AddCreateBy(i)
	return uruo
}

// SetUpdateBy sets the "updateBy" field.
func (uruo *UserResourceUpdateOne) SetUpdateBy(i int64) *UserResourceUpdateOne {
	uruo.mutation.ResetUpdateBy()
	uruo.mutation.SetUpdateBy(i)
	return uruo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (uruo *UserResourceUpdateOne) SetNillableUpdateBy(i *int64) *UserResourceUpdateOne {
	if i != nil {
		uruo.SetUpdateBy(*i)
	}
	return uruo
}

// AddUpdateBy adds i to the "updateBy" field.
func (uruo *UserResourceUpdateOne) AddUpdateBy(i int64) *UserResourceUpdateOne {
	uruo.mutation.AddUpdateBy(i)
	return uruo
}

// SetTenantId sets the "tenantId" field.
func (uruo *UserResourceUpdateOne) SetTenantId(i int64) *UserResourceUpdateOne {
	uruo.mutation.ResetTenantId()
	uruo.mutation.SetTenantId(i)
	return uruo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (uruo *UserResourceUpdateOne) SetNillableTenantId(i *int64) *UserResourceUpdateOne {
	if i != nil {
		uruo.SetTenantId(*i)
	}
	return uruo
}

// AddTenantId adds i to the "tenantId" field.
func (uruo *UserResourceUpdateOne) AddTenantId(i int64) *UserResourceUpdateOne {
	uruo.mutation.AddTenantId(i)
	return uruo
}

// Mutation returns the UserResourceMutation object of the builder.
func (uruo *UserResourceUpdateOne) Mutation() *UserResourceMutation {
	return uruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uruo *UserResourceUpdateOne) Select(field string, fields ...string) *UserResourceUpdateOne {
	uruo.fields = append([]string{field}, fields...)
	return uruo
}

// Save executes the query and returns the updated UserResource entity.
func (uruo *UserResourceUpdateOne) Save(ctx context.Context) (*UserResource, error) {
	var (
		err  error
		node *UserResource
	)
	uruo.defaults()
	if len(uruo.hooks) == 0 {
		node, err = uruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uruo.mutation = mutation
			node, err = uruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uruo.hooks) - 1; i >= 0; i-- {
			if uruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uruo *UserResourceUpdateOne) SaveX(ctx context.Context) *UserResource {
	node, err := uruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uruo *UserResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := uruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uruo *UserResourceUpdateOne) ExecX(ctx context.Context) {
	if err := uruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uruo *UserResourceUpdateOne) defaults() {
	if _, ok := uruo.mutation.UpdatedAt(); !ok {
		v := userresource.UpdateDefaultUpdatedAt()
		uruo.mutation.SetUpdatedAt(v)
	}
}

func (uruo *UserResourceUpdateOne) sqlSave(ctx context.Context) (_node *UserResource, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userresource.Table,
			Columns: userresource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userresource.FieldID,
			},
		},
	}
	id, ok := uruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserResource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userresource.FieldID)
		for _, f := range fields {
			if !userresource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userresource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uruo.mutation.ResType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userresource.FieldResType,
		})
	}
	if uruo.mutation.ResTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userresource.FieldResType,
		})
	}
	if value, ok := uruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userresource.FieldName,
		})
	}
	if uruo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userresource.FieldName,
		})
	}
	if value, ok := uruo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userresource.FieldURL,
		})
	}
	if uruo.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userresource.FieldURL,
		})
	}
	if value, ok := uruo.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userresource.FieldSummary,
		})
	}
	if uruo.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userresource.FieldSummary,
		})
	}
	if value, ok := uruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userresource.FieldUpdatedAt,
		})
	}
	if value, ok := uruo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldCreateBy,
		})
	}
	if value, ok := uruo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldCreateBy,
		})
	}
	if value, ok := uruo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldUpdateBy,
		})
	}
	if value, ok := uruo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldUpdateBy,
		})
	}
	if value, ok := uruo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldTenantId,
		})
	}
	if value, ok := uruo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userresource.FieldTenantId,
		})
	}
	_node = &UserResource{config: uruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
