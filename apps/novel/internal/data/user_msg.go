package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/usermsg/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/usermsg"
	"hope/pkg/pagin"
	"time"
)

type userMsgRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserMsgRepo .
func NewUserMsgRepo(data *Data, logger log.Logger) biz.UserMsgRepo {
	return &userMsgRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUserMsg 创建
func (r *userMsgRepo) CreateUserMsg(ctx context.Context, req *v1.UserMsgCreateReq) (*ent.UserMsg, error) {
	now := time.Now()
	return r.data.db.UserMsg.Create().
		SetUserId(req.UserId).
		SetMsgId(req.MsgId).
		SetIsRead(req.IsRead).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteUserMsg 删除
func (r *userMsgRepo) DeleteUserMsg(ctx context.Context, req *v1.UserMsgDeleteReq) error {
	return r.data.db.UserMsg.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteUserMsg 批量删除
func (r *userMsgRepo) BatchDeleteUserMsg(ctx context.Context, req *v1.UserMsgBatchDeleteReq) (int, error) {
	return r.data.db.UserMsg.Delete().Where(usermsg.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateUserMsg 更新
func (r *userMsgRepo) UpdateUserMsg(ctx context.Context, req *v1.UserMsgUpdateReq) (*ent.UserMsg, error) {
	return r.data.db.UserMsg.UpdateOne(convert.UserMsgUpdateReq2Data(req)).Save(ctx)
}

// GetUserMsg 根据Id查询
func (r *userMsgRepo) GetUserMsg(ctx context.Context, req *v1.UserMsgReq) (*ent.UserMsg, error) {
	return r.data.db.UserMsg.Get(ctx, req.Id)
}

// PageUserMsg 分页查询
func (r *userMsgRepo) PageUserMsg(ctx context.Context, req *v1.UserMsgPageReq) ([]*ent.UserMsg, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{}
	}
	query := r.data.db.UserMsg.
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
	query.Limit(int(p.GetPage())).
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
func (r *userMsgRepo) genCondition(req *v1.UserMsgReq) []predicate.UserMsg {
	if req == nil {
		return nil
	}
	list := make([]predicate.UserMsg, 0)
	if req.Id > 0 {
		list = append(list, usermsg.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, usermsg.UserId(req.UserId))
	}
	if req.MsgId > 0 {
		list = append(list, usermsg.MsgId(req.MsgId))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, usermsg.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, usermsg.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, usermsg.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, usermsg.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, usermsg.TenantId(req.TenantId))
	}

	return list
}
