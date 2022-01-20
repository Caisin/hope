// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/qiniuconfig"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// QiniuConfigDelete is the builder for deleting a QiniuConfig entity.
type QiniuConfigDelete struct {
	config
	hooks    []Hook
	mutation *QiniuConfigMutation
}

// Where appends a list predicates to the QiniuConfigDelete builder.
func (qcd *QiniuConfigDelete) Where(ps ...predicate.QiniuConfig) *QiniuConfigDelete {
	qcd.mutation.Where(ps...)
	return qcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qcd *QiniuConfigDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(qcd.hooks) == 0 {
		affected, err = qcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QiniuConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qcd.mutation = mutation
			affected, err = qcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qcd.hooks) - 1; i >= 0; i-- {
			if qcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcd *QiniuConfigDelete) ExecX(ctx context.Context) int {
	n, err := qcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qcd *QiniuConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: qiniuconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: qiniuconfig.FieldID,
			},
		},
	}
	if ps := qcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, qcd.driver, _spec)
}

// QiniuConfigDeleteOne is the builder for deleting a single QiniuConfig entity.
type QiniuConfigDeleteOne struct {
	qcd *QiniuConfigDelete
}

// Exec executes the deletion query.
func (qcdo *QiniuConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := qcdo.qcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{qiniuconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qcdo *QiniuConfigDeleteOne) ExecX(ctx context.Context) {
	qcdo.qcd.ExecX(ctx)
}
