// Code generated by Caisin. DO NOT EDIT.
// source: apps/admin/internal/data/ent/schema/sys_job.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/admin/sysjob/v1"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/sysjob"
)

func SysJobUpdateReq2Data(v *v1.SysJobUpdateReq) *ent.SysJob {
	if v == nil {
		return nil
	}
	return &ent.SysJob{
		ID:             v.Id,
		JobName:        v.JobName,
		JobGroup:       v.JobGroup,
		JobType:        v.JobType,
		CronExpression: v.CronExpression,
		InvokeTarget:   v.InvokeTarget,
		Args:           v.Args,
		ExecPolicy:     v.ExecPolicy,
		Concurrent:     v.Concurrent,
		State:          sysjob.State(v.State),
		EntryId:        v.EntryId,
	}
}

func SysJobData2UpdateReq(v *ent.SysJob) *v1.SysJobUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.SysJobUpdateReq{
		Id:             v.ID,
		JobName:        v.JobName,
		JobGroup:       v.JobGroup,
		JobType:        v.JobType,
		CronExpression: v.CronExpression,
		InvokeTarget:   v.InvokeTarget,
		Args:           v.Args,
		ExecPolicy:     v.ExecPolicy,
		Concurrent:     v.Concurrent,
		State:          string(v.State),
		EntryId:        v.EntryId,
	}
}

func SysJobCreateReq2Data(v *v1.SysJobCreateReq) *ent.SysJob {
	if v == nil {
		return nil
	}
	return &ent.SysJob{
		JobName:        v.JobName,
		JobGroup:       v.JobGroup,
		JobType:        v.JobType,
		CronExpression: v.CronExpression,
		InvokeTarget:   v.InvokeTarget,
		Args:           v.Args,
		ExecPolicy:     v.ExecPolicy,
		Concurrent:     v.Concurrent,
		State:          sysjob.State(v.State),
		EntryId:        v.EntryId,
	}
}

func SysJobData2CreateReq(v *ent.SysJob) *v1.SysJobCreateReq {
	if v == nil {
		return nil
	}
	return &v1.SysJobCreateReq{
		JobName:        v.JobName,
		JobGroup:       v.JobGroup,
		JobType:        v.JobType,
		CronExpression: v.CronExpression,
		InvokeTarget:   v.InvokeTarget,
		Args:           v.Args,
		ExecPolicy:     v.ExecPolicy,
		Concurrent:     v.Concurrent,
		State:          string(v.State),
		EntryId:        v.EntryId,
	}
}

func SysJobReq2Data(v *v1.SysJobReq) *ent.SysJob {
	if v == nil {
		return nil
	}
	return &ent.SysJob{
		JobName:        v.JobName,
		JobGroup:       v.JobGroup,
		JobType:        v.JobType,
		CronExpression: v.CronExpression,
		InvokeTarget:   v.InvokeTarget,
		Args:           v.Args,
		ExecPolicy:     v.ExecPolicy,
		Concurrent:     v.Concurrent,
		State:          sysjob.State(v.State),
		EntryId:        v.EntryId,
	}
}

func SysJobData2Req(v *ent.SysJob) *v1.SysJobReq {
	if v == nil {
		return nil
	}
	return &v1.SysJobReq{
		JobName:        v.JobName,
		JobGroup:       v.JobGroup,
		JobType:        v.JobType,
		CronExpression: v.CronExpression,
		InvokeTarget:   v.InvokeTarget,
		Args:           v.Args,
		ExecPolicy:     v.ExecPolicy,
		Concurrent:     v.Concurrent,
		State:          string(v.State),
		EntryId:        v.EntryId,
	}
}

func SysJobData2Reply(v *ent.SysJob) *v1.SysJobData {
	if v == nil {
		return nil
	}
	return &v1.SysJobData{
		Id:             v.ID,
		JobName:        v.JobName,
		JobGroup:       v.JobGroup,
		JobType:        v.JobType,
		CronExpression: v.CronExpression,
		InvokeTarget:   v.InvokeTarget,
		Args:           v.Args,
		ExecPolicy:     v.ExecPolicy,
		Concurrent:     v.Concurrent,
		State:          string(v.State),
		EntryId:        v.EntryId,
		CreatedAt:      timestamppb.New(v.CreatedAt),
		UpdatedAt:      timestamppb.New(v.UpdatedAt),
		CreateBy:       v.CreateBy,
		UpdateBy:       v.UpdateBy,
		TenantId:       v.TenantId,
	}
}
