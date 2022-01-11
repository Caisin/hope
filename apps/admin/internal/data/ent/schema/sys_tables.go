package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysTables holds the schema definition for the SysTables entity.
type SysTables struct {
	ent.Schema
}

// Fields of the SysTables.
func (SysTables) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("tableName").Optional().
			Comment(`表名称`),
		field.String("tableComment").Optional().
			Comment(`表备注`),
		field.String("className").Optional().
			Comment(`类名`),
		field.String("tplCategory").Optional().
			Comment(``),
		field.String("packageName").Optional().
			Comment(`包名`),
		field.String("moduleName").Optional().
			Comment(`go文件名`),
		field.String("moduleFrontName").Optional().
			Comment(`前端文件名`),
		field.String("businessName").Optional().
			Comment(``),
		field.String("functionName").Optional().
			Comment(`功能名称`),
		field.String("functionAuthor").Optional().
			Comment(`功能作者`),
		field.String("pkColumn").Optional().
			Comment(``),
		field.String("pkGoField").Optional().
			Comment(``),
		field.String("pkJsonField").Optional().
			Comment(``),
		field.String("options").Optional().
			Comment(``),
		field.String("treeCode").Optional().
			Comment(``),
		field.String("treeParentCode").Optional().
			Comment(``),
		field.String("treeName").Optional().
			Comment(``),
		field.Bool("tree").Optional().
			Comment(``),
		field.Bool("crud").Optional().
			Comment(``),
		field.String("remark").Optional().
			Comment(``),
		field.Int("isDataScope").Optional().
			Comment(``),
		field.Int("isActions").Optional().
			Comment(``),
		field.Int("isAuth").Optional().
			Comment(``),
		field.String("isLogicalDelete").Optional().
			Comment(``),
		field.Bool("logicalDelete").Optional().
			Comment(``),
		field.String("logicalDeleteColumn").Optional().
			Comment(``),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysTables.
func (SysTables) Edges() []ent.Edge {
	return []ent.Edge{}
}
