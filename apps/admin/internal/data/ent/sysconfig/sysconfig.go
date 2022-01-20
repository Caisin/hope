// Code generated by entc, DO NOT EDIT.

package sysconfig

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the sysconfig type in the database.
	Label = "sys_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldConfigName holds the string denoting the configname field in the database.
	FieldConfigName = "config_name"
	// FieldConfigKey holds the string denoting the configkey field in the database.
	FieldConfigKey = "config_key"
	// FieldConfigValue holds the string denoting the configvalue field in the database.
	FieldConfigValue = "config_value"
	// FieldConfigType holds the string denoting the configtype field in the database.
	FieldConfigType = "config_type"
	// FieldIsFrontend holds the string denoting the isfrontend field in the database.
	FieldIsFrontend = "is_frontend"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
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
	// Table holds the table name of the sysconfig in the database.
	Table = "sys_configs"
)

// Columns holds all SQL columns for sysconfig fields.
var Columns = []string{
	FieldID,
	FieldConfigName,
	FieldConfigKey,
	FieldConfigValue,
	FieldConfigType,
	FieldIsFrontend,
	FieldState,
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

// State defines the type for the "state" enum field.
type State string

// StateU is the default value of the State enum.
const DefaultState = StateU

// State values.
const (
	StateU State = "U"
	StateE State = "E"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateU, StateE:
		return nil
	default:
		return fmt.Errorf("sysconfig: invalid enum value for state field: %q", s)
	}
}
