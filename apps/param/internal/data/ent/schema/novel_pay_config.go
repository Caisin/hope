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
		field.Int64("id").
			Comment(`主键编码`),
		field.String("productId").Optional().
			Comment(`谷歌商品ID`),
		field.String("paymentName").Optional().
			Comment(`商品名称`),
		field.Int64("firstPayment").Optional().
			Comment(`首次充值金额`),
		field.Int64("payment").Optional().
			Comment(`充值金额 分`),
		field.Int64("oldPrice").Optional().
			Comment(`原价`),
		field.String("cfgType").Optional().
			Comment(`配置类型,activity`),
		field.Int64("coin").Optional().
			Comment(`书币`),
		field.String("currency").Optional().
			Comment(`货币类型`),
		field.Int64("coupon").Optional().
			Comment(`书券`),
		field.Int("coinItem").Optional().
			Comment(`现金科目,查看asset_item表cash_flag=1`),
		field.Int("couponItem").Optional().
			Comment(`赠送科目,查看asset_item表cash_flag=0`),
		field.Int("sort").Optional().
			Comment(`排序`),
		field.Bool("state").Optional().
			Comment(`状态 1`),
		field.Int("isSend").Optional().
			Comment(`巨量引擎是否回传`),
		field.Int("payType").Optional().
			Comment(`充值类型,0`),
		field.Int64("vipType").Optional().
			Comment(`vip类型`),
		field.Bool("isHot").Optional().
			Comment(`展示火,0`),
		field.Int("cycleDay").Optional().
			Comment(`扣款周期天数,有这个值就周期扣款`),
		field.Int("width").Optional().
			Comment(`6`),
		field.String("summary").Optional().
			Comment(`描述`),
		field.String("remark").Optional().
			Comment(`备注`),
		field.Time("effectTime").Optional().
			Comment(`生效时间`),
		field.Time("expiredTime").Optional().
			Comment(`失效时间`),
	}
	fields = append(fields, mixin.HopeMixin{}.Fields()...)
	return fields
}

// Edges of the NovelPayConfig.
func (NovelPayConfig) Edges() []ent.Edge {
	return []ent.Edge{}
}
