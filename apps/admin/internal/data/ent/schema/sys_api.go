package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysApi holds the schema definition for the SysApi entity.
type SysApi struct {
	ent.Schema
}

// Fields of the SysApi.
func (SysApi) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int("id").
			Comment(`主键编码`),
		field.String("handle").Optional().
			Comment(`处理器`),
		field.String("title").Optional().
			Comment(`标题`),
		field.String("path").Optional().
			Comment(`地址`),
		field.String("action").Optional().
			Comment(`请求类型`),
		field.String("type").Optional().
			Comment(`接口类型`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the SysApi.
func (SysApi) Edges() []ent.Edge {
	return []ent.Edge{}
}
