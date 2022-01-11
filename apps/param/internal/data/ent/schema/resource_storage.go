package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// ResourceStorage holds the schema definition for the ResourceStorage entity.
type ResourceStorage struct {
	ent.Schema
}

// Fields of the ResourceStorage.
func (ResourceStorage) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int("groupId").Optional().
			Comment(`分组`),
		field.Int("storageType").Optional().
			Comment(`存储类型 1`),
		field.String("realName").Optional().
			Comment(`文件真实的名称`),
		field.String("bucket").Optional().
			Comment(`Bucket 识别符(七牛)`),
		field.String("name").Optional().
			Comment(`文件名称`),
		field.String("suffix").Optional().
			Comment(`后缀`),
		field.String("path").Optional().
			Comment(`路径`),
		field.String("type").Optional().
			Comment(`类型`),
		field.String("size").Optional().
			Comment(`大小`),
		field.String("deleteUrl").Optional().
			Comment(`(sm.sm)删除的URL`),
		field.String("filename").Optional().
			Comment(`(sm.sm)图片名称`),
		field.String("key").Optional().
			Comment(`文件名(七牛)`),
		field.String("height").Optional().
			Comment(`(sm.sm)图片高度`),
		field.String("url").Optional().
			Comment(`(sm.sm)图片地址`),
		field.String("username").Optional().
			Comment(`(sm.sm)用户名称`),
		field.String("width").Optional().
			Comment(`(sm.sm)图片宽度`),
		field.String("md5code").Optional().
			Comment(`(sm.sm)文件的MD5值`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the ResourceStorage.
func (ResourceStorage) Edges() []ent.Edge {
	return []ent.Edge{}
}
