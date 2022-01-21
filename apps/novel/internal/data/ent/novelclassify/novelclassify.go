// Code generated by entc, DO NOT EDIT.

package novelclassify

import (
	"time"
)

const (
	// Label holds the string label denoting the novelclassify type in the database.
	Label = "novel_classify"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldClassifyName holds the string denoting the classifyname field in the database.
	FieldClassifyName = "classify_name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldOrderNum holds the string denoting the ordernum field in the database.
	FieldOrderNum = "order_num"
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
	// EdgeNovels holds the string denoting the novels edge name in mutations.
	EdgeNovels = "novels"
	// Table holds the table name of the novelclassify in the database.
	Table = "novel_classifies"
	// NovelsTable is the table that holds the novels relation/edge.
	NovelsTable = "novels"
	// NovelsInverseTable is the table name for the Novel entity.
	// It exists in this package in order to avoid circular dependency with the "novel" package.
	NovelsInverseTable = "novels"
	// NovelsColumn is the table column denoting the novels relation/edge.
	NovelsColumn = "classify_id"
)

// Columns holds all SQL columns for novelclassify fields.
var Columns = []string{
	FieldID,
	FieldPid,
	FieldClassifyName,
	FieldStatus,
	FieldOrderNum,
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
