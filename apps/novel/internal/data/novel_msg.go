package data
import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novelmsg/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novelmsg"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"
	"hope/pkg/pagin"
	"time"
)

type novelMsgRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelMsgRepo .
func NewNovelMsgRepo(data *Data, logger log.Logger) biz.NovelMsgRepo {
	return &novelMsgRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelMsg 创建
func (r *novelMsgRepo) CreateNovelMsg(ctx context.Context, req *v1.NovelMsgCreateReq) (*ent.NovelMsg, error) {
	now := time.Now()
	return r.data.db.NovelMsg.Create().
    SetTitle(req.Title).
    SetMsg(req.Msg).
    SetMsgType(req.MsgType).
    SetStatus(req.Status).
	SetCreatedAt(now).
	SetUpdatedAt(now).
	Save(ctx)

}

// DeleteNovelMsg 删除
func (r *novelMsgRepo) DeleteNovelMsg(ctx context.Context, req *v1.NovelMsgDeleteReq) error {
	return r.data.db.NovelMsg.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelMsg 批量删除
func (r *novelMsgRepo) BatchDeleteNovelMsg(ctx context.Context, req *v1.NovelMsgBatchDeleteReq) (int, error) {
	return r.data.db.NovelMsg.Delete().Where(novelmsg.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelMsg 更新
func (r *novelMsgRepo) UpdateNovelMsg(ctx context.Context, req *v1.NovelMsgUpdateReq) (*ent.NovelMsg, error) {
	return r.data.db.NovelMsg.UpdateOne(convert.NovelMsgUpdateReq2Data(req)).Save(ctx)
}

// GetNovelMsg 根据Id查询
func (r *novelMsgRepo) GetNovelMsg(ctx context.Context, req *v1.NovelMsgReq) (*ent.NovelMsg, error) {
	return r.data.db.NovelMsg.Get(ctx, req.Id)
}

// PageNovelMsg 分页查询
func (r *novelMsgRepo) PageNovelMsg(ctx context.Context, req *v1.NovelMsgPageReq) ([]*ent.NovelMsg, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin=&pagin.Pagination{}
	}
	query := r.data.db.NovelMsg.
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
func (r *novelMsgRepo) genCondition(req *v1.NovelMsgReq) []predicate.NovelMsg {
	if req == nil {
		return nil
	}
	list := make([]predicate.NovelMsg, 0)
	if req.Id > 0 {
		list = append(list, novelmsg.ID(req.Id))
	}
	if str.IsBlank(req.Title) {
		list = append(list, novelmsg.TitleContains(req.Title))
	}
	if str.IsBlank(req.Msg) {
		list = append(list, novelmsg.MsgContains(req.Msg))
	}
	if str.IsBlank(req.MsgType) {
		list = append(list, novelmsg.MsgTypeContains(req.MsgType))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novelmsg.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novelmsg.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novelmsg.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novelmsg.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novelmsg.TenantId(req.TenantId))
	}
	
	return list
}
