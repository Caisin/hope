package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// UserResourceRecord holds the schema definition for the UserResourceRecord entity.
type UserResourceRecord struct {
	ent.Schema
}

// Fields of the UserResourceRecord.
func (UserResourceRecord) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.Int64("resId").Optional().
			Comment(`资源ID`),
		field.Bool("def").Optional().
			Comment(`是否默认`),
		field.String("name").Optional().
			Comment(`头像名称`),
		field.String("url").Optional().
			Comment(`资源地址`),
		field.String("resType").Optional().
			Comment(`资源类型,avatar`),
		field.String("remark").Optional().
			Comment(`备注`),
		field.Int32("state").Optional().
			Comment(`使用状态,1`),
	}
	fields = append(fields, mixin.EETimeFields()...)
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the UserResourceRecord.
func (UserResourceRecord) Edges() []ent.Edge {
	return []ent.Edge{}
}
