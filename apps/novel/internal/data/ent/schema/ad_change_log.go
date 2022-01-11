package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// AdChangeLog holds the schema definition for the AdChangeLog entity.
type AdChangeLog struct {
	ent.Schema
}

// Fields of the AdChangeLog.
func (AdChangeLog) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.String("adId").Optional().
			Comment(`广告ID`),
		field.Int64("chId").Optional().
			Comment(`渠道ID`),
		field.String("deviceId").Optional().
			Comment(`设备号`),
		field.String("extInfo").Optional().
			Comment(`手机拓展信息`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the AdChangeLog.
func (AdChangeLog) Edges() []ent.Edge {
	return []ent.Edge{}
}
