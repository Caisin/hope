package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Int64("parentId").
			Default(0).
			Comment(`父菜单Id`),
		field.String("name").
			Comment(`菜单名`),
		field.String("title").
			Comment(`菜单标题`),
		field.String("redirect").Optional().
			Comment(`跳转路径`),
		field.String("icon").Optional().
			Comment(`图标`),
		field.String("path").Optional().
			Comment(`路径`),
		field.String("paths").Optional().
			Comment(`多级路径`),
		field.String("menuType").Optional().
			Comment(`D-目录M-菜单F-按钮`),
		field.String("action").Optional().
			Comment(``),
		field.String("permission").Optional().
			Comment(`权限`),
		field.Bool("ignoreKeepAlive").Optional().
			Comment(`不缓存`),
		field.Bool("hideBreadcrumb").Optional().
			Comment(`隐藏面包屑`),
		field.Bool("hideChildrenInMenu").Optional().
			Comment(`隐藏子菜单`),
		field.String("component").Optional().
			Comment(`组件`),
		field.Int32("sort").Optional().
			Comment(`排序`),
		field.Bool("hideMenu").Optional().
			Comment(`影藏菜单`),
		field.String("frameSrc").Optional().
			Comment(`外链地址`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysMenu.
func (SysMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", SysRole.Type).Ref("menus"),
		edge.To("children", SysMenu.Type).From("parent").Field("parentId").Required().Unique(),
	}
}
