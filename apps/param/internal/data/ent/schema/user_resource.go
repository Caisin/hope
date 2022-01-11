package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// UserResource holds the schema definition for the UserResource entity.
type UserResource struct {
	ent.Schema
}

// Fields of the UserResource.
func (UserResource) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("resType").Optional().
			Comment(`资源类型,avatar`),
		field.String("Name").Optional().
			Comment(`头像名称`),
		field.String("url").Optional().
			Comment(`图片地址`),
		field.String("summary").Optional().
			Comment(`描述`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the UserResource.
func (UserResource) Edges() []ent.Edge {
	return []ent.Edge{}
}
