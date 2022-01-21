package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/agreementlog/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/agreementlog"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type agreementLogRepo struct {
	data *Data
	log  *log.Helper
}

// NewAgreementLogRepo .
func NewAgreementLogRepo(data *Data, logger log.Logger) biz.AgreementLogRepo {
	return &agreementLogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateAgreementLog 创建
func (r *agreementLogRepo) CreateAgreementLog(ctx context.Context, req *v1.AgreementLogCreateReq) (*ent.AgreementLog, error) {
	now := time.Now()
	return r.data.db.AgreementLog.Create().
		SetOuterAgreementNo(req.OuterAgreementNo).
		SetOrderId(req.OrderId).
		SetUserId(req.UserId).
		SetChId(req.ChId).
		SetUserName(req.UserName).
		SetPaymentName(req.PaymentName).
		SetPaymentId(req.PaymentId).
		SetState(req.State).
		SetPayment(req.Payment).
		SetAgreementType(agreementlog.AgreementType(req.AgreementType)).
		SetVipType(req.VipType).
		SetTimes(req.Times).
		SetCycleDays(req.CycleDays).
		SetNextExecTime(req.NextExecTime.AsTime()).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteAgreementLog 删除
func (r *agreementLogRepo) DeleteAgreementLog(ctx context.Context, req *v1.AgreementLogDeleteReq) error {
	return r.data.db.AgreementLog.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteAgreementLog 批量删除
func (r *agreementLogRepo) BatchDeleteAgreementLog(ctx context.Context, req *v1.AgreementLogBatchDeleteReq) (int, error) {
	return r.data.db.AgreementLog.Delete().Where(agreementlog.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateAgreementLog 更新
func (r *agreementLogRepo) UpdateAgreementLog(ctx context.Context, req *v1.AgreementLogUpdateReq) (*ent.AgreementLog, error) {
	return r.data.db.AgreementLog.UpdateOne(convert.AgreementLogUpdateReq2Data(req)).Save(ctx)
}

// GetAgreementLog 根据Id查询
func (r *agreementLogRepo) GetAgreementLog(ctx context.Context, req *v1.AgreementLogReq) (*ent.AgreementLog, error) {
	return r.data.db.AgreementLog.Get(ctx, req.Id)
}

// PageAgreementLog 分页查询
func (r *agreementLogRepo) PageAgreementLog(ctx context.Context, req *v1.AgreementLogPageReq) ([]*ent.AgreementLog, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.AgreementLog.
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
func (r *agreementLogRepo) genCondition(req *v1.AgreementLogReq) []predicate.AgreementLog {
	if req == nil {
		return nil
	}
	list := make([]predicate.AgreementLog, 0)
	if req.Id > 0 {
		list = append(list, agreementlog.ID(req.Id))
	}
	if str.IsBlank(req.OuterAgreementNo) {
		list = append(list, agreementlog.OuterAgreementNoContains(req.OuterAgreementNo))
	}
	if str.IsBlank(req.OrderId) {
		list = append(list, agreementlog.OrderIdContains(req.OrderId))
	}
	if req.UserId > 0 {
		list = append(list, agreementlog.UserId(req.UserId))
	}
	if req.ChId > 0 {
		list = append(list, agreementlog.ChId(req.ChId))
	}
	if str.IsBlank(req.UserName) {
		list = append(list, agreementlog.UserNameContains(req.UserName))
	}
	if str.IsBlank(req.PaymentName) {
		list = append(list, agreementlog.PaymentNameContains(req.PaymentName))
	}
	if req.PaymentId > 0 {
		list = append(list, agreementlog.PaymentId(req.PaymentId))
	}
	if req.State > 0 {
		list = append(list, agreementlog.State(req.State))
	}
	if req.Payment > 0 {
		list = append(list, agreementlog.Payment(req.Payment))
	}
	agreementType := agreementlog.AgreementType(req.AgreementType)
	if agreementlog.AgreementTypeValidator(agreementType) == nil {
		list = append(list, agreementlog.AgreementTypeEQ(agreementType))
	}
	if req.VipType > 0 {
		list = append(list, agreementlog.VipType(req.VipType))
	}
	if req.Times > 0 {
		list = append(list, agreementlog.Times(req.Times))
	}
	if req.CycleDays > 0 {
		list = append(list, agreementlog.CycleDays(req.CycleDays))
	}
	if req.NextExecTime.IsValid() && !req.NextExecTime.AsTime().IsZero() {
		list = append(list, agreementlog.NextExecTimeGTE(req.NextExecTime.AsTime()))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, agreementlog.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, agreementlog.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, agreementlog.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, agreementlog.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, agreementlog.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, agreementlog.TenantId(req.TenantId))
	}

	return list
}
