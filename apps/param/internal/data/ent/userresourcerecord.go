// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/param/internal/data/ent/userresourcerecord"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// UserResourceRecord is the model entity for the UserResourceRecord schema.
type UserResourceRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserId holds the value of the "userId" field.
	// 用户ID
	UserId int64 `json:"userId,omitempty"`
	// ResId holds the value of the "resId" field.
	// 资源ID
	ResId int64 `json:"resId,omitempty"`
	// Def holds the value of the "def" field.
	// 是否默认
	Def bool `json:"def,omitempty"`
	// Name holds the value of the "name" field.
	// 头像名称
	Name string `json:"name,omitempty"`
	// URL holds the value of the "url" field.
	// 资源地址
	URL string `json:"url,omitempty"`
	// ResType holds the value of the "resType" field.
	// 资源类型,avatar
	ResType string `json:"resType,omitempty"`
	// Remark holds the value of the "remark" field.
	// 备注
	Remark string `json:"remark,omitempty"`
	// State holds the value of the "state" field.
	// 使用状态,1
	State int32 `json:"state,omitempty"`
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
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserResourceRecord) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userresourcerecord.FieldDef:
			values[i] = new(sql.NullBool)
		case userresourcerecord.FieldID, userresourcerecord.FieldUserId, userresourcerecord.FieldResId, userresourcerecord.FieldState, userresourcerecord.FieldCreateBy, userresourcerecord.FieldUpdateBy, userresourcerecord.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case userresourcerecord.FieldName, userresourcerecord.FieldURL, userresourcerecord.FieldResType, userresourcerecord.FieldRemark:
			values[i] = new(sql.NullString)
		case userresourcerecord.FieldEffectTime, userresourcerecord.FieldExpiredTime, userresourcerecord.FieldCreatedAt, userresourcerecord.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserResourceRecord", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserResourceRecord fields.
func (urr *UserResourceRecord) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userresourcerecord.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			urr.ID = int64(value.Int64)
		case userresourcerecord.FieldUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				urr.UserId = value.Int64
			}
		case userresourcerecord.FieldResId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field resId", values[i])
			} else if value.Valid {
				urr.ResId = value.Int64
			}
		case userresourcerecord.FieldDef:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field def", values[i])
			} else if value.Valid {
				urr.Def = value.Bool
			}
		case userresourcerecord.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				urr.Name = value.String
			}
		case userresourcerecord.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				urr.URL = value.String
			}
		case userresourcerecord.FieldResType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resType", values[i])
			} else if value.Valid {
				urr.ResType = value.String
			}
		case userresourcerecord.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				urr.Remark = value.String
			}
		case userresourcerecord.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				urr.State = int32(value.Int64)
			}
		case userresourcerecord.FieldEffectTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field effectTime", values[i])
			} else if value.Valid {
				urr.EffectTime = value.Time
			}
		case userresourcerecord.FieldExpiredTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiredTime", values[i])
			} else if value.Valid {
				urr.ExpiredTime = value.Time
			}
		case userresourcerecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				urr.CreatedAt = value.Time
			}
		case userresourcerecord.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				urr.UpdatedAt = value.Time
			}
		case userresourcerecord.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				urr.CreateBy = value.Int64
			}
		case userresourcerecord.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				urr.UpdateBy = value.Int64
			}
		case userresourcerecord.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				urr.TenantId = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserResourceRecord.
// Note that you need to call UserResourceRecord.Unwrap() before calling this method if this UserResourceRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (urr *UserResourceRecord) Update() *UserResourceRecordUpdateOne {
	return (&UserResourceRecordClient{config: urr.config}).UpdateOne(urr)
}

// Unwrap unwraps the UserResourceRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (urr *UserResourceRecord) Unwrap() *UserResourceRecord {
	tx, ok := urr.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserResourceRecord is not a transactional entity")
	}
	urr.config.driver = tx.drv
	return urr
}

// String implements the fmt.Stringer.
func (urr *UserResourceRecord) String() string {
	var builder strings.Builder
	builder.WriteString("UserResourceRecord(")
	builder.WriteString(fmt.Sprintf("id=%v", urr.ID))
	builder.WriteString(", userId=")
	builder.WriteString(fmt.Sprintf("%v", urr.UserId))
	builder.WriteString(", resId=")
	builder.WriteString(fmt.Sprintf("%v", urr.ResId))
	builder.WriteString(", def=")
	builder.WriteString(fmt.Sprintf("%v", urr.Def))
	builder.WriteString(", name=")
	builder.WriteString(urr.Name)
	builder.WriteString(", url=")
	builder.WriteString(urr.URL)
	builder.WriteString(", resType=")
	builder.WriteString(urr.ResType)
	builder.WriteString(", remark=")
	builder.WriteString(urr.Remark)
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", urr.State))
	builder.WriteString(", effectTime=")
	builder.WriteString(urr.EffectTime.Format(time.ANSIC))
	builder.WriteString(", expiredTime=")
	builder.WriteString(urr.ExpiredTime.Format(time.ANSIC))
	builder.WriteString(", createdAt=")
	builder.WriteString(urr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(urr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", urr.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", urr.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", urr.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// UserResourceRecords is a parsable slice of UserResourceRecord.
type UserResourceRecords []*UserResourceRecord

func (urr UserResourceRecords) config(cfg config) {
	for _i := range urr {
		urr[_i].config = cfg
	}
}
