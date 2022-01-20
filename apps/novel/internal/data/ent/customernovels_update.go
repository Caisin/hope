// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/novel/internal/data/ent/customernovels"
	"hope/apps/novel/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerNovelsUpdate is the builder for updating CustomerNovels entities.
type CustomerNovelsUpdate struct {
	config
	hooks    []Hook
	mutation *CustomerNovelsMutation
}

// Where appends a list predicates to the CustomerNovelsUpdate builder.
func (cnu *CustomerNovelsUpdate) Where(ps ...predicate.CustomerNovels) *CustomerNovelsUpdate {
	cnu.mutation.Where(ps...)
	return cnu
}

// SetNovelId sets the "novelId" field.
func (cnu *CustomerNovelsUpdate) SetNovelId(i int64) *CustomerNovelsUpdate {
	cnu.mutation.ResetNovelId()
	cnu.mutation.SetNovelId(i)
	return cnu
}

// SetNillableNovelId sets the "novelId" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableNovelId(i *int64) *CustomerNovelsUpdate {
	if i != nil {
		cnu.SetNovelId(*i)
	}
	return cnu
}

// AddNovelId adds i to the "novelId" field.
func (cnu *CustomerNovelsUpdate) AddNovelId(i int64) *CustomerNovelsUpdate {
	cnu.mutation.AddNovelId(i)
	return cnu
}

// ClearNovelId clears the value of the "novelId" field.
func (cnu *CustomerNovelsUpdate) ClearNovelId() *CustomerNovelsUpdate {
	cnu.mutation.ClearNovelId()
	return cnu
}

// SetTypeId sets the "typeId" field.
func (cnu *CustomerNovelsUpdate) SetTypeId(i int32) *CustomerNovelsUpdate {
	cnu.mutation.ResetTypeId()
	cnu.mutation.SetTypeId(i)
	return cnu
}

// SetNillableTypeId sets the "typeId" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableTypeId(i *int32) *CustomerNovelsUpdate {
	if i != nil {
		cnu.SetTypeId(*i)
	}
	return cnu
}

// AddTypeId adds i to the "typeId" field.
func (cnu *CustomerNovelsUpdate) AddTypeId(i int32) *CustomerNovelsUpdate {
	cnu.mutation.AddTypeId(i)
	return cnu
}

// ClearTypeId clears the value of the "typeId" field.
func (cnu *CustomerNovelsUpdate) ClearTypeId() *CustomerNovelsUpdate {
	cnu.mutation.ClearTypeId()
	return cnu
}

// SetTypeCode sets the "typeCode" field.
func (cnu *CustomerNovelsUpdate) SetTypeCode(s string) *CustomerNovelsUpdate {
	cnu.mutation.SetTypeCode(s)
	return cnu
}

// SetNillableTypeCode sets the "typeCode" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableTypeCode(s *string) *CustomerNovelsUpdate {
	if s != nil {
		cnu.SetTypeCode(*s)
	}
	return cnu
}

// ClearTypeCode clears the value of the "typeCode" field.
func (cnu *CustomerNovelsUpdate) ClearTypeCode() *CustomerNovelsUpdate {
	cnu.mutation.ClearTypeCode()
	return cnu
}

// SetGroupCode sets the "groupCode" field.
func (cnu *CustomerNovelsUpdate) SetGroupCode(s string) *CustomerNovelsUpdate {
	cnu.mutation.SetGroupCode(s)
	return cnu
}

// SetNillableGroupCode sets the "groupCode" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableGroupCode(s *string) *CustomerNovelsUpdate {
	if s != nil {
		cnu.SetGroupCode(*s)
	}
	return cnu
}

// ClearGroupCode clears the value of the "groupCode" field.
func (cnu *CustomerNovelsUpdate) ClearGroupCode() *CustomerNovelsUpdate {
	cnu.mutation.ClearGroupCode()
	return cnu
}

// SetFieldName sets the "fieldName" field.
func (cnu *CustomerNovelsUpdate) SetFieldName(s string) *CustomerNovelsUpdate {
	cnu.mutation.SetFieldName(s)
	return cnu
}

// SetNillableFieldName sets the "fieldName" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableFieldName(s *string) *CustomerNovelsUpdate {
	if s != nil {
		cnu.SetFieldName(*s)
	}
	return cnu
}

// ClearFieldName clears the value of the "fieldName" field.
func (cnu *CustomerNovelsUpdate) ClearFieldName() *CustomerNovelsUpdate {
	cnu.mutation.ClearFieldName()
	return cnu
}

// SetCover sets the "cover" field.
func (cnu *CustomerNovelsUpdate) SetCover(s string) *CustomerNovelsUpdate {
	cnu.mutation.SetCover(s)
	return cnu
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableCover(s *string) *CustomerNovelsUpdate {
	if s != nil {
		cnu.SetCover(*s)
	}
	return cnu
}

// ClearCover clears the value of the "cover" field.
func (cnu *CustomerNovelsUpdate) ClearCover() *CustomerNovelsUpdate {
	cnu.mutation.ClearCover()
	return cnu
}

// SetOrderNum sets the "orderNum" field.
func (cnu *CustomerNovelsUpdate) SetOrderNum(i int32) *CustomerNovelsUpdate {
	cnu.mutation.ResetOrderNum()
	cnu.mutation.SetOrderNum(i)
	return cnu
}

// SetNillableOrderNum sets the "orderNum" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableOrderNum(i *int32) *CustomerNovelsUpdate {
	if i != nil {
		cnu.SetOrderNum(*i)
	}
	return cnu
}

// AddOrderNum adds i to the "orderNum" field.
func (cnu *CustomerNovelsUpdate) AddOrderNum(i int32) *CustomerNovelsUpdate {
	cnu.mutation.AddOrderNum(i)
	return cnu
}

// ClearOrderNum clears the value of the "orderNum" field.
func (cnu *CustomerNovelsUpdate) ClearOrderNum() *CustomerNovelsUpdate {
	cnu.mutation.ClearOrderNum()
	return cnu
}

// SetRemark sets the "remark" field.
func (cnu *CustomerNovelsUpdate) SetRemark(s string) *CustomerNovelsUpdate {
	cnu.mutation.SetRemark(s)
	return cnu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableRemark(s *string) *CustomerNovelsUpdate {
	if s != nil {
		cnu.SetRemark(*s)
	}
	return cnu
}

// ClearRemark clears the value of the "remark" field.
func (cnu *CustomerNovelsUpdate) ClearRemark() *CustomerNovelsUpdate {
	cnu.mutation.ClearRemark()
	return cnu
}

// SetEffectTime sets the "effectTime" field.
func (cnu *CustomerNovelsUpdate) SetEffectTime(t time.Time) *CustomerNovelsUpdate {
	cnu.mutation.SetEffectTime(t)
	return cnu
}

// SetNillableEffectTime sets the "effectTime" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableEffectTime(t *time.Time) *CustomerNovelsUpdate {
	if t != nil {
		cnu.SetEffectTime(*t)
	}
	return cnu
}

// SetExpiredTime sets the "expiredTime" field.
func (cnu *CustomerNovelsUpdate) SetExpiredTime(t time.Time) *CustomerNovelsUpdate {
	cnu.mutation.SetExpiredTime(t)
	return cnu
}

// SetNillableExpiredTime sets the "expiredTime" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableExpiredTime(t *time.Time) *CustomerNovelsUpdate {
	if t != nil {
		cnu.SetExpiredTime(*t)
	}
	return cnu
}

// SetUpdatedAt sets the "updatedAt" field.
func (cnu *CustomerNovelsUpdate) SetUpdatedAt(t time.Time) *CustomerNovelsUpdate {
	cnu.mutation.SetUpdatedAt(t)
	return cnu
}

// SetCreateBy sets the "createBy" field.
func (cnu *CustomerNovelsUpdate) SetCreateBy(i int64) *CustomerNovelsUpdate {
	cnu.mutation.ResetCreateBy()
	cnu.mutation.SetCreateBy(i)
	return cnu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableCreateBy(i *int64) *CustomerNovelsUpdate {
	if i != nil {
		cnu.SetCreateBy(*i)
	}
	return cnu
}

// AddCreateBy adds i to the "createBy" field.
func (cnu *CustomerNovelsUpdate) AddCreateBy(i int64) *CustomerNovelsUpdate {
	cnu.mutation.AddCreateBy(i)
	return cnu
}

// SetUpdateBy sets the "updateBy" field.
func (cnu *CustomerNovelsUpdate) SetUpdateBy(i int64) *CustomerNovelsUpdate {
	cnu.mutation.ResetUpdateBy()
	cnu.mutation.SetUpdateBy(i)
	return cnu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableUpdateBy(i *int64) *CustomerNovelsUpdate {
	if i != nil {
		cnu.SetUpdateBy(*i)
	}
	return cnu
}

// AddUpdateBy adds i to the "updateBy" field.
func (cnu *CustomerNovelsUpdate) AddUpdateBy(i int64) *CustomerNovelsUpdate {
	cnu.mutation.AddUpdateBy(i)
	return cnu
}

// SetTenantId sets the "tenantId" field.
func (cnu *CustomerNovelsUpdate) SetTenantId(i int64) *CustomerNovelsUpdate {
	cnu.mutation.ResetTenantId()
	cnu.mutation.SetTenantId(i)
	return cnu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (cnu *CustomerNovelsUpdate) SetNillableTenantId(i *int64) *CustomerNovelsUpdate {
	if i != nil {
		cnu.SetTenantId(*i)
	}
	return cnu
}

// AddTenantId adds i to the "tenantId" field.
func (cnu *CustomerNovelsUpdate) AddTenantId(i int64) *CustomerNovelsUpdate {
	cnu.mutation.AddTenantId(i)
	return cnu
}

// Mutation returns the CustomerNovelsMutation object of the builder.
func (cnu *CustomerNovelsUpdate) Mutation() *CustomerNovelsMutation {
	return cnu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cnu *CustomerNovelsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cnu.defaults()
	if len(cnu.hooks) == 0 {
		affected, err = cnu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerNovelsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cnu.mutation = mutation
			affected, err = cnu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cnu.hooks) - 1; i >= 0; i-- {
			if cnu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cnu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cnu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cnu *CustomerNovelsUpdate) SaveX(ctx context.Context) int {
	affected, err := cnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cnu *CustomerNovelsUpdate) Exec(ctx context.Context) error {
	_, err := cnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnu *CustomerNovelsUpdate) ExecX(ctx context.Context) {
	if err := cnu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cnu *CustomerNovelsUpdate) defaults() {
	if _, ok := cnu.mutation.UpdatedAt(); !ok {
		v := customernovels.UpdateDefaultUpdatedAt()
		cnu.mutation.SetUpdatedAt(v)
	}
}

func (cnu *CustomerNovelsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customernovels.Table,
			Columns: customernovels.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: customernovels.FieldID,
			},
		},
	}
	if ps := cnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cnu.mutation.NovelId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldNovelId,
		})
	}
	if value, ok := cnu.mutation.AddedNovelId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldNovelId,
		})
	}
	if cnu.mutation.NovelIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: customernovels.FieldNovelId,
		})
	}
	if value, ok := cnu.mutation.TypeId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: customernovels.FieldTypeId,
		})
	}
	if value, ok := cnu.mutation.AddedTypeId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: customernovels.FieldTypeId,
		})
	}
	if cnu.mutation.TypeIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: customernovels.FieldTypeId,
		})
	}
	if value, ok := cnu.mutation.TypeCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldTypeCode,
		})
	}
	if cnu.mutation.TypeCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customernovels.FieldTypeCode,
		})
	}
	if value, ok := cnu.mutation.GroupCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldGroupCode,
		})
	}
	if cnu.mutation.GroupCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customernovels.FieldGroupCode,
		})
	}
	if value, ok := cnu.mutation.FieldName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldFieldName,
		})
	}
	if cnu.mutation.FieldNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customernovels.FieldFieldName,
		})
	}
	if value, ok := cnu.mutation.Cover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldCover,
		})
	}
	if cnu.mutation.CoverCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customernovels.FieldCover,
		})
	}
	if value, ok := cnu.mutation.OrderNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: customernovels.FieldOrderNum,
		})
	}
	if value, ok := cnu.mutation.AddedOrderNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: customernovels.FieldOrderNum,
		})
	}
	if cnu.mutation.OrderNumCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: customernovels.FieldOrderNum,
		})
	}
	if value, ok := cnu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldRemark,
		})
	}
	if cnu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customernovels.FieldRemark,
		})
	}
	if value, ok := cnu.mutation.EffectTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customernovels.FieldEffectTime,
		})
	}
	if value, ok := cnu.mutation.ExpiredTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customernovels.FieldExpiredTime,
		})
	}
	if value, ok := cnu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customernovels.FieldUpdatedAt,
		})
	}
	if value, ok := cnu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldCreateBy,
		})
	}
	if value, ok := cnu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldCreateBy,
		})
	}
	if value, ok := cnu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldUpdateBy,
		})
	}
	if value, ok := cnu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldUpdateBy,
		})
	}
	if value, ok := cnu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldTenantId,
		})
	}
	if value, ok := cnu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldTenantId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customernovels.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CustomerNovelsUpdateOne is the builder for updating a single CustomerNovels entity.
type CustomerNovelsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomerNovelsMutation
}

// SetNovelId sets the "novelId" field.
func (cnuo *CustomerNovelsUpdateOne) SetNovelId(i int64) *CustomerNovelsUpdateOne {
	cnuo.mutation.ResetNovelId()
	cnuo.mutation.SetNovelId(i)
	return cnuo
}

// SetNillableNovelId sets the "novelId" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableNovelId(i *int64) *CustomerNovelsUpdateOne {
	if i != nil {
		cnuo.SetNovelId(*i)
	}
	return cnuo
}

// AddNovelId adds i to the "novelId" field.
func (cnuo *CustomerNovelsUpdateOne) AddNovelId(i int64) *CustomerNovelsUpdateOne {
	cnuo.mutation.AddNovelId(i)
	return cnuo
}

// ClearNovelId clears the value of the "novelId" field.
func (cnuo *CustomerNovelsUpdateOne) ClearNovelId() *CustomerNovelsUpdateOne {
	cnuo.mutation.ClearNovelId()
	return cnuo
}

// SetTypeId sets the "typeId" field.
func (cnuo *CustomerNovelsUpdateOne) SetTypeId(i int32) *CustomerNovelsUpdateOne {
	cnuo.mutation.ResetTypeId()
	cnuo.mutation.SetTypeId(i)
	return cnuo
}

// SetNillableTypeId sets the "typeId" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableTypeId(i *int32) *CustomerNovelsUpdateOne {
	if i != nil {
		cnuo.SetTypeId(*i)
	}
	return cnuo
}

// AddTypeId adds i to the "typeId" field.
func (cnuo *CustomerNovelsUpdateOne) AddTypeId(i int32) *CustomerNovelsUpdateOne {
	cnuo.mutation.AddTypeId(i)
	return cnuo
}

// ClearTypeId clears the value of the "typeId" field.
func (cnuo *CustomerNovelsUpdateOne) ClearTypeId() *CustomerNovelsUpdateOne {
	cnuo.mutation.ClearTypeId()
	return cnuo
}

// SetTypeCode sets the "typeCode" field.
func (cnuo *CustomerNovelsUpdateOne) SetTypeCode(s string) *CustomerNovelsUpdateOne {
	cnuo.mutation.SetTypeCode(s)
	return cnuo
}

// SetNillableTypeCode sets the "typeCode" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableTypeCode(s *string) *CustomerNovelsUpdateOne {
	if s != nil {
		cnuo.SetTypeCode(*s)
	}
	return cnuo
}

// ClearTypeCode clears the value of the "typeCode" field.
func (cnuo *CustomerNovelsUpdateOne) ClearTypeCode() *CustomerNovelsUpdateOne {
	cnuo.mutation.ClearTypeCode()
	return cnuo
}

// SetGroupCode sets the "groupCode" field.
func (cnuo *CustomerNovelsUpdateOne) SetGroupCode(s string) *CustomerNovelsUpdateOne {
	cnuo.mutation.SetGroupCode(s)
	return cnuo
}

// SetNillableGroupCode sets the "groupCode" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableGroupCode(s *string) *CustomerNovelsUpdateOne {
	if s != nil {
		cnuo.SetGroupCode(*s)
	}
	return cnuo
}

// ClearGroupCode clears the value of the "groupCode" field.
func (cnuo *CustomerNovelsUpdateOne) ClearGroupCode() *CustomerNovelsUpdateOne {
	cnuo.mutation.ClearGroupCode()
	return cnuo
}

// SetFieldName sets the "fieldName" field.
func (cnuo *CustomerNovelsUpdateOne) SetFieldName(s string) *CustomerNovelsUpdateOne {
	cnuo.mutation.SetFieldName(s)
	return cnuo
}

// SetNillableFieldName sets the "fieldName" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableFieldName(s *string) *CustomerNovelsUpdateOne {
	if s != nil {
		cnuo.SetFieldName(*s)
	}
	return cnuo
}

// ClearFieldName clears the value of the "fieldName" field.
func (cnuo *CustomerNovelsUpdateOne) ClearFieldName() *CustomerNovelsUpdateOne {
	cnuo.mutation.ClearFieldName()
	return cnuo
}

// SetCover sets the "cover" field.
func (cnuo *CustomerNovelsUpdateOne) SetCover(s string) *CustomerNovelsUpdateOne {
	cnuo.mutation.SetCover(s)
	return cnuo
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableCover(s *string) *CustomerNovelsUpdateOne {
	if s != nil {
		cnuo.SetCover(*s)
	}
	return cnuo
}

// ClearCover clears the value of the "cover" field.
func (cnuo *CustomerNovelsUpdateOne) ClearCover() *CustomerNovelsUpdateOne {
	cnuo.mutation.ClearCover()
	return cnuo
}

// SetOrderNum sets the "orderNum" field.
func (cnuo *CustomerNovelsUpdateOne) SetOrderNum(i int32) *CustomerNovelsUpdateOne {
	cnuo.mutation.ResetOrderNum()
	cnuo.mutation.SetOrderNum(i)
	return cnuo
}

// SetNillableOrderNum sets the "orderNum" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableOrderNum(i *int32) *CustomerNovelsUpdateOne {
	if i != nil {
		cnuo.SetOrderNum(*i)
	}
	return cnuo
}

// AddOrderNum adds i to the "orderNum" field.
func (cnuo *CustomerNovelsUpdateOne) AddOrderNum(i int32) *CustomerNovelsUpdateOne {
	cnuo.mutation.AddOrderNum(i)
	return cnuo
}

// ClearOrderNum clears the value of the "orderNum" field.
func (cnuo *CustomerNovelsUpdateOne) ClearOrderNum() *CustomerNovelsUpdateOne {
	cnuo.mutation.ClearOrderNum()
	return cnuo
}

// SetRemark sets the "remark" field.
func (cnuo *CustomerNovelsUpdateOne) SetRemark(s string) *CustomerNovelsUpdateOne {
	cnuo.mutation.SetRemark(s)
	return cnuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableRemark(s *string) *CustomerNovelsUpdateOne {
	if s != nil {
		cnuo.SetRemark(*s)
	}
	return cnuo
}

// ClearRemark clears the value of the "remark" field.
func (cnuo *CustomerNovelsUpdateOne) ClearRemark() *CustomerNovelsUpdateOne {
	cnuo.mutation.ClearRemark()
	return cnuo
}

// SetEffectTime sets the "effectTime" field.
func (cnuo *CustomerNovelsUpdateOne) SetEffectTime(t time.Time) *CustomerNovelsUpdateOne {
	cnuo.mutation.SetEffectTime(t)
	return cnuo
}

// SetNillableEffectTime sets the "effectTime" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableEffectTime(t *time.Time) *CustomerNovelsUpdateOne {
	if t != nil {
		cnuo.SetEffectTime(*t)
	}
	return cnuo
}

// SetExpiredTime sets the "expiredTime" field.
func (cnuo *CustomerNovelsUpdateOne) SetExpiredTime(t time.Time) *CustomerNovelsUpdateOne {
	cnuo.mutation.SetExpiredTime(t)
	return cnuo
}

// SetNillableExpiredTime sets the "expiredTime" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableExpiredTime(t *time.Time) *CustomerNovelsUpdateOne {
	if t != nil {
		cnuo.SetExpiredTime(*t)
	}
	return cnuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (cnuo *CustomerNovelsUpdateOne) SetUpdatedAt(t time.Time) *CustomerNovelsUpdateOne {
	cnuo.mutation.SetUpdatedAt(t)
	return cnuo
}

// SetCreateBy sets the "createBy" field.
func (cnuo *CustomerNovelsUpdateOne) SetCreateBy(i int64) *CustomerNovelsUpdateOne {
	cnuo.mutation.ResetCreateBy()
	cnuo.mutation.SetCreateBy(i)
	return cnuo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableCreateBy(i *int64) *CustomerNovelsUpdateOne {
	if i != nil {
		cnuo.SetCreateBy(*i)
	}
	return cnuo
}

// AddCreateBy adds i to the "createBy" field.
func (cnuo *CustomerNovelsUpdateOne) AddCreateBy(i int64) *CustomerNovelsUpdateOne {
	cnuo.mutation.AddCreateBy(i)
	return cnuo
}

// SetUpdateBy sets the "updateBy" field.
func (cnuo *CustomerNovelsUpdateOne) SetUpdateBy(i int64) *CustomerNovelsUpdateOne {
	cnuo.mutation.ResetUpdateBy()
	cnuo.mutation.SetUpdateBy(i)
	return cnuo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableUpdateBy(i *int64) *CustomerNovelsUpdateOne {
	if i != nil {
		cnuo.SetUpdateBy(*i)
	}
	return cnuo
}

// AddUpdateBy adds i to the "updateBy" field.
func (cnuo *CustomerNovelsUpdateOne) AddUpdateBy(i int64) *CustomerNovelsUpdateOne {
	cnuo.mutation.AddUpdateBy(i)
	return cnuo
}

// SetTenantId sets the "tenantId" field.
func (cnuo *CustomerNovelsUpdateOne) SetTenantId(i int64) *CustomerNovelsUpdateOne {
	cnuo.mutation.ResetTenantId()
	cnuo.mutation.SetTenantId(i)
	return cnuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (cnuo *CustomerNovelsUpdateOne) SetNillableTenantId(i *int64) *CustomerNovelsUpdateOne {
	if i != nil {
		cnuo.SetTenantId(*i)
	}
	return cnuo
}

// AddTenantId adds i to the "tenantId" field.
func (cnuo *CustomerNovelsUpdateOne) AddTenantId(i int64) *CustomerNovelsUpdateOne {
	cnuo.mutation.AddTenantId(i)
	return cnuo
}

// Mutation returns the CustomerNovelsMutation object of the builder.
func (cnuo *CustomerNovelsUpdateOne) Mutation() *CustomerNovelsMutation {
	return cnuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cnuo *CustomerNovelsUpdateOne) Select(field string, fields ...string) *CustomerNovelsUpdateOne {
	cnuo.fields = append([]string{field}, fields...)
	return cnuo
}

// Save executes the query and returns the updated CustomerNovels entity.
func (cnuo *CustomerNovelsUpdateOne) Save(ctx context.Context) (*CustomerNovels, error) {
	var (
		err  error
		node *CustomerNovels
	)
	cnuo.defaults()
	if len(cnuo.hooks) == 0 {
		node, err = cnuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerNovelsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cnuo.mutation = mutation
			node, err = cnuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cnuo.hooks) - 1; i >= 0; i-- {
			if cnuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cnuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cnuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cnuo *CustomerNovelsUpdateOne) SaveX(ctx context.Context) *CustomerNovels {
	node, err := cnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cnuo *CustomerNovelsUpdateOne) Exec(ctx context.Context) error {
	_, err := cnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cnuo *CustomerNovelsUpdateOne) ExecX(ctx context.Context) {
	if err := cnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cnuo *CustomerNovelsUpdateOne) defaults() {
	if _, ok := cnuo.mutation.UpdatedAt(); !ok {
		v := customernovels.UpdateDefaultUpdatedAt()
		cnuo.mutation.SetUpdatedAt(v)
	}
}

func (cnuo *CustomerNovelsUpdateOne) sqlSave(ctx context.Context) (_node *CustomerNovels, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customernovels.Table,
			Columns: customernovels.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: customernovels.FieldID,
			},
		},
	}
	id, ok := cnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CustomerNovels.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customernovels.FieldID)
		for _, f := range fields {
			if !customernovels.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != customernovels.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cnuo.mutation.NovelId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldNovelId,
		})
	}
	if value, ok := cnuo.mutation.AddedNovelId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldNovelId,
		})
	}
	if cnuo.mutation.NovelIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: customernovels.FieldNovelId,
		})
	}
	if value, ok := cnuo.mutation.TypeId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: customernovels.FieldTypeId,
		})
	}
	if value, ok := cnuo.mutation.AddedTypeId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: customernovels.FieldTypeId,
		})
	}
	if cnuo.mutation.TypeIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: customernovels.FieldTypeId,
		})
	}
	if value, ok := cnuo.mutation.TypeCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldTypeCode,
		})
	}
	if cnuo.mutation.TypeCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customernovels.FieldTypeCode,
		})
	}
	if value, ok := cnuo.mutation.GroupCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldGroupCode,
		})
	}
	if cnuo.mutation.GroupCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customernovels.FieldGroupCode,
		})
	}
	if value, ok := cnuo.mutation.FieldName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldFieldName,
		})
	}
	if cnuo.mutation.FieldNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customernovels.FieldFieldName,
		})
	}
	if value, ok := cnuo.mutation.Cover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldCover,
		})
	}
	if cnuo.mutation.CoverCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customernovels.FieldCover,
		})
	}
	if value, ok := cnuo.mutation.OrderNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: customernovels.FieldOrderNum,
		})
	}
	if value, ok := cnuo.mutation.AddedOrderNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: customernovels.FieldOrderNum,
		})
	}
	if cnuo.mutation.OrderNumCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: customernovels.FieldOrderNum,
		})
	}
	if value, ok := cnuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customernovels.FieldRemark,
		})
	}
	if cnuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customernovels.FieldRemark,
		})
	}
	if value, ok := cnuo.mutation.EffectTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customernovels.FieldEffectTime,
		})
	}
	if value, ok := cnuo.mutation.ExpiredTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customernovels.FieldExpiredTime,
		})
	}
	if value, ok := cnuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customernovels.FieldUpdatedAt,
		})
	}
	if value, ok := cnuo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldCreateBy,
		})
	}
	if value, ok := cnuo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldCreateBy,
		})
	}
	if value, ok := cnuo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldUpdateBy,
		})
	}
	if value, ok := cnuo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldUpdateBy,
		})
	}
	if value, ok := cnuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldTenantId,
		})
	}
	if value, ok := cnuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: customernovels.FieldTenantId,
		})
	}
	_node = &CustomerNovels{config: cnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customernovels.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
