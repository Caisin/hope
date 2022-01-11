package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// AmBalance holds the schema definition for the AmBalance entity.
type AmBalance struct {
	ent.Schema
}

// Fields of the AmBalance.
func (AmBalance) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.String("orderId").Optional().
			Comment(`订单号`),
		field.Int64("eventId").Optional().
			Comment(`关联用户事件Id`),
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.Int("cashTag").Optional().
			Comment(`现金标识,0优惠券 1书币`),
		field.Int("assetItemId").Optional().
			Comment(`账本科目`),
		field.Int64("amount").Optional().
			Comment(`原始金额`),
		field.Int64("balance").Optional().
			Comment(`剩余可用`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.EETimeFields()...)
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the AmBalance.
func (AmBalance) Edges() []ent.Edge {
	return []ent.Edge{}
}
