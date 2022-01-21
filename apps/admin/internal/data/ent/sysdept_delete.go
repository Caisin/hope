// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysdept"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysDeptDelete is the builder for deleting a SysDept entity.
type SysDeptDelete struct {
	config
	hooks    []Hook
	mutation *SysDeptMutation
}

// Where appends a list predicates to the SysDeptDelete builder.
func (sdd *SysDeptDelete) Where(ps ...predicate.SysDept) *SysDeptDelete {
	sdd.mutation.Where(ps...)
	return sdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdd *SysDeptDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sdd.hooks) == 0 {
		affected, err = sdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDeptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdd.mutation = mutation
			affected, err = sdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdd.hooks) - 1; i >= 0; i-- {
			if sdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdd *SysDeptDelete) ExecX(ctx context.Context) int {
	n, err := sdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdd *SysDeptDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sysdept.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysdept.FieldID,
			},
		},
	}
	if ps := sdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sdd.driver, _spec)
}

// SysDeptDeleteOne is the builder for deleting a single SysDept entity.
type SysDeptDeleteOne struct {
	sdd *SysDeptDelete
}

// Exec executes the deletion query.
func (sddo *SysDeptDeleteOne) Exec(ctx context.Context) error {
	n, err := sddo.sdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysdept.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sddo *SysDeptDeleteOne) ExecX(ctx context.Context) {
	sddo.sdd.ExecX(ctx)
}