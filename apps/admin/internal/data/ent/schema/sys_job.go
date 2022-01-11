package schema

import (
	"entgo.io/ent"
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
		field.Int("jobId").
			Comment(`编码`),
		field.String("jobName").Optional().
			Comment(`名称`),
		field.String("jobGroup").Optional().
			Comment(`任务分组`),
		field.Int("jobType").Optional().
			Comment(`任务类型 1`),
		field.String("cronExpression").Optional().
			Comment(`cron表达式`),
		field.String("invokeTarget").Optional().
			Comment(`调用目标`),
		field.String("args").Optional().
			Comment(`目标参数`),
		field.Int("misfirePolicy").Optional().
			Comment(`执行策略`),
		field.Int("concurrent").Optional().
			Comment(`是否并发`),
		field.Int("status").Optional().
			Comment(`状态`),
		field.Int("entry_id").Optional().
			Comment(`job启动时返回的id`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the SysJob.
func (SysJob) Edges() []ent.Edge {
	return []ent.Edge{}
}
