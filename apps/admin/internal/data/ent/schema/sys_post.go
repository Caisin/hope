package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysPost holds the schema definition for the SysPost entity.
type SysPost struct {
	ent.Schema
}

// Fields of the SysPost.
func (SysPost) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("postName").Optional().
			Comment(`岗位名称`),
		field.String("postCode").Optional().
			Comment(`岗位代码`),
		field.Int32("sort").Optional().
			Comment(`岗位排序`),
		field.Int32("status").Optional().
			Comment(`状态`),
		field.String("remark").Optional().
			Comment(`描述`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysPost.
func (SysPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", SysUser.Type).Comment("岗位上的用户"),
	}
}
