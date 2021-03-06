// Code generated by Caisin. DO NOT EDIT.
// source: apps/novel/internal/data/ent/schema/asset_change_log.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/assetchangelog/v1"
	"hope/apps/novel/internal/data/ent"
)

func AssetChangeLogUpdateReq2Data(v *v1.AssetChangeLogUpdateReq) *ent.AssetChangeLog {
	if v == nil {
		return nil
	}
	return &ent.AssetChangeLog{
		ID:          v.Id,
		OrderId:     v.OrderId,
		BalanceId:   v.BalanceId,
		EventId:     v.EventId,
		UserId:      v.UserId,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		OldBalance:  v.OldBalance,
		NewBalance:  v.NewBalance,
		Remark:      v.Remark,
	}
}

func AssetChangeLogData2UpdateReq(v *ent.AssetChangeLog) *v1.AssetChangeLogUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.AssetChangeLogUpdateReq{
		Id:          v.ID,
		OrderId:     v.OrderId,
		BalanceId:   v.BalanceId,
		EventId:     v.EventId,
		UserId:      v.UserId,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		OldBalance:  v.OldBalance,
		NewBalance:  v.NewBalance,
		Remark:      v.Remark,
	}
}

func AssetChangeLogCreateReq2Data(v *v1.AssetChangeLogCreateReq) *ent.AssetChangeLog {
	if v == nil {
		return nil
	}
	return &ent.AssetChangeLog{
		OrderId:     v.OrderId,
		BalanceId:   v.BalanceId,
		EventId:     v.EventId,
		UserId:      v.UserId,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		OldBalance:  v.OldBalance,
		NewBalance:  v.NewBalance,
		Remark:      v.Remark,
	}
}

func AssetChangeLogData2CreateReq(v *ent.AssetChangeLog) *v1.AssetChangeLogCreateReq {
	if v == nil {
		return nil
	}
	return &v1.AssetChangeLogCreateReq{
		OrderId:     v.OrderId,
		BalanceId:   v.BalanceId,
		EventId:     v.EventId,
		UserId:      v.UserId,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		OldBalance:  v.OldBalance,
		NewBalance:  v.NewBalance,
		Remark:      v.Remark,
	}
}

func AssetChangeLogReq2Data(v *v1.AssetChangeLogReq) *ent.AssetChangeLog {
	if v == nil {
		return nil
	}
	return &ent.AssetChangeLog{
		OrderId:     v.OrderId,
		BalanceId:   v.BalanceId,
		EventId:     v.EventId,
		UserId:      v.UserId,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		OldBalance:  v.OldBalance,
		NewBalance:  v.NewBalance,
		Remark:      v.Remark,
	}
}

func AssetChangeLogData2Req(v *ent.AssetChangeLog) *v1.AssetChangeLogReq {
	if v == nil {
		return nil
	}
	return &v1.AssetChangeLogReq{
		OrderId:     v.OrderId,
		BalanceId:   v.BalanceId,
		EventId:     v.EventId,
		UserId:      v.UserId,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		OldBalance:  v.OldBalance,
		NewBalance:  v.NewBalance,
		Remark:      v.Remark,
	}
}

func AssetChangeLogData2Reply(v *ent.AssetChangeLog) *v1.AssetChangeLogData {
	if v == nil {
		return nil
	}
	return &v1.AssetChangeLogData{
		Id:          v.ID,
		OrderId:     v.OrderId,
		BalanceId:   v.BalanceId,
		EventId:     v.EventId,
		UserId:      v.UserId,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		OldBalance:  v.OldBalance,
		NewBalance:  v.NewBalance,
		Remark:      v.Remark,
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}
