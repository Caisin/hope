// Code generated by Caisin. DO NOT EDIT.
// source: apps/param/internal/data/ent/schema/user_consume.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/param/userconsume/v1"
	"hope/apps/param/internal/data/ent"
)

func UserConsumeUpdateReq2Data(v *v1.UserConsumeUpdateReq) *ent.UserConsume {
	if v == nil {
		return nil
	}
	return &ent.UserConsume{
		ID:      v.Id,
		NovelId: v.NovelId,
		Coin:    v.Coin,
		Coupon:  v.Coupon,
	}
}

func UserConsumeData2UpdateReq(v *ent.UserConsume) *v1.UserConsumeUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.UserConsumeUpdateReq{
		Id:      v.ID,
		NovelId: v.NovelId,
		Coin:    v.Coin,
		Coupon:  v.Coupon,
	}
}

func UserConsumeCreateReq2Data(v *v1.UserConsumeCreateReq) *ent.UserConsume {
	if v == nil {
		return nil
	}
	return &ent.UserConsume{
		NovelId: v.NovelId,
		Coin:    v.Coin,
		Coupon:  v.Coupon,
	}
}

func UserConsumeData2CreateReq(v *ent.UserConsume) *v1.UserConsumeCreateReq {
	if v == nil {
		return nil
	}
	return &v1.UserConsumeCreateReq{
		NovelId: v.NovelId,
		Coin:    v.Coin,
		Coupon:  v.Coupon,
	}
}

func UserConsumeReq2Data(v *v1.UserConsumeReq) *ent.UserConsume {
	if v == nil {
		return nil
	}
	return &ent.UserConsume{
		NovelId: v.NovelId,
		Coin:    v.Coin,
		Coupon:  v.Coupon,
	}
}

func UserConsumeData2Req(v *ent.UserConsume) *v1.UserConsumeReq {
	if v == nil {
		return nil
	}
	return &v1.UserConsumeReq{
		NovelId: v.NovelId,
		Coin:    v.Coin,
		Coupon:  v.Coupon,
	}
}

func UserConsumeData2Reply(v *ent.UserConsume) *v1.UserConsumeData {
	if v == nil {
		return nil
	}
	return &v1.UserConsumeData{
		Id:        v.ID,
		NovelId:   v.NovelId,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
		Discount:  v.Discount,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
