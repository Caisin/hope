// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/novel/internal/data/ent/vipuser"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// VipUser is the model entity for the VipUser schema.
type VipUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// VipType holds the value of the "vipType" field.
	// vip类型
	VipType int64 `json:"vipType,omitempty"`
	// SvipType holds the value of the "svipType" field.
	// svip类型
	SvipType int64 `json:"svipType,omitempty"`
	// SvipEffectTime holds the value of the "svipEffectTime" field.
	// 超级VIP生效时间
	SvipEffectTime time.Time `json:"svipEffectTime,omitempty"`
	// SvipExpiredTime holds the value of the "svipExpiredTime" field.
	// 超级VIP失效时间
	SvipExpiredTime time.Time `json:"svipExpiredTime,omitempty"`
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
	// The values are being populated by the VipUserQuery when eager-loading is set.
	Edges VipUserEdges `json:"edges"`
}

// VipUserEdges holds the relations/edges for other nodes in the graph.
type VipUserEdges struct {
	// User holds the value of the user edge.
	User []*SocialUser `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e VipUserEdges) UserOrErr() ([]*SocialUser, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VipUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case vipuser.FieldID, vipuser.FieldVipType, vipuser.FieldSvipType, vipuser.FieldCreateBy, vipuser.FieldUpdateBy, vipuser.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case vipuser.FieldRemark:
			values[i] = new(sql.NullString)
		case vipuser.FieldSvipEffectTime, vipuser.FieldSvipExpiredTime, vipuser.FieldEffectTime, vipuser.FieldExpiredTime, vipuser.FieldCreatedAt, vipuser.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type VipUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VipUser fields.
func (vu *VipUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vipuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vu.ID = int64(value.Int64)
		case vipuser.FieldVipType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vipType", values[i])
			} else if value.Valid {
				vu.VipType = value.Int64
			}
		case vipuser.FieldSvipType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field svipType", values[i])
			} else if value.Valid {
				vu.SvipType = value.Int64
			}
		case vipuser.FieldSvipEffectTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field svipEffectTime", values[i])
			} else if value.Valid {
				vu.SvipEffectTime = value.Time
			}
		case vipuser.FieldSvipExpiredTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field svipExpiredTime", values[i])
			} else if value.Valid {
				vu.SvipExpiredTime = value.Time
			}
		case vipuser.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				vu.Remark = value.String
			}
		case vipuser.FieldEffectTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field effectTime", values[i])
			} else if value.Valid {
				vu.EffectTime = value.Time
			}
		case vipuser.FieldExpiredTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiredTime", values[i])
			} else if value.Valid {
				vu.ExpiredTime = value.Time
			}
		case vipuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				vu.CreatedAt = value.Time
			}
		case vipuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				vu.UpdatedAt = value.Time
			}
		case vipuser.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				vu.CreateBy = value.Int64
			}
		case vipuser.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				vu.UpdateBy = value.Int64
			}
		case vipuser.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				vu.TenantId = value.Int64
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the VipUser entity.
func (vu *VipUser) QueryUser() *SocialUserQuery {
	return (&VipUserClient{config: vu.config}).QueryUser(vu)
}

// Update returns a builder for updating this VipUser.
// Note that you need to call VipUser.Unwrap() before calling this method if this VipUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (vu *VipUser) Update() *VipUserUpdateOne {
	return (&VipUserClient{config: vu.config}).UpdateOne(vu)
}

// Unwrap unwraps the VipUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vu *VipUser) Unwrap() *VipUser {
	tx, ok := vu.config.driver.(*txDriver)
	if !ok {
		panic("ent: VipUser is not a transactional entity")
	}
	vu.config.driver = tx.drv
	return vu
}

// String implements the fmt.Stringer.
func (vu *VipUser) String() string {
	var builder strings.Builder
	builder.WriteString("VipUser(")
	builder.WriteString(fmt.Sprintf("id=%v", vu.ID))
	builder.WriteString(", vipType=")
	builder.WriteString(fmt.Sprintf("%v", vu.VipType))
	builder.WriteString(", svipType=")
	builder.WriteString(fmt.Sprintf("%v", vu.SvipType))
	builder.WriteString(", svipEffectTime=")
	builder.WriteString(vu.SvipEffectTime.Format(time.ANSIC))
	builder.WriteString(", svipExpiredTime=")
	builder.WriteString(vu.SvipExpiredTime.Format(time.ANSIC))
	builder.WriteString(", remark=")
	builder.WriteString(vu.Remark)
	builder.WriteString(", effectTime=")
	builder.WriteString(vu.EffectTime.Format(time.ANSIC))
	builder.WriteString(", expiredTime=")
	builder.WriteString(vu.ExpiredTime.Format(time.ANSIC))
	builder.WriteString(", createdAt=")
	builder.WriteString(vu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(vu.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", vu.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", vu.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", vu.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// VipUsers is a parsable slice of VipUser.
type VipUsers []*VipUser

func (vu VipUsers) config(cfg config) {
	for _i := range vu {
		vu[_i].config = cfg
	}
}