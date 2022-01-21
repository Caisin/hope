// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/novel/internal/data/ent/adchannel"
	"hope/apps/novel/internal/data/ent/agreementlog"
	"hope/apps/novel/internal/data/ent/payorder"
	"hope/apps/novel/internal/data/ent/schema"
	"hope/apps/novel/internal/data/ent/socialuser"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// PayOrder is the model entity for the PayOrder schema.
type PayOrder struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// OrderId holds the value of the "orderId" field.
	// 订单号
	OrderId string `json:"orderId,omitempty"`
	// UserId holds the value of the "userId" field.
	// 用户ID
	UserId int64 `json:"userId,omitempty"`
	// ChId holds the value of the "chId" field.
	// 渠道ID
	ChId int64 `json:"chId,omitempty"`
	// AgreementId holds the value of the "agreementId" field.
	// 签约协议号
	AgreementId int64 `json:"agreementId,omitempty"`
	// LastRead holds the value of the "lastRead" field.
	// 最后阅读书籍
	LastRead string `json:"lastRead,omitempty"`
	// LastChapter holds the value of the "lastChapter" field.
	// 最后阅读章节
	LastChapter string `json:"lastChapter,omitempty"`
	// PaymentName holds the value of the "paymentName" field.
	// 充值配置名称
	PaymentName string `json:"paymentName,omitempty"`
	// PaymentId holds the value of the "paymentId" field.
	// 充值配置ID
	PaymentId string `json:"paymentId,omitempty"`
	// State holds the value of the "state" field.
	// 订单状态:1、待付款，2、已付款，3、取消，4、退款
	State schema.OrderState `json:"state,omitempty"`
	// Payment holds the value of the "payment" field.
	// 充值金额
	Payment int64 `json:"payment,omitempty"`
	// PaymentTime holds the value of the "paymentTime" field.
	// 支付时间
	PaymentTime time.Time `json:"paymentTime,omitempty"`
	// CloseTime holds the value of the "closeTime" field.
	// 订单关闭时间
	CloseTime time.Time `json:"closeTime,omitempty"`
	// PayType holds the value of the "payType" field.
	// 支付类型
	PayType payorder.PayType `json:"payType,omitempty"`
	// Coin holds the value of the "coin" field.
	// 书币
	Coin int64 `json:"coin,omitempty"`
	// Coupon holds the value of the "coupon" field.
	// 书券
	Coupon int64 `json:"coupon,omitempty"`
	// VipDays holds the value of the "vipDays" field.
	// vip天数
	VipDays string `json:"vipDays,omitempty"`
	// VipType holds the value of the "vipType" field.
	// vip类型
	VipType int64 `json:"vipType,omitempty"`
	// VipName holds the value of the "vipName" field.
	// vip名称
	VipName string `json:"vipName,omitempty"`
	// Times holds the value of the "times" field.
	// 充值次数
	Times int32 `json:"times,omitempty"`
	// OtherOrderId holds the value of the "otherOrderId" field.
	// 第三方订单
	OtherOrderId string `json:"otherOrderId,omitempty"`
	// Remark holds the value of the "remark" field.
	// 备注
	Remark string `json:"remark,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	// 创建时间
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	// 更新时间
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// CreateBy holds the value of the "createBy" field.
	// 创建者
	CreateBy int64 `json:"createBy,omitempty"`
	// UpdateBy holds the value of the "updateBy" field.
	// 更新者
	UpdateBy int64 `json:"updateBy,omitempty"`
	// TenantId holds the value of the "tenantId" field.
	// 租户
	TenantId int64 `json:"tenantId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PayOrderQuery when eager-loading is set.
	Edges PayOrderEdges `json:"edges"`
}

// PayOrderEdges holds the relations/edges for other nodes in the graph.
type PayOrderEdges struct {
	// User holds the value of the user edge.
	User *SocialUser `json:"user,omitempty"`
	// Channel holds the value of the channel edge.
	Channel *AdChannel `json:"channel,omitempty"`
	// Agreement holds the value of the agreement edge.
	Agreement *AgreementLog `json:"agreement,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PayOrderEdges) UserOrErr() (*SocialUser, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: socialuser.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ChannelOrErr returns the Channel value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PayOrderEdges) ChannelOrErr() (*AdChannel, error) {
	if e.loadedTypes[1] {
		if e.Channel == nil {
			// The edge channel was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: adchannel.Label}
		}
		return e.Channel, nil
	}
	return nil, &NotLoadedError{edge: "channel"}
}

// AgreementOrErr returns the Agreement value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PayOrderEdges) AgreementOrErr() (*AgreementLog, error) {
	if e.loadedTypes[2] {
		if e.Agreement == nil {
			// The edge agreement was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: agreementlog.Label}
		}
		return e.Agreement, nil
	}
	return nil, &NotLoadedError{edge: "agreement"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PayOrder) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case payorder.FieldID, payorder.FieldUserId, payorder.FieldChId, payorder.FieldAgreementId, payorder.FieldState, payorder.FieldPayment, payorder.FieldCoin, payorder.FieldCoupon, payorder.FieldVipType, payorder.FieldTimes, payorder.FieldCreateBy, payorder.FieldUpdateBy, payorder.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case payorder.FieldOrderId, payorder.FieldLastRead, payorder.FieldLastChapter, payorder.FieldPaymentName, payorder.FieldPaymentId, payorder.FieldPayType, payorder.FieldVipDays, payorder.FieldVipName, payorder.FieldOtherOrderId, payorder.FieldRemark:
			values[i] = new(sql.NullString)
		case payorder.FieldPaymentTime, payorder.FieldCloseTime, payorder.FieldCreatedAt, payorder.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PayOrder", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PayOrder fields.
func (po *PayOrder) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case payorder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int64(value.Int64)
		case payorder.FieldOrderId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orderId", values[i])
			} else if value.Valid {
				po.OrderId = value.String
			}
		case payorder.FieldUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				po.UserId = value.Int64
			}
		case payorder.FieldChId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chId", values[i])
			} else if value.Valid {
				po.ChId = value.Int64
			}
		case payorder.FieldAgreementId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field agreementId", values[i])
			} else if value.Valid {
				po.AgreementId = value.Int64
			}
		case payorder.FieldLastRead:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastRead", values[i])
			} else if value.Valid {
				po.LastRead = value.String
			}
		case payorder.FieldLastChapter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastChapter", values[i])
			} else if value.Valid {
				po.LastChapter = value.String
			}
		case payorder.FieldPaymentName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentName", values[i])
			} else if value.Valid {
				po.PaymentName = value.String
			}
		case payorder.FieldPaymentId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentId", values[i])
			} else if value.Valid {
				po.PaymentId = value.String
			}
		case payorder.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				po.State = schema.OrderState(value.Int64)
			}
		case payorder.FieldPayment:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field payment", values[i])
			} else if value.Valid {
				po.Payment = value.Int64
			}
		case payorder.FieldPaymentTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field paymentTime", values[i])
			} else if value.Valid {
				po.PaymentTime = value.Time
			}
		case payorder.FieldCloseTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field closeTime", values[i])
			} else if value.Valid {
				po.CloseTime = value.Time
			}
		case payorder.FieldPayType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payType", values[i])
			} else if value.Valid {
				po.PayType = payorder.PayType(value.String)
			}
		case payorder.FieldCoin:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coin", values[i])
			} else if value.Valid {
				po.Coin = value.Int64
			}
		case payorder.FieldCoupon:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coupon", values[i])
			} else if value.Valid {
				po.Coupon = value.Int64
			}
		case payorder.FieldVipDays:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vipDays", values[i])
			} else if value.Valid {
				po.VipDays = value.String
			}
		case payorder.FieldVipType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vipType", values[i])
			} else if value.Valid {
				po.VipType = value.Int64
			}
		case payorder.FieldVipName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vipName", values[i])
			} else if value.Valid {
				po.VipName = value.String
			}
		case payorder.FieldTimes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field times", values[i])
			} else if value.Valid {
				po.Times = int32(value.Int64)
			}
		case payorder.FieldOtherOrderId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field otherOrderId", values[i])
			} else if value.Valid {
				po.OtherOrderId = value.String
			}
		case payorder.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				po.Remark = value.String
			}
		case payorder.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case payorder.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case payorder.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				po.CreateBy = value.Int64
			}
		case payorder.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				po.UpdateBy = value.Int64
			}
		case payorder.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				po.TenantId = value.Int64
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the PayOrder entity.
func (po *PayOrder) QueryUser() *SocialUserQuery {
	return (&PayOrderClient{config: po.config}).QueryUser(po)
}

// QueryChannel queries the "channel" edge of the PayOrder entity.
func (po *PayOrder) QueryChannel() *AdChannelQuery {
	return (&PayOrderClient{config: po.config}).QueryChannel(po)
}

// QueryAgreement queries the "agreement" edge of the PayOrder entity.
func (po *PayOrder) QueryAgreement() *AgreementLogQuery {
	return (&PayOrderClient{config: po.config}).QueryAgreement(po)
}

// Update returns a builder for updating this PayOrder.
// Note that you need to call PayOrder.Unwrap() before calling this method if this PayOrder
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *PayOrder) Update() *PayOrderUpdateOne {
	return (&PayOrderClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the PayOrder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *PayOrder) Unwrap() *PayOrder {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: PayOrder is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *PayOrder) String() string {
	var builder strings.Builder
	builder.WriteString("PayOrder(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", orderId=")
	builder.WriteString(po.OrderId)
	builder.WriteString(", userId=")
	builder.WriteString(fmt.Sprintf("%v", po.UserId))
	builder.WriteString(", chId=")
	builder.WriteString(fmt.Sprintf("%v", po.ChId))
	builder.WriteString(", agreementId=")
	builder.WriteString(fmt.Sprintf("%v", po.AgreementId))
	builder.WriteString(", lastRead=")
	builder.WriteString(po.LastRead)
	builder.WriteString(", lastChapter=")
	builder.WriteString(po.LastChapter)
	builder.WriteString(", paymentName=")
	builder.WriteString(po.PaymentName)
	builder.WriteString(", paymentId=")
	builder.WriteString(po.PaymentId)
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", po.State))
	builder.WriteString(", payment=")
	builder.WriteString(fmt.Sprintf("%v", po.Payment))
	builder.WriteString(", paymentTime=")
	builder.WriteString(po.PaymentTime.Format(time.ANSIC))
	builder.WriteString(", closeTime=")
	builder.WriteString(po.CloseTime.Format(time.ANSIC))
	builder.WriteString(", payType=")
	builder.WriteString(fmt.Sprintf("%v", po.PayType))
	builder.WriteString(", coin=")
	builder.WriteString(fmt.Sprintf("%v", po.Coin))
	builder.WriteString(", coupon=")
	builder.WriteString(fmt.Sprintf("%v", po.Coupon))
	builder.WriteString(", vipDays=")
	builder.WriteString(po.VipDays)
	builder.WriteString(", vipType=")
	builder.WriteString(fmt.Sprintf("%v", po.VipType))
	builder.WriteString(", vipName=")
	builder.WriteString(po.VipName)
	builder.WriteString(", times=")
	builder.WriteString(fmt.Sprintf("%v", po.Times))
	builder.WriteString(", otherOrderId=")
	builder.WriteString(po.OtherOrderId)
	builder.WriteString(", remark=")
	builder.WriteString(po.Remark)
	builder.WriteString(", createdAt=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", po.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", po.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", po.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// PayOrders is a parsable slice of PayOrder.
type PayOrders []*PayOrder

func (po PayOrders) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
