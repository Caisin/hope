package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/customernovelconfig/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/customernovelconfig"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"
	"time"
)

type customerNovelConfigRepo struct {
	data *Data
	log  *log.Helper
}

// NewCustomerNovelConfigRepo .
func NewCustomerNovelConfigRepo(data *Data, logger log.Logger) biz.CustomerNovelConfigRepo {
	return &customerNovelConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateCustomerNovelConfig 创建
func (r *customerNovelConfigRepo) CreateCustomerNovelConfig(ctx context.Context, req *v1.CustomerNovelConfigCreateReq) (*ent.CustomerNovelConfig, error) {
	now := time.Now()
	return r.data.db.CustomerNovelConfig.Create().
		SetGroupCode(req.GroupCode).
		SetInnerGroupCode(req.InnerGroupCode).
		SetGroupName(req.GroupName).
		SetTypeId(req.TypeId).
		SetTypeCode(req.TypeCode).
		SetTypeName(req.TypeName).
		SetFieldName(req.FieldName).
		SetDefaultNum(req.DefaultNum).
		SetState(req.State).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteCustomerNovelConfig 删除
func (r *customerNovelConfigRepo) DeleteCustomerNovelConfig(ctx context.Context, req *v1.CustomerNovelConfigDeleteReq) error {
	return r.data.db.CustomerNovelConfig.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteCustomerNovelConfig 批量删除
func (r *customerNovelConfigRepo) BatchDeleteCustomerNovelConfig(ctx context.Context, req *v1.CustomerNovelConfigBatchDeleteReq) (int, error) {
	return r.data.db.CustomerNovelConfig.Delete().Where(customernovelconfig.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateCustomerNovelConfig 更新
func (r *customerNovelConfigRepo) UpdateCustomerNovelConfig(ctx context.Context, req *v1.CustomerNovelConfigUpdateReq) (*ent.CustomerNovelConfig, error) {
	return r.data.db.CustomerNovelConfig.UpdateOne(convert.CustomerNovelConfigUpdateReq2Data(req)).Save(ctx)
}

// GetCustomerNovelConfig 根据Id查询
func (r *customerNovelConfigRepo) GetCustomerNovelConfig(ctx context.Context, req *v1.CustomerNovelConfigReq) (*ent.CustomerNovelConfig, error) {
	return r.data.db.CustomerNovelConfig.Get(ctx, req.Id)
}

// PageCustomerNovelConfig 分页查询
func (r *customerNovelConfigRepo) PageCustomerNovelConfig(ctx context.Context, req *v1.CustomerNovelConfigPageReq) ([]*ent.CustomerNovelConfig, error) {
	pagin := req.Pagin
	query := r.data.db.CustomerNovelConfig.
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
func (r *customerNovelConfigRepo) genCondition(req *v1.CustomerNovelConfigReq) []predicate.CustomerNovelConfig {
	list := make([]predicate.CustomerNovelConfig, 0)
	if req.Id > 0 {
		list = append(list, customernovelconfig.ID(req.Id))
	}
	if str.IsBlank(req.GroupCode) {
		list = append(list, customernovelconfig.GroupCodeContains(req.GroupCode))
	}
	if str.IsBlank(req.InnerGroupCode) {
		list = append(list, customernovelconfig.InnerGroupCodeContains(req.InnerGroupCode))
	}
	if str.IsBlank(req.GroupName) {
		list = append(list, customernovelconfig.GroupNameContains(req.GroupName))
	}
	if req.TypeId > 0 {
		list = append(list, customernovelconfig.TypeId(req.TypeId))
	}
	if str.IsBlank(req.TypeCode) {
		list = append(list, customernovelconfig.TypeCodeContains(req.TypeCode))
	}
	if str.IsBlank(req.TypeName) {
		list = append(list, customernovelconfig.TypeNameContains(req.TypeName))
	}
	if str.IsBlank(req.FieldName) {
		list = append(list, customernovelconfig.FieldNameContains(req.FieldName))
	}
	if req.DefaultNum > 0 {
		list = append(list, customernovelconfig.DefaultNum(req.DefaultNum))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, customernovelconfig.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, customernovelconfig.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, customernovelconfig.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, customernovelconfig.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, customernovelconfig.TenantId(req.TenantId))
	}

	return list
}
