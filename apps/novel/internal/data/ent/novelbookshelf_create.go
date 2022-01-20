// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelbookshelf"
	"hope/apps/novel/internal/data/ent/socialuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelBookshelfCreate is the builder for creating a NovelBookshelf entity.
type NovelBookshelfCreate struct {
	config
	mutation *NovelBookshelfMutation
	hooks    []Hook
}

// SetUserId sets the "userId" field.
func (nbc *NovelBookshelfCreate) SetUserId(i int64) *NovelBookshelfCreate {
	nbc.mutation.SetUserId(i)
	return nbc
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableUserId(i *int64) *NovelBookshelfCreate {
	if i != nil {
		nbc.SetUserId(*i)
	}
	return nbc
}

// SetUserName sets the "userName" field.
func (nbc *NovelBookshelfCreate) SetUserName(s string) *NovelBookshelfCreate {
	nbc.mutation.SetUserName(s)
	return nbc
}

// SetNillableUserName sets the "userName" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableUserName(s *string) *NovelBookshelfCreate {
	if s != nil {
		nbc.SetUserName(*s)
	}
	return nbc
}

// SetNovelId sets the "novelId" field.
func (nbc *NovelBookshelfCreate) SetNovelId(i int64) *NovelBookshelfCreate {
	nbc.mutation.SetNovelId(i)
	return nbc
}

// SetNillableNovelId sets the "novelId" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableNovelId(i *int64) *NovelBookshelfCreate {
	if i != nil {
		nbc.SetNovelId(*i)
	}
	return nbc
}

// SetLastReadTime sets the "lastReadTime" field.
func (nbc *NovelBookshelfCreate) SetLastReadTime(t time.Time) *NovelBookshelfCreate {
	nbc.mutation.SetLastReadTime(t)
	return nbc
}

// SetNillableLastReadTime sets the "lastReadTime" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableLastReadTime(t *time.Time) *NovelBookshelfCreate {
	if t != nil {
		nbc.SetLastReadTime(*t)
	}
	return nbc
}

// SetLastChapterOrder sets the "lastChapterOrder" field.
func (nbc *NovelBookshelfCreate) SetLastChapterOrder(i int32) *NovelBookshelfCreate {
	nbc.mutation.SetLastChapterOrder(i)
	return nbc
}

// SetNillableLastChapterOrder sets the "lastChapterOrder" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableLastChapterOrder(i *int32) *NovelBookshelfCreate {
	if i != nil {
		nbc.SetLastChapterOrder(*i)
	}
	return nbc
}

// SetLastChapterId sets the "lastChapterId" field.
func (nbc *NovelBookshelfCreate) SetLastChapterId(i int64) *NovelBookshelfCreate {
	nbc.mutation.SetLastChapterId(i)
	return nbc
}

// SetNillableLastChapterId sets the "lastChapterId" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableLastChapterId(i *int64) *NovelBookshelfCreate {
	if i != nil {
		nbc.SetLastChapterId(*i)
	}
	return nbc
}

// SetLastChapterName sets the "lastChapterName" field.
func (nbc *NovelBookshelfCreate) SetLastChapterName(s string) *NovelBookshelfCreate {
	nbc.mutation.SetLastChapterName(s)
	return nbc
}

// SetNillableLastChapterName sets the "lastChapterName" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableLastChapterName(s *string) *NovelBookshelfCreate {
	if s != nil {
		nbc.SetLastChapterName(*s)
	}
	return nbc
}

// SetRemark sets the "remark" field.
func (nbc *NovelBookshelfCreate) SetRemark(s string) *NovelBookshelfCreate {
	nbc.mutation.SetRemark(s)
	return nbc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableRemark(s *string) *NovelBookshelfCreate {
	if s != nil {
		nbc.SetRemark(*s)
	}
	return nbc
}

// SetCreatedAt sets the "createdAt" field.
func (nbc *NovelBookshelfCreate) SetCreatedAt(t time.Time) *NovelBookshelfCreate {
	nbc.mutation.SetCreatedAt(t)
	return nbc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableCreatedAt(t *time.Time) *NovelBookshelfCreate {
	if t != nil {
		nbc.SetCreatedAt(*t)
	}
	return nbc
}

// SetUpdatedAt sets the "updatedAt" field.
func (nbc *NovelBookshelfCreate) SetUpdatedAt(t time.Time) *NovelBookshelfCreate {
	nbc.mutation.SetUpdatedAt(t)
	return nbc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableUpdatedAt(t *time.Time) *NovelBookshelfCreate {
	if t != nil {
		nbc.SetUpdatedAt(*t)
	}
	return nbc
}

// SetCreateBy sets the "createBy" field.
func (nbc *NovelBookshelfCreate) SetCreateBy(i int64) *NovelBookshelfCreate {
	nbc.mutation.SetCreateBy(i)
	return nbc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableCreateBy(i *int64) *NovelBookshelfCreate {
	if i != nil {
		nbc.SetCreateBy(*i)
	}
	return nbc
}

// SetUpdateBy sets the "updateBy" field.
func (nbc *NovelBookshelfCreate) SetUpdateBy(i int64) *NovelBookshelfCreate {
	nbc.mutation.SetUpdateBy(i)
	return nbc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableUpdateBy(i *int64) *NovelBookshelfCreate {
	if i != nil {
		nbc.SetUpdateBy(*i)
	}
	return nbc
}

// SetTenantId sets the "tenantId" field.
func (nbc *NovelBookshelfCreate) SetTenantId(i int64) *NovelBookshelfCreate {
	nbc.mutation.SetTenantId(i)
	return nbc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (nbc *NovelBookshelfCreate) SetNillableTenantId(i *int64) *NovelBookshelfCreate {
	if i != nil {
		nbc.SetTenantId(*i)
	}
	return nbc
}

// SetUserID sets the "user" edge to the SocialUser entity by ID.
func (nbc *NovelBookshelfCreate) SetUserID(id int64) *NovelBookshelfCreate {
	nbc.mutation.SetUserID(id)
	return nbc
}

// SetUser sets the "user" edge to the SocialUser entity.
func (nbc *NovelBookshelfCreate) SetUser(s *SocialUser) *NovelBookshelfCreate {
	return nbc.SetUserID(s.ID)
}

// Mutation returns the NovelBookshelfMutation object of the builder.
func (nbc *NovelBookshelfCreate) Mutation() *NovelBookshelfMutation {
	return nbc.mutation
}

// Save creates the NovelBookshelf in the database.
func (nbc *NovelBookshelfCreate) Save(ctx context.Context) (*NovelBookshelf, error) {
	var (
		err  error
		node *NovelBookshelf
	)
	nbc.defaults()
	if len(nbc.hooks) == 0 {
		if err = nbc.check(); err != nil {
			return nil, err
		}
		node, err = nbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelBookshelfMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nbc.check(); err != nil {
				return nil, err
			}
			nbc.mutation = mutation
			if node, err = nbc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nbc.hooks) - 1; i >= 0; i-- {
			if nbc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nbc *NovelBookshelfCreate) SaveX(ctx context.Context) *NovelBookshelf {
	v, err := nbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nbc *NovelBookshelfCreate) Exec(ctx context.Context) error {
	_, err := nbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nbc *NovelBookshelfCreate) ExecX(ctx context.Context) {
	if err := nbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nbc *NovelBookshelfCreate) defaults() {
	if _, ok := nbc.mutation.CreatedAt(); !ok {
		v := novelbookshelf.DefaultCreatedAt()
		nbc.mutation.SetCreatedAt(v)
	}
	if _, ok := nbc.mutation.UpdatedAt(); !ok {
		v := novelbookshelf.DefaultUpdatedAt()
		nbc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nbc.mutation.CreateBy(); !ok {
		v := novelbookshelf.DefaultCreateBy
		nbc.mutation.SetCreateBy(v)
	}
	if _, ok := nbc.mutation.UpdateBy(); !ok {
		v := novelbookshelf.DefaultUpdateBy
		nbc.mutation.SetUpdateBy(v)
	}
	if _, ok := nbc.mutation.TenantId(); !ok {
		v := novelbookshelf.DefaultTenantId
		nbc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nbc *NovelBookshelfCreate) check() error {
	if _, ok := nbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := nbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := nbc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := nbc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := nbc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	if _, ok := nbc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (nbc *NovelBookshelfCreate) sqlSave(ctx context.Context) (*NovelBookshelf, error) {
	_node, _spec := nbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (nbc *NovelBookshelfCreate) createSpec() (*NovelBookshelf, *sqlgraph.CreateSpec) {
	var (
		_node = &NovelBookshelf{config: nbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: novelbookshelf.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelbookshelf.FieldID,
			},
		}
	)
	if value, ok := nbc.mutation.UserId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldUserId,
		})
		_node.UserId = value
	}
	if value, ok := nbc.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbookshelf.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := nbc.mutation.NovelId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldNovelId,
		})
		_node.NovelId = value
	}
	if value, ok := nbc.mutation.LastReadTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbookshelf.FieldLastReadTime,
		})
		_node.LastReadTime = value
	}
	if value, ok := nbc.mutation.LastChapterOrder(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterOrder,
		})
		_node.LastChapterOrder = value
	}
	if value, ok := nbc.mutation.LastChapterId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterId,
		})
		_node.LastChapterId = value
	}
	if value, ok := nbc.mutation.LastChapterName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterName,
		})
		_node.LastChapterName = value
	}
	if value, ok := nbc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbookshelf.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := nbc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbookshelf.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := nbc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbookshelf.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := nbc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := nbc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := nbc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := nbc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novelbookshelf.UserTable,
			Columns: []string{novelbookshelf.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: socialuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.social_user_bookshelves = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NovelBookshelfCreateBulk is the builder for creating many NovelBookshelf entities in bulk.
type NovelBookshelfCreateBulk struct {
	config
	builders []*NovelBookshelfCreate
}

// Save creates the NovelBookshelf entities in the database.
func (nbcb *NovelBookshelfCreateBulk) Save(ctx context.Context) ([]*NovelBookshelf, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nbcb.builders))
	nodes := make([]*NovelBookshelf, len(nbcb.builders))
	mutators := make([]Mutator, len(nbcb.builders))
	for i := range nbcb.builders {
		func(i int, root context.Context) {
			builder := nbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NovelBookshelfMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, nbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nbcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, nbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nbcb *NovelBookshelfCreateBulk) SaveX(ctx context.Context) []*NovelBookshelf {
	v, err := nbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nbcb *NovelBookshelfCreateBulk) Exec(ctx context.Context) error {
	_, err := nbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nbcb *NovelBookshelfCreateBulk) ExecX(ctx context.Context) {
	if err := nbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
