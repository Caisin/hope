package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// UserEvent holds the schema definition for the UserEvent entity.
type UserEvent struct {
	ent.Schema
}

// Fields of the UserEvent.
func (UserEvent) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.String("eventType").Optional().
			Comment(`事件类型,read`),
		field.Int64("novelId").Optional().
			Comment(`书号`),
		field.Int64("chapterId").Optional().
			Comment(`章节号`),
		field.Int64("coin").Optional().
			Comment(`书币`),
		field.Int64("coupon").Optional().
			Comment(`书券`),
		field.Int64("money").Optional().
			Comment(`金额,充值金额`),
		field.String("keyword").Optional().
			Comment(`搜索关键字`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the UserEvent.
func (UserEvent) Edges() []ent.Edge {
	return []ent.Edge{}
}
