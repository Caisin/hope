// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/customernovels"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerNovelsCreate is the builder for creating a CustomerNovels entity.
type CustomerNovelsCreate struct {
	config
	mutation *CustomerNovelsMutation
	hooks    []Hook
}

// SetNovelId sets the "novelId" field.
func (cnc *CustomerNovelsCreate) SetNovelId(i int64) *CustomerNovelsCreate {
	cnc.mutation.SetNovelId(i)
	return cnc
}

// SetNillableNovelId sets the "novelId" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableNovelId(i *int64) *CustomerNovelsCreate {
	if i != nil {
		cnc.SetNovelId(*i)
	}
	return cnc
}

// SetTypeId sets the "typeId" field.
func (cnc *CustomerNovelsCreate) SetTypeId(i int32) *CustomerNovelsCreate {
	cnc.mutation.SetTypeId(i)
	return cnc
}

// SetNillableTypeId sets the "typeId" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableTypeId(i *int32) *CustomerNovelsCreate {
	if i != nil {
		cnc.SetTypeId(*i)
	}
	return cnc
}

// SetTypeCode sets the "typeCode" field.
func (cnc *CustomerNovelsCreate) SetTypeCode(s string) *CustomerNovelsCreate {
	cnc.mutation.SetTypeCode(s)
	return cnc
}

// SetNillableTypeCode sets the "typeCode" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableTypeCode(s *string) *CustomerNovelsCreate {
	if s != nil {
		cnc.SetTypeCode(*s)
	}
	return cnc
}

// SetGroupCode sets the "groupCode" field.
func (cnc *CustomerNovelsCreate) SetGroupCode(s string) *CustomerNovelsCreate {
	cnc.mutation.SetGroupCode(s)
	return cnc
}

// SetNillableGroupCode sets the "groupCode" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableGroupCode(s *string) *CustomerNovelsCreate {
	if s != nil {
		cnc.SetGroupCode(*s)
	}
	return cnc
}

// SetFieldName sets the "fieldName" field.
func (cnc *CustomerNovelsCreate) SetFieldName(s string) *CustomerNovelsCreate {
	cnc.mutation.SetFieldName(s)
	return cnc
}

// SetNillableFieldName sets the "fieldName" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableFieldName(s *string) *CustomerNovelsCreate {
	if s != nil {
		cnc.SetFieldName(*s)
	}
	return cnc
}

// SetCover sets the "cover" field.
func (cnc *CustomerNovelsCreate) SetCover(s string) *CustomerNovelsCreate {
	cnc.mutation.SetCover(s)
	return cnc
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableCover(s *string) *CustomerNovelsCreate {
	if s != nil {
		cnc.SetCover(*s)
	}
	return cnc
}

// SetOrderNum sets the "orderNum" field.
func (cnc *CustomerNovelsCreate) SetOrderNum(i int32) *CustomerNovelsCreate {
	cnc.mutation.SetOrderNum(i)
	return cnc
}

// SetNillableOrderNum sets the "orderNum" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableOrderNum(i *int32) *CustomerNovelsCreate {
	if i != nil {
		cnc.SetOrderNum(*i)
	}
	return cnc
}

// SetRemark sets the "remark" field.
func (cnc *CustomerNovelsCreate) SetRemark(s string) *CustomerNovelsCreate {
	cnc.mutation.SetRemark(s)
	return cnc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableRemark(s *string) *CustomerNovelsCreate {
	if s != nil {
		cnc.SetRemark(*s)
	}
	return cnc
}

// SetEffectTime sets the "effectTime" field.
func (cnc *CustomerNovelsCreate) SetEffectTime(t time.Time) *CustomerNovelsCreate {
	cnc.mutation.SetEffectTime(t)
	return cnc
}

// SetNillableEffectTime sets the "effectTime" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableEffectTime(t *time.Time) *CustomerNovelsCreate {
	if t != nil {
		cnc.SetEffectTime(*t)
	}
	return cnc
}

// SetExpiredTime sets the "expiredTime" field.
func (cnc *CustomerNovelsCreate) SetExpiredTime(t time.Time) *CustomerNovelsCreate {
	cnc.mutation.SetExpiredTime(t)
	return cnc
}

// SetNillableExpiredTime sets the "expiredTime" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableExpiredTime(t *time.Time) *CustomerNovelsCreate {
	if t != nil {
		cnc.SetExpiredTime(*t)
	}
	return cnc
}

// SetCreatedAt sets the "createdAt" field.
func (cnc *CustomerNovelsCreate) SetCreatedAt(t time.Time) *CustomerNovelsCreate {
	cnc.mutation.SetCreatedAt(t)
	return cnc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableCreatedAt(t *time.Time) *CustomerNovelsCreate {
	if t != nil {
		cnc.SetCreatedAt(*t)
	}
	return cnc
}

// SetUpdatedAt sets the "updatedAt" field.
func (cnc *CustomerNovelsCreate) SetUpdatedAt(t time.Time) *CustomerNovelsCreate {
	cnc.mutation.SetUpdatedAt(t)
	return cnc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableUpdatedAt(t *time.Time) *CustomerNovelsCreate {
	if t != nil {
		cnc.SetUpdatedAt(*t)
	}
	return cnc
}

// SetCreateBy sets the "createBy" field.
func (cnc *CustomerNovelsCreate) SetCreateBy(i int64) *CustomerNovelsCreate {
	cnc.mutation.SetCreateBy(i)
	return cnc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableCreateBy(i *int64) *CustomerNovelsCreate {
	if i != nil {
		cnc.SetCreateBy(*i)
	}
	return cnc
}

// SetUpdateBy sets the "updateBy" field.
func (cnc *CustomerNovelsCreate) SetUpdateBy(i int64) *CustomerNovelsCreate {
	cnc.mutation.SetUpdateBy(i)
	return cnc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableUpdateBy(i *int64) *CustomerNovelsCreate {
	if i != nil {
		cnc.SetUpdateBy(*i)
	}
	return cnc
}

// SetTenantId sets the "tenantId" field.
func (cnc *CustomerNovelsCreate) SetTenantId(i int64) *CustomerNovelsCreate {
	cnc.mutation.SetTenantId(i)
	return cnc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (cnc *CustomerNovelsCreate) SetNillableTenantId(i *int64) *CustomerNovelsCreate {
	if i != nil {
		cnc.SetTenantId(*i)
	}
	return cnc
}

// Mutation returns the CustomerNovelsMutation object of the builder.
func (cnc *CustomerNovelsCreate) Mutation() *CustomerNovelsMutation {
	return cnc.mutation
}

// Save creates the CustomerNovels in the database.
func (cnc *CustomerNovelsCreate) Save(ctx context.Context) (*CustomerNovels, error) {
	var (
		err  error
		node *CustomerNovels
	)
	cnc.defaults()
	if len(cnc.hooks) == 0 {
		if err = cnc.check(); err != nil {
			return nil, err
		}
		node, err = cnc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerNovelsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cnc.check(); err != nil {
				return nil, err
			}
			cnc.mutation = mutation
			if node, err = cnc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cnc.hooks) - 1; i >= 0; i-- {
			if cnc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cnc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cnc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cnc *CustomerNovelsCreate) SaveX(ctx context.Context) *CustomerNovels {
	v, err := cnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cnc *CustomerNovelsCreate) Exec(ctx context.Context) error {
	_, err := cnc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnc *CustomerNovelsCreate) ExecX(ctx context.Context) {
	if err := cnc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cnc *CustomerNovelsCreate) defaults() {
	if _, ok := cnc.mutation.EffectTime(); !ok {
		v := customernovels.DefaultEffectTime()
		cnc.mutation.SetEffectTime(v)
	}
	if _, ok := cnc.mutation.ExpiredTime(); !ok {
		v := customernovels.DefaultExpiredTime()
		cnc.mutation.SetExpiredTime(v)
	}
	if _, ok := cnc.mutation.CreatedAt(); !ok {
		v := customernovels.DefaultCreatedAt()
		cnc.mutation.SetCreatedAt(v)
	}
	if _, ok := cnc.mutation.UpdatedAt(); !ok {
		v := customernovels.DefaultUpdatedAt()
		cnc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cnc.mutation.CreateBy(); !ok {
		v := customernovels.DefaultCreateBy
		cnc.mutation.SetCreateBy(v)
	}
	if _, ok := cnc.mutation.UpdateBy(); !ok {
		v := customernovels.DefaultUpdateBy
		cnc.mutation.SetUpdateBy(v)
	}
	if _, ok := cnc.mutation.TenantId(); !ok {
		v := customernovels.DefaultTenantId
		cnc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cnc *CustomerNovelsCreate) check() error {
	if _, ok := cnc.mutation.EffectTime(); !ok {
		return &ValidationError{Name: "effectTime", err: errors.New(`ent: missing required field "CustomerNovels.effectTime"`)}
	}
	if _, ok := cnc.mutation.ExpiredTime(); !ok {
		return &ValidationError{Name: "expiredTime", err: errors.New(`ent: missing required field "CustomerNovels.expiredTime"`)}
	}
	if _, ok := cnc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "CustomerNovels.createdAt"`)}
	}
	if _, ok := cnc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "CustomerNovels.updatedAt"`)}
	}
	if _, ok := cnc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "CustomerNovels.createBy"`)}
	}
	if _, ok := cnc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "CustomerNovels.updateBy"`)}
	}
	if _, ok := cnc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "CustomerNovels.tenantId"`)}
	}
	return nil
}

func (cnc *CustomerNovelsCreate) sqlSave(ctx context.Context) (*CustomerNovels, error) {
	_node, _spec := cnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cnc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (cnc *CustomerNovelsCreate) createSpec() (*CustomerNovels, *sqlgraph.CreateSpec) {
	var (
		_node = &CustomerNovels{config: cnc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: customernovels.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: customernovels.FieldID,
			},
		}
	)
	if value, ok := cnc.mutation.NovelId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldNovelId,
		})
		_node.NovelId = value
	}
	if value, ok := cnc.mutation.TypeId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: customernovels.FieldTypeId,
		})
		_node.TypeId = value
	}
	if value, ok := cnc.mutation.TypeCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldTypeCode,
		})
		_node.TypeCode = value
	}
	if value, ok := cnc.mutation.GroupCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldGroupCode,
		})
		_node.GroupCode = value
	}
	if value, ok := cnc.mutation.FieldName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldFieldName,
		})
		_node.FieldName = value
	}
	if value, ok := cnc.mutation.Cover(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldCover,
		})
		_node.Cover = value
	}
	if value, ok := cnc.mutation.OrderNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: customernovels.FieldOrderNum,
		})
		_node.OrderNum = value
	}
	if value, ok := cnc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := cnc.mutation.EffectTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customernovels.FieldEffectTime,
		})
		_node.EffectTime = value
	}
	if value, ok := cnc.mutation.ExpiredTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customernovels.FieldExpiredTime,
		})
		_node.ExpiredTime = value
	}
	if value, ok := cnc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customernovels.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cnc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customernovels.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cnc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := cnc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := cnc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldTenantId,
		})
		_node.TenantId = value
	}
	return _node, _spec
}

// CustomerNovelsCreateBulk is the builder for creating many CustomerNovels entities in bulk.
type CustomerNovelsCreateBulk struct {
	config
	builders []*CustomerNovelsCreate
}

// Save creates the CustomerNovels entities in the database.
func (cncb *CustomerNovelsCreateBulk) Save(ctx context.Context) ([]*CustomerNovels, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cncb.builders))
	nodes := make([]*CustomerNovels, len(cncb.builders))
	mutators := make([]Mutator, len(cncb.builders))
	for i := range cncb.builders {
		func(i int, root context.Context) {
			builder := cncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerNovelsMutation)
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
					_, err = mutators[i+1].Mutate(root, cncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cncb *CustomerNovelsCreateBulk) SaveX(ctx context.Context) []*CustomerNovels {
	v, err := cncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cncb *CustomerNovelsCreateBulk) Exec(ctx context.Context) error {
	_, err := cncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cncb *CustomerNovelsCreateBulk) ExecX(ctx context.Context) {
	if err := cncb.Exec(ctx); err != nil {
		panic(err)
	}
}
