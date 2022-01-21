package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/datasource/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/convert"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/datasource"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type dataSourceRepo struct {
	data *Data
	log  *log.Helper
}

// NewDataSourceRepo .
func NewDataSourceRepo(data *Data, logger log.Logger) biz.DataSourceRepo {
	return &dataSourceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateDataSource 创建
func (r *dataSourceRepo) CreateDataSource(ctx context.Context, req *v1.DataSourceCreateReq) (*ent.DataSource, error) {
	now := time.Now()
	return r.data.db.DataSource.Create().
		SetDbName(req.DbName).
		SetHost(req.Host).
		SetPort(req.Port).
		SetDatabase(req.Database).
		SetUserName(req.UserName).
		SetPwd(req.Pwd).
		SetStatus(req.Status).
		SetDbType(datasource.DbType(req.DbType)).
		SetConnMaxIdleTime(req.ConnMaxIdleTime).
		SetConnMaxLifeTime(req.ConnMaxLifeTime).
		SetMaxIdleConns(req.MaxIdleConns).
		SetMaxOpenConns(req.MaxOpenConns).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

}

// DeleteDataSource 删除
func (r *dataSourceRepo) DeleteDataSource(ctx context.Context, req *v1.DataSourceDeleteReq) error {
	return r.data.db.DataSource.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteDataSource 批量删除
func (r *dataSourceRepo) BatchDeleteDataSource(ctx context.Context, req *v1.DataSourceBatchDeleteReq) (int, error) {
	return r.data.db.DataSource.Delete().Where(datasource.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateDataSource 更新
func (r *dataSourceRepo) UpdateDataSource(ctx context.Context, req *v1.DataSourceUpdateReq) (*ent.DataSource, error) {
	return r.data.db.DataSource.UpdateOne(convert.DataSourceUpdateReq2Data(req)).Save(ctx)
}

// GetDataSource 根据Id查询
func (r *dataSourceRepo) GetDataSource(ctx context.Context, req *v1.DataSourceReq) (*ent.DataSource, error) {
	return r.data.db.DataSource.Get(ctx, req.Id)
}

// PageDataSource 分页查询
func (r *dataSourceRepo) PageDataSource(ctx context.Context, req *v1.DataSourcePageReq) ([]*ent.DataSource, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.DataSource.
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
func (r *dataSourceRepo) genCondition(req *v1.DataSourceReq) []predicate.DataSource {
	if req == nil {
		return nil
	}
	list := make([]predicate.DataSource, 0)
	if req.Id > 0 {
		list = append(list, datasource.ID(req.Id))
	}
	if str.IsBlank(req.DbName) {
		list = append(list, datasource.DbNameContains(req.DbName))
	}
	if str.IsBlank(req.Host) {
		list = append(list, datasource.HostContains(req.Host))
	}
	if req.Port > 0 {
		list = append(list, datasource.Port(req.Port))
	}
	if str.IsBlank(req.Database) {
		list = append(list, datasource.DatabaseContains(req.Database))
	}
	if str.IsBlank(req.UserName) {
		list = append(list, datasource.UserNameContains(req.UserName))
	}
	if str.IsBlank(req.Pwd) {
		list = append(list, datasource.PwdContains(req.Pwd))
	}
	list = append(list, datasource.Status(req.Status))
	dbType := datasource.DbType(req.DbType)
	if datasource.DbTypeValidator(dbType) == nil {
		list = append(list, datasource.DbTypeEQ(dbType))
	}
	if req.ConnMaxIdleTime > 0 {
		list = append(list, datasource.ConnMaxIdleTime(req.ConnMaxIdleTime))
	}
	if req.ConnMaxLifeTime > 0 {
		list = append(list, datasource.ConnMaxLifeTime(req.ConnMaxLifeTime))
	}
	if req.MaxIdleConns > 0 {
		list = append(list, datasource.MaxIdleConns(req.MaxIdleConns))
	}
	if req.MaxOpenConns > 0 {
		list = append(list, datasource.MaxOpenConns(req.MaxOpenConns))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, datasource.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, datasource.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, datasource.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, datasource.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, datasource.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, datasource.TenantId(req.TenantId))
	}

	return list
}
