package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelBuyRecord holds the schema definition for the NovelBuyRecord entity.
type NovelBuyRecord struct {
	ent.Schema
}

// Fields of the NovelBuyRecord.
func (NovelBuyRecord) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.String("userName").Optional().
			Comment(`用户名称`),
		field.Int64("novelId").Optional().
			Comment(`书号`),
		field.String("novelName").Optional().
			Comment(`书名`),
		field.Int64("packageId").Optional().
			Comment(`书籍打包ID`),
		field.String("cover").Optional().
			Comment(`书封面`),
		field.Int64("coin").Optional().
			Comment(`花费书币`),
		field.Int64("coupon").Optional().
			Comment(`花费书券`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the NovelBuyRecord.
func (NovelBuyRecord) Edges() []ent.Edge {
	return []ent.Edge{}
}
