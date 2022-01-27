// Code generated by Caisin. DO NOT EDIT.
// source: apps/admin/internal/data/ent/schema/sys_api.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/admin/sysapi/v1"
	"hope/apps/admin/internal/data/ent"
)

func SysApiUpdateReq2Data(v *v1.SysApiUpdateReq) *ent.SysApi {
	if v == nil {
		return nil
	}
	return &ent.SysApi{
		ID:     v.Id,
		Handle: v.Handle,
		Title:  v.Title,
		Path:   v.Path,
		Action: v.Action,
		Type:   v.Type,
	}
}

func SysApiData2UpdateReq(v *ent.SysApi) *v1.SysApiUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.SysApiUpdateReq{
		Id:     v.ID,
		Handle: v.Handle,
		Title:  v.Title,
		Path:   v.Path,
		Action: v.Action,
		Type:   v.Type,
	}
}

func SysApiCreateReq2Data(v *v1.SysApiCreateReq) *ent.SysApi {
	if v == nil {
		return nil
	}
	return &ent.SysApi{
		Handle: v.Handle,
		Title:  v.Title,
		Path:   v.Path,
		Action: v.Action,
		Type:   v.Type,
	}
}

func SysApiData2CreateReq(v *ent.SysApi) *v1.SysApiCreateReq {
	if v == nil {
		return nil
	}
	return &v1.SysApiCreateReq{
		Handle: v.Handle,
		Title:  v.Title,
		Path:   v.Path,
		Action: v.Action,
		Type:   v.Type,
	}
}

func SysApiReq2Data(v *v1.SysApiReq) *ent.SysApi {
	if v == nil {
		return nil
	}
	return &ent.SysApi{
		Handle: v.Handle,
		Title:  v.Title,
		Path:   v.Path,
		Action: v.Action,
		Type:   v.Type,
	}
}

func SysApiData2Req(v *ent.SysApi) *v1.SysApiReq {
	if v == nil {
		return nil
	}
	return &v1.SysApiReq{
		Handle: v.Handle,
		Title:  v.Title,
		Path:   v.Path,
		Action: v.Action,
		Type:   v.Type,
	}
}

func SysApiData2Reply(v *ent.SysApi) *v1.SysApiData {
	if v == nil {
		return nil
	}
	return &v1.SysApiData{
		Id:        v.ID,
		Handle:    v.Handle,
		Title:     v.Title,
		Path:      v.Path,
		Action:    v.Action,
		Type:      v.Type,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
