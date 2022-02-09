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

// SysDictTypeUpdate is the builder for updating SysDictType entities.
type SysDictTypeUpdate struct {
	config
	hooks    []Hook
	mutation *SysDictTypeMutation
}

// Where appends a list predicates to the SysDictTypeUpdate builder.
func (sdtu *SysDictTypeUpdate) Where(ps ...predicate.SysDictType) *SysDictTypeUpdate {
	sdtu.mutation.Where(ps...)
	return sdtu
}

// SetDictName sets the "dictName" field.
func (sdtu *SysDictTypeUpdate) SetDictName(s string) *SysDictTypeUpdate {
	sdtu.mutation.SetDictName(s)
	return sdtu
}

// SetNillableDictName sets the "dictName" field if the given value is not nil.
func (sdtu *SysDictTypeUpdate) SetNillableDictName(s *string) *SysDictTypeUpdate {
	if s != nil {
		sdtu.SetDictName(*s)
	}
	return sdtu
}

// ClearDictName clears the value of the "dictName" field.
func (sdtu *SysDictTypeUpdate) ClearDictName() *SysDictTypeUpdate {
	sdtu.mutation.ClearDictName()
	return sdtu
}

// SetTypeCode sets the "typeCode" field.
func (sdtu *SysDictTypeUpdate) SetTypeCode(s string) *SysDictTypeUpdate {
	sdtu.mutation.SetTypeCode(s)
	return sdtu
}

// SetStatus sets the "status" field.
func (sdtu *SysDictTypeUpdate) SetStatus(i int32) *SysDictTypeUpdate {
	sdtu.mutation.ResetStatus()
	sdtu.mutation.SetStatus(i)
	return sdtu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sdtu *SysDictTypeUpdate) SetNillableStatus(i *int32) *SysDictTypeUpdate {
	if i != nil {
		sdtu.SetStatus(*i)
	}
	return sdtu
}

// AddStatus adds i to the "status" field.
func (sdtu *SysDictTypeUpdate) AddStatus(i int32) *SysDictTypeUpdate {
	sdtu.mutation.AddStatus(i)
	return sdtu
}

// ClearStatus clears the value of the "status" field.
func (sdtu *SysDictTypeUpdate) ClearStatus() *SysDictTypeUpdate {
	sdtu.mutation.ClearStatus()
	return sdtu
}

// SetRemark sets the "remark" field.
func (sdtu *SysDictTypeUpdate) SetRemark(s string) *SysDictTypeUpdate {
	sdtu.mutation.SetRemark(s)
	return sdtu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sdtu *SysDictTypeUpdate) SetNillableRemark(s *string) *SysDictTypeUpdate {
	if s != nil {
		sdtu.SetRemark(*s)
	}
	return sdtu
}

// ClearRemark clears the value of the "remark" field.
func (sdtu *SysDictTypeUpdate) ClearRemark() *SysDictTypeUpdate {
	sdtu.mutation.ClearRemark()
	return sdtu
}

// SetUpdatedAt sets the "updatedAt" field.
func (sdtu *SysDictTypeUpdate) SetUpdatedAt(t time.Time) *SysDictTypeUpdate {
	sdtu.mutation.SetUpdatedAt(t)
	return sdtu
}

// SetCreateBy sets the "createBy" field.
func (sdtu *SysDictTypeUpdate) SetCreateBy(i int64) *SysDictTypeUpdate {
	sdtu.mutation.ResetCreateBy()
	sdtu.mutation.SetCreateBy(i)
	return sdtu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (sdtu *SysDictTypeUpdate) SetNillableCreateBy(i *int64) *SysDictTypeUpdate {
	if i != nil {
		sdtu.SetCreateBy(*i)
	}
	return sdtu
}

// AddCreateBy adds i to the "createBy" field.
func (sdtu *SysDictTypeUpdate) AddCreateBy(i int64) *SysDictTypeUpdate {
	sdtu.mutation.AddCreateBy(i)
	return sdtu
}

// SetUpdateBy sets the "updateBy" field.
func (sdtu *SysDictTypeUpdate) SetUpdateBy(i int64) *SysDictTypeUpdate {
	sdtu.mutation.ResetUpdateBy()
	sdtu.mutation.SetUpdateBy(i)
	return sdtu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (sdtu *SysDictTypeUpdate) SetNillableUpdateBy(i *int64) *SysDictTypeUpdate {
	if i != nil {
		sdtu.SetUpdateBy(*i)
	}
	return sdtu
}

// AddUpdateBy adds i to the "updateBy" field.
func (sdtu *SysDictTypeUpdate) AddUpdateBy(i int64) *SysDictTypeUpdate {
	sdtu.mutation.AddUpdateBy(i)
	return sdtu
}

// SetTenantId sets the "tenantId" field.
func (sdtu *SysDictTypeUpdate) SetTenantId(i int64) *SysDictTypeUpdate {
	sdtu.mutation.ResetTenantId()
	sdtu.mutation.SetTenantId(i)
	return sdtu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (sdtu *SysDictTypeUpdate) SetNillableTenantId(i *int64) *SysDictTypeUpdate {
	if i != nil {
		sdtu.SetTenantId(*i)
	}
	return sdtu
}

// AddTenantId adds i to the "tenantId" field.
func (sdtu *SysDictTypeUpdate) AddTenantId(i int64) *SysDictTypeUpdate {
	sdtu.mutation.AddTenantId(i)
	return sdtu
}

// AddDataListIDs adds the "dataList" edge to the SysDictData entity by IDs.
func (sdtu *SysDictTypeUpdate) AddDataListIDs(ids ...int64) *SysDictTypeUpdate {
	sdtu.mutation.AddDataListIDs(ids...)
	return sdtu
}

// AddDataList adds the "dataList" edges to the SysDictData entity.
func (sdtu *SysDictTypeUpdate) AddDataList(s ...*SysDictData) *SysDictTypeUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdtu.AddDataListIDs(ids...)
}

// Mutation returns the SysDictTypeMutation object of the builder.
func (sdtu *SysDictTypeUpdate) Mutation() *SysDictTypeMutation {
	return sdtu.mutation
}

// ClearDataList clears all "dataList" edges to the SysDictData entity.
func (sdtu *SysDictTypeUpdate) ClearDataList() *SysDictTypeUpdate {
	sdtu.mutation.ClearDataList()
	return sdtu
}

// RemoveDataListIDs removes the "dataList" edge to SysDictData entities by IDs.
func (sdtu *SysDictTypeUpdate) RemoveDataListIDs(ids ...int64) *SysDictTypeUpdate {
	sdtu.mutation.RemoveDataListIDs(ids...)
	return sdtu
}

// RemoveDataList removes "dataList" edges to SysDictData entities.
func (sdtu *SysDictTypeUpdate) RemoveDataList(s ...*SysDictData) *SysDictTypeUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdtu.RemoveDataListIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sdtu *SysDictTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sdtu.defaults()
	if len(sdtu.hooks) == 0 {
		affected, err = sdtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdtu.mutation = mutation
			affected, err = sdtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdtu.hooks) - 1; i >= 0; i-- {
			if sdtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sdtu *SysDictTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := sdtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sdtu *SysDictTypeUpdate) Exec(ctx context.Context) error {
	_, err := sdtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdtu *SysDictTypeUpdate) ExecX(ctx context.Context) {
	if err := sdtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdtu *SysDictTypeUpdate) defaults() {
	if _, ok := sdtu.mutation.UpdatedAt(); !ok {
		v := sysdicttype.UpdateDefaultUpdatedAt()
		sdtu.mutation.SetUpdatedAt(v)
	}
}

func (sdtu *SysDictTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdicttype.Table,
			Columns: sysdicttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysdicttype.FieldID,
			},
		},
	}
	if ps := sdtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdtu.mutation.DictName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdicttype.FieldDictName,
		})
	}
	if sdtu.mutation.DictNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdicttype.FieldDictName,
		})
	}
	if value, ok := sdtu.mutation.TypeCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdicttype.FieldTypeCode,
		})
	}
	if value, ok := sdtu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdicttype.FieldStatus,
		})
	}
	if value, ok := sdtu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdicttype.FieldStatus,
		})
	}
	if sdtu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: sysdicttype.FieldStatus,
		})
	}
	if value, ok := sdtu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdicttype.FieldRemark,
		})
	}
	if sdtu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdicttype.FieldRemark,
		})
	}
	if value, ok := sdtu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdicttype.FieldUpdatedAt,
		})
	}
	if value, ok := sdtu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldCreateBy,
		})
	}
	if value, ok := sdtu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldCreateBy,
		})
	}
	if value, ok := sdtu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldUpdateBy,
		})
	}
	if value, ok := sdtu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldUpdateBy,
		})
	}
	if value, ok := sdtu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldTenantId,
		})
	}
	if value, ok := sdtu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldTenantId,
		})
	}
	if sdtu.mutation.DataListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysdicttype.DataListTable,
			Columns: []string{sysdicttype.DataListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysdictdata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdtu.mutation.RemovedDataListIDs(); len(nodes) > 0 && !sdtu.mutation.DataListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysdicttype.DataListTable,
			Columns: []string{sysdicttype.DataListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysdictdata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdtu.mutation.DataListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysdicttype.DataListTable,
			Columns: []string{sysdicttype.DataListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysdictdata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sdtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdicttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SysDictTypeUpdateOne is the builder for updating a single SysDictType entity.
type SysDictTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysDictTypeMutation
}

// SetDictName sets the "dictName" field.
func (sdtuo *SysDictTypeUpdateOne) SetDictName(s string) *SysDictTypeUpdateOne {
	sdtuo.mutation.SetDictName(s)
	return sdtuo
}

// SetNillableDictName sets the "dictName" field if the given value is not nil.
func (sdtuo *SysDictTypeUpdateOne) SetNillableDictName(s *string) *SysDictTypeUpdateOne {
	if s != nil {
		sdtuo.SetDictName(*s)
	}
	return sdtuo
}

// ClearDictName clears the value of the "dictName" field.
func (sdtuo *SysDictTypeUpdateOne) ClearDictName() *SysDictTypeUpdateOne {
	sdtuo.mutation.ClearDictName()
	return sdtuo
}

// SetTypeCode sets the "typeCode" field.
func (sdtuo *SysDictTypeUpdateOne) SetTypeCode(s string) *SysDictTypeUpdateOne {
	sdtuo.mutation.SetTypeCode(s)
	return sdtuo
}

// SetStatus sets the "status" field.
func (sdtuo *SysDictTypeUpdateOne) SetStatus(i int32) *SysDictTypeUpdateOne {
	sdtuo.mutation.ResetStatus()
	sdtuo.mutation.SetStatus(i)
	return sdtuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sdtuo *SysDictTypeUpdateOne) SetNillableStatus(i *int32) *SysDictTypeUpdateOne {
	if i != nil {
		sdtuo.SetStatus(*i)
	}
	return sdtuo
}

// AddStatus adds i to the "status" field.
func (sdtuo *SysDictTypeUpdateOne) AddStatus(i int32) *SysDictTypeUpdateOne {
	sdtuo.mutation.AddStatus(i)
	return sdtuo
}

// ClearStatus clears the value of the "status" field.
func (sdtuo *SysDictTypeUpdateOne) ClearStatus() *SysDictTypeUpdateOne {
	sdtuo.mutation.ClearStatus()
	return sdtuo
}

// SetRemark sets the "remark" field.
func (sdtuo *SysDictTypeUpdateOne) SetRemark(s string) *SysDictTypeUpdateOne {
	sdtuo.mutation.SetRemark(s)
	return sdtuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sdtuo *SysDictTypeUpdateOne) SetNillableRemark(s *string) *SysDictTypeUpdateOne {
	if s != nil {
		sdtuo.SetRemark(*s)
	}
	return sdtuo
}

// ClearRemark clears the value of the "remark" field.
func (sdtuo *SysDictTypeUpdateOne) ClearRemark() *SysDictTypeUpdateOne {
	sdtuo.mutation.ClearRemark()
	return sdtuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (sdtuo *SysDictTypeUpdateOne) SetUpdatedAt(t time.Time) *SysDictTypeUpdateOne {
	sdtuo.mutation.SetUpdatedAt(t)
	return sdtuo
}

// SetCreateBy sets the "createBy" field.
func (sdtuo *SysDictTypeUpdateOne) SetCreateBy(i int64) *SysDictTypeUpdateOne {
	sdtuo.mutation.ResetCreateBy()
	sdtuo.mutation.SetCreateBy(i)
	return sdtuo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (sdtuo *SysDictTypeUpdateOne) SetNillableCreateBy(i *int64) *SysDictTypeUpdateOne {
	if i != nil {
		sdtuo.SetCreateBy(*i)
	}
	return sdtuo
}

// AddCreateBy adds i to the "createBy" field.
func (sdtuo *SysDictTypeUpdateOne) AddCreateBy(i int64) *SysDictTypeUpdateOne {
	sdtuo.mutation.AddCreateBy(i)
	return sdtuo
}

// SetUpdateBy sets the "updateBy" field.
func (sdtuo *SysDictTypeUpdateOne) SetUpdateBy(i int64) *SysDictTypeUpdateOne {
	sdtuo.mutation.ResetUpdateBy()
	sdtuo.mutation.SetUpdateBy(i)
	return sdtuo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (sdtuo *SysDictTypeUpdateOne) SetNillableUpdateBy(i *int64) *SysDictTypeUpdateOne {
	if i != nil {
		sdtuo.SetUpdateBy(*i)
	}
	return sdtuo
}

// AddUpdateBy adds i to the "updateBy" field.
func (sdtuo *SysDictTypeUpdateOne) AddUpdateBy(i int64) *SysDictTypeUpdateOne {
	sdtuo.mutation.AddUpdateBy(i)
	return sdtuo
}

// SetTenantId sets the "tenantId" field.
func (sdtuo *SysDictTypeUpdateOne) SetTenantId(i int64) *SysDictTypeUpdateOne {
	sdtuo.mutation.ResetTenantId()
	sdtuo.mutation.SetTenantId(i)
	return sdtuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (sdtuo *SysDictTypeUpdateOne) SetNillableTenantId(i *int64) *SysDictTypeUpdateOne {
	if i != nil {
		sdtuo.SetTenantId(*i)
	}
	return sdtuo
}

// AddTenantId adds i to the "tenantId" field.
func (sdtuo *SysDictTypeUpdateOne) AddTenantId(i int64) *SysDictTypeUpdateOne {
	sdtuo.mutation.AddTenantId(i)
	return sdtuo
}

// AddDataListIDs adds the "dataList" edge to the SysDictData entity by IDs.
func (sdtuo *SysDictTypeUpdateOne) AddDataListIDs(ids ...int64) *SysDictTypeUpdateOne {
	sdtuo.mutation.AddDataListIDs(ids...)
	return sdtuo
}

// AddDataList adds the "dataList" edges to the SysDictData entity.
func (sdtuo *SysDictTypeUpdateOne) AddDataList(s ...*SysDictData) *SysDictTypeUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdtuo.AddDataListIDs(ids...)
}

// Mutation returns the SysDictTypeMutation object of the builder.
func (sdtuo *SysDictTypeUpdateOne) Mutation() *SysDictTypeMutation {
	return sdtuo.mutation
}

// ClearDataList clears all "dataList" edges to the SysDictData entity.
func (sdtuo *SysDictTypeUpdateOne) ClearDataList() *SysDictTypeUpdateOne {
	sdtuo.mutation.ClearDataList()
	return sdtuo
}

// RemoveDataListIDs removes the "dataList" edge to SysDictData entities by IDs.
func (sdtuo *SysDictTypeUpdateOne) RemoveDataListIDs(ids ...int64) *SysDictTypeUpdateOne {
	sdtuo.mutation.RemoveDataListIDs(ids...)
	return sdtuo
}

// RemoveDataList removes "dataList" edges to SysDictData entities.
func (sdtuo *SysDictTypeUpdateOne) RemoveDataList(s ...*SysDictData) *SysDictTypeUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdtuo.RemoveDataListIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sdtuo *SysDictTypeUpdateOne) Select(field string, fields ...string) *SysDictTypeUpdateOne {
	sdtuo.fields = append([]string{field}, fields...)
	return sdtuo
}

// Save executes the query and returns the updated SysDictType entity.
func (sdtuo *SysDictTypeUpdateOne) Save(ctx context.Context) (*SysDictType, error) {
	var (
		err  error
		node *SysDictType
	)
	sdtuo.defaults()
	if len(sdtuo.hooks) == 0 {
		node, err = sdtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysDictTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdtuo.mutation = mutation
			node, err = sdtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sdtuo.hooks) - 1; i >= 0; i-- {
			if sdtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sdtuo *SysDictTypeUpdateOne) SaveX(ctx context.Context) *SysDictType {
	node, err := sdtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sdtuo *SysDictTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := sdtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdtuo *SysDictTypeUpdateOne) ExecX(ctx context.Context) {
	if err := sdtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdtuo *SysDictTypeUpdateOne) defaults() {
	if _, ok := sdtuo.mutation.UpdatedAt(); !ok {
		v := sysdicttype.UpdateDefaultUpdatedAt()
		sdtuo.mutation.SetUpdatedAt(v)
	}
}

func (sdtuo *SysDictTypeUpdateOne) sqlSave(ctx context.Context) (_node *SysDictType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdicttype.Table,
			Columns: sysdicttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysdicttype.FieldID,
			},
		},
	}
	id, ok := sdtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SysDictType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sdtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdicttype.FieldID)
		for _, f := range fields {
			if !sysdicttype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysdicttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sdtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdtuo.mutation.DictName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdicttype.FieldDictName,
		})
	}
	if sdtuo.mutation.DictNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdicttype.FieldDictName,
		})
	}
	if value, ok := sdtuo.mutation.TypeCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdicttype.FieldTypeCode,
		})
	}
	if value, ok := sdtuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdicttype.FieldStatus,
		})
	}
	if value, ok := sdtuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysdicttype.FieldStatus,
		})
	}
	if sdtuo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: sysdicttype.FieldStatus,
		})
	}
	if value, ok := sdtuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysdicttype.FieldRemark,
		})
	}
	if sdtuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysdicttype.FieldRemark,
		})
	}
	if value, ok := sdtuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysdicttype.FieldUpdatedAt,
		})
	}
	if value, ok := sdtuo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldCreateBy,
		})
	}
	if value, ok := sdtuo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldCreateBy,
		})
	}
	if value, ok := sdtuo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldUpdateBy,
		})
	}
	if value, ok := sdtuo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldUpdateBy,
		})
	}
	if value, ok := sdtuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldTenantId,
		})
	}
	if value, ok := sdtuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: sysdicttype.FieldTenantId,
		})
	}
	if sdtuo.mutation.DataListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysdicttype.DataListTable,
			Columns: []string{sysdicttype.DataListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysdictdata.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdtuo.mutation.RemovedDataListIDs(); len(nodes) > 0 && !sdtuo.mutation.DataListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysdicttype.DataListTable,
			Columns: []string{sysdicttype.DataListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysdictdata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdtuo.mutation.DataListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysdicttype.DataListTable,
			Columns: []string{sysdicttype.DataListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysdictdata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SysDictType{config: sdtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sdtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysdicttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
