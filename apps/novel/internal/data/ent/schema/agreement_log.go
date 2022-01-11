package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// AgreementLog holds the schema definition for the AgreementLog entity.
type AgreementLog struct {
	ent.Schema
}

// Fields of the AgreementLog.
func (AgreementLog) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("agreementNo").
			Comment(`协议号`),
		field.String("outerAgreementNo").Optional().
			Comment(`外部签约协议号`),
		field.String("orderId").Optional().
			Comment(`订单号`),
		field.Int64("userId").Optional().
			Comment(`用户ID`),
		field.Int64("chId").Optional().
			Comment(`渠道`),
		field.String("userName").Optional().
			Comment(`用户名称`),
		field.String("paymentName").Optional().
			Comment(`商品名称`),
		field.Int64("paymentId").Optional().
			Comment(`支付配置ID payment_config表的ID`),
		field.Int("state").Optional().
			Comment(`状态：1、待签约，2、签约成功，3、取消签约`),
		field.Int64("payment").Optional().
			Comment(`每期扣款金额,单位：分`),
		field.Enum("agreementType").Values("Alipay", "Google", "WeChat", "Apple").Optional().
			Comment(`支付方式：Alipay,Google,WeChat`),
		field.Int64("vipType").Optional().
			Comment(`vip类型`),
		field.Int64("times").Optional().
			Comment(`第几期`),
		field.Int("cycleDays").Optional().
			Comment(`周期天数`),
		field.Time("nextExecTime").Optional().
			Comment(`下次执行时间`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the AgreementLog.
func (AgreementLog) Edges() []ent.Edge {
	return []ent.Edge{}
}
