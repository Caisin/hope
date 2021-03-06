// Code generated by entc, DO NOT EDIT.

package casbinrule

import (
	"time"
)

const (
	// Label holds the string label denoting the casbinrule type in the database.
	Label = "casbin_rule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPType holds the string denoting the p_type field in the database.
	FieldPType = "p_type"
	// FieldV0 holds the string denoting the v0 field in the database.
	FieldV0 = "v0"
	// FieldV1 holds the string denoting the v1 field in the database.
	FieldV1 = "v1"
	// FieldV2 holds the string denoting the v2 field in the database.
	FieldV2 = "v2"
	// FieldV3 holds the string denoting the v3 field in the database.
	FieldV3 = "v3"
	// FieldV4 holds the string denoting the v4 field in the database.
	FieldV4 = "v4"
	// FieldV5 holds the string denoting the v5 field in the database.
	FieldV5 = "v5"
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
	// Table holds the table name of the casbinrule in the database.
	Table = "casbin_rules"
)

// Columns holds all SQL columns for casbinrule fields.
var Columns = []string{
	FieldID,
	FieldPType,
	FieldV0,
	FieldV1,
	FieldV2,
	FieldV3,
	FieldV4,
	FieldV5,
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
