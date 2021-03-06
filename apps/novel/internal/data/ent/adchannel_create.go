// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/adchannel"
	"hope/apps/novel/internal/data/ent/payorder"
	"hope/apps/novel/internal/data/ent/socialuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdChannelCreate is the builder for creating a AdChannel entity.
type AdChannelCreate struct {
	config
	mutation *AdChannelMutation
	hooks    []Hook
}

// SetChannelName sets the "channelName" field.
func (acc *AdChannelCreate) SetChannelName(s string) *AdChannelCreate {
	acc.mutation.SetChannelName(s)
	return acc
}

// SetNillableChannelName sets the "channelName" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableChannelName(s *string) *AdChannelCreate {
	if s != nil {
		acc.SetChannelName(*s)
	}
	return acc
}

// SetNovelId sets the "novelId" field.
func (acc *AdChannelCreate) SetNovelId(i int64) *AdChannelCreate {
	acc.mutation.SetNovelId(i)
	return acc
}

// SetNillableNovelId sets the "novelId" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableNovelId(i *int64) *AdChannelCreate {
	if i != nil {
		acc.SetNovelId(*i)
	}
	return acc
}

// SetReg sets the "reg" field.
func (acc *AdChannelCreate) SetReg(i int64) *AdChannelCreate {
	acc.mutation.SetReg(i)
	return acc
}

// SetNillableReg sets the "reg" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableReg(i *int64) *AdChannelCreate {
	if i != nil {
		acc.SetReg(*i)
	}
	return acc
}

// SetPay sets the "pay" field.
func (acc *AdChannelCreate) SetPay(i int64) *AdChannelCreate {
	acc.mutation.SetPay(i)
	return acc
}

// SetNillablePay sets the "pay" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillablePay(i *int64) *AdChannelCreate {
	if i != nil {
		acc.SetPay(*i)
	}
	return acc
}

// SetNovelName sets the "novelName" field.
func (acc *AdChannelCreate) SetNovelName(s string) *AdChannelCreate {
	acc.mutation.SetNovelName(s)
	return acc
}

// SetNillableNovelName sets the "novelName" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableNovelName(s *string) *AdChannelCreate {
	if s != nil {
		acc.SetNovelName(*s)
	}
	return acc
}

// SetChapterId sets the "chapterId" field.
func (acc *AdChannelCreate) SetChapterId(i int64) *AdChannelCreate {
	acc.mutation.SetChapterId(i)
	return acc
}

// SetNillableChapterId sets the "chapterId" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableChapterId(i *int64) *AdChannelCreate {
	if i != nil {
		acc.SetChapterId(*i)
	}
	return acc
}

// SetChapterNum sets the "chapterNum" field.
func (acc *AdChannelCreate) SetChapterNum(i int32) *AdChannelCreate {
	acc.mutation.SetChapterNum(i)
	return acc
}

// SetNillableChapterNum sets the "chapterNum" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableChapterNum(i *int32) *AdChannelCreate {
	if i != nil {
		acc.SetChapterNum(*i)
	}
	return acc
}

// SetAdType sets the "adType" field.
func (acc *AdChannelCreate) SetAdType(s string) *AdChannelCreate {
	acc.mutation.SetAdType(s)
	return acc
}

// SetNillableAdType sets the "adType" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableAdType(s *string) *AdChannelCreate {
	if s != nil {
		acc.SetAdType(*s)
	}
	return acc
}

// SetImg sets the "img" field.
func (acc *AdChannelCreate) SetImg(s string) *AdChannelCreate {
	acc.mutation.SetImg(s)
	return acc
}

// SetNillableImg sets the "img" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableImg(s *string) *AdChannelCreate {
	if s != nil {
		acc.SetImg(*s)
	}
	return acc
}

// SetCreatedAt sets the "createdAt" field.
func (acc *AdChannelCreate) SetCreatedAt(t time.Time) *AdChannelCreate {
	acc.mutation.SetCreatedAt(t)
	return acc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableCreatedAt(t *time.Time) *AdChannelCreate {
	if t != nil {
		acc.SetCreatedAt(*t)
	}
	return acc
}

// SetUpdatedAt sets the "updatedAt" field.
func (acc *AdChannelCreate) SetUpdatedAt(t time.Time) *AdChannelCreate {
	acc.mutation.SetUpdatedAt(t)
	return acc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableUpdatedAt(t *time.Time) *AdChannelCreate {
	if t != nil {
		acc.SetUpdatedAt(*t)
	}
	return acc
}

// SetCreateBy sets the "createBy" field.
func (acc *AdChannelCreate) SetCreateBy(i int64) *AdChannelCreate {
	acc.mutation.SetCreateBy(i)
	return acc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableCreateBy(i *int64) *AdChannelCreate {
	if i != nil {
		acc.SetCreateBy(*i)
	}
	return acc
}

// SetUpdateBy sets the "updateBy" field.
func (acc *AdChannelCreate) SetUpdateBy(i int64) *AdChannelCreate {
	acc.mutation.SetUpdateBy(i)
	return acc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableUpdateBy(i *int64) *AdChannelCreate {
	if i != nil {
		acc.SetUpdateBy(*i)
	}
	return acc
}

// SetTenantId sets the "tenantId" field.
func (acc *AdChannelCreate) SetTenantId(i int64) *AdChannelCreate {
	acc.mutation.SetTenantId(i)
	return acc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (acc *AdChannelCreate) SetNillableTenantId(i *int64) *AdChannelCreate {
	if i != nil {
		acc.SetTenantId(*i)
	}
	return acc
}

// AddUserIDs adds the "users" edge to the SocialUser entity by IDs.
func (acc *AdChannelCreate) AddUserIDs(ids ...int64) *AdChannelCreate {
	acc.mutation.AddUserIDs(ids...)
	return acc
}

// AddUsers adds the "users" edges to the SocialUser entity.
func (acc *AdChannelCreate) AddUsers(s ...*SocialUser) *AdChannelCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return acc.AddUserIDs(ids...)
}

// AddOrderIDs adds the "orders" edge to the PayOrder entity by IDs.
func (acc *AdChannelCreate) AddOrderIDs(ids ...int64) *AdChannelCreate {
	acc.mutation.AddOrderIDs(ids...)
	return acc
}

// AddOrders adds the "orders" edges to the PayOrder entity.
func (acc *AdChannelCreate) AddOrders(p ...*PayOrder) *AdChannelCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return acc.AddOrderIDs(ids...)
}

// Mutation returns the AdChannelMutation object of the builder.
func (acc *AdChannelCreate) Mutation() *AdChannelMutation {
	return acc.mutation
}

// Save creates the AdChannel in the database.
func (acc *AdChannelCreate) Save(ctx context.Context) (*AdChannel, error) {
	var (
		err  error
		node *AdChannel
	)
	acc.defaults()
	if len(acc.hooks) == 0 {
		if err = acc.check(); err != nil {
			return nil, err
		}
		node, err = acc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acc.check(); err != nil {
				return nil, err
			}
			acc.mutation = mutation
			if node, err = acc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(acc.hooks) - 1; i >= 0; i-- {
			if acc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (acc *AdChannelCreate) SaveX(ctx context.Context) *AdChannel {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *AdChannelCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *AdChannelCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acc *AdChannelCreate) defaults() {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		v := adchannel.DefaultCreatedAt()
		acc.mutation.SetCreatedAt(v)
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		v := adchannel.DefaultUpdatedAt()
		acc.mutation.SetUpdatedAt(v)
	}
	if _, ok := acc.mutation.CreateBy(); !ok {
		v := adchannel.DefaultCreateBy
		acc.mutation.SetCreateBy(v)
	}
	if _, ok := acc.mutation.UpdateBy(); !ok {
		v := adchannel.DefaultUpdateBy
		acc.mutation.SetUpdateBy(v)
	}
	if _, ok := acc.mutation.TenantId(); !ok {
		v := adchannel.DefaultTenantId
		acc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *AdChannelCreate) check() error {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "AdChannel.createdAt"`)}
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "AdChannel.updatedAt"`)}
	}
	if _, ok := acc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "AdChannel.createBy"`)}
	}
	if _, ok := acc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "AdChannel.updateBy"`)}
	}
	if _, ok := acc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "AdChannel.tenantId"`)}
	}
	return nil
}

func (acc *AdChannelCreate) sqlSave(ctx context.Context) (*AdChannel, error) {
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (acc *AdChannelCreate) createSpec() (*AdChannel, *sqlgraph.CreateSpec) {
	var (
		_node = &AdChannel{config: acc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: adchannel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: adchannel.FieldID,
			},
		}
	)
	if value, ok := acc.mutation.ChannelName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adchannel.FieldChannelName,
		})
		_node.ChannelName = value
	}
	if value, ok := acc.mutation.NovelId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchannel.FieldNovelId,
		})
		_node.NovelId = value
	}
	if value, ok := acc.mutation.Reg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchannel.FieldReg,
		})
		_node.Reg = value
	}
	if value, ok := acc.mutation.Pay(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchannel.FieldPay,
		})
		_node.Pay = value
	}
	if value, ok := acc.mutation.NovelName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adchannel.FieldNovelName,
		})
		_node.NovelName = value
	}
	if value, ok := acc.mutation.ChapterId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchannel.FieldChapterId,
		})
		_node.ChapterId = value
	}
	if value, ok := acc.mutation.ChapterNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: adchannel.FieldChapterNum,
		})
		_node.ChapterNum = value
	}
	if value, ok := acc.mutation.AdType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adchannel.FieldAdType,
		})
		_node.AdType = value
	}
	if value, ok := acc.mutation.Img(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adchannel.FieldImg,
		})
		_node.Img = value
	}
	if value, ok := acc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: adchannel.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := acc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: adchannel.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := acc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchannel.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := acc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchannel.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := acc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: adchannel.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := acc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adchannel.UsersTable,
			Columns: []string{adchannel.UsersColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := acc.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   adchannel.OrdersTable,
			Columns: []string{adchannel.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: payorder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AdChannelCreateBulk is the builder for creating many AdChannel entities in bulk.
type AdChannelCreateBulk struct {
	config
	builders []*AdChannelCreate
}

// Save creates the AdChannel entities in the database.
func (accb *AdChannelCreateBulk) Save(ctx context.Context) ([]*AdChannel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*AdChannel, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdChannelMutation)
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
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *AdChannelCreateBulk) SaveX(ctx context.Context) []*AdChannel {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *AdChannelCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *AdChannelCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}
