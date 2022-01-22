// Code generated by Caisin. DO NOT EDIT.
// source: apps/param/internal/data/ent/schema/user_resource.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/param/userresource/v1"
	"hope/apps/param/internal/data/ent"
)

func UserResourceUpdateReq2Data(v *v1.UserResourceUpdateReq) *ent.UserResource {
	if v == nil {
		return nil
	}
	return &ent.UserResource{
		ID:      v.Id,
		ResType: v.ResType,
		Name:    v.Name,
		URL:     v.Url,
	}
}

func UserResourceData2UpdateReq(v *ent.UserResource) *v1.UserResourceUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.UserResourceUpdateReq{
		Id:      v.ID,
		ResType: v.ResType,
		Name:    v.Name,
		Url:     v.URL,
	}
}

func UserResourceCreateReq2Data(v *v1.UserResourceCreateReq) *ent.UserResource {
	if v == nil {
		return nil
	}
	return &ent.UserResource{
		ResType: v.ResType,
		Name:    v.Name,
		URL:     v.Url,
	}
}

func UserResourceData2CreateReq(v *ent.UserResource) *v1.UserResourceCreateReq {
	if v == nil {
		return nil
	}
	return &v1.UserResourceCreateReq{
		ResType: v.ResType,
		Name:    v.Name,
		Url:     v.URL,
	}
}

func UserResourceReq2Data(v *v1.UserResourceReq) *ent.UserResource {
	if v == nil {
		return nil
	}
	return &ent.UserResource{
		ResType: v.ResType,
		Name:    v.Name,
		URL:     v.Url,
	}
}

func UserResourceData2Req(v *ent.UserResource) *v1.UserResourceReq {
	if v == nil {
		return nil
	}
	return &v1.UserResourceReq{
		ResType: v.ResType,
		Name:    v.Name,
		Url:     v.URL,
	}
}

func UserResourceData2Reply(v *ent.UserResource) *v1.UserResourceData {
	if v == nil {
		return nil
	}
	return &v1.UserResourceData{
		Id:        v.ID,
		ResType:   v.ResType,
		Name:      v.Name,
		Url:       v.URL,
		Summary:   v.Summary,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
