package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysConfig holds the schema definition for the SysConfig entity.
type SysConfig struct {
	ent.Schema
}

// Fields of the SysConfig.
func (SysConfig) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("configName").Optional().
			Comment(`配置名称`),
		field.String("configKey").Optional().
			Comment(`配置Key`),
		field.String("configValue").Optional().
			Comment(`配置值`),
		field.String("configType").Optional().
			Comment(`配置类型`),
		field.Int32("isFrontend").Optional().
			Comment(`是否前台`),
		field.Enum("state").
			Values("U", "E").
			Default("U").
			Comment(`状态:U:使用状态,E:失效状态`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysConfig.
func (SysConfig) Edges() []ent.Edge {
	return []ent.Edge{}
}
