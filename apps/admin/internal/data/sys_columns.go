package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/syscolumns/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/syscolumns"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type sysColumnsRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysColumnsRepo .
func NewSysColumnsRepo(data *Data, logger log.Logger) biz.SysColumnsRepo {
	return &sysColumnsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysColumns 创建
func (r *sysColumnsRepo) CreateSysColumns(ctx context.Context, req *v1.SysColumnsCreateReq) (*ent.SysColumns, error) {
	now := time.Now()
	return r.data.db.SysColumns.Create().
		SetColumnId(req.ColumnId).
		SetColumnName(req.ColumnName).
		SetColumnComment(req.ColumnComment).
		SetColumnType(req.ColumnType).
		SetGoType(req.GoType).
		SetGoField(req.GoField).
		SetJsonField(req.JsonField).
		SetIsPk(req.IsPk).
		SetIsIncrement(req.IsIncrement).
		SetIsRequired(req.IsRequired).
		SetIsInsert(req.IsInsert).
		SetIsEdit(req.IsEdit).
		SetIsList(req.IsList).
		SetIsQuery(req.IsQuery).
		SetQueryType(req.QueryType).
		SetHtmlType(req.HtmlType).
		SetDictType(req.DictType).
		SetSort(req.Sort).
		SetList(req.List).
		SetPk(req.Pk).
		SetRequired(req.Required).
		SetSuperColumn(req.SuperColumn).
		SetUsableColumn(req.UsableColumn).
		SetIncrement(req.Increment).
		SetInsert(req.Insert).
		SetEdit(req.Edit).
		SetQuery(req.Query).
		SetRemark(req.Remark).
		SetFkLabelName(req.FkLabelName).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteSysColumns 删除
func (r *sysColumnsRepo) DeleteSysColumns(ctx context.Context, req *v1.SysColumnsDeleteReq) error {
	return r.data.db.SysColumns.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysColumns 批量删除
func (r *sysColumnsRepo) BatchDeleteSysColumns(ctx context.Context, req *v1.SysColumnsBatchDeleteReq) (int, error) {
	return r.data.db.SysColumns.Delete().Where(syscolumns.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysColumns 更新
func (r *sysColumnsRepo) UpdateSysColumns(ctx context.Context, req *v1.SysColumnsUpdateReq) (*ent.SysColumns, error) {
	return r.data.db.SysColumns.UpdateOne(convert.SysColumnsUpdateReq2Data(req)).Save(ctx)
}

// GetSysColumns 根据Id查询
func (r *sysColumnsRepo) GetSysColumns(ctx context.Context, req *v1.SysColumnsReq) (*ent.SysColumns, error) {
	return r.data.db.SysColumns.Get(ctx, req.Id)
}

// PageSysColumns 分页查询
func (r *sysColumnsRepo) PageSysColumns(ctx context.Context, req *v1.SysColumnsPageReq) ([]*ent.SysColumns, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.SysColumns.
		Query().
		Where(
			//查询条件构造
			r.genCondition(req.Param)...,
		)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	req.GetPagin().Total = int64(count)
	if count == 0 {
		return nil, nil
	}
	query.Limit(int(p.GetPageSize())).
		Offset(int(p.GetOffSet()))
	if p.NeedOrder() {
		if p.IsDesc() {
			query.Order(ent.Desc(p.GetField()))
		} else {
			query.Order(ent.Asc(p.GetField()))
		}
	}
	return query.All(ctx)
}

// genCondition 构造查询条件
func (r *sysColumnsRepo) genCondition(req *v1.SysColumnsReq) []predicate.SysColumns {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysColumns, 0)
	if req.Id > 0 {
		list = append(list, syscolumns.ID(req.Id))
	}
	if req.ColumnId > 0 {
		list = append(list, syscolumns.ColumnId(req.ColumnId))
	}
	if str.IsBlank(req.ColumnName) {
		list = append(list, syscolumns.ColumnNameContains(req.ColumnName))
	}
	if str.IsBlank(req.ColumnComment) {
		list = append(list, syscolumns.ColumnCommentContains(req.ColumnComment))
	}
	if str.IsBlank(req.ColumnType) {
		list = append(list, syscolumns.ColumnTypeContains(req.ColumnType))
	}
	if str.IsBlank(req.GoType) {
		list = append(list, syscolumns.GoTypeContains(req.GoType))
	}
	if str.IsBlank(req.GoField) {
		list = append(list, syscolumns.GoFieldContains(req.GoField))
	}
	if str.IsBlank(req.JsonField) {
		list = append(list, syscolumns.JsonFieldContains(req.JsonField))
	}
	if str.IsBlank(req.IsPk) {
		list = append(list, syscolumns.IsPkContains(req.IsPk))
	}
	if str.IsBlank(req.IsIncrement) {
		list = append(list, syscolumns.IsIncrementContains(req.IsIncrement))
	}
	if str.IsBlank(req.IsRequired) {
		list = append(list, syscolumns.IsRequiredContains(req.IsRequired))
	}
	if str.IsBlank(req.IsInsert) {
		list = append(list, syscolumns.IsInsertContains(req.IsInsert))
	}
	if str.IsBlank(req.IsEdit) {
		list = append(list, syscolumns.IsEditContains(req.IsEdit))
	}
	if str.IsBlank(req.IsList) {
		list = append(list, syscolumns.IsListContains(req.IsList))
	}
	if str.IsBlank(req.IsQuery) {
		list = append(list, syscolumns.IsQueryContains(req.IsQuery))
	}
	if str.IsBlank(req.QueryType) {
		list = append(list, syscolumns.QueryTypeContains(req.QueryType))
	}
	if str.IsBlank(req.HtmlType) {
		list = append(list, syscolumns.HtmlTypeContains(req.HtmlType))
	}
	if str.IsBlank(req.DictType) {
		list = append(list, syscolumns.DictTypeContains(req.DictType))
	}
	if req.Sort > 0 {
		list = append(list, syscolumns.Sort(req.Sort))
	}
	if str.IsBlank(req.List) {
		list = append(list, syscolumns.ListContains(req.List))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, syscolumns.RemarkContains(req.Remark))
	}
	if str.IsBlank(req.FkLabelName) {
		list = append(list, syscolumns.FkLabelNameContains(req.FkLabelName))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, syscolumns.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, syscolumns.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, syscolumns.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, syscolumns.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, syscolumns.TenantId(req.TenantId))
	}

	return list
}
