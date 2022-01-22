// Code generated by Caisin. DO NOT EDIT.
// source: apps/param/internal/data/ent/schema/novel_tag.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/param/noveltag/v1"
	"hope/apps/param/internal/data/ent"
)

func NovelTagUpdateReq2Data(v *v1.NovelTagUpdateReq) *ent.NovelTag {
	if v == nil {
		return nil
	}
	return &ent.NovelTag{
		ID:      v.Id,
		TagId:   v.TagId,
		TagName: v.TagName,
	}
}

func NovelTagData2UpdateReq(v *ent.NovelTag) *v1.NovelTagUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.NovelTagUpdateReq{
		Id:      v.ID,
		TagId:   v.TagId,
		TagName: v.TagName,
	}
}

func NovelTagCreateReq2Data(v *v1.NovelTagCreateReq) *ent.NovelTag {
	if v == nil {
		return nil
	}
	return &ent.NovelTag{
		TagId:   v.TagId,
		TagName: v.TagName,
	}
}

func NovelTagData2CreateReq(v *ent.NovelTag) *v1.NovelTagCreateReq {
	if v == nil {
		return nil
	}
	return &v1.NovelTagCreateReq{
		TagId:   v.TagId,
		TagName: v.TagName,
	}
}

func NovelTagReq2Data(v *v1.NovelTagReq) *ent.NovelTag {
	if v == nil {
		return nil
	}
	return &ent.NovelTag{
		TagId:   v.TagId,
		TagName: v.TagName,
	}
}

func NovelTagData2Req(v *ent.NovelTag) *v1.NovelTagReq {
	if v == nil {
		return nil
	}
	return &v1.NovelTagReq{
		TagId:   v.TagId,
		TagName: v.TagName,
	}
}

func NovelTagData2Reply(v *ent.NovelTag) *v1.NovelTagData {
	if v == nil {
		return nil
	}
	return &v1.NovelTagData{
		Id:        v.ID,
		TagId:     v.TagId,
		TagName:   v.TagName,
		Remark:    v.Remark,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
