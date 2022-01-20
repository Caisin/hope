package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/agreementlog/v1"
	"hope/apps/novel/internal/data/ent"
)

type AgreementLogRepo interface {
	CreateAgreementLog(context.Context, *v1.AgreementLogCreateReq) (*ent.AgreementLog, error)
	DeleteAgreementLog(context.Context, *v1.AgreementLogDeleteReq) error
	BatchDeleteAgreementLog(context.Context, *v1.AgreementLogBatchDeleteReq) (int, error)
	UpdateAgreementLog(context.Context, *v1.AgreementLogUpdateReq) (*ent.AgreementLog, error)
	GetAgreementLog(context.Context, *v1.AgreementLogReq) (*ent.AgreementLog, error)
	PageAgreementLog(context.Context, *v1.AgreementLogPageReq) ([]*ent.AgreementLog, error)
}

type AgreementLogUseCase struct {
	repo AgreementLogRepo
	log  *log.Helper
}

func NewAgreementLogUseCase(repo AgreementLogRepo, logger log.Logger) *AgreementLogUseCase {
	return &AgreementLogUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AgreementLogUseCase) Create(ctx context.Context, req *v1.AgreementLogCreateReq) (*ent.AgreementLog, error) {
	return uc.repo.CreateAgreementLog(ctx, req)
}
func (uc *AgreementLogUseCase) Delete(ctx context.Context, req *v1.AgreementLogDeleteReq) error {
	return uc.repo.DeleteAgreementLog(ctx, req)
}
func (uc *AgreementLogUseCase) BatchDelete(ctx context.Context, req *v1.AgreementLogBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteAgreementLog(ctx, req)
}
func (uc *AgreementLogUseCase) Update(ctx context.Context, req *v1.AgreementLogUpdateReq) (*ent.AgreementLog, error) {
	return uc.repo.UpdateAgreementLog(ctx, req)
}
func (uc *AgreementLogUseCase) Get(ctx context.Context, req *v1.AgreementLogReq) (*ent.AgreementLog, error) {
	return uc.repo.GetAgreementLog(ctx, req)
}
func (uc *AgreementLogUseCase) Page(ctx context.Context, req *v1.AgreementLogPageReq) ([]*ent.AgreementLog, error) {
	return uc.repo.PageAgreementLog(ctx, req)
}
