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
		field.Int32("assetItemId").
			Comment(`账本科目`),
		field.String("assetName").Optional().
			Comment(`账本名称`),
		field.Int32("cashTag").Optional().
			Comment(`现金标记,0书券1现金`),
		field.Int32("validDays").Optional().
			Comment(`有效天数`),
	}
	fields = append(fields, mixin.EETimeFields()...)
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the AssetItem.
func (AssetItem) Edges() []ent.Edge {
	return []ent.Edge{}
}
