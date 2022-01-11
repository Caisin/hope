package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// Activity holds the schema definition for the Activity entity.
type Activity struct {
	ent.Schema
}

// Fields of the Activity.
func (Activity) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("activityCode").Optional().
			Comment(`活动编码`),
		field.String("activityName").Optional().
			Comment(`活动名称`),
		field.String("summary").Optional().
			Comment(`活动描述`),
		field.String("ruleImgSc").Optional().
			Comment(`活动规则简体中文图`),
		field.String("ruleImgTc").Optional().
			Comment(`活动规则繁体中文图`),
		field.String("popupImg").Optional().
			Comment(`弹出框图片地址`),
		field.Int("regDays").Optional().
			Comment(`注册天数`),
		field.String("cycleType").Optional().
			Comment(`循环类型,默认`),
		field.Time("effectTime").Optional().
			Comment(`活动生效时间`),
		field.Time("expiredTime").Optional().
			Comment(`活动失效时间`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the Activity.
func (Activity) Edges() []ent.Edge {
	return []ent.Edge{}
}
