// Code generated by entc, DO NOT EDIT.

package vipuser

import (
	"time"
)

const (
	// Label holds the string label denoting the vipuser type in the database.
	Label = "vip_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVipType holds the string denoting the viptype field in the database.
	FieldVipType = "vip_type"
	// FieldSvipType holds the string denoting the sviptype field in the database.
	FieldSvipType = "svip_type"
	// FieldSvipEffectTime holds the string denoting the svipeffecttime field in the database.
	FieldSvipEffectTime = "svip_effect_time"
	// FieldSvipExpiredTime holds the string denoting the svipexpiredtime field in the database.
	FieldSvipExpiredTime = "svip_expired_time"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldEffectTime holds the string denoting the effecttime field in the database.
	FieldEffectTime = "effect_time"
	// FieldExpiredTime holds the string denoting the expiredtime field in the database.
	FieldExpiredTime = "expired_time"
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
	// Table holds the table name of the vipuser in the database.
	Table = "vip_users"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "social_user_vips"
	// UserInverseTable is the table name for the SocialUser entity.
	// It exists in this package in order to avoid circular dependency with the "socialuser" package.
	UserInverseTable = "social_users"
)

// Columns holds all SQL columns for vipuser fields.
var Columns = []string{
	FieldID,
	FieldVipType,
	FieldSvipType,
	FieldSvipEffectTime,
	FieldSvipExpiredTime,
	FieldRemark,
	FieldEffectTime,
	FieldExpiredTime,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreateBy,
	FieldUpdateBy,
	FieldTenantId,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"social_user_id", "vip_user_id"}
)

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
	// DefaultEffectTime holds the default value on creation for the "effectTime" field.
	DefaultEffectTime func() time.Time
	// DefaultExpiredTime holds the default value on creation for the "expiredTime" field.
	DefaultExpiredTime func() time.Time
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