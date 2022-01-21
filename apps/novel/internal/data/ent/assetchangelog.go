// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/novel/internal/data/ent/assetchangelog"
	"hope/apps/novel/internal/data/ent/socialuser"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// AssetChangeLog is the model entity for the AssetChangeLog schema.
type AssetChangeLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// OrderId holds the value of the "orderId" field.
	// 订单号
	OrderId string `json:"orderId,omitempty"`
	// BalanceId holds the value of the "balanceId" field.
	// 账本ID
	BalanceId int64 `json:"balanceId,omitempty"`
	// EventId holds the value of the "eventId" field.
	// 关联用户事件Id
	EventId int64 `json:"eventId,omitempty"`
	// UserId holds the value of the "userId" field.
	// 用户ID
	UserId int64 `json:"userId,omitempty"`
	// AssetItemId holds the value of the "assetItemId" field.
	// 账本科目
	AssetItemId int32 `json:"assetItemId,omitempty"`
	// Amount holds the value of the "amount" field.
	// 变化金额
	Amount int64 `json:"amount,omitempty"`
	// OldBalance holds the value of the "oldBalance" field.
	// 变化前剩余可用
	OldBalance int64 `json:"oldBalance,omitempty"`
	// NewBalance holds the value of the "newBalance" field.
	// 变化后剩余可用
	NewBalance int64 `json:"newBalance,omitempty"`
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
	// The values are being populated by the AssetChangeLogQuery when eager-loading is set.
	Edges AssetChangeLogEdges `json:"edges"`
}

// AssetChangeLogEdges holds the relations/edges for other nodes in the graph.
type AssetChangeLogEdges struct {
	// User holds the value of the user edge.
	User *SocialUser `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AssetChangeLogEdges) UserOrErr() (*SocialUser, error) {
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
func (*AssetChangeLog) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case assetchangelog.FieldID, assetchangelog.FieldBalanceId, assetchangelog.FieldEventId, assetchangelog.FieldUserId, assetchangelog.FieldAssetItemId, assetchangelog.FieldAmount, assetchangelog.FieldOldBalance, assetchangelog.FieldNewBalance, assetchangelog.FieldCreateBy, assetchangelog.FieldUpdateBy, assetchangelog.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case assetchangelog.FieldOrderId, assetchangelog.FieldRemark:
			values[i] = new(sql.NullString)
		case assetchangelog.FieldCreatedAt, assetchangelog.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AssetChangeLog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AssetChangeLog fields.
func (acl *AssetChangeLog) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case assetchangelog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			acl.ID = int64(value.Int64)
		case assetchangelog.FieldOrderId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orderId", values[i])
			} else if value.Valid {
				acl.OrderId = value.String
			}
		case assetchangelog.FieldBalanceId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field balanceId", values[i])
			} else if value.Valid {
				acl.BalanceId = value.Int64
			}
		case assetchangelog.FieldEventId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field eventId", values[i])
			} else if value.Valid {
				acl.EventId = value.Int64
			}
		case assetchangelog.FieldUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				acl.UserId = value.Int64
			}
		case assetchangelog.FieldAssetItemId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field assetItemId", values[i])
			} else if value.Valid {
				acl.AssetItemId = int32(value.Int64)
			}
		case assetchangelog.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				acl.Amount = value.Int64
			}
		case assetchangelog.FieldOldBalance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field oldBalance", values[i])
			} else if value.Valid {
				acl.OldBalance = value.Int64
			}
		case assetchangelog.FieldNewBalance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field newBalance", values[i])
			} else if value.Valid {
				acl.NewBalance = value.Int64
			}
		case assetchangelog.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				acl.Remark = value.String
			}
		case assetchangelog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				acl.CreatedAt = value.Time
			}
		case assetchangelog.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				acl.UpdatedAt = value.Time
			}
		case assetchangelog.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				acl.CreateBy = value.Int64
			}
		case assetchangelog.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				acl.UpdateBy = value.Int64
			}
		case assetchangelog.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				acl.TenantId = value.Int64
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the AssetChangeLog entity.
func (acl *AssetChangeLog) QueryUser() *SocialUserQuery {
	return (&AssetChangeLogClient{config: acl.config}).QueryUser(acl)
}

// Update returns a builder for updating this AssetChangeLog.
// Note that you need to call AssetChangeLog.Unwrap() before calling this method if this AssetChangeLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (acl *AssetChangeLog) Update() *AssetChangeLogUpdateOne {
	return (&AssetChangeLogClient{config: acl.config}).UpdateOne(acl)
}

// Unwrap unwraps the AssetChangeLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (acl *AssetChangeLog) Unwrap() *AssetChangeLog {
	tx, ok := acl.config.driver.(*txDriver)
	if !ok {
		panic("ent: AssetChangeLog is not a transactional entity")
	}
	acl.config.driver = tx.drv
	return acl
}

// String implements the fmt.Stringer.
func (acl *AssetChangeLog) String() string {
	var builder strings.Builder
	builder.WriteString("AssetChangeLog(")
	builder.WriteString(fmt.Sprintf("id=%v", acl.ID))
	builder.WriteString(", orderId=")
	builder.WriteString(acl.OrderId)
	builder.WriteString(", balanceId=")
	builder.WriteString(fmt.Sprintf("%v", acl.BalanceId))
	builder.WriteString(", eventId=")
	builder.WriteString(fmt.Sprintf("%v", acl.EventId))
	builder.WriteString(", userId=")
	builder.WriteString(fmt.Sprintf("%v", acl.UserId))
	builder.WriteString(", assetItemId=")
	builder.WriteString(fmt.Sprintf("%v", acl.AssetItemId))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", acl.Amount))
	builder.WriteString(", oldBalance=")
	builder.WriteString(fmt.Sprintf("%v", acl.OldBalance))
	builder.WriteString(", newBalance=")
	builder.WriteString(fmt.Sprintf("%v", acl.NewBalance))
	builder.WriteString(", remark=")
	builder.WriteString(acl.Remark)
	builder.WriteString(", createdAt=")
	builder.WriteString(acl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(acl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", acl.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", acl.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", acl.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// AssetChangeLogs is a parsable slice of AssetChangeLog.
type AssetChangeLogs []*AssetChangeLog

func (acl AssetChangeLogs) config(cfg config) {
	for _i := range acl {
		acl[_i].config = cfg
	}
}
