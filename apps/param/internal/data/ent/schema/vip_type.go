package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// VipType holds the schema definition for the VipType entity.
type VipType struct {
	ent.Schema
}

// Fields of the VipType.
func (VipType) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("vipName").Optional().
			Comment(`账本名称`),
		field.Bool("isSuper").Optional().
			Comment(`是否超级vip`),
		field.Int("validDays").Optional().
			Comment(`有效天数`),
		field.Int64("discountRate").Optional().
			Comment(`折扣率`),
		field.Int64("avatarId").Optional().
			Comment(`头像ID`),
		field.String("summary").Optional().
			Comment(`描述`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the VipType.
func (VipType) Edges() []ent.Edge {
	return []ent.Edge{}
}
