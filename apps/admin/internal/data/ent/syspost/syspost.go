// Code generated by entc, DO NOT EDIT.

package syspost

import (
	"time"
)

const (
	// Label holds the string label denoting the syspost type in the database.
	Label = "sys_post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPostName holds the string denoting the postname field in the database.
	FieldPostName = "post_name"
	// FieldPostCode holds the string denoting the postcode field in the database.
	FieldPostCode = "post_code"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
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
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the syspost in the database.
	Table = "sys_posts"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "sys_users"
	// UsersInverseTable is the table name for the SysUser entity.
	// It exists in this package in order to avoid circular dependency with the "sysuser" package.
	UsersInverseTable = "sys_users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "sys_post_users"
)

// Columns holds all SQL columns for syspost fields.
var Columns = []string{
	FieldID,
	FieldPostName,
	FieldPostCode,
	FieldSort,
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
