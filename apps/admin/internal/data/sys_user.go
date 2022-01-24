package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/admin/sysuser/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/convert"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysuser"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type sysUserRepo struct {
	data *Data
	log  *log.Helper
}

// NewSysUserRepo .
func NewSysUserRepo(data *Data, logger log.Logger) biz.SysUserRepo {
	return &sysUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSysUser 创建
func (r *sysUserRepo) CreateSysUser(ctx context.Context, req *v1.SysUserCreateReq) (*ent.SysUser, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	encrypt, err := auth.Encrypt(req.Password)
	if err != nil {
		return nil, err
	}
	return r.data.db.SysUser.Create().
		SetUsername(req.Username).
		SetNickName(req.NickName).
		SetPhone(req.Phone).
		SetDeptId(req.DeptId).
		SetPostId(req.PostId).
		SetRoleId(req.RoleId).
		SetAvatar(req.Avatar).
		SetSex(req.Sex).
		SetEmail(req.Email).
		SetRemark(req.Remark).
		SetStatus(req.Status).
		SetExtInfo(req.ExtInfo).
		SetHomePath(req.HomePath).
		SetDesc(req.Desc).
		SetPassword(encrypt).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteSysUser 删除
func (r *sysUserRepo) DeleteSysUser(ctx context.Context, req *v1.SysUserDeleteReq) error {
	return r.data.db.SysUser.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSysUser 批量删除
func (r *sysUserRepo) BatchDeleteSysUser(ctx context.Context, req *v1.SysUserBatchDeleteReq) (int, error) {
	return r.data.db.SysUser.Delete().Where(sysuser.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSysUser 更新
func (r *sysUserRepo) UpdateSysUser(ctx context.Context, req *v1.SysUserUpdateReq) (*ent.SysUser, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	data := convert.SysUserUpdateReq2Data(req)
	data.UpdateBy = claims.UserId
	return r.data.db.SysUser.UpdateOne(data).Save(ctx)
}

// GetSysUser 根据Id查询
func (r *sysUserRepo) GetSysUser(ctx context.Context, req *v1.SysUserReq) (*ent.SysUser, error) {
	return r.data.db.SysUser.Get(ctx, req.Id)
}

// PageSysUser 分页查询
func (r *sysUserRepo) PageSysUser(ctx context.Context, req *v1.SysUserPageReq) ([]*ent.SysUser, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.SysUser.
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
func (r *sysUserRepo) genCondition(req *v1.SysUserReq) []predicate.SysUser {
	if req == nil {
		return nil
	}
	list := make([]predicate.SysUser, 0)
	if req.Id > 0 {
		list = append(list, sysuser.ID(req.Id))
	}
	if str.IsBlank(req.Username) {
		list = append(list, sysuser.UsernameContains(req.Username))
	}
	if str.IsBlank(req.NickName) {
		list = append(list, sysuser.NickNameContains(req.NickName))
	}
	if str.IsBlank(req.Phone) {
		list = append(list, sysuser.PhoneContains(req.Phone))
	}
	if req.DeptId > 0 {
		list = append(list, sysuser.DeptId(req.DeptId))
	}
	if req.PostId > 0 {
		list = append(list, sysuser.PostId(req.PostId))
	}
	if req.RoleId > 0 {
		list = append(list, sysuser.RoleId(req.RoleId))
	}
	if str.IsBlank(req.Avatar) {
		list = append(list, sysuser.AvatarContains(req.Avatar))
	}
	if req.Sex > 0 {
		list = append(list, sysuser.Sex(req.Sex))
	}
	if str.IsBlank(req.Email) {
		list = append(list, sysuser.EmailContains(req.Email))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, sysuser.RemarkContains(req.Remark))
	}
	if str.IsBlank(req.Status) {
		list = append(list, sysuser.StatusContains(req.Status))
	}
	if str.IsBlank(req.ExtInfo) {
		list = append(list, sysuser.ExtInfoContains(req.ExtInfo))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, sysuser.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, sysuser.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, sysuser.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, sysuser.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, sysuser.TenantId(req.TenantId))
	}

	return list
}
