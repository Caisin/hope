package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysColumns holds the schema definition for the SysColumns entity.
type SysColumns struct {
	ent.Schema
}

// Fields of the SysColumns.
func (SysColumns) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int32("columnId").
			Comment(``),
		field.String("columnName").Optional().
			Comment(``),
		field.String("columnComment").Optional().
			Comment(``),
		field.String("columnType").Optional().
			Comment(``),
		field.String("goType").Optional().
			Comment(``),
		field.String("goField").Optional().
			Comment(``),
		field.String("jsonField").Optional().
			Comment(``),
		field.String("isPk").Optional().
			Comment(``),
		field.String("isIncrement").Optional().
			Comment(``),
		field.String("isRequired").Optional().
			Comment(``),
		field.String("isInsert").Optional().
			Comment(``),
		field.String("isEdit").Optional().
			Comment(``),
		field.String("isList").Optional().
			Comment(``),
		field.String("isQuery").Optional().
			Comment(``),
		field.String("queryType").Optional().
			Comment(``),
		field.String("htmlType").Optional().
			Comment(``),
		field.String("dictType").Optional().
			Comment(``),
		field.Int32("sort").Optional().
			Comment(``),
		field.String("list").Optional().
			Comment(``),
		field.Bool("pk").Optional().
			Comment(``),
		field.Bool("required").Optional().
			Comment(``),
		field.Bool("superColumn").Optional().
			Comment(``),
		field.Bool("usableColumn").Optional().
			Comment(``),
		field.Bool("increment").Optional().
			Comment(``),
		field.Bool("insert").Optional().
			Comment(``),
		field.Bool("edit").Optional().
			Comment(``),
		field.Bool("query").Optional().
			Comment(``),
		field.String("remark").Optional().
			Comment(``),
		field.String("fkLabelName").Optional().
			Comment(``),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the SysColumns.
func (SysColumns) Edges() []ent.Edge {
	return []ent.Edge{}
}
