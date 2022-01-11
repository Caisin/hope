package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// CasbinRule holds the schema definition for the CasbinRule entity.
type CasbinRule struct {
	ent.Schema
}

// Fields of the CasbinRule.
func (CasbinRule) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("p_type").Optional().
			Comment(``),
		field.String("v0").Optional().
			Comment(``),
		field.String("v1").Optional().
			Comment(``),
		field.String("v2").Optional().
			Comment(``),
		field.String("v3").Optional().
			Comment(``),
		field.String("v4").Optional().
			Comment(``),
		field.String("v5").Optional().
			Comment(``),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the CasbinRule.
func (CasbinRule) Edges() []ent.Edge {
	return []ent.Edge{}
}
