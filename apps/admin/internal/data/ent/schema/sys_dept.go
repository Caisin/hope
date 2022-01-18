package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.String("deptPath").Optional().
			Comment(`部门路径`),
		field.String("deptName").Optional().
			Comment(`部门名称`),
		field.Int32("sort").Optional().
			Comment(`排序`),
		field.String("leader").Optional().
			Comment(`负责人`),
		field.String("phone").Optional().
			Comment(`手机`),
		field.String("email").Optional().
			Comment(`邮箱`),
		field.Int32("status").Optional().
			Comment(`状态`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysDept.
func (SysDept) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("childes", SysDept.Type).From("parent").Unique(),
		edge.To("users", SysUser.Type),
	}
}
