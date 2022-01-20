// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/ambalance"
	"hope/apps/novel/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AmBalanceDelete is the builder for deleting a AmBalance entity.
type AmBalanceDelete struct {
	config
	hooks    []Hook
	mutation *AmBalanceMutation
}

// Where appends a list predicates to the AmBalanceDelete builder.
func (abd *AmBalanceDelete) Where(ps ...predicate.AmBalance) *AmBalanceDelete {
	abd.mutation.Where(ps...)
	return abd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (abd *AmBalanceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(abd.hooks) == 0 {
		affected, err = abd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AmBalanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			abd.mutation = mutation
			affected, err = abd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(abd.hooks) - 1; i >= 0; i-- {
			if abd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = abd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, abd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (abd *AmBalanceDelete) ExecX(ctx context.Context) int {
	n, err := abd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (abd *AmBalanceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: ambalance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: ambalance.FieldID,
			},
		},
	}
	if ps := abd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, abd.driver, _spec)
}

// AmBalanceDeleteOne is the builder for deleting a single AmBalance entity.
type AmBalanceDeleteOne struct {
	abd *AmBalanceDelete
}

// Exec executes the deletion query.
func (abdo *AmBalanceDeleteOne) Exec(ctx context.Context) error {
	n, err := abdo.abd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ambalance.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (abdo *AmBalanceDeleteOne) ExecX(ctx context.Context) {
	abdo.abd.ExecX(ctx)
}
