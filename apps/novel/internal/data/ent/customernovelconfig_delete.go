// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/customernovelconfig"
	"hope/apps/novel/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerNovelConfigDelete is the builder for deleting a CustomerNovelConfig entity.
type CustomerNovelConfigDelete struct {
	config
	hooks    []Hook
	mutation *CustomerNovelConfigMutation
}

// Where appends a list predicates to the CustomerNovelConfigDelete builder.
func (cncd *CustomerNovelConfigDelete) Where(ps ...predicate.CustomerNovelConfig) *CustomerNovelConfigDelete {
	cncd.mutation.Where(ps...)
	return cncd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cncd *CustomerNovelConfigDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cncd.hooks) == 0 {
		affected, err = cncd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerNovelConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cncd.mutation = mutation
			affected, err = cncd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cncd.hooks) - 1; i >= 0; i-- {
			if cncd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cncd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cncd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cncd *CustomerNovelConfigDelete) ExecX(ctx context.Context) int {
	n, err := cncd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cncd *CustomerNovelConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: customernovelconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: customernovelconfig.FieldID,
			},
		},
	}
	if ps := cncd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cncd.driver, _spec)
}

// CustomerNovelConfigDeleteOne is the builder for deleting a single CustomerNovelConfig entity.
type CustomerNovelConfigDeleteOne struct {
	cncd *CustomerNovelConfigDelete
}

// Exec executes the deletion query.
func (cncdo *CustomerNovelConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := cncdo.cncd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{customernovelconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cncdo *CustomerNovelConfigDeleteOne) ExecX(ctx context.Context) {
	cncdo.cncd.ExecX(ctx)
}
