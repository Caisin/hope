package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelChapter holds the schema definition for the NovelChapter entity.
type NovelChapter struct {
	ent.Schema
}

// Fields of the NovelChapter.
func (NovelChapter) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("chapterId").
			Comment(`主键编码`),
		field.Int64("novelId").Optional().
			Comment(`小说编号`),
		field.Int("orderNum").Optional().
			Comment(`章节序号`),
		field.String("chapterName").Optional().
			Comment(`章节名称`),
		field.String("content").Optional().
			Comment(`章节内容`),
		field.String("mediaKey").Optional().
			Comment(`阿里云音频目录`),
		field.String("duration").Optional().
			Comment(`音频时长`),
		field.Time("publishTime").Optional().
			Comment(`发布时间`),
		field.Int("status").Optional().
			Comment(`状态：0 草稿 ，1 发布`),
		field.Bool("isFree").Optional().
			Comment(`0`),
		field.Int64("price").Optional().
			Comment(`价格`),
		field.Int("wordNum").Optional().
			Comment(`章节字数`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.TimeMixin{}.Fields()...)
	return fields
}

// Edges of the NovelChapter.
func (NovelChapter) Edges() []ent.Edge {
	return []ent.Edge{}
}
