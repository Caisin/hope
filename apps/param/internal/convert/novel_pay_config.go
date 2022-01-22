// Code generated by Caisin. DO NOT EDIT.
// source: apps/param/internal/data/ent/schema/novel_pay_config.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/param/novelpayconfig/v1"
	"hope/apps/param/internal/data/ent"
)

func NovelPayConfigUpdateReq2Data(v *v1.NovelPayConfigUpdateReq) *ent.NovelPayConfig {
	if v == nil {
		return nil
	}
	return &ent.NovelPayConfig{
		ID:            v.Id,
		ProductId:     v.ProductId,
		PaymentName:   v.PaymentName,
		FirstPayment:  v.FirstPayment,
		Payment:       v.Payment,
		OriginalPrice: v.OriginalPrice,
		CfgType:       v.CfgType,
		Coin:          v.Coin,
		Currency:      v.Currency,
		Coupon:        v.Coupon,
		CoinItem:      v.CoinItem,
		CouponItem:    v.CouponItem,
		Sort:          v.Sort,
		State:         v.State,
		IsSend:        v.IsSend,
		PayType:       v.PayType,
		VipType:       v.VipType,
		IsHot:         v.IsHot,
		CycleDay:      v.CycleDay,
		Summary:       v.Summary,
		Remark:        v.Remark,
		EffectTime:    v.EffectTime.AsTime(),
	}
}

func NovelPayConfigData2UpdateReq(v *ent.NovelPayConfig) *v1.NovelPayConfigUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.NovelPayConfigUpdateReq{
		Id:            v.ID,
		ProductId:     v.ProductId,
		PaymentName:   v.PaymentName,
		FirstPayment:  v.FirstPayment,
		Payment:       v.Payment,
		OriginalPrice: v.OriginalPrice,
		CfgType:       v.CfgType,
		Coin:          v.Coin,
		Currency:      v.Currency,
		Coupon:        v.Coupon,
		CoinItem:      v.CoinItem,
		CouponItem:    v.CouponItem,
		Sort:          v.Sort,
		State:         v.State,
		IsSend:        v.IsSend,
		PayType:       v.PayType,
		VipType:       v.VipType,
		IsHot:         v.IsHot,
		CycleDay:      v.CycleDay,
		Summary:       v.Summary,
		Remark:        v.Remark,
		EffectTime:    timestamppb.New(v.EffectTime),
	}
}

func NovelPayConfigCreateReq2Data(v *v1.NovelPayConfigCreateReq) *ent.NovelPayConfig {
	if v == nil {
		return nil
	}
	return &ent.NovelPayConfig{
		ProductId:     v.ProductId,
		PaymentName:   v.PaymentName,
		FirstPayment:  v.FirstPayment,
		Payment:       v.Payment,
		OriginalPrice: v.OriginalPrice,
		CfgType:       v.CfgType,
		Coin:          v.Coin,
		Currency:      v.Currency,
		Coupon:        v.Coupon,
		CoinItem:      v.CoinItem,
		CouponItem:    v.CouponItem,
		Sort:          v.Sort,
		State:         v.State,
		IsSend:        v.IsSend,
		PayType:       v.PayType,
		VipType:       v.VipType,
		IsHot:         v.IsHot,
		CycleDay:      v.CycleDay,
		Summary:       v.Summary,
		Remark:        v.Remark,
		EffectTime:    v.EffectTime.AsTime(),
	}
}

func NovelPayConfigData2CreateReq(v *ent.NovelPayConfig) *v1.NovelPayConfigCreateReq {
	if v == nil {
		return nil
	}
	return &v1.NovelPayConfigCreateReq{
		ProductId:     v.ProductId,
		PaymentName:   v.PaymentName,
		FirstPayment:  v.FirstPayment,
		Payment:       v.Payment,
		OriginalPrice: v.OriginalPrice,
		CfgType:       v.CfgType,
		Coin:          v.Coin,
		Currency:      v.Currency,
		Coupon:        v.Coupon,
		CoinItem:      v.CoinItem,
		CouponItem:    v.CouponItem,
		Sort:          v.Sort,
		State:         v.State,
		IsSend:        v.IsSend,
		PayType:       v.PayType,
		VipType:       v.VipType,
		IsHot:         v.IsHot,
		CycleDay:      v.CycleDay,
		Summary:       v.Summary,
		Remark:        v.Remark,
		EffectTime:    timestamppb.New(v.EffectTime),
	}
}

func NovelPayConfigReq2Data(v *v1.NovelPayConfigReq) *ent.NovelPayConfig {
	if v == nil {
		return nil
	}
	return &ent.NovelPayConfig{
		ProductId:     v.ProductId,
		PaymentName:   v.PaymentName,
		FirstPayment:  v.FirstPayment,
		Payment:       v.Payment,
		OriginalPrice: v.OriginalPrice,
		CfgType:       v.CfgType,
		Coin:          v.Coin,
		Currency:      v.Currency,
		Coupon:        v.Coupon,
		CoinItem:      v.CoinItem,
		CouponItem:    v.CouponItem,
		Sort:          v.Sort,
		State:         v.State,
		IsSend:        v.IsSend,
		PayType:       v.PayType,
		VipType:       v.VipType,
		IsHot:         v.IsHot,
		CycleDay:      v.CycleDay,
		Summary:       v.Summary,
		Remark:        v.Remark,
		EffectTime:    v.EffectTime.AsTime(),
	}
}

func NovelPayConfigData2Req(v *ent.NovelPayConfig) *v1.NovelPayConfigReq {
	if v == nil {
		return nil
	}
	return &v1.NovelPayConfigReq{
		ProductId:     v.ProductId,
		PaymentName:   v.PaymentName,
		FirstPayment:  v.FirstPayment,
		Payment:       v.Payment,
		OriginalPrice: v.OriginalPrice,
		CfgType:       v.CfgType,
		Coin:          v.Coin,
		Currency:      v.Currency,
		Coupon:        v.Coupon,
		CoinItem:      v.CoinItem,
		CouponItem:    v.CouponItem,
		Sort:          v.Sort,
		State:         v.State,
		IsSend:        v.IsSend,
		PayType:       v.PayType,
		VipType:       v.VipType,
		IsHot:         v.IsHot,
		CycleDay:      v.CycleDay,
		Summary:       v.Summary,
		Remark:        v.Remark,
		EffectTime:    timestamppb.New(v.EffectTime),
	}
}

func NovelPayConfigData2Reply(v *ent.NovelPayConfig) *v1.NovelPayConfigData {
	if v == nil {
		return nil
	}
	return &v1.NovelPayConfigData{
		Id:            v.ID,
		ProductId:     v.ProductId,
		PaymentName:   v.PaymentName,
		FirstPayment:  v.FirstPayment,
		Payment:       v.Payment,
		OriginalPrice: v.OriginalPrice,
		CfgType:       v.CfgType,
		Coin:          v.Coin,
		Currency:      v.Currency,
		Coupon:        v.Coupon,
		CoinItem:      v.CoinItem,
		CouponItem:    v.CouponItem,
		Sort:          v.Sort,
		State:         v.State,
		IsSend:        v.IsSend,
		PayType:       v.PayType,
		VipType:       v.VipType,
		IsHot:         v.IsHot,
		CycleDay:      v.CycleDay,
		Summary:       v.Summary,
		Remark:        v.Remark,
		EffectTime:    timestamppb.New(v.EffectTime),
		ExpiredTime:   timestamppb.New(v.ExpiredTime),
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}
