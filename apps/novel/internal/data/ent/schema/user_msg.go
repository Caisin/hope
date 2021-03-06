package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// UserMsg holds the schema definition for the UserMsg entity.
type UserMsg struct {
	ent.Schema
}

// Fields of the UserMsg.
func (UserMsg) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("userId").
			Comment(`用户ID`),
		field.Int64("msgId").Optional().
			Comment(`消息编码`),
		field.Bool("isRead").Optional().
			Comment(`是否阅读`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the UserMsg.
func (UserMsg) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", SocialUser.Type).Field("userId").Required().Ref("msgs").Unique(),
	}
}
