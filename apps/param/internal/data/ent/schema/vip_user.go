package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// VipUser holds the schema definition for the VipUser entity.
type VipUser struct {
	ent.Schema
}

// Fields of the VipUser.
func (VipUser) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("UserId").
			Comment(`用户ID`),
		field.Int64("vipType").Optional().
			Comment(`vip类型`),
		field.Int64("svipType").Optional().
			Comment(`svip类型`),
		field.Time("effectTime").Optional().
			Comment(`生效时间`),
		field.Time("expiredTime").Optional().
			Comment(`失效时间`),
		field.Time("svipEffectTime").Optional().
			Comment(`超级VIP生效时间`),
		field.Time("svipExpiredTime").Optional().
			Comment(`超级VIP失效时间`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.TimeMixin{}.Fields()...)
	return fields
}

// Edges of the VipUser.
func (VipUser) Edges() []ent.Edge {
	return []ent.Edge{}
}
