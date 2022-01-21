// Code generated by entc, DO NOT EDIT.

package sysloginlog

import (
	"time"
)

const (
	// Label holds the string label denoting the sysloginlog type in the database.
	Label = "sys_login_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldIpaddr holds the string denoting the ipaddr field in the database.
	FieldIpaddr = "ipaddr"
	// FieldLoginLocation holds the string denoting the loginlocation field in the database.
	FieldLoginLocation = "login_location"
	// FieldBrowser holds the string denoting the browser field in the database.
	FieldBrowser = "browser"
	// FieldOs holds the string denoting the os field in the database.
	FieldOs = "os"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldLoginTime holds the string denoting the logintime field in the database.
	FieldLoginTime = "login_time"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldMsg holds the string denoting the msg field in the database.
	FieldMsg = "msg"
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
	// Table holds the table name of the sysloginlog in the database.
	Table = "sys_login_logs"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "sys_login_logs"
	// UserInverseTable is the table name for the SysUser entity.
	// It exists in this package in order to avoid circular dependency with the "sysuser" package.
	UserInverseTable = "sys_users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for sysloginlog fields.
var Columns = []string{
	FieldID,
	FieldUserId,
	FieldStatus,
	FieldIpaddr,
	FieldLoginLocation,
	FieldBrowser,
	FieldOs,
	FieldPlatform,
	FieldLoginTime,
	FieldRemark,
	FieldMsg,
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
