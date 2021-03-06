// Code generated by entc, DO NOT EDIT.

package viptype

import (
	"time"
)

const (
	// Label holds the string label denoting the viptype type in the database.
	Label = "vip_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVipName holds the string denoting the vipname field in the database.
	FieldVipName = "vip_name"
	// FieldIsSuper holds the string denoting the issuper field in the database.
	FieldIsSuper = "is_super"
	// FieldValidDays holds the string denoting the validdays field in the database.
	FieldValidDays = "valid_days"
	// FieldDiscountRate holds the string denoting the discountrate field in the database.
	FieldDiscountRate = "discount_rate"
	// FieldAvatarId holds the string denoting the avatarid field in the database.
	FieldAvatarId = "avatar_id"
	// FieldSummary holds the string denoting the summary field in the database.
	FieldSummary = "summary"
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
	// Table holds the table name of the viptype in the database.
	Table = "vip_types"
)

// Columns holds all SQL columns for viptype fields.
var Columns = []string{
	FieldID,
	FieldVipName,
	FieldIsSuper,
	FieldValidDays,
	FieldDiscountRate,
	FieldAvatarId,
	FieldSummary,
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
