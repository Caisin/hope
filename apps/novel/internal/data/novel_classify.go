package data
import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/novelclassify/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/novelclassify"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"
	"hope/pkg/pagin"
	"time"
)

type novelClassifyRepo struct {
	data *Data
	log  *log.Helper
}

// NewNovelClassifyRepo .
func NewNovelClassifyRepo(data *Data, logger log.Logger) biz.NovelClassifyRepo {
	return &novelClassifyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateNovelClassify 创建
func (r *novelClassifyRepo) CreateNovelClassify(ctx context.Context, req *v1.NovelClassifyCreateReq) (*ent.NovelClassify, error) {
	now := time.Now()
	return r.data.db.NovelClassify.Create().
    SetPid(req.Pid).
    SetClassifyName(req.ClassifyName).
    SetStatus(req.Status).
    SetOrderNum(req.OrderNum).
    SetRemark(req.Remark).
	SetCreatedAt(now).
	SetUpdatedAt(now).
	Save(ctx)

}

// DeleteNovelClassify 删除
func (r *novelClassifyRepo) DeleteNovelClassify(ctx context.Context, req *v1.NovelClassifyDeleteReq) error {
	return r.data.db.NovelClassify.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteNovelClassify 批量删除
func (r *novelClassifyRepo) BatchDeleteNovelClassify(ctx context.Context, req *v1.NovelClassifyBatchDeleteReq) (int, error) {
	return r.data.db.NovelClassify.Delete().Where(novelclassify.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateNovelClassify 更新
func (r *novelClassifyRepo) UpdateNovelClassify(ctx context.Context, req *v1.NovelClassifyUpdateReq) (*ent.NovelClassify, error) {
	return r.data.db.NovelClassify.UpdateOne(convert.NovelClassifyUpdateReq2Data(req)).Save(ctx)
}

// GetNovelClassify 根据Id查询
func (r *novelClassifyRepo) GetNovelClassify(ctx context.Context, req *v1.NovelClassifyReq) (*ent.NovelClassify, error) {
	return r.data.db.NovelClassify.Get(ctx, req.Id)
}

// PageNovelClassify 分页查询
func (r *novelClassifyRepo) PageNovelClassify(ctx context.Context, req *v1.NovelClassifyPageReq) ([]*ent.NovelClassify, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin=&pagin.Pagination{}
	}
	query := r.data.db.NovelClassify.
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
func (r *novelClassifyRepo) genCondition(req *v1.NovelClassifyReq) []predicate.NovelClassify {
	if req == nil {
		return nil
	}
	list := make([]predicate.NovelClassify, 0)
	if req.Id > 0 {
		list = append(list, novelclassify.ID(req.Id))
	}
	if req.Pid > 0 {
		list = append(list, novelclassify.Pid(req.Pid))
	}
	if str.IsBlank(req.ClassifyName) {
		list = append(list, novelclassify.ClassifyNameContains(req.ClassifyName))
	}
	if req.Status > 0 {
		list = append(list, novelclassify.Status(req.Status))
	}
	if req.OrderNum > 0 {
		list = append(list, novelclassify.OrderNum(req.OrderNum))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, novelclassify.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, novelclassify.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, novelclassify.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, novelclassify.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, novelclassify.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, novelclassify.TenantId(req.TenantId))
	}
	
	return list
}
