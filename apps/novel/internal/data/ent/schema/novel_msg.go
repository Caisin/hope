package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelMsg holds the schema definition for the NovelMsg entity.
type NovelMsg struct {
	ent.Schema
}

// Fields of the NovelMsg.
func (NovelMsg) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.String("title").Optional().
			Comment(`消息标题`),
		field.String("msg").Optional().
			Comment(`消息内容`),
		field.String("msgType").Optional().
			Comment(`消息类型,user`),
		field.Bool("status").Optional().
			Comment(`状态`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the NovelMsg.
func (NovelMsg) Edges() []ent.Edge {
	return []ent.Edge{}
}
