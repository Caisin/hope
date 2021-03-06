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
		Password: v.Password,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Desc:     v.Desc,
		HomePath: v.HomePath,
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
		Password: v.Password,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Desc:     v.Desc,
		HomePath: v.HomePath,
		Status:   v.Status,
	}
}

func SysUserCreateReq2Data(v *v1.SysUserCreateReq) *ent.SysUser {
	if v == nil {
		return nil
	}
	return &ent.SysUser{
		Username: v.Username,
		Password: v.Password,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Desc:     v.Desc,
		HomePath: v.HomePath,
		Status:   v.Status,
	}
}

func SysUserData2CreateReq(v *ent.SysUser) *v1.SysUserCreateReq {
	if v == nil {
		return nil
	}
	return &v1.SysUserCreateReq{
		Username: v.Username,
		Password: v.Password,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Desc:     v.Desc,
		HomePath: v.HomePath,
		Status:   v.Status,
	}
}

func SysUserReq2Data(v *v1.SysUserReq) *ent.SysUser {
	if v == nil {
		return nil
	}
	return &ent.SysUser{
		Username: v.Username,
		Password: v.Password,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Desc:     v.Desc,
		HomePath: v.HomePath,
		Status:   v.Status,
	}
}

func SysUserData2Req(v *ent.SysUser) *v1.SysUserReq {
	if v == nil {
		return nil
	}
	return &v1.SysUserReq{
		Username: v.Username,
		Password: v.Password,
		NickName: v.NickName,
		Phone:    v.Phone,
		DeptId:   v.DeptId,
		PostId:   v.PostId,
		RoleId:   v.RoleId,
		Avatar:   v.Avatar,
		Sex:      v.Sex,
		Email:    v.Email,
		Remark:   v.Remark,
		Desc:     v.Desc,
		HomePath: v.HomePath,
		Status:   v.Status,
	}
}

func SysUserData2Reply(v *ent.SysUser) *v1.SysUserData {
	if v == nil {
		return nil
	}
	return &v1.SysUserData{
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
		Desc:      v.Desc,
		HomePath:  v.HomePath,
		Status:    v.Status,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
