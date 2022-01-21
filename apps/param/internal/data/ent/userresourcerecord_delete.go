// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/userresourcerecord"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserResourceRecordDelete is the builder for deleting a UserResourceRecord entity.
type UserResourceRecordDelete struct {
	config
	hooks    []Hook
	mutation *UserResourceRecordMutation
}

// Where appends a list predicates to the UserResourceRecordDelete builder.
func (urrd *UserResourceRecordDelete) Where(ps ...predicate.UserResourceRecord) *UserResourceRecordDelete {
	urrd.mutation.Where(ps...)
	return urrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (urrd *UserResourceRecordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(urrd.hooks) == 0 {
		affected, err = urrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserResourceRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			urrd.mutation = mutation
			affected, err = urrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(urrd.hooks) - 1; i >= 0; i-- {
			if urrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = urrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, urrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (urrd *UserResourceRecordDelete) ExecX(ctx context.Context) int {
	n, err := urrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (urrd *UserResourceRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userresourcerecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userresourcerecord.FieldID,
			},
		},
	}
	if ps := urrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, urrd.driver, _spec)
}

// UserResourceRecordDeleteOne is the builder for deleting a single UserResourceRecord entity.
type UserResourceRecordDeleteOne struct {
	urrd *UserResourceRecordDelete
}

// Exec executes the deletion query.
func (urrdo *UserResourceRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := urrdo.urrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userresourcerecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (urrdo *UserResourceRecordDeleteOne) ExecX(ctx context.Context) {
	urrdo.urrd.ExecX(ctx)
}