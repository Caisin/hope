// Code generated by Caisin. DO NOT EDIT.
// source: apps/admin/internal/data/ent/schema/sys_user.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/admin/sysuser/v1"
	"hope/apps/admin/internal/data/ent"
)

func SysUserUpdateReq2Data(v *v1.SysUserUpdateReq) *ent.SysUser {
	if v == nil {
		return nil
	}
	return &ent.SysUser{
		ID:       v.Id,
		Username: v.Username,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Status:   v.Status,
	}
}

func SysUserData2UpdateReq(v *ent.SysUser) *v1.SysUserUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.SysUserUpdateReq{
		Id:       v.ID,
		Username: v.Username,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Status:   v.Status,
	}
}

func SysUserCreateReq2Data(v *v1.SysUserCreateReq) *ent.SysUser {
	if v == nil {
		return nil
	}
	return &ent.SysUser{
		Username: v.Username,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Status:   v.Status,
	}
}

func SysUserData2CreateReq(v *ent.SysUser) *v1.SysUserCreateReq {
	if v == nil {
		return nil
	}
	return &v1.SysUserCreateReq{
		Username: v.Username,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Status:   v.Status,
	}
}

func SysUserReq2Data(v *v1.SysUserReq) *ent.SysUser {
	if v == nil {
		return nil
	}
	return &ent.SysUser{
		Username: v.Username,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Status:   v.Status,
	}
}

func SysUserData2Req(v *ent.SysUser) *v1.SysUserReq {
	if v == nil {
		return nil
	}
	return &v1.SysUserReq{
		Username: v.Username,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Status:   v.Status,
	}
}

func SysUserReply2Data(v *v1.SysUserReply) *ent.SysUser {
	if v == nil {
		return nil
	}
	return &ent.SysUser{
		ID:        v.Id,
		Username:  v.Username,
		NickName:  v.NickName,
		Phone:     v.Phone,
		DeptId:    v.DeptId,
		PostId:    v.PostId,
		RoleId:    v.RoleId,
		Avatar:    v.Avatar,
		Sex:       v.Sex,
		Email:     v.Email,
		Remark:    v.Remark,
		Status:    v.Status,
		ExtInfo:   v.ExtInfo,
		CreatedAt: v.CreatedAt.AsTime(),
		UpdatedAt: v.UpdatedAt.AsTime(),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func SysUserData2Reply(v *ent.SysUser) *v1.SysUserReply {
	if v == nil {
		return nil
	}
	return &v1.SysUserReply{
		Id:        v.ID,
		Username:  v.Username,
		NickName:  v.NickName,
		Phone:     v.Phone,
		DeptId:    v.DeptId,
		PostId:    v.PostId,
		RoleId:    v.RoleId,
		Avatar:    v.Avatar,
		Sex:       v.Sex,
		Email:     v.Email,
		Remark:    v.Remark,
		Status:    v.Status,
		ExtInfo:   v.ExtInfo,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func SysUserUpdateReply2Data(v *v1.SysUserUpdateReply) *ent.SysUser {
	if v == nil {
		return nil
	}
	return &ent.SysUser{
		ID:       v.Id,
		Username: v.Username,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Status:   v.Status,
	}
}

func SysUserData2UpdateReply(v *ent.SysUser) *v1.SysUserUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.SysUserUpdateReply{
		Id:       v.ID,
		Username: v.Username,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Status:   v.Status,
	}
}

func SysUserCreateReply2Data(v *v1.SysUserCreateReply) *ent.SysUser {
	if v == nil {
		return nil
	}
	return &ent.SysUser{
		ID:        v.Id,
		Username:  v.Username,
		NickName:  v.NickName,
		Phone:     v.Phone,
		DeptId:    v.DeptId,
		PostId:    v.PostId,
		RoleId:    v.RoleId,
		Avatar:    v.Avatar,
		Sex:       v.Sex,
		Email:     v.Email,
		Remark:    v.Remark,
		Status:    v.Status,
		ExtInfo:   v.ExtInfo,
		CreatedAt: v.CreatedAt.AsTime(),
		UpdatedAt: v.UpdatedAt.AsTime(),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}

func SysUserData2CreateReply(v *ent.SysUser) *v1.SysUserCreateReply {
	if v == nil {
		return nil
	}
	return &v1.SysUserCreateReply{
		Id:        v.ID,
		Username:  v.Username,
		NickName:  v.NickName,
		Phone:     v.Phone,
		DeptId:    v.DeptId,
		PostId:    v.PostId,
		RoleId:    v.RoleId,
		Avatar:    v.Avatar,
		Sex:       v.Sex,
		Email:     v.Email,
		Remark:    v.Remark,
		Status:    v.Status,
		ExtInfo:   v.ExtInfo,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
