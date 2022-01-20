// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/qiniuconfig"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QiniuConfigUpdate is the builder for updating QiniuConfig entities.
type QiniuConfigUpdate struct {
	config
	hooks    []Hook
	mutation *QiniuConfigMutation
}

// Where appends a list predicates to the QiniuConfigUpdate builder.
func (qcu *QiniuConfigUpdate) Where(ps ...predicate.QiniuConfig) *QiniuConfigUpdate {
	qcu.mutation.Where(ps...)
	return qcu
}

// SetAccessKey sets the "accessKey" field.
func (qcu *QiniuConfigUpdate) SetAccessKey(s string) *QiniuConfigUpdate {
	qcu.mutation.SetAccessKey(s)
	return qcu
}

// SetNillableAccessKey sets the "accessKey" field if the given value is not nil.
func (qcu *QiniuConfigUpdate) SetNillableAccessKey(s *string) *QiniuConfigUpdate {
	if s != nil {
		qcu.SetAccessKey(*s)
	}
	return qcu
}

// ClearAccessKey clears the value of the "accessKey" field.
func (qcu *QiniuConfigUpdate) ClearAccessKey() *QiniuConfigUpdate {
	qcu.mutation.ClearAccessKey()
	return qcu
}

// SetBucket sets the "bucket" field.
func (qcu *QiniuConfigUpdate) SetBucket(s string) *QiniuConfigUpdate {
	qcu.mutation.SetBucket(s)
	return qcu
}

// SetNillableBucket sets the "bucket" field if the given value is not nil.
func (qcu *QiniuConfigUpdate) SetNillableBucket(s *string) *QiniuConfigUpdate {
	if s != nil {
		qcu.SetBucket(*s)
	}
	return qcu
}

// ClearBucket clears the value of the "bucket" field.
func (qcu *QiniuConfigUpdate) ClearBucket() *QiniuConfigUpdate {
	qcu.mutation.ClearBucket()
	return qcu
}

// SetHost sets the "host" field.
func (qcu *QiniuConfigUpdate) SetHost(s string) *QiniuConfigUpdate {
	qcu.mutation.SetHost(s)
	return qcu
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (qcu *QiniuConfigUpdate) SetNillableHost(s *string) *QiniuConfigUpdate {
	if s != nil {
		qcu.SetHost(*s)
	}
	return qcu
}

// ClearHost clears the value of the "host" field.
func (qcu *QiniuConfigUpdate) ClearHost() *QiniuConfigUpdate {
	qcu.mutation.ClearHost()
	return qcu
}

// SetSecretKey sets the "secretKey" field.
func (qcu *QiniuConfigUpdate) SetSecretKey(s string) *QiniuConfigUpdate {
	qcu.mutation.SetSecretKey(s)
	return qcu
}

// SetNillableSecretKey sets the "secretKey" field if the given value is not nil.
func (qcu *QiniuConfigUpdate) SetNillableSecretKey(s *string) *QiniuConfigUpdate {
	if s != nil {
		qcu.SetSecretKey(*s)
	}
	return qcu
}

// ClearSecretKey clears the value of the "secretKey" field.
func (qcu *QiniuConfigUpdate) ClearSecretKey() *QiniuConfigUpdate {
	qcu.mutation.ClearSecretKey()
	return qcu
}

// SetType sets the "type" field.
func (qcu *QiniuConfigUpdate) SetType(s string) *QiniuConfigUpdate {
	qcu.mutation.SetType(s)
	return qcu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (qcu *QiniuConfigUpdate) SetNillableType(s *string) *QiniuConfigUpdate {
	if s != nil {
		qcu.SetType(*s)
	}
	return qcu
}

// ClearType clears the value of the "type" field.
func (qcu *QiniuConfigUpdate) ClearType() *QiniuConfigUpdate {
	qcu.mutation.ClearType()
	return qcu
}

// SetZone sets the "zone" field.
func (qcu *QiniuConfigUpdate) SetZone(s string) *QiniuConfigUpdate {
	qcu.mutation.SetZone(s)
	return qcu
}

// SetNillableZone sets the "zone" field if the given value is not nil.
func (qcu *QiniuConfigUpdate) SetNillableZone(s *string) *QiniuConfigUpdate {
	if s != nil {
		qcu.SetZone(*s)
	}
	return qcu
}

// ClearZone clears the value of the "zone" field.
func (qcu *QiniuConfigUpdate) ClearZone() *QiniuConfigUpdate {
	qcu.mutation.ClearZone()
	return qcu
}

// SetUpdatedAt sets the "updatedAt" field.
func (qcu *QiniuConfigUpdate) SetUpdatedAt(t time.Time) *QiniuConfigUpdate {
	qcu.mutation.SetUpdatedAt(t)
	return qcu
}

// SetCreateBy sets the "createBy" field.
func (qcu *QiniuConfigUpdate) SetCreateBy(i int64) *QiniuConfigUpdate {
	qcu.mutation.ResetCreateBy()
	qcu.mutation.SetCreateBy(i)
	return qcu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (qcu *QiniuConfigUpdate) SetNillableCreateBy(i *int64) *QiniuConfigUpdate {
	if i != nil {
		qcu.SetCreateBy(*i)
	}
	return qcu
}

// AddCreateBy adds i to the "createBy" field.
func (qcu *QiniuConfigUpdate) AddCreateBy(i int64) *QiniuConfigUpdate {
	qcu.mutation.AddCreateBy(i)
	return qcu
}

// SetUpdateBy sets the "updateBy" field.
func (qcu *QiniuConfigUpdate) SetUpdateBy(i int64) *QiniuConfigUpdate {
	qcu.mutation.ResetUpdateBy()
	qcu.mutation.SetUpdateBy(i)
	return qcu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (qcu *QiniuConfigUpdate) SetNillableUpdateBy(i *int64) *QiniuConfigUpdate {
	if i != nil {
		qcu.SetUpdateBy(*i)
	}
	return qcu
}

// AddUpdateBy adds i to the "updateBy" field.
func (qcu *QiniuConfigUpdate) AddUpdateBy(i int64) *QiniuConfigUpdate {
	qcu.mutation.AddUpdateBy(i)
	return qcu
}

// SetTenantId sets the "tenantId" field.
func (qcu *QiniuConfigUpdate) SetTenantId(i int64) *QiniuConfigUpdate {
	qcu.mutation.ResetTenantId()
	qcu.mutation.SetTenantId(i)
	return qcu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (qcu *QiniuConfigUpdate) SetNillableTenantId(i *int64) *QiniuConfigUpdate {
	if i != nil {
		qcu.SetTenantId(*i)
	}
	return qcu
}

// AddTenantId adds i to the "tenantId" field.
func (qcu *QiniuConfigUpdate) AddTenantId(i int64) *QiniuConfigUpdate {
	qcu.mutation.AddTenantId(i)
	return qcu
}

// Mutation returns the QiniuConfigMutation object of the builder.
func (qcu *QiniuConfigUpdate) Mutation() *QiniuConfigMutation {
	return qcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qcu *QiniuConfigUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	qcu.defaults()
	if len(qcu.hooks) == 0 {
		affected, err = qcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QiniuConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qcu.mutation = mutation
			affected, err = qcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qcu.hooks) - 1; i >= 0; i-- {
			if qcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (qcu *QiniuConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := qcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qcu *QiniuConfigUpdate) Exec(ctx context.Context) error {
	_, err := qcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcu *QiniuConfigUpdate) ExecX(ctx context.Context) {
	if err := qcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qcu *QiniuConfigUpdate) defaults() {
	if _, ok := qcu.mutation.UpdatedAt(); !ok {
		v := qiniuconfig.UpdateDefaultUpdatedAt()
		qcu.mutation.SetUpdatedAt(v)
	}
}

func (qcu *QiniuConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   qiniuconfig.Table,
			Columns: qiniuconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: qiniuconfig.FieldID,
			},
		},
	}
	if ps := qcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qcu.mutation.AccessKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldAccessKey,
		})
	}
	if qcu.mutation.AccessKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldAccessKey,
		})
	}
	if value, ok := qcu.mutation.Bucket(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldBucket,
		})
	}
	if qcu.mutation.BucketCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldBucket,
		})
	}
	if value, ok := qcu.mutation.Host(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldHost,
		})
	}
	if qcu.mutation.HostCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldHost,
		})
	}
	if value, ok := qcu.mutation.SecretKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldSecretKey,
		})
	}
	if qcu.mutation.SecretKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldSecretKey,
		})
	}
	if value, ok := qcu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldType,
		})
	}
	if qcu.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldType,
		})
	}
	if value, ok := qcu.mutation.Zone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldZone,
		})
	}
	if qcu.mutation.ZoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldZone,
		})
	}
	if value, ok := qcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: qiniuconfig.FieldUpdatedAt,
		})
	}
	if value, ok := qcu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldCreateBy,
		})
	}
	if value, ok := qcu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldCreateBy,
		})
	}
	if value, ok := qcu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldUpdateBy,
		})
	}
	if value, ok := qcu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldUpdateBy,
		})
	}
	if value, ok := qcu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldTenantId,
		})
	}
	if value, ok := qcu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldTenantId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{qiniuconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// QiniuConfigUpdateOne is the builder for updating a single QiniuConfig entity.
type QiniuConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *QiniuConfigMutation
}

// SetAccessKey sets the "accessKey" field.
func (qcuo *QiniuConfigUpdateOne) SetAccessKey(s string) *QiniuConfigUpdateOne {
	qcuo.mutation.SetAccessKey(s)
	return qcuo
}

// SetNillableAccessKey sets the "accessKey" field if the given value is not nil.
func (qcuo *QiniuConfigUpdateOne) SetNillableAccessKey(s *string) *QiniuConfigUpdateOne {
	if s != nil {
		qcuo.SetAccessKey(*s)
	}
	return qcuo
}

// ClearAccessKey clears the value of the "accessKey" field.
func (qcuo *QiniuConfigUpdateOne) ClearAccessKey() *QiniuConfigUpdateOne {
	qcuo.mutation.ClearAccessKey()
	return qcuo
}

// SetBucket sets the "bucket" field.
func (qcuo *QiniuConfigUpdateOne) SetBucket(s string) *QiniuConfigUpdateOne {
	qcuo.mutation.SetBucket(s)
	return qcuo
}

// SetNillableBucket sets the "bucket" field if the given value is not nil.
func (qcuo *QiniuConfigUpdateOne) SetNillableBucket(s *string) *QiniuConfigUpdateOne {
	if s != nil {
		qcuo.SetBucket(*s)
	}
	return qcuo
}

// ClearBucket clears the value of the "bucket" field.
func (qcuo *QiniuConfigUpdateOne) ClearBucket() *QiniuConfigUpdateOne {
	qcuo.mutation.ClearBucket()
	return qcuo
}

// SetHost sets the "host" field.
func (qcuo *QiniuConfigUpdateOne) SetHost(s string) *QiniuConfigUpdateOne {
	qcuo.mutation.SetHost(s)
	return qcuo
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (qcuo *QiniuConfigUpdateOne) SetNillableHost(s *string) *QiniuConfigUpdateOne {
	if s != nil {
		qcuo.SetHost(*s)
	}
	return qcuo
}

// ClearHost clears the value of the "host" field.
func (qcuo *QiniuConfigUpdateOne) ClearHost() *QiniuConfigUpdateOne {
	qcuo.mutation.ClearHost()
	return qcuo
}

// SetSecretKey sets the "secretKey" field.
func (qcuo *QiniuConfigUpdateOne) SetSecretKey(s string) *QiniuConfigUpdateOne {
	qcuo.mutation.SetSecretKey(s)
	return qcuo
}

// SetNillableSecretKey sets the "secretKey" field if the given value is not nil.
func (qcuo *QiniuConfigUpdateOne) SetNillableSecretKey(s *string) *QiniuConfigUpdateOne {
	if s != nil {
		qcuo.SetSecretKey(*s)
	}
	return qcuo
}

// ClearSecretKey clears the value of the "secretKey" field.
func (qcuo *QiniuConfigUpdateOne) ClearSecretKey() *QiniuConfigUpdateOne {
	qcuo.mutation.ClearSecretKey()
	return qcuo
}

// SetType sets the "type" field.
func (qcuo *QiniuConfigUpdateOne) SetType(s string) *QiniuConfigUpdateOne {
	qcuo.mutation.SetType(s)
	return qcuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (qcuo *QiniuConfigUpdateOne) SetNillableType(s *string) *QiniuConfigUpdateOne {
	if s != nil {
		qcuo.SetType(*s)
	}
	return qcuo
}

// ClearType clears the value of the "type" field.
func (qcuo *QiniuConfigUpdateOne) ClearType() *QiniuConfigUpdateOne {
	qcuo.mutation.ClearType()
	return qcuo
}

// SetZone sets the "zone" field.
func (qcuo *QiniuConfigUpdateOne) SetZone(s string) *QiniuConfigUpdateOne {
	qcuo.mutation.SetZone(s)
	return qcuo
}

// SetNillableZone sets the "zone" field if the given value is not nil.
func (qcuo *QiniuConfigUpdateOne) SetNillableZone(s *string) *QiniuConfigUpdateOne {
	if s != nil {
		qcuo.SetZone(*s)
	}
	return qcuo
}

// ClearZone clears the value of the "zone" field.
func (qcuo *QiniuConfigUpdateOne) ClearZone() *QiniuConfigUpdateOne {
	qcuo.mutation.ClearZone()
	return qcuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (qcuo *QiniuConfigUpdateOne) SetUpdatedAt(t time.Time) *QiniuConfigUpdateOne {
	qcuo.mutation.SetUpdatedAt(t)
	return qcuo
}

// SetCreateBy sets the "createBy" field.
func (qcuo *QiniuConfigUpdateOne) SetCreateBy(i int64) *QiniuConfigUpdateOne {
	qcuo.mutation.ResetCreateBy()
	qcuo.mutation.SetCreateBy(i)
	return qcuo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (qcuo *QiniuConfigUpdateOne) SetNillableCreateBy(i *int64) *QiniuConfigUpdateOne {
	if i != nil {
		qcuo.SetCreateBy(*i)
	}
	return qcuo
}

// AddCreateBy adds i to the "createBy" field.
func (qcuo *QiniuConfigUpdateOne) AddCreateBy(i int64) *QiniuConfigUpdateOne {
	qcuo.mutation.AddCreateBy(i)
	return qcuo
}

// SetUpdateBy sets the "updateBy" field.
func (qcuo *QiniuConfigUpdateOne) SetUpdateBy(i int64) *QiniuConfigUpdateOne {
	qcuo.mutation.ResetUpdateBy()
	qcuo.mutation.SetUpdateBy(i)
	return qcuo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (qcuo *QiniuConfigUpdateOne) SetNillableUpdateBy(i *int64) *QiniuConfigUpdateOne {
	if i != nil {
		qcuo.SetUpdateBy(*i)
	}
	return qcuo
}

// AddUpdateBy adds i to the "updateBy" field.
func (qcuo *QiniuConfigUpdateOne) AddUpdateBy(i int64) *QiniuConfigUpdateOne {
	qcuo.mutation.AddUpdateBy(i)
	return qcuo
}

// SetTenantId sets the "tenantId" field.
func (qcuo *QiniuConfigUpdateOne) SetTenantId(i int64) *QiniuConfigUpdateOne {
	qcuo.mutation.ResetTenantId()
	qcuo.mutation.SetTenantId(i)
	return qcuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (qcuo *QiniuConfigUpdateOne) SetNillableTenantId(i *int64) *QiniuConfigUpdateOne {
	if i != nil {
		qcuo.SetTenantId(*i)
	}
	return qcuo
}

// AddTenantId adds i to the "tenantId" field.
func (qcuo *QiniuConfigUpdateOne) AddTenantId(i int64) *QiniuConfigUpdateOne {
	qcuo.mutation.AddTenantId(i)
	return qcuo
}

// Mutation returns the QiniuConfigMutation object of the builder.
func (qcuo *QiniuConfigUpdateOne) Mutation() *QiniuConfigMutation {
	return qcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (qcuo *QiniuConfigUpdateOne) Select(field string, fields ...string) *QiniuConfigUpdateOne {
	qcuo.fields = append([]string{field}, fields...)
	return qcuo
}

// Save executes the query and returns the updated QiniuConfig entity.
func (qcuo *QiniuConfigUpdateOne) Save(ctx context.Context) (*QiniuConfig, error) {
	var (
		err  error
		node *QiniuConfig
	)
	qcuo.defaults()
	if len(qcuo.hooks) == 0 {
		node, err = qcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QiniuConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qcuo.mutation = mutation
			node, err = qcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qcuo.hooks) - 1; i >= 0; i-- {
			if qcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (qcuo *QiniuConfigUpdateOne) SaveX(ctx context.Context) *QiniuConfig {
	node, err := qcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (qcuo *QiniuConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := qcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcuo *QiniuConfigUpdateOne) ExecX(ctx context.Context) {
	if err := qcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qcuo *QiniuConfigUpdateOne) defaults() {
	if _, ok := qcuo.mutation.UpdatedAt(); !ok {
		v := qiniuconfig.UpdateDefaultUpdatedAt()
		qcuo.mutation.SetUpdatedAt(v)
	}
}

func (qcuo *QiniuConfigUpdateOne) sqlSave(ctx context.Context) (_node *QiniuConfig, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   qiniuconfig.Table,
			Columns: qiniuconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: qiniuconfig.FieldID,
			},
		},
	}
	id, ok := qcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing QiniuConfig.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := qcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, qiniuconfig.FieldID)
		for _, f := range fields {
			if !qiniuconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != qiniuconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := qcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qcuo.mutation.AccessKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldAccessKey,
		})
	}
	if qcuo.mutation.AccessKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldAccessKey,
		})
	}
	if value, ok := qcuo.mutation.Bucket(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldBucket,
		})
	}
	if qcuo.mutation.BucketCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldBucket,
		})
	}
	if value, ok := qcuo.mutation.Host(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldHost,
		})
	}
	if qcuo.mutation.HostCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldHost,
		})
	}
	if value, ok := qcuo.mutation.SecretKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldSecretKey,
		})
	}
	if qcuo.mutation.SecretKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldSecretKey,
		})
	}
	if value, ok := qcuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldType,
		})
	}
	if qcuo.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldType,
		})
	}
	if value, ok := qcuo.mutation.Zone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: qiniuconfig.FieldZone,
		})
	}
	if qcuo.mutation.ZoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: qiniuconfig.FieldZone,
		})
	}
	if value, ok := qcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: qiniuconfig.FieldUpdatedAt,
		})
	}
	if value, ok := qcuo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldCreateBy,
		})
	}
	if value, ok := qcuo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldCreateBy,
		})
	}
	if value, ok := qcuo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldUpdateBy,
		})
	}
	if value, ok := qcuo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldUpdateBy,
		})
	}
	if value, ok := qcuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldTenantId,
		})
	}
	if value, ok := qcuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: qiniuconfig.FieldTenantId,
		})
	}
	_node = &QiniuConfig{config: qcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, qcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{qiniuconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
