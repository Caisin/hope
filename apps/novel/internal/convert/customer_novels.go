// Code generated by Caisin. DO NOT EDIT.
// source: apps/novel/internal/data/ent/schema/customer_novels.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/customernovels/v1"
	"hope/apps/novel/internal/data/ent"
)

func CustomerNovelsUpdateReq2Data(v *v1.CustomerNovelsUpdateReq) *ent.CustomerNovels {
	if v == nil {
		return nil
	}
	return &ent.CustomerNovels{
		ID:         v.Id,
		NovelId:    v.NovelId,
		TypeId:     v.TypeId,
		TypeCode:   v.TypeCode,
		GroupCode:  v.GroupCode,
		FieldName:  v.FieldName,
		Cover:      v.Cover,
		OrderNum:   v.OrderNum,
		Remark:     v.Remark,
		EffectTime: v.EffectTime.AsTime(),
	}
}

func CustomerNovelsData2UpdateReq(v *ent.CustomerNovels) *v1.CustomerNovelsUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.CustomerNovelsUpdateReq{
		Id:         v.ID,
		NovelId:    v.NovelId,
		TypeId:     v.TypeId,
		TypeCode:   v.TypeCode,
		GroupCode:  v.GroupCode,
		FieldName:  v.FieldName,
		Cover:      v.Cover,
		OrderNum:   v.OrderNum,
		Remark:     v.Remark,
		EffectTime: timestamppb.New(v.EffectTime),
	}
}

func CustomerNovelsCreateReq2Data(v *v1.CustomerNovelsCreateReq) *ent.CustomerNovels {
	if v == nil {
		return nil
	}
	return &ent.CustomerNovels{
		NovelId:    v.NovelId,
		TypeId:     v.TypeId,
		TypeCode:   v.TypeCode,
		GroupCode:  v.GroupCode,
		FieldName:  v.FieldName,
		Cover:      v.Cover,
		OrderNum:   v.OrderNum,
		Remark:     v.Remark,
		EffectTime: v.EffectTime.AsTime(),
	}
}

func CustomerNovelsData2CreateReq(v *ent.CustomerNovels) *v1.CustomerNovelsCreateReq {
	if v == nil {
		return nil
	}
	return &v1.CustomerNovelsCreateReq{
		NovelId:    v.NovelId,
		TypeId:     v.TypeId,
		TypeCode:   v.TypeCode,
		GroupCode:  v.GroupCode,
		FieldName:  v.FieldName,
		Cover:      v.Cover,
		OrderNum:   v.OrderNum,
		Remark:     v.Remark,
		EffectTime: timestamppb.New(v.EffectTime),
	}
}

func CustomerNovelsReq2Data(v *v1.CustomerNovelsReq) *ent.CustomerNovels {
	if v == nil {
		return nil
	}
	return &ent.CustomerNovels{
		NovelId:    v.NovelId,
		TypeId:     v.TypeId,
		TypeCode:   v.TypeCode,
		GroupCode:  v.GroupCode,
		FieldName:  v.FieldName,
		Cover:      v.Cover,
		OrderNum:   v.OrderNum,
		Remark:     v.Remark,
		EffectTime: v.EffectTime.AsTime(),
	}
}

func CustomerNovelsData2Req(v *ent.CustomerNovels) *v1.CustomerNovelsReq {
	if v == nil {
		return nil
	}
	return &v1.CustomerNovelsReq{
		NovelId:    v.NovelId,
		TypeId:     v.TypeId,
		TypeCode:   v.TypeCode,
		GroupCode:  v.GroupCode,
		FieldName:  v.FieldName,
		Cover:      v.Cover,
		OrderNum:   v.OrderNum,
		Remark:     v.Remark,
		EffectTime: timestamppb.New(v.EffectTime),
	}
}

func CustomerNovelsReply2Data(v *v1.CustomerNovelsReply) *ent.CustomerNovels {
	if v == nil {
		return nil
	}
	return &ent.CustomerNovels{
		ID:          v.Id,
		NovelId:     v.NovelId,
		TypeId:      v.TypeId,
		TypeCode:    v.TypeCode,
		GroupCode:   v.GroupCode,
		FieldName:   v.FieldName,
		Cover:       v.Cover,
		OrderNum:    v.OrderNum,
		Remark:      v.Remark,
		EffectTime:  v.EffectTime.AsTime(),
		ExpiredTime: v.ExpiredTime.AsTime(),
		CreatedAt:   v.CreatedAt.AsTime(),
		UpdatedAt:   v.UpdatedAt.AsTime(),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}

func CustomerNovelsData2Reply(v *ent.CustomerNovels) *v1.CustomerNovelsReply {
	if v == nil {
		return nil
	}
	return &v1.CustomerNovelsReply{
		Id:          v.ID,
		NovelId:     v.NovelId,
		TypeId:      v.TypeId,
		TypeCode:    v.TypeCode,
		GroupCode:   v.GroupCode,
		FieldName:   v.FieldName,
		Cover:       v.Cover,
		OrderNum:    v.OrderNum,
		Remark:      v.Remark,
		EffectTime:  timestamppb.New(v.EffectTime),
		ExpiredTime: timestamppb.New(v.ExpiredTime),
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}

func CustomerNovelsUpdateReply2Data(v *v1.CustomerNovelsUpdateReply) *ent.CustomerNovels {
	if v == nil {
		return nil
	}
	return &ent.CustomerNovels{
		ID:         v.Id,
		NovelId:    v.NovelId,
		TypeId:     v.TypeId,
		TypeCode:   v.TypeCode,
		GroupCode:  v.GroupCode,
		FieldName:  v.FieldName,
		Cover:      v.Cover,
		OrderNum:   v.OrderNum,
		Remark:     v.Remark,
		EffectTime: v.EffectTime.AsTime(),
	}
}

func CustomerNovelsData2UpdateReply(v *ent.CustomerNovels) *v1.CustomerNovelsUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.CustomerNovelsUpdateReply{
		Id:         v.ID,
		NovelId:    v.NovelId,
		TypeId:     v.TypeId,
		TypeCode:   v.TypeCode,
		GroupCode:  v.GroupCode,
		FieldName:  v.FieldName,
		Cover:      v.Cover,
		OrderNum:   v.OrderNum,
		Remark:     v.Remark,
		EffectTime: timestamppb.New(v.EffectTime),
	}
}

func CustomerNovelsCreateReply2Data(v *v1.CustomerNovelsCreateReply) *ent.CustomerNovels {
	if v == nil {
		return nil
	}
	return &ent.CustomerNovels{
		ID:          v.Id,
		NovelId:     v.NovelId,
		TypeId:      v.TypeId,
		TypeCode:    v.TypeCode,
		GroupCode:   v.GroupCode,
		FieldName:   v.FieldName,
		Cover:       v.Cover,
		OrderNum:    v.OrderNum,
		Remark:      v.Remark,
		EffectTime:  v.EffectTime.AsTime(),
		ExpiredTime: v.ExpiredTime.AsTime(),
		CreatedAt:   v.CreatedAt.AsTime(),
		UpdatedAt:   v.UpdatedAt.AsTime(),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}

func CustomerNovelsData2CreateReply(v *ent.CustomerNovels) *v1.CustomerNovelsCreateReply {
	if v == nil {
		return nil
	}
	return &v1.CustomerNovelsCreateReply{
		Id:          v.ID,
		NovelId:     v.NovelId,
		TypeId:      v.TypeId,
		TypeCode:    v.TypeCode,
		GroupCode:   v.GroupCode,
		FieldName:   v.FieldName,
		Cover:       v.Cover,
		OrderNum:    v.OrderNum,
		Remark:      v.Remark,
		EffectTime:  timestamppb.New(v.EffectTime),
		ExpiredTime: timestamppb.New(v.ExpiredTime),
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}