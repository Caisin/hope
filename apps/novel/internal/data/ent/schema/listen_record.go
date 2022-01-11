package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
	"time"
)

// ListenRecord holds the schema definition for the ListenRecord entity.
type ListenRecord struct {
	ent.Schema
}

// Fields of the ListenRecord.
func (ListenRecord) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("userId").
			Comment(`用户ID`),
		field.Int64("chapterId").
			Comment(`主键编码`),
		field.Int64("novelId").Optional().
			Comment(`小说编号`),
		field.Int64("listenTimes").Optional().
			Comment(`播放次数`),
		field.Int64("duration").GoType(time.Duration(0)).Optional().
			Comment(`音频时长`),
		field.Int64("allDuration").GoType(time.Duration(0)).Optional().
			Comment(`总播放时长`),
		field.Int64("dayDuration").GoType(time.Duration(0)).Optional().
			Comment(`当天播放时长`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the ListenRecord.
func (ListenRecord) Edges() []ent.Edge {
	return []ent.Edge{}
}
