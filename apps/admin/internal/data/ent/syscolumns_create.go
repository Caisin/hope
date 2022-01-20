// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/syscolumns"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysColumnsCreate is the builder for creating a SysColumns entity.
type SysColumnsCreate struct {
	config
	mutation *SysColumnsMutation
	hooks    []Hook
}

// SetColumnId sets the "columnId" field.
func (scc *SysColumnsCreate) SetColumnId(i int32) *SysColumnsCreate {
	scc.mutation.SetColumnId(i)
	return scc
}

// SetColumnName sets the "columnName" field.
func (scc *SysColumnsCreate) SetColumnName(s string) *SysColumnsCreate {
	scc.mutation.SetColumnName(s)
	return scc
}

// SetNillableColumnName sets the "columnName" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableColumnName(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetColumnName(*s)
	}
	return scc
}

// SetColumnComment sets the "columnComment" field.
func (scc *SysColumnsCreate) SetColumnComment(s string) *SysColumnsCreate {
	scc.mutation.SetColumnComment(s)
	return scc
}

// SetNillableColumnComment sets the "columnComment" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableColumnComment(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetColumnComment(*s)
	}
	return scc
}

// SetColumnType sets the "columnType" field.
func (scc *SysColumnsCreate) SetColumnType(s string) *SysColumnsCreate {
	scc.mutation.SetColumnType(s)
	return scc
}

// SetNillableColumnType sets the "columnType" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableColumnType(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetColumnType(*s)
	}
	return scc
}

// SetGoType sets the "goType" field.
func (scc *SysColumnsCreate) SetGoType(s string) *SysColumnsCreate {
	scc.mutation.SetGoType(s)
	return scc
}

// SetNillableGoType sets the "goType" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableGoType(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetGoType(*s)
	}
	return scc
}

// SetGoField sets the "goField" field.
func (scc *SysColumnsCreate) SetGoField(s string) *SysColumnsCreate {
	scc.mutation.SetGoField(s)
	return scc
}

// SetNillableGoField sets the "goField" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableGoField(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetGoField(*s)
	}
	return scc
}

// SetJsonField sets the "jsonField" field.
func (scc *SysColumnsCreate) SetJsonField(s string) *SysColumnsCreate {
	scc.mutation.SetJsonField(s)
	return scc
}

// SetNillableJsonField sets the "jsonField" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableJsonField(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetJsonField(*s)
	}
	return scc
}

// SetIsPk sets the "isPk" field.
func (scc *SysColumnsCreate) SetIsPk(s string) *SysColumnsCreate {
	scc.mutation.SetIsPk(s)
	return scc
}

// SetNillableIsPk sets the "isPk" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableIsPk(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetIsPk(*s)
	}
	return scc
}

// SetIsIncrement sets the "isIncrement" field.
func (scc *SysColumnsCreate) SetIsIncrement(s string) *SysColumnsCreate {
	scc.mutation.SetIsIncrement(s)
	return scc
}

// SetNillableIsIncrement sets the "isIncrement" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableIsIncrement(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetIsIncrement(*s)
	}
	return scc
}

// SetIsRequired sets the "isRequired" field.
func (scc *SysColumnsCreate) SetIsRequired(s string) *SysColumnsCreate {
	scc.mutation.SetIsRequired(s)
	return scc
}

// SetNillableIsRequired sets the "isRequired" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableIsRequired(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetIsRequired(*s)
	}
	return scc
}

// SetIsInsert sets the "isInsert" field.
func (scc *SysColumnsCreate) SetIsInsert(s string) *SysColumnsCreate {
	scc.mutation.SetIsInsert(s)
	return scc
}

// SetNillableIsInsert sets the "isInsert" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableIsInsert(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetIsInsert(*s)
	}
	return scc
}

// SetIsEdit sets the "isEdit" field.
func (scc *SysColumnsCreate) SetIsEdit(s string) *SysColumnsCreate {
	scc.mutation.SetIsEdit(s)
	return scc
}

// SetNillableIsEdit sets the "isEdit" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableIsEdit(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetIsEdit(*s)
	}
	return scc
}

// SetIsList sets the "isList" field.
func (scc *SysColumnsCreate) SetIsList(s string) *SysColumnsCreate {
	scc.mutation.SetIsList(s)
	return scc
}

// SetNillableIsList sets the "isList" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableIsList(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetIsList(*s)
	}
	return scc
}

// SetIsQuery sets the "isQuery" field.
func (scc *SysColumnsCreate) SetIsQuery(s string) *SysColumnsCreate {
	scc.mutation.SetIsQuery(s)
	return scc
}

// SetNillableIsQuery sets the "isQuery" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableIsQuery(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetIsQuery(*s)
	}
	return scc
}

// SetQueryType sets the "queryType" field.
func (scc *SysColumnsCreate) SetQueryType(s string) *SysColumnsCreate {
	scc.mutation.SetQueryType(s)
	return scc
}

// SetNillableQueryType sets the "queryType" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableQueryType(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetQueryType(*s)
	}
	return scc
}

// SetHtmlType sets the "htmlType" field.
func (scc *SysColumnsCreate) SetHtmlType(s string) *SysColumnsCreate {
	scc.mutation.SetHtmlType(s)
	return scc
}

// SetNillableHtmlType sets the "htmlType" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableHtmlType(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetHtmlType(*s)
	}
	return scc
}

// SetDictType sets the "dictType" field.
func (scc *SysColumnsCreate) SetDictType(s string) *SysColumnsCreate {
	scc.mutation.SetDictType(s)
	return scc
}

// SetNillableDictType sets the "dictType" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableDictType(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetDictType(*s)
	}
	return scc
}

// SetSort sets the "sort" field.
func (scc *SysColumnsCreate) SetSort(i int32) *SysColumnsCreate {
	scc.mutation.SetSort(i)
	return scc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableSort(i *int32) *SysColumnsCreate {
	if i != nil {
		scc.SetSort(*i)
	}
	return scc
}

// SetList sets the "list" field.
func (scc *SysColumnsCreate) SetList(s string) *SysColumnsCreate {
	scc.mutation.SetList(s)
	return scc
}

// SetNillableList sets the "list" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableList(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetList(*s)
	}
	return scc
}

// SetPk sets the "pk" field.
func (scc *SysColumnsCreate) SetPk(b bool) *SysColumnsCreate {
	scc.mutation.SetPk(b)
	return scc
}

// SetNillablePk sets the "pk" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillablePk(b *bool) *SysColumnsCreate {
	if b != nil {
		scc.SetPk(*b)
	}
	return scc
}

// SetRequired sets the "required" field.
func (scc *SysColumnsCreate) SetRequired(b bool) *SysColumnsCreate {
	scc.mutation.SetRequired(b)
	return scc
}

// SetNillableRequired sets the "required" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableRequired(b *bool) *SysColumnsCreate {
	if b != nil {
		scc.SetRequired(*b)
	}
	return scc
}

// SetSuperColumn sets the "superColumn" field.
func (scc *SysColumnsCreate) SetSuperColumn(b bool) *SysColumnsCreate {
	scc.mutation.SetSuperColumn(b)
	return scc
}

// SetNillableSuperColumn sets the "superColumn" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableSuperColumn(b *bool) *SysColumnsCreate {
	if b != nil {
		scc.SetSuperColumn(*b)
	}
	return scc
}

// SetUsableColumn sets the "usableColumn" field.
func (scc *SysColumnsCreate) SetUsableColumn(b bool) *SysColumnsCreate {
	scc.mutation.SetUsableColumn(b)
	return scc
}

// SetNillableUsableColumn sets the "usableColumn" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableUsableColumn(b *bool) *SysColumnsCreate {
	if b != nil {
		scc.SetUsableColumn(*b)
	}
	return scc
}

// SetIncrement sets the "increment" field.
func (scc *SysColumnsCreate) SetIncrement(b bool) *SysColumnsCreate {
	scc.mutation.SetIncrement(b)
	return scc
}

// SetNillableIncrement sets the "increment" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableIncrement(b *bool) *SysColumnsCreate {
	if b != nil {
		scc.SetIncrement(*b)
	}
	return scc
}

// SetInsert sets the "insert" field.
func (scc *SysColumnsCreate) SetInsert(b bool) *SysColumnsCreate {
	scc.mutation.SetInsert(b)
	return scc
}

// SetNillableInsert sets the "insert" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableInsert(b *bool) *SysColumnsCreate {
	if b != nil {
		scc.SetInsert(*b)
	}
	return scc
}

// SetEdit sets the "edit" field.
func (scc *SysColumnsCreate) SetEdit(b bool) *SysColumnsCreate {
	scc.mutation.SetEdit(b)
	return scc
}

// SetNillableEdit sets the "edit" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableEdit(b *bool) *SysColumnsCreate {
	if b != nil {
		scc.SetEdit(*b)
	}
	return scc
}

// SetQuery sets the "query" field.
func (scc *SysColumnsCreate) SetQuery(b bool) *SysColumnsCreate {
	scc.mutation.SetQuery(b)
	return scc
}

// SetNillableQuery sets the "query" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableQuery(b *bool) *SysColumnsCreate {
	if b != nil {
		scc.SetQuery(*b)
	}
	return scc
}

// SetRemark sets the "remark" field.
func (scc *SysColumnsCreate) SetRemark(s string) *SysColumnsCreate {
	scc.mutation.SetRemark(s)
	return scc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableRemark(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetRemark(*s)
	}
	return scc
}

// SetFkLabelName sets the "fkLabelName" field.
func (scc *SysColumnsCreate) SetFkLabelName(s string) *SysColumnsCreate {
	scc.mutation.SetFkLabelName(s)
	return scc
}

// SetNillableFkLabelName sets the "fkLabelName" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableFkLabelName(s *string) *SysColumnsCreate {
	if s != nil {
		scc.SetFkLabelName(*s)
	}
	return scc
}

// SetCreatedAt sets the "createdAt" field.
func (scc *SysColumnsCreate) SetCreatedAt(t time.Time) *SysColumnsCreate {
	scc.mutation.SetCreatedAt(t)
	return scc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableCreatedAt(t *time.Time) *SysColumnsCreate {
	if t != nil {
		scc.SetCreatedAt(*t)
	}
	return scc
}

// SetUpdatedAt sets the "updatedAt" field.
func (scc *SysColumnsCreate) SetUpdatedAt(t time.Time) *SysColumnsCreate {
	scc.mutation.SetUpdatedAt(t)
	return scc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableUpdatedAt(t *time.Time) *SysColumnsCreate {
	if t != nil {
		scc.SetUpdatedAt(*t)
	}
	return scc
}

// SetCreateBy sets the "createBy" field.
func (scc *SysColumnsCreate) SetCreateBy(i int64) *SysColumnsCreate {
	scc.mutation.SetCreateBy(i)
	return scc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableCreateBy(i *int64) *SysColumnsCreate {
	if i != nil {
		scc.SetCreateBy(*i)
	}
	return scc
}

// SetUpdateBy sets the "updateBy" field.
func (scc *SysColumnsCreate) SetUpdateBy(i int64) *SysColumnsCreate {
	scc.mutation.SetUpdateBy(i)
	return scc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableUpdateBy(i *int64) *SysColumnsCreate {
	if i != nil {
		scc.SetUpdateBy(*i)
	}
	return scc
}

// SetTenantId sets the "tenantId" field.
func (scc *SysColumnsCreate) SetTenantId(i int64) *SysColumnsCreate {
	scc.mutation.SetTenantId(i)
	return scc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (scc *SysColumnsCreate) SetNillableTenantId(i *int64) *SysColumnsCreate {
	if i != nil {
		scc.SetTenantId(*i)
	}
	return scc
}

// Mutation returns the SysColumnsMutation object of the builder.
func (scc *SysColumnsCreate) Mutation() *SysColumnsMutation {
	return scc.mutation
}

// Save creates the SysColumns in the database.
func (scc *SysColumnsCreate) Save(ctx context.Context) (*SysColumns, error) {
	var (
		err  error
		node *SysColumns
	)
	scc.defaults()
	if len(scc.hooks) == 0 {
		if err = scc.check(); err != nil {
			return nil, err
		}
		node, err = scc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysColumnsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = scc.check(); err != nil {
				return nil, err
			}
			scc.mutation = mutation
			if node, err = scc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(scc.hooks) - 1; i >= 0; i-- {
			if scc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = scc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, scc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (scc *SysColumnsCreate) SaveX(ctx context.Context) *SysColumns {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *SysColumnsCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *SysColumnsCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *SysColumnsCreate) defaults() {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		v := syscolumns.DefaultCreatedAt()
		scc.mutation.SetCreatedAt(v)
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		v := syscolumns.DefaultUpdatedAt()
		scc.mutation.SetUpdatedAt(v)
	}
	if _, ok := scc.mutation.CreateBy(); !ok {
		v := syscolumns.DefaultCreateBy
		scc.mutation.SetCreateBy(v)
	}
	if _, ok := scc.mutation.UpdateBy(); !ok {
		v := syscolumns.DefaultUpdateBy
		scc.mutation.SetUpdateBy(v)
	}
	if _, ok := scc.mutation.TenantId(); !ok {
		v := syscolumns.DefaultTenantId
		scc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *SysColumnsCreate) check() error {
	if _, ok := scc.mutation.ColumnId(); !ok {
		return &ValidationError{Name: "columnId", err: errors.New(`ent: missing required field "columnId"`)}
	}
	if _, ok := scc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := scc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := scc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := scc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (scc *SysColumnsCreate) sqlSave(ctx context.Context) (*SysColumns, error) {
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (scc *SysColumnsCreate) createSpec() (*SysColumns, *sqlgraph.CreateSpec) {
	var (
		_node = &SysColumns{config: scc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: syscolumns.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: syscolumns.FieldID,
			},
		}
	)
	if value, ok := scc.mutation.ColumnId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syscolumns.FieldColumnId,
		})
		_node.ColumnId = value
	}
	if value, ok := scc.mutation.ColumnName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldColumnName,
		})
		_node.ColumnName = value
	}
	if value, ok := scc.mutation.ColumnComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldColumnComment,
		})
		_node.ColumnComment = value
	}
	if value, ok := scc.mutation.ColumnType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldColumnType,
		})
		_node.ColumnType = value
	}
	if value, ok := scc.mutation.GoType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldGoType,
		})
		_node.GoType = value
	}
	if value, ok := scc.mutation.GoField(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldGoField,
		})
		_node.GoField = value
	}
	if value, ok := scc.mutation.JsonField(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldJsonField,
		})
		_node.JsonField = value
	}
	if value, ok := scc.mutation.IsPk(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldIsPk,
		})
		_node.IsPk = value
	}
	if value, ok := scc.mutation.IsIncrement(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldIsIncrement,
		})
		_node.IsIncrement = value
	}
	if value, ok := scc.mutation.IsRequired(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldIsRequired,
		})
		_node.IsRequired = value
	}
	if value, ok := scc.mutation.IsInsert(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldIsInsert,
		})
		_node.IsInsert = value
	}
	if value, ok := scc.mutation.IsEdit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldIsEdit,
		})
		_node.IsEdit = value
	}
	if value, ok := scc.mutation.IsList(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldIsList,
		})
		_node.IsList = value
	}
	if value, ok := scc.mutation.IsQuery(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldIsQuery,
		})
		_node.IsQuery = value
	}
	if value, ok := scc.mutation.QueryType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldQueryType,
		})
		_node.QueryType = value
	}
	if value, ok := scc.mutation.HtmlType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldHtmlType,
		})
		_node.HtmlType = value
	}
	if value, ok := scc.mutation.DictType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldDictType,
		})
		_node.DictType = value
	}
	if value, ok := scc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syscolumns.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := scc.mutation.List(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldList,
		})
		_node.List = value
	}
	if value, ok := scc.mutation.Pk(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: syscolumns.FieldPk,
		})
		_node.Pk = value
	}
	if value, ok := scc.mutation.Required(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: syscolumns.FieldRequired,
		})
		_node.Required = value
	}
	if value, ok := scc.mutation.SuperColumn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: syscolumns.FieldSuperColumn,
		})
		_node.SuperColumn = value
	}
	if value, ok := scc.mutation.UsableColumn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: syscolumns.FieldUsableColumn,
		})
		_node.UsableColumn = value
	}
	if value, ok := scc.mutation.Increment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: syscolumns.FieldIncrement,
		})
		_node.Increment = value
	}
	if value, ok := scc.mutation.Insert(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: syscolumns.FieldInsert,
		})
		_node.Insert = value
	}
	if value, ok := scc.mutation.Edit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: syscolumns.FieldEdit,
		})
		_node.Edit = value
	}
	if value, ok := scc.mutation.Query(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: syscolumns.FieldQuery,
		})
		_node.Query = value
	}
	if value, ok := scc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := scc.mutation.FkLabelName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syscolumns.FieldFkLabelName,
		})
		_node.FkLabelName = value
	}
	if value, ok := scc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: syscolumns.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := scc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: syscolumns.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := scc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syscolumns.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := scc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syscolumns.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := scc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syscolumns.FieldTenantId,
		})
		_node.TenantId = value
	}
	return _node, _spec
}

// SysColumnsCreateBulk is the builder for creating many SysColumns entities in bulk.
type SysColumnsCreateBulk struct {
	config
	builders []*SysColumnsCreate
}

// Save creates the SysColumns entities in the database.
func (sccb *SysColumnsCreateBulk) Save(ctx context.Context) ([]*SysColumns, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*SysColumns, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysColumnsMutation)
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
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *SysColumnsCreateBulk) SaveX(ctx context.Context) []*SysColumns {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *SysColumnsCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *SysColumnsCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}
