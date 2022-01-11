package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// UserAnalysisStatistics holds the schema definition for the UserAnalysisStatistics entity.
type UserAnalysisStatistics struct {
	ent.Schema
}

// Fields of the UserAnalysisStatistics.
func (UserAnalysisStatistics) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("statisticsDate").
			Comment(`统计日期`),
		field.Int64("allUserNum").Optional().
			Comment(`累计注册`),
		field.Int64("allPayment").Optional().
			Comment(`累计充值`),
		field.Int64("allPayUser").Optional().
			Comment(`总充值人数`),
		field.Int64("allOrderNum").Optional().
			Comment(`总订单数`),
		field.Int64("dayUserNum").Optional().
			Comment(`当日注册人数`),
		field.Int64("dayPayment").Optional().
			Comment(`当日订单金额`),
		field.Int64("dayOrderNum").Optional().
			Comment(`当日订单数`),
		field.Int64("dayPayUser").Optional().
			Comment(`当日充值人数`),
		field.Int64("dayRegPayment").Optional().
			Comment(`当日注册充值`),
		field.Int64("dayRegUserNum").Optional().
			Comment(`当日注册充值人数`),
		field.Int64("dayRegOrderNum").Optional().
			Comment(`当日注册充值笔数`),
		field.Int64("oldRegPayment").Optional().
			Comment(`当日老用户充值`),
		field.Int64("oldRegUserNum").Optional().
			Comment(`当日老用户充值人数`),
		field.Int64("oldRegOrderNum").Optional().
			Comment(`当日老用户充值笔数`),
		field.Int64("payRate").Optional().
			Comment(`付费率=当日充值笔数/当日注册`),
		field.Int64("arpu").Optional().
			Comment(`ARPU(average revenue per user)=充值/注册`),
		field.Int64("dayRegArpu").Optional().
			Comment(`当日注册充值客单价=当日注册充值/当日注册充值笔数`),
		field.Int64("dayArpu").Optional().
			Comment(`今日充值客单价=今日充值金额/充值笔数`),
		field.Int64("dayOldArpu").Optional().
			Comment(`老用户充值客单价=老用户充值/老用户充值笔数`),
	}
	fields = append(fields, mixin.TimeMixin{}.Fields()...)
	return fields
}

// Edges of the UserAnalysisStatistics.
func (UserAnalysisStatistics) Edges() []ent.Edge {
	return []ent.Edge{}
}
