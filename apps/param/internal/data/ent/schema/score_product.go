package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// ScoreProduct holds the schema definition for the ScoreProduct entity.
type ScoreProduct struct {
	ent.Schema
}

// Fields of the ScoreProduct.
func (ScoreProduct) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("productName").Optional().
			Comment(`产品名称`),
		field.String("summary").Optional().
			Comment(`产品描述`),
		field.String("cardUrl").Optional().
			Comment(`vip卡图片`),
		field.Int64("score").Optional().
			Comment(`需消耗积分价格`),
		field.Int64("vipType").Optional().
			Comment(`兑换VIP类型`),
	}
	fields = append(fields, mixin.EETimeFields()...)
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the ScoreProduct.
func (ScoreProduct) Edges() []ent.Edge {
	return []ent.Edge{}
}
