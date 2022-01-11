package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelAutoBuy holds the schema definition for the NovelAutoBuy entity.
type NovelAutoBuy struct {
	ent.Schema
}

// Fields of the NovelAutoBuy.
func (NovelAutoBuy) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.Int64("novelId").Optional().
			Comment(`小说编号`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the NovelAutoBuy.
func (NovelAutoBuy) Edges() []ent.Edge {
	return []ent.Edge{}
}
