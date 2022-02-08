package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysdept/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysdept"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type sysDeptRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysDeptRepo .
func NewSysDeptRepo(data *Data, logger log.Logger) biz.SysDeptRepo {
	return &sysDeptRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysDept 创建
func (r *sysDeptRepo) CreateSysDept(ctx context.Context, req *v1.SysDeptCreateReq) (*ent.SysDept, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.SysDept.Create().
		SetDeptPath(req.DeptPath).
		SetDeptName(req.DeptName).
		SetSort(req.Sort).
		SetLeader(req.Leader).
		SetPhone(req.Phone).
		SetEmail(req.Email).
		SetStatus(req.Status).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteSysDept 删除
func (r *sysDeptRepo) DeleteSysDept(ctx context.Context, req *v1.SysDeptDeleteReq) error {
	return r.data.db.SysDept.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysDept 批量删除
func (r *sysDeptRepo) BatchDeleteSysDept(ctx context.Context, req *v1.SysDeptBatchDeleteReq) (int, error) {
	return r.data.db.SysDept.Delete().Where(sysdept.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysDept 更新
func (r *sysDeptRepo) UpdateSysDept(ctx context.Context, req *v1.SysDeptUpdateReq) (*ent.SysDept, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.SysDept.UpdateOneID(req.Id).
		SetDeptPath(req.DeptPath).
		SetDeptName(req.DeptName).
		SetSort(req.Sort).
		SetLeader(req.Leader).
		SetPhone(req.Phone).
		SetEmail(req.Email).
		SetStatus(req.Status).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetSysDept 根据Id查询
func (r *sysDeptRepo) GetSysDept(ctx context.Context, req *v1.SysDeptReq) (*ent.SysDept, error) {
	return r.data.db.SysDept.Get(ctx, req.Id)
}

// PageSysDept 分页查询
func (r *sysDeptRepo) PageSysDept(ctx context.Context, req *v1.SysDeptPageReq) ([]*ent.SysDept, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.SysDept.
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
func (r *sysDeptRepo) genCondition(req *v1.SysDeptReq) []predicate.SysDept {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysDept, 0)
	if req.Id > 0 {
		list = append(list, sysdept.ID(req.Id))
	}
	if str.IsNotBlank(req.DeptPath) {
		list = append(list, sysdept.DeptPathContains(req.DeptPath))
	}
	if str.IsNotBlank(req.DeptName) {
		list = append(list, sysdept.DeptNameContains(req.DeptName))
	}
	if req.Sort > 0 {
		list = append(list, sysdept.Sort(req.Sort))
	}
	if str.IsNotBlank(req.Leader) {
		list = append(list, sysdept.LeaderContains(req.Leader))
	}
	if str.IsNotBlank(req.Phone) {
		list = append(list, sysdept.PhoneContains(req.Phone))
	}
	if str.IsNotBlank(req.Email) {
		list = append(list, sysdept.EmailContains(req.Email))
	}
	if req.Status > 0 {
		list = append(list, sysdept.Status(req.Status))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysdept.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysdept.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysdept.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysdept.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysdept.TenantId(req.TenantId))
	}

	return list
}
