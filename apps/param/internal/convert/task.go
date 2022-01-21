// Code generated by Caisin. DO NOT EDIT.
// source: apps/param/internal/data/ent/schema/task.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/param/task/v1"
	"hope/apps/param/internal/data/ent"
)

func TaskUpdateReq2Data(v *v1.TaskUpdateReq) *ent.Task {
	if v == nil {
		return nil
	}
	return &ent.Task{
		ID:            v.Id,
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    v.EffectTime.AsTime(),
	}
}

func TaskData2UpdateReq(v *ent.Task) *v1.TaskUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.TaskUpdateReq{
		Id:            v.ID,
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    timestamppb.New(v.EffectTime),
	}
}

func TaskCreateReq2Data(v *v1.TaskCreateReq) *ent.Task {
	if v == nil {
		return nil
	}
	return &ent.Task{
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    v.EffectTime.AsTime(),
	}
}

func TaskData2CreateReq(v *ent.Task) *v1.TaskCreateReq {
	if v == nil {
		return nil
	}
	return &v1.TaskCreateReq{
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    timestamppb.New(v.EffectTime),
	}
}

func TaskReq2Data(v *v1.TaskReq) *ent.Task {
	if v == nil {
		return nil
	}
	return &ent.Task{
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    v.EffectTime.AsTime(),
	}
}

func TaskData2Req(v *ent.Task) *v1.TaskReq {
	if v == nil {
		return nil
	}
	return &v1.TaskReq{
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    timestamppb.New(v.EffectTime),
	}
}

func TaskReply2Data(v *v1.TaskReply) *ent.Task {
	if v == nil {
		return nil
	}
	return &ent.Task{
		ID:            v.Id,
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    v.EffectTime.AsTime(),
		ExpiredTime:   v.ExpiredTime.AsTime(),
		CreatedAt:     v.CreatedAt.AsTime(),
		UpdatedAt:     v.UpdatedAt.AsTime(),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func TaskData2Reply(v *ent.Task) *v1.TaskReply {
	if v == nil {
		return nil
	}
	return &v1.TaskReply{
		Id:            v.ID,
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    timestamppb.New(v.EffectTime),
		ExpiredTime:   timestamppb.New(v.ExpiredTime),
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func TaskUpdateReply2Data(v *v1.TaskUpdateReply) *ent.Task {
	if v == nil {
		return nil
	}
	return &ent.Task{
		ID:            v.Id,
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    v.EffectTime.AsTime(),
	}
}

func TaskData2UpdateReply(v *ent.Task) *v1.TaskUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.TaskUpdateReply{
		Id:            v.ID,
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    timestamppb.New(v.EffectTime),
	}
}

func TaskCreateReply2Data(v *v1.TaskCreateReply) *ent.Task {
	if v == nil {
		return nil
	}
	return &ent.Task{
		ID:            v.Id,
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    v.EffectTime.AsTime(),
		ExpiredTime:   v.ExpiredTime.AsTime(),
		CreatedAt:     v.CreatedAt.AsTime(),
		UpdatedAt:     v.UpdatedAt.AsTime(),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func TaskData2CreateReply(v *ent.Task) *v1.TaskCreateReply {
	if v == nil {
		return nil
	}
	return &v1.TaskCreateReply{
		Id:            v.ID,
		TaskName:      v.TaskName,
		TaskGroup:     v.TaskGroup,
		Unit:          v.Unit,
		Topic:         v.Topic,
		Function:      v.Function,
		TaskCode:      v.TaskCode,
		PreTask:       v.PreTask,
		NovelId:       v.NovelId,
		CycleType:     v.CycleType,
		Remark:        v.Remark,
		Amount:        v.Amount,
		Reward:        v.Reward,
		AmountItem:    v.AmountItem,
		RewardItem:    v.RewardItem,
		TargetNames:   v.TargetNames,
		TargetAmounts: v.TargetAmounts,
		Status:        v.Status,
		SortNum:       v.SortNum,
		ActionType:    v.ActionType,
		EffectTime:    timestamppb.New(v.EffectTime),
		ExpiredTime:   timestamppb.New(v.ExpiredTime),
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}