// Code generated by Caisin. DO NOT EDIT.
// source: apps/admin/internal/data/ent/schema/sys_dict_data.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/admin/sysdictdata/v1"
	"hope/apps/admin/internal/data/ent"
)

func SysDictDataUpdateReq2Data(v *v1.SysDictDataUpdateReq) *ent.SysDictData {
	if v == nil {
		return nil
	}
	return &ent.SysDictData{
		ID:        v.Id,
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
	}
}

func SysDictDataData2UpdateReq(v *ent.SysDictData) *v1.SysDictDataUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.SysDictDataUpdateReq{
		Id:        v.ID,
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
	}
}

func SysDictDataCreateReq2Data(v *v1.SysDictDataCreateReq) *ent.SysDictData {
	if v == nil {
		return nil
	}
	return &ent.SysDictData{
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
	}
}

func SysDictDataData2CreateReq(v *ent.SysDictData) *v1.SysDictDataCreateReq {
	if v == nil {
		return nil
	}
	return &v1.SysDictDataCreateReq{
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
	}
}

func SysDictDataReq2Data(v *v1.SysDictDataReq) *ent.SysDictData {
	if v == nil {
		return nil
	}
	return &ent.SysDictData{
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
	}
}

func SysDictDataData2Req(v *ent.SysDictData) *v1.SysDictDataReq {
	if v == nil {
		return nil
	}
	return &v1.SysDictDataReq{
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
	}
}

func SysDictDataReply2Data(v *v1.SysDictDataReply) *ent.SysDictData {
	if v == nil {
		return nil
	}
	return &ent.SysDictData{
		ID:        v.Id,
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
		Remark:    v.Remark,
		CreatedAt: v.CreatedAt.AsTime(),
		UpdatedAt: v.UpdatedAt.AsTime(),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func SysDictDataData2Reply(v *ent.SysDictData) *v1.SysDictDataReply {
	if v == nil {
		return nil
	}
	return &v1.SysDictDataReply{
		Id:        v.ID,
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
		Remark:    v.Remark,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func SysDictDataUpdateReply2Data(v *v1.SysDictDataUpdateReply) *ent.SysDictData {
	if v == nil {
		return nil
	}
	return &ent.SysDictData{
		ID:        v.Id,
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
	}
}

func SysDictDataData2UpdateReply(v *ent.SysDictData) *v1.SysDictDataUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.SysDictDataUpdateReply{
		Id:        v.ID,
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
	}
}

func SysDictDataCreateReply2Data(v *v1.SysDictDataCreateReply) *ent.SysDictData {
	if v == nil {
		return nil
	}
	return &ent.SysDictData{
		ID:        v.Id,
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
		Remark:    v.Remark,
		CreatedAt: v.CreatedAt.AsTime(),
		UpdatedAt: v.UpdatedAt.AsTime(),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func SysDictDataData2CreateReply(v *ent.SysDictData) *v1.SysDictDataCreateReply {
	if v == nil {
		return nil
	}
	return &v1.SysDictDataCreateReply{
		Id:        v.ID,
		DictSort:  v.DictSort,
		DictLabel: v.DictLabel,
		DictValue: v.DictValue,
		IsDefault: v.IsDefault,
		Status:    v.Status,
		Default:   v.Default,
		Remark:    v.Remark,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
