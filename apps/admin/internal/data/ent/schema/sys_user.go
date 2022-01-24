package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.String("username").
			Comment(`用户名`),
		field.String("password").
			Comment(`密码`),
		field.String("nickName").
			Comment(`昵称`),
		field.String("phone").Optional().
			Comment(`手机号`),
		field.Int64("deptId").Optional().
			Comment(`部门ID`),
		field.Int64("postId").Optional().
			Comment(`岗位ID`),
		field.Int64("roleId").Optional().
			Comment(`角色ID`),
		field.String("avatar").Optional().
			Comment(`头像`),
		field.Int32("sex").Optional().
			Comment(`性别,0保密,1男,2女`),
		field.String("email").Optional().
			Comment(`邮箱`),
		field.String("remark").Optional().
			Comment(`备注`),
		field.String("desc").Optional().
			Comment(`个人简介`),
		field.String("homePath").Optional().
			Comment(`登陆默认打开页面`),
		field.String("status").Optional().
			Comment(`状态`),
		field.String("extInfo").Optional().
			Comment(`拓展信息,facebook回传用的`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysUser.
func (SysUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dept", SysDept.Type).Field("deptId").Comment("部门").Ref("users").Unique(),
		edge.From("post", SysPost.Type).Field("postId").Comment("岗位").Ref("users").Unique(),
		edge.From("role", SysRole.Type).Comment("角色").Ref("users"),
		edge.To("loginLogs", SysLoginLog.Type).Comment("登陆日志"),
		edge.To("operaLogs", SysOperaLog.Type).Comment("操作日志"),
	}
}
