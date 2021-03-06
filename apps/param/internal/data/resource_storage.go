package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/param/resourcestorage/v1"
	"hope/apps/param/internal/biz"
	"hope/apps/param/internal/data/ent"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/resourcestorage"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type resourceStorageRepo struct {
	data *Data
	log  *log.Helper
}

// NewResourceStorageRepo .
func NewResourceStorageRepo(data *Data, logger log.Logger) biz.ResourceStorageRepo {
	return &resourceStorageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateResourceStorage 创建
func (r *resourceStorageRepo) CreateResourceStorage(ctx context.Context, req *v1.ResourceStorageCreateReq) (*ent.ResourceStorage, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.ResourceStorage.Create().
		SetGroupId(req.GroupId).
		SetStorageType(req.StorageType).
		SetRealName(req.RealName).
		SetBucket(req.Bucket).
		SetName(req.Name).
		SetSuffix(req.Suffix).
		SetPath(req.Path).
		SetType(req.Type).
		SetSize(req.Size).
		SetDeleteUrl(req.DeleteUrl).
		SetFilename(req.Filename).
		SetKey(req.Key).
		SetHeight(req.Height).
		SetURL(req.Url).
		SetUsername(req.Username).
		SetWidth(req.Width).
		SetMd5code(req.Md5Code).
		SetRemark(req.Remark).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteResourceStorage 删除
func (r *resourceStorageRepo) DeleteResourceStorage(ctx context.Context, req *v1.ResourceStorageDeleteReq) error {
	return r.data.db.ResourceStorage.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteResourceStorage 批量删除
func (r *resourceStorageRepo) BatchDeleteResourceStorage(ctx context.Context, req *v1.ResourceStorageBatchDeleteReq) (int, error) {
	return r.data.db.ResourceStorage.Delete().Where(resourcestorage.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateResourceStorage 更新
func (r *resourceStorageRepo) UpdateResourceStorage(ctx context.Context, req *v1.ResourceStorageUpdateReq) (*ent.ResourceStorage, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.ResourceStorage.UpdateOneID(req.Id).
		SetGroupId(req.GroupId).
		SetStorageType(req.StorageType).
		SetRealName(req.RealName).
		SetBucket(req.Bucket).
		SetName(req.Name).
		SetSuffix(req.Suffix).
		SetPath(req.Path).
		SetType(req.Type).
		SetSize(req.Size).
		SetDeleteUrl(req.DeleteUrl).
		SetFilename(req.Filename).
		SetKey(req.Key).
		SetHeight(req.Height).
		SetURL(req.Url).
		SetUsername(req.Username).
		SetWidth(req.Width).
		SetMd5code(req.Md5Code).
		SetRemark(req.Remark).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetResourceStorage 根据Id查询
func (r *resourceStorageRepo) GetResourceStorage(ctx context.Context, req *v1.ResourceStorageReq) (*ent.ResourceStorage, error) {
	return r.data.db.ResourceStorage.Get(ctx, req.Id)
}

// PageResourceStorage 分页查询
func (r *resourceStorageRepo) PageResourceStorage(ctx context.Context, req *v1.ResourceStoragePageReq) ([]*ent.ResourceStorage, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.ResourceStorage.
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
func (r *resourceStorageRepo) genCondition(req *v1.ResourceStorageReq) []predicate.ResourceStorage {
	if req == nil {
		return nil
	}
	list := make([]predicate.ResourceStorage, 0)
	if req.Id > 0 {
		list = append(list, resourcestorage.ID(req.Id))
	}
	if req.GroupId > 0 {
		list = append(list, resourcestorage.GroupId(req.GroupId))
	}
	if req.StorageType > 0 {
		list = append(list, resourcestorage.StorageType(req.StorageType))
	}
	if str.IsNotBlank(req.RealName) {
		list = append(list, resourcestorage.RealNameContains(req.RealName))
	}
	if str.IsNotBlank(req.Bucket) {
		list = append(list, resourcestorage.BucketContains(req.Bucket))
	}
	if str.IsNotBlank(req.Name) {
		list = append(list, resourcestorage.NameContains(req.Name))
	}
	if str.IsNotBlank(req.Suffix) {
		list = append(list, resourcestorage.SuffixContains(req.Suffix))
	}
	if str.IsNotBlank(req.Path) {
		list = append(list, resourcestorage.PathContains(req.Path))
	}
	if str.IsNotBlank(req.Type) {
		list = append(list, resourcestorage.TypeContains(req.Type))
	}
	if str.IsNotBlank(req.Size) {
		list = append(list, resourcestorage.SizeContains(req.Size))
	}
	if str.IsNotBlank(req.DeleteUrl) {
		list = append(list, resourcestorage.DeleteUrlContains(req.DeleteUrl))
	}
	if str.IsNotBlank(req.Filename) {
		list = append(list, resourcestorage.FilenameContains(req.Filename))
	}
	if str.IsNotBlank(req.Key) {
		list = append(list, resourcestorage.KeyContains(req.Key))
	}
	if str.IsNotBlank(req.Height) {
		list = append(list, resourcestorage.HeightContains(req.Height))
	}
	if str.IsNotBlank(req.Url) {
		list = append(list, resourcestorage.URLContains(req.Url))
	}
	if str.IsNotBlank(req.Username) {
		list = append(list, resourcestorage.UsernameContains(req.Username))
	}
	if str.IsNotBlank(req.Width) {
		list = append(list, resourcestorage.WidthContains(req.Width))
	}
	if str.IsNotBlank(req.Md5Code) {
		list = append(list, resourcestorage.Md5codeContains(req.Md5Code))
	}
	if str.IsNotBlank(req.Remark) {
		list = append(list, resourcestorage.RemarkContains(req.Remark))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, resourcestorage.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, resourcestorage.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, resourcestorage.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, resourcestorage.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, resourcestorage.TenantId(req.TenantId))
	}

	return list
}
