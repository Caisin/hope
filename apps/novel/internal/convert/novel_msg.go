// Code generated by Caisin. DO NOT EDIT.
// source: apps/novel/internal/data/ent/schema/novel_msg.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/novelmsg/v1"
	"hope/apps/novel/internal/data/ent"
)

func NovelMsgUpdateReq2Data(v *v1.NovelMsgUpdateReq) *ent.NovelMsg {
	if v == nil {
		return nil
	}
	return &ent.NovelMsg{
		ID:      v.Id,
		Title:   v.Title,
		Msg:     v.Msg,
		MsgType: v.MsgType,
	}
}

func NovelMsgData2UpdateReq(v *ent.NovelMsg) *v1.NovelMsgUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.NovelMsgUpdateReq{
		Id:      v.ID,
		Title:   v.Title,
		Msg:     v.Msg,
		MsgType: v.MsgType,
	}
}

func NovelMsgCreateReq2Data(v *v1.NovelMsgCreateReq) *ent.NovelMsg {
	if v == nil {
		return nil
	}
	return &ent.NovelMsg{
		Title:   v.Title,
		Msg:     v.Msg,
		MsgType: v.MsgType,
	}
}

func NovelMsgData2CreateReq(v *ent.NovelMsg) *v1.NovelMsgCreateReq {
	if v == nil {
		return nil
	}
	return &v1.NovelMsgCreateReq{
		Title:   v.Title,
		Msg:     v.Msg,
		MsgType: v.MsgType,
	}
}

func NovelMsgReq2Data(v *v1.NovelMsgReq) *ent.NovelMsg {
	if v == nil {
		return nil
	}
	return &ent.NovelMsg{
		Title:   v.Title,
		Msg:     v.Msg,
		MsgType: v.MsgType,
	}
}

func NovelMsgData2Req(v *ent.NovelMsg) *v1.NovelMsgReq {
	if v == nil {
		return nil
	}
	return &v1.NovelMsgReq{
		Title:   v.Title,
		Msg:     v.Msg,
		MsgType: v.MsgType,
	}
}

func NovelMsgData2Reply(v *ent.NovelMsg) *v1.NovelMsgData {
	if v == nil {
		return nil
	}
	return &v1.NovelMsgData{
		Id:        v.ID,
		Title:     v.Title,
		Msg:       v.Msg,
		MsgType:   v.MsgType,
		Status:    v.Status,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
