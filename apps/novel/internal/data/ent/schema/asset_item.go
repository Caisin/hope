package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// AssetItem holds the schema definition for the AssetItem entity.
type AssetItem struct {
	ent.Schema
}

// Fields of the AssetItem.
func (AssetItem) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int("assetItemId").
			Comment(`账本科目`),
		field.String("assetName").Optional().
			Comment(`账本名称`),
		field.Int("cashTag").Optional().
			Comment(`现金标记,0书券1现金`),
		field.Int("validDays").Optional().
			Comment(`有效天数`),
		field.Time("effectTime").Optional().
			Comment(`生效时间`),
		field.Time("expiredTime").Optional().
			Comment(`失效时间`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the AssetItem.
func (AssetItem) Edges() []ent.Edge {
	return []ent.Edge{}
}
