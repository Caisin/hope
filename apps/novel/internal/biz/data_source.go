package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/datasource/v1"
	"hope/apps/novel/internal/data/ent"
)

type DataSourceRepo interface {
	CreateDataSource(context.Context, *v1.DataSourceCreateReq) (*ent.DataSource, error)
	DeleteDataSource(context.Context, *v1.DataSourceDeleteReq) error
	BatchDeleteDataSource(context.Context, *v1.DataSourceBatchDeleteReq) (int, error)
	UpdateDataSource(context.Context, *v1.DataSourceUpdateReq) (*ent.DataSource, error)
	GetDataSource(context.Context, *v1.DataSourceReq) (*ent.DataSource, error)
	PageDataSource(context.Context, *v1.DataSourcePageReq) ([]*ent.DataSource, error)
}

type DataSourceUseCase struct {
	repo DataSourceRepo
	log  *log.Helper
}

func NewDataSourceUseCase(repo DataSourceRepo, logger log.Logger) *DataSourceUseCase {
	return &DataSourceUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DataSourceUseCase) Create(ctx context.Context, req *v1.DataSourceCreateReq) (*ent.DataSource, error) {
	return uc.repo.CreateDataSource(ctx, req)
}
func (uc *DataSourceUseCase) Delete(ctx context.Context, req *v1.DataSourceDeleteReq) error {
	return uc.repo.DeleteDataSource(ctx, req)
}
func (uc *DataSourceUseCase) BatchDelete(ctx context.Context, req *v1.DataSourceBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteDataSource(ctx, req)
}
func (uc *DataSourceUseCase) Update(ctx context.Context, req *v1.DataSourceUpdateReq) (*ent.DataSource, error) {
	return uc.repo.UpdateDataSource(ctx, req)
}
func (uc *DataSourceUseCase) Get(ctx context.Context, req *v1.DataSourceReq) (*ent.DataSource, error) {
	return uc.repo.GetDataSource(ctx, req)
}
func (uc *DataSourceUseCase) Page(ctx context.Context, req *v1.DataSourcePageReq) ([]*ent.DataSource, error) {
	return uc.repo.PageDataSource(ctx, req)
}
