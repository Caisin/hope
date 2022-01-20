// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/payorder"
	"hope/apps/novel/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PayOrderDelete is the builder for deleting a PayOrder entity.
type PayOrderDelete struct {
	config
	hooks    []Hook
	mutation *PayOrderMutation
}

// Where appends a list predicates to the PayOrderDelete builder.
func (pod *PayOrderDelete) Where(ps ...predicate.PayOrder) *PayOrderDelete {
	pod.mutation.Where(ps...)
	return pod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pod *PayOrderDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pod.hooks) == 0 {
		affected, err = pod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PayOrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pod.mutation = mutation
			affected, err = pod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pod.hooks) - 1; i >= 0; i-- {
			if pod.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pod *PayOrderDelete) ExecX(ctx context.Context) int {
	n, err := pod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pod *PayOrderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: payorder.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: payorder.FieldID,
			},
		},
	}
	if ps := pod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pod.driver, _spec)
}

// PayOrderDeleteOne is the builder for deleting a single PayOrder entity.
type PayOrderDeleteOne struct {
	pod *PayOrderDelete
}

// Exec executes the deletion query.
func (podo *PayOrderDeleteOne) Exec(ctx context.Context) error {
	n, err := podo.pod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{payorder.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (podo *PayOrderDeleteOne) ExecX(ctx context.Context) {
	podo.pod.ExecX(ctx)
}
