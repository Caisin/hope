// Code generated by entc, DO NOT EDIT.

package clienterror

import (
	"time"
)

const (
	// Label holds the string label denoting the clienterror type in the database.
	Label = "client_error"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppVersion holds the string denoting the appversion field in the database.
	FieldAppVersion = "app_version"
	// FieldDeviceName holds the string denoting the devicename field in the database.
	FieldDeviceName = "device_name"
	// FieldOsName holds the string denoting the osname field in the database.
	FieldOsName = "os_name"
	// FieldErrorInfo holds the string denoting the errorinfo field in the database.
	FieldErrorInfo = "error_info"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
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
	// Table holds the table name of the clienterror in the database.
	Table = "client_errors"
)

// Columns holds all SQL columns for clienterror fields.
var Columns = []string{
	FieldID,
	FieldAppVersion,
	FieldDeviceName,
	FieldOsName,
	FieldErrorInfo,
	FieldUserId,
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
