// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/param/internal/data/ent/pageconfig"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// PageConfig is the model entity for the PageConfig schema.
type PageConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// PageCode holds the value of the "pageCode" field.
	// 页面编码
	PageCode string `json:"pageCode,omitempty"`
	// PageName holds the value of the "pageName" field.
	// 页面名称
	PageName string `json:"pageName,omitempty"`
	// GroupCodes holds the value of the "groupCodes" field.
	// 分组编码集,逗号分割
	GroupCodes string `json:"groupCodes,omitempty"`
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
func (*PageConfig) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pageconfig.FieldID, pageconfig.FieldCreateBy, pageconfig.FieldUpdateBy, pageconfig.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case pageconfig.FieldPageCode, pageconfig.FieldPageName, pageconfig.FieldGroupCodes:
			values[i] = new(sql.NullString)
		case pageconfig.FieldCreatedAt, pageconfig.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PageConfig", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PageConfig fields.
func (pc *PageConfig) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pageconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = int64(value.Int64)
		case pageconfig.FieldPageCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pageCode", values[i])
			} else if value.Valid {
				pc.PageCode = value.String
			}
		case pageconfig.FieldPageName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pageName", values[i])
			} else if value.Valid {
				pc.PageName = value.String
			}
		case pageconfig.FieldGroupCodes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field groupCodes", values[i])
			} else if value.Valid {
				pc.GroupCodes = value.String
			}
		case pageconfig.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				pc.CreatedAt = value.Time
			}
		case pageconfig.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				pc.UpdatedAt = value.Time
			}
		case pageconfig.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				pc.CreateBy = value.Int64
			}
		case pageconfig.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				pc.UpdateBy = value.Int64
			}
		case pageconfig.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				pc.TenantId = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PageConfig.
// Note that you need to call PageConfig.Unwrap() before calling this method if this PageConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *PageConfig) Update() *PageConfigUpdateOne {
	return (&PageConfigClient{config: pc.config}).UpdateOne(pc)
}

// Unwrap unwraps the PageConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *PageConfig) Unwrap() *PageConfig {
	tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PageConfig is not a transactional entity")
	}
	pc.config.driver = tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *PageConfig) String() string {
	var builder strings.Builder
	builder.WriteString("PageConfig(")
	builder.WriteString(fmt.Sprintf("id=%v", pc.ID))
	builder.WriteString(", pageCode=")
	builder.WriteString(pc.PageCode)
	builder.WriteString(", pageName=")
	builder.WriteString(pc.PageName)
	builder.WriteString(", groupCodes=")
	builder.WriteString(pc.GroupCodes)
	builder.WriteString(", createdAt=")
	builder.WriteString(pc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(pc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", pc.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", pc.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", pc.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// PageConfigs is a parsable slice of PageConfig.
type PageConfigs []*PageConfig

func (pc PageConfigs) config(cfg config) {
	for _i := range pc {
		pc[_i].config = cfg
	}
}
