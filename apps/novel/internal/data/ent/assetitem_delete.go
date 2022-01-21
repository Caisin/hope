// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/assetitem"
	"hope/apps/novel/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AssetItemDelete is the builder for deleting a AssetItem entity.
type AssetItemDelete struct {
	config
	hooks    []Hook
	mutation *AssetItemMutation
}

// Where appends a list predicates to the AssetItemDelete builder.
func (aid *AssetItemDelete) Where(ps ...predicate.AssetItem) *AssetItemDelete {
	aid.mutation.Where(ps...)
	return aid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aid *AssetItemDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aid.hooks) == 0 {
		affected, err = aid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aid.mutation = mutation
			affected, err = aid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aid.hooks) - 1; i >= 0; i-- {
			if aid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aid *AssetItemDelete) ExecX(ctx context.Context) int {
	n, err := aid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aid *AssetItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: assetitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: assetitem.FieldID,
			},
		},
	}
	if ps := aid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aid.driver, _spec)
}

// AssetItemDeleteOne is the builder for deleting a single AssetItem entity.
type AssetItemDeleteOne struct {
	aid *AssetItemDelete
}

// Exec executes the deletion query.
func (aido *AssetItemDeleteOne) Exec(ctx context.Context) error {
	n, err := aido.aid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{assetitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aido *AssetItemDeleteOne) ExecX(ctx context.Context) {
	aido.aid.ExecX(ctx)
}