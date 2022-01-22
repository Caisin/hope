// Code generated by Caisin. DO NOT EDIT.
// source: apps/novel/internal/data/ent/schema/app_version.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/appversion/v1"
	"hope/apps/novel/internal/data/ent"
)

func AppVersionUpdateReq2Data(v *v1.AppVersionUpdateReq) *ent.AppVersion {
	if v == nil {
		return nil
	}
	return &ent.AppVersion{
		ID:          v.Id,
		Title:       v.Title,
		Version:     v.Version,
		UpdateInfo:  v.UpdateInfo,
		DownloadUrl: v.DownloadUrl,
	}
}

func AppVersionData2UpdateReq(v *ent.AppVersion) *v1.AppVersionUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.AppVersionUpdateReq{
		Id:          v.ID,
		Title:       v.Title,
		Version:     v.Version,
		UpdateInfo:  v.UpdateInfo,
		DownloadUrl: v.DownloadUrl,
	}
}

func AppVersionCreateReq2Data(v *v1.AppVersionCreateReq) *ent.AppVersion {
	if v == nil {
		return nil
	}
	return &ent.AppVersion{
		Title:       v.Title,
		Version:     v.Version,
		UpdateInfo:  v.UpdateInfo,
		DownloadUrl: v.DownloadUrl,
	}
}

func AppVersionData2CreateReq(v *ent.AppVersion) *v1.AppVersionCreateReq {
	if v == nil {
		return nil
	}
	return &v1.AppVersionCreateReq{
		Title:       v.Title,
		Version:     v.Version,
		UpdateInfo:  v.UpdateInfo,
		DownloadUrl: v.DownloadUrl,
	}
}

func AppVersionReq2Data(v *v1.AppVersionReq) *ent.AppVersion {
	if v == nil {
		return nil
	}
	return &ent.AppVersion{
		Title:       v.Title,
		Version:     v.Version,
		UpdateInfo:  v.UpdateInfo,
		DownloadUrl: v.DownloadUrl,
	}
}

func AppVersionData2Req(v *ent.AppVersion) *v1.AppVersionReq {
	if v == nil {
		return nil
	}
	return &v1.AppVersionReq{
		Title:       v.Title,
		Version:     v.Version,
		UpdateInfo:  v.UpdateInfo,
		DownloadUrl: v.DownloadUrl,
	}
}

func AppVersionData2Reply(v *ent.AppVersion) *v1.AppVersionData {
	if v == nil {
		return nil
	}
	return &v1.AppVersionData{
		Id:          v.ID,
		Title:       v.Title,
		Version:     v.Version,
		UpdateInfo:  v.UpdateInfo,
		DownloadUrl: v.DownloadUrl,
		Platform:    v.Platform,
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}
