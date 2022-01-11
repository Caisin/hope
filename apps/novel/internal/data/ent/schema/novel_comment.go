package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelComment holds the schema definition for the NovelComment entity.
type NovelComment struct {
	ent.Schema
}

// Fields of the NovelComment.
func (NovelComment) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("novelId").Optional().
			Comment(`小说编号`),
		field.Int64("userId").Optional().
			Comment(`用户Id`),
		field.String("avatar").Optional().
			Comment(`评论用户头像`),
		field.String("userName").Optional().
			Comment(`用户名`),
		field.Int64("repUserId").Optional().
			Comment(`回复用户`),
		field.String("repUserName").Optional().
			Comment(`回复用户ID`),
		field.String("content").Optional().
			Comment(`回复内容`),
		field.Int("score").Optional().
			Comment(`评分,除以10`),
		field.Int64("pId").Optional().
			Comment(`回复评论ID`),
		field.Bool("isTop").Optional().
			Comment(`置顶`),
		field.Enum("state").Values("0", "1", "2").Optional().
			Comment(`状态,0`),
		field.Bool("isHighlight").Optional().
			Comment(`高亮`),
		field.Bool("isHot").Optional().
			Comment(`是否热门`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the NovelComment.
func (NovelComment) Edges() []ent.Edge {
	return []ent.Edge{}
}
