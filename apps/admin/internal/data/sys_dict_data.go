package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysdictdata/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysdictdata"
	"hope/pkg/util/str"
	"time"
)

type sysDictDataRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysDictDataRepo .
func NewSysDictDataRepo(data *Data, logger log.Logger) biz.SysDictDataRepo {
	return &sysDictDataRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysDictData 创建
func (r *sysDictDataRepo) CreateSysDictData(ctx context.Context, req *v1.SysDictDataCreateReq) (*ent.SysDictData, error) {
	now := time.Now()
	return r.data.db.SysDictData.Create().
		SetDictSort(req.DictSort).
		SetDictLabel(req.DictLabel).
		SetDictValue(req.DictValue).
		SetIsDefault(req.IsDefault).
		SetStatus(req.Status).
		SetDefault(req.Default).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteSysDictData 删除
func (r *sysDictDataRepo) DeleteSysDictData(ctx context.Context, req *v1.SysDictDataDeleteReq) error {
	return r.data.db.SysDictData.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysDictData 批量删除
func (r *sysDictDataRepo) BatchDeleteSysDictData(ctx context.Context, req *v1.SysDictDataBatchDeleteReq) (int, error) {
	return r.data.db.SysDictData.Delete().Where(sysdictdata.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysDictData 更新
func (r *sysDictDataRepo) UpdateSysDictData(ctx context.Context, req *v1.SysDictDataUpdateReq) (*ent.SysDictData, error) {
	return r.data.db.SysDictData.UpdateOne(convert.SysDictDataUpdateReq2Data(req)).Save(ctx)
}

// GetSysDictData 根据Id查询
func (r *sysDictDataRepo) GetSysDictData(ctx context.Context, req *v1.SysDictDataReq) (*ent.SysDictData, error) {
	return r.data.db.SysDictData.Get(ctx, req.Id)
}

// PageSysDictData 分页查询
func (r *sysDictDataRepo) PageSysDictData(ctx context.Context, req *v1.SysDictDataPageReq) ([]*ent.SysDictData, error) {
	pagin := req.Pagin
	query := r.data.db.SysDictData.
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
	query.Limit(int(pagin.GetPage())).
		Offset(int(pagin.GetOffSet()))
	if pagin.NeedOrder() {
		if pagin.IsDesc() {
			query.Order(ent.Desc(pagin.GetField()))
		} else {
			query.Order(ent.Asc(pagin.GetField()))
		}
	}
	return query.All(ctx)
}

// genCondition 构造查询条件
func (r *sysDictDataRepo) genCondition(req *v1.SysDictDataReq) []predicate.SysDictData {
	list := make([]predicate.SysDictData, 0)
	if req.Id > 0 {
		list = append(list, sysdictdata.ID(req.Id))
	}
	if req.DictSort > 0 {
		list = append(list, sysdictdata.DictSort(req.DictSort))
	}
	if str.IsBlank(req.DictLabel) {
		list = append(list, sysdictdata.DictLabelContains(req.DictLabel))
	}
	if str.IsBlank(req.DictValue) {
		list = append(list, sysdictdata.DictValueContains(req.DictValue))
	}
	if str.IsBlank(req.IsDefault) {
		list = append(list, sysdictdata.IsDefaultContains(req.IsDefault))
	}
	if req.Status > 0 {
		list = append(list, sysdictdata.Status(req.Status))
	}
	if str.IsBlank(req.Default) {
		list = append(list, sysdictdata.DefaultContains(req.Default))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, sysdictdata.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysdictdata.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysdictdata.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysdictdata.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysdictdata.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysdictdata.TenantId(req.TenantId))
	}

	return list
}
