package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/adchannel/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/adchannel"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type adChannelRepo struct {
	data *Data
	log  *log.Helper
}

// NewAdChannelRepo .
func NewAdChannelRepo(data *Data, logger log.Logger) biz.AdChannelRepo {
	return &adChannelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateAdChannel 创建
func (r *adChannelRepo) CreateAdChannel(ctx context.Context, req *v1.AdChannelCreateReq) (*ent.AdChannel, error) {
	now := time.Now()
	return r.data.db.AdChannel.Create().
		SetChannelName(req.ChannelName).
		SetNovelId(req.NovelId).
		SetReg(req.Reg).
		SetPay(req.Pay).
		SetNovelName(req.NovelName).
		SetChapterId(req.ChapterId).
		SetChapterNum(req.ChapterNum).
		SetAdType(req.AdType).
		SetImg(req.Img).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteAdChannel 删除
func (r *adChannelRepo) DeleteAdChannel(ctx context.Context, req *v1.AdChannelDeleteReq) error {
	return r.data.db.AdChannel.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteAdChannel 批量删除
func (r *adChannelRepo) BatchDeleteAdChannel(ctx context.Context, req *v1.AdChannelBatchDeleteReq) (int, error) {
	return r.data.db.AdChannel.Delete().Where(adchannel.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateAdChannel 更新
func (r *adChannelRepo) UpdateAdChannel(ctx context.Context, req *v1.AdChannelUpdateReq) (*ent.AdChannel, error) {
	return r.data.db.AdChannel.UpdateOne(convert.AdChannelUpdateReq2Data(req)).Save(ctx)
}

// GetAdChannel 根据Id查询
func (r *adChannelRepo) GetAdChannel(ctx context.Context, req *v1.AdChannelReq) (*ent.AdChannel, error) {
	return r.data.db.AdChannel.Get(ctx, req.Id)
}

// PageAdChannel 分页查询
func (r *adChannelRepo) PageAdChannel(ctx context.Context, req *v1.AdChannelPageReq) ([]*ent.AdChannel, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.AdChannel.
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
func (r *adChannelRepo) genCondition(req *v1.AdChannelReq) []predicate.AdChannel {
	if req == nil {
		return nil
	}
	list := make([]predicate.AdChannel, 0)
	if req.Id > 0 {
		list = append(list, adchannel.ID(req.Id))
	}
	if str.IsBlank(req.ChannelName) {
		list = append(list, adchannel.ChannelNameContains(req.ChannelName))
	}
	if req.NovelId > 0 {
		list = append(list, adchannel.NovelId(req.NovelId))
	}
	if req.Reg > 0 {
		list = append(list, adchannel.Reg(req.Reg))
	}
	if req.Pay > 0 {
		list = append(list, adchannel.Pay(req.Pay))
	}
	if str.IsBlank(req.NovelName) {
		list = append(list, adchannel.NovelNameContains(req.NovelName))
	}
	if req.ChapterId > 0 {
		list = append(list, adchannel.ChapterId(req.ChapterId))
	}
	if req.ChapterNum > 0 {
		list = append(list, adchannel.ChapterNum(req.ChapterNum))
	}
	if str.IsBlank(req.AdType) {
		list = append(list, adchannel.AdTypeContains(req.AdType))
	}
	if str.IsBlank(req.Img) {
		list = append(list, adchannel.ImgContains(req.Img))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, adchannel.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, adchannel.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, adchannel.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, adchannel.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, adchannel.TenantId(req.TenantId))
	}

	return list
}
