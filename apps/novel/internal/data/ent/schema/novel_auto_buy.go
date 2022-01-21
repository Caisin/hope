package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Int64("userId").
			Comment(`用户ID`),
		field.Int64("novelId").Optional().
			Comment(`小说编号`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the NovelAutoBuy.
func (NovelAutoBuy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", SocialUser.Type).Field("userId").Required().Ref("autoBuyNovels").Unique(),
	}
}
