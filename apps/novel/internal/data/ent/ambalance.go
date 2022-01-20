// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/novel/internal/data/ent/ambalance"
	"hope/apps/novel/internal/data/ent/socialuser"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// AmBalance is the model entity for the AmBalance schema.
type AmBalance struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// OrderId holds the value of the "orderId" field.
	// 订单号
	OrderId string `json:"orderId,omitempty"`
	// EventId holds the value of the "eventId" field.
	// 关联用户事件Id
	EventId int64 `json:"eventId,omitempty"`
	// CashTag holds the value of the "cashTag" field.
	// 现金标识,0优惠券 1书币
	CashTag int32 `json:"cashTag,omitempty"`
	// AssetItemId holds the value of the "assetItemId" field.
	// 账本科目
	AssetItemId int32 `json:"assetItemId,omitempty"`
	// Amount holds the value of the "amount" field.
	// 原始金额
	Amount int64 `json:"amount,omitempty"`
	// Balance holds the value of the "balance" field.
	// 剩余可用
	Balance int64 `json:"balance,omitempty"`
	// Remark holds the value of the "remark" field.
	// 备注
	Remark string `json:"remark,omitempty"`
	// EffectTime holds the value of the "effectTime" field.
	// 生效时间
	EffectTime time.Time `json:"effectTime,omitempty"`
	// ExpiredTime holds the value of the "expiredTime" field.
	// 失效时间
	ExpiredTime time.Time `json:"expiredTime,omitempty"`
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
	// The values are being populated by the AmBalanceQuery when eager-loading is set.
	Edges                AmBalanceEdges `json:"edges"`
	social_user_balances *int64
}

// AmBalanceEdges holds the relations/edges for other nodes in the graph.
type AmBalanceEdges struct {
	// User holds the value of the user edge.
	User *SocialUser `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AmBalanceEdges) UserOrErr() (*SocialUser, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*AmBalance) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ambalance.FieldID, ambalance.FieldEventId, ambalance.FieldCashTag, ambalance.FieldAssetItemId, ambalance.FieldAmount, ambalance.FieldBalance, ambalance.FieldCreateBy, ambalance.FieldUpdateBy, ambalance.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case ambalance.FieldOrderId, ambalance.FieldRemark:
			values[i] = new(sql.NullString)
		case ambalance.FieldEffectTime, ambalance.FieldExpiredTime, ambalance.FieldCreatedAt, ambalance.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case ambalance.ForeignKeys[0]: // social_user_balances
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AmBalance", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AmBalance fields.
func (ab *AmBalance) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ambalance.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ab.ID = int64(value.Int64)
		case ambalance.FieldOrderId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orderId", values[i])
			} else if value.Valid {
				ab.OrderId = value.String
			}
		case ambalance.FieldEventId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field eventId", values[i])
			} else if value.Valid {
				ab.EventId = value.Int64
			}
		case ambalance.FieldCashTag:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cashTag", values[i])
			} else if value.Valid {
				ab.CashTag = int32(value.Int64)
			}
		case ambalance.FieldAssetItemId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field assetItemId", values[i])
			} else if value.Valid {
				ab.AssetItemId = int32(value.Int64)
			}
		case ambalance.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				ab.Amount = value.Int64
			}
		case ambalance.FieldBalance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				ab.Balance = value.Int64
			}
		case ambalance.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				ab.Remark = value.String
			}
		case ambalance.FieldEffectTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field effectTime", values[i])
			} else if value.Valid {
				ab.EffectTime = value.Time
			}
		case ambalance.FieldExpiredTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiredTime", values[i])
			} else if value.Valid {
				ab.ExpiredTime = value.Time
			}
		case ambalance.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				ab.CreatedAt = value.Time
			}
		case ambalance.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				ab.UpdatedAt = value.Time
			}
		case ambalance.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				ab.CreateBy = value.Int64
			}
		case ambalance.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				ab.UpdateBy = value.Int64
			}
		case ambalance.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				ab.TenantId = value.Int64
			}
		case ambalance.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field social_user_balances", value)
			} else if value.Valid {
				ab.social_user_balances = new(int64)
				*ab.social_user_balances = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the AmBalance entity.
func (ab *AmBalance) QueryUser() *SocialUserQuery {
	return (&AmBalanceClient{config: ab.config}).QueryUser(ab)
}

// Update returns a builder for updating this AmBalance.
// Note that you need to call AmBalance.Unwrap() before calling this method if this AmBalance
// was returned from a transaction, and the transaction was committed or rolled back.
func (ab *AmBalance) Update() *AmBalanceUpdateOne {
	return (&AmBalanceClient{config: ab.config}).UpdateOne(ab)
}

// Unwrap unwraps the AmBalance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ab *AmBalance) Unwrap() *AmBalance {
	tx, ok := ab.config.driver.(*txDriver)
	if !ok {
		panic("ent: AmBalance is not a transactional entity")
	}
	ab.config.driver = tx.drv
	return ab
}

// String implements the fmt.Stringer.
func (ab *AmBalance) String() string {
	var builder strings.Builder
	builder.WriteString("AmBalance(")
	builder.WriteString(fmt.Sprintf("id=%v", ab.ID))
	builder.WriteString(", orderId=")
	builder.WriteString(ab.OrderId)
	builder.WriteString(", eventId=")
	builder.WriteString(fmt.Sprintf("%v", ab.EventId))
	builder.WriteString(", cashTag=")
	builder.WriteString(fmt.Sprintf("%v", ab.CashTag))
	builder.WriteString(", assetItemId=")
	builder.WriteString(fmt.Sprintf("%v", ab.AssetItemId))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", ab.Amount))
	builder.WriteString(", balance=")
	builder.WriteString(fmt.Sprintf("%v", ab.Balance))
	builder.WriteString(", remark=")
	builder.WriteString(ab.Remark)
	builder.WriteString(", effectTime=")
	builder.WriteString(ab.EffectTime.Format(time.ANSIC))
	builder.WriteString(", expiredTime=")
	builder.WriteString(ab.ExpiredTime.Format(time.ANSIC))
	builder.WriteString(", createdAt=")
	builder.WriteString(ab.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(ab.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", ab.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", ab.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", ab.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// AmBalances is a parsable slice of AmBalance.
type AmBalances []*AmBalance

func (ab AmBalances) config(cfg config) {
	for _i := range ab {
		ab[_i].config = cfg
	}
}
