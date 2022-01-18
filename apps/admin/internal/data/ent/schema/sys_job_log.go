package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
	"time"
)

// SysJobLog holds the schema definition for the SysJobLog entity.
type SysJobLog struct {
	ent.Schema
}

// Fields of the SysJobLog.
func (SysJobLog) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int32("jobId").Optional().
			Comment(`编码`),
		field.String("jobName").Optional().
			Comment(`名称`),
		field.Int32("entry_id").Optional().
			Comment(`job启动时返回的id`),
		field.Bool("status").Optional().
			Comment(`状态`),
		field.Int64("duration").GoType(time.Duration(0)).Optional().
			Comment(`耗时`),
		field.String("info").Optional().
			Comment(`执行详情,错误信息`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysJobLog.
func (SysJobLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("job", SysJob.Type).Ref("logs").Unique(),
	}
}
