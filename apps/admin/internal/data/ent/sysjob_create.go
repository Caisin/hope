// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/sysjob"
	"hope/apps/admin/internal/data/ent/sysjoblog"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysJobCreate is the builder for creating a SysJob entity.
type SysJobCreate struct {
	config
	mutation *SysJobMutation
	hooks    []Hook
}

// SetJobName sets the "jobName" field.
func (sjc *SysJobCreate) SetJobName(s string) *SysJobCreate {
	sjc.mutation.SetJobName(s)
	return sjc
}

// SetNillableJobName sets the "jobName" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableJobName(s *string) *SysJobCreate {
	if s != nil {
		sjc.SetJobName(*s)
	}
	return sjc
}

// SetJobGroup sets the "jobGroup" field.
func (sjc *SysJobCreate) SetJobGroup(s string) *SysJobCreate {
	sjc.mutation.SetJobGroup(s)
	return sjc
}

// SetNillableJobGroup sets the "jobGroup" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableJobGroup(s *string) *SysJobCreate {
	if s != nil {
		sjc.SetJobGroup(*s)
	}
	return sjc
}

// SetJobType sets the "jobType" field.
func (sjc *SysJobCreate) SetJobType(i int32) *SysJobCreate {
	sjc.mutation.SetJobType(i)
	return sjc
}

// SetNillableJobType sets the "jobType" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableJobType(i *int32) *SysJobCreate {
	if i != nil {
		sjc.SetJobType(*i)
	}
	return sjc
}

// SetCronExpression sets the "cronExpression" field.
func (sjc *SysJobCreate) SetCronExpression(s string) *SysJobCreate {
	sjc.mutation.SetCronExpression(s)
	return sjc
}

// SetNillableCronExpression sets the "cronExpression" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableCronExpression(s *string) *SysJobCreate {
	if s != nil {
		sjc.SetCronExpression(*s)
	}
	return sjc
}

// SetInvokeTarget sets the "invokeTarget" field.
func (sjc *SysJobCreate) SetInvokeTarget(s string) *SysJobCreate {
	sjc.mutation.SetInvokeTarget(s)
	return sjc
}

// SetNillableInvokeTarget sets the "invokeTarget" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableInvokeTarget(s *string) *SysJobCreate {
	if s != nil {
		sjc.SetInvokeTarget(*s)
	}
	return sjc
}

// SetArgs sets the "args" field.
func (sjc *SysJobCreate) SetArgs(s string) *SysJobCreate {
	sjc.mutation.SetArgs(s)
	return sjc
}

// SetNillableArgs sets the "args" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableArgs(s *string) *SysJobCreate {
	if s != nil {
		sjc.SetArgs(*s)
	}
	return sjc
}

// SetExecPolicy sets the "execPolicy" field.
func (sjc *SysJobCreate) SetExecPolicy(i int32) *SysJobCreate {
	sjc.mutation.SetExecPolicy(i)
	return sjc
}

// SetNillableExecPolicy sets the "execPolicy" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableExecPolicy(i *int32) *SysJobCreate {
	if i != nil {
		sjc.SetExecPolicy(*i)
	}
	return sjc
}

// SetConcurrent sets the "concurrent" field.
func (sjc *SysJobCreate) SetConcurrent(i int32) *SysJobCreate {
	sjc.mutation.SetConcurrent(i)
	return sjc
}

// SetNillableConcurrent sets the "concurrent" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableConcurrent(i *int32) *SysJobCreate {
	if i != nil {
		sjc.SetConcurrent(*i)
	}
	return sjc
}

// SetState sets the "state" field.
func (sjc *SysJobCreate) SetState(s sysjob.State) *SysJobCreate {
	sjc.mutation.SetState(s)
	return sjc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableState(s *sysjob.State) *SysJobCreate {
	if s != nil {
		sjc.SetState(*s)
	}
	return sjc
}

// SetEntryId sets the "entryId" field.
func (sjc *SysJobCreate) SetEntryId(i int32) *SysJobCreate {
	sjc.mutation.SetEntryId(i)
	return sjc
}

// SetNillableEntryId sets the "entryId" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableEntryId(i *int32) *SysJobCreate {
	if i != nil {
		sjc.SetEntryId(*i)
	}
	return sjc
}

// SetCreatedAt sets the "createdAt" field.
func (sjc *SysJobCreate) SetCreatedAt(t time.Time) *SysJobCreate {
	sjc.mutation.SetCreatedAt(t)
	return sjc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableCreatedAt(t *time.Time) *SysJobCreate {
	if t != nil {
		sjc.SetCreatedAt(*t)
	}
	return sjc
}

// SetUpdatedAt sets the "updatedAt" field.
func (sjc *SysJobCreate) SetUpdatedAt(t time.Time) *SysJobCreate {
	sjc.mutation.SetUpdatedAt(t)
	return sjc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableUpdatedAt(t *time.Time) *SysJobCreate {
	if t != nil {
		sjc.SetUpdatedAt(*t)
	}
	return sjc
}

// SetCreateBy sets the "createBy" field.
func (sjc *SysJobCreate) SetCreateBy(i int64) *SysJobCreate {
	sjc.mutation.SetCreateBy(i)
	return sjc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableCreateBy(i *int64) *SysJobCreate {
	if i != nil {
		sjc.SetCreateBy(*i)
	}
	return sjc
}

// SetUpdateBy sets the "updateBy" field.
func (sjc *SysJobCreate) SetUpdateBy(i int64) *SysJobCreate {
	sjc.mutation.SetUpdateBy(i)
	return sjc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableUpdateBy(i *int64) *SysJobCreate {
	if i != nil {
		sjc.SetUpdateBy(*i)
	}
	return sjc
}

// SetTenantId sets the "tenantId" field.
func (sjc *SysJobCreate) SetTenantId(i int64) *SysJobCreate {
	sjc.mutation.SetTenantId(i)
	return sjc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (sjc *SysJobCreate) SetNillableTenantId(i *int64) *SysJobCreate {
	if i != nil {
		sjc.SetTenantId(*i)
	}
	return sjc
}

// AddLogIDs adds the "logs" edge to the SysJobLog entity by IDs.
func (sjc *SysJobCreate) AddLogIDs(ids ...int64) *SysJobCreate {
	sjc.mutation.AddLogIDs(ids...)
	return sjc
}

// AddLogs adds the "logs" edges to the SysJobLog entity.
func (sjc *SysJobCreate) AddLogs(s ...*SysJobLog) *SysJobCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sjc.AddLogIDs(ids...)
}

// Mutation returns the SysJobMutation object of the builder.
func (sjc *SysJobCreate) Mutation() *SysJobMutation {
	return sjc.mutation
}

// Save creates the SysJob in the database.
func (sjc *SysJobCreate) Save(ctx context.Context) (*SysJob, error) {
	var (
		err  error
		node *SysJob
	)
	sjc.defaults()
	if len(sjc.hooks) == 0 {
		if err = sjc.check(); err != nil {
			return nil, err
		}
		node, err = sjc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysJobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sjc.check(); err != nil {
				return nil, err
			}
			sjc.mutation = mutation
			if node, err = sjc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sjc.hooks) - 1; i >= 0; i-- {
			if sjc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sjc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sjc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sjc *SysJobCreate) SaveX(ctx context.Context) *SysJob {
	v, err := sjc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sjc *SysJobCreate) Exec(ctx context.Context) error {
	_, err := sjc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sjc *SysJobCreate) ExecX(ctx context.Context) {
	if err := sjc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sjc *SysJobCreate) defaults() {
	if _, ok := sjc.mutation.State(); !ok {
		v := sysjob.DefaultState
		sjc.mutation.SetState(v)
	}
	if _, ok := sjc.mutation.CreatedAt(); !ok {
		v := sysjob.DefaultCreatedAt()
		sjc.mutation.SetCreatedAt(v)
	}
	if _, ok := sjc.mutation.UpdatedAt(); !ok {
		v := sysjob.DefaultUpdatedAt()
		sjc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sjc.mutation.CreateBy(); !ok {
		v := sysjob.DefaultCreateBy
		sjc.mutation.SetCreateBy(v)
	}
	if _, ok := sjc.mutation.UpdateBy(); !ok {
		v := sysjob.DefaultUpdateBy
		sjc.mutation.SetUpdateBy(v)
	}
	if _, ok := sjc.mutation.TenantId(); !ok {
		v := sysjob.DefaultTenantId
		sjc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sjc *SysJobCreate) check() error {
	if _, ok := sjc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "SysJob.state"`)}
	}
	if v, ok := sjc.mutation.State(); ok {
		if err := sysjob.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "SysJob.state": %w`, err)}
		}
	}
	if _, ok := sjc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "SysJob.createdAt"`)}
	}
	if _, ok := sjc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "SysJob.updatedAt"`)}
	}
	if _, ok := sjc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "SysJob.createBy"`)}
	}
	if _, ok := sjc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "SysJob.updateBy"`)}
	}
	if _, ok := sjc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "SysJob.tenantId"`)}
	}
	return nil
}

func (sjc *SysJobCreate) sqlSave(ctx context.Context) (*SysJob, error) {
	_node, _spec := sjc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sjc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (sjc *SysJobCreate) createSpec() (*SysJob, *sqlgraph.CreateSpec) {
	var (
		_node = &SysJob{config: sjc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysjob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysjob.FieldID,
			},
		}
	)
	if value, ok := sjc.mutation.JobName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysjob.FieldJobName,
		})
		_node.JobName = value
	}
	if value, ok := sjc.mutation.JobGroup(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysjob.FieldJobGroup,
		})
		_node.JobGroup = value
	}
	if value, ok := sjc.mutation.JobType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysjob.FieldJobType,
		})
		_node.JobType = value
	}
	if value, ok := sjc.mutation.CronExpression(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysjob.FieldCronExpression,
		})
		_node.CronExpression = value
	}
	if value, ok := sjc.mutation.InvokeTarget(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysjob.FieldInvokeTarget,
		})
		_node.InvokeTarget = value
	}
	if value, ok := sjc.mutation.Args(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysjob.FieldArgs,
		})
		_node.Args = value
	}
	if value, ok := sjc.mutation.ExecPolicy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysjob.FieldExecPolicy,
		})
		_node.ExecPolicy = value
	}
	if value, ok := sjc.mutation.Concurrent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysjob.FieldConcurrent,
		})
		_node.Concurrent = value
	}
	if value, ok := sjc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: sysjob.FieldState,
		})
		_node.State = value
	}
	if value, ok := sjc.mutation.EntryId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysjob.FieldEntryId,
		})
		_node.EntryId = value
	}
	if value, ok := sjc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysjob.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sjc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysjob.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sjc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysjob.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := sjc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysjob.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := sjc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysjob.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := sjc.mutation.LogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysjob.LogsTable,
			Columns: []string{sysjob.LogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysjoblog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SysJobCreateBulk is the builder for creating many SysJob entities in bulk.
type SysJobCreateBulk struct {
	config
	builders []*SysJobCreate
}

// Save creates the SysJob entities in the database.
func (sjcb *SysJobCreateBulk) Save(ctx context.Context) ([]*SysJob, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sjcb.builders))
	nodes := make([]*SysJob, len(sjcb.builders))
	mutators := make([]Mutator, len(sjcb.builders))
	for i := range sjcb.builders {
		func(i int, root context.Context) {
			builder := sjcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysJobMutation)
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
					_, err = mutators[i+1].Mutate(root, sjcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sjcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sjcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sjcb *SysJobCreateBulk) SaveX(ctx context.Context) []*SysJob {
	v, err := sjcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sjcb *SysJobCreateBulk) Exec(ctx context.Context) error {
	_, err := sjcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sjcb *SysJobCreateBulk) ExecX(ctx context.Context) {
	if err := sjcb.Exec(ctx); err != nil {
		panic(err)
	}
}
