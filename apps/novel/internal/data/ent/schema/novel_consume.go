package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelConsume holds the schema definition for the NovelConsume entity.
type NovelConsume struct {
	ent.Schema
}

// Fields of the NovelConsume.
func (NovelConsume) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("novelId").
			Comment(`书号`),
		field.Int64("coin").Optional().
			Comment(`书币`),
		field.Int64("coupon").Optional().
			Comment(`书券`),
		field.Int64("discount").Optional().
			Comment(`VIP折扣金额`),
		field.Int64("reward").Optional().
			Comment(`打赏金额`),
	}
	fields = append(fields, mixin.TimeMixin{}.Fields()...)
	return fields
}

// Edges of the NovelConsume.
func (NovelConsume) Edges() []ent.Edge {
	return []ent.Edge{}
}
