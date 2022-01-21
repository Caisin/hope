package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Int64("userId").
			Comment(`用户ID`),
		field.Int64("vipType").Optional().
			Comment(`vip类型`),
		field.Int64("svipType").Optional().
			Comment(`svip类型`),
		field.Time("svipEffectTime").Optional().
			Comment(`超级VIP生效时间`),
		field.Time("svipExpiredTime").Optional().
			Comment(`超级VIP失效时间`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.EETimeFields()...)
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the VipUser.
func (VipUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", SocialUser.Type).Field("userId").Required().Unique().Comment("用户").Ref("vips"),
	}
}
