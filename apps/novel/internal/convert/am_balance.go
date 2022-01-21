// Code generated by Caisin. DO NOT EDIT.
// source: apps/novel/internal/data/ent/schema/am_balance.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/ambalance/v1"
	"hope/apps/novel/internal/data/ent"
)

func AmBalanceUpdateReq2Data(v *v1.AmBalanceUpdateReq) *ent.AmBalance {
	if v == nil {
		return nil
	}
	return &ent.AmBalance{
		ID:          v.Id,
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
		Remark:      v.Remark,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func AmBalanceData2UpdateReq(v *ent.AmBalance) *v1.AmBalanceUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.AmBalanceUpdateReq{
		Id:          v.ID,
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
		Remark:      v.Remark,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func AmBalanceCreateReq2Data(v *v1.AmBalanceCreateReq) *ent.AmBalance {
	if v == nil {
		return nil
	}
	return &ent.AmBalance{
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
		Remark:      v.Remark,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func AmBalanceData2CreateReq(v *ent.AmBalance) *v1.AmBalanceCreateReq {
	if v == nil {
		return nil
	}
	return &v1.AmBalanceCreateReq{
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
		Remark:      v.Remark,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func AmBalanceReq2Data(v *v1.AmBalanceReq) *ent.AmBalance {
	if v == nil {
		return nil
	}
	return &ent.AmBalance{
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
		Remark:      v.Remark,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func AmBalanceData2Req(v *ent.AmBalance) *v1.AmBalanceReq {
	if v == nil {
		return nil
	}
	return &v1.AmBalanceReq{
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
		Remark:      v.Remark,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func AmBalanceReply2Data(v *v1.AmBalanceReply) *ent.AmBalance {
	if v == nil {
		return nil
	}
	return &ent.AmBalance{
		ID:          v.Id,
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
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

func AmBalanceData2Reply(v *ent.AmBalance) *v1.AmBalanceReply {
	if v == nil {
		return nil
	}
	return &v1.AmBalanceReply{
		Id:          v.ID,
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
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

func AmBalanceUpdateReply2Data(v *v1.AmBalanceUpdateReply) *ent.AmBalance {
	if v == nil {
		return nil
	}
	return &ent.AmBalance{
		ID:          v.Id,
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
		Remark:      v.Remark,
		EffectTime:  v.EffectTime.AsTime(),
	}
}

func AmBalanceData2UpdateReply(v *ent.AmBalance) *v1.AmBalanceUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.AmBalanceUpdateReply{
		Id:          v.ID,
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
		Remark:      v.Remark,
		EffectTime:  timestamppb.New(v.EffectTime),
	}
}

func AmBalanceCreateReply2Data(v *v1.AmBalanceCreateReply) *ent.AmBalance {
	if v == nil {
		return nil
	}
	return &ent.AmBalance{
		ID:          v.Id,
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
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

func AmBalanceData2CreateReply(v *ent.AmBalance) *v1.AmBalanceCreateReply {
	if v == nil {
		return nil
	}
	return &v1.AmBalanceCreateReply{
		Id:          v.ID,
		UserId:      v.UserId,
		OrderId:     v.OrderId,
		EventId:     v.EventId,
		CashTag:     v.CashTag,
		AssetItemId: v.AssetItemId,
		Amount:      v.Amount,
		Balance:     v.Balance,
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
