package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/systables/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/systables"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type sysTablesRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysTablesRepo .
func NewSysTablesRepo(data *Data, logger log.Logger) biz.SysTablesRepo {
	return &sysTablesRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysTables 创建
func (r *sysTablesRepo) CreateSysTables(ctx context.Context, req *v1.SysTablesCreateReq) (*ent.SysTables, error) {
	now := time.Now()
	return r.data.db.SysTables.Create().
		SetTableName(req.TableName).
		SetTableComment(req.TableComment).
		SetClassName(req.ClassName).
		SetTplCategory(req.TplCategory).
		SetPackageName(req.PackageName).
		SetModuleName(req.ModuleName).
		SetModuleFrontName(req.ModuleFrontName).
		SetBusinessName(req.BusinessName).
		SetFunctionName(req.FunctionName).
		SetFunctionAuthor(req.FunctionAuthor).
		SetPkColumn(req.PkColumn).
		SetPkGoField(req.PkGoField).
		SetPkJsonField(req.PkJsonField).
		SetOptions(req.Options).
		SetTreeCode(req.TreeCode).
		SetTreeParentCode(req.TreeParentCode).
		SetTreeName(req.TreeName).
		SetTree(req.Tree).
		SetCrud(req.Crud).
		SetRemark(req.Remark).
		SetIsDataScope(req.IsDataScope).
		SetIsActions(req.IsActions).
		SetIsAuth(req.IsAuth).
		SetIsLogicalDelete(req.IsLogicalDelete).
		SetLogicalDelete(req.LogicalDelete).
		SetLogicalDeleteColumn(req.LogicalDeleteColumn).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteSysTables 删除
func (r *sysTablesRepo) DeleteSysTables(ctx context.Context, req *v1.SysTablesDeleteReq) error {
	return r.data.db.SysTables.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysTables 批量删除
func (r *sysTablesRepo) BatchDeleteSysTables(ctx context.Context, req *v1.SysTablesBatchDeleteReq) (int, error) {
	return r.data.db.SysTables.Delete().Where(systables.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysTables 更新
func (r *sysTablesRepo) UpdateSysTables(ctx context.Context, req *v1.SysTablesUpdateReq) (*ent.SysTables, error) {
	return r.data.db.SysTables.UpdateOne(convert.SysTablesUpdateReq2Data(req)).Save(ctx)
}

// GetSysTables 根据Id查询
func (r *sysTablesRepo) GetSysTables(ctx context.Context, req *v1.SysTablesReq) (*ent.SysTables, error) {
	return r.data.db.SysTables.Get(ctx, req.Id)
}

// PageSysTables 分页查询
func (r *sysTablesRepo) PageSysTables(ctx context.Context, req *v1.SysTablesPageReq) ([]*ent.SysTables, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.SysTables.
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
func (r *sysTablesRepo) genCondition(req *v1.SysTablesReq) []predicate.SysTables {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysTables, 0)
	if req.Id > 0 {
		list = append(list, systables.ID(req.Id))
	}
	if str.IsBlank(req.TableName) {
		list = append(list, systables.TableNameContains(req.TableName))
	}
	if str.IsBlank(req.TableComment) {
		list = append(list, systables.TableCommentContains(req.TableComment))
	}
	if str.IsBlank(req.ClassName) {
		list = append(list, systables.ClassNameContains(req.ClassName))
	}
	if str.IsBlank(req.TplCategory) {
		list = append(list, systables.TplCategoryContains(req.TplCategory))
	}
	if str.IsBlank(req.PackageName) {
		list = append(list, systables.PackageNameContains(req.PackageName))
	}
	if str.IsBlank(req.ModuleName) {
		list = append(list, systables.ModuleNameContains(req.ModuleName))
	}
	if str.IsBlank(req.ModuleFrontName) {
		list = append(list, systables.ModuleFrontNameContains(req.ModuleFrontName))
	}
	if str.IsBlank(req.BusinessName) {
		list = append(list, systables.BusinessNameContains(req.BusinessName))
	}
	if str.IsBlank(req.FunctionName) {
		list = append(list, systables.FunctionNameContains(req.FunctionName))
	}
	if str.IsBlank(req.FunctionAuthor) {
		list = append(list, systables.FunctionAuthorContains(req.FunctionAuthor))
	}
	if str.IsBlank(req.PkColumn) {
		list = append(list, systables.PkColumnContains(req.PkColumn))
	}
	if str.IsBlank(req.PkGoField) {
		list = append(list, systables.PkGoFieldContains(req.PkGoField))
	}
	if str.IsBlank(req.PkJsonField) {
		list = append(list, systables.PkJsonFieldContains(req.PkJsonField))
	}
	if str.IsBlank(req.Options) {
		list = append(list, systables.OptionsContains(req.Options))
	}
	if str.IsBlank(req.TreeCode) {
		list = append(list, systables.TreeCodeContains(req.TreeCode))
	}
	if str.IsBlank(req.TreeParentCode) {
		list = append(list, systables.TreeParentCodeContains(req.TreeParentCode))
	}
	if str.IsBlank(req.TreeName) {
		list = append(list, systables.TreeNameContains(req.TreeName))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, systables.RemarkContains(req.Remark))
	}
	if req.IsDataScope > 0 {
		list = append(list, systables.IsDataScope(req.IsDataScope))
	}
	if req.IsActions > 0 {
		list = append(list, systables.IsActions(req.IsActions))
	}
	if req.IsAuth > 0 {
		list = append(list, systables.IsAuth(req.IsAuth))
	}
	if str.IsBlank(req.IsLogicalDelete) {
		list = append(list, systables.IsLogicalDeleteContains(req.IsLogicalDelete))
	}
	if str.IsBlank(req.LogicalDeleteColumn) {
		list = append(list, systables.LogicalDeleteColumnContains(req.LogicalDeleteColumn))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, systables.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, systables.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, systables.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, systables.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, systables.TenantId(req.TenantId))
	}

	return list
}
