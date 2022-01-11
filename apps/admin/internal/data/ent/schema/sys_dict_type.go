package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysDictType holds the schema definition for the SysDictType entity.
type SysDictType struct {
	ent.Schema
}

// Fields of the SysDictType.
func (SysDictType) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int("id").
			Comment(`主键编码`),
		field.String("dictName").Optional().
			Comment(`字典名称`),
		field.String("dictType").Optional().
			Comment(`字典类型`),
		field.Int("status").Optional().
			Comment(`状态`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysDictType.
func (SysDictType) Edges() []ent.Edge {
	return []ent.Edge{}
}
