package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysLoginLog holds the schema definition for the SysLoginLog entity.
type SysLoginLog struct {
	ent.Schema
}

// Fields of the SysLoginLog.
func (SysLoginLog) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("username").Optional().
			Comment(`用户名`),
		field.String("status").Optional().
			Comment(`状态`),
		field.String("ipaddr").Optional().
			Comment(`ip地址`),
		field.String("loginLocation").Optional().
			Comment(`归属地`),
		field.String("browser").Optional().
			Comment(`浏览器`),
		field.String("os").Optional().
			Comment(`系统`),
		field.String("platform").Optional().
			Comment(`固件`),
		field.Time("loginTime").Optional().
			Comment(`登录时间`),
		field.String("remark").Optional().
			Comment(`备注`),
		field.String("msg").Optional().
			Comment(`信息`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysLoginLog.
func (SysLoginLog) Edges() []ent.Edge {
	return []ent.Edge{}
}
