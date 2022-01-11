package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// AdChannel holds the schema definition for the AdChannel entity.
type AdChannel struct {
	ent.Schema
}

// Fields of the AdChannel.
func (AdChannel) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Unique().
			Comment(`主键编码`),
		field.String("channelName").Optional().
			Comment(`渠道名称`),
		field.Int64("novelId").Optional().
			Comment(`书号`),
		field.Int64("reg").Optional().
			Comment(`充值回传比例`),
		field.Int64("pay").Optional().
			Comment(`付费回传比例`),
		field.String("novelName").Optional().
			Comment(`书名`),
		field.Int64("chapterId").Optional().
			Comment(`章节号`),
		field.Int("chapterNum").Optional().
			Comment(`章节数`),
		field.String("adType").Optional().
			Comment(`投放渠道`),
		field.String("img").Optional().
			Comment(`落地页图片uri`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the AdChannel.
func (AdChannel) Edges() []ent.Edge {
	return []ent.Edge{}
}
