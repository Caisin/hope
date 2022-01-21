// Code generated by entc, DO NOT EDIT.

package sysmenu

import (
	"time"
)

const (
	// Label holds the string label denoting the sysmenu type in the database.
	Label = "sys_menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMenuName holds the string denoting the menuname field in the database.
	FieldMenuName = "menu_name"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldPaths holds the string denoting the paths field in the database.
	FieldPaths = "paths"
	// FieldMenuType holds the string denoting the menutype field in the database.
	FieldMenuType = "menu_type"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldPermission holds the string denoting the permission field in the database.
	FieldPermission = "permission"
	// FieldNoCache holds the string denoting the nocache field in the database.
	FieldNoCache = "no_cache"
	// FieldBreadcrumb holds the string denoting the breadcrumb field in the database.
	FieldBreadcrumb = "breadcrumb"
	// FieldComponent holds the string denoting the component field in the database.
	FieldComponent = "component"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldVisible holds the string denoting the visible field in the database.
	FieldVisible = "visible"
	// FieldIsFrame holds the string denoting the isframe field in the database.
	FieldIsFrame = "is_frame"
	// FieldSysApi holds the string denoting the sysapi field in the database.
	FieldSysApi = "sys_api"
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
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildes holds the string denoting the childes edge name in mutations.
	EdgeChildes = "childes"
	// Table holds the table name of the sysmenu in the database.
	Table = "sys_menus"
	// RoleTable is the table that holds the role relation/edge. The primary key declared below.
	RoleTable = "sys_role_menus"
	// RoleInverseTable is the table name for the SysRole entity.
	// It exists in this package in order to avoid circular dependency with the "sysrole" package.
	RoleInverseTable = "sys_roles"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "sys_menus"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "sys_menu_childes"
	// ChildesTable is the table that holds the childes relation/edge.
	ChildesTable = "sys_menus"
	// ChildesColumn is the table column denoting the childes relation/edge.
	ChildesColumn = "sys_menu_childes"
)

// Columns holds all SQL columns for sysmenu fields.
var Columns = []string{
	FieldID,
	FieldMenuName,
	FieldTitle,
	FieldIcon,
	FieldPath,
	FieldPaths,
	FieldMenuType,
	FieldAction,
	FieldPermission,
	FieldNoCache,
	FieldBreadcrumb,
	FieldComponent,
	FieldSort,
	FieldVisible,
	FieldIsFrame,
	FieldSysApi,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreateBy,
	FieldUpdateBy,
	FieldTenantId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sys_menus"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"sys_menu_childes",
}

var (
	// RolePrimaryKey and RoleColumn2 are the table columns denoting the
	// primary key for the role relation (M2M).
	RolePrimaryKey = []string{"sys_role_id", "sys_menu_id"}
)

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