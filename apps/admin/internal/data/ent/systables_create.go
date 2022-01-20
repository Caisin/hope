// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/systables"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysTablesCreate is the builder for creating a SysTables entity.
type SysTablesCreate struct {
	config
	mutation *SysTablesMutation
	hooks    []Hook
}

// SetTableName sets the "tableName" field.
func (stc *SysTablesCreate) SetTableName(s string) *SysTablesCreate {
	stc.mutation.SetTableName(s)
	return stc
}

// SetNillableTableName sets the "tableName" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableTableName(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetTableName(*s)
	}
	return stc
}

// SetTableComment sets the "tableComment" field.
func (stc *SysTablesCreate) SetTableComment(s string) *SysTablesCreate {
	stc.mutation.SetTableComment(s)
	return stc
}

// SetNillableTableComment sets the "tableComment" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableTableComment(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetTableComment(*s)
	}
	return stc
}

// SetClassName sets the "className" field.
func (stc *SysTablesCreate) SetClassName(s string) *SysTablesCreate {
	stc.mutation.SetClassName(s)
	return stc
}

// SetNillableClassName sets the "className" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableClassName(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetClassName(*s)
	}
	return stc
}

// SetTplCategory sets the "tplCategory" field.
func (stc *SysTablesCreate) SetTplCategory(s string) *SysTablesCreate {
	stc.mutation.SetTplCategory(s)
	return stc
}

// SetNillableTplCategory sets the "tplCategory" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableTplCategory(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetTplCategory(*s)
	}
	return stc
}

// SetPackageName sets the "packageName" field.
func (stc *SysTablesCreate) SetPackageName(s string) *SysTablesCreate {
	stc.mutation.SetPackageName(s)
	return stc
}

// SetNillablePackageName sets the "packageName" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillablePackageName(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetPackageName(*s)
	}
	return stc
}

// SetModuleName sets the "moduleName" field.
func (stc *SysTablesCreate) SetModuleName(s string) *SysTablesCreate {
	stc.mutation.SetModuleName(s)
	return stc
}

// SetNillableModuleName sets the "moduleName" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableModuleName(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetModuleName(*s)
	}
	return stc
}

// SetModuleFrontName sets the "moduleFrontName" field.
func (stc *SysTablesCreate) SetModuleFrontName(s string) *SysTablesCreate {
	stc.mutation.SetModuleFrontName(s)
	return stc
}

// SetNillableModuleFrontName sets the "moduleFrontName" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableModuleFrontName(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetModuleFrontName(*s)
	}
	return stc
}

// SetBusinessName sets the "businessName" field.
func (stc *SysTablesCreate) SetBusinessName(s string) *SysTablesCreate {
	stc.mutation.SetBusinessName(s)
	return stc
}

// SetNillableBusinessName sets the "businessName" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableBusinessName(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetBusinessName(*s)
	}
	return stc
}

// SetFunctionName sets the "functionName" field.
func (stc *SysTablesCreate) SetFunctionName(s string) *SysTablesCreate {
	stc.mutation.SetFunctionName(s)
	return stc
}

// SetNillableFunctionName sets the "functionName" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableFunctionName(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetFunctionName(*s)
	}
	return stc
}

// SetFunctionAuthor sets the "functionAuthor" field.
func (stc *SysTablesCreate) SetFunctionAuthor(s string) *SysTablesCreate {
	stc.mutation.SetFunctionAuthor(s)
	return stc
}

// SetNillableFunctionAuthor sets the "functionAuthor" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableFunctionAuthor(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetFunctionAuthor(*s)
	}
	return stc
}

// SetPkColumn sets the "pkColumn" field.
func (stc *SysTablesCreate) SetPkColumn(s string) *SysTablesCreate {
	stc.mutation.SetPkColumn(s)
	return stc
}

// SetNillablePkColumn sets the "pkColumn" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillablePkColumn(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetPkColumn(*s)
	}
	return stc
}

// SetPkGoField sets the "pkGoField" field.
func (stc *SysTablesCreate) SetPkGoField(s string) *SysTablesCreate {
	stc.mutation.SetPkGoField(s)
	return stc
}

// SetNillablePkGoField sets the "pkGoField" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillablePkGoField(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetPkGoField(*s)
	}
	return stc
}

// SetPkJsonField sets the "pkJsonField" field.
func (stc *SysTablesCreate) SetPkJsonField(s string) *SysTablesCreate {
	stc.mutation.SetPkJsonField(s)
	return stc
}

// SetNillablePkJsonField sets the "pkJsonField" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillablePkJsonField(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetPkJsonField(*s)
	}
	return stc
}

// SetOptions sets the "options" field.
func (stc *SysTablesCreate) SetOptions(s string) *SysTablesCreate {
	stc.mutation.SetOptions(s)
	return stc
}

// SetNillableOptions sets the "options" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableOptions(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetOptions(*s)
	}
	return stc
}

// SetTreeCode sets the "treeCode" field.
func (stc *SysTablesCreate) SetTreeCode(s string) *SysTablesCreate {
	stc.mutation.SetTreeCode(s)
	return stc
}

// SetNillableTreeCode sets the "treeCode" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableTreeCode(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetTreeCode(*s)
	}
	return stc
}

// SetTreeParentCode sets the "treeParentCode" field.
func (stc *SysTablesCreate) SetTreeParentCode(s string) *SysTablesCreate {
	stc.mutation.SetTreeParentCode(s)
	return stc
}

// SetNillableTreeParentCode sets the "treeParentCode" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableTreeParentCode(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetTreeParentCode(*s)
	}
	return stc
}

// SetTreeName sets the "treeName" field.
func (stc *SysTablesCreate) SetTreeName(s string) *SysTablesCreate {
	stc.mutation.SetTreeName(s)
	return stc
}

// SetNillableTreeName sets the "treeName" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableTreeName(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetTreeName(*s)
	}
	return stc
}

// SetTree sets the "tree" field.
func (stc *SysTablesCreate) SetTree(b bool) *SysTablesCreate {
	stc.mutation.SetTree(b)
	return stc
}

// SetNillableTree sets the "tree" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableTree(b *bool) *SysTablesCreate {
	if b != nil {
		stc.SetTree(*b)
	}
	return stc
}

// SetCrud sets the "crud" field.
func (stc *SysTablesCreate) SetCrud(b bool) *SysTablesCreate {
	stc.mutation.SetCrud(b)
	return stc
}

// SetNillableCrud sets the "crud" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableCrud(b *bool) *SysTablesCreate {
	if b != nil {
		stc.SetCrud(*b)
	}
	return stc
}

// SetRemark sets the "remark" field.
func (stc *SysTablesCreate) SetRemark(s string) *SysTablesCreate {
	stc.mutation.SetRemark(s)
	return stc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableRemark(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetRemark(*s)
	}
	return stc
}

// SetIsDataScope sets the "isDataScope" field.
func (stc *SysTablesCreate) SetIsDataScope(i int32) *SysTablesCreate {
	stc.mutation.SetIsDataScope(i)
	return stc
}

// SetNillableIsDataScope sets the "isDataScope" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableIsDataScope(i *int32) *SysTablesCreate {
	if i != nil {
		stc.SetIsDataScope(*i)
	}
	return stc
}

// SetIsActions sets the "isActions" field.
func (stc *SysTablesCreate) SetIsActions(i int32) *SysTablesCreate {
	stc.mutation.SetIsActions(i)
	return stc
}

// SetNillableIsActions sets the "isActions" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableIsActions(i *int32) *SysTablesCreate {
	if i != nil {
		stc.SetIsActions(*i)
	}
	return stc
}

// SetIsAuth sets the "isAuth" field.
func (stc *SysTablesCreate) SetIsAuth(i int32) *SysTablesCreate {
	stc.mutation.SetIsAuth(i)
	return stc
}

// SetNillableIsAuth sets the "isAuth" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableIsAuth(i *int32) *SysTablesCreate {
	if i != nil {
		stc.SetIsAuth(*i)
	}
	return stc
}

// SetIsLogicalDelete sets the "isLogicalDelete" field.
func (stc *SysTablesCreate) SetIsLogicalDelete(s string) *SysTablesCreate {
	stc.mutation.SetIsLogicalDelete(s)
	return stc
}

// SetNillableIsLogicalDelete sets the "isLogicalDelete" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableIsLogicalDelete(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetIsLogicalDelete(*s)
	}
	return stc
}

// SetLogicalDelete sets the "logicalDelete" field.
func (stc *SysTablesCreate) SetLogicalDelete(b bool) *SysTablesCreate {
	stc.mutation.SetLogicalDelete(b)
	return stc
}

// SetNillableLogicalDelete sets the "logicalDelete" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableLogicalDelete(b *bool) *SysTablesCreate {
	if b != nil {
		stc.SetLogicalDelete(*b)
	}
	return stc
}

// SetLogicalDeleteColumn sets the "logicalDeleteColumn" field.
func (stc *SysTablesCreate) SetLogicalDeleteColumn(s string) *SysTablesCreate {
	stc.mutation.SetLogicalDeleteColumn(s)
	return stc
}

// SetNillableLogicalDeleteColumn sets the "logicalDeleteColumn" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableLogicalDeleteColumn(s *string) *SysTablesCreate {
	if s != nil {
		stc.SetLogicalDeleteColumn(*s)
	}
	return stc
}

// SetCreatedAt sets the "createdAt" field.
func (stc *SysTablesCreate) SetCreatedAt(t time.Time) *SysTablesCreate {
	stc.mutation.SetCreatedAt(t)
	return stc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableCreatedAt(t *time.Time) *SysTablesCreate {
	if t != nil {
		stc.SetCreatedAt(*t)
	}
	return stc
}

// SetUpdatedAt sets the "updatedAt" field.
func (stc *SysTablesCreate) SetUpdatedAt(t time.Time) *SysTablesCreate {
	stc.mutation.SetUpdatedAt(t)
	return stc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableUpdatedAt(t *time.Time) *SysTablesCreate {
	if t != nil {
		stc.SetUpdatedAt(*t)
	}
	return stc
}

// SetCreateBy sets the "createBy" field.
func (stc *SysTablesCreate) SetCreateBy(i int64) *SysTablesCreate {
	stc.mutation.SetCreateBy(i)
	return stc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableCreateBy(i *int64) *SysTablesCreate {
	if i != nil {
		stc.SetCreateBy(*i)
	}
	return stc
}

// SetUpdateBy sets the "updateBy" field.
func (stc *SysTablesCreate) SetUpdateBy(i int64) *SysTablesCreate {
	stc.mutation.SetUpdateBy(i)
	return stc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableUpdateBy(i *int64) *SysTablesCreate {
	if i != nil {
		stc.SetUpdateBy(*i)
	}
	return stc
}

// SetTenantId sets the "tenantId" field.
func (stc *SysTablesCreate) SetTenantId(i int64) *SysTablesCreate {
	stc.mutation.SetTenantId(i)
	return stc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (stc *SysTablesCreate) SetNillableTenantId(i *int64) *SysTablesCreate {
	if i != nil {
		stc.SetTenantId(*i)
	}
	return stc
}

// Mutation returns the SysTablesMutation object of the builder.
func (stc *SysTablesCreate) Mutation() *SysTablesMutation {
	return stc.mutation
}

// Save creates the SysTables in the database.
func (stc *SysTablesCreate) Save(ctx context.Context) (*SysTables, error) {
	var (
		err  error
		node *SysTables
	)
	stc.defaults()
	if len(stc.hooks) == 0 {
		if err = stc.check(); err != nil {
			return nil, err
		}
		node, err = stc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysTablesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = stc.check(); err != nil {
				return nil, err
			}
			stc.mutation = mutation
			if node, err = stc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(stc.hooks) - 1; i >= 0; i-- {
			if stc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (stc *SysTablesCreate) SaveX(ctx context.Context) *SysTables {
	v, err := stc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stc *SysTablesCreate) Exec(ctx context.Context) error {
	_, err := stc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stc *SysTablesCreate) ExecX(ctx context.Context) {
	if err := stc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stc *SysTablesCreate) defaults() {
	if _, ok := stc.mutation.CreatedAt(); !ok {
		v := systables.DefaultCreatedAt()
		stc.mutation.SetCreatedAt(v)
	}
	if _, ok := stc.mutation.UpdatedAt(); !ok {
		v := systables.DefaultUpdatedAt()
		stc.mutation.SetUpdatedAt(v)
	}
	if _, ok := stc.mutation.CreateBy(); !ok {
		v := systables.DefaultCreateBy
		stc.mutation.SetCreateBy(v)
	}
	if _, ok := stc.mutation.UpdateBy(); !ok {
		v := systables.DefaultUpdateBy
		stc.mutation.SetUpdateBy(v)
	}
	if _, ok := stc.mutation.TenantId(); !ok {
		v := systables.DefaultTenantId
		stc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stc *SysTablesCreate) check() error {
	if _, ok := stc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := stc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := stc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := stc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := stc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (stc *SysTablesCreate) sqlSave(ctx context.Context) (*SysTables, error) {
	_node, _spec := stc.createSpec()
	if err := sqlgraph.CreateNode(ctx, stc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (stc *SysTablesCreate) createSpec() (*SysTables, *sqlgraph.CreateSpec) {
	var (
		_node = &SysTables{config: stc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: systables.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: systables.FieldID,
			},
		}
	)
	if value, ok := stc.mutation.TableName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldTableName,
		})
		_node.TableName = value
	}
	if value, ok := stc.mutation.TableComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldTableComment,
		})
		_node.TableComment = value
	}
	if value, ok := stc.mutation.ClassName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldClassName,
		})
		_node.ClassName = value
	}
	if value, ok := stc.mutation.TplCategory(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldTplCategory,
		})
		_node.TplCategory = value
	}
	if value, ok := stc.mutation.PackageName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldPackageName,
		})
		_node.PackageName = value
	}
	if value, ok := stc.mutation.ModuleName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldModuleName,
		})
		_node.ModuleName = value
	}
	if value, ok := stc.mutation.ModuleFrontName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldModuleFrontName,
		})
		_node.ModuleFrontName = value
	}
	if value, ok := stc.mutation.BusinessName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldBusinessName,
		})
		_node.BusinessName = value
	}
	if value, ok := stc.mutation.FunctionName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldFunctionName,
		})
		_node.FunctionName = value
	}
	if value, ok := stc.mutation.FunctionAuthor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldFunctionAuthor,
		})
		_node.FunctionAuthor = value
	}
	if value, ok := stc.mutation.PkColumn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldPkColumn,
		})
		_node.PkColumn = value
	}
	if value, ok := stc.mutation.PkGoField(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldPkGoField,
		})
		_node.PkGoField = value
	}
	if value, ok := stc.mutation.PkJsonField(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldPkJsonField,
		})
		_node.PkJsonField = value
	}
	if value, ok := stc.mutation.Options(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldOptions,
		})
		_node.Options = value
	}
	if value, ok := stc.mutation.TreeCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldTreeCode,
		})
		_node.TreeCode = value
	}
	if value, ok := stc.mutation.TreeParentCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldTreeParentCode,
		})
		_node.TreeParentCode = value
	}
	if value, ok := stc.mutation.TreeName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldTreeName,
		})
		_node.TreeName = value
	}
	if value, ok := stc.mutation.Tree(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: systables.FieldTree,
		})
		_node.Tree = value
	}
	if value, ok := stc.mutation.Crud(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: systables.FieldCrud,
		})
		_node.Crud = value
	}
	if value, ok := stc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := stc.mutation.IsDataScope(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systables.FieldIsDataScope,
		})
		_node.IsDataScope = value
	}
	if value, ok := stc.mutation.IsActions(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systables.FieldIsActions,
		})
		_node.IsActions = value
	}
	if value, ok := stc.mutation.IsAuth(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: systables.FieldIsAuth,
		})
		_node.IsAuth = value
	}
	if value, ok := stc.mutation.IsLogicalDelete(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldIsLogicalDelete,
		})
		_node.IsLogicalDelete = value
	}
	if value, ok := stc.mutation.LogicalDelete(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: systables.FieldLogicalDelete,
		})
		_node.LogicalDelete = value
	}
	if value, ok := stc.mutation.LogicalDeleteColumn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systables.FieldLogicalDeleteColumn,
		})
		_node.LogicalDeleteColumn = value
	}
	if value, ok := stc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systables.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := stc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systables.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := stc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systables.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := stc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systables.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := stc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: systables.FieldTenantId,
		})
		_node.TenantId = value
	}
	return _node, _spec
}

// SysTablesCreateBulk is the builder for creating many SysTables entities in bulk.
type SysTablesCreateBulk struct {
	config
	builders []*SysTablesCreate
}

// Save creates the SysTables entities in the database.
func (stcb *SysTablesCreateBulk) Save(ctx context.Context) ([]*SysTables, error) {
	specs := make([]*sqlgraph.CreateSpec, len(stcb.builders))
	nodes := make([]*SysTables, len(stcb.builders))
	mutators := make([]Mutator, len(stcb.builders))
	for i := range stcb.builders {
		func(i int, root context.Context) {
			builder := stcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysTablesMutation)
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
					_, err = mutators[i+1].Mutate(root, stcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, stcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, stcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (stcb *SysTablesCreateBulk) SaveX(ctx context.Context) []*SysTables {
	v, err := stcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stcb *SysTablesCreateBulk) Exec(ctx context.Context) error {
	_, err := stcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stcb *SysTablesCreateBulk) ExecX(ctx context.Context) {
	if err := stcb.Exec(ctx); err != nil {
		panic(err)
	}
}
