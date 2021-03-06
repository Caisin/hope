// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysoperalog"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysOperaLogDelete is the builder for deleting a SysOperaLog entity.
type SysOperaLogDelete struct {
	config
	hooks    []Hook
	mutation *SysOperaLogMutation
}

// Where appends a list predicates to the SysOperaLogDelete builder.
func (sold *SysOperaLogDelete) Where(ps ...predicate.SysOperaLog) *SysOperaLogDelete {
	sold.mutation.Where(ps...)
	return sold
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sold *SysOperaLogDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sold.hooks) == 0 {
		affected, err = sold.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysOperaLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sold.mutation = mutation
			affected, err = sold.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sold.hooks) - 1; i >= 0; i-- {
			if sold.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sold.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sold.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sold *SysOperaLogDelete) ExecX(ctx context.Context) int {
	n, err := sold.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sold *SysOperaLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sysoperalog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysoperalog.FieldID,
			},
		},
	}
	if ps := sold.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sold.driver, _spec)
}

// SysOperaLogDeleteOne is the builder for deleting a single SysOperaLog entity.
type SysOperaLogDeleteOne struct {
	sold *SysOperaLogDelete
}

// Exec executes the deletion query.
func (soldo *SysOperaLogDeleteOne) Exec(ctx context.Context) error {
	n, err := soldo.sold.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysoperalog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (soldo *SysOperaLogDeleteOne) ExecX(ctx context.Context) {
	soldo.sold.ExecX(ctx)
}
