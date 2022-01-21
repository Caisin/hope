// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/admin/internal/data/ent/sysjob"
	"hope/apps/admin/internal/data/ent/sysjoblog"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// SysJobLog is the model entity for the SysJobLog schema.
type SysJobLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// JobId holds the value of the "jobId" field.
	// 编码
	JobId int32 `json:"jobId,omitempty"`
	// JobName holds the value of the "jobName" field.
	// 名称
	JobName string `json:"jobName,omitempty"`
	// EntryId holds the value of the "entryId" field.
	// job启动时返回的id
	EntryId int32 `json:"entryId,omitempty"`
	// Status holds the value of the "status" field.
	// 状态
	Status bool `json:"status,omitempty"`
	// Duration holds the value of the "duration" field.
	// 耗时
	Duration time.Duration `json:"duration,omitempty"`
	// Info holds the value of the "info" field.
	// 执行详情,错误信息
	Info string `json:"info,omitempty"`
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
	// The values are being populated by the SysJobLogQuery when eager-loading is set.
	Edges        SysJobLogEdges `json:"edges"`
	sys_job_logs *int64
}

// SysJobLogEdges holds the relations/edges for other nodes in the graph.
type SysJobLogEdges struct {
	// Job holds the value of the job edge.
	Job *SysJob `json:"job,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// JobOrErr returns the Job value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysJobLogEdges) JobOrErr() (*SysJob, error) {
	if e.loadedTypes[0] {
		if e.Job == nil {
			// The edge job was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: sysjob.Label}
		}
		return e.Job, nil
	}
	return nil, &NotLoadedError{edge: "job"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysJobLog) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysjoblog.FieldStatus:
			values[i] = new(sql.NullBool)
		case sysjoblog.FieldID, sysjoblog.FieldJobId, sysjoblog.FieldEntryId, sysjoblog.FieldDuration, sysjoblog.FieldCreateBy, sysjoblog.FieldUpdateBy, sysjoblog.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case sysjoblog.FieldJobName, sysjoblog.FieldInfo:
			values[i] = new(sql.NullString)
		case sysjoblog.FieldCreatedAt, sysjoblog.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case sysjoblog.ForeignKeys[0]: // sys_job_logs
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SysJobLog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysJobLog fields.
func (sjl *SysJobLog) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysjoblog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sjl.ID = int64(value.Int64)
		case sysjoblog.FieldJobId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field jobId", values[i])
			} else if value.Valid {
				sjl.JobId = int32(value.Int64)
			}
		case sysjoblog.FieldJobName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field jobName", values[i])
			} else if value.Valid {
				sjl.JobName = value.String
			}
		case sysjoblog.FieldEntryId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field entryId", values[i])
			} else if value.Valid {
				sjl.EntryId = int32(value.Int64)
			}
		case sysjoblog.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sjl.Status = value.Bool
			}
		case sysjoblog.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				sjl.Duration = time.Duration(value.Int64)
			}
		case sysjoblog.FieldInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field info", values[i])
			} else if value.Valid {
				sjl.Info = value.String
			}
		case sysjoblog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				sjl.CreatedAt = value.Time
			}
		case sysjoblog.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				sjl.UpdatedAt = value.Time
			}
		case sysjoblog.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				sjl.CreateBy = value.Int64
			}
		case sysjoblog.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				sjl.UpdateBy = value.Int64
			}
		case sysjoblog.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				sjl.TenantId = value.Int64
			}
		case sysjoblog.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field sys_job_logs", value)
			} else if value.Valid {
				sjl.sys_job_logs = new(int64)
				*sjl.sys_job_logs = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryJob queries the "job" edge of the SysJobLog entity.
func (sjl *SysJobLog) QueryJob() *SysJobQuery {
	return (&SysJobLogClient{config: sjl.config}).QueryJob(sjl)
}

// Update returns a builder for updating this SysJobLog.
// Note that you need to call SysJobLog.Unwrap() before calling this method if this SysJobLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (sjl *SysJobLog) Update() *SysJobLogUpdateOne {
	return (&SysJobLogClient{config: sjl.config}).UpdateOne(sjl)
}

// Unwrap unwraps the SysJobLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sjl *SysJobLog) Unwrap() *SysJobLog {
	tx, ok := sjl.config.driver.(*txDriver)
	if !ok {
		panic("ent: SysJobLog is not a transactional entity")
	}
	sjl.config.driver = tx.drv
	return sjl
}

// String implements the fmt.Stringer.
func (sjl *SysJobLog) String() string {
	var builder strings.Builder
	builder.WriteString("SysJobLog(")
	builder.WriteString(fmt.Sprintf("id=%v", sjl.ID))
	builder.WriteString(", jobId=")
	builder.WriteString(fmt.Sprintf("%v", sjl.JobId))
	builder.WriteString(", jobName=")
	builder.WriteString(sjl.JobName)
	builder.WriteString(", entryId=")
	builder.WriteString(fmt.Sprintf("%v", sjl.EntryId))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", sjl.Status))
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", sjl.Duration))
	builder.WriteString(", info=")
	builder.WriteString(sjl.Info)
	builder.WriteString(", createdAt=")
	builder.WriteString(sjl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(sjl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", sjl.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", sjl.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", sjl.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// SysJobLogs is a parsable slice of SysJobLog.
type SysJobLogs []*SysJobLog

func (sjl SysJobLogs) config(cfg config) {
	for _i := range sjl {
		sjl[_i].config = cfg
	}
}