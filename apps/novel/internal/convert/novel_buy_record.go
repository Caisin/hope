// Code generated by Caisin. DO NOT EDIT.
// source: apps/novel/internal/data/ent/schema/novel_buy_record.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/novelbuyrecord/v1"
	"hope/apps/novel/internal/data/ent"
)

func NovelBuyRecordUpdateReq2Data(v *v1.NovelBuyRecordUpdateReq) *ent.NovelBuyRecord {
	if v == nil {
		return nil
	}
	return &ent.NovelBuyRecord{
		ID:        v.Id,
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
	}
}

func NovelBuyRecordData2UpdateReq(v *ent.NovelBuyRecord) *v1.NovelBuyRecordUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.NovelBuyRecordUpdateReq{
		Id:        v.ID,
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
	}
}

func NovelBuyRecordCreateReq2Data(v *v1.NovelBuyRecordCreateReq) *ent.NovelBuyRecord {
	if v == nil {
		return nil
	}
	return &ent.NovelBuyRecord{
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
	}
}

func NovelBuyRecordData2CreateReq(v *ent.NovelBuyRecord) *v1.NovelBuyRecordCreateReq {
	if v == nil {
		return nil
	}
	return &v1.NovelBuyRecordCreateReq{
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
	}
}

func NovelBuyRecordReq2Data(v *v1.NovelBuyRecordReq) *ent.NovelBuyRecord {
	if v == nil {
		return nil
	}
	return &ent.NovelBuyRecord{
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
	}
}

func NovelBuyRecordData2Req(v *ent.NovelBuyRecord) *v1.NovelBuyRecordReq {
	if v == nil {
		return nil
	}
	return &v1.NovelBuyRecordReq{
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
	}
}

func NovelBuyRecordReply2Data(v *v1.NovelBuyRecordReply) *ent.NovelBuyRecord {
	if v == nil {
		return nil
	}
	return &ent.NovelBuyRecord{
		ID:        v.Id,
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
		Remark:    v.Remark,
		CreatedAt: v.CreatedAt.AsTime(),
		UpdatedAt: v.UpdatedAt.AsTime(),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func NovelBuyRecordData2Reply(v *ent.NovelBuyRecord) *v1.NovelBuyRecordReply {
	if v == nil {
		return nil
	}
	return &v1.NovelBuyRecordReply{
		Id:        v.ID,
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
		Remark:    v.Remark,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func NovelBuyRecordUpdateReply2Data(v *v1.NovelBuyRecordUpdateReply) *ent.NovelBuyRecord {
	if v == nil {
		return nil
	}
	return &ent.NovelBuyRecord{
		ID:        v.Id,
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
	}
}

func NovelBuyRecordData2UpdateReply(v *ent.NovelBuyRecord) *v1.NovelBuyRecordUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.NovelBuyRecordUpdateReply{
		Id:        v.ID,
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
	}
}

func NovelBuyRecordCreateReply2Data(v *v1.NovelBuyRecordCreateReply) *ent.NovelBuyRecord {
	if v == nil {
		return nil
	}
	return &ent.NovelBuyRecord{
		ID:        v.Id,
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
		Remark:    v.Remark,
		CreatedAt: v.CreatedAt.AsTime(),
		UpdatedAt: v.UpdatedAt.AsTime(),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func NovelBuyRecordData2CreateReply(v *ent.NovelBuyRecord) *v1.NovelBuyRecordCreateReply {
	if v == nil {
		return nil
	}
	return &v1.NovelBuyRecordCreateReply{
		Id:        v.ID,
		UserId:    v.UserId,
		UserName:  v.UserName,
		NovelId:   v.NovelId,
		NovelName: v.NovelName,
		PackageId: v.PackageId,
		Cover:     v.Cover,
		Coin:      v.Coin,
		Coupon:    v.Coupon,
		Remark:    v.Remark,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
