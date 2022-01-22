// Code generated by Caisin. DO NOT EDIT.
// source: apps/param/internal/data/ent/schema/score_product.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/param/scoreproduct/v1"
	"hope/apps/param/internal/data/ent"
)

func ScoreProductUpdateReq2Data(v *v1.ScoreProductUpdateReq) *ent.ScoreProduct {
	if v == nil {
		return nil
	}
	return &ent.ScoreProduct{
		ID:          v.Id,
		ProductName: v.ProductName,
		Summary:     v.Summary,
		CardUrl:     v.CardUrl,
		Score:       v.Score,
		VipType:     v.VipType,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func ScoreProductData2UpdateReq(v *ent.ScoreProduct) *v1.ScoreProductUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.ScoreProductUpdateReq{
		Id:          v.ID,
		ProductName: v.ProductName,
		Summary:     v.Summary,
		CardUrl:     v.CardUrl,
		Score:       v.Score,
		VipType:     v.VipType,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func ScoreProductCreateReq2Data(v *v1.ScoreProductCreateReq) *ent.ScoreProduct {
	if v == nil {
		return nil
	}
	return &ent.ScoreProduct{
		ProductName: v.ProductName,
		Summary:     v.Summary,
		CardUrl:     v.CardUrl,
		Score:       v.Score,
		VipType:     v.VipType,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func ScoreProductData2CreateReq(v *ent.ScoreProduct) *v1.ScoreProductCreateReq {
	if v == nil {
		return nil
	}
	return &v1.ScoreProductCreateReq{
		ProductName: v.ProductName,
		Summary:     v.Summary,
		CardUrl:     v.CardUrl,
		Score:       v.Score,
		VipType:     v.VipType,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func ScoreProductReq2Data(v *v1.ScoreProductReq) *ent.ScoreProduct {
	if v == nil {
		return nil
	}
	return &ent.ScoreProduct{
		ProductName: v.ProductName,
		Summary:     v.Summary,
		CardUrl:     v.CardUrl,
		Score:       v.Score,
		VipType:     v.VipType,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func ScoreProductData2Req(v *ent.ScoreProduct) *v1.ScoreProductReq {
	if v == nil {
		return nil
	}
	return &v1.ScoreProductReq{
		ProductName: v.ProductName,
		Summary:     v.Summary,
		CardUrl:     v.CardUrl,
		Score:       v.Score,
		VipType:     v.VipType,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func ScoreProductData2Reply(v *ent.ScoreProduct) *v1.ScoreProductData {
	if v == nil {
		return nil
	}
	return &v1.ScoreProductData{
		Id:          v.ID,
		ProductName: v.ProductName,
		Summary:     v.Summary,
		CardUrl:     v.CardUrl,
		Score:       v.Score,
		VipType:     v.VipType,
		EffectTime:  timestamppb.New(v.EffectTime),
		ExpiredTime: timestamppb.New(v.ExpiredTime),
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}
