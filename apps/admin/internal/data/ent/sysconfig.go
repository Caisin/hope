// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/admin/internal/data/ent/sysconfig"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// SysConfig is the model entity for the SysConfig schema.
type SysConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// ConfigName holds the value of the "configName" field.
	// 配置名称
	ConfigName string `json:"configName,omitempty"`
	// ConfigKey holds the value of the "configKey" field.
	// 配置Key
	ConfigKey string `json:"configKey,omitempty"`
	// ConfigValue holds the value of the "configValue" field.
	// 配置值
	ConfigValue string `json:"configValue,omitempty"`
	// ConfigType holds the value of the "configType" field.
	// 配置类型
	ConfigType string `json:"configType,omitempty"`
	// IsFrontend holds the value of the "isFrontend" field.
	// 是否前台
	IsFrontend int32 `json:"isFrontend,omitempty"`
	// State holds the value of the "state" field.
	// 状态:U:使用状态,E:失效状态
	State sysconfig.State `json:"state,omitempty"`
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
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysConfig) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysconfig.FieldID, sysconfig.FieldIsFrontend, sysconfig.FieldCreateBy, sysconfig.FieldUpdateBy, sysconfig.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case sysconfig.FieldConfigName, sysconfig.FieldConfigKey, sysconfig.FieldConfigValue, sysconfig.FieldConfigType, sysconfig.FieldState, sysconfig.FieldRemark:
			values[i] = new(sql.NullString)
		case sysconfig.FieldCreatedAt, sysconfig.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysConfig", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysConfig fields.
func (sc *SysConfig) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = int64(value.Int64)
		case sysconfig.FieldConfigName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field configName", values[i])
			} else if value.Valid {
				sc.ConfigName = value.String
			}
		case sysconfig.FieldConfigKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field configKey", values[i])
			} else if value.Valid {
				sc.ConfigKey = value.String
			}
		case sysconfig.FieldConfigValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field configValue", values[i])
			} else if value.Valid {
				sc.ConfigValue = value.String
			}
		case sysconfig.FieldConfigType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field configType", values[i])
			} else if value.Valid {
				sc.ConfigType = value.String
			}
		case sysconfig.FieldIsFrontend:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field isFrontend", values[i])
			} else if value.Valid {
				sc.IsFrontend = int32(value.Int64)
			}
		case sysconfig.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				sc.State = sysconfig.State(value.String)
			}
		case sysconfig.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				sc.Remark = value.String
			}
		case sysconfig.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				sc.CreatedAt = value.Time
			}
		case sysconfig.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				sc.UpdatedAt = value.Time
			}
		case sysconfig.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				sc.CreateBy = value.Int64
			}
		case sysconfig.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				sc.UpdateBy = value.Int64
			}
		case sysconfig.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				sc.TenantId = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SysConfig.
// Note that you need to call SysConfig.Unwrap() before calling this method if this SysConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *SysConfig) Update() *SysConfigUpdateOne {
	return (&SysConfigClient{config: sc.config}).UpdateOne(sc)
}

// Unwrap unwraps the SysConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *SysConfig) Unwrap() *SysConfig {
	tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysConfig is not a transactional entity")
	}
	sc.config.driver = tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *SysConfig) String() string {
	var builder strings.Builder
	builder.WriteString("SysConfig(")
	builder.WriteString(fmt.Sprintf("id=%v", sc.ID))
	builder.WriteString(", configName=")
	builder.WriteString(sc.ConfigName)
	builder.WriteString(", configKey=")
	builder.WriteString(sc.ConfigKey)
	builder.WriteString(", configValue=")
	builder.WriteString(sc.ConfigValue)
	builder.WriteString(", configType=")
	builder.WriteString(sc.ConfigType)
	builder.WriteString(", isFrontend=")
	builder.WriteString(fmt.Sprintf("%v", sc.IsFrontend))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", sc.State))
	builder.WriteString(", remark=")
	builder.WriteString(sc.Remark)
	builder.WriteString(", createdAt=")
	builder.WriteString(sc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(sc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", sc.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", sc.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", sc.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// SysConfigs is a parsable slice of SysConfig.
type SysConfigs []*SysConfig

func (sc SysConfigs) config(cfg config) {
	for _i := range sc {
		sc[_i].config = cfg
	}
}
