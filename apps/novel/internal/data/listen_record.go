package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/listenrecord/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/listenrecord"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/pagin"
	"time"
)

type listenRecordRepo struct {
	data *Data
	log  *log.Helper
}

// NewListenRecordRepo .
func NewListenRecordRepo(data *Data, logger log.Logger) biz.ListenRecordRepo {
	return &listenRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateListenRecord 创建
func (r *listenRecordRepo) CreateListenRecord(ctx context.Context, req *v1.ListenRecordCreateReq) (*ent.ListenRecord, error) {
	now := time.Now()
	return r.data.db.ListenRecord.Create().
		SetUserId(req.UserId).
		SetChapterId(req.ChapterId).
		SetNovelId(req.NovelId).
		SetListenTimes(req.ListenTimes).
		SetDuration(req.Duration.AsDuration()).
		SetAllDuration(req.AllDuration.AsDuration()).
		SetDayDuration(req.DayDuration.AsDuration()).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteListenRecord 删除
func (r *listenRecordRepo) DeleteListenRecord(ctx context.Context, req *v1.ListenRecordDeleteReq) error {
	return r.data.db.ListenRecord.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteListenRecord 批量删除
func (r *listenRecordRepo) BatchDeleteListenRecord(ctx context.Context, req *v1.ListenRecordBatchDeleteReq) (int, error) {
	return r.data.db.ListenRecord.Delete().Where(listenrecord.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateListenRecord 更新
func (r *listenRecordRepo) UpdateListenRecord(ctx context.Context, req *v1.ListenRecordUpdateReq) (*ent.ListenRecord, error) {
	return r.data.db.ListenRecord.UpdateOne(convert.ListenRecordUpdateReq2Data(req)).Save(ctx)
}

// GetListenRecord 根据Id查询
func (r *listenRecordRepo) GetListenRecord(ctx context.Context, req *v1.ListenRecordReq) (*ent.ListenRecord, error) {
	return r.data.db.ListenRecord.Get(ctx, req.Id)
}

// PageListenRecord 分页查询
func (r *listenRecordRepo) PageListenRecord(ctx context.Context, req *v1.ListenRecordPageReq) ([]*ent.ListenRecord, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.ListenRecord.
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
func (r *listenRecordRepo) genCondition(req *v1.ListenRecordReq) []predicate.ListenRecord {
	if req == nil {
		return nil
	}
	list := make([]predicate.ListenRecord, 0)
	if req.Id > 0 {
		list = append(list, listenrecord.ID(req.Id))
	}
	if req.UserId > 0 {
		list = append(list, listenrecord.UserId(req.UserId))
	}
	if req.ChapterId > 0 {
		list = append(list, listenrecord.ChapterId(req.ChapterId))
	}
	if req.NovelId > 0 {
		list = append(list, listenrecord.NovelId(req.NovelId))
	}
	if req.ListenTimes > 0 {
		list = append(list, listenrecord.ListenTimes(req.ListenTimes))
	}
	if req.Duration.AsDuration() > 0 {
		list = append(list, listenrecord.Duration(req.Duration.AsDuration()))
	}
	if req.AllDuration.AsDuration() > 0 {
		list = append(list, listenrecord.AllDuration(req.AllDuration.AsDuration()))
	}
	if req.DayDuration.AsDuration() > 0 {
		list = append(list, listenrecord.DayDuration(req.DayDuration.AsDuration()))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, listenrecord.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, listenrecord.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, listenrecord.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, listenrecord.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, listenrecord.TenantId(req.TenantId))
	}

	return list
}
