// Code generated by entc, DO NOT EDIT.

package ambalance

import (
	"time"
)

const (
	// Label holds the string label denoting the ambalance type in the database.
	Label = "am_balance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldOrderId holds the string denoting the orderid field in the database.
	FieldOrderId = "order_id"
	// FieldEventId holds the string denoting the eventid field in the database.
	FieldEventId = "event_id"
	// FieldCashTag holds the string denoting the cashtag field in the database.
	FieldCashTag = "cash_tag"
	// FieldAssetItemId holds the string denoting the assetitemid field in the database.
	FieldAssetItemId = "asset_item_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
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
	// Table holds the table name of the ambalance in the database.
	Table = "am_balances"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "am_balances"
	// UserInverseTable is the table name for the SocialUser entity.
	// It exists in this package in order to avoid circular dependency with the "socialuser" package.
	UserInverseTable = "social_users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for ambalance fields.
var Columns = []string{
	FieldID,
	FieldUserId,
	FieldOrderId,
	FieldEventId,
	FieldCashTag,
	FieldAssetItemId,
	FieldAmount,
	FieldBalance,
	FieldRemark,
	FieldEffectTime,
	FieldExpiredTime,
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
