// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelautobuy"
	"hope/apps/novel/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelAutoBuyDelete is the builder for deleting a NovelAutoBuy entity.
type NovelAutoBuyDelete struct {
	config
	hooks    []Hook
	mutation *NovelAutoBuyMutation
}

// Where appends a list predicates to the NovelAutoBuyDelete builder.
func (nabd *NovelAutoBuyDelete) Where(ps ...predicate.NovelAutoBuy) *NovelAutoBuyDelete {
	nabd.mutation.Where(ps...)
	return nabd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nabd *NovelAutoBuyDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nabd.hooks) == 0 {
		affected, err = nabd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelAutoBuyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nabd.mutation = mutation
			affected, err = nabd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nabd.hooks) - 1; i >= 0; i-- {
			if nabd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nabd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nabd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nabd *NovelAutoBuyDelete) ExecX(ctx context.Context) int {
	n, err := nabd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nabd *NovelAutoBuyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: novelautobuy.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelautobuy.FieldID,
			},
		},
	}
	if ps := nabd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, nabd.driver, _spec)
}

// NovelAutoBuyDeleteOne is the builder for deleting a single NovelAutoBuy entity.
type NovelAutoBuyDeleteOne struct {
	nabd *NovelAutoBuyDelete
}

// Exec executes the deletion query.
func (nabdo *NovelAutoBuyDeleteOne) Exec(ctx context.Context) error {
	n, err := nabdo.nabd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{novelautobuy.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nabdo *NovelAutoBuyDeleteOne) ExecX(ctx context.Context) {
	nabdo.nabd.ExecX(ctx)
}
