package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelBuyChapterRecord holds the schema definition for the NovelBuyChapterRecord entity.
type NovelBuyChapterRecord struct {
	ent.Schema
}

// Fields of the NovelBuyChapterRecord.
func (NovelBuyChapterRecord) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("recordId").
			Comment(`主键编码`),
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.String("userName").Optional().
			Comment(`用户名称`),
		field.Int64("chapterId").Optional().
			Comment(`章节ID`),
		field.Int("chapterOrderNum").Optional().
			Comment(`章节序号`),
		field.Int64("novelId").Optional().
			Comment(`书号`),
		field.String("novelName").Optional().
			Comment(`书名`),
		field.String("chapterName").Optional().
			Comment(`章节名称`),
		field.Bool("isSvip").Optional().
			Comment(`是否超级vip记录`),
		field.Int64("coin").Optional().
			Comment(`花费书币`),
		field.Int64("coupon").Optional().
			Comment(`花费书券`),
		field.Int64("discount").Optional().
			Comment(`折扣金额`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the NovelBuyChapterRecord.
func (NovelBuyChapterRecord) Edges() []ent.Edge {
	return []ent.Edge{}
}
