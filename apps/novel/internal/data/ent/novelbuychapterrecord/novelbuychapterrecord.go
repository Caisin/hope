// Code generated by entc, DO NOT EDIT.

package novelbuychapterrecord

import (
	"time"
)

const (
	// Label holds the string label denoting the novelbuychapterrecord type in the database.
	Label = "novel_buy_chapter_record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldUserName holds the string denoting the username field in the database.
	FieldUserName = "user_name"
	// FieldChapterId holds the string denoting the chapterid field in the database.
	FieldChapterId = "chapter_id"
	// FieldChapterOrderNum holds the string denoting the chapterordernum field in the database.
	FieldChapterOrderNum = "chapter_order_num"
	// FieldNovelId holds the string denoting the novelid field in the database.
	FieldNovelId = "novel_id"
	// FieldNovelName holds the string denoting the novelname field in the database.
	FieldNovelName = "novel_name"
	// FieldChapterName holds the string denoting the chaptername field in the database.
	FieldChapterName = "chapter_name"
	// FieldIsSvip holds the string denoting the issvip field in the database.
	FieldIsSvip = "is_svip"
	// FieldCoin holds the string denoting the coin field in the database.
	FieldCoin = "coin"
	// FieldCoupon holds the string denoting the coupon field in the database.
	FieldCoupon = "coupon"
	// FieldDiscount holds the string denoting the discount field in the database.
	FieldDiscount = "discount"
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
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the novelbuychapterrecord in the database.
	Table = "novel_buy_chapter_records"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "novel_buy_chapter_records"
	// UserInverseTable is the table name for the SocialUser entity.
	// It exists in this package in order to avoid circular dependency with the "socialuser" package.
	UserInverseTable = "social_users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "social_user_buy_chapter_records"
)

// Columns holds all SQL columns for novelbuychapterrecord fields.
var Columns = []string{
	FieldID,
	FieldUserId,
	FieldUserName,
	FieldChapterId,
	FieldChapterOrderNum,
	FieldNovelId,
	FieldNovelName,
	FieldChapterName,
	FieldIsSvip,
	FieldCoin,
	FieldCoupon,
	FieldDiscount,
	FieldRemark,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreateBy,
	FieldUpdateBy,
	FieldTenantId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "novel_buy_chapter_records"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"social_user_buy_chapter_records",
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
