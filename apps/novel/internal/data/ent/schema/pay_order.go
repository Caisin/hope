package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"hope/pkg/ent/mixin"
)

// PayOrder holds the schema definition for the PayOrder entity.
type PayOrder struct {
	ent.Schema
}

func (PayOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.HopeMixin{},
	}
}

type OrderState int32

const (
	ToBePaid OrderState = iota + 1 //待支付
	Payed                          //已支付
	Cancel                         //取消
	Refund                         //退款
)

func (p OrderState) String() string {
	s := "unknown"
	switch p {
	case ToBePaid:
		s = "待支付"
	case Payed:
		s = "已支付"
	case Cancel:
		s = "取消"
	case Refund:
		s = "退款"
	}
	return s
}

func (p OrderState) Validate() error {
	switch p {
	case ToBePaid, Payed, Cancel, Refund:
		return nil
	default:
		return fmt.Errorf("invalid orderState value: %v", p)
	}
}

// Fields of the PayOrder.
func (PayOrder) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("orderId").Unique().
			Comment("订单号"),
		field.Int64("userId").
			Comment(`用户ID`),
		field.Int64("chId").
			Comment(`渠道ID`),
		field.Int64("agreementId").
			Comment(`签约协议号`),
		field.String("lastRead").
			Comment("最后阅读书籍"),
		field.String("lastChapter").
			Comment("最后阅读章节"),
		field.String("paymentName").
			Comment("充值配置名称"),
		field.String("paymentId").
			Comment("充值配置ID"),
		field.Int32("state").
			GoType(OrderState(0)).
			Default(int32(ToBePaid)).
			Validate(func(i int32) error {
				return OrderState(i).
					Validate()
			}).
			Comment("订单状态:1、待付款，2、已付款，3、取消，4、退款"),
		field.Int64("payment").
			Min(0).
			Comment("充值金额"),
		field.Time("paymentTime").
			Comment("支付时间"),
		field.Time("closeTime").
			Comment("订单关闭时间"),
		field.Enum("payType").
			Values("Google", "WeChat", "Alipay", "Apple").
			Comment("支付类型"),
		field.Int64("coin").
			Comment("书币"),
		field.Int64("coupon").
			Comment("书券"),
		field.String("vipDays").
			Optional().
			Comment("vip天数"),
		field.Int64("vipType").
			Optional().
			Comment("vip类型"),
		field.String("vipName").
			Optional().
			Comment("vip名称"),
		field.Int32("times").
			Comment("充值次数"),
		field.String("otherOrderId").
			Optional().
			Comment("第三方订单"),
		field.String("remark").
			Optional().
			Comment("备注"),
	}
	fields = append(fields, mixin.Fields()...)
	return fields
}

// Edges of the PayOrder.
func (PayOrder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("user", SocialUser.Type).Field("userId").Required().Comment("所属用户").
			Ref("orders").
			Unique(),
		edge.
			From("channel", AdChannel.Type).Field("chId").Required().Comment("进入渠道").
			Ref("orders").
			Unique(),
		edge.
			From("agreement", AgreementLog.Type).Field("agreementId").Required().Comment("自动扣费协议").
			Ref("orders").
			Unique(),
	}
}
