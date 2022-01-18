package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// Novel holds the schema definition for the Novel entity.
type Novel struct {
	ent.Schema
}

// Fields of the Novel.
func (Novel) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("novelId").
			Comment(`小说编码`),
		field.Int64("classifyId").Optional().
			Comment(`分类ID`),
		field.String("classifyName").Optional().
			Comment(`分类名称`),
		field.String("authorId").Optional().
			Comment(`作者ID,本站小说有`),
		field.String("title").Optional().
			Comment(`小说标题`),
		field.String("summary").Optional().
			Comment(`小说摘要`),
		field.String("author").Optional().
			Comment(`作者`),
		field.String("anchor").Optional().
			Comment(`主播`),
		field.Int64("hits").Optional().
			Comment(`点击量`),
		field.String("keywords").Optional().
			Comment(`关键字，用逗号隔开`),
		field.String("source").Optional().
			Comment(`来源`),
		field.Int32("score").Optional().
			Comment(`评分`),
		field.String("cover").Optional().
			Comment(`封面`),
		field.String("tagIds").Optional().
			Comment(`标签,关联标签表novel_tag`),
		field.Int32("wordNum").Optional().
			Comment(`书本字数`),
		field.Int32("freeNum").Optional().
			Comment(`免费章节数量`),
		field.Int32("onlineState").Optional().
			Comment(`连载状态`),
		field.Int64("price").Optional().
			Comment(`整本价格`),
		field.Int32("publish").Optional().
			Comment(`发布状态`),
		field.Int64("originalPrice").Optional().
			Comment(`原价,展示用`),
		field.Int32("chapterPrice").Optional().
			Comment(`千字价格`),
		field.Int32("chapterCount").Optional().
			Comment(`章节数量`),
		field.Int32("signType").Optional().
			Comment(`签约类型 0`),
		field.Time("signDate").Optional().
			Comment(`签约时间`),
		field.String("leadingMan").Optional().
			Comment(`男主角`),
		field.String("leadingLady").Optional().
			Comment(`女主角`),
		field.String("remark").Optional().
			Comment(`备注`),
		field.String("mediaKey").Optional().
			Comment(`阿里云音频目录`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the Novel.
func (Novel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chapters", NovelChapter.Type).Comment("章节列表"),
		edge.To("pkgs", BookPackage.Type).Comment("打包购买包"),
	}
}
