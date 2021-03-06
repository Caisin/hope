package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysJob holds the schema definition for the SysJob entity.
type SysJob struct {
	ent.Schema
}

// Fields of the SysJob.
func (SysJob) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("jobName").Optional().
			Comment(`名称`),
		field.String("jobGroup").Optional().
			Comment(`任务分组`),
		field.Int32("jobType").Optional().
			Comment(`任务类型 1`),
		field.String("cronExpression").Optional().
			Comment(`cron表达式`),
		field.String("invokeTarget").Optional().
			Comment(`调用目标`),
		field.String("args").Optional().
			Comment(`目标参数`),
		field.Int32("execPolicy").Optional().
			Comment(`执行策略`),
		field.Int32("concurrent").Optional().
			Comment(`是否并发`),
		field.Enum("state").Values("Pause", "Run", "Stop").Default("Stop").
			Comment(`状态`),
		field.Int32("entryId").Optional().
			Comment(`job启动时返回的id`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysJob.
func (SysJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("logs", SysJobLog.Type).Comment("执行日志"),
	}
}
