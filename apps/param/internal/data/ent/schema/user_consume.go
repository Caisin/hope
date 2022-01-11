package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// UserConsume holds the schema definition for the UserConsume entity.
type UserConsume struct {
	ent.Schema
}

// Fields of the UserConsume.
func (UserConsume) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("novelId").
			Comment(`用户ID`),
		field.Int64("coin").Optional().
			Comment(`书币`),
		field.Int64("coupon").Optional().
			Comment(`书券`),
		field.Int64("discount").Optional().
			Comment(`VIP折扣金额`),
	}
	fields = append(fields, mixin.TimeMixin{}.Fields()...)
	return fields
}

// Edges of the UserConsume.
func (UserConsume) Edges() []ent.Edge {
	return []ent.Edge{}
}
