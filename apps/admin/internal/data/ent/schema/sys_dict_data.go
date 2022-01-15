package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysDictData holds the schema definition for the SysDictData entity.
type SysDictData struct {
	ent.Schema
}

// Fields of the SysDictData.
func (SysDictData) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int("dictSort").Optional().
			Comment(`字典排序`),
		field.String("dictLabel").Optional().
			Comment(`字典标签`),
		field.String("dictValue").Optional().
			Comment(`字典值`),
		field.String("dictType").Optional().
			Comment(`字典类型`),
		field.String("isDefault").Optional().
			Comment(`是否默认`),
		field.Int("status").Optional().
			Comment(`状态`),
		field.String("default").Optional().
			Comment(`默认`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysDictData.
func (SysDictData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dictType", SysDictType.Type).Ref("dataList").Unique(),
	}
}
