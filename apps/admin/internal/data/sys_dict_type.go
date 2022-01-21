package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysdicttype/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysdicttype"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type sysDictTypeRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysDictTypeRepo .
func NewSysDictTypeRepo(data *Data, logger log.Logger) biz.SysDictTypeRepo {
	return &sysDictTypeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysDictType 创建
func (r *sysDictTypeRepo) CreateSysDictType(ctx context.Context, req *v1.SysDictTypeCreateReq) (*ent.SysDictType, error) {
	now := time.Now()
	return r.data.db.SysDictType.Create().
		SetDictName(req.DictName).
		SetTypeCode(req.TypeCode).
		SetStatus(req.Status).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteSysDictType 删除
func (r *sysDictTypeRepo) DeleteSysDictType(ctx context.Context, req *v1.SysDictTypeDeleteReq) error {
	return r.data.db.SysDictType.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysDictType 批量删除
func (r *sysDictTypeRepo) BatchDeleteSysDictType(ctx context.Context, req *v1.SysDictTypeBatchDeleteReq) (int, error) {
	return r.data.db.SysDictType.Delete().Where(sysdicttype.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysDictType 更新
func (r *sysDictTypeRepo) UpdateSysDictType(ctx context.Context, req *v1.SysDictTypeUpdateReq) (*ent.SysDictType, error) {
	return r.data.db.SysDictType.UpdateOne(convert.SysDictTypeUpdateReq2Data(req)).Save(ctx)
}

// GetSysDictType 根据Id查询
func (r *sysDictTypeRepo) GetSysDictType(ctx context.Context, req *v1.SysDictTypeReq) (*ent.SysDictType, error) {
	return r.data.db.SysDictType.Get(ctx, req.Id)
}

// PageSysDictType 分页查询
func (r *sysDictTypeRepo) PageSysDictType(ctx context.Context, req *v1.SysDictTypePageReq) ([]*ent.SysDictType, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.SysDictType.
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
func (r *sysDictTypeRepo) genCondition(req *v1.SysDictTypeReq) []predicate.SysDictType {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysDictType, 0)
	if req.Id > 0 {
		list = append(list, sysdicttype.ID(req.Id))
	}
	if str.IsBlank(req.DictName) {
		list = append(list, sysdicttype.DictNameContains(req.DictName))
	}
	if str.IsBlank(req.TypeCode) {
		list = append(list, sysdicttype.TypeCodeContains(req.TypeCode))
	}
	if req.Status > 0 {
		list = append(list, sysdicttype.Status(req.Status))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, sysdicttype.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysdicttype.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysdicttype.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysdicttype.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysdicttype.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysdicttype.TenantId(req.TenantId))
	}

	return list
}
