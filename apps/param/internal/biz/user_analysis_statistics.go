package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/useranalysisstatistics/v1"
	"hope/apps/param/internal/data/ent"
)

type UserAnalysisStatisticsRepo interface {
	CreateUserAnalysisStatistics(context.Context, *v1.UserAnalysisStatisticsCreateReq) (*ent.UserAnalysisStatistics, error)
	DeleteUserAnalysisStatistics(context.Context, *v1.UserAnalysisStatisticsDeleteReq) error
	BatchDeleteUserAnalysisStatistics(context.Context, *v1.UserAnalysisStatisticsBatchDeleteReq) (int, error)
	UpdateUserAnalysisStatistics(context.Context, *v1.UserAnalysisStatisticsUpdateReq) (*ent.UserAnalysisStatistics, error)
	GetUserAnalysisStatistics(context.Context, *v1.UserAnalysisStatisticsReq) (*ent.UserAnalysisStatistics, error)
	PageUserAnalysisStatistics(context.Context, *v1.UserAnalysisStatisticsPageReq) ([]*ent.UserAnalysisStatistics, error)
}

type UserAnalysisStatisticsUseCase struct {
	repo UserAnalysisStatisticsRepo
	log  *log.Helper
}

func NewUserAnalysisStatisticsUseCase(repo UserAnalysisStatisticsRepo, logger log.Logger) *UserAnalysisStatisticsUseCase {
	return &UserAnalysisStatisticsUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserAnalysisStatisticsUseCase) Create(ctx context.Context, req *v1.UserAnalysisStatisticsCreateReq) (*ent.UserAnalysisStatistics, error) {
	return uc.repo.CreateUserAnalysisStatistics(ctx, req)
}
func (uc *UserAnalysisStatisticsUseCase) Delete(ctx context.Context, req *v1.UserAnalysisStatisticsDeleteReq) error {
	return uc.repo.DeleteUserAnalysisStatistics(ctx, req)
}
func (uc *UserAnalysisStatisticsUseCase) BatchDelete(ctx context.Context, req *v1.UserAnalysisStatisticsBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteUserAnalysisStatistics(ctx, req)
}
func (uc *UserAnalysisStatisticsUseCase) Update(ctx context.Context, req *v1.UserAnalysisStatisticsUpdateReq) (*ent.UserAnalysisStatistics, error) {
	return uc.repo.UpdateUserAnalysisStatistics(ctx, req)
}
func (uc *UserAnalysisStatisticsUseCase) Get(ctx context.Context, req *v1.UserAnalysisStatisticsReq) (*ent.UserAnalysisStatistics, error) {
	return uc.repo.GetUserAnalysisStatistics(ctx, req)
}
func (uc *UserAnalysisStatisticsUseCase) Page(ctx context.Context, req *v1.UserAnalysisStatisticsPageReq) ([]*ent.UserAnalysisStatistics, error) {
	return uc.repo.PageUserAnalysisStatistics(ctx, req)
}
