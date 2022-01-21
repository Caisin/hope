package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/syspost/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/syspost"
	"hope/pkg/pagin"
	"hope/pkg/util/str"
	"time"
)

type sysPostRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysPostRepo .
func NewSysPostRepo(data *Data, logger log.Logger) biz.SysPostRepo {
	return &sysPostRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysPost 创建
func (r *sysPostRepo) CreateSysPost(ctx context.Context, req *v1.SysPostCreateReq) (*ent.SysPost, error) {
	now := time.Now()
	return r.data.db.SysPost.Create().
		SetPostName(req.PostName).
		SetPostCode(req.PostCode).
		SetSort(req.Sort).
		SetStatus(req.Status).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteSysPost 删除
func (r *sysPostRepo) DeleteSysPost(ctx context.Context, req *v1.SysPostDeleteReq) error {
	return r.data.db.SysPost.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysPost 批量删除
func (r *sysPostRepo) BatchDeleteSysPost(ctx context.Context, req *v1.SysPostBatchDeleteReq) (int, error) {
	return r.data.db.SysPost.Delete().Where(syspost.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysPost 更新
func (r *sysPostRepo) UpdateSysPost(ctx context.Context, req *v1.SysPostUpdateReq) (*ent.SysPost, error) {
	return r.data.db.SysPost.UpdateOne(convert.SysPostUpdateReq2Data(req)).Save(ctx)
}

// GetSysPost 根据Id查询
func (r *sysPostRepo) GetSysPost(ctx context.Context, req *v1.SysPostReq) (*ent.SysPost, error) {
	return r.data.db.SysPost.Get(ctx, req.Id)
}

// PageSysPost 分页查询
func (r *sysPostRepo) PageSysPost(ctx context.Context, req *v1.SysPostPageReq) ([]*ent.SysPost, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.SysPost.
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
func (r *sysPostRepo) genCondition(req *v1.SysPostReq) []predicate.SysPost {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysPost, 0)
	if req.Id > 0 {
		list = append(list, syspost.ID(req.Id))
	}
	if str.IsBlank(req.PostName) {
		list = append(list, syspost.PostNameContains(req.PostName))
	}
	if str.IsBlank(req.PostCode) {
		list = append(list, syspost.PostCodeContains(req.PostCode))
	}
	if req.Sort > 0 {
		list = append(list, syspost.Sort(req.Sort))
	}
	if req.Status > 0 {
		list = append(list, syspost.Status(req.Status))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, syspost.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, syspost.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, syspost.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, syspost.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, syspost.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, syspost.TenantId(req.TenantId))
	}

	return list
}
