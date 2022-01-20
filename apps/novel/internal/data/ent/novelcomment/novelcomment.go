// Code generated by entc, DO NOT EDIT.

package novelcomment

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the novelcomment type in the database.
	Label = "novel_comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNovelId holds the string denoting the novelid field in the database.
	FieldNovelId = "novel_id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldUserName holds the string denoting the username field in the database.
	FieldUserName = "user_name"
	// FieldRepUserId holds the string denoting the repuserid field in the database.
	FieldRepUserId = "rep_user_id"
	// FieldRepUserName holds the string denoting the repusername field in the database.
	FieldRepUserName = "rep_user_name"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldPId holds the string denoting the pid field in the database.
	FieldPId = "p_id"
	// FieldIsTop holds the string denoting the istop field in the database.
	FieldIsTop = "is_top"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldIsHighlight holds the string denoting the ishighlight field in the database.
	FieldIsHighlight = "is_highlight"
	// FieldIsHot holds the string denoting the ishot field in the database.
	FieldIsHot = "is_hot"
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
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildes holds the string denoting the childes edge name in mutations.
	EdgeChildes = "childes"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the novelcomment in the database.
	Table = "novel_comments"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "novel_comments"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "novel_comment_childes"
	// ChildesTable is the table that holds the childes relation/edge.
	ChildesTable = "novel_comments"
	// ChildesColumn is the table column denoting the childes relation/edge.
	ChildesColumn = "novel_comment_childes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "novel_comments"
	// UserInverseTable is the table name for the SocialUser entity.
	// It exists in this package in order to avoid circular dependency with the "socialuser" package.
	UserInverseTable = "social_users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "social_user_comments"
)

// Columns holds all SQL columns for novelcomment fields.
var Columns = []string{
	FieldID,
	FieldNovelId,
	FieldUserId,
	FieldAvatar,
	FieldUserName,
	FieldRepUserId,
	FieldRepUserName,
	FieldContent,
	FieldScore,
	FieldPId,
	FieldIsTop,
	FieldState,
	FieldIsHighlight,
	FieldIsHot,
	FieldRemark,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreateBy,
	FieldUpdateBy,
	FieldTenantId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "novel_comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"novel_comment_childes",
	"social_user_comments",
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

// State defines the type for the "state" enum field.
type State string

// State values.
const (
	State0 State = "0"
	State1 State = "1"
	State2 State = "2"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case State0, State1, State2:
		return nil
	default:
		return fmt.Errorf("novelcomment: invalid enum value for state field: %q", s)
	}
}
