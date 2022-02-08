package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/useranalysisstatistics/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/useranalysisstatistics"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type userAnalysisStatisticsRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserAnalysisStatisticsRepo .
func NewUserAnalysisStatisticsRepo(data *Data, logger log.Logger) biz.UserAnalysisStatisticsRepo {
	return &userAnalysisStatisticsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUserAnalysisStatistics 创建
func (r *userAnalysisStatisticsRepo) CreateUserAnalysisStatistics(ctx context.Context, req *v1.UserAnalysisStatisticsCreateReq) (*ent.UserAnalysisStatistics, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.UserAnalysisStatistics.Create().
		SetStatisticsDate(req.StatisticsDate).
		SetAllUserNum(req.AllUserNum).
		SetAllPayment(req.AllPayment).
		SetAllPayUser(req.AllPayUser).
		SetAllOrderNum(req.AllOrderNum).
		SetDayUserNum(req.DayUserNum).
		SetDayPayment(req.DayPayment).
		SetDayOrderNum(req.DayOrderNum).
		SetDayPayUser(req.DayPayUser).
		SetDayRegPayment(req.DayRegPayment).
		SetDayRegUserNum(req.DayRegUserNum).
		SetDayRegOrderNum(req.DayRegOrderNum).
		SetOldRegPayment(req.OldRegPayment).
		SetOldRegUserNum(req.OldRegUserNum).
		SetOldRegOrderNum(req.OldRegOrderNum).
		SetPayRate(req.PayRate).
		SetArpu(req.Arpu).
		SetDayRegArpu(req.DayRegArpu).
		SetDayArpu(req.DayArpu).
		SetDayOldArpu(req.DayOldArpu).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteUserAnalysisStatistics 删除
func (r *userAnalysisStatisticsRepo) DeleteUserAnalysisStatistics(ctx context.Context, req *v1.UserAnalysisStatisticsDeleteReq) error {
	return r.data.db.UserAnalysisStatistics.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteUserAnalysisStatistics 批量删除
func (r *userAnalysisStatisticsRepo) BatchDeleteUserAnalysisStatistics(ctx context.Context, req *v1.UserAnalysisStatisticsBatchDeleteReq) (int, error) {
	return r.data.db.UserAnalysisStatistics.Delete().Where(useranalysisstatistics.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateUserAnalysisStatistics 更新
func (r *userAnalysisStatisticsRepo) UpdateUserAnalysisStatistics(ctx context.Context, req *v1.UserAnalysisStatisticsUpdateReq) (*ent.UserAnalysisStatistics, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.UserAnalysisStatistics.UpdateOneID(req.Id).
		SetStatisticsDate(req.StatisticsDate).
		SetAllUserNum(req.AllUserNum).
		SetAllPayment(req.AllPayment).
		SetAllPayUser(req.AllPayUser).
		SetAllOrderNum(req.AllOrderNum).
		SetDayUserNum(req.DayUserNum).
		SetDayPayment(req.DayPayment).
		SetDayOrderNum(req.DayOrderNum).
		SetDayPayUser(req.DayPayUser).
		SetDayRegPayment(req.DayRegPayment).
		SetDayRegUserNum(req.DayRegUserNum).
		SetDayRegOrderNum(req.DayRegOrderNum).
		SetOldRegPayment(req.OldRegPayment).
		SetOldRegUserNum(req.OldRegUserNum).
		SetOldRegOrderNum(req.OldRegOrderNum).
		SetPayRate(req.PayRate).
		SetArpu(req.Arpu).
		SetDayRegArpu(req.DayRegArpu).
		SetDayArpu(req.DayArpu).
		SetDayOldArpu(req.DayOldArpu).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetUserAnalysisStatistics 根据Id查询
func (r *userAnalysisStatisticsRepo) GetUserAnalysisStatistics(ctx context.Context, req *v1.UserAnalysisStatisticsReq) (*ent.UserAnalysisStatistics, error) {
	return r.data.db.UserAnalysisStatistics.Get(ctx, req.Id)
}

// PageUserAnalysisStatistics 分页查询
func (r *userAnalysisStatisticsRepo) PageUserAnalysisStatistics(ctx context.Context, req *v1.UserAnalysisStatisticsPageReq) ([]*ent.UserAnalysisStatistics, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.UserAnalysisStatistics.
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
func (r *userAnalysisStatisticsRepo) genCondition(req *v1.UserAnalysisStatisticsReq) []predicate.UserAnalysisStatistics {
	if req == nil {
		return nil
	}
	list := make([]predicate.UserAnalysisStatistics, 0)
	if req.Id > 0 {
		list = append(list, useranalysisstatistics.ID(req.Id))
	}
	if str.IsNotBlank(req.StatisticsDate) {
		list = append(list, useranalysisstatistics.StatisticsDateContains(req.StatisticsDate))
	}
	if req.AllUserNum > 0 {
		list = append(list, useranalysisstatistics.AllUserNum(req.AllUserNum))
	}
	if req.AllPayment > 0 {
		list = append(list, useranalysisstatistics.AllPayment(req.AllPayment))
	}
	if req.AllPayUser > 0 {
		list = append(list, useranalysisstatistics.AllPayUser(req.AllPayUser))
	}
	if req.AllOrderNum > 0 {
		list = append(list, useranalysisstatistics.AllOrderNum(req.AllOrderNum))
	}
	if req.DayUserNum > 0 {
		list = append(list, useranalysisstatistics.DayUserNum(req.DayUserNum))
	}
	if req.DayPayment > 0 {
		list = append(list, useranalysisstatistics.DayPayment(req.DayPayment))
	}
	if req.DayOrderNum > 0 {
		list = append(list, useranalysisstatistics.DayOrderNum(req.DayOrderNum))
	}
	if req.DayPayUser > 0 {
		list = append(list, useranalysisstatistics.DayPayUser(req.DayPayUser))
	}
	if req.DayRegPayment > 0 {
		list = append(list, useranalysisstatistics.DayRegPayment(req.DayRegPayment))
	}
	if req.DayRegUserNum > 0 {
		list = append(list, useranalysisstatistics.DayRegUserNum(req.DayRegUserNum))
	}
	if req.DayRegOrderNum > 0 {
		list = append(list, useranalysisstatistics.DayRegOrderNum(req.DayRegOrderNum))
	}
	if req.OldRegPayment > 0 {
		list = append(list, useranalysisstatistics.OldRegPayment(req.OldRegPayment))
	}
	if req.OldRegUserNum > 0 {
		list = append(list, useranalysisstatistics.OldRegUserNum(req.OldRegUserNum))
	}
	if req.OldRegOrderNum > 0 {
		list = append(list, useranalysisstatistics.OldRegOrderNum(req.OldRegOrderNum))
	}
	if req.PayRate > 0 {
		list = append(list, useranalysisstatistics.PayRate(req.PayRate))
	}
	if req.Arpu > 0 {
		list = append(list, useranalysisstatistics.Arpu(req.Arpu))
	}
	if req.DayRegArpu > 0 {
		list = append(list, useranalysisstatistics.DayRegArpu(req.DayRegArpu))
	}
	if req.DayArpu > 0 {
		list = append(list, useranalysisstatistics.DayArpu(req.DayArpu))
	}
	if req.DayOldArpu > 0 {
		list = append(list, useranalysisstatistics.DayOldArpu(req.DayOldArpu))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, useranalysisstatistics.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, useranalysisstatistics.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, useranalysisstatistics.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, useranalysisstatistics.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, useranalysisstatistics.TenantId(req.TenantId))
	}

	return list
}
