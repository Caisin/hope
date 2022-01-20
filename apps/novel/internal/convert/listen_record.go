// Code generated by Caisin. DO NOT EDIT.
// source: apps/novel/internal/data/ent/schema/listen_record.go

package convert

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/listenrecord/v1"
	"hope/apps/novel/internal/data/ent"
)

func ListenRecordUpdateReq2Data(v *v1.ListenRecordUpdateReq) *ent.ListenRecord {
	if v == nil {
		return nil
	}
	return &ent.ListenRecord{
		ID:          v.Id,
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    v.Duration.AsDuration(),
		AllDuration: v.AllDuration.AsDuration(),
	}
}

func ListenRecordData2UpdateReq(v *ent.ListenRecord) *v1.ListenRecordUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.ListenRecordUpdateReq{
		Id:          v.ID,
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    durationpb.New(v.Duration),
		AllDuration: durationpb.New(v.AllDuration),
	}
}

func ListenRecordCreateReq2Data(v *v1.ListenRecordCreateReq) *ent.ListenRecord {
	if v == nil {
		return nil
	}
	return &ent.ListenRecord{
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    v.Duration.AsDuration(),
		AllDuration: v.AllDuration.AsDuration(),
	}
}

func ListenRecordData2CreateReq(v *ent.ListenRecord) *v1.ListenRecordCreateReq {
	if v == nil {
		return nil
	}
	return &v1.ListenRecordCreateReq{
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    durationpb.New(v.Duration),
		AllDuration: durationpb.New(v.AllDuration),
	}
}

func ListenRecordReq2Data(v *v1.ListenRecordReq) *ent.ListenRecord {
	if v == nil {
		return nil
	}
	return &ent.ListenRecord{
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    v.Duration.AsDuration(),
		AllDuration: v.AllDuration.AsDuration(),
	}
}

func ListenRecordData2Req(v *ent.ListenRecord) *v1.ListenRecordReq {
	if v == nil {
		return nil
	}
	return &v1.ListenRecordReq{
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    durationpb.New(v.Duration),
		AllDuration: durationpb.New(v.AllDuration),
	}
}

func ListenRecordReply2Data(v *v1.ListenRecordReply) *ent.ListenRecord {
	if v == nil {
		return nil
	}
	return &ent.ListenRecord{
		ID:          v.Id,
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    v.Duration.AsDuration(),
		AllDuration: v.AllDuration.AsDuration(),
		DayDuration: v.DayDuration.AsDuration(),
		CreatedAt:   v.CreatedAt.AsTime(),
		UpdatedAt:   v.UpdatedAt.AsTime(),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}

func ListenRecordData2Reply(v *ent.ListenRecord) *v1.ListenRecordReply {
	if v == nil {
		return nil
	}
	return &v1.ListenRecordReply{
		Id:          v.ID,
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    durationpb.New(v.Duration),
		AllDuration: durationpb.New(v.AllDuration),
		DayDuration: durationpb.New(v.DayDuration),
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}

func ListenRecordUpdateReply2Data(v *v1.ListenRecordUpdateReply) *ent.ListenRecord {
	if v == nil {
		return nil
	}
	return &ent.ListenRecord{
		ID:          v.Id,
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    v.Duration.AsDuration(),
		AllDuration: v.AllDuration.AsDuration(),
	}
}

func ListenRecordData2UpdateReply(v *ent.ListenRecord) *v1.ListenRecordUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.ListenRecordUpdateReply{
		Id:          v.ID,
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    durationpb.New(v.Duration),
		AllDuration: durationpb.New(v.AllDuration),
	}
}

func ListenRecordCreateReply2Data(v *v1.ListenRecordCreateReply) *ent.ListenRecord {
	if v == nil {
		return nil
	}
	return &ent.ListenRecord{
		ID:          v.Id,
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    v.Duration.AsDuration(),
		AllDuration: v.AllDuration.AsDuration(),
		DayDuration: v.DayDuration.AsDuration(),
		CreatedAt:   v.CreatedAt.AsTime(),
		UpdatedAt:   v.UpdatedAt.AsTime(),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}

func ListenRecordData2CreateReply(v *ent.ListenRecord) *v1.ListenRecordCreateReply {
	if v == nil {
		return nil
	}
	return &v1.ListenRecordCreateReply{
		Id:          v.ID,
		UserId:      v.UserId,
		ChapterId:   v.ChapterId,
		NovelId:     v.NovelId,
		ListenTimes: v.ListenTimes,
		Duration:    durationpb.New(v.Duration),
		AllDuration: durationpb.New(v.AllDuration),
		DayDuration: durationpb.New(v.DayDuration),
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
		CreateBy:    v.CreateBy,
		UpdateBy:    v.UpdateBy,
		TenantId:    v.TenantId,
	}
}
