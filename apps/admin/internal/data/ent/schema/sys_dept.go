package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysDept holds the schema definition for the SysDept entity.
type SysDept struct {
	ent.Schema
}

// Fields of the SysDept.
func (SysDept) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int("deptId").
			Comment(`部门编码`),
		field.Int("parentId").Optional().
			Comment(`上级部门`),
		field.String("deptPath").Optional().
			Comment(`部门路径`),
		field.String("deptName").Optional().
			Comment(`部门名称`),
		field.Int("sort").Optional().
			Comment(`排序`),
		field.String("leader").Optional().
			Comment(`负责人`),
		field.String("phone").Optional().
			Comment(`手机`),
		field.String("email").Optional().
			Comment(`邮箱`),
		field.Int("status").Optional().
			Comment(`状态`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysDept.
func (SysDept) Edges() []ent.Edge {
	return []ent.Edge{}
}
