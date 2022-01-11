package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysUser holds the schema definition for the SysUser entity.
type SysUser struct {
	ent.Schema
}

// Fields of the SysUser.
func (SysUser) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("userId").
			Comment(`用户ID`),
		field.String("username").Optional().
			Comment(`用户名`),
		field.String("nickName").Optional().
			Comment(`昵称`),
		field.String("phone").Optional().
			Comment(`手机号`),
		field.Int("roleId").Optional().
			Comment(`角色ID`),
		field.String("avatar").Optional().
			Comment(`头像`),
		field.Int("sex").Optional().
			Comment(`性别,0保密,1男,2女`),
		field.String("email").Optional().
			Comment(`邮箱`),
		field.Int("deptId").Optional().
			Comment(`部门`),
		field.Int("postId").Optional().
			Comment(`岗位`),
		field.String("remark").Optional().
			Comment(`备注`),
		field.String("status").Optional().
			Comment(`状态`),
		field.String("extInfo").Optional().
			Comment(`拓展信息,facebook回传用的`),
	}
	fields = append(fields, mixin.TimeMixin{}.Fields()...)
	return fields
}

// Edges of the SysUser.
func (SysUser) Edges() []ent.Edge {
	return []ent.Edge{}
}
