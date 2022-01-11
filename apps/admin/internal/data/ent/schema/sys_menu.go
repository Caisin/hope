package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysMenu holds the schema definition for the SysMenu entity.
type SysMenu struct {
	ent.Schema
}

// Fields of the SysMenu.
func (SysMenu) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("menuName").Optional().
			Comment(`菜单名`),
		field.String("title").Optional().
			Comment(`菜单标题`),
		field.String("icon").Optional().
			Comment(`图标`),
		field.String("path").Optional().
			Comment(`路径`),
		field.String("paths").Optional().
			Comment(`多级路径`),
		field.String("menuType").Optional().
			Comment(`M-目录C-菜单F-按钮`),
		field.String("action").Optional().
			Comment(``),
		field.String("permission").Optional().
			Comment(`权限`),
		field.Int("parentId").Optional().
			Comment(`父节点`),
		field.Bool("noCache").Optional().
			Comment(`无缓存`),
		field.String("breadcrumb").Optional().
			Comment(`面包屑`),
		field.String("component").Optional().
			Comment(`组件`),
		field.Int("sort").Optional().
			Comment(`排序`),
		field.Bool("visible").Optional().
			Comment(`是否可见`),
		field.Bool("isFrame").Optional().
			Comment(`是否外链1是0否`),
		field.String("sysApi").Optional().
			Comment(``),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysMenu.
func (SysMenu) Edges() []ent.Edge {
	return []ent.Edge{}
}
