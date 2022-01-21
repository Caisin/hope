// Code generated by Caisin. DO NOT EDIT.
// source: apps/novel/internal/data/ent/schema/data_source.go

package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "hope/api/novel/datasource/v1"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/datasource"
)

func DataSourceUpdateReq2Data(v *v1.DataSourceUpdateReq) *ent.DataSource {
	if v == nil {
		return nil
	}
	return &ent.DataSource{
		ID:              v.Id,
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          datasource.DbType(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
	}
}

func DataSourceData2UpdateReq(v *ent.DataSource) *v1.DataSourceUpdateReq {
	if v == nil {
		return nil
	}
	return &v1.DataSourceUpdateReq{
		Id:              v.ID,
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          string(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
	}
}

func DataSourceCreateReq2Data(v *v1.DataSourceCreateReq) *ent.DataSource {
	if v == nil {
		return nil
	}
	return &ent.DataSource{
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          datasource.DbType(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
	}
}

func DataSourceData2CreateReq(v *ent.DataSource) *v1.DataSourceCreateReq {
	if v == nil {
		return nil
	}
	return &v1.DataSourceCreateReq{
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          string(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
	}
}

func DataSourceReq2Data(v *v1.DataSourceReq) *ent.DataSource {
	if v == nil {
		return nil
	}
	return &ent.DataSource{
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          datasource.DbType(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
	}
}

func DataSourceData2Req(v *ent.DataSource) *v1.DataSourceReq {
	if v == nil {
		return nil
	}
	return &v1.DataSourceReq{
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          string(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
	}
}

func DataSourceReply2Data(v *v1.DataSourceReply) *ent.DataSource {
	if v == nil {
		return nil
	}
	return &ent.DataSource{
		ID:              v.Id,
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          datasource.DbType(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
		Remark:          v.Remark,
		CreatedAt:       v.CreatedAt.AsTime(),
		UpdatedAt:       v.UpdatedAt.AsTime(),
		CreateBy:        v.CreateBy,
		UpdateBy:        v.UpdateBy,
		TenantId:        v.TenantId,
	}
}

func DataSourceData2Reply(v *ent.DataSource) *v1.DataSourceReply {
	if v == nil {
		return nil
	}
	return &v1.DataSourceReply{
		Id:              v.ID,
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          string(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
		Remark:          v.Remark,
		CreatedAt:       timestamppb.New(v.CreatedAt),
		UpdatedAt:       timestamppb.New(v.UpdatedAt),
		CreateBy:        v.CreateBy,
		UpdateBy:        v.UpdateBy,
		TenantId:        v.TenantId,
	}
}

func DataSourceUpdateReply2Data(v *v1.DataSourceUpdateReply) *ent.DataSource {
	if v == nil {
		return nil
	}
	return &ent.DataSource{
		ID:              v.Id,
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          datasource.DbType(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
	}
}

func DataSourceData2UpdateReply(v *ent.DataSource) *v1.DataSourceUpdateReply {
	if v == nil {
		return nil
	}
	return &v1.DataSourceUpdateReply{
		Id:              v.ID,
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          string(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
	}
}

func DataSourceCreateReply2Data(v *v1.DataSourceCreateReply) *ent.DataSource {
	if v == nil {
		return nil
	}
	return &ent.DataSource{
		ID:              v.Id,
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          datasource.DbType(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
		Remark:          v.Remark,
		CreatedAt:       v.CreatedAt.AsTime(),
		UpdatedAt:       v.UpdatedAt.AsTime(),
		CreateBy:        v.CreateBy,
		UpdateBy:        v.UpdateBy,
		TenantId:        v.TenantId,
	}
}

func DataSourceData2CreateReply(v *ent.DataSource) *v1.DataSourceCreateReply {
	if v == nil {
		return nil
	}
	return &v1.DataSourceCreateReply{
		Id:              v.ID,
		DbName:          v.DbName,
		Host:            v.Host,
		Port:            v.Port,
		Database:        v.Database,
		UserName:        v.UserName,
		Pwd:             v.Pwd,
		Status:          v.Status,
		DbType:          string(v.DbType),
		ConnMaxIdleTime: v.ConnMaxIdleTime,
		ConnMaxLifeTime: v.ConnMaxLifeTime,
		MaxIdleConns:    v.MaxIdleConns,
		MaxOpenConns:    v.MaxOpenConns,
		Remark:          v.Remark,
		CreatedAt:       timestamppb.New(v.CreatedAt),
		UpdatedAt:       timestamppb.New(v.UpdatedAt),
		CreateBy:        v.CreateBy,
		UpdateBy:        v.UpdateBy,
		TenantId:        v.TenantId,
	}
}