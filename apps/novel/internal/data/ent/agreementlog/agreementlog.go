// Code generated by entc, DO NOT EDIT.

package agreementlog

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the agreementlog type in the database.
	Label = "agreement_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOuterAgreementNo holds the string denoting the outeragreementno field in the database.
	FieldOuterAgreementNo = "outer_agreement_no"
	// FieldOrderId holds the string denoting the orderid field in the database.
	FieldOrderId = "order_id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldChId holds the string denoting the chid field in the database.
	FieldChId = "ch_id"
	// FieldUserName holds the string denoting the username field in the database.
	FieldUserName = "user_name"
	// FieldPaymentName holds the string denoting the paymentname field in the database.
	FieldPaymentName = "payment_name"
	// FieldPaymentId holds the string denoting the paymentid field in the database.
	FieldPaymentId = "payment_id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldPayment holds the string denoting the payment field in the database.
	FieldPayment = "payment"
	// FieldAgreementType holds the string denoting the agreementtype field in the database.
	FieldAgreementType = "agreement_type"
	// FieldVipType holds the string denoting the viptype field in the database.
	FieldVipType = "vip_type"
	// FieldTimes holds the string denoting the times field in the database.
	FieldTimes = "times"
	// FieldCycleDays holds the string denoting the cycledays field in the database.
	FieldCycleDays = "cycle_days"
	// FieldNextExecTime holds the string denoting the nextexectime field in the database.
	FieldNextExecTime = "next_exec_time"
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
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// Table holds the table name of the agreementlog in the database.
	Table = "agreement_logs"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "pay_orders"
	// OrdersInverseTable is the table name for the PayOrder entity.
	// It exists in this package in order to avoid circular dependency with the "payorder" package.
	OrdersInverseTable = "pay_orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "agreement_id"
)

// Columns holds all SQL columns for agreementlog fields.
var Columns = []string{
	FieldID,
	FieldOuterAgreementNo,
	FieldOrderId,
	FieldUserId,
	FieldChId,
	FieldUserName,
	FieldPaymentName,
	FieldPaymentId,
	FieldState,
	FieldPayment,
	FieldAgreementType,
	FieldVipType,
	FieldTimes,
	FieldCycleDays,
	FieldNextExecTime,
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

// AgreementType defines the type for the "agreementType" enum field.
type AgreementType string

// AgreementType values.
const (
	AgreementTypeAlipay AgreementType = "Alipay"
	AgreementTypeGoogle AgreementType = "Google"
	AgreementTypeWeChat AgreementType = "WeChat"
	AgreementTypeApple  AgreementType = "Apple"
)

func (at AgreementType) String() string {
	return string(at)
}

// AgreementTypeValidator is a validator for the "agreementType" field enum values. It is called by the builders before save.
func AgreementTypeValidator(at AgreementType) error {
	switch at {
	case AgreementTypeAlipay, AgreementTypeGoogle, AgreementTypeWeChat, AgreementTypeApple:
		return nil
	default:
		return fmt.Errorf("agreementlog: invalid enum value for agreementType field: %q", at)
	}
}
