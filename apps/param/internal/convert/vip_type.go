// Code generated by Caisin. DO NOT EDIT.
// source: apps/param/internal/data/ent/schema/vip_type.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/param/viptype/v1"
	"hope/apps/param/internal/data/ent"
)

func VipTypeUpdateReq2Data(v *v1.VipTypeUpdateReq) *ent.VipType {
	if v == nil {
		return nil
	}
	return &ent.VipType{
		ID:           v.Id,
		VipName:      v.VipName,
		IsSuper:      v.IsSuper,
		ValidDays:    v.ValidDays,
		DiscountRate: v.DiscountRate,
		AvatarId:     v.AvatarId,
	}
}

func VipTypeData2UpdateReq(v *ent.VipType) *v1.VipTypeUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.VipTypeUpdateReq{
		Id:           v.ID,
		VipName:      v.VipName,
		IsSuper:      v.IsSuper,
		ValidDays:    v.ValidDays,
		DiscountRate: v.DiscountRate,
		AvatarId:     v.AvatarId,
	}
}

func VipTypeCreateReq2Data(v *v1.VipTypeCreateReq) *ent.VipType {
	if v == nil {
		return nil
	}
	return &ent.VipType{
		VipName:      v.VipName,
		IsSuper:      v.IsSuper,
		ValidDays:    v.ValidDays,
		DiscountRate: v.DiscountRate,
		AvatarId:     v.AvatarId,
	}
}

func VipTypeData2CreateReq(v *ent.VipType) *v1.VipTypeCreateReq {
	if v == nil {
		return nil
	}
	return &v1.VipTypeCreateReq{
		VipName:      v.VipName,
		IsSuper:      v.IsSuper,
		ValidDays:    v.ValidDays,
		DiscountRate: v.DiscountRate,
		AvatarId:     v.AvatarId,
	}
}

func VipTypeReq2Data(v *v1.VipTypeReq) *ent.VipType {
	if v == nil {
		return nil
	}
	return &ent.VipType{
		VipName:      v.VipName,
		IsSuper:      v.IsSuper,
		ValidDays:    v.ValidDays,
		DiscountRate: v.DiscountRate,
		AvatarId:     v.AvatarId,
	}
}

func VipTypeData2Req(v *ent.VipType) *v1.VipTypeReq {
	if v == nil {
		return nil
	}
	return &v1.VipTypeReq{
		VipName:      v.VipName,
		IsSuper:      v.IsSuper,
		ValidDays:    v.ValidDays,
		DiscountRate: v.DiscountRate,
		AvatarId:     v.AvatarId,
	}
}

func VipTypeData2Reply(v *ent.VipType) *v1.VipTypeData {
	if v == nil {
		return nil
	}
	return &v1.VipTypeData{
		Id:           v.ID,
		VipName:      v.VipName,
		IsSuper:      v.IsSuper,
		ValidDays:    v.ValidDays,
		DiscountRate: v.DiscountRate,
		AvatarId:     v.AvatarId,
		Summary:      v.Summary,
		CreatedAt:    timestamppb.New(v.CreatedAt),
		UpdatedAt:    timestamppb.New(v.UpdatedAt),
		CreateBy:     v.CreateBy,
		UpdateBy:     v.UpdateBy,
		TenantId:     v.TenantId,
	}
}
