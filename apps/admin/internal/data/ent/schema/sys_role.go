package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysRole holds the schema definition for the SysRole entity.
type SysRole struct {
	ent.Schema
}

// Fields of the SysRole.
func (SysRole) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int("roleId").
			Comment(`角色编码`),
		field.String("roleName").Optional().
			Comment(`角色名称`),
		field.String("status").Optional().
			Comment(`状态`),
		field.String("roleKey").Optional().
			Comment(`角色代码`),
		field.Int("roleSort").Optional().
			Comment(`角色排序`),
		field.String("flag").Optional().
			Comment(`标记`),
		field.String("remark").Optional().
			Comment(`备注`),
		field.Bool("admin").Optional().
			Comment(`是否管理员`),
		field.String("dataScope").Optional().
			Comment(`1.全部数据权限
2.自定数据权限
3.本部门数据权限
4.本部门及以下数据权限
5.仅本人数据权限`),
		field.String("sysDept").Optional().
			Comment(``),
		field.String("sysMenu").Optional().
			Comment(``),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysRole.
func (SysRole) Edges() []ent.Edge {
	return []ent.Edge{}
}
