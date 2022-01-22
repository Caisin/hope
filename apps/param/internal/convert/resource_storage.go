// Code generated by Caisin. DO NOT EDIT.
// source: apps/param/internal/data/ent/schema/resource_storage.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/param/resourcestorage/v1"
	"hope/apps/param/internal/data/ent"
)

func ResourceStorageUpdateReq2Data(v *v1.ResourceStorageUpdateReq) *ent.ResourceStorage {
	if v == nil {
		return nil
	}
	return &ent.ResourceStorage{
		ID:          v.Id,
		GroupId:     v.GroupId,
		StorageType: v.StorageType,
		RealName:    v.RealName,
		Bucket:      v.Bucket,
		Name:        v.Name,
		Suffix:      v.Suffix,
		Path:        v.Path,
		Type:        v.Type,
		Size:        v.Size,
		DeleteUrl:   v.DeleteUrl,
		Filename:    v.Filename,
		Key:         v.Key,
		Height:      v.Height,
		URL:         v.Url,
		Username:    v.Username,
		Width:       v.Width,
		Md5code:     v.Md5Code,
	}
}

func ResourceStorageData2UpdateReq(v *ent.ResourceStorage) *v1.ResourceStorageUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.ResourceStorageUpdateReq{
		Id:          v.ID,
		GroupId:     v.GroupId,
		StorageType: v.StorageType,
		RealName:    v.RealName,
		Bucket:      v.Bucket,
		Name:        v.Name,
		Suffix:      v.Suffix,
		Path:        v.Path,
		Type:        v.Type,
		Size:        v.Size,
		DeleteUrl:   v.DeleteUrl,
		Filename:    v.Filename,
		Key:         v.Key,
		Height:      v.Height,
		Url:         v.URL,
		Username:    v.Username,
		Width:       v.Width,
		Md5Code:     v.Md5code,
	}
}

func ResourceStorageCreateReq2Data(v *v1.ResourceStorageCreateReq) *ent.ResourceStorage {
	if v == nil {
		return nil
	}
	return &ent.ResourceStorage{
		GroupId:     v.GroupId,
		StorageType: v.StorageType,
		RealName:    v.RealName,
		Bucket:      v.Bucket,
		Name:        v.Name,
		Suffix:      v.Suffix,
		Path:        v.Path,
		Type:        v.Type,
		Size:        v.Size,
		DeleteUrl:   v.DeleteUrl,
		Filename:    v.Filename,
		Key:         v.Key,
		Height:      v.Height,
		URL:         v.Url,
		Username:    v.Username,
		Width:       v.Width,
		Md5code:     v.Md5Code,
	}
}

func ResourceStorageData2CreateReq(v *ent.ResourceStorage) *v1.ResourceStorageCreateReq {
	if v == nil {
		return nil
	}
	return &v1.ResourceStorageCreateReq{
		GroupId:     v.GroupId,
		StorageType: v.StorageType,
		RealName:    v.RealName,
		Bucket:      v.Bucket,
		Name:        v.Name,
		Suffix:      v.Suffix,
		Path:        v.Path,
		Type:        v.Type,
		Size:        v.Size,
		DeleteUrl:   v.DeleteUrl,
		Filename:    v.Filename,
		Key:         v.Key,
		Height:      v.Height,
		Url:         v.URL,
		Username:    v.Username,
		Width:       v.Width,
		Md5Code:     v.Md5code,
	}
}

func ResourceStorageReq2Data(v *v1.ResourceStorageReq) *ent.ResourceStorage {
	if v == nil {
		return nil
	}
	return &ent.ResourceStorage{
		GroupId:     v.GroupId,
		StorageType: v.StorageType,
		RealName:    v.RealName,
		Bucket:      v.Bucket,
		Name:        v.Name,
		Suffix:      v.Suffix,
		Path:        v.Path,
		Type:        v.Type,
		Size:        v.Size,
		DeleteUrl:   v.DeleteUrl,
		Filename:    v.Filename,
		Key:         v.Key,
		Height:      v.Height,
		URL:         v.Url,
		Username:    v.Username,
		Width:       v.Width,
		Md5code:     v.Md5Code,
	}
}

func ResourceStorageData2Req(v *ent.ResourceStorage) *v1.ResourceStorageReq {
	if v == nil {
		return nil
	}
	return &v1.ResourceStorageReq{
		GroupId:     v.GroupId,
		StorageType: v.StorageType,
		RealName:    v.RealName,
		Bucket:      v.Bucket,
		Name:        v.Name,
		Suffix:      v.Suffix,
		Path:        v.Path,
		Type:        v.Type,
		Size:        v.Size,
		DeleteUrl:   v.DeleteUrl,
		Filename:    v.Filename,
		Key:         v.Key,
		Height:      v.Height,
		Url:         v.URL,
		Username:    v.Username,
		Width:       v.Width,
		Md5Code:     v.Md5code,
	}
}

func ResourceStorageData2Reply(v *ent.ResourceStorage) *v1.ResourceStorageData {
	if v == nil {
		return nil
	}
	return &v1.ResourceStorageData{
		Id:          v.ID,
		GroupId:     v.GroupId,
		StorageType: v.StorageType,
		RealName:    v.RealName,
		Bucket:      v.Bucket,
		Name:        v.Name,
		Suffix:      v.Suffix,
		Path:        v.Path,
		Type:        v.Type,
		Size:        v.Size,
		DeleteUrl:   v.DeleteUrl,
		Filename:    v.Filename,
		Key:         v.Key,
		Height:      v.Height,
		Url:         v.URL,
		Username:    v.Username,
		Width:       v.Width,
		Md5Code:     v.Md5code,
		Remark:      v.Remark,
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}
