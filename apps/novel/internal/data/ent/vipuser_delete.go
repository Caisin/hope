// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/vipuser"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VipUserDelete is the builder for deleting a VipUser entity.
type VipUserDelete struct {
	config
	hooks    []Hook
	mutation *VipUserMutation
}

// Where appends a list predicates to the VipUserDelete builder.
func (vud *VipUserDelete) Where(ps ...predicate.VipUser) *VipUserDelete {
	vud.mutation.Where(ps...)
	return vud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vud *VipUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vud.hooks) == 0 {
		affected, err = vud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VipUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vud.mutation = mutation
			affected, err = vud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vud.hooks) - 1; i >= 0; i-- {
			if vud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vud *VipUserDelete) ExecX(ctx context.Context) int {
	n, err := vud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vud *VipUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vipuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: vipuser.FieldID,
			},
		},
	}
	if ps := vud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vud.driver, _spec)
}

// VipUserDeleteOne is the builder for deleting a single VipUser entity.
type VipUserDeleteOne struct {
	vud *VipUserDelete
}

// Exec executes the deletion query.
func (vudo *VipUserDeleteOne) Exec(ctx context.Context) error {
	n, err := vudo.vud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vipuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vudo *VipUserDeleteOne) ExecX(ctx context.Context) {
	vudo.vud.ExecX(ctx)
}
