// Code generated by entc, DO NOT EDIT.

package userevent

import (
	"time"
)

const (
	// Label holds the string label denoting the userevent type in the database.
	Label = "user_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldEventType holds the string denoting the eventtype field in the database.
	FieldEventType = "event_type"
	// FieldNovelId holds the string denoting the novelid field in the database.
	FieldNovelId = "novel_id"
	// FieldChapterId holds the string denoting the chapterid field in the database.
	FieldChapterId = "chapter_id"
	// FieldCoin holds the string denoting the coin field in the database.
	FieldCoin = "coin"
	// FieldCoupon holds the string denoting the coupon field in the database.
	FieldCoupon = "coupon"
	// FieldMoney holds the string denoting the money field in the database.
	FieldMoney = "money"
	// FieldKeyword holds the string denoting the keyword field in the database.
	FieldKeyword = "keyword"
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
	// Table holds the table name of the userevent in the database.
	Table = "user_events"
)

// Columns holds all SQL columns for userevent fields.
var Columns = []string{
	FieldID,
	FieldUserId,
	FieldEventType,
	FieldNovelId,
	FieldChapterId,
	FieldCoin,
	FieldCoupon,
	FieldMoney,
	FieldKeyword,
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