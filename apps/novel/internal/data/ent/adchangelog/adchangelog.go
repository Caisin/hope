// Code generated by entc, DO NOT EDIT.

package adchangelog

import (
	"time"
)

const (
	// Label holds the string label denoting the adchangelog type in the database.
	Label = "ad_change_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldAdId holds the string denoting the adid field in the database.
	FieldAdId = "ad_id"
	// FieldChId holds the string denoting the chid field in the database.
	FieldChId = "ch_id"
	// FieldDeviceId holds the string denoting the deviceid field in the database.
	FieldDeviceId = "device_id"
	// FieldExtInfo holds the string denoting the extinfo field in the database.
	FieldExtInfo = "ext_info"
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
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the adchangelog in the database.
	Table = "ad_change_logs"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "ad_change_logs"
	// UserInverseTable is the table name for the SocialUser entity.
	// It exists in this package in order to avoid circular dependency with the "socialuser" package.
	UserInverseTable = "social_users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for adchangelog fields.
var Columns = []string{
	FieldID,
	FieldUserId,
	FieldAdId,
	FieldChId,
	FieldDeviceId,
	FieldExtInfo,
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
