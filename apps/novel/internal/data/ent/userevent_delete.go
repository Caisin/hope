// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/userevent"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserEventDelete is the builder for deleting a UserEvent entity.
type UserEventDelete struct {
	config
	hooks    []Hook
	mutation *UserEventMutation
}

// Where appends a list predicates to the UserEventDelete builder.
func (ued *UserEventDelete) Where(ps ...predicate.UserEvent) *UserEventDelete {
	ued.mutation.Where(ps...)
	return ued
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ued *UserEventDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ued.hooks) == 0 {
		affected, err = ued.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ued.mutation = mutation
			affected, err = ued.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ued.hooks) - 1; i >= 0; i-- {
			if ued.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ued.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ued.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ued *UserEventDelete) ExecX(ctx context.Context) int {
	n, err := ued.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ued *UserEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userevent.FieldID,
			},
		},
	}
	if ps := ued.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ued.driver, _spec)
}

// UserEventDeleteOne is the builder for deleting a single UserEvent entity.
type UserEventDeleteOne struct {
	ued *UserEventDelete
}

// Exec executes the deletion query.
func (uedo *UserEventDeleteOne) Exec(ctx context.Context) error {
	n, err := uedo.ued.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uedo *UserEventDeleteOne) ExecX(ctx context.Context) {
	uedo.ued.ExecX(ctx)
}