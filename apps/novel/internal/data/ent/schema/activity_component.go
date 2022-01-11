package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// ActivityComponent holds the schema definition for the ActivityComponent entity.
type ActivityComponent struct {
	ent.Schema
}

// Fields of the ActivityComponent.
func (ActivityComponent) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.String("activityCode").Optional().
			Comment(`活动编码`),
		field.String("componentType").Optional().
			Comment(`活动组成类型,reg`),
		field.String("policy").Optional().
			Comment(`策略,all`),
		field.Bool("vipDays").Optional().
			Comment(`vip拓展天数,当用户为vip时生效`),
		field.Int64("minConsume").Optional().
			Comment(`最小消耗金额`),
		field.Int64("maxConsume").Optional().
			Comment(`最大消耗金额`),
		field.Int64("minPayNum").Optional().
			Comment(`充值次数`),
		field.Int64("payTimes").Optional().
			Comment(`第几次充值`),
		field.Int64("payAmount").Optional().
			Comment(`充值金额`),
		field.Int64("regDays").Optional().
			Comment(`注册天数`),
		field.String("summary").Optional().
			Comment(`活动描述`),
		field.Int("assetItemId").Optional().
			Comment(`活动期间赠送资金科目`),
		field.Int64("amount").Optional().
			Comment(`活动期间赠送金额,-1为赠送等额消耗书券`),
		field.Int64("resId").Optional().
			Comment(`获得资产编号`),
		field.Int("resDays").Optional().
			Comment(`获得资产有效时间`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the ActivityComponent.
func (ActivityComponent) Edges() []ent.Edge {
	return []ent.Edge{}
}
