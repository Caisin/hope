package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// BookPackage holds the schema definition for the BookPackage entity.
type BookPackage struct {
	ent.Schema
}

// Fields of the BookPackage.
func (BookPackage) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("activityCode").Optional().
			Comment(`活动编码`),
		field.String("packageName").Optional().
			Comment(`书籍包名称`),
		field.Int64("price").Optional().
			Comment(`活动打包价格`),
		field.Int64("dailyPrice").Optional().
			Comment(`日常价格,既所有书本正常购买的价格`),
	}
	fields = append(fields, mixin.EETimeFields()...)
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the BookPackage.
func (BookPackage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("books", Novel.Type).Ref("pkgs"),
	}
}
