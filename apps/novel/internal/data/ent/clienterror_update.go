// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/clienterror"
	"hope/apps/novel/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ClientErrorUpdate is the builder for updating ClientError entities.
type ClientErrorUpdate struct {
	config
	hooks    []Hook
	mutation *ClientErrorMutation
}

// Where appends a list predicates to the ClientErrorUpdate builder.
func (ceu *ClientErrorUpdate) Where(ps ...predicate.ClientError) *ClientErrorUpdate {
	ceu.mutation.Where(ps...)
	return ceu
}

// SetAppVersion sets the "appVersion" field.
func (ceu *ClientErrorUpdate) SetAppVersion(s string) *ClientErrorUpdate {
	ceu.mutation.SetAppVersion(s)
	return ceu
}

// SetNillableAppVersion sets the "appVersion" field if the given value is not nil.
func (ceu *ClientErrorUpdate) SetNillableAppVersion(s *string) *ClientErrorUpdate {
	if s != nil {
		ceu.SetAppVersion(*s)
	}
	return ceu
}

// ClearAppVersion clears the value of the "appVersion" field.
func (ceu *ClientErrorUpdate) ClearAppVersion() *ClientErrorUpdate {
	ceu.mutation.ClearAppVersion()
	return ceu
}

// SetDeviceName sets the "deviceName" field.
func (ceu *ClientErrorUpdate) SetDeviceName(s string) *ClientErrorUpdate {
	ceu.mutation.SetDeviceName(s)
	return ceu
}

// SetNillableDeviceName sets the "deviceName" field if the given value is not nil.
func (ceu *ClientErrorUpdate) SetNillableDeviceName(s *string) *ClientErrorUpdate {
	if s != nil {
		ceu.SetDeviceName(*s)
	}
	return ceu
}

// ClearDeviceName clears the value of the "deviceName" field.
func (ceu *ClientErrorUpdate) ClearDeviceName() *ClientErrorUpdate {
	ceu.mutation.ClearDeviceName()
	return ceu
}

// SetOsName sets the "osName" field.
func (ceu *ClientErrorUpdate) SetOsName(s string) *ClientErrorUpdate {
	ceu.mutation.SetOsName(s)
	return ceu
}

// SetNillableOsName sets the "osName" field if the given value is not nil.
func (ceu *ClientErrorUpdate) SetNillableOsName(s *string) *ClientErrorUpdate {
	if s != nil {
		ceu.SetOsName(*s)
	}
	return ceu
}

// ClearOsName clears the value of the "osName" field.
func (ceu *ClientErrorUpdate) ClearOsName() *ClientErrorUpdate {
	ceu.mutation.ClearOsName()
	return ceu
}

// SetErrorInfo sets the "errorInfo" field.
func (ceu *ClientErrorUpdate) SetErrorInfo(s string) *ClientErrorUpdate {
	ceu.mutation.SetErrorInfo(s)
	return ceu
}

// SetNillableErrorInfo sets the "errorInfo" field if the given value is not nil.
func (ceu *ClientErrorUpdate) SetNillableErrorInfo(s *string) *ClientErrorUpdate {
	if s != nil {
		ceu.SetErrorInfo(*s)
	}
	return ceu
}

// ClearErrorInfo clears the value of the "errorInfo" field.
func (ceu *ClientErrorUpdate) ClearErrorInfo() *ClientErrorUpdate {
	ceu.mutation.ClearErrorInfo()
	return ceu
}

// SetUserId sets the "userId" field.
func (ceu *ClientErrorUpdate) SetUserId(i int64) *ClientErrorUpdate {
	ceu.mutation.ResetUserId()
	ceu.mutation.SetUserId(i)
	return ceu
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (ceu *ClientErrorUpdate) SetNillableUserId(i *int64) *ClientErrorUpdate {
	if i != nil {
		ceu.SetUserId(*i)
	}
	return ceu
}

// AddUserId adds i to the "userId" field.
func (ceu *ClientErrorUpdate) AddUserId(i int64) *ClientErrorUpdate {
	ceu.mutation.AddUserId(i)
	return ceu
}

// ClearUserId clears the value of the "userId" field.
func (ceu *ClientErrorUpdate) ClearUserId() *ClientErrorUpdate {
	ceu.mutation.ClearUserId()
	return ceu
}

// SetUpdatedAt sets the "updatedAt" field.
func (ceu *ClientErrorUpdate) SetUpdatedAt(t time.Time) *ClientErrorUpdate {
	ceu.mutation.SetUpdatedAt(t)
	return ceu
}

// SetCreateBy sets the "createBy" field.
func (ceu *ClientErrorUpdate) SetCreateBy(i int64) *ClientErrorUpdate {
	ceu.mutation.ResetCreateBy()
	ceu.mutation.SetCreateBy(i)
	return ceu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (ceu *ClientErrorUpdate) SetNillableCreateBy(i *int64) *ClientErrorUpdate {
	if i != nil {
		ceu.SetCreateBy(*i)
	}
	return ceu
}

// AddCreateBy adds i to the "createBy" field.
func (ceu *ClientErrorUpdate) AddCreateBy(i int64) *ClientErrorUpdate {
	ceu.mutation.AddCreateBy(i)
	return ceu
}

// SetUpdateBy sets the "updateBy" field.
func (ceu *ClientErrorUpdate) SetUpdateBy(i int64) *ClientErrorUpdate {
	ceu.mutation.ResetUpdateBy()
	ceu.mutation.SetUpdateBy(i)
	return ceu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (ceu *ClientErrorUpdate) SetNillableUpdateBy(i *int64) *ClientErrorUpdate {
	if i != nil {
		ceu.SetUpdateBy(*i)
	}
	return ceu
}

// AddUpdateBy adds i to the "updateBy" field.
func (ceu *ClientErrorUpdate) AddUpdateBy(i int64) *ClientErrorUpdate {
	ceu.mutation.AddUpdateBy(i)
	return ceu
}

// SetTenantId sets the "tenantId" field.
func (ceu *ClientErrorUpdate) SetTenantId(i int64) *ClientErrorUpdate {
	ceu.mutation.ResetTenantId()
	ceu.mutation.SetTenantId(i)
	return ceu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (ceu *ClientErrorUpdate) SetNillableTenantId(i *int64) *ClientErrorUpdate {
	if i != nil {
		ceu.SetTenantId(*i)
	}
	return ceu
}

// AddTenantId adds i to the "tenantId" field.
func (ceu *ClientErrorUpdate) AddTenantId(i int64) *ClientErrorUpdate {
	ceu.mutation.AddTenantId(i)
	return ceu
}

// Mutation returns the ClientErrorMutation object of the builder.
func (ceu *ClientErrorUpdate) Mutation() *ClientErrorMutation {
	return ceu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ceu *ClientErrorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ceu.defaults()
	if len(ceu.hooks) == 0 {
		affected, err = ceu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClientErrorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ceu.mutation = mutation
			affected, err = ceu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ceu.hooks) - 1; i >= 0; i-- {
			if ceu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ceu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ceu *ClientErrorUpdate) SaveX(ctx context.Context) int {
	affected, err := ceu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ceu *ClientErrorUpdate) Exec(ctx context.Context) error {
	_, err := ceu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceu *ClientErrorUpdate) ExecX(ctx context.Context) {
	if err := ceu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ceu *ClientErrorUpdate) defaults() {
	if _, ok := ceu.mutation.UpdatedAt(); !ok {
		v := clienterror.UpdateDefaultUpdatedAt()
		ceu.mutation.SetUpdatedAt(v)
	}
}

func (ceu *ClientErrorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clienterror.Table,
			Columns: clienterror.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: clienterror.FieldID,
			},
		},
	}
	if ps := ceu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceu.mutation.AppVersion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldAppVersion,
		})
	}
	if ceu.mutation.AppVersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: clienterror.FieldAppVersion,
		})
	}
	if value, ok := ceu.mutation.DeviceName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldDeviceName,
		})
	}
	if ceu.mutation.DeviceNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: clienterror.FieldDeviceName,
		})
	}
	if value, ok := ceu.mutation.OsName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldOsName,
		})
	}
	if ceu.mutation.OsNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: clienterror.FieldOsName,
		})
	}
	if value, ok := ceu.mutation.ErrorInfo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldErrorInfo,
		})
	}
	if ceu.mutation.ErrorInfoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: clienterror.FieldErrorInfo,
		})
	}
	if value, ok := ceu.mutation.UserId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldUserId,
		})
	}
	if value, ok := ceu.mutation.AddedUserId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldUserId,
		})
	}
	if ceu.mutation.UserIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: clienterror.FieldUserId,
		})
	}
	if value, ok := ceu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clienterror.FieldUpdatedAt,
		})
	}
	if value, ok := ceu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldCreateBy,
		})
	}
	if value, ok := ceu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldCreateBy,
		})
	}
	if value, ok := ceu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldUpdateBy,
		})
	}
	if value, ok := ceu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldUpdateBy,
		})
	}
	if value, ok := ceu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldTenantId,
		})
	}
	if value, ok := ceu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldTenantId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ceu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clienterror.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ClientErrorUpdateOne is the builder for updating a single ClientError entity.
type ClientErrorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClientErrorMutation
}

// SetAppVersion sets the "appVersion" field.
func (ceuo *ClientErrorUpdateOne) SetAppVersion(s string) *ClientErrorUpdateOne {
	ceuo.mutation.SetAppVersion(s)
	return ceuo
}

// SetNillableAppVersion sets the "appVersion" field if the given value is not nil.
func (ceuo *ClientErrorUpdateOne) SetNillableAppVersion(s *string) *ClientErrorUpdateOne {
	if s != nil {
		ceuo.SetAppVersion(*s)
	}
	return ceuo
}

// ClearAppVersion clears the value of the "appVersion" field.
func (ceuo *ClientErrorUpdateOne) ClearAppVersion() *ClientErrorUpdateOne {
	ceuo.mutation.ClearAppVersion()
	return ceuo
}

// SetDeviceName sets the "deviceName" field.
func (ceuo *ClientErrorUpdateOne) SetDeviceName(s string) *ClientErrorUpdateOne {
	ceuo.mutation.SetDeviceName(s)
	return ceuo
}

// SetNillableDeviceName sets the "deviceName" field if the given value is not nil.
func (ceuo *ClientErrorUpdateOne) SetNillableDeviceName(s *string) *ClientErrorUpdateOne {
	if s != nil {
		ceuo.SetDeviceName(*s)
	}
	return ceuo
}

// ClearDeviceName clears the value of the "deviceName" field.
func (ceuo *ClientErrorUpdateOne) ClearDeviceName() *ClientErrorUpdateOne {
	ceuo.mutation.ClearDeviceName()
	return ceuo
}

// SetOsName sets the "osName" field.
func (ceuo *ClientErrorUpdateOne) SetOsName(s string) *ClientErrorUpdateOne {
	ceuo.mutation.SetOsName(s)
	return ceuo
}

// SetNillableOsName sets the "osName" field if the given value is not nil.
func (ceuo *ClientErrorUpdateOne) SetNillableOsName(s *string) *ClientErrorUpdateOne {
	if s != nil {
		ceuo.SetOsName(*s)
	}
	return ceuo
}

// ClearOsName clears the value of the "osName" field.
func (ceuo *ClientErrorUpdateOne) ClearOsName() *ClientErrorUpdateOne {
	ceuo.mutation.ClearOsName()
	return ceuo
}

// SetErrorInfo sets the "errorInfo" field.
func (ceuo *ClientErrorUpdateOne) SetErrorInfo(s string) *ClientErrorUpdateOne {
	ceuo.mutation.SetErrorInfo(s)
	return ceuo
}

// SetNillableErrorInfo sets the "errorInfo" field if the given value is not nil.
func (ceuo *ClientErrorUpdateOne) SetNillableErrorInfo(s *string) *ClientErrorUpdateOne {
	if s != nil {
		ceuo.SetErrorInfo(*s)
	}
	return ceuo
}

// ClearErrorInfo clears the value of the "errorInfo" field.
func (ceuo *ClientErrorUpdateOne) ClearErrorInfo() *ClientErrorUpdateOne {
	ceuo.mutation.ClearErrorInfo()
	return ceuo
}

// SetUserId sets the "userId" field.
func (ceuo *ClientErrorUpdateOne) SetUserId(i int64) *ClientErrorUpdateOne {
	ceuo.mutation.ResetUserId()
	ceuo.mutation.SetUserId(i)
	return ceuo
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (ceuo *ClientErrorUpdateOne) SetNillableUserId(i *int64) *ClientErrorUpdateOne {
	if i != nil {
		ceuo.SetUserId(*i)
	}
	return ceuo
}

// AddUserId adds i to the "userId" field.
func (ceuo *ClientErrorUpdateOne) AddUserId(i int64) *ClientErrorUpdateOne {
	ceuo.mutation.AddUserId(i)
	return ceuo
}

// ClearUserId clears the value of the "userId" field.
func (ceuo *ClientErrorUpdateOne) ClearUserId() *ClientErrorUpdateOne {
	ceuo.mutation.ClearUserId()
	return ceuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (ceuo *ClientErrorUpdateOne) SetUpdatedAt(t time.Time) *ClientErrorUpdateOne {
	ceuo.mutation.SetUpdatedAt(t)
	return ceuo
}

// SetCreateBy sets the "createBy" field.
func (ceuo *ClientErrorUpdateOne) SetCreateBy(i int64) *ClientErrorUpdateOne {
	ceuo.mutation.ResetCreateBy()
	ceuo.mutation.SetCreateBy(i)
	return ceuo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (ceuo *ClientErrorUpdateOne) SetNillableCreateBy(i *int64) *ClientErrorUpdateOne {
	if i != nil {
		ceuo.SetCreateBy(*i)
	}
	return ceuo
}

// AddCreateBy adds i to the "createBy" field.
func (ceuo *ClientErrorUpdateOne) AddCreateBy(i int64) *ClientErrorUpdateOne {
	ceuo.mutation.AddCreateBy(i)
	return ceuo
}

// SetUpdateBy sets the "updateBy" field.
func (ceuo *ClientErrorUpdateOne) SetUpdateBy(i int64) *ClientErrorUpdateOne {
	ceuo.mutation.ResetUpdateBy()
	ceuo.mutation.SetUpdateBy(i)
	return ceuo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (ceuo *ClientErrorUpdateOne) SetNillableUpdateBy(i *int64) *ClientErrorUpdateOne {
	if i != nil {
		ceuo.SetUpdateBy(*i)
	}
	return ceuo
}

// AddUpdateBy adds i to the "updateBy" field.
func (ceuo *ClientErrorUpdateOne) AddUpdateBy(i int64) *ClientErrorUpdateOne {
	ceuo.mutation.AddUpdateBy(i)
	return ceuo
}

// SetTenantId sets the "tenantId" field.
func (ceuo *ClientErrorUpdateOne) SetTenantId(i int64) *ClientErrorUpdateOne {
	ceuo.mutation.ResetTenantId()
	ceuo.mutation.SetTenantId(i)
	return ceuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (ceuo *ClientErrorUpdateOne) SetNillableTenantId(i *int64) *ClientErrorUpdateOne {
	if i != nil {
		ceuo.SetTenantId(*i)
	}
	return ceuo
}

// AddTenantId adds i to the "tenantId" field.
func (ceuo *ClientErrorUpdateOne) AddTenantId(i int64) *ClientErrorUpdateOne {
	ceuo.mutation.AddTenantId(i)
	return ceuo
}

// Mutation returns the ClientErrorMutation object of the builder.
func (ceuo *ClientErrorUpdateOne) Mutation() *ClientErrorMutation {
	return ceuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ceuo *ClientErrorUpdateOne) Select(field string, fields ...string) *ClientErrorUpdateOne {
	ceuo.fields = append([]string{field}, fields...)
	return ceuo
}

// Save executes the query and returns the updated ClientError entity.
func (ceuo *ClientErrorUpdateOne) Save(ctx context.Context) (*ClientError, error) {
	var (
		err  error
		node *ClientError
	)
	ceuo.defaults()
	if len(ceuo.hooks) == 0 {
		node, err = ceuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClientErrorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ceuo.mutation = mutation
			node, err = ceuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ceuo.hooks) - 1; i >= 0; i-- {
			if ceuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ceuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ceuo *ClientErrorUpdateOne) SaveX(ctx context.Context) *ClientError {
	node, err := ceuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ceuo *ClientErrorUpdateOne) Exec(ctx context.Context) error {
	_, err := ceuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceuo *ClientErrorUpdateOne) ExecX(ctx context.Context) {
	if err := ceuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ceuo *ClientErrorUpdateOne) defaults() {
	if _, ok := ceuo.mutation.UpdatedAt(); !ok {
		v := clienterror.UpdateDefaultUpdatedAt()
		ceuo.mutation.SetUpdatedAt(v)
	}
}

func (ceuo *ClientErrorUpdateOne) sqlSave(ctx context.Context) (_node *ClientError, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   clienterror.Table,
			Columns: clienterror.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: clienterror.FieldID,
			},
		},
	}
	id, ok := ceuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ClientError.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ceuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, clienterror.FieldID)
		for _, f := range fields {
			if !clienterror.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != clienterror.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ceuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceuo.mutation.AppVersion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldAppVersion,
		})
	}
	if ceuo.mutation.AppVersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: clienterror.FieldAppVersion,
		})
	}
	if value, ok := ceuo.mutation.DeviceName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldDeviceName,
		})
	}
	if ceuo.mutation.DeviceNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: clienterror.FieldDeviceName,
		})
	}
	if value, ok := ceuo.mutation.OsName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldOsName,
		})
	}
	if ceuo.mutation.OsNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: clienterror.FieldOsName,
		})
	}
	if value, ok := ceuo.mutation.ErrorInfo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldErrorInfo,
		})
	}
	if ceuo.mutation.ErrorInfoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: clienterror.FieldErrorInfo,
		})
	}
	if value, ok := ceuo.mutation.UserId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldUserId,
		})
	}
	if value, ok := ceuo.mutation.AddedUserId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldUserId,
		})
	}
	if ceuo.mutation.UserIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: clienterror.FieldUserId,
		})
	}
	if value, ok := ceuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clienterror.FieldUpdatedAt,
		})
	}
	if value, ok := ceuo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldCreateBy,
		})
	}
	if value, ok := ceuo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldCreateBy,
		})
	}
	if value, ok := ceuo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldUpdateBy,
		})
	}
	if value, ok := ceuo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldUpdateBy,
		})
	}
	if value, ok := ceuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldTenantId,
		})
	}
	if value, ok := ceuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldTenantId,
		})
	}
	_node = &ClientError{config: ceuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ceuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{clienterror.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
