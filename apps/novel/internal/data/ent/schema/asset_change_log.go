package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// AssetChangeLog holds the schema definition for the AssetChangeLog entity.
type AssetChangeLog struct {
	ent.Schema
}

// Fields of the AssetChangeLog.
func (AssetChangeLog) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("orderId").Optional().
			Comment(`订单号`),
		field.Int64("balanceId").Optional().
			Comment(`账本ID`),
		field.Int64("eventId").Optional().
			Comment(`关联用户事件Id`),
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.Int("assetItemId").Optional().
			Comment(`账本科目`),
		field.Int64("amount").Optional().
			Comment(`变化金额`),
		field.Int64("oldBalance").Optional().
			Comment(`变化前剩余可用`),
		field.Int64("newBalance").Optional().
			Comment(`变化后剩余可用`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the AssetChangeLog.
func (AssetChangeLog) Edges() []ent.Edge {
	return []ent.Edge{}
}
