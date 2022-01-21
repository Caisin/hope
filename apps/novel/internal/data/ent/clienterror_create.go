// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/clienterror"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ClientErrorCreate is the builder for creating a ClientError entity.
type ClientErrorCreate struct {
	config
	mutation *ClientErrorMutation
	hooks    []Hook
}

// SetAppVersion sets the "appVersion" field.
func (cec *ClientErrorCreate) SetAppVersion(s string) *ClientErrorCreate {
	cec.mutation.SetAppVersion(s)
	return cec
}

// SetNillableAppVersion sets the "appVersion" field if the given value is not nil.
func (cec *ClientErrorCreate) SetNillableAppVersion(s *string) *ClientErrorCreate {
	if s != nil {
		cec.SetAppVersion(*s)
	}
	return cec
}

// SetDeviceName sets the "deviceName" field.
func (cec *ClientErrorCreate) SetDeviceName(s string) *ClientErrorCreate {
	cec.mutation.SetDeviceName(s)
	return cec
}

// SetNillableDeviceName sets the "deviceName" field if the given value is not nil.
func (cec *ClientErrorCreate) SetNillableDeviceName(s *string) *ClientErrorCreate {
	if s != nil {
		cec.SetDeviceName(*s)
	}
	return cec
}

// SetOsName sets the "osName" field.
func (cec *ClientErrorCreate) SetOsName(s string) *ClientErrorCreate {
	cec.mutation.SetOsName(s)
	return cec
}

// SetNillableOsName sets the "osName" field if the given value is not nil.
func (cec *ClientErrorCreate) SetNillableOsName(s *string) *ClientErrorCreate {
	if s != nil {
		cec.SetOsName(*s)
	}
	return cec
}

// SetErrorInfo sets the "errorInfo" field.
func (cec *ClientErrorCreate) SetErrorInfo(s string) *ClientErrorCreate {
	cec.mutation.SetErrorInfo(s)
	return cec
}

// SetNillableErrorInfo sets the "errorInfo" field if the given value is not nil.
func (cec *ClientErrorCreate) SetNillableErrorInfo(s *string) *ClientErrorCreate {
	if s != nil {
		cec.SetErrorInfo(*s)
	}
	return cec
}

// SetUserId sets the "userId" field.
func (cec *ClientErrorCreate) SetUserId(i int64) *ClientErrorCreate {
	cec.mutation.SetUserId(i)
	return cec
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (cec *ClientErrorCreate) SetNillableUserId(i *int64) *ClientErrorCreate {
	if i != nil {
		cec.SetUserId(*i)
	}
	return cec
}

// SetCreatedAt sets the "createdAt" field.
func (cec *ClientErrorCreate) SetCreatedAt(t time.Time) *ClientErrorCreate {
	cec.mutation.SetCreatedAt(t)
	return cec
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (cec *ClientErrorCreate) SetNillableCreatedAt(t *time.Time) *ClientErrorCreate {
	if t != nil {
		cec.SetCreatedAt(*t)
	}
	return cec
}

// SetUpdatedAt sets the "updatedAt" field.
func (cec *ClientErrorCreate) SetUpdatedAt(t time.Time) *ClientErrorCreate {
	cec.mutation.SetUpdatedAt(t)
	return cec
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (cec *ClientErrorCreate) SetNillableUpdatedAt(t *time.Time) *ClientErrorCreate {
	if t != nil {
		cec.SetUpdatedAt(*t)
	}
	return cec
}

// SetCreateBy sets the "createBy" field.
func (cec *ClientErrorCreate) SetCreateBy(i int64) *ClientErrorCreate {
	cec.mutation.SetCreateBy(i)
	return cec
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (cec *ClientErrorCreate) SetNillableCreateBy(i *int64) *ClientErrorCreate {
	if i != nil {
		cec.SetCreateBy(*i)
	}
	return cec
}

// SetUpdateBy sets the "updateBy" field.
func (cec *ClientErrorCreate) SetUpdateBy(i int64) *ClientErrorCreate {
	cec.mutation.SetUpdateBy(i)
	return cec
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (cec *ClientErrorCreate) SetNillableUpdateBy(i *int64) *ClientErrorCreate {
	if i != nil {
		cec.SetUpdateBy(*i)
	}
	return cec
}

// SetTenantId sets the "tenantId" field.
func (cec *ClientErrorCreate) SetTenantId(i int64) *ClientErrorCreate {
	cec.mutation.SetTenantId(i)
	return cec
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (cec *ClientErrorCreate) SetNillableTenantId(i *int64) *ClientErrorCreate {
	if i != nil {
		cec.SetTenantId(*i)
	}
	return cec
}

// Mutation returns the ClientErrorMutation object of the builder.
func (cec *ClientErrorCreate) Mutation() *ClientErrorMutation {
	return cec.mutation
}

// Save creates the ClientError in the database.
func (cec *ClientErrorCreate) Save(ctx context.Context) (*ClientError, error) {
	var (
		err  error
		node *ClientError
	)
	cec.defaults()
	if len(cec.hooks) == 0 {
		if err = cec.check(); err != nil {
			return nil, err
		}
		node, err = cec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClientErrorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cec.check(); err != nil {
				return nil, err
			}
			cec.mutation = mutation
			if node, err = cec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cec.hooks) - 1; i >= 0; i-- {
			if cec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cec *ClientErrorCreate) SaveX(ctx context.Context) *ClientError {
	v, err := cec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cec *ClientErrorCreate) Exec(ctx context.Context) error {
	_, err := cec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cec *ClientErrorCreate) ExecX(ctx context.Context) {
	if err := cec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cec *ClientErrorCreate) defaults() {
	if _, ok := cec.mutation.CreatedAt(); !ok {
		v := clienterror.DefaultCreatedAt()
		cec.mutation.SetCreatedAt(v)
	}
	if _, ok := cec.mutation.UpdatedAt(); !ok {
		v := clienterror.DefaultUpdatedAt()
		cec.mutation.SetUpdatedAt(v)
	}
	if _, ok := cec.mutation.CreateBy(); !ok {
		v := clienterror.DefaultCreateBy
		cec.mutation.SetCreateBy(v)
	}
	if _, ok := cec.mutation.UpdateBy(); !ok {
		v := clienterror.DefaultUpdateBy
		cec.mutation.SetUpdateBy(v)
	}
	if _, ok := cec.mutation.TenantId(); !ok {
		v := clienterror.DefaultTenantId
		cec.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cec *ClientErrorCreate) check() error {
	if _, ok := cec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := cec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := cec.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := cec.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := cec.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (cec *ClientErrorCreate) sqlSave(ctx context.Context) (*ClientError, error) {
	_node, _spec := cec.createSpec()
	if err := sqlgraph.CreateNode(ctx, cec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (cec *ClientErrorCreate) createSpec() (*ClientError, *sqlgraph.CreateSpec) {
	var (
		_node = &ClientError{config: cec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: clienterror.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: clienterror.FieldID,
			},
		}
	)
	if value, ok := cec.mutation.AppVersion(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldAppVersion,
		})
		_node.AppVersion = value
	}
	if value, ok := cec.mutation.DeviceName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldDeviceName,
		})
		_node.DeviceName = value
	}
	if value, ok := cec.mutation.OsName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldOsName,
		})
		_node.OsName = value
	}
	if value, ok := cec.mutation.ErrorInfo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: clienterror.FieldErrorInfo,
		})
		_node.ErrorInfo = value
	}
	if value, ok := cec.mutation.UserId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldUserId,
		})
		_node.UserId = value
	}
	if value, ok := cec.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clienterror.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cec.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: clienterror.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cec.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := cec.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := cec.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: clienterror.FieldTenantId,
		})
		_node.TenantId = value
	}
	return _node, _spec
}

// ClientErrorCreateBulk is the builder for creating many ClientError entities in bulk.
type ClientErrorCreateBulk struct {
	config
	builders []*ClientErrorCreate
}

// Save creates the ClientError entities in the database.
func (cecb *ClientErrorCreateBulk) Save(ctx context.Context) ([]*ClientError, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cecb.builders))
	nodes := make([]*ClientError, len(cecb.builders))
	mutators := make([]Mutator, len(cecb.builders))
	for i := range cecb.builders {
		func(i int, root context.Context) {
			builder := cecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClientErrorMutation)
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
					_, err = mutators[i+1].Mutate(root, cecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cecb *ClientErrorCreateBulk) SaveX(ctx context.Context) []*ClientError {
	v, err := cecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cecb *ClientErrorCreateBulk) Exec(ctx context.Context) error {
	_, err := cecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cecb *ClientErrorCreateBulk) ExecX(ctx context.Context) {
	if err := cecb.Exec(ctx); err != nil {
		panic(err)
	}
}