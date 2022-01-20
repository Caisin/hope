// Code generated by entc, DO NOT EDIT.

package sysdictdata

import (
	"time"
)

const (
	// Label holds the string label denoting the sysdictdata type in the database.
	Label = "sys_dict_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDictSort holds the string denoting the dictsort field in the database.
	FieldDictSort = "dict_sort"
	// FieldDictLabel holds the string denoting the dictlabel field in the database.
	FieldDictLabel = "dict_label"
	// FieldDictValue holds the string denoting the dictvalue field in the database.
	FieldDictValue = "dict_value"
	// FieldIsDefault holds the string denoting the isdefault field in the database.
	FieldIsDefault = "is_default"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDefault holds the string denoting the default field in the database.
	FieldDefault = "default"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreateBy holds the string denoting the createby field in the database.
	FieldCreateBy = "create_by"
	// FieldUpdateBy holds the string denoting the updateby field in the database.
	FieldUpdateBy = "update_by"
	// FieldTenantId holds the string denoting the tenantid field in the database.
	FieldTenantId = "tenant_id"
	// EdgeDictType holds the string denoting the dicttype edge name in mutations.
	EdgeDictType = "dictType"
	// Table holds the table name of the sysdictdata in the database.
	Table = "sys_dict_data"
	// DictTypeTable is the table that holds the dictType relation/edge.
	DictTypeTable = "sys_dict_data"
	// DictTypeInverseTable is the table name for the SysDictType entity.
	// It exists in this package in order to avoid circular dependency with the "sysdicttype" package.
	DictTypeInverseTable = "sys_dict_types"
	// DictTypeColumn is the table column denoting the dictType relation/edge.
	DictTypeColumn = "sys_dict_type_data_list"
)

// Columns holds all SQL columns for sysdictdata fields.
var Columns = []string{
	FieldID,
	FieldDictSort,
	FieldDictLabel,
	FieldDictValue,
	FieldIsDefault,
	FieldStatus,
	FieldDefault,
	FieldRemark,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreateBy,
	FieldUpdateBy,
	FieldTenantId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sys_dict_data"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"sys_dict_type_data_list",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultCreateBy holds the default value on creation for the "createBy" field.
	DefaultCreateBy int64
	// DefaultUpdateBy holds the default value on creation for the "updateBy" field.
	DefaultUpdateBy int64
	// DefaultTenantId holds the default value on creation for the "tenantId" field.
	DefaultTenantId int64
)
