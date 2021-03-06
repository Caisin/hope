package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelClassify holds the schema definition for the NovelClassify entity.
type NovelClassify struct {
	ent.Schema
}

// Fields of the NovelClassify.
func (NovelClassify) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("pid").Optional().
			Comment(`父类ID,默认0,表示无`),
		field.String("classifyName").Optional().
			Comment(`分类名称`),
		field.Int32("status").Optional().
			Comment(`状态 0`),
		field.Int32("orderNum").Optional().
			Comment(`排序字段`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the NovelClassify.
func (NovelClassify) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("novels", Novel.Type).Comment("当前分类下的小说"),
	}
}
