package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// SysOperaLog holds the schema definition for the SysOperaLog entity.
type SysOperaLog struct {
	ent.Schema
}

// Fields of the SysOperaLog.
func (SysOperaLog) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").
			Comment(`主键编码`),
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.String("title").Optional().
			Comment(`操作模块`),
		field.String("requestId").Optional().
			Comment(`请求ID`),
		field.String("businessType").Optional().
			Comment(`操作类型`),
		field.String("businessTypes").Optional().
			Comment(`BusinessTypes`),
		field.String("method").Optional().
			Comment(`函数`),
		field.String("requestMethod").Optional().
			Comment(`请求方式`),
		field.String("operatorType").Optional().
			Comment(`操作类型`),
		field.String("operName").Optional().
			Comment(`操作者`),
		field.String("deptName").Optional().
			Comment(`部门名称`),
		field.String("operUrl").Optional().
			Comment(`访问地址`),
		field.String("operIp").Optional().
			Comment(`客户端ip`),
		field.String("browser").Optional().
			Comment(`浏览器`),
		field.String("os").Optional().
			Comment(`系统`),
		field.String("platform").Optional().
			Comment(`固件`),
		field.String("operLocation").Optional().
			Comment(`访问位置`),
		field.String("operParam").Optional().
			Comment(`请求参数`),
		field.String("status").Optional().
			Comment(`操作状态`),
		field.Time("operTime").Optional().
			Comment(`操作时间`),
		field.String("jsonResult").Optional().
			Comment(`返回数据`),
		field.String("remark").Optional().
			Comment(`备注`),
		field.String("latencyTime").Optional().
			Comment(`耗时`),
		field.String("userAgent").Optional().
			Comment(`ua`),
	}
	fields = append(fields, mixin.TimeMixin{}.Fields()...)
	return fields
}

// Edges of the SysOperaLog.
func (SysOperaLog) Edges() []ent.Edge {
	return []ent.Edge{}
}
