// Code generated by entc, DO NOT EDIT.

package sysdept

import (
	"time"
)

const (
	// Label holds the string label denoting the sysdept type in the database.
	Label = "sys_dept"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeptPath holds the string denoting the deptpath field in the database.
	FieldDeptPath = "dept_path"
	// FieldDeptName holds the string denoting the deptname field in the database.
	FieldDeptName = "dept_name"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldLeader holds the string denoting the leader field in the database.
	FieldLeader = "leader"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
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
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildes holds the string denoting the childes edge name in mutations.
	EdgeChildes = "childes"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the sysdept in the database.
	Table = "sys_depts"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "sys_depts"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "sys_dept_childes"
	// ChildesTable is the table that holds the childes relation/edge.
	ChildesTable = "sys_depts"
	// ChildesColumn is the table column denoting the childes relation/edge.
	ChildesColumn = "sys_dept_childes"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "sys_users"
	// UsersInverseTable is the table name for the SysUser entity.
	// It exists in this package in order to avoid circular dependency with the "sysuser" package.
	UsersInverseTable = "sys_users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "sys_dept_users"
)

// Columns holds all SQL columns for sysdept fields.
var Columns = []string{
	FieldID,
	FieldDeptPath,
	FieldDeptName,
	FieldSort,
	FieldLeader,
	FieldPhone,
	FieldEmail,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreateBy,
	FieldUpdateBy,
	FieldTenantId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sys_depts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"sys_dept_childes",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
