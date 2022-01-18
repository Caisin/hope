package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("taskName").Optional().
			Comment(`任务名称`),
		field.String("taskGroup").Optional().
			Comment(`任务分组`),
		field.String("unit").Optional().
			Comment(`单位`),
		field.String("topic").Optional().
			Comment(`订阅主题`),
		field.String("function").Optional().
			Comment(`处理方法`),
		field.String("taskCode").Optional().
			Comment(`任务类型`),
		field.Int64("preTask").Optional().
			Comment(`前置任务ID`),
		field.Int64("novelId").Optional().
			Comment(`点击跳转书号`),
		field.String("cycleType").Optional().
			Comment(`循环类型,once`),
		field.String("remark").Optional().
			Comment(`描述信息`),
		field.Int64("amount").Optional().
			Comment(`奖励金额`),
		field.Int64("reward").Optional().
			Comment(`额外奖励`),
		field.Int32("amountItem").Optional().
			Comment(`奖励资金科目`),
		field.Int32("rewardItem").Optional().
			Comment(`额外奖励资金科目`),
		field.String("targetNames").Optional().
			Comment(`目标任务名,当任务为一天多次时,逗号分割`),
		field.String("targetAmounts").Optional().
			Comment(`目标值,当任务为一天多次时,逗号分割`),
		field.Bool("status").Optional().
			Comment(`是否启用`),
		field.Int64("sortNum").Optional().
			Comment(`排序字段`),
		field.String("actionType").Optional().
			Comment(`跳转动作,pay`),
	}
	fields = append(fields, mixin.EETimeFields()...)
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{}
}
