// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/listenrecord"
	"hope/apps/novel/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ListenRecordDelete is the builder for deleting a ListenRecord entity.
type ListenRecordDelete struct {
	config
	hooks    []Hook
	mutation *ListenRecordMutation
}

// Where appends a list predicates to the ListenRecordDelete builder.
func (lrd *ListenRecordDelete) Where(ps ...predicate.ListenRecord) *ListenRecordDelete {
	lrd.mutation.Where(ps...)
	return lrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lrd *ListenRecordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lrd.hooks) == 0 {
		affected, err = lrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ListenRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lrd.mutation = mutation
			affected, err = lrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lrd.hooks) - 1; i >= 0; i-- {
			if lrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (lrd *ListenRecordDelete) ExecX(ctx context.Context) int {
	n, err := lrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lrd *ListenRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: listenrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: listenrecord.FieldID,
			},
		},
	}
	if ps := lrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, lrd.driver, _spec)
}

// ListenRecordDeleteOne is the builder for deleting a single ListenRecord entity.
type ListenRecordDeleteOne struct {
	lrd *ListenRecordDelete
}

// Exec executes the deletion query.
func (lrdo *ListenRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := lrdo.lrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{listenrecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lrdo *ListenRecordDeleteOne) ExecX(ctx context.Context) {
	lrdo.lrd.ExecX(ctx)
}
