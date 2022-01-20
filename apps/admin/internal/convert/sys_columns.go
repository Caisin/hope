// Code generated by Caisin. DO NOT EDIT.
// source: apps/admin/internal/data/ent/schema/sys_columns.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/admin/syscolumns/v1"
	"hope/apps/admin/internal/data/ent"
)

func SysColumnsUpdateReq2Data(v *v1.SysColumnsUpdateReq) *ent.SysColumns {
	if v == nil {
		return nil
	}
	return &ent.SysColumns{
		ID:            v.Id,
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
	}
}

func SysColumnsData2UpdateReq(v *ent.SysColumns) *v1.SysColumnsUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.SysColumnsUpdateReq{
		Id:            v.ID,
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
	}
}

func SysColumnsCreateReq2Data(v *v1.SysColumnsCreateReq) *ent.SysColumns {
	if v == nil {
		return nil
	}
	return &ent.SysColumns{
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
	}
}

func SysColumnsData2CreateReq(v *ent.SysColumns) *v1.SysColumnsCreateReq {
	if v == nil {
		return nil
	}
	return &v1.SysColumnsCreateReq{
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
	}
}

func SysColumnsReq2Data(v *v1.SysColumnsReq) *ent.SysColumns {
	if v == nil {
		return nil
	}
	return &ent.SysColumns{
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
	}
}

func SysColumnsData2Req(v *ent.SysColumns) *v1.SysColumnsReq {
	if v == nil {
		return nil
	}
	return &v1.SysColumnsReq{
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
	}
}

func SysColumnsReply2Data(v *v1.SysColumnsReply) *ent.SysColumns {
	if v == nil {
		return nil
	}
	return &ent.SysColumns{
		ID:            v.Id,
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
		FkLabelName:   v.FkLabelName,
		CreatedAt:     v.CreatedAt.AsTime(),
		UpdatedAt:     v.UpdatedAt.AsTime(),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func SysColumnsData2Reply(v *ent.SysColumns) *v1.SysColumnsReply {
	if v == nil {
		return nil
	}
	return &v1.SysColumnsReply{
		Id:            v.ID,
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
		FkLabelName:   v.FkLabelName,
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func SysColumnsUpdateReply2Data(v *v1.SysColumnsUpdateReply) *ent.SysColumns {
	if v == nil {
		return nil
	}
	return &ent.SysColumns{
		ID:            v.Id,
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
	}
}

func SysColumnsData2UpdateReply(v *ent.SysColumns) *v1.SysColumnsUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.SysColumnsUpdateReply{
		Id:            v.ID,
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
	}
}

func SysColumnsCreateReply2Data(v *v1.SysColumnsCreateReply) *ent.SysColumns {
	if v == nil {
		return nil
	}
	return &ent.SysColumns{
		ID:            v.Id,
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
		FkLabelName:   v.FkLabelName,
		CreatedAt:     v.CreatedAt.AsTime(),
		UpdatedAt:     v.UpdatedAt.AsTime(),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}

func SysColumnsData2CreateReply(v *ent.SysColumns) *v1.SysColumnsCreateReply {
	if v == nil {
		return nil
	}
	return &v1.SysColumnsCreateReply{
		Id:            v.ID,
		ColumnId:      v.ColumnId,
		ColumnName:    v.ColumnName,
		ColumnComment: v.ColumnComment,
		ColumnType:    v.ColumnType,
		GoType:        v.GoType,
		GoField:       v.GoField,
		JsonField:     v.JsonField,
		IsPk:          v.IsPk,
		IsIncrement:   v.IsIncrement,
		IsRequired:    v.IsRequired,
		IsInsert:      v.IsInsert,
		IsEdit:        v.IsEdit,
		IsList:        v.IsList,
		IsQuery:       v.IsQuery,
		QueryType:     v.QueryType,
		HtmlType:      v.HtmlType,
		DictType:      v.DictType,
		Sort:          v.Sort,
		List:          v.List,
		Pk:            v.Pk,
		Required:      v.Required,
		SuperColumn:   v.SuperColumn,
		UsableColumn:  v.UsableColumn,
		Increment:     v.Increment,
		Insert:        v.Insert,
		Edit:          v.Edit,
		Query:         v.Query,
		Remark:        v.Remark,
		FkLabelName:   v.FkLabelName,
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
		CreateBy:      v.CreateBy,
		UpdateBy:      v.UpdateBy,
		TenantId:      v.TenantId,
	}
}
