package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// QiniuConfig holds the schema definition for the QiniuConfig entity.
type QiniuConfig struct {
	ent.Schema
}

// Fields of the QiniuConfig.
func (QiniuConfig) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.String("accessKey").Optional().
			Comment(`accessKey`),
		field.String("bucket").Optional().
			Comment(`Bucket 识别符`),
		field.String("host").Optional().
			Comment(`外链域名`),
		field.String("secretKey").Optional().
			Comment(`secretKey`),
		field.String("type").Optional().
			Comment(`空间类型`),
		field.String("zone").Optional().
			Comment(`机房`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the QiniuConfig.
func (QiniuConfig) Edges() []ent.Edge {
	return []ent.Edge{}
}
