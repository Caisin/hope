package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// ResourceGroup holds the schema definition for the ResourceGroup entity.
type ResourceGroup struct {
	ent.Schema
}

// Fields of the ResourceGroup.
func (ResourceGroup) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.String("name").Optional().
			Comment(`Name`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the ResourceGroup.
func (ResourceGroup) Edges() []ent.Edge {
	return []ent.Edge{}
}
