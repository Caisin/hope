package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// PageConfig holds the schema definition for the PageConfig entity.
type PageConfig struct {
	ent.Schema
}

// Fields of the PageConfig.
func (PageConfig) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.String("pageCode").Optional().
			Comment(`页面编码`),
		field.String("pageName").Optional().
			Comment(`页面名称`),
		field.String("groupCodes").Optional().
			Comment(`分组编码集,逗号分割`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the PageConfig.
func (PageConfig) Edges() []ent.Edge {
	return []ent.Edge{}
}
