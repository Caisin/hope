// Code generated by entc, DO NOT EDIT.

package sysapi

import (
	"time"
)

const (
	// Label holds the string label denoting the sysapi type in the database.
	Label = "sys_api"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHandle holds the string denoting the handle field in the database.
	FieldHandle = "handle"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
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
	// Table holds the table name of the sysapi in the database.
	Table = "sys_apis"
)

// Columns holds all SQL columns for sysapi fields.
var Columns = []string{
	FieldID,
	FieldHandle,
	FieldTitle,
	FieldPath,
	FieldAction,
	FieldType,
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
