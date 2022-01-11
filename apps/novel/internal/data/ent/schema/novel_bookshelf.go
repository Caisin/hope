package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelBookshelf holds the schema definition for the NovelBookshelf entity.
type NovelBookshelf struct {
	ent.Schema
}

// Fields of the NovelBookshelf.
func (NovelBookshelf) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.String("userName").Optional().
			Comment(`用户名`),
		field.Int64("novelId").Optional().
			Comment(`小说编号`),
		field.Time("lastReadTime").Optional().
			Comment(`最后阅读时间`),
		field.Int("lastChapterOrder").Optional().
			Comment(`章节序号`),
		field.Int64("lastChapterId").Optional().
			Comment(`章节ID`),
		field.String("lastChapterName").Optional().
			Comment(`最后阅读章节名称`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the NovelBookshelf.
func (NovelBookshelf) Edges() []ent.Edge {
	return []ent.Edge{}
}
