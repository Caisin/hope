// Code generated by Caisin. DO NOT EDIT.
// source: apps/novel/internal/data/ent/schema/activity_component.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/activitycomponent/v1"
	"hope/apps/novel/internal/data/ent"
)

func ActivityComponentUpdateReq2Data(v *v1.ActivityComponentUpdateReq) *ent.ActivityComponent {
	if v == nil {
		return nil
	}
	return &ent.ActivityComponent{
		ID:            v.Id,
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
	}
}

func ActivityComponentData2UpdateReq(v *ent.ActivityComponent) *v1.ActivityComponentUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.ActivityComponentUpdateReq{
		Id:            v.ID,
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
	}
}

func ActivityComponentCreateReq2Data(v *v1.ActivityComponentCreateReq) *ent.ActivityComponent {
	if v == nil {
		return nil
	}
	return &ent.ActivityComponent{
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
	}
}

func ActivityComponentData2CreateReq(v *ent.ActivityComponent) *v1.ActivityComponentCreateReq {
	if v == nil {
		return nil
	}
	return &v1.ActivityComponentCreateReq{
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
	}
}

func ActivityComponentReq2Data(v *v1.ActivityComponentReq) *ent.ActivityComponent {
	if v == nil {
		return nil
	}
	return &ent.ActivityComponent{
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
	}
}

func ActivityComponentData2Req(v *ent.ActivityComponent) *v1.ActivityComponentReq {
	if v == nil {
		return nil
	}
	return &v1.ActivityComponentReq{
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
	}
}

func ActivityComponentReply2Data(v *v1.ActivityComponentReply) *ent.ActivityComponent {
	if v == nil {
		return nil
	}
	return &ent.ActivityComponent{
		ID:            v.Id,
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
		ResDays:       v.ResDays,
		CreatedAt:     v.CreatedAt.AsTime(),
		UpdatedAt:     v.UpdatedAt.AsTime(),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func ActivityComponentData2Reply(v *ent.ActivityComponent) *v1.ActivityComponentReply {
	if v == nil {
		return nil
	}
	return &v1.ActivityComponentReply{
		Id:            v.ID,
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
		ResDays:       v.ResDays,
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func ActivityComponentUpdateReply2Data(v *v1.ActivityComponentUpdateReply) *ent.ActivityComponent {
	if v == nil {
		return nil
	}
	return &ent.ActivityComponent{
		ID:            v.Id,
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
	}
}

func ActivityComponentData2UpdateReply(v *ent.ActivityComponent) *v1.ActivityComponentUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.ActivityComponentUpdateReply{
		Id:            v.ID,
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
	}
}

func ActivityComponentCreateReply2Data(v *v1.ActivityComponentCreateReply) *ent.ActivityComponent {
	if v == nil {
		return nil
	}
	return &ent.ActivityComponent{
		ID:            v.Id,
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
		ResDays:       v.ResDays,
		CreatedAt:     v.CreatedAt.AsTime(),
		UpdatedAt:     v.UpdatedAt.AsTime(),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func ActivityComponentData2CreateReply(v *ent.ActivityComponent) *v1.ActivityComponentCreateReply {
	if v == nil {
		return nil
	}
	return &v1.ActivityComponentCreateReply{
		Id:            v.ID,
		ActivityCode:  v.ActivityCode,
		ComponentType: v.ComponentType,
		Policy:        v.Policy,
		VipDays:       v.VipDays,
		MinConsume:    v.MinConsume,
		MaxConsume:    v.MaxConsume,
		MinPayNum:     v.MinPayNum,
		PayTimes:      v.PayTimes,
		PayAmount:     v.PayAmount,
		RegDays:       v.RegDays,
		Summary:       v.Summary,
		AssetItemId:   v.AssetItemId,
		Amount:        v.Amount,
		ResId:         v.ResId,
		ResDays:       v.ResDays,
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}