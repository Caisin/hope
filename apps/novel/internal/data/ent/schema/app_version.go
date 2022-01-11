package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// AppVersion holds the schema definition for the AppVersion entity.
type AppVersion struct {
	ent.Schema
}

// Fields of the AppVersion.
func (AppVersion) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.String("title").Optional().
			Comment(`标题`),
		field.Int("version").Optional().
			Comment(`版本号`),
		field.String("updateInfo").Optional().
			Comment(`更新信息`),
		field.String("downloadUrl").Optional().
			Comment(`下载地址`),
		field.String("platform").Optional().
			Comment(`平台`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the AppVersion.
func (AppVersion) Edges() []ent.Edge {
	return []ent.Edge{}
}
