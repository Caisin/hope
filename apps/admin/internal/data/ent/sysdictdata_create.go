// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/sysdictdata"
	"hope/apps/admin/internal/data/ent/sysdicttype"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysDictDataCreate is the builder for creating a SysDictData entity.
type SysDictDataCreate struct {
	config
	mutation *SysDictDataMutation
	hooks    []Hook
}

// SetTypeId sets the "typeId" field.
func (sddc *SysDictDataCreate) SetTypeId(i int64) *SysDictDataCreate {
	sddc.mutation.SetTypeId(i)
	return sddc
}

// SetTypeCode sets the "typeCode" field.
func (sddc *SysDictDataCreate) SetTypeCode(s string) *SysDictDataCreate {
	sddc.mutation.SetTypeCode(s)
	return sddc
}

// SetDictSort sets the "dictSort" field.
func (sddc *SysDictDataCreate) SetDictSort(i int32) *SysDictDataCreate {
	sddc.mutation.SetDictSort(i)
	return sddc
}

// SetNillableDictSort sets the "dictSort" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableDictSort(i *int32) *SysDictDataCreate {
	if i != nil {
		sddc.SetDictSort(*i)
	}
	return sddc
}

// SetDictLabel sets the "dictLabel" field.
func (sddc *SysDictDataCreate) SetDictLabel(s string) *SysDictDataCreate {
	sddc.mutation.SetDictLabel(s)
	return sddc
}

// SetNillableDictLabel sets the "dictLabel" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableDictLabel(s *string) *SysDictDataCreate {
	if s != nil {
		sddc.SetDictLabel(*s)
	}
	return sddc
}

// SetDictValue sets the "dictValue" field.
func (sddc *SysDictDataCreate) SetDictValue(s string) *SysDictDataCreate {
	sddc.mutation.SetDictValue(s)
	return sddc
}

// SetNillableDictValue sets the "dictValue" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableDictValue(s *string) *SysDictDataCreate {
	if s != nil {
		sddc.SetDictValue(*s)
	}
	return sddc
}

// SetIsDefault sets the "isDefault" field.
func (sddc *SysDictDataCreate) SetIsDefault(s string) *SysDictDataCreate {
	sddc.mutation.SetIsDefault(s)
	return sddc
}

// SetNillableIsDefault sets the "isDefault" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableIsDefault(s *string) *SysDictDataCreate {
	if s != nil {
		sddc.SetIsDefault(*s)
	}
	return sddc
}

// SetStatus sets the "status" field.
func (sddc *SysDictDataCreate) SetStatus(i int32) *SysDictDataCreate {
	sddc.mutation.SetStatus(i)
	return sddc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableStatus(i *int32) *SysDictDataCreate {
	if i != nil {
		sddc.SetStatus(*i)
	}
	return sddc
}

// SetDefault sets the "default" field.
func (sddc *SysDictDataCreate) SetDefault(s string) *SysDictDataCreate {
	sddc.mutation.SetDefault(s)
	return sddc
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableDefault(s *string) *SysDictDataCreate {
	if s != nil {
		sddc.SetDefault(*s)
	}
	return sddc
}

// SetRemark sets the "remark" field.
func (sddc *SysDictDataCreate) SetRemark(s string) *SysDictDataCreate {
	sddc.mutation.SetRemark(s)
	return sddc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableRemark(s *string) *SysDictDataCreate {
	if s != nil {
		sddc.SetRemark(*s)
	}
	return sddc
}

// SetCreatedAt sets the "createdAt" field.
func (sddc *SysDictDataCreate) SetCreatedAt(t time.Time) *SysDictDataCreate {
	sddc.mutation.SetCreatedAt(t)
	return sddc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableCreatedAt(t *time.Time) *SysDictDataCreate {
	if t != nil {
		sddc.SetCreatedAt(*t)
	}
	return sddc
}

// SetUpdatedAt sets the "updatedAt" field.
func (sddc *SysDictDataCreate) SetUpdatedAt(t time.Time) *SysDictDataCreate {
	sddc.mutation.SetUpdatedAt(t)
	return sddc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableUpdatedAt(t *time.Time) *SysDictDataCreate {
	if t != nil {
		sddc.SetUpdatedAt(*t)
	}
	return sddc
}

// SetCreateBy sets the "createBy" field.
func (sddc *SysDictDataCreate) SetCreateBy(i int64) *SysDictDataCreate {
	sddc.mutation.SetCreateBy(i)
	return sddc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableCreateBy(i *int64) *SysDictDataCreate {
	if i != nil {
		sddc.SetCreateBy(*i)
	}
	return sddc
}

// SetUpdateBy sets the "updateBy" field.
func (sddc *SysDictDataCreate) SetUpdateBy(i int64) *SysDictDataCreate {
	sddc.mutation.SetUpdateBy(i)
	return sddc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableUpdateBy(i *int64) *SysDictDataCreate {
	if i != nil {
		sddc.SetUpdateBy(*i)
	}
	return sddc
}

// SetTenantId sets the "tenantId" field.
func (sddc *SysDictDataCreate) SetTenantId(i int64) *SysDictDataCreate {
	sddc.mutation.SetTenantId(i)
	return sddc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (sddc *SysDictDataCreate) SetNillableTenantId(i *int64) *SysDictDataCreate {
	if i != nil {
		sddc.SetTenantId(*i)
	}
	return sddc
}

// SetDictTypeID sets the "dictType" edge to the SysDictType entity by ID.
func (sddc *SysDictDataCreate) SetDictTypeID(id int64) *SysDictDataCreate {
	sddc.mutation.SetDictTypeID(id)
	return sddc
}

// SetDictType sets the "dictType" edge to the SysDictType entity.
func (sddc *SysDictDataCreate) SetDictType(s *SysDictType) *SysDictDataCreate {
	return sddc.SetDictTypeID(s.ID)
}

// Mutation returns the SysDictDataMutation object of the builder.
func (sddc *SysDictDataCreate) Mutation() *SysDictDataMutation {
	return sddc.mutation
}

// Save creates the SysDictData in the database.
func (sddc *SysDictDataCreate) Save(ctx context.Context) (*SysDictData, error) {
	var (
		err  error
		node *SysDictData
	)
	sddc.defaults()
	if len(sddc.hooks) == 0 {
		if err = sddc.check(); err != nil {
			return nil, err
		}
		node, err = sddc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sddc.check(); err != nil {
				return nil, err
			}
			sddc.mutation = mutation
			if node, err = sddc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sddc.hooks) - 1; i >= 0; i-- {
			if sddc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sddc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sddc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sddc *SysDictDataCreate) SaveX(ctx context.Context) *SysDictData {
	v, err := sddc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sddc *SysDictDataCreate) Exec(ctx context.Context) error {
	_, err := sddc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sddc *SysDictDataCreate) ExecX(ctx context.Context) {
	if err := sddc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sddc *SysDictDataCreate) defaults() {
	if _, ok := sddc.mutation.CreatedAt(); !ok {
		v := sysdictdata.DefaultCreatedAt()
		sddc.mutation.SetCreatedAt(v)
	}
	if _, ok := sddc.mutation.UpdatedAt(); !ok {
		v := sysdictdata.DefaultUpdatedAt()
		sddc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sddc.mutation.CreateBy(); !ok {
		v := sysdictdata.DefaultCreateBy
		sddc.mutation.SetCreateBy(v)
	}
	if _, ok := sddc.mutation.UpdateBy(); !ok {
		v := sysdictdata.DefaultUpdateBy
		sddc.mutation.SetUpdateBy(v)
	}
	if _, ok := sddc.mutation.TenantId(); !ok {
		v := sysdictdata.DefaultTenantId
		sddc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sddc *SysDictDataCreate) check() error {
	if _, ok := sddc.mutation.TypeId(); !ok {
		return &ValidationError{Name: "typeId", err: errors.New(`ent: missing required field "typeId"`)}
	}
	if _, ok := sddc.mutation.TypeCode(); !ok {
		return &ValidationError{Name: "typeCode", err: errors.New(`ent: missing required field "typeCode"`)}
	}
	if _, ok := sddc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := sddc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := sddc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := sddc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := sddc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	if _, ok := sddc.mutation.DictTypeID(); !ok {
		return &ValidationError{Name: "dictType", err: errors.New("ent: missing required edge \"dictType\"")}
	}
	return nil
}

func (sddc *SysDictDataCreate) sqlSave(ctx context.Context) (*SysDictData, error) {
	_node, _spec := sddc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sddc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (sddc *SysDictDataCreate) createSpec() (*SysDictData, *sqlgraph.CreateSpec) {
	var (
		_node = &SysDictData{config: sddc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sysdictdata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysdictdata.FieldID,
			},
		}
	)
	if value, ok := sddc.mutation.TypeCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldTypeCode,
		})
		_node.TypeCode = value
	}
	if value, ok := sddc.mutation.DictSort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictdata.FieldDictSort,
		})
		_node.DictSort = value
	}
	if value, ok := sddc.mutation.DictLabel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldDictLabel,
		})
		_node.DictLabel = value
	}
	if value, ok := sddc.mutation.DictValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldDictValue,
		})
		_node.DictValue = value
	}
	if value, ok := sddc.mutation.IsDefault(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldIsDefault,
		})
		_node.IsDefault = value
	}
	if value, ok := sddc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictdata.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := sddc.mutation.Default(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldDefault,
		})
		_node.Default = value
	}
	if value, ok := sddc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := sddc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictdata.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sddc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictdata.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sddc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := sddc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := sddc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := sddc.mutation.DictTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysdictdata.DictTypeTable,
			Columns: []string{sysdictdata.DictTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysdicttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TypeId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SysDictDataCreateBulk is the builder for creating many SysDictData entities in bulk.
type SysDictDataCreateBulk struct {
	config
	builders []*SysDictDataCreate
}

// Save creates the SysDictData entities in the database.
func (sddcb *SysDictDataCreateBulk) Save(ctx context.Context) ([]*SysDictData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sddcb.builders))
	nodes := make([]*SysDictData, len(sddcb.builders))
	mutators := make([]Mutator, len(sddcb.builders))
	for i := range sddcb.builders {
		func(i int, root context.Context) {
			builder := sddcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysDictDataMutation)
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
					_, err = mutators[i+1].Mutate(root, sddcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sddcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sddcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sddcb *SysDictDataCreateBulk) SaveX(ctx context.Context) []*SysDictData {
	v, err := sddcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sddcb *SysDictDataCreateBulk) Exec(ctx context.Context) error {
	_, err := sddcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sddcb *SysDictDataCreateBulk) ExecX(ctx context.Context) {
	if err := sddcb.Exec(ctx); err != nil {
		panic(err)
	}
}
