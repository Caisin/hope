package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// TaskLog holds the schema definition for the TaskLog entity.
type TaskLog struct {
	ent.Schema
}

// Fields of the TaskLog.
func (TaskLog) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("userId").
			Comment(`用户ID`),
		field.String("taskGroup").Optional().
			Comment(`任务分组`),
		field.String("taskCode").Optional().
			Comment(`任务编码`),
		field.Int64("taskId").Optional().
			Comment(`任务Id`),
		field.String("taskName").Optional().
			Comment(`任务名称`),
		field.Int64("amount").Optional().
			Comment(`奖励金额`),
		field.Int64("reward").Optional().
			Comment(`额外奖励金额`),
		field.Int32("amountItem").Optional().
			Comment(`奖励资产科目`),
		field.Int32("rewardItem").Optional().
			Comment(`额外奖励资产科目`),
		field.Int64("targetAmount").Optional().
			Comment(`目标值`),
		field.Int64("doneAmount").Optional().
			Comment(`完成值`),
		field.Int64("state").Optional().
			Comment(`状态,0`),
		field.Time("doneAt").Optional().
			Comment(`完成时间`),
		field.Time("obtainAt").Optional().
			Comment(`领取奖励时间`),
		field.Int32("doneTimes").Optional().
			Comment(`完成次数`),
		field.Int32("allTimes").Optional().
			Comment(`可完成次数`),
	}
	fields = append(fields, mixin.EETimeFields()...)
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the TaskLog.
func (TaskLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", SocialUser.Type).Field("userId").Required().Ref("tasks").Unique(),
	}
}
