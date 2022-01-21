// Code generated by Caisin. DO NOT EDIT.
// source: apps/admin/internal/data/ent/schema/sys_opera_log.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/admin/sysoperalog/v1"
	"hope/apps/admin/internal/data/ent"
)

func SysOperaLogUpdateReq2Data(v *v1.SysOperaLogUpdateReq) *ent.SysOperaLog {
	if v == nil {
		return nil
	}
	return &ent.SysOperaLog{
		ID:            v.Id,
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      v.OperTime.AsTime(),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
	}
}

func SysOperaLogData2UpdateReq(v *ent.SysOperaLog) *v1.SysOperaLogUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.SysOperaLogUpdateReq{
		Id:            v.ID,
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      timestamppb.New(v.OperTime),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
	}
}

func SysOperaLogCreateReq2Data(v *v1.SysOperaLogCreateReq) *ent.SysOperaLog {
	if v == nil {
		return nil
	}
	return &ent.SysOperaLog{
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      v.OperTime.AsTime(),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
	}
}

func SysOperaLogData2CreateReq(v *ent.SysOperaLog) *v1.SysOperaLogCreateReq {
	if v == nil {
		return nil
	}
	return &v1.SysOperaLogCreateReq{
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      timestamppb.New(v.OperTime),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
	}
}

func SysOperaLogReq2Data(v *v1.SysOperaLogReq) *ent.SysOperaLog {
	if v == nil {
		return nil
	}
	return &ent.SysOperaLog{
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      v.OperTime.AsTime(),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
	}
}

func SysOperaLogData2Req(v *ent.SysOperaLog) *v1.SysOperaLogReq {
	if v == nil {
		return nil
	}
	return &v1.SysOperaLogReq{
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      timestamppb.New(v.OperTime),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
	}
}

func SysOperaLogReply2Data(v *v1.SysOperaLogReply) *ent.SysOperaLog {
	if v == nil {
		return nil
	}
	return &ent.SysOperaLog{
		ID:            v.Id,
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      v.OperTime.AsTime(),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
		UserAgent:     v.UserAgent,
		CreatedAt:     v.CreatedAt.AsTime(),
		UpdatedAt:     v.UpdatedAt.AsTime(),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func SysOperaLogData2Reply(v *ent.SysOperaLog) *v1.SysOperaLogReply {
	if v == nil {
		return nil
	}
	return &v1.SysOperaLogReply{
		Id:            v.ID,
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      timestamppb.New(v.OperTime),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
		UserAgent:     v.UserAgent,
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func SysOperaLogUpdateReply2Data(v *v1.SysOperaLogUpdateReply) *ent.SysOperaLog {
	if v == nil {
		return nil
	}
	return &ent.SysOperaLog{
		ID:            v.Id,
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      v.OperTime.AsTime(),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
	}
}

func SysOperaLogData2UpdateReply(v *ent.SysOperaLog) *v1.SysOperaLogUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.SysOperaLogUpdateReply{
		Id:            v.ID,
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      timestamppb.New(v.OperTime),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
	}
}

func SysOperaLogCreateReply2Data(v *v1.SysOperaLogCreateReply) *ent.SysOperaLog {
	if v == nil {
		return nil
	}
	return &ent.SysOperaLog{
		ID:            v.Id,
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      v.OperTime.AsTime(),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
		UserAgent:     v.UserAgent,
		CreatedAt:     v.CreatedAt.AsTime(),
		UpdatedAt:     v.UpdatedAt.AsTime(),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func SysOperaLogData2CreateReply(v *ent.SysOperaLog) *v1.SysOperaLogCreateReply {
	if v == nil {
		return nil
	}
	return &v1.SysOperaLogCreateReply{
		Id:            v.ID,
		UserId:        v.UserId,
		Title:         v.Title,
		RequestId:     v.RequestId,
		BusinessType:  v.BusinessType,
		BusinessTypes: v.BusinessTypes,
		Method:        v.Method,
		RequestMethod: v.RequestMethod,
		OperatorType:  v.OperatorType,
		OperName:      v.OperName,
		DeptName:      v.DeptName,
		OperUrl:       v.OperUrl,
		OperIp:        v.OperIp,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		OperLocation:  v.OperLocation,
		OperParam:     v.OperParam,
		Status:        v.Status,
		OperTime:      timestamppb.New(v.OperTime),
		JsonResult:    v.JsonResult,
		Remark:        v.Remark,
		LatencyTime:   v.LatencyTime,
		UserAgent:     v.UserAgent,
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}