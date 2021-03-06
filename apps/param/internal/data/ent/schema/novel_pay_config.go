package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hope/pkg/ent/mixin"
)

// NovelPayConfig holds the schema definition for the NovelPayConfig entity.
type NovelPayConfig struct {
	ent.Schema
}

// Fields of the NovelPayConfig.
func (NovelPayConfig) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("productId").Optional().
			Comment(`谷歌商品ID`),
		field.String("paymentName").Optional().
			Comment(`商品名称`),
		field.Int64("firstPayment").Optional().
			Comment(`首次充值金额`),
		field.Int64("payment").Optional().
			Comment(`充值金额:分`),
		field.Int64("originalPrice").Optional().
			Comment(`原价`),
		field.String("cfgType").Optional().
			Comment(`配置类型,activity`),
		field.Int64("coin").Optional().
			Comment(`书币`),
		field.String("currency").Optional().
			Comment(`货币类型`),
		field.Int64("coupon").Optional().
			Comment(`书券`),
		field.Int32("coinItem").Optional().
			Comment(`现金科目,查看asset_item表cash_flag=1`),
		field.Int32("couponItem").Optional().
			Comment(`赠送科目,查看asset_item表cash_flag=0`),
		field.Int32("sort").Optional().
			Comment(`排序`),
		field.Bool("state").Optional().
			Comment(`状态`),
		field.Int32("isSend").Optional().
			Comment(`巨量引擎是否回传`),
		field.Int32("payType").Optional().
			Comment(`充值类型`),
		field.Int64("vipType").Optional().
			Comment(`vip类型`),
		field.Bool("isHot").Optional().
			Comment(`展示火`),
		field.Int32("cycleDay").Optional().
			Comment(`扣款周期天数,有这个值就周期扣款`),
		field.String("summary").Optional().
			Comment(`描述`),
		field.String("remark").Optional().
			Comment(`备注`),
	}
	fields = append(fields, mixin.EETimeFields()...)
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the NovelPayConfig.
func (NovelPayConfig) Edges() []ent.Edge {
	return []ent.Edge{}
}
