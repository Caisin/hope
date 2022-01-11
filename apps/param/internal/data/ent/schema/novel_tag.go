package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelTag holds the schema definition for the NovelTag entity.
type NovelTag struct {
	ent.Schema
}

// Fields of the NovelTag.
func (NovelTag) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("tagId").
			Comment(`主键`),
		field.String("tagName").Optional().
			Comment(`标签名称`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.TimeMixin{}.Fields()...)
	return fields
}

// Edges of the NovelTag.
func (NovelTag) Edges() []ent.Edge {
	return []ent.Edge{}
}
