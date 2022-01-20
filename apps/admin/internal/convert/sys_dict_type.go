// Code generated by Caisin. DO NOT EDIT.
// source: apps/admin/internal/data/ent/schema/sys_dict_type.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/admin/sysdicttype/v1"
	"hope/apps/admin/internal/data/ent"
)

func SysDictTypeUpdateReq2Data(v *v1.SysDictTypeUpdateReq) *ent.SysDictType {
	if v == nil {
		return nil
	}
	return &ent.SysDictType{
		ID:       v.Id,
		DictName: v.DictName,
		DictType: v.DictType,
		Status:   v.Status,
	}
}

func SysDictTypeData2UpdateReq(v *ent.SysDictType) *v1.SysDictTypeUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.SysDictTypeUpdateReq{
		Id:       v.ID,
		DictName: v.DictName,
		DictType: v.DictType,
		Status:   v.Status,
	}
}

func SysDictTypeCreateReq2Data(v *v1.SysDictTypeCreateReq) *ent.SysDictType {
	if v == nil {
		return nil
	}
	return &ent.SysDictType{
		DictName: v.DictName,
		DictType: v.DictType,
		Status:   v.Status,
	}
}

func SysDictTypeData2CreateReq(v *ent.SysDictType) *v1.SysDictTypeCreateReq {
	if v == nil {
		return nil
	}
	return &v1.SysDictTypeCreateReq{
		DictName: v.DictName,
		DictType: v.DictType,
		Status:   v.Status,
	}
}

func SysDictTypeReq2Data(v *v1.SysDictTypeReq) *ent.SysDictType {
	if v == nil {
		return nil
	}
	return &ent.SysDictType{
		DictName: v.DictName,
		DictType: v.DictType,
		Status:   v.Status,
	}
}

func SysDictTypeData2Req(v *ent.SysDictType) *v1.SysDictTypeReq {
	if v == nil {
		return nil
	}
	return &v1.SysDictTypeReq{
		DictName: v.DictName,
		DictType: v.DictType,
		Status:   v.Status,
	}
}

func SysDictTypeReply2Data(v *v1.SysDictTypeReply) *ent.SysDictType {
	if v == nil {
		return nil
	}
	return &ent.SysDictType{
		ID:        v.Id,
		DictName:  v.DictName,
		DictType:  v.DictType,
		Status:    v.Status,
		Remark:    v.Remark,
		CreatedAt: v.CreatedAt.AsTime(),
		UpdatedAt: v.UpdatedAt.AsTime(),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func SysDictTypeData2Reply(v *ent.SysDictType) *v1.SysDictTypeReply {
	if v == nil {
		return nil
	}
	return &v1.SysDictTypeReply{
		Id:        v.ID,
		DictName:  v.DictName,
		DictType:  v.DictType,
		Status:    v.Status,
		Remark:    v.Remark,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func SysDictTypeUpdateReply2Data(v *v1.SysDictTypeUpdateReply) *ent.SysDictType {
	if v == nil {
		return nil
	}
	return &ent.SysDictType{
		ID:       v.Id,
		DictName: v.DictName,
		DictType: v.DictType,
		Status:   v.Status,
	}
}

func SysDictTypeData2UpdateReply(v *ent.SysDictType) *v1.SysDictTypeUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.SysDictTypeUpdateReply{
		Id:       v.ID,
		DictName: v.DictName,
		DictType: v.DictType,
		Status:   v.Status,
	}
}

func SysDictTypeCreateReply2Data(v *v1.SysDictTypeCreateReply) *ent.SysDictType {
	if v == nil {
		return nil
	}
	return &ent.SysDictType{
		ID:        v.Id,
		DictName:  v.DictName,
		DictType:  v.DictType,
		Status:    v.Status,
		Remark:    v.Remark,
		CreatedAt: v.CreatedAt.AsTime(),
		UpdatedAt: v.UpdatedAt.AsTime(),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func SysDictTypeData2CreateReply(v *ent.SysDictType) *v1.SysDictTypeCreateReply {
	if v == nil {
		return nil
	}
	return &v1.SysDictTypeCreateReply{
		Id:        v.ID,
		DictName:  v.DictName,
		DictType:  v.DictType,
		Status:    v.Status,
		Remark:    v.Remark,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
