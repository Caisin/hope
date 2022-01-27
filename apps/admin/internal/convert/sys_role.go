// Code generated by Caisin. DO NOT EDIT.
// source: apps/admin/internal/data/ent/schema/sys_role.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/admin/sysrole/v1"
	"hope/apps/admin/internal/data/ent"
)

func SysRoleUpdateReq2Data(v *v1.SysRoleUpdateReq) *ent.SysRole {
	if v == nil {
		return nil
	}
	return &ent.SysRole{
		ID:        v.Id,
		RoleName:  v.RoleName,
		Status:    v.Status,
		RoleKey:   v.RoleKey,
		RoleSort:  v.RoleSort,
		Flag:      v.Flag,
		Remark:    v.Remark,
		Admin:     v.Admin,
		DataScope: v.DataScope,
		SysDept:   v.SysDept,
		SysMenu:   v.SysMenu,
	}
}

func SysRoleData2UpdateReq(v *ent.SysRole) *v1.SysRoleUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.SysRoleUpdateReq{
		Id:        v.ID,
		RoleName:  v.RoleName,
		Status:    v.Status,
		RoleKey:   v.RoleKey,
		RoleSort:  v.RoleSort,
		Flag:      v.Flag,
		Remark:    v.Remark,
		Admin:     v.Admin,
		DataScope: v.DataScope,
		SysDept:   v.SysDept,
		SysMenu:   v.SysMenu,
	}
}

func SysRoleCreateReq2Data(v *v1.SysRoleCreateReq) *ent.SysRole {
	if v == nil {
		return nil
	}
	return &ent.SysRole{
		RoleName:  v.RoleName,
		Status:    v.Status,
		RoleKey:   v.RoleKey,
		RoleSort:  v.RoleSort,
		Flag:      v.Flag,
		Remark:    v.Remark,
		Admin:     v.Admin,
		DataScope: v.DataScope,
		SysDept:   v.SysDept,
		SysMenu:   v.SysMenu,
	}
}

func SysRoleData2CreateReq(v *ent.SysRole) *v1.SysRoleCreateReq {
	if v == nil {
		return nil
	}
	return &v1.SysRoleCreateReq{
		RoleName:  v.RoleName,
		Status:    v.Status,
		RoleKey:   v.RoleKey,
		RoleSort:  v.RoleSort,
		Flag:      v.Flag,
		Remark:    v.Remark,
		Admin:     v.Admin,
		DataScope: v.DataScope,
		SysDept:   v.SysDept,
		SysMenu:   v.SysMenu,
	}
}

func SysRoleReq2Data(v *v1.SysRoleReq) *ent.SysRole {
	if v == nil {
		return nil
	}
	return &ent.SysRole{
		RoleName:  v.RoleName,
		Status:    v.Status,
		RoleKey:   v.RoleKey,
		RoleSort:  v.RoleSort,
		Flag:      v.Flag,
		Remark:    v.Remark,
		Admin:     v.Admin,
		DataScope: v.DataScope,
		SysDept:   v.SysDept,
		SysMenu:   v.SysMenu,
	}
}

func SysRoleData2Req(v *ent.SysRole) *v1.SysRoleReq {
	if v == nil {
		return nil
	}
	return &v1.SysRoleReq{
		RoleName:  v.RoleName,
		Status:    v.Status,
		RoleKey:   v.RoleKey,
		RoleSort:  v.RoleSort,
		Flag:      v.Flag,
		Remark:    v.Remark,
		Admin:     v.Admin,
		DataScope: v.DataScope,
		SysDept:   v.SysDept,
		SysMenu:   v.SysMenu,
	}
}

func SysRoleData2Reply(v *ent.SysRole) *v1.SysRoleData {
	if v == nil {
		return nil
	}
	return &v1.SysRoleData{
		Id:        v.ID,
		RoleName:  v.RoleName,
		Status:    v.Status,
		RoleKey:   v.RoleKey,
		RoleSort:  v.RoleSort,
		Flag:      v.Flag,
		Remark:    v.Remark,
		Admin:     v.Admin,
		DataScope: v.DataScope,
		SysDept:   v.SysDept,
		SysMenu:   v.SysMenu,
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
		CreateBy:  v.CreateBy,
		UpdateBy:  v.UpdateBy,
		TenantId:  v.TenantId,
	}
}
