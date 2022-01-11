package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// CustomerNovelConfig holds the schema definition for the CustomerNovelConfig entity.
type CustomerNovelConfig struct {
	ent.Schema
}

// Fields of the CustomerNovelConfig.
func (CustomerNovelConfig) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("groupCode").Optional().
			Comment(`分组编码`),
		field.String("innerGroupCode").Optional().
			Comment(`嵌套分组`),
		field.String("groupName").Optional().
			Comment(`分组名称`),
		field.Int("typeId").Optional().
			Comment(`类型ID`),
		field.String("TypeCode").Optional().
			Comment(`类型编码`),
		field.String("typeName").Optional().
			Comment(`类型名称`),
		field.String("fieldName").Optional().
			Comment(`字段名称`),
		field.Int("defaultNum").Optional().
			Comment(`默认数量`),
		field.Bool("state").Optional().
			Comment(`是否可用`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the CustomerNovelConfig.
func (CustomerNovelConfig) Edges() []ent.Edge {
	return []ent.Edge{}
}
