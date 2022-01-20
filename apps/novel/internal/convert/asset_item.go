// Code generated by Caisin. DO NOT EDIT.
// source: apps/novel/internal/data/ent/schema/asset_item.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/assetitem/v1"
	"hope/apps/novel/internal/data/ent"
)

func AssetItemUpdateReq2Data(v *v1.AssetItemUpdateReq) *ent.AssetItem {
	if v == nil {
		return nil
	}
	return &ent.AssetItem{
		ID:          v.Id,
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func AssetItemData2UpdateReq(v *ent.AssetItem) *v1.AssetItemUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.AssetItemUpdateReq{
		Id:          v.ID,
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func AssetItemCreateReq2Data(v *v1.AssetItemCreateReq) *ent.AssetItem {
	if v == nil {
		return nil
	}
	return &ent.AssetItem{
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func AssetItemData2CreateReq(v *ent.AssetItem) *v1.AssetItemCreateReq {
	if v == nil {
		return nil
	}
	return &v1.AssetItemCreateReq{
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func AssetItemReq2Data(v *v1.AssetItemReq) *ent.AssetItem {
	if v == nil {
		return nil
	}
	return &ent.AssetItem{
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func AssetItemData2Req(v *ent.AssetItem) *v1.AssetItemReq {
	if v == nil {
		return nil
	}
	return &v1.AssetItemReq{
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func AssetItemReply2Data(v *v1.AssetItemReply) *ent.AssetItem {
	if v == nil {
		return nil
	}
	return &ent.AssetItem{
		ID:          v.Id,
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  v.EffectTime.AsTime(),
		ExpiredTime: v.ExpiredTime.AsTime(),
		CreatedAt:   v.CreatedAt.AsTime(),
		UpdatedAt:   v.UpdatedAt.AsTime(),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}

func AssetItemData2Reply(v *ent.AssetItem) *v1.AssetItemReply {
	if v == nil {
		return nil
	}
	return &v1.AssetItemReply{
		Id:          v.ID,
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  timestamppb.New(v.EffectTime),
		ExpiredTime: timestamppb.New(v.ExpiredTime),
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}

func AssetItemUpdateReply2Data(v *v1.AssetItemUpdateReply) *ent.AssetItem {
	if v == nil {
		return nil
	}
	return &ent.AssetItem{
		ID:          v.Id,
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func AssetItemData2UpdateReply(v *ent.AssetItem) *v1.AssetItemUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.AssetItemUpdateReply{
		Id:          v.ID,
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func AssetItemCreateReply2Data(v *v1.AssetItemCreateReply) *ent.AssetItem {
	if v == nil {
		return nil
	}
	return &ent.AssetItem{
		ID:          v.Id,
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  v.EffectTime.AsTime(),
		ExpiredTime: v.ExpiredTime.AsTime(),
		CreatedAt:   v.CreatedAt.AsTime(),
		UpdatedAt:   v.UpdatedAt.AsTime(),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}

func AssetItemData2CreateReply(v *ent.AssetItem) *v1.AssetItemCreateReply {
	if v == nil {
		return nil
	}
	return &v1.AssetItemCreateReply{
		Id:          v.ID,
		AssetItemId: v.AssetItemId,
		AssetName:   v.AssetName,
		CashTag:     v.CashTag,
		ValidDays:   v.ValidDays,
		EffectTime:  timestamppb.New(v.EffectTime),
		ExpiredTime: timestamppb.New(v.ExpiredTime),
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}
