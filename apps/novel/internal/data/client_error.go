package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/clienterror/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/clienterror"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type clientErrorRepo struct {
	data *Data
	log  *log.Helper
}

// NewClientErrorRepo .
func NewClientErrorRepo(data *Data, logger log.Logger) biz.ClientErrorRepo {
	return &clientErrorRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateClientError 创建
func (r *clientErrorRepo) CreateClientError(ctx context.Context, req *v1.ClientErrorCreateReq) (*ent.ClientError, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.ClientError.Create().
		SetAppVersion(req.AppVersion).
		SetDeviceName(req.DeviceName).
		SetOsName(req.OsName).
		SetErrorInfo(req.ErrorInfo).
		SetUserId(req.UserId).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteClientError 删除
func (r *clientErrorRepo) DeleteClientError(ctx context.Context, req *v1.ClientErrorDeleteReq) error {
	return r.data.db.ClientError.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteClientError 批量删除
func (r *clientErrorRepo) BatchDeleteClientError(ctx context.Context, req *v1.ClientErrorBatchDeleteReq) (int, error) {
	return r.data.db.ClientError.Delete().Where(clienterror.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateClientError 更新
func (r *clientErrorRepo) UpdateClientError(ctx context.Context, req *v1.ClientErrorUpdateReq) (*ent.ClientError, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.ClientError.UpdateOneID(req.Id).
		SetAppVersion(req.AppVersion).
		SetDeviceName(req.DeviceName).
		SetOsName(req.OsName).
		SetErrorInfo(req.ErrorInfo).
		SetUserId(req.UserId).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetClientError 根据Id查询
func (r *clientErrorRepo) GetClientError(ctx context.Context, req *v1.ClientErrorReq) (*ent.ClientError, error) {
	return r.data.db.ClientError.Get(ctx, req.Id)
}

// PageClientError 分页查询
func (r *clientErrorRepo) PageClientError(ctx context.Context, req *v1.ClientErrorPageReq) ([]*ent.ClientError, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.ClientError.
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
func (r *clientErrorRepo) genCondition(req *v1.ClientErrorReq) []predicate.ClientError {
	if req == nil {
		return nil
	}
	list := make([]predicate.ClientError, 0)
	if req.Id > 0 {
		list = append(list, clienterror.ID(req.Id))
	}
	if str.IsBlank(req.AppVersion) {
		list = append(list, clienterror.AppVersionContains(req.AppVersion))
	}
	if str.IsBlank(req.DeviceName) {
		list = append(list, clienterror.DeviceNameContains(req.DeviceName))
	}
	if str.IsBlank(req.OsName) {
		list = append(list, clienterror.OsNameContains(req.OsName))
	}
	if str.IsBlank(req.ErrorInfo) {
		list = append(list, clienterror.ErrorInfoContains(req.ErrorInfo))
	}
	if req.UserId > 0 {
		list = append(list, clienterror.UserId(req.UserId))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, clienterror.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, clienterror.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, clienterror.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, clienterror.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, clienterror.TenantId(req.TenantId))
	}

	return list
}
