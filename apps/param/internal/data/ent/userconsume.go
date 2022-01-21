// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/param/internal/data/ent/userconsume"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// UserConsume is the model entity for the UserConsume schema.
type UserConsume struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// NovelId holds the value of the "novelId" field.
	// 用户ID
	NovelId int64 `json:"novelId,omitempty"`
	// Coin holds the value of the "coin" field.
	// 书币
	Coin int64 `json:"coin,omitempty"`
	// Coupon holds the value of the "coupon" field.
	// 书券
	Coupon int64 `json:"coupon,omitempty"`
	// Discount holds the value of the "discount" field.
	// VIP折扣金额
	Discount int64 `json:"discount,omitempty"`
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
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserConsume) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userconsume.FieldID, userconsume.FieldNovelId, userconsume.FieldCoin, userconsume.FieldCoupon, userconsume.FieldDiscount, userconsume.FieldCreateBy, userconsume.FieldUpdateBy, userconsume.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case userconsume.FieldCreatedAt, userconsume.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserConsume", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserConsume fields.
func (uc *UserConsume) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userconsume.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uc.ID = int64(value.Int64)
		case userconsume.FieldNovelId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field novelId", values[i])
			} else if value.Valid {
				uc.NovelId = value.Int64
			}
		case userconsume.FieldCoin:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coin", values[i])
			} else if value.Valid {
				uc.Coin = value.Int64
			}
		case userconsume.FieldCoupon:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coupon", values[i])
			} else if value.Valid {
				uc.Coupon = value.Int64
			}
		case userconsume.FieldDiscount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field discount", values[i])
			} else if value.Valid {
				uc.Discount = value.Int64
			}
		case userconsume.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				uc.CreatedAt = value.Time
			}
		case userconsume.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				uc.UpdatedAt = value.Time
			}
		case userconsume.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				uc.CreateBy = value.Int64
			}
		case userconsume.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				uc.UpdateBy = value.Int64
			}
		case userconsume.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				uc.TenantId = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserConsume.
// Note that you need to call UserConsume.Unwrap() before calling this method if this UserConsume
// was returned from a transaction, and the transaction was committed or rolled back.
func (uc *UserConsume) Update() *UserConsumeUpdateOne {
	return (&UserConsumeClient{config: uc.config}).UpdateOne(uc)
}

// Unwrap unwraps the UserConsume entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uc *UserConsume) Unwrap() *UserConsume {
	tx, ok := uc.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserConsume is not a transactional entity")
	}
	uc.config.driver = tx.drv
	return uc
}

// String implements the fmt.Stringer.
func (uc *UserConsume) String() string {
	var builder strings.Builder
	builder.WriteString("UserConsume(")
	builder.WriteString(fmt.Sprintf("id=%v", uc.ID))
	builder.WriteString(", novelId=")
	builder.WriteString(fmt.Sprintf("%v", uc.NovelId))
	builder.WriteString(", coin=")
	builder.WriteString(fmt.Sprintf("%v", uc.Coin))
	builder.WriteString(", coupon=")
	builder.WriteString(fmt.Sprintf("%v", uc.Coupon))
	builder.WriteString(", discount=")
	builder.WriteString(fmt.Sprintf("%v", uc.Discount))
	builder.WriteString(", createdAt=")
	builder.WriteString(uc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(uc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", uc.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", uc.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", uc.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// UserConsumes is a parsable slice of UserConsume.
type UserConsumes []*UserConsume

func (uc UserConsumes) config(cfg config) {
	for _i := range uc {
		uc[_i].config = cfg
	}
}