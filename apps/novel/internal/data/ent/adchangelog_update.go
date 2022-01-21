// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/adchangelog"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdChangeLogUpdate is the builder for updating AdChangeLog entities.
type AdChangeLogUpdate struct {
	config
	hooks    []Hook
	mutation *AdChangeLogMutation
}

// Where appends a list predicates to the AdChangeLogUpdate builder.
func (aclu *AdChangeLogUpdate) Where(ps ...predicate.AdChangeLog) *AdChangeLogUpdate {
	aclu.mutation.Where(ps...)
	return aclu
}

// SetUserId sets the "userId" field.
func (aclu *AdChangeLogUpdate) SetUserId(i int64) *AdChangeLogUpdate {
	aclu.mutation.SetUserId(i)
	return aclu
}

// SetAdId sets the "adId" field.
func (aclu *AdChangeLogUpdate) SetAdId(s string) *AdChangeLogUpdate {
	aclu.mutation.SetAdId(s)
	return aclu
}

// SetNillableAdId sets the "adId" field if the given value is not nil.
func (aclu *AdChangeLogUpdate) SetNillableAdId(s *string) *AdChangeLogUpdate {
	if s != nil {
		aclu.SetAdId(*s)
	}
	return aclu
}

// ClearAdId clears the value of the "adId" field.
func (aclu *AdChangeLogUpdate) ClearAdId() *AdChangeLogUpdate {
	aclu.mutation.ClearAdId()
	return aclu
}

// SetChId sets the "chId" field.
func (aclu *AdChangeLogUpdate) SetChId(i int64) *AdChangeLogUpdate {
	aclu.mutation.ResetChId()
	aclu.mutation.SetChId(i)
	return aclu
}

// SetNillableChId sets the "chId" field if the given value is not nil.
func (aclu *AdChangeLogUpdate) SetNillableChId(i *int64) *AdChangeLogUpdate {
	if i != nil {
		aclu.SetChId(*i)
	}
	return aclu
}

// AddChId adds i to the "chId" field.
func (aclu *AdChangeLogUpdate) AddChId(i int64) *AdChangeLogUpdate {
	aclu.mutation.AddChId(i)
	return aclu
}

// ClearChId clears the value of the "chId" field.
func (aclu *AdChangeLogUpdate) ClearChId() *AdChangeLogUpdate {
	aclu.mutation.ClearChId()
	return aclu
}

// SetDeviceId sets the "deviceId" field.
func (aclu *AdChangeLogUpdate) SetDeviceId(s string) *AdChangeLogUpdate {
	aclu.mutation.SetDeviceId(s)
	return aclu
}

// SetNillableDeviceId sets the "deviceId" field if the given value is not nil.
func (aclu *AdChangeLogUpdate) SetNillableDeviceId(s *string) *AdChangeLogUpdate {
	if s != nil {
		aclu.SetDeviceId(*s)
	}
	return aclu
}

// ClearDeviceId clears the value of the "deviceId" field.
func (aclu *AdChangeLogUpdate) ClearDeviceId() *AdChangeLogUpdate {
	aclu.mutation.ClearDeviceId()
	return aclu
}

// SetExtInfo sets the "extInfo" field.
func (aclu *AdChangeLogUpdate) SetExtInfo(s string) *AdChangeLogUpdate {
	aclu.mutation.SetExtInfo(s)
	return aclu
}

// SetNillableExtInfo sets the "extInfo" field if the given value is not nil.
func (aclu *AdChangeLogUpdate) SetNillableExtInfo(s *string) *AdChangeLogUpdate {
	if s != nil {
		aclu.SetExtInfo(*s)
	}
	return aclu
}

// ClearExtInfo clears the value of the "extInfo" field.
func (aclu *AdChangeLogUpdate) ClearExtInfo() *AdChangeLogUpdate {
	aclu.mutation.ClearExtInfo()
	return aclu
}

// SetUpdatedAt sets the "updatedAt" field.
func (aclu *AdChangeLogUpdate) SetUpdatedAt(t time.Time) *AdChangeLogUpdate {
	aclu.mutation.SetUpdatedAt(t)
	return aclu
}

// SetCreateBy sets the "createBy" field.
func (aclu *AdChangeLogUpdate) SetCreateBy(i int64) *AdChangeLogUpdate {
	aclu.mutation.ResetCreateBy()
	aclu.mutation.SetCreateBy(i)
	return aclu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (aclu *AdChangeLogUpdate) SetNillableCreateBy(i *int64) *AdChangeLogUpdate {
	if i != nil {
		aclu.SetCreateBy(*i)
	}
	return aclu
}

// AddCreateBy adds i to the "createBy" field.
func (aclu *AdChangeLogUpdate) AddCreateBy(i int64) *AdChangeLogUpdate {
	aclu.mutation.AddCreateBy(i)
	return aclu
}

// SetUpdateBy sets the "updateBy" field.
func (aclu *AdChangeLogUpdate) SetUpdateBy(i int64) *AdChangeLogUpdate {
	aclu.mutation.ResetUpdateBy()
	aclu.mutation.SetUpdateBy(i)
	return aclu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (aclu *AdChangeLogUpdate) SetNillableUpdateBy(i *int64) *AdChangeLogUpdate {
	if i != nil {
		aclu.SetUpdateBy(*i)
	}
	return aclu
}

// AddUpdateBy adds i to the "updateBy" field.
func (aclu *AdChangeLogUpdate) AddUpdateBy(i int64) *AdChangeLogUpdate {
	aclu.mutation.AddUpdateBy(i)
	return aclu
}

// SetTenantId sets the "tenantId" field.
func (aclu *AdChangeLogUpdate) SetTenantId(i int64) *AdChangeLogUpdate {
	aclu.mutation.ResetTenantId()
	aclu.mutation.SetTenantId(i)
	return aclu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (aclu *AdChangeLogUpdate) SetNillableTenantId(i *int64) *AdChangeLogUpdate {
	if i != nil {
		aclu.SetTenantId(*i)
	}
	return aclu
}

// AddTenantId adds i to the "tenantId" field.
func (aclu *AdChangeLogUpdate) AddTenantId(i int64) *AdChangeLogUpdate {
	aclu.mutation.AddTenantId(i)
	return aclu
}

// SetUserID sets the "user" edge to the SocialUser entity by ID.
func (aclu *AdChangeLogUpdate) SetUserID(id int64) *AdChangeLogUpdate {
	aclu.mutation.SetUserID(id)
	return aclu
}

// SetUser sets the "user" edge to the SocialUser entity.
func (aclu *AdChangeLogUpdate) SetUser(s *SocialUser) *AdChangeLogUpdate {
	return aclu.SetUserID(s.ID)
}

// Mutation returns the AdChangeLogMutation object of the builder.
func (aclu *AdChangeLogUpdate) Mutation() *AdChangeLogMutation {
	return aclu.mutation
}

// ClearUser clears the "user" edge to the SocialUser entity.
func (aclu *AdChangeLogUpdate) ClearUser() *AdChangeLogUpdate {
	aclu.mutation.ClearUser()
	return aclu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aclu *AdChangeLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aclu.defaults()
	if len(aclu.hooks) == 0 {
		if err = aclu.check(); err != nil {
			return 0, err
		}
		affected, err = aclu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdChangeLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aclu.check(); err != nil {
				return 0, err
			}
			aclu.mutation = mutation
			affected, err = aclu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aclu.hooks) - 1; i >= 0; i-- {
			if aclu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aclu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aclu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aclu *AdChangeLogUpdate) SaveX(ctx context.Context) int {
	affected, err := aclu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aclu *AdChangeLogUpdate) Exec(ctx context.Context) error {
	_, err := aclu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aclu *AdChangeLogUpdate) ExecX(ctx context.Context) {
	if err := aclu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aclu *AdChangeLogUpdate) defaults() {
	if _, ok := aclu.mutation.UpdatedAt(); !ok {
		v := adchangelog.UpdateDefaultUpdatedAt()
		aclu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aclu *AdChangeLogUpdate) check() error {
	if _, ok := aclu.mutation.UserID(); aclu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (aclu *AdChangeLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adchangelog.Table,
			Columns: adchangelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: adchangelog.FieldID,
			},
		},
	}
	if ps := aclu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aclu.mutation.AdId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adchangelog.FieldAdId,
		})
	}
	if aclu.mutation.AdIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: adchangelog.FieldAdId,
		})
	}
	if value, ok := aclu.mutation.ChId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldChId,
		})
	}
	if value, ok := aclu.mutation.AddedChId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldChId,
		})
	}
	if aclu.mutation.ChIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: adchangelog.FieldChId,
		})
	}
	if value, ok := aclu.mutation.DeviceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adchangelog.FieldDeviceId,
		})
	}
	if aclu.mutation.DeviceIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: adchangelog.FieldDeviceId,
		})
	}
	if value, ok := aclu.mutation.ExtInfo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adchangelog.FieldExtInfo,
		})
	}
	if aclu.mutation.ExtInfoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: adchangelog.FieldExtInfo,
		})
	}
	if value, ok := aclu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: adchangelog.FieldUpdatedAt,
		})
	}
	if value, ok := aclu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldCreateBy,
		})
	}
	if value, ok := aclu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldCreateBy,
		})
	}
	if value, ok := aclu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldUpdateBy,
		})
	}
	if value, ok := aclu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldUpdateBy,
		})
	}
	if value, ok := aclu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldTenantId,
		})
	}
	if value, ok := aclu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldTenantId,
		})
	}
	if aclu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adchangelog.UserTable,
			Columns: []string{adchangelog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: socialuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aclu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adchangelog.UserTable,
			Columns: []string{adchangelog.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aclu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adchangelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AdChangeLogUpdateOne is the builder for updating a single AdChangeLog entity.
type AdChangeLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdChangeLogMutation
}

// SetUserId sets the "userId" field.
func (acluo *AdChangeLogUpdateOne) SetUserId(i int64) *AdChangeLogUpdateOne {
	acluo.mutation.SetUserId(i)
	return acluo
}

// SetAdId sets the "adId" field.
func (acluo *AdChangeLogUpdateOne) SetAdId(s string) *AdChangeLogUpdateOne {
	acluo.mutation.SetAdId(s)
	return acluo
}

// SetNillableAdId sets the "adId" field if the given value is not nil.
func (acluo *AdChangeLogUpdateOne) SetNillableAdId(s *string) *AdChangeLogUpdateOne {
	if s != nil {
		acluo.SetAdId(*s)
	}
	return acluo
}

// ClearAdId clears the value of the "adId" field.
func (acluo *AdChangeLogUpdateOne) ClearAdId() *AdChangeLogUpdateOne {
	acluo.mutation.ClearAdId()
	return acluo
}

// SetChId sets the "chId" field.
func (acluo *AdChangeLogUpdateOne) SetChId(i int64) *AdChangeLogUpdateOne {
	acluo.mutation.ResetChId()
	acluo.mutation.SetChId(i)
	return acluo
}

// SetNillableChId sets the "chId" field if the given value is not nil.
func (acluo *AdChangeLogUpdateOne) SetNillableChId(i *int64) *AdChangeLogUpdateOne {
	if i != nil {
		acluo.SetChId(*i)
	}
	return acluo
}

// AddChId adds i to the "chId" field.
func (acluo *AdChangeLogUpdateOne) AddChId(i int64) *AdChangeLogUpdateOne {
	acluo.mutation.AddChId(i)
	return acluo
}

// ClearChId clears the value of the "chId" field.
func (acluo *AdChangeLogUpdateOne) ClearChId() *AdChangeLogUpdateOne {
	acluo.mutation.ClearChId()
	return acluo
}

// SetDeviceId sets the "deviceId" field.
func (acluo *AdChangeLogUpdateOne) SetDeviceId(s string) *AdChangeLogUpdateOne {
	acluo.mutation.SetDeviceId(s)
	return acluo
}

// SetNillableDeviceId sets the "deviceId" field if the given value is not nil.
func (acluo *AdChangeLogUpdateOne) SetNillableDeviceId(s *string) *AdChangeLogUpdateOne {
	if s != nil {
		acluo.SetDeviceId(*s)
	}
	return acluo
}

// ClearDeviceId clears the value of the "deviceId" field.
func (acluo *AdChangeLogUpdateOne) ClearDeviceId() *AdChangeLogUpdateOne {
	acluo.mutation.ClearDeviceId()
	return acluo
}

// SetExtInfo sets the "extInfo" field.
func (acluo *AdChangeLogUpdateOne) SetExtInfo(s string) *AdChangeLogUpdateOne {
	acluo.mutation.SetExtInfo(s)
	return acluo
}

// SetNillableExtInfo sets the "extInfo" field if the given value is not nil.
func (acluo *AdChangeLogUpdateOne) SetNillableExtInfo(s *string) *AdChangeLogUpdateOne {
	if s != nil {
		acluo.SetExtInfo(*s)
	}
	return acluo
}

// ClearExtInfo clears the value of the "extInfo" field.
func (acluo *AdChangeLogUpdateOne) ClearExtInfo() *AdChangeLogUpdateOne {
	acluo.mutation.ClearExtInfo()
	return acluo
}

// SetUpdatedAt sets the "updatedAt" field.
func (acluo *AdChangeLogUpdateOne) SetUpdatedAt(t time.Time) *AdChangeLogUpdateOne {
	acluo.mutation.SetUpdatedAt(t)
	return acluo
}

// SetCreateBy sets the "createBy" field.
func (acluo *AdChangeLogUpdateOne) SetCreateBy(i int64) *AdChangeLogUpdateOne {
	acluo.mutation.ResetCreateBy()
	acluo.mutation.SetCreateBy(i)
	return acluo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (acluo *AdChangeLogUpdateOne) SetNillableCreateBy(i *int64) *AdChangeLogUpdateOne {
	if i != nil {
		acluo.SetCreateBy(*i)
	}
	return acluo
}

// AddCreateBy adds i to the "createBy" field.
func (acluo *AdChangeLogUpdateOne) AddCreateBy(i int64) *AdChangeLogUpdateOne {
	acluo.mutation.AddCreateBy(i)
	return acluo
}

// SetUpdateBy sets the "updateBy" field.
func (acluo *AdChangeLogUpdateOne) SetUpdateBy(i int64) *AdChangeLogUpdateOne {
	acluo.mutation.ResetUpdateBy()
	acluo.mutation.SetUpdateBy(i)
	return acluo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (acluo *AdChangeLogUpdateOne) SetNillableUpdateBy(i *int64) *AdChangeLogUpdateOne {
	if i != nil {
		acluo.SetUpdateBy(*i)
	}
	return acluo
}

// AddUpdateBy adds i to the "updateBy" field.
func (acluo *AdChangeLogUpdateOne) AddUpdateBy(i int64) *AdChangeLogUpdateOne {
	acluo.mutation.AddUpdateBy(i)
	return acluo
}

// SetTenantId sets the "tenantId" field.
func (acluo *AdChangeLogUpdateOne) SetTenantId(i int64) *AdChangeLogUpdateOne {
	acluo.mutation.ResetTenantId()
	acluo.mutation.SetTenantId(i)
	return acluo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (acluo *AdChangeLogUpdateOne) SetNillableTenantId(i *int64) *AdChangeLogUpdateOne {
	if i != nil {
		acluo.SetTenantId(*i)
	}
	return acluo
}

// AddTenantId adds i to the "tenantId" field.
func (acluo *AdChangeLogUpdateOne) AddTenantId(i int64) *AdChangeLogUpdateOne {
	acluo.mutation.AddTenantId(i)
	return acluo
}

// SetUserID sets the "user" edge to the SocialUser entity by ID.
func (acluo *AdChangeLogUpdateOne) SetUserID(id int64) *AdChangeLogUpdateOne {
	acluo.mutation.SetUserID(id)
	return acluo
}

// SetUser sets the "user" edge to the SocialUser entity.
func (acluo *AdChangeLogUpdateOne) SetUser(s *SocialUser) *AdChangeLogUpdateOne {
	return acluo.SetUserID(s.ID)
}

// Mutation returns the AdChangeLogMutation object of the builder.
func (acluo *AdChangeLogUpdateOne) Mutation() *AdChangeLogMutation {
	return acluo.mutation
}

// ClearUser clears the "user" edge to the SocialUser entity.
func (acluo *AdChangeLogUpdateOne) ClearUser() *AdChangeLogUpdateOne {
	acluo.mutation.ClearUser()
	return acluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acluo *AdChangeLogUpdateOne) Select(field string, fields ...string) *AdChangeLogUpdateOne {
	acluo.fields = append([]string{field}, fields...)
	return acluo
}

// Save executes the query and returns the updated AdChangeLog entity.
func (acluo *AdChangeLogUpdateOne) Save(ctx context.Context) (*AdChangeLog, error) {
	var (
		err  error
		node *AdChangeLog
	)
	acluo.defaults()
	if len(acluo.hooks) == 0 {
		if err = acluo.check(); err != nil {
			return nil, err
		}
		node, err = acluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdChangeLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acluo.check(); err != nil {
				return nil, err
			}
			acluo.mutation = mutation
			node, err = acluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(acluo.hooks) - 1; i >= 0; i-- {
			if acluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (acluo *AdChangeLogUpdateOne) SaveX(ctx context.Context) *AdChangeLog {
	node, err := acluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acluo *AdChangeLogUpdateOne) Exec(ctx context.Context) error {
	_, err := acluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acluo *AdChangeLogUpdateOne) ExecX(ctx context.Context) {
	if err := acluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acluo *AdChangeLogUpdateOne) defaults() {
	if _, ok := acluo.mutation.UpdatedAt(); !ok {
		v := adchangelog.UpdateDefaultUpdatedAt()
		acluo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acluo *AdChangeLogUpdateOne) check() error {
	if _, ok := acluo.mutation.UserID(); acluo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (acluo *AdChangeLogUpdateOne) sqlSave(ctx context.Context) (_node *AdChangeLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adchangelog.Table,
			Columns: adchangelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: adchangelog.FieldID,
			},
		},
	}
	id, ok := acluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AdChangeLog.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := acluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adchangelog.FieldID)
		for _, f := range fields {
			if !adchangelog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adchangelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acluo.mutation.AdId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adchangelog.FieldAdId,
		})
	}
	if acluo.mutation.AdIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: adchangelog.FieldAdId,
		})
	}
	if value, ok := acluo.mutation.ChId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldChId,
		})
	}
	if value, ok := acluo.mutation.AddedChId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldChId,
		})
	}
	if acluo.mutation.ChIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: adchangelog.FieldChId,
		})
	}
	if value, ok := acluo.mutation.DeviceId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adchangelog.FieldDeviceId,
		})
	}
	if acluo.mutation.DeviceIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: adchangelog.FieldDeviceId,
		})
	}
	if value, ok := acluo.mutation.ExtInfo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adchangelog.FieldExtInfo,
		})
	}
	if acluo.mutation.ExtInfoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: adchangelog.FieldExtInfo,
		})
	}
	if value, ok := acluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: adchangelog.FieldUpdatedAt,
		})
	}
	if value, ok := acluo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldCreateBy,
		})
	}
	if value, ok := acluo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldCreateBy,
		})
	}
	if value, ok := acluo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldUpdateBy,
		})
	}
	if value, ok := acluo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldUpdateBy,
		})
	}
	if value, ok := acluo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldTenantId,
		})
	}
	if value, ok := acluo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchangelog.FieldTenantId,
		})
	}
	if acluo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adchangelog.UserTable,
			Columns: []string{adchangelog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: socialuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acluo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adchangelog.UserTable,
			Columns: []string{adchangelog.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AdChangeLog{config: acluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adchangelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
