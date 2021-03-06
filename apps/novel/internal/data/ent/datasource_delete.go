// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/datasource"
	"hope/apps/novel/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DataSourceDelete is the builder for deleting a DataSource entity.
type DataSourceDelete struct {
	config
	hooks    []Hook
	mutation *DataSourceMutation
}

// Where appends a list predicates to the DataSourceDelete builder.
func (dsd *DataSourceDelete) Where(ps ...predicate.DataSource) *DataSourceDelete {
	dsd.mutation.Where(ps...)
	return dsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dsd *DataSourceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dsd.hooks) == 0 {
		affected, err = dsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DataSourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dsd.mutation = mutation
			affected, err = dsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dsd.hooks) - 1; i >= 0; i-- {
			if dsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dsd *DataSourceDelete) ExecX(ctx context.Context) int {
	n, err := dsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dsd *DataSourceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: datasource.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: datasource.FieldID,
			},
		},
	}
	if ps := dsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dsd.driver, _spec)
}

// DataSourceDeleteOne is the builder for deleting a single DataSource entity.
type DataSourceDeleteOne struct {
	dsd *DataSourceDelete
}

// Exec executes the deletion query.
func (dsdo *DataSourceDeleteOne) Exec(ctx context.Context) error {
	n, err := dsdo.dsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{datasource.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dsdo *DataSourceDeleteOne) ExecX(ctx context.Context) {
	dsdo.dsd.ExecX(ctx)
}
