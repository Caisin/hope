package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// ClientError holds the schema definition for the ClientError entity.
type ClientError struct {
	ent.Schema
}

// Fields of the ClientError.
func (ClientError) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.String("appVersion").Optional().
			Comment(`App版本`),
		field.String("deviceName").Optional().
			Comment(`设备名称`),
		field.String("osName").Optional().
			Comment(`操作系统名称`),
		field.String("errorInfo").Optional().
			Comment(`错误信息`),
		field.Int64("userId").Optional().
			Comment(`用户ID`),
	}
	fields = append(fields, mixin.TimeMixin{}.Fields()...)
	return fields
}

// Edges of the ClientError.
func (ClientError) Edges() []ent.Edge {
	return []ent.Edge{}
}
