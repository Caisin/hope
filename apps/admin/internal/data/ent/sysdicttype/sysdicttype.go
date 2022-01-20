// Code generated by entc, DO NOT EDIT.

package sysdicttype

import (
	"time"
)

const (
	// Label holds the string label denoting the sysdicttype type in the database.
	Label = "sys_dict_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDictName holds the string denoting the dictname field in the database.
	FieldDictName = "dict_name"
	// FieldDictType holds the string denoting the dicttype field in the database.
	FieldDictType = "dict_type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
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
	// EdgeDataList holds the string denoting the datalist edge name in mutations.
	EdgeDataList = "dataList"
	// Table holds the table name of the sysdicttype in the database.
	Table = "sys_dict_types"
	// DataListTable is the table that holds the dataList relation/edge.
	DataListTable = "sys_dict_data"
	// DataListInverseTable is the table name for the SysDictData entity.
	// It exists in this package in order to avoid circular dependency with the "sysdictdata" package.
	DataListInverseTable = "sys_dict_data"
	// DataListColumn is the table column denoting the dataList relation/edge.
	DataListColumn = "sys_dict_type_data_list"
)

// Columns holds all SQL columns for sysdicttype fields.
var Columns = []string{
	FieldID,
	FieldDictName,
	FieldDictType,
	FieldStatus,
	FieldRemark,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreateBy,
	FieldUpdateBy,
	FieldTenantId,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
