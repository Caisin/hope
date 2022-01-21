// Code generated by entc, DO NOT EDIT.

package ent

import (
	"hope/apps/param/internal/data/ent/novelpayconfig"
	"hope/apps/param/internal/data/ent/noveltag"
	"hope/apps/param/internal/data/ent/pageconfig"
	"hope/apps/param/internal/data/ent/qiniuconfig"
	"hope/apps/param/internal/data/ent/resourcegroup"
	"hope/apps/param/internal/data/ent/resourcestorage"
	"hope/apps/param/internal/data/ent/schema"
	"hope/apps/param/internal/data/ent/scoreproduct"
	"hope/apps/param/internal/data/ent/task"
	"hope/apps/param/internal/data/ent/useranalysisstatistics"
	"hope/apps/param/internal/data/ent/userconsume"
	"hope/apps/param/internal/data/ent/userresource"
	"hope/apps/param/internal/data/ent/userresourcerecord"
	"hope/apps/param/internal/data/ent/viptype"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	novelpayconfigFields := schema.NovelPayConfig{}.Fields()
	_ = novelpayconfigFields
	// novelpayconfigDescEffectTime is the schema descriptor for effectTime field.
	novelpayconfigDescEffectTime := novelpayconfigFields[20].Descriptor()
	// novelpayconfig.DefaultEffectTime holds the default value on creation for the effectTime field.
	novelpayconfig.DefaultEffectTime = novelpayconfigDescEffectTime.Default.(func() time.Time)
	// novelpayconfigDescExpiredTime is the schema descriptor for expiredTime field.
	novelpayconfigDescExpiredTime := novelpayconfigFields[21].Descriptor()
	// novelpayconfig.DefaultExpiredTime holds the default value on creation for the expiredTime field.
	novelpayconfig.DefaultExpiredTime = novelpayconfigDescExpiredTime.Default.(func() time.Time)
	// novelpayconfigDescCreatedAt is the schema descriptor for createdAt field.
	novelpayconfigDescCreatedAt := novelpayconfigFields[22].Descriptor()
	// novelpayconfig.DefaultCreatedAt holds the default value on creation for the createdAt field.
	novelpayconfig.DefaultCreatedAt = novelpayconfigDescCreatedAt.Default.(func() time.Time)
	// novelpayconfigDescUpdatedAt is the schema descriptor for updatedAt field.
	novelpayconfigDescUpdatedAt := novelpayconfigFields[23].Descriptor()
	// novelpayconfig.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	novelpayconfig.DefaultUpdatedAt = novelpayconfigDescUpdatedAt.Default.(func() time.Time)
	// novelpayconfig.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	novelpayconfig.UpdateDefaultUpdatedAt = novelpayconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// novelpayconfigDescCreateBy is the schema descriptor for createBy field.
	novelpayconfigDescCreateBy := novelpayconfigFields[24].Descriptor()
	// novelpayconfig.DefaultCreateBy holds the default value on creation for the createBy field.
	novelpayconfig.DefaultCreateBy = novelpayconfigDescCreateBy.Default.(int64)
	// novelpayconfigDescUpdateBy is the schema descriptor for updateBy field.
	novelpayconfigDescUpdateBy := novelpayconfigFields[25].Descriptor()
	// novelpayconfig.DefaultUpdateBy holds the default value on creation for the updateBy field.
	novelpayconfig.DefaultUpdateBy = novelpayconfigDescUpdateBy.Default.(int64)
	// novelpayconfigDescTenantId is the schema descriptor for tenantId field.
	novelpayconfigDescTenantId := novelpayconfigFields[26].Descriptor()
	// novelpayconfig.DefaultTenantId holds the default value on creation for the tenantId field.
	novelpayconfig.DefaultTenantId = novelpayconfigDescTenantId.Default.(int64)
	noveltagFields := schema.NovelTag{}.Fields()
	_ = noveltagFields
	// noveltagDescCreatedAt is the schema descriptor for createdAt field.
	noveltagDescCreatedAt := noveltagFields[3].Descriptor()
	// noveltag.DefaultCreatedAt holds the default value on creation for the createdAt field.
	noveltag.DefaultCreatedAt = noveltagDescCreatedAt.Default.(func() time.Time)
	// noveltagDescUpdatedAt is the schema descriptor for updatedAt field.
	noveltagDescUpdatedAt := noveltagFields[4].Descriptor()
	// noveltag.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	noveltag.DefaultUpdatedAt = noveltagDescUpdatedAt.Default.(func() time.Time)
	// noveltag.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	noveltag.UpdateDefaultUpdatedAt = noveltagDescUpdatedAt.UpdateDefault.(func() time.Time)
	// noveltagDescCreateBy is the schema descriptor for createBy field.
	noveltagDescCreateBy := noveltagFields[5].Descriptor()
	// noveltag.DefaultCreateBy holds the default value on creation for the createBy field.
	noveltag.DefaultCreateBy = noveltagDescCreateBy.Default.(int64)
	// noveltagDescUpdateBy is the schema descriptor for updateBy field.
	noveltagDescUpdateBy := noveltagFields[6].Descriptor()
	// noveltag.DefaultUpdateBy holds the default value on creation for the updateBy field.
	noveltag.DefaultUpdateBy = noveltagDescUpdateBy.Default.(int64)
	// noveltagDescTenantId is the schema descriptor for tenantId field.
	noveltagDescTenantId := noveltagFields[7].Descriptor()
	// noveltag.DefaultTenantId holds the default value on creation for the tenantId field.
	noveltag.DefaultTenantId = noveltagDescTenantId.Default.(int64)
	pageconfigFields := schema.PageConfig{}.Fields()
	_ = pageconfigFields
	// pageconfigDescCreatedAt is the schema descriptor for createdAt field.
	pageconfigDescCreatedAt := pageconfigFields[3].Descriptor()
	// pageconfig.DefaultCreatedAt holds the default value on creation for the createdAt field.
	pageconfig.DefaultCreatedAt = pageconfigDescCreatedAt.Default.(func() time.Time)
	// pageconfigDescUpdatedAt is the schema descriptor for updatedAt field.
	pageconfigDescUpdatedAt := pageconfigFields[4].Descriptor()
	// pageconfig.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	pageconfig.DefaultUpdatedAt = pageconfigDescUpdatedAt.Default.(func() time.Time)
	// pageconfig.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	pageconfig.UpdateDefaultUpdatedAt = pageconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// pageconfigDescCreateBy is the schema descriptor for createBy field.
	pageconfigDescCreateBy := pageconfigFields[5].Descriptor()
	// pageconfig.DefaultCreateBy holds the default value on creation for the createBy field.
	pageconfig.DefaultCreateBy = pageconfigDescCreateBy.Default.(int64)
	// pageconfigDescUpdateBy is the schema descriptor for updateBy field.
	pageconfigDescUpdateBy := pageconfigFields[6].Descriptor()
	// pageconfig.DefaultUpdateBy holds the default value on creation for the updateBy field.
	pageconfig.DefaultUpdateBy = pageconfigDescUpdateBy.Default.(int64)
	// pageconfigDescTenantId is the schema descriptor for tenantId field.
	pageconfigDescTenantId := pageconfigFields[7].Descriptor()
	// pageconfig.DefaultTenantId holds the default value on creation for the tenantId field.
	pageconfig.DefaultTenantId = pageconfigDescTenantId.Default.(int64)
	qiniuconfigFields := schema.QiniuConfig{}.Fields()
	_ = qiniuconfigFields
	// qiniuconfigDescCreatedAt is the schema descriptor for createdAt field.
	qiniuconfigDescCreatedAt := qiniuconfigFields[6].Descriptor()
	// qiniuconfig.DefaultCreatedAt holds the default value on creation for the createdAt field.
	qiniuconfig.DefaultCreatedAt = qiniuconfigDescCreatedAt.Default.(func() time.Time)
	// qiniuconfigDescUpdatedAt is the schema descriptor for updatedAt field.
	qiniuconfigDescUpdatedAt := qiniuconfigFields[7].Descriptor()
	// qiniuconfig.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	qiniuconfig.DefaultUpdatedAt = qiniuconfigDescUpdatedAt.Default.(func() time.Time)
	// qiniuconfig.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	qiniuconfig.UpdateDefaultUpdatedAt = qiniuconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// qiniuconfigDescCreateBy is the schema descriptor for createBy field.
	qiniuconfigDescCreateBy := qiniuconfigFields[8].Descriptor()
	// qiniuconfig.DefaultCreateBy holds the default value on creation for the createBy field.
	qiniuconfig.DefaultCreateBy = qiniuconfigDescCreateBy.Default.(int64)
	// qiniuconfigDescUpdateBy is the schema descriptor for updateBy field.
	qiniuconfigDescUpdateBy := qiniuconfigFields[9].Descriptor()
	// qiniuconfig.DefaultUpdateBy holds the default value on creation for the updateBy field.
	qiniuconfig.DefaultUpdateBy = qiniuconfigDescUpdateBy.Default.(int64)
	// qiniuconfigDescTenantId is the schema descriptor for tenantId field.
	qiniuconfigDescTenantId := qiniuconfigFields[10].Descriptor()
	// qiniuconfig.DefaultTenantId holds the default value on creation for the tenantId field.
	qiniuconfig.DefaultTenantId = qiniuconfigDescTenantId.Default.(int64)
	resourcegroupFields := schema.ResourceGroup{}.Fields()
	_ = resourcegroupFields
	// resourcegroupDescCreatedAt is the schema descriptor for createdAt field.
	resourcegroupDescCreatedAt := resourcegroupFields[1].Descriptor()
	// resourcegroup.DefaultCreatedAt holds the default value on creation for the createdAt field.
	resourcegroup.DefaultCreatedAt = resourcegroupDescCreatedAt.Default.(func() time.Time)
	// resourcegroupDescUpdatedAt is the schema descriptor for updatedAt field.
	resourcegroupDescUpdatedAt := resourcegroupFields[2].Descriptor()
	// resourcegroup.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	resourcegroup.DefaultUpdatedAt = resourcegroupDescUpdatedAt.Default.(func() time.Time)
	// resourcegroup.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	resourcegroup.UpdateDefaultUpdatedAt = resourcegroupDescUpdatedAt.UpdateDefault.(func() time.Time)
	// resourcegroupDescCreateBy is the schema descriptor for createBy field.
	resourcegroupDescCreateBy := resourcegroupFields[3].Descriptor()
	// resourcegroup.DefaultCreateBy holds the default value on creation for the createBy field.
	resourcegroup.DefaultCreateBy = resourcegroupDescCreateBy.Default.(int64)
	// resourcegroupDescUpdateBy is the schema descriptor for updateBy field.
	resourcegroupDescUpdateBy := resourcegroupFields[4].Descriptor()
	// resourcegroup.DefaultUpdateBy holds the default value on creation for the updateBy field.
	resourcegroup.DefaultUpdateBy = resourcegroupDescUpdateBy.Default.(int64)
	// resourcegroupDescTenantId is the schema descriptor for tenantId field.
	resourcegroupDescTenantId := resourcegroupFields[5].Descriptor()
	// resourcegroup.DefaultTenantId holds the default value on creation for the tenantId field.
	resourcegroup.DefaultTenantId = resourcegroupDescTenantId.Default.(int64)
	resourcestorageFields := schema.ResourceStorage{}.Fields()
	_ = resourcestorageFields
	// resourcestorageDescCreatedAt is the schema descriptor for createdAt field.
	resourcestorageDescCreatedAt := resourcestorageFields[18].Descriptor()
	// resourcestorage.DefaultCreatedAt holds the default value on creation for the createdAt field.
	resourcestorage.DefaultCreatedAt = resourcestorageDescCreatedAt.Default.(func() time.Time)
	// resourcestorageDescUpdatedAt is the schema descriptor for updatedAt field.
	resourcestorageDescUpdatedAt := resourcestorageFields[19].Descriptor()
	// resourcestorage.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	resourcestorage.DefaultUpdatedAt = resourcestorageDescUpdatedAt.Default.(func() time.Time)
	// resourcestorage.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	resourcestorage.UpdateDefaultUpdatedAt = resourcestorageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// resourcestorageDescCreateBy is the schema descriptor for createBy field.
	resourcestorageDescCreateBy := resourcestorageFields[20].Descriptor()
	// resourcestorage.DefaultCreateBy holds the default value on creation for the createBy field.
	resourcestorage.DefaultCreateBy = resourcestorageDescCreateBy.Default.(int64)
	// resourcestorageDescUpdateBy is the schema descriptor for updateBy field.
	resourcestorageDescUpdateBy := resourcestorageFields[21].Descriptor()
	// resourcestorage.DefaultUpdateBy holds the default value on creation for the updateBy field.
	resourcestorage.DefaultUpdateBy = resourcestorageDescUpdateBy.Default.(int64)
	// resourcestorageDescTenantId is the schema descriptor for tenantId field.
	resourcestorageDescTenantId := resourcestorageFields[22].Descriptor()
	// resourcestorage.DefaultTenantId holds the default value on creation for the tenantId field.
	resourcestorage.DefaultTenantId = resourcestorageDescTenantId.Default.(int64)
	scoreproductFields := schema.ScoreProduct{}.Fields()
	_ = scoreproductFields
	// scoreproductDescEffectTime is the schema descriptor for effectTime field.
	scoreproductDescEffectTime := scoreproductFields[5].Descriptor()
	// scoreproduct.DefaultEffectTime holds the default value on creation for the effectTime field.
	scoreproduct.DefaultEffectTime = scoreproductDescEffectTime.Default.(func() time.Time)
	// scoreproductDescExpiredTime is the schema descriptor for expiredTime field.
	scoreproductDescExpiredTime := scoreproductFields[6].Descriptor()
	// scoreproduct.DefaultExpiredTime holds the default value on creation for the expiredTime field.
	scoreproduct.DefaultExpiredTime = scoreproductDescExpiredTime.Default.(func() time.Time)
	// scoreproductDescCreatedAt is the schema descriptor for createdAt field.
	scoreproductDescCreatedAt := scoreproductFields[7].Descriptor()
	// scoreproduct.DefaultCreatedAt holds the default value on creation for the createdAt field.
	scoreproduct.DefaultCreatedAt = scoreproductDescCreatedAt.Default.(func() time.Time)
	// scoreproductDescUpdatedAt is the schema descriptor for updatedAt field.
	scoreproductDescUpdatedAt := scoreproductFields[8].Descriptor()
	// scoreproduct.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	scoreproduct.DefaultUpdatedAt = scoreproductDescUpdatedAt.Default.(func() time.Time)
	// scoreproduct.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	scoreproduct.UpdateDefaultUpdatedAt = scoreproductDescUpdatedAt.UpdateDefault.(func() time.Time)
	// scoreproductDescCreateBy is the schema descriptor for createBy field.
	scoreproductDescCreateBy := scoreproductFields[9].Descriptor()
	// scoreproduct.DefaultCreateBy holds the default value on creation for the createBy field.
	scoreproduct.DefaultCreateBy = scoreproductDescCreateBy.Default.(int64)
	// scoreproductDescUpdateBy is the schema descriptor for updateBy field.
	scoreproductDescUpdateBy := scoreproductFields[10].Descriptor()
	// scoreproduct.DefaultUpdateBy holds the default value on creation for the updateBy field.
	scoreproduct.DefaultUpdateBy = scoreproductDescUpdateBy.Default.(int64)
	// scoreproductDescTenantId is the schema descriptor for tenantId field.
	scoreproductDescTenantId := scoreproductFields[11].Descriptor()
	// scoreproduct.DefaultTenantId holds the default value on creation for the tenantId field.
	scoreproduct.DefaultTenantId = scoreproductDescTenantId.Default.(int64)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescEffectTime is the schema descriptor for effectTime field.
	taskDescEffectTime := taskFields[19].Descriptor()
	// task.DefaultEffectTime holds the default value on creation for the effectTime field.
	task.DefaultEffectTime = taskDescEffectTime.Default.(func() time.Time)
	// taskDescExpiredTime is the schema descriptor for expiredTime field.
	taskDescExpiredTime := taskFields[20].Descriptor()
	// task.DefaultExpiredTime holds the default value on creation for the expiredTime field.
	task.DefaultExpiredTime = taskDescExpiredTime.Default.(func() time.Time)
	// taskDescCreatedAt is the schema descriptor for createdAt field.
	taskDescCreatedAt := taskFields[21].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the createdAt field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescUpdatedAt is the schema descriptor for updatedAt field.
	taskDescUpdatedAt := taskFields[22].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() time.Time)
	// task.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	task.UpdateDefaultUpdatedAt = taskDescUpdatedAt.UpdateDefault.(func() time.Time)
	// taskDescCreateBy is the schema descriptor for createBy field.
	taskDescCreateBy := taskFields[23].Descriptor()
	// task.DefaultCreateBy holds the default value on creation for the createBy field.
	task.DefaultCreateBy = taskDescCreateBy.Default.(int64)
	// taskDescUpdateBy is the schema descriptor for updateBy field.
	taskDescUpdateBy := taskFields[24].Descriptor()
	// task.DefaultUpdateBy holds the default value on creation for the updateBy field.
	task.DefaultUpdateBy = taskDescUpdateBy.Default.(int64)
	// taskDescTenantId is the schema descriptor for tenantId field.
	taskDescTenantId := taskFields[25].Descriptor()
	// task.DefaultTenantId holds the default value on creation for the tenantId field.
	task.DefaultTenantId = taskDescTenantId.Default.(int64)
	useranalysisstatisticsFields := schema.UserAnalysisStatistics{}.Fields()
	_ = useranalysisstatisticsFields
	// useranalysisstatisticsDescCreatedAt is the schema descriptor for createdAt field.
	useranalysisstatisticsDescCreatedAt := useranalysisstatisticsFields[20].Descriptor()
	// useranalysisstatistics.DefaultCreatedAt holds the default value on creation for the createdAt field.
	useranalysisstatistics.DefaultCreatedAt = useranalysisstatisticsDescCreatedAt.Default.(func() time.Time)
	// useranalysisstatisticsDescUpdatedAt is the schema descriptor for updatedAt field.
	useranalysisstatisticsDescUpdatedAt := useranalysisstatisticsFields[21].Descriptor()
	// useranalysisstatistics.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	useranalysisstatistics.DefaultUpdatedAt = useranalysisstatisticsDescUpdatedAt.Default.(func() time.Time)
	// useranalysisstatistics.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	useranalysisstatistics.UpdateDefaultUpdatedAt = useranalysisstatisticsDescUpdatedAt.UpdateDefault.(func() time.Time)
	// useranalysisstatisticsDescCreateBy is the schema descriptor for createBy field.
	useranalysisstatisticsDescCreateBy := useranalysisstatisticsFields[22].Descriptor()
	// useranalysisstatistics.DefaultCreateBy holds the default value on creation for the createBy field.
	useranalysisstatistics.DefaultCreateBy = useranalysisstatisticsDescCreateBy.Default.(int64)
	// useranalysisstatisticsDescUpdateBy is the schema descriptor for updateBy field.
	useranalysisstatisticsDescUpdateBy := useranalysisstatisticsFields[23].Descriptor()
	// useranalysisstatistics.DefaultUpdateBy holds the default value on creation for the updateBy field.
	useranalysisstatistics.DefaultUpdateBy = useranalysisstatisticsDescUpdateBy.Default.(int64)
	// useranalysisstatisticsDescTenantId is the schema descriptor for tenantId field.
	useranalysisstatisticsDescTenantId := useranalysisstatisticsFields[24].Descriptor()
	// useranalysisstatistics.DefaultTenantId holds the default value on creation for the tenantId field.
	useranalysisstatistics.DefaultTenantId = useranalysisstatisticsDescTenantId.Default.(int64)
	userconsumeFields := schema.UserConsume{}.Fields()
	_ = userconsumeFields
	// userconsumeDescCreatedAt is the schema descriptor for createdAt field.
	userconsumeDescCreatedAt := userconsumeFields[4].Descriptor()
	// userconsume.DefaultCreatedAt holds the default value on creation for the createdAt field.
	userconsume.DefaultCreatedAt = userconsumeDescCreatedAt.Default.(func() time.Time)
	// userconsumeDescUpdatedAt is the schema descriptor for updatedAt field.
	userconsumeDescUpdatedAt := userconsumeFields[5].Descriptor()
	// userconsume.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	userconsume.DefaultUpdatedAt = userconsumeDescUpdatedAt.Default.(func() time.Time)
	// userconsume.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	userconsume.UpdateDefaultUpdatedAt = userconsumeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userconsumeDescCreateBy is the schema descriptor for createBy field.
	userconsumeDescCreateBy := userconsumeFields[6].Descriptor()
	// userconsume.DefaultCreateBy holds the default value on creation for the createBy field.
	userconsume.DefaultCreateBy = userconsumeDescCreateBy.Default.(int64)
	// userconsumeDescUpdateBy is the schema descriptor for updateBy field.
	userconsumeDescUpdateBy := userconsumeFields[7].Descriptor()
	// userconsume.DefaultUpdateBy holds the default value on creation for the updateBy field.
	userconsume.DefaultUpdateBy = userconsumeDescUpdateBy.Default.(int64)
	// userconsumeDescTenantId is the schema descriptor for tenantId field.
	userconsumeDescTenantId := userconsumeFields[8].Descriptor()
	// userconsume.DefaultTenantId holds the default value on creation for the tenantId field.
	userconsume.DefaultTenantId = userconsumeDescTenantId.Default.(int64)
	userresourceFields := schema.UserResource{}.Fields()
	_ = userresourceFields
	// userresourceDescCreatedAt is the schema descriptor for createdAt field.
	userresourceDescCreatedAt := userresourceFields[4].Descriptor()
	// userresource.DefaultCreatedAt holds the default value on creation for the createdAt field.
	userresource.DefaultCreatedAt = userresourceDescCreatedAt.Default.(func() time.Time)
	// userresourceDescUpdatedAt is the schema descriptor for updatedAt field.
	userresourceDescUpdatedAt := userresourceFields[5].Descriptor()
	// userresource.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	userresource.DefaultUpdatedAt = userresourceDescUpdatedAt.Default.(func() time.Time)
	// userresource.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	userresource.UpdateDefaultUpdatedAt = userresourceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userresourceDescCreateBy is the schema descriptor for createBy field.
	userresourceDescCreateBy := userresourceFields[6].Descriptor()
	// userresource.DefaultCreateBy holds the default value on creation for the createBy field.
	userresource.DefaultCreateBy = userresourceDescCreateBy.Default.(int64)
	// userresourceDescUpdateBy is the schema descriptor for updateBy field.
	userresourceDescUpdateBy := userresourceFields[7].Descriptor()
	// userresource.DefaultUpdateBy holds the default value on creation for the updateBy field.
	userresource.DefaultUpdateBy = userresourceDescUpdateBy.Default.(int64)
	// userresourceDescTenantId is the schema descriptor for tenantId field.
	userresourceDescTenantId := userresourceFields[8].Descriptor()
	// userresource.DefaultTenantId holds the default value on creation for the tenantId field.
	userresource.DefaultTenantId = userresourceDescTenantId.Default.(int64)
	userresourcerecordFields := schema.UserResourceRecord{}.Fields()
	_ = userresourcerecordFields
	// userresourcerecordDescEffectTime is the schema descriptor for effectTime field.
	userresourcerecordDescEffectTime := userresourcerecordFields[8].Descriptor()
	// userresourcerecord.DefaultEffectTime holds the default value on creation for the effectTime field.
	userresourcerecord.DefaultEffectTime = userresourcerecordDescEffectTime.Default.(func() time.Time)
	// userresourcerecordDescExpiredTime is the schema descriptor for expiredTime field.
	userresourcerecordDescExpiredTime := userresourcerecordFields[9].Descriptor()
	// userresourcerecord.DefaultExpiredTime holds the default value on creation for the expiredTime field.
	userresourcerecord.DefaultExpiredTime = userresourcerecordDescExpiredTime.Default.(func() time.Time)
	// userresourcerecordDescCreatedAt is the schema descriptor for createdAt field.
	userresourcerecordDescCreatedAt := userresourcerecordFields[10].Descriptor()
	// userresourcerecord.DefaultCreatedAt holds the default value on creation for the createdAt field.
	userresourcerecord.DefaultCreatedAt = userresourcerecordDescCreatedAt.Default.(func() time.Time)
	// userresourcerecordDescUpdatedAt is the schema descriptor for updatedAt field.
	userresourcerecordDescUpdatedAt := userresourcerecordFields[11].Descriptor()
	// userresourcerecord.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	userresourcerecord.DefaultUpdatedAt = userresourcerecordDescUpdatedAt.Default.(func() time.Time)
	// userresourcerecord.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	userresourcerecord.UpdateDefaultUpdatedAt = userresourcerecordDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userresourcerecordDescCreateBy is the schema descriptor for createBy field.
	userresourcerecordDescCreateBy := userresourcerecordFields[12].Descriptor()
	// userresourcerecord.DefaultCreateBy holds the default value on creation for the createBy field.
	userresourcerecord.DefaultCreateBy = userresourcerecordDescCreateBy.Default.(int64)
	// userresourcerecordDescUpdateBy is the schema descriptor for updateBy field.
	userresourcerecordDescUpdateBy := userresourcerecordFields[13].Descriptor()
	// userresourcerecord.DefaultUpdateBy holds the default value on creation for the updateBy field.
	userresourcerecord.DefaultUpdateBy = userresourcerecordDescUpdateBy.Default.(int64)
	// userresourcerecordDescTenantId is the schema descriptor for tenantId field.
	userresourcerecordDescTenantId := userresourcerecordFields[14].Descriptor()
	// userresourcerecord.DefaultTenantId holds the default value on creation for the tenantId field.
	userresourcerecord.DefaultTenantId = userresourcerecordDescTenantId.Default.(int64)
	viptypeFields := schema.VipType{}.Fields()
	_ = viptypeFields
	// viptypeDescCreatedAt is the schema descriptor for createdAt field.
	viptypeDescCreatedAt := viptypeFields[6].Descriptor()
	// viptype.DefaultCreatedAt holds the default value on creation for the createdAt field.
	viptype.DefaultCreatedAt = viptypeDescCreatedAt.Default.(func() time.Time)
	// viptypeDescUpdatedAt is the schema descriptor for updatedAt field.
	viptypeDescUpdatedAt := viptypeFields[7].Descriptor()
	// viptype.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	viptype.DefaultUpdatedAt = viptypeDescUpdatedAt.Default.(func() time.Time)
	// viptype.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	viptype.UpdateDefaultUpdatedAt = viptypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// viptypeDescCreateBy is the schema descriptor for createBy field.
	viptypeDescCreateBy := viptypeFields[8].Descriptor()
	// viptype.DefaultCreateBy holds the default value on creation for the createBy field.
	viptype.DefaultCreateBy = viptypeDescCreateBy.Default.(int64)
	// viptypeDescUpdateBy is the schema descriptor for updateBy field.
	viptypeDescUpdateBy := viptypeFields[9].Descriptor()
	// viptype.DefaultUpdateBy holds the default value on creation for the updateBy field.
	viptype.DefaultUpdateBy = viptypeDescUpdateBy.Default.(int64)
	// viptypeDescTenantId is the schema descriptor for tenantId field.
	viptypeDescTenantId := viptypeFields[10].Descriptor()
	// viptype.DefaultTenantId holds the default value on creation for the tenantId field.
	viptype.DefaultTenantId = viptypeDescTenantId.Default.(int64)
}