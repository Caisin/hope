// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/param/internal/data/ent/resourcegroup"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ResourceGroup is the model entity for the ResourceGroup schema.
type ResourceGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	// Name
	Name string `json:"name,omitempty"`
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
func (*ResourceGroup) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case resourcegroup.FieldID, resourcegroup.FieldCreateBy, resourcegroup.FieldUpdateBy, resourcegroup.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case resourcegroup.FieldName:
			values[i] = new(sql.NullString)
		case resourcegroup.FieldCreatedAt, resourcegroup.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ResourceGroup", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResourceGroup fields.
func (rg *ResourceGroup) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resourcegroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rg.ID = int64(value.Int64)
		case resourcegroup.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rg.Name = value.String
			}
		case resourcegroup.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				rg.CreatedAt = value.Time
			}
		case resourcegroup.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				rg.UpdatedAt = value.Time
			}
		case resourcegroup.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				rg.CreateBy = value.Int64
			}
		case resourcegroup.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				rg.UpdateBy = value.Int64
			}
		case resourcegroup.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				rg.TenantId = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ResourceGroup.
// Note that you need to call ResourceGroup.Unwrap() before calling this method if this ResourceGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (rg *ResourceGroup) Update() *ResourceGroupUpdateOne {
	return (&ResourceGroupClient{config: rg.config}).UpdateOne(rg)
}

// Unwrap unwraps the ResourceGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rg *ResourceGroup) Unwrap() *ResourceGroup {
	tx, ok := rg.config.driver.(*txDriver)
	if !ok {
		panic("ent: ResourceGroup is not a transactional entity")
	}
	rg.config.driver = tx.drv
	return rg
}

// String implements the fmt.Stringer.
func (rg *ResourceGroup) String() string {
	var builder strings.Builder
	builder.WriteString("ResourceGroup(")
	builder.WriteString(fmt.Sprintf("id=%v", rg.ID))
	builder.WriteString(", name=")
	builder.WriteString(rg.Name)
	builder.WriteString(", createdAt=")
	builder.WriteString(rg.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(rg.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", rg.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", rg.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", rg.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// ResourceGroups is a parsable slice of ResourceGroup.
type ResourceGroups []*ResourceGroup

func (rg ResourceGroups) config(cfg config) {
	for _i := range rg {
		rg[_i].config = cfg
	}
}
