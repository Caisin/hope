package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// CustomerNovels holds the schema definition for the CustomerNovels entity.
type CustomerNovels struct {
	ent.Schema
}

// Fields of the CustomerNovels.
func (CustomerNovels) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.Int64("novelId").Optional().
			Comment(`小说编码`),
		field.Int("typeId").Optional().
			Comment(`类型编号,见参数表sys_dic_data,dicType=customer_type,dic_sort值`),
		field.String("typeCode").Optional().
			Comment(`类型编码,见参数表sys_dic_data,dicType=customer_type,dic_type值`),
		field.String("groupCode").Optional().
			Comment(`分组编码`),
		field.String("fieldName").Optional().
			Comment(`字段`),
		field.String("cover").Optional().
			Comment(`封面,无值使用小说封面`),
		field.Int("orderNum").Optional().
			Comment(`排序字段,值越大,越靠前`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.EETimeFields()...)
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the CustomerNovels.
func (CustomerNovels) Edges() []ent.Edge {
	return []ent.Edge{}
}
