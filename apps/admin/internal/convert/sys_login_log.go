// Code generated by Caisin. DO NOT EDIT.
// source: apps/admin/internal/data/ent/schema/sys_login_log.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/admin/sysloginlog/v1"
	"hope/apps/admin/internal/data/ent"
)

func SysLoginLogUpdateReq2Data(v *v1.SysLoginLogUpdateReq) *ent.SysLoginLog {
	if v == nil {
		return nil
	}
	return &ent.SysLoginLog{
		ID:            v.Id,
		UserId:        v.UserId,
		Status:        v.Status,
		Ipaddr:        v.Ipaddr,
		LoginLocation: v.LoginLocation,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		LoginTime:     v.LoginTime.AsTime(),
		Remark:        v.Remark,
		Msg:           v.Msg,
	}
}

func SysLoginLogData2UpdateReq(v *ent.SysLoginLog) *v1.SysLoginLogUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.SysLoginLogUpdateReq{
		Id:            v.ID,
		UserId:        v.UserId,
		Status:        v.Status,
		Ipaddr:        v.Ipaddr,
		LoginLocation: v.LoginLocation,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		LoginTime:     timestamppb.New(v.LoginTime),
		Remark:        v.Remark,
		Msg:           v.Msg,
	}
}

func SysLoginLogCreateReq2Data(v *v1.SysLoginLogCreateReq) *ent.SysLoginLog {
	if v == nil {
		return nil
	}
	return &ent.SysLoginLog{
		UserId:        v.UserId,
		Status:        v.Status,
		Ipaddr:        v.Ipaddr,
		LoginLocation: v.LoginLocation,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		LoginTime:     v.LoginTime.AsTime(),
		Remark:        v.Remark,
		Msg:           v.Msg,
	}
}

func SysLoginLogData2CreateReq(v *ent.SysLoginLog) *v1.SysLoginLogCreateReq {
	if v == nil {
		return nil
	}
	return &v1.SysLoginLogCreateReq{
		UserId:        v.UserId,
		Status:        v.Status,
		Ipaddr:        v.Ipaddr,
		LoginLocation: v.LoginLocation,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		LoginTime:     timestamppb.New(v.LoginTime),
		Remark:        v.Remark,
		Msg:           v.Msg,
	}
}

func SysLoginLogReq2Data(v *v1.SysLoginLogReq) *ent.SysLoginLog {
	if v == nil {
		return nil
	}
	return &ent.SysLoginLog{
		UserId:        v.UserId,
		Status:        v.Status,
		Ipaddr:        v.Ipaddr,
		LoginLocation: v.LoginLocation,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		LoginTime:     v.LoginTime.AsTime(),
		Remark:        v.Remark,
		Msg:           v.Msg,
	}
}

func SysLoginLogData2Req(v *ent.SysLoginLog) *v1.SysLoginLogReq {
	if v == nil {
		return nil
	}
	return &v1.SysLoginLogReq{
		UserId:        v.UserId,
		Status:        v.Status,
		Ipaddr:        v.Ipaddr,
		LoginLocation: v.LoginLocation,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		LoginTime:     timestamppb.New(v.LoginTime),
		Remark:        v.Remark,
		Msg:           v.Msg,
	}
}

func SysLoginLogData2Reply(v *ent.SysLoginLog) *v1.SysLoginLogData {
	if v == nil {
		return nil
	}
	return &v1.SysLoginLogData{
		Id:            v.ID,
		UserId:        v.UserId,
		Status:        v.Status,
		Ipaddr:        v.Ipaddr,
		LoginLocation: v.LoginLocation,
		Browser:       v.Browser,
		Os:            v.Os,
		Platform:      v.Platform,
		LoginTime:     timestamppb.New(v.LoginTime),
		Remark:        v.Remark,
		Msg:           v.Msg,
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}
