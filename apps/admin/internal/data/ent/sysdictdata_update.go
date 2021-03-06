// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysdictdata"
	"hope/apps/admin/internal/data/ent/sysdicttype"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysDictDataUpdate is the builder for updating SysDictData entities.
type SysDictDataUpdate struct {
	config
	hooks    []Hook
	mutation *SysDictDataMutation
}

// Where appends a list predicates to the SysDictDataUpdate builder.
func (sddu *SysDictDataUpdate) Where(ps ...predicate.SysDictData) *SysDictDataUpdate {
	sddu.mutation.Where(ps...)
	return sddu
}

// SetTypeId sets the "typeId" field.
func (sddu *SysDictDataUpdate) SetTypeId(i int64) *SysDictDataUpdate {
	sddu.mutation.SetTypeId(i)
	return sddu
}

// SetTypeCode sets the "typeCode" field.
func (sddu *SysDictDataUpdate) SetTypeCode(s string) *SysDictDataUpdate {
	sddu.mutation.SetTypeCode(s)
	return sddu
}

// SetDictSort sets the "dictSort" field.
func (sddu *SysDictDataUpdate) SetDictSort(i int32) *SysDictDataUpdate {
	sddu.mutation.ResetDictSort()
	sddu.mutation.SetDictSort(i)
	return sddu
}

// SetNillableDictSort sets the "dictSort" field if the given value is not nil.
func (sddu *SysDictDataUpdate) SetNillableDictSort(i *int32) *SysDictDataUpdate {
	if i != nil {
		sddu.SetDictSort(*i)
	}
	return sddu
}

// AddDictSort adds i to the "dictSort" field.
func (sddu *SysDictDataUpdate) AddDictSort(i int32) *SysDictDataUpdate {
	sddu.mutation.AddDictSort(i)
	return sddu
}

// ClearDictSort clears the value of the "dictSort" field.
func (sddu *SysDictDataUpdate) ClearDictSort() *SysDictDataUpdate {
	sddu.mutation.ClearDictSort()
	return sddu
}

// SetDictLabel sets the "dictLabel" field.
func (sddu *SysDictDataUpdate) SetDictLabel(s string) *SysDictDataUpdate {
	sddu.mutation.SetDictLabel(s)
	return sddu
}

// SetNillableDictLabel sets the "dictLabel" field if the given value is not nil.
func (sddu *SysDictDataUpdate) SetNillableDictLabel(s *string) *SysDictDataUpdate {
	if s != nil {
		sddu.SetDictLabel(*s)
	}
	return sddu
}

// ClearDictLabel clears the value of the "dictLabel" field.
func (sddu *SysDictDataUpdate) ClearDictLabel() *SysDictDataUpdate {
	sddu.mutation.ClearDictLabel()
	return sddu
}

// SetDictValue sets the "dictValue" field.
func (sddu *SysDictDataUpdate) SetDictValue(s string) *SysDictDataUpdate {
	sddu.mutation.SetDictValue(s)
	return sddu
}

// SetNillableDictValue sets the "dictValue" field if the given value is not nil.
func (sddu *SysDictDataUpdate) SetNillableDictValue(s *string) *SysDictDataUpdate {
	if s != nil {
		sddu.SetDictValue(*s)
	}
	return sddu
}

// ClearDictValue clears the value of the "dictValue" field.
func (sddu *SysDictDataUpdate) ClearDictValue() *SysDictDataUpdate {
	sddu.mutation.ClearDictValue()
	return sddu
}

// SetIsDefault sets the "isDefault" field.
func (sddu *SysDictDataUpdate) SetIsDefault(s string) *SysDictDataUpdate {
	sddu.mutation.SetIsDefault(s)
	return sddu
}

// SetNillableIsDefault sets the "isDefault" field if the given value is not nil.
func (sddu *SysDictDataUpdate) SetNillableIsDefault(s *string) *SysDictDataUpdate {
	if s != nil {
		sddu.SetIsDefault(*s)
	}
	return sddu
}

// ClearIsDefault clears the value of the "isDefault" field.
func (sddu *SysDictDataUpdate) ClearIsDefault() *SysDictDataUpdate {
	sddu.mutation.ClearIsDefault()
	return sddu
}

// SetStatus sets the "status" field.
func (sddu *SysDictDataUpdate) SetStatus(i int32) *SysDictDataUpdate {
	sddu.mutation.ResetStatus()
	sddu.mutation.SetStatus(i)
	return sddu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sddu *SysDictDataUpdate) SetNillableStatus(i *int32) *SysDictDataUpdate {
	if i != nil {
		sddu.SetStatus(*i)
	}
	return sddu
}

// AddStatus adds i to the "status" field.
func (sddu *SysDictDataUpdate) AddStatus(i int32) *SysDictDataUpdate {
	sddu.mutation.AddStatus(i)
	return sddu
}

// ClearStatus clears the value of the "status" field.
func (sddu *SysDictDataUpdate) ClearStatus() *SysDictDataUpdate {
	sddu.mutation.ClearStatus()
	return sddu
}

// SetDefault sets the "default" field.
func (sddu *SysDictDataUpdate) SetDefault(s string) *SysDictDataUpdate {
	sddu.mutation.SetDefault(s)
	return sddu
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (sddu *SysDictDataUpdate) SetNillableDefault(s *string) *SysDictDataUpdate {
	if s != nil {
		sddu.SetDefault(*s)
	}
	return sddu
}

// ClearDefault clears the value of the "default" field.
func (sddu *SysDictDataUpdate) ClearDefault() *SysDictDataUpdate {
	sddu.mutation.ClearDefault()
	return sddu
}

// SetRemark sets the "remark" field.
func (sddu *SysDictDataUpdate) SetRemark(s string) *SysDictDataUpdate {
	sddu.mutation.SetRemark(s)
	return sddu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sddu *SysDictDataUpdate) SetNillableRemark(s *string) *SysDictDataUpdate {
	if s != nil {
		sddu.SetRemark(*s)
	}
	return sddu
}

// ClearRemark clears the value of the "remark" field.
func (sddu *SysDictDataUpdate) ClearRemark() *SysDictDataUpdate {
	sddu.mutation.ClearRemark()
	return sddu
}

// SetUpdatedAt sets the "updatedAt" field.
func (sddu *SysDictDataUpdate) SetUpdatedAt(t time.Time) *SysDictDataUpdate {
	sddu.mutation.SetUpdatedAt(t)
	return sddu
}

// SetCreateBy sets the "createBy" field.
func (sddu *SysDictDataUpdate) SetCreateBy(i int64) *SysDictDataUpdate {
	sddu.mutation.ResetCreateBy()
	sddu.mutation.SetCreateBy(i)
	return sddu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (sddu *SysDictDataUpdate) SetNillableCreateBy(i *int64) *SysDictDataUpdate {
	if i != nil {
		sddu.SetCreateBy(*i)
	}
	return sddu
}

// AddCreateBy adds i to the "createBy" field.
func (sddu *SysDictDataUpdate) AddCreateBy(i int64) *SysDictDataUpdate {
	sddu.mutation.AddCreateBy(i)
	return sddu
}

// SetUpdateBy sets the "updateBy" field.
func (sddu *SysDictDataUpdate) SetUpdateBy(i int64) *SysDictDataUpdate {
	sddu.mutation.ResetUpdateBy()
	sddu.mutation.SetUpdateBy(i)
	return sddu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (sddu *SysDictDataUpdate) SetNillableUpdateBy(i *int64) *SysDictDataUpdate {
	if i != nil {
		sddu.SetUpdateBy(*i)
	}
	return sddu
}

// AddUpdateBy adds i to the "updateBy" field.
func (sddu *SysDictDataUpdate) AddUpdateBy(i int64) *SysDictDataUpdate {
	sddu.mutation.AddUpdateBy(i)
	return sddu
}

// SetTenantId sets the "tenantId" field.
func (sddu *SysDictDataUpdate) SetTenantId(i int64) *SysDictDataUpdate {
	sddu.mutation.ResetTenantId()
	sddu.mutation.SetTenantId(i)
	return sddu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (sddu *SysDictDataUpdate) SetNillableTenantId(i *int64) *SysDictDataUpdate {
	if i != nil {
		sddu.SetTenantId(*i)
	}
	return sddu
}

// AddTenantId adds i to the "tenantId" field.
func (sddu *SysDictDataUpdate) AddTenantId(i int64) *SysDictDataUpdate {
	sddu.mutation.AddTenantId(i)
	return sddu
}

// SetDictTypeID sets the "dictType" edge to the SysDictType entity by ID.
func (sddu *SysDictDataUpdate) SetDictTypeID(id int64) *SysDictDataUpdate {
	sddu.mutation.SetDictTypeID(id)
	return sddu
}

// SetDictType sets the "dictType" edge to the SysDictType entity.
func (sddu *SysDictDataUpdate) SetDictType(s *SysDictType) *SysDictDataUpdate {
	return sddu.SetDictTypeID(s.ID)
}

// Mutation returns the SysDictDataMutation object of the builder.
func (sddu *SysDictDataUpdate) Mutation() *SysDictDataMutation {
	return sddu.mutation
}

// ClearDictType clears the "dictType" edge to the SysDictType entity.
func (sddu *SysDictDataUpdate) ClearDictType() *SysDictDataUpdate {
	sddu.mutation.ClearDictType()
	return sddu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sddu *SysDictDataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sddu.defaults()
	if len(sddu.hooks) == 0 {
		if err = sddu.check(); err != nil {
			return 0, err
		}
		affected, err = sddu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sddu.check(); err != nil {
				return 0, err
			}
			sddu.mutation = mutation
			affected, err = sddu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sddu.hooks) - 1; i >= 0; i-- {
			if sddu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sddu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sddu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sddu *SysDictDataUpdate) SaveX(ctx context.Context) int {
	affected, err := sddu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sddu *SysDictDataUpdate) Exec(ctx context.Context) error {
	_, err := sddu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sddu *SysDictDataUpdate) ExecX(ctx context.Context) {
	if err := sddu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sddu *SysDictDataUpdate) defaults() {
	if _, ok := sddu.mutation.UpdatedAt(); !ok {
		v := sysdictdata.UpdateDefaultUpdatedAt()
		sddu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sddu *SysDictDataUpdate) check() error {
	if _, ok := sddu.mutation.DictTypeID(); sddu.mutation.DictTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SysDictData.dictType"`)
	}
	return nil
}

func (sddu *SysDictDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdictdata.Table,
			Columns: sysdictdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysdictdata.FieldID,
			},
		},
	}
	if ps := sddu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sddu.mutation.TypeCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldTypeCode,
		})
	}
	if value, ok := sddu.mutation.DictSort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictdata.FieldDictSort,
		})
	}
	if value, ok := sddu.mutation.AddedDictSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictdata.FieldDictSort,
		})
	}
	if sddu.mutation.DictSortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: sysdictdata.FieldDictSort,
		})
	}
	if value, ok := sddu.mutation.DictLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldDictLabel,
		})
	}
	if sddu.mutation.DictLabelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdictdata.FieldDictLabel,
		})
	}
	if value, ok := sddu.mutation.DictValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldDictValue,
		})
	}
	if sddu.mutation.DictValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdictdata.FieldDictValue,
		})
	}
	if value, ok := sddu.mutation.IsDefault(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldIsDefault,
		})
	}
	if sddu.mutation.IsDefaultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdictdata.FieldIsDefault,
		})
	}
	if value, ok := sddu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictdata.FieldStatus,
		})
	}
	if value, ok := sddu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictdata.FieldStatus,
		})
	}
	if sddu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: sysdictdata.FieldStatus,
		})
	}
	if value, ok := sddu.mutation.Default(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldDefault,
		})
	}
	if sddu.mutation.DefaultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdictdata.FieldDefault,
		})
	}
	if value, ok := sddu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldRemark,
		})
	}
	if sddu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdictdata.FieldRemark,
		})
	}
	if value, ok := sddu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictdata.FieldUpdatedAt,
		})
	}
	if value, ok := sddu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldCreateBy,
		})
	}
	if value, ok := sddu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldCreateBy,
		})
	}
	if value, ok := sddu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldUpdateBy,
		})
	}
	if value, ok := sddu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldUpdateBy,
		})
	}
	if value, ok := sddu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldTenantId,
		})
	}
	if value, ok := sddu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldTenantId,
		})
	}
	if sddu.mutation.DictTypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sddu.mutation.DictTypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sddu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdictdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SysDictDataUpdateOne is the builder for updating a single SysDictData entity.
type SysDictDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysDictDataMutation
}

// SetTypeId sets the "typeId" field.
func (sdduo *SysDictDataUpdateOne) SetTypeId(i int64) *SysDictDataUpdateOne {
	sdduo.mutation.SetTypeId(i)
	return sdduo
}

// SetTypeCode sets the "typeCode" field.
func (sdduo *SysDictDataUpdateOne) SetTypeCode(s string) *SysDictDataUpdateOne {
	sdduo.mutation.SetTypeCode(s)
	return sdduo
}

// SetDictSort sets the "dictSort" field.
func (sdduo *SysDictDataUpdateOne) SetDictSort(i int32) *SysDictDataUpdateOne {
	sdduo.mutation.ResetDictSort()
	sdduo.mutation.SetDictSort(i)
	return sdduo
}

// SetNillableDictSort sets the "dictSort" field if the given value is not nil.
func (sdduo *SysDictDataUpdateOne) SetNillableDictSort(i *int32) *SysDictDataUpdateOne {
	if i != nil {
		sdduo.SetDictSort(*i)
	}
	return sdduo
}

// AddDictSort adds i to the "dictSort" field.
func (sdduo *SysDictDataUpdateOne) AddDictSort(i int32) *SysDictDataUpdateOne {
	sdduo.mutation.AddDictSort(i)
	return sdduo
}

// ClearDictSort clears the value of the "dictSort" field.
func (sdduo *SysDictDataUpdateOne) ClearDictSort() *SysDictDataUpdateOne {
	sdduo.mutation.ClearDictSort()
	return sdduo
}

// SetDictLabel sets the "dictLabel" field.
func (sdduo *SysDictDataUpdateOne) SetDictLabel(s string) *SysDictDataUpdateOne {
	sdduo.mutation.SetDictLabel(s)
	return sdduo
}

// SetNillableDictLabel sets the "dictLabel" field if the given value is not nil.
func (sdduo *SysDictDataUpdateOne) SetNillableDictLabel(s *string) *SysDictDataUpdateOne {
	if s != nil {
		sdduo.SetDictLabel(*s)
	}
	return sdduo
}

// ClearDictLabel clears the value of the "dictLabel" field.
func (sdduo *SysDictDataUpdateOne) ClearDictLabel() *SysDictDataUpdateOne {
	sdduo.mutation.ClearDictLabel()
	return sdduo
}

// SetDictValue sets the "dictValue" field.
func (sdduo *SysDictDataUpdateOne) SetDictValue(s string) *SysDictDataUpdateOne {
	sdduo.mutation.SetDictValue(s)
	return sdduo
}

// SetNillableDictValue sets the "dictValue" field if the given value is not nil.
func (sdduo *SysDictDataUpdateOne) SetNillableDictValue(s *string) *SysDictDataUpdateOne {
	if s != nil {
		sdduo.SetDictValue(*s)
	}
	return sdduo
}

// ClearDictValue clears the value of the "dictValue" field.
func (sdduo *SysDictDataUpdateOne) ClearDictValue() *SysDictDataUpdateOne {
	sdduo.mutation.ClearDictValue()
	return sdduo
}

// SetIsDefault sets the "isDefault" field.
func (sdduo *SysDictDataUpdateOne) SetIsDefault(s string) *SysDictDataUpdateOne {
	sdduo.mutation.SetIsDefault(s)
	return sdduo
}

// SetNillableIsDefault sets the "isDefault" field if the given value is not nil.
func (sdduo *SysDictDataUpdateOne) SetNillableIsDefault(s *string) *SysDictDataUpdateOne {
	if s != nil {
		sdduo.SetIsDefault(*s)
	}
	return sdduo
}

// ClearIsDefault clears the value of the "isDefault" field.
func (sdduo *SysDictDataUpdateOne) ClearIsDefault() *SysDictDataUpdateOne {
	sdduo.mutation.ClearIsDefault()
	return sdduo
}

// SetStatus sets the "status" field.
func (sdduo *SysDictDataUpdateOne) SetStatus(i int32) *SysDictDataUpdateOne {
	sdduo.mutation.ResetStatus()
	sdduo.mutation.SetStatus(i)
	return sdduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sdduo *SysDictDataUpdateOne) SetNillableStatus(i *int32) *SysDictDataUpdateOne {
	if i != nil {
		sdduo.SetStatus(*i)
	}
	return sdduo
}

// AddStatus adds i to the "status" field.
func (sdduo *SysDictDataUpdateOne) AddStatus(i int32) *SysDictDataUpdateOne {
	sdduo.mutation.AddStatus(i)
	return sdduo
}

// ClearStatus clears the value of the "status" field.
func (sdduo *SysDictDataUpdateOne) ClearStatus() *SysDictDataUpdateOne {
	sdduo.mutation.ClearStatus()
	return sdduo
}

// SetDefault sets the "default" field.
func (sdduo *SysDictDataUpdateOne) SetDefault(s string) *SysDictDataUpdateOne {
	sdduo.mutation.SetDefault(s)
	return sdduo
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (sdduo *SysDictDataUpdateOne) SetNillableDefault(s *string) *SysDictDataUpdateOne {
	if s != nil {
		sdduo.SetDefault(*s)
	}
	return sdduo
}

// ClearDefault clears the value of the "default" field.
func (sdduo *SysDictDataUpdateOne) ClearDefault() *SysDictDataUpdateOne {
	sdduo.mutation.ClearDefault()
	return sdduo
}

// SetRemark sets the "remark" field.
func (sdduo *SysDictDataUpdateOne) SetRemark(s string) *SysDictDataUpdateOne {
	sdduo.mutation.SetRemark(s)
	return sdduo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sdduo *SysDictDataUpdateOne) SetNillableRemark(s *string) *SysDictDataUpdateOne {
	if s != nil {
		sdduo.SetRemark(*s)
	}
	return sdduo
}

// ClearRemark clears the value of the "remark" field.
func (sdduo *SysDictDataUpdateOne) ClearRemark() *SysDictDataUpdateOne {
	sdduo.mutation.ClearRemark()
	return sdduo
}

// SetUpdatedAt sets the "updatedAt" field.
func (sdduo *SysDictDataUpdateOne) SetUpdatedAt(t time.Time) *SysDictDataUpdateOne {
	sdduo.mutation.SetUpdatedAt(t)
	return sdduo
}

// SetCreateBy sets the "createBy" field.
func (sdduo *SysDictDataUpdateOne) SetCreateBy(i int64) *SysDictDataUpdateOne {
	sdduo.mutation.ResetCreateBy()
	sdduo.mutation.SetCreateBy(i)
	return sdduo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (sdduo *SysDictDataUpdateOne) SetNillableCreateBy(i *int64) *SysDictDataUpdateOne {
	if i != nil {
		sdduo.SetCreateBy(*i)
	}
	return sdduo
}

// AddCreateBy adds i to the "createBy" field.
func (sdduo *SysDictDataUpdateOne) AddCreateBy(i int64) *SysDictDataUpdateOne {
	sdduo.mutation.AddCreateBy(i)
	return sdduo
}

// SetUpdateBy sets the "updateBy" field.
func (sdduo *SysDictDataUpdateOne) SetUpdateBy(i int64) *SysDictDataUpdateOne {
	sdduo.mutation.ResetUpdateBy()
	sdduo.mutation.SetUpdateBy(i)
	return sdduo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (sdduo *SysDictDataUpdateOne) SetNillableUpdateBy(i *int64) *SysDictDataUpdateOne {
	if i != nil {
		sdduo.SetUpdateBy(*i)
	}
	return sdduo
}

// AddUpdateBy adds i to the "updateBy" field.
func (sdduo *SysDictDataUpdateOne) AddUpdateBy(i int64) *SysDictDataUpdateOne {
	sdduo.mutation.AddUpdateBy(i)
	return sdduo
}

// SetTenantId sets the "tenantId" field.
func (sdduo *SysDictDataUpdateOne) SetTenantId(i int64) *SysDictDataUpdateOne {
	sdduo.mutation.ResetTenantId()
	sdduo.mutation.SetTenantId(i)
	return sdduo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (sdduo *SysDictDataUpdateOne) SetNillableTenantId(i *int64) *SysDictDataUpdateOne {
	if i != nil {
		sdduo.SetTenantId(*i)
	}
	return sdduo
}

// AddTenantId adds i to the "tenantId" field.
func (sdduo *SysDictDataUpdateOne) AddTenantId(i int64) *SysDictDataUpdateOne {
	sdduo.mutation.AddTenantId(i)
	return sdduo
}

// SetDictTypeID sets the "dictType" edge to the SysDictType entity by ID.
func (sdduo *SysDictDataUpdateOne) SetDictTypeID(id int64) *SysDictDataUpdateOne {
	sdduo.mutation.SetDictTypeID(id)
	return sdduo
}

// SetDictType sets the "dictType" edge to the SysDictType entity.
func (sdduo *SysDictDataUpdateOne) SetDictType(s *SysDictType) *SysDictDataUpdateOne {
	return sdduo.SetDictTypeID(s.ID)
}

// Mutation returns the SysDictDataMutation object of the builder.
func (sdduo *SysDictDataUpdateOne) Mutation() *SysDictDataMutation {
	return sdduo.mutation
}

// ClearDictType clears the "dictType" edge to the SysDictType entity.
func (sdduo *SysDictDataUpdateOne) ClearDictType() *SysDictDataUpdateOne {
	sdduo.mutation.ClearDictType()
	return sdduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sdduo *SysDictDataUpdateOne) Select(field string, fields ...string) *SysDictDataUpdateOne {
	sdduo.fields = append([]string{field}, fields...)
	return sdduo
}

// Save executes the query and returns the updated SysDictData entity.
func (sdduo *SysDictDataUpdateOne) Save(ctx context.Context) (*SysDictData, error) {
	var (
		err  error
		node *SysDictData
	)
	sdduo.defaults()
	if len(sdduo.hooks) == 0 {
		if err = sdduo.check(); err != nil {
			return nil, err
		}
		node, err = sdduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdduo.check(); err != nil {
				return nil, err
			}
			sdduo.mutation = mutation
			node, err = sdduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sdduo.hooks) - 1; i >= 0; i-- {
			if sdduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sdduo *SysDictDataUpdateOne) SaveX(ctx context.Context) *SysDictData {
	node, err := sdduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sdduo *SysDictDataUpdateOne) Exec(ctx context.Context) error {
	_, err := sdduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdduo *SysDictDataUpdateOne) ExecX(ctx context.Context) {
	if err := sdduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdduo *SysDictDataUpdateOne) defaults() {
	if _, ok := sdduo.mutation.UpdatedAt(); !ok {
		v := sysdictdata.UpdateDefaultUpdatedAt()
		sdduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdduo *SysDictDataUpdateOne) check() error {
	if _, ok := sdduo.mutation.DictTypeID(); sdduo.mutation.DictTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SysDictData.dictType"`)
	}
	return nil
}

func (sdduo *SysDictDataUpdateOne) sqlSave(ctx context.Context) (_node *SysDictData, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdictdata.Table,
			Columns: sysdictdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysdictdata.FieldID,
			},
		},
	}
	id, ok := sdduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysDictData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sdduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdictdata.FieldID)
		for _, f := range fields {
			if !sysdictdata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysdictdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sdduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdduo.mutation.TypeCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldTypeCode,
		})
	}
	if value, ok := sdduo.mutation.DictSort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictdata.FieldDictSort,
		})
	}
	if value, ok := sdduo.mutation.AddedDictSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictdata.FieldDictSort,
		})
	}
	if sdduo.mutation.DictSortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: sysdictdata.FieldDictSort,
		})
	}
	if value, ok := sdduo.mutation.DictLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldDictLabel,
		})
	}
	if sdduo.mutation.DictLabelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdictdata.FieldDictLabel,
		})
	}
	if value, ok := sdduo.mutation.DictValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldDictValue,
		})
	}
	if sdduo.mutation.DictValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdictdata.FieldDictValue,
		})
	}
	if value, ok := sdduo.mutation.IsDefault(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldIsDefault,
		})
	}
	if sdduo.mutation.IsDefaultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdictdata.FieldIsDefault,
		})
	}
	if value, ok := sdduo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictdata.FieldStatus,
		})
	}
	if value, ok := sdduo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdictdata.FieldStatus,
		})
	}
	if sdduo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: sysdictdata.FieldStatus,
		})
	}
	if value, ok := sdduo.mutation.Default(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldDefault,
		})
	}
	if sdduo.mutation.DefaultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdictdata.FieldDefault,
		})
	}
	if value, ok := sdduo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdictdata.FieldRemark,
		})
	}
	if sdduo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdictdata.FieldRemark,
		})
	}
	if value, ok := sdduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdictdata.FieldUpdatedAt,
		})
	}
	if value, ok := sdduo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldCreateBy,
		})
	}
	if value, ok := sdduo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldCreateBy,
		})
	}
	if value, ok := sdduo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldUpdateBy,
		})
	}
	if value, ok := sdduo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldUpdateBy,
		})
	}
	if value, ok := sdduo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldTenantId,
		})
	}
	if value, ok := sdduo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdictdata.FieldTenantId,
		})
	}
	if sdduo.mutation.DictTypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdduo.mutation.DictTypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SysDictData{config: sdduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sdduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdictdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
