// Code generated by entc, DO NOT EDIT.

package resourcestorage

import (
	"time"
)

const (
	// Label holds the string label denoting the resourcestorage type in the database.
	Label = "resource_storage"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGroupId holds the string denoting the groupid field in the database.
	FieldGroupId = "group_id"
	// FieldStorageType holds the string denoting the storagetype field in the database.
	FieldStorageType = "storage_type"
	// FieldRealName holds the string denoting the realname field in the database.
	FieldRealName = "real_name"
	// FieldBucket holds the string denoting the bucket field in the database.
	FieldBucket = "bucket"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSuffix holds the string denoting the suffix field in the database.
	FieldSuffix = "suffix"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldDeleteUrl holds the string denoting the deleteurl field in the database.
	FieldDeleteUrl = "delete_url"
	// FieldFilename holds the string denoting the filename field in the database.
	FieldFilename = "filename"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldMd5code holds the string denoting the md5code field in the database.
	FieldMd5code = "md5code"
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
	// Table holds the table name of the resourcestorage in the database.
	Table = "resource_storages"
)

// Columns holds all SQL columns for resourcestorage fields.
var Columns = []string{
	FieldID,
	FieldGroupId,
	FieldStorageType,
	FieldRealName,
	FieldBucket,
	FieldName,
	FieldSuffix,
	FieldPath,
	FieldType,
	FieldSize,
	FieldDeleteUrl,
	FieldFilename,
	FieldKey,
	FieldHeight,
	FieldURL,
	FieldUsername,
	FieldWidth,
	FieldMd5code,
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
