package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type HopeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").
			Immutable().
			Default(time.Now).Comment("创建时间"),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).Comment("更新时间"),
		field.Int64("createBy").
			Default(0).Comment("创建者"),
		field.Int64("controlBy").
			Default(0),
		field.Int64("tenantId").
			Default(0).Comment("租户"),
	}
}

func EETimeFields() []ent.Field {
	return []ent.Field{
		field.Time("effectTime").
			Default(time.Now).Comment("生效时间"),
		field.Time("expiredTime").
			Default(time.Now).
			Comment("失效时间"),
	}
}
