package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novelautobuy/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novelautobuy"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/pagin"
	"time"
)

type novelAutoBuyRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelAutoBuyRepo .
func NewNovelAutoBuyRepo(data *Data, logger log.Logger) biz.NovelAutoBuyRepo {
	return &novelAutoBuyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelAutoBuy 创建
func (r *novelAutoBuyRepo) CreateNovelAutoBuy(ctx context.Context, req *v1.NovelAutoBuyCreateReq) (*ent.NovelAutoBuy, error) {
	now := time.Now()
	return r.data.db.NovelAutoBuy.Create().
		SetUserId(req.UserId).
		SetNovelId(req.NovelId).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteNovelAutoBuy 删除
func (r *novelAutoBuyRepo) DeleteNovelAutoBuy(ctx context.Context, req *v1.NovelAutoBuyDeleteReq) error {
	return r.data.db.NovelAutoBuy.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelAutoBuy 批量删除
func (r *novelAutoBuyRepo) BatchDeleteNovelAutoBuy(ctx context.Context, req *v1.NovelAutoBuyBatchDeleteReq) (int, error) {
	return r.data.db.NovelAutoBuy.Delete().Where(novelautobuy.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelAutoBuy 更新
func (r *novelAutoBuyRepo) UpdateNovelAutoBuy(ctx context.Context, req *v1.NovelAutoBuyUpdateReq) (*ent.NovelAutoBuy, error) {
	return r.data.db.NovelAutoBuy.UpdateOne(convert.NovelAutoBuyUpdateReq2Data(req)).Save(ctx)
}

// GetNovelAutoBuy 根据Id查询
func (r *novelAutoBuyRepo) GetNovelAutoBuy(ctx context.Context, req *v1.NovelAutoBuyReq) (*ent.NovelAutoBuy, error) {
	return r.data.db.NovelAutoBuy.Get(ctx, req.Id)
}

// PageNovelAutoBuy 分页查询
func (r *novelAutoBuyRepo) PageNovelAutoBuy(ctx context.Context, req *v1.NovelAutoBuyPageReq) ([]*ent.NovelAutoBuy, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.NovelAutoBuy.
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
func (r *novelAutoBuyRepo) genCondition(req *v1.NovelAutoBuyReq) []predicate.NovelAutoBuy {
	if req == nil {
		return nil
	}
	list := make([]predicate.NovelAutoBuy, 0)
	if req.Id > 0 {
		list = append(list, novelautobuy.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, novelautobuy.UserId(req.UserId))
	}
	if req.NovelId > 0 {
		list = append(list, novelautobuy.NovelId(req.NovelId))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novelautobuy.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novelautobuy.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novelautobuy.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novelautobuy.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novelautobuy.TenantId(req.TenantId))
	}

	return list
}
