package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// DataSource holds the schema definition for the DataSource entity.
type DataSource struct {
	ent.Schema
}

// Fields of the DataSource.
func (DataSource) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.String("dbName").Optional().
			Comment(`数据源名称`),
		field.String("host").Optional().
			Comment(`数据源名称`),
		field.Int("port").Optional().
			Comment(`数据源名称`),
		field.String("database").Optional().
			Comment(`数据源名称`),
		field.String("userName").Optional().
			Comment(`用户名`),
		field.String("pwd").Optional().
			Comment(`密码`),
		field.Bool("status").Optional().
			Comment(`状态`),
		field.String("driver").Optional().
			Comment(`数据库类型`),
		field.Int("connMaxIdleTime").Optional().
			Comment(`最大空闲连接数`),
		field.Int("connMaxLifeTime").Optional().
			Comment(`连接可重用的最大时间长度`),
		field.Int("maxIdleConns").Optional().
			Comment(`最大空闲`),
		field.Int("maxOpenConns").Optional().
			Comment(`最大打开连接数`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the DataSource.
func (DataSource) Edges() []ent.Edge {
	return []ent.Edge{}
}
