// Code generated by entc, DO NOT EDIT.

package ent

import (
	"hope/apps/admin/internal/data/ent/casbinrule"
	"hope/apps/admin/internal/data/ent/schema"
	"hope/apps/admin/internal/data/ent/sysapi"
	"hope/apps/admin/internal/data/ent/sysconfig"
	"hope/apps/admin/internal/data/ent/sysdept"
	"hope/apps/admin/internal/data/ent/sysdictdata"
	"hope/apps/admin/internal/data/ent/sysdicttype"
	"hope/apps/admin/internal/data/ent/sysjob"
	"hope/apps/admin/internal/data/ent/sysjoblog"
	"hope/apps/admin/internal/data/ent/sysloginlog"
	"hope/apps/admin/internal/data/ent/sysmenu"
	"hope/apps/admin/internal/data/ent/sysoperalog"
	"hope/apps/admin/internal/data/ent/syspost"
	"hope/apps/admin/internal/data/ent/sysrole"
	"hope/apps/admin/internal/data/ent/sysuser"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	casbinruleFields := schema.CasbinRule{}.Fields()
	_ = casbinruleFields
	// casbinruleDescCreatedAt is the schema descriptor for createdAt field.
	casbinruleDescCreatedAt := casbinruleFields[7].Descriptor()
	// casbinrule.DefaultCreatedAt holds the default value on creation for the createdAt field.
	casbinrule.DefaultCreatedAt = casbinruleDescCreatedAt.Default.(func() time.Time)
	// casbinruleDescUpdatedAt is the schema descriptor for updatedAt field.
	casbinruleDescUpdatedAt := casbinruleFields[8].Descriptor()
	// casbinrule.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	casbinrule.DefaultUpdatedAt = casbinruleDescUpdatedAt.Default.(func() time.Time)
	// casbinrule.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	casbinrule.UpdateDefaultUpdatedAt = casbinruleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// casbinruleDescCreateBy is the schema descriptor for createBy field.
	casbinruleDescCreateBy := casbinruleFields[9].Descriptor()
	// casbinrule.DefaultCreateBy holds the default value on creation for the createBy field.
	casbinrule.DefaultCreateBy = casbinruleDescCreateBy.Default.(int64)
	// casbinruleDescUpdateBy is the schema descriptor for updateBy field.
	casbinruleDescUpdateBy := casbinruleFields[10].Descriptor()
	// casbinrule.DefaultUpdateBy holds the default value on creation for the updateBy field.
	casbinrule.DefaultUpdateBy = casbinruleDescUpdateBy.Default.(int64)
	// casbinruleDescTenantId is the schema descriptor for tenantId field.
	casbinruleDescTenantId := casbinruleFields[11].Descriptor()
	// casbinrule.DefaultTenantId holds the default value on creation for the tenantId field.
	casbinrule.DefaultTenantId = casbinruleDescTenantId.Default.(int64)
	sysapiFields := schema.SysApi{}.Fields()
	_ = sysapiFields
	// sysapiDescCreatedAt is the schema descriptor for createdAt field.
	sysapiDescCreatedAt := sysapiFields[5].Descriptor()
	// sysapi.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysapi.DefaultCreatedAt = sysapiDescCreatedAt.Default.(func() time.Time)
	// sysapiDescUpdatedAt is the schema descriptor for updatedAt field.
	sysapiDescUpdatedAt := sysapiFields[6].Descriptor()
	// sysapi.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysapi.DefaultUpdatedAt = sysapiDescUpdatedAt.Default.(func() time.Time)
	// sysapi.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysapi.UpdateDefaultUpdatedAt = sysapiDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysapiDescCreateBy is the schema descriptor for createBy field.
	sysapiDescCreateBy := sysapiFields[7].Descriptor()
	// sysapi.DefaultCreateBy holds the default value on creation for the createBy field.
	sysapi.DefaultCreateBy = sysapiDescCreateBy.Default.(int64)
	// sysapiDescUpdateBy is the schema descriptor for updateBy field.
	sysapiDescUpdateBy := sysapiFields[8].Descriptor()
	// sysapi.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysapi.DefaultUpdateBy = sysapiDescUpdateBy.Default.(int64)
	// sysapiDescTenantId is the schema descriptor for tenantId field.
	sysapiDescTenantId := sysapiFields[9].Descriptor()
	// sysapi.DefaultTenantId holds the default value on creation for the tenantId field.
	sysapi.DefaultTenantId = sysapiDescTenantId.Default.(int64)
	sysconfigFields := schema.SysConfig{}.Fields()
	_ = sysconfigFields
	// sysconfigDescCreatedAt is the schema descriptor for createdAt field.
	sysconfigDescCreatedAt := sysconfigFields[7].Descriptor()
	// sysconfig.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysconfig.DefaultCreatedAt = sysconfigDescCreatedAt.Default.(func() time.Time)
	// sysconfigDescUpdatedAt is the schema descriptor for updatedAt field.
	sysconfigDescUpdatedAt := sysconfigFields[8].Descriptor()
	// sysconfig.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysconfig.DefaultUpdatedAt = sysconfigDescUpdatedAt.Default.(func() time.Time)
	// sysconfig.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysconfig.UpdateDefaultUpdatedAt = sysconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysconfigDescCreateBy is the schema descriptor for createBy field.
	sysconfigDescCreateBy := sysconfigFields[9].Descriptor()
	// sysconfig.DefaultCreateBy holds the default value on creation for the createBy field.
	sysconfig.DefaultCreateBy = sysconfigDescCreateBy.Default.(int64)
	// sysconfigDescUpdateBy is the schema descriptor for updateBy field.
	sysconfigDescUpdateBy := sysconfigFields[10].Descriptor()
	// sysconfig.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysconfig.DefaultUpdateBy = sysconfigDescUpdateBy.Default.(int64)
	// sysconfigDescTenantId is the schema descriptor for tenantId field.
	sysconfigDescTenantId := sysconfigFields[11].Descriptor()
	// sysconfig.DefaultTenantId holds the default value on creation for the tenantId field.
	sysconfig.DefaultTenantId = sysconfigDescTenantId.Default.(int64)
	sysdeptFields := schema.SysDept{}.Fields()
	_ = sysdeptFields
	// sysdeptDescCreatedAt is the schema descriptor for createdAt field.
	sysdeptDescCreatedAt := sysdeptFields[7].Descriptor()
	// sysdept.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysdept.DefaultCreatedAt = sysdeptDescCreatedAt.Default.(func() time.Time)
	// sysdeptDescUpdatedAt is the schema descriptor for updatedAt field.
	sysdeptDescUpdatedAt := sysdeptFields[8].Descriptor()
	// sysdept.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysdept.DefaultUpdatedAt = sysdeptDescUpdatedAt.Default.(func() time.Time)
	// sysdept.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysdept.UpdateDefaultUpdatedAt = sysdeptDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysdeptDescCreateBy is the schema descriptor for createBy field.
	sysdeptDescCreateBy := sysdeptFields[9].Descriptor()
	// sysdept.DefaultCreateBy holds the default value on creation for the createBy field.
	sysdept.DefaultCreateBy = sysdeptDescCreateBy.Default.(int64)
	// sysdeptDescUpdateBy is the schema descriptor for updateBy field.
	sysdeptDescUpdateBy := sysdeptFields[10].Descriptor()
	// sysdept.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysdept.DefaultUpdateBy = sysdeptDescUpdateBy.Default.(int64)
	// sysdeptDescTenantId is the schema descriptor for tenantId field.
	sysdeptDescTenantId := sysdeptFields[11].Descriptor()
	// sysdept.DefaultTenantId holds the default value on creation for the tenantId field.
	sysdept.DefaultTenantId = sysdeptDescTenantId.Default.(int64)
	sysdictdataFields := schema.SysDictData{}.Fields()
	_ = sysdictdataFields
	// sysdictdataDescCreatedAt is the schema descriptor for createdAt field.
	sysdictdataDescCreatedAt := sysdictdataFields[9].Descriptor()
	// sysdictdata.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysdictdata.DefaultCreatedAt = sysdictdataDescCreatedAt.Default.(func() time.Time)
	// sysdictdataDescUpdatedAt is the schema descriptor for updatedAt field.
	sysdictdataDescUpdatedAt := sysdictdataFields[10].Descriptor()
	// sysdictdata.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysdictdata.DefaultUpdatedAt = sysdictdataDescUpdatedAt.Default.(func() time.Time)
	// sysdictdata.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysdictdata.UpdateDefaultUpdatedAt = sysdictdataDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysdictdataDescCreateBy is the schema descriptor for createBy field.
	sysdictdataDescCreateBy := sysdictdataFields[11].Descriptor()
	// sysdictdata.DefaultCreateBy holds the default value on creation for the createBy field.
	sysdictdata.DefaultCreateBy = sysdictdataDescCreateBy.Default.(int64)
	// sysdictdataDescUpdateBy is the schema descriptor for updateBy field.
	sysdictdataDescUpdateBy := sysdictdataFields[12].Descriptor()
	// sysdictdata.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysdictdata.DefaultUpdateBy = sysdictdataDescUpdateBy.Default.(int64)
	// sysdictdataDescTenantId is the schema descriptor for tenantId field.
	sysdictdataDescTenantId := sysdictdataFields[13].Descriptor()
	// sysdictdata.DefaultTenantId holds the default value on creation for the tenantId field.
	sysdictdata.DefaultTenantId = sysdictdataDescTenantId.Default.(int64)
	sysdicttypeFields := schema.SysDictType{}.Fields()
	_ = sysdicttypeFields
	// sysdicttypeDescCreatedAt is the schema descriptor for createdAt field.
	sysdicttypeDescCreatedAt := sysdicttypeFields[4].Descriptor()
	// sysdicttype.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysdicttype.DefaultCreatedAt = sysdicttypeDescCreatedAt.Default.(func() time.Time)
	// sysdicttypeDescUpdatedAt is the schema descriptor for updatedAt field.
	sysdicttypeDescUpdatedAt := sysdicttypeFields[5].Descriptor()
	// sysdicttype.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysdicttype.DefaultUpdatedAt = sysdicttypeDescUpdatedAt.Default.(func() time.Time)
	// sysdicttype.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysdicttype.UpdateDefaultUpdatedAt = sysdicttypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysdicttypeDescCreateBy is the schema descriptor for createBy field.
	sysdicttypeDescCreateBy := sysdicttypeFields[6].Descriptor()
	// sysdicttype.DefaultCreateBy holds the default value on creation for the createBy field.
	sysdicttype.DefaultCreateBy = sysdicttypeDescCreateBy.Default.(int64)
	// sysdicttypeDescUpdateBy is the schema descriptor for updateBy field.
	sysdicttypeDescUpdateBy := sysdicttypeFields[7].Descriptor()
	// sysdicttype.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysdicttype.DefaultUpdateBy = sysdicttypeDescUpdateBy.Default.(int64)
	// sysdicttypeDescTenantId is the schema descriptor for tenantId field.
	sysdicttypeDescTenantId := sysdicttypeFields[8].Descriptor()
	// sysdicttype.DefaultTenantId holds the default value on creation for the tenantId field.
	sysdicttype.DefaultTenantId = sysdicttypeDescTenantId.Default.(int64)
	sysjobFields := schema.SysJob{}.Fields()
	_ = sysjobFields
	// sysjobDescCreatedAt is the schema descriptor for createdAt field.
	sysjobDescCreatedAt := sysjobFields[10].Descriptor()
	// sysjob.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysjob.DefaultCreatedAt = sysjobDescCreatedAt.Default.(func() time.Time)
	// sysjobDescUpdatedAt is the schema descriptor for updatedAt field.
	sysjobDescUpdatedAt := sysjobFields[11].Descriptor()
	// sysjob.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysjob.DefaultUpdatedAt = sysjobDescUpdatedAt.Default.(func() time.Time)
	// sysjob.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysjob.UpdateDefaultUpdatedAt = sysjobDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysjobDescCreateBy is the schema descriptor for createBy field.
	sysjobDescCreateBy := sysjobFields[12].Descriptor()
	// sysjob.DefaultCreateBy holds the default value on creation for the createBy field.
	sysjob.DefaultCreateBy = sysjobDescCreateBy.Default.(int64)
	// sysjobDescUpdateBy is the schema descriptor for updateBy field.
	sysjobDescUpdateBy := sysjobFields[13].Descriptor()
	// sysjob.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysjob.DefaultUpdateBy = sysjobDescUpdateBy.Default.(int64)
	// sysjobDescTenantId is the schema descriptor for tenantId field.
	sysjobDescTenantId := sysjobFields[14].Descriptor()
	// sysjob.DefaultTenantId holds the default value on creation for the tenantId field.
	sysjob.DefaultTenantId = sysjobDescTenantId.Default.(int64)
	sysjoblogFields := schema.SysJobLog{}.Fields()
	_ = sysjoblogFields
	// sysjoblogDescCreatedAt is the schema descriptor for createdAt field.
	sysjoblogDescCreatedAt := sysjoblogFields[6].Descriptor()
	// sysjoblog.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysjoblog.DefaultCreatedAt = sysjoblogDescCreatedAt.Default.(func() time.Time)
	// sysjoblogDescUpdatedAt is the schema descriptor for updatedAt field.
	sysjoblogDescUpdatedAt := sysjoblogFields[7].Descriptor()
	// sysjoblog.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysjoblog.DefaultUpdatedAt = sysjoblogDescUpdatedAt.Default.(func() time.Time)
	// sysjoblog.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysjoblog.UpdateDefaultUpdatedAt = sysjoblogDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysjoblogDescCreateBy is the schema descriptor for createBy field.
	sysjoblogDescCreateBy := sysjoblogFields[8].Descriptor()
	// sysjoblog.DefaultCreateBy holds the default value on creation for the createBy field.
	sysjoblog.DefaultCreateBy = sysjoblogDescCreateBy.Default.(int64)
	// sysjoblogDescUpdateBy is the schema descriptor for updateBy field.
	sysjoblogDescUpdateBy := sysjoblogFields[9].Descriptor()
	// sysjoblog.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysjoblog.DefaultUpdateBy = sysjoblogDescUpdateBy.Default.(int64)
	// sysjoblogDescTenantId is the schema descriptor for tenantId field.
	sysjoblogDescTenantId := sysjoblogFields[10].Descriptor()
	// sysjoblog.DefaultTenantId holds the default value on creation for the tenantId field.
	sysjoblog.DefaultTenantId = sysjoblogDescTenantId.Default.(int64)
	sysloginlogFields := schema.SysLoginLog{}.Fields()
	_ = sysloginlogFields
	// sysloginlogDescCreatedAt is the schema descriptor for createdAt field.
	sysloginlogDescCreatedAt := sysloginlogFields[10].Descriptor()
	// sysloginlog.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysloginlog.DefaultCreatedAt = sysloginlogDescCreatedAt.Default.(func() time.Time)
	// sysloginlogDescUpdatedAt is the schema descriptor for updatedAt field.
	sysloginlogDescUpdatedAt := sysloginlogFields[11].Descriptor()
	// sysloginlog.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysloginlog.DefaultUpdatedAt = sysloginlogDescUpdatedAt.Default.(func() time.Time)
	// sysloginlog.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysloginlog.UpdateDefaultUpdatedAt = sysloginlogDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysloginlogDescCreateBy is the schema descriptor for createBy field.
	sysloginlogDescCreateBy := sysloginlogFields[12].Descriptor()
	// sysloginlog.DefaultCreateBy holds the default value on creation for the createBy field.
	sysloginlog.DefaultCreateBy = sysloginlogDescCreateBy.Default.(int64)
	// sysloginlogDescUpdateBy is the schema descriptor for updateBy field.
	sysloginlogDescUpdateBy := sysloginlogFields[13].Descriptor()
	// sysloginlog.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysloginlog.DefaultUpdateBy = sysloginlogDescUpdateBy.Default.(int64)
	// sysloginlogDescTenantId is the schema descriptor for tenantId field.
	sysloginlogDescTenantId := sysloginlogFields[14].Descriptor()
	// sysloginlog.DefaultTenantId holds the default value on creation for the tenantId field.
	sysloginlog.DefaultTenantId = sysloginlogDescTenantId.Default.(int64)
	sysmenuFields := schema.SysMenu{}.Fields()
	_ = sysmenuFields
	// sysmenuDescParentId is the schema descriptor for parentId field.
	sysmenuDescParentId := sysmenuFields[0].Descriptor()
	// sysmenu.DefaultParentId holds the default value on creation for the parentId field.
	sysmenu.DefaultParentId = sysmenuDescParentId.Default.(int64)
	// sysmenuDescCheckPermission is the schema descriptor for checkPermission field.
	sysmenuDescCheckPermission := sysmenuFields[18].Descriptor()
	// sysmenu.DefaultCheckPermission holds the default value on creation for the checkPermission field.
	sysmenu.DefaultCheckPermission = sysmenuDescCheckPermission.Default.(bool)
	// sysmenuDescCreatedAt is the schema descriptor for createdAt field.
	sysmenuDescCreatedAt := sysmenuFields[20].Descriptor()
	// sysmenu.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysmenu.DefaultCreatedAt = sysmenuDescCreatedAt.Default.(func() time.Time)
	// sysmenuDescUpdatedAt is the schema descriptor for updatedAt field.
	sysmenuDescUpdatedAt := sysmenuFields[21].Descriptor()
	// sysmenu.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysmenu.DefaultUpdatedAt = sysmenuDescUpdatedAt.Default.(func() time.Time)
	// sysmenu.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysmenu.UpdateDefaultUpdatedAt = sysmenuDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysmenuDescCreateBy is the schema descriptor for createBy field.
	sysmenuDescCreateBy := sysmenuFields[22].Descriptor()
	// sysmenu.DefaultCreateBy holds the default value on creation for the createBy field.
	sysmenu.DefaultCreateBy = sysmenuDescCreateBy.Default.(int64)
	// sysmenuDescUpdateBy is the schema descriptor for updateBy field.
	sysmenuDescUpdateBy := sysmenuFields[23].Descriptor()
	// sysmenu.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysmenu.DefaultUpdateBy = sysmenuDescUpdateBy.Default.(int64)
	// sysmenuDescTenantId is the schema descriptor for tenantId field.
	sysmenuDescTenantId := sysmenuFields[24].Descriptor()
	// sysmenu.DefaultTenantId holds the default value on creation for the tenantId field.
	sysmenu.DefaultTenantId = sysmenuDescTenantId.Default.(int64)
	sysoperalogFields := schema.SysOperaLog{}.Fields()
	_ = sysoperalogFields
	// sysoperalogDescCreatedAt is the schema descriptor for createdAt field.
	sysoperalogDescCreatedAt := sysoperalogFields[23].Descriptor()
	// sysoperalog.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysoperalog.DefaultCreatedAt = sysoperalogDescCreatedAt.Default.(func() time.Time)
	// sysoperalogDescUpdatedAt is the schema descriptor for updatedAt field.
	sysoperalogDescUpdatedAt := sysoperalogFields[24].Descriptor()
	// sysoperalog.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysoperalog.DefaultUpdatedAt = sysoperalogDescUpdatedAt.Default.(func() time.Time)
	// sysoperalog.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysoperalog.UpdateDefaultUpdatedAt = sysoperalogDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysoperalogDescCreateBy is the schema descriptor for createBy field.
	sysoperalogDescCreateBy := sysoperalogFields[25].Descriptor()
	// sysoperalog.DefaultCreateBy holds the default value on creation for the createBy field.
	sysoperalog.DefaultCreateBy = sysoperalogDescCreateBy.Default.(int64)
	// sysoperalogDescUpdateBy is the schema descriptor for updateBy field.
	sysoperalogDescUpdateBy := sysoperalogFields[26].Descriptor()
	// sysoperalog.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysoperalog.DefaultUpdateBy = sysoperalogDescUpdateBy.Default.(int64)
	// sysoperalogDescTenantId is the schema descriptor for tenantId field.
	sysoperalogDescTenantId := sysoperalogFields[27].Descriptor()
	// sysoperalog.DefaultTenantId holds the default value on creation for the tenantId field.
	sysoperalog.DefaultTenantId = sysoperalogDescTenantId.Default.(int64)
	syspostFields := schema.SysPost{}.Fields()
	_ = syspostFields
	// syspostDescCreatedAt is the schema descriptor for createdAt field.
	syspostDescCreatedAt := syspostFields[5].Descriptor()
	// syspost.DefaultCreatedAt holds the default value on creation for the createdAt field.
	syspost.DefaultCreatedAt = syspostDescCreatedAt.Default.(func() time.Time)
	// syspostDescUpdatedAt is the schema descriptor for updatedAt field.
	syspostDescUpdatedAt := syspostFields[6].Descriptor()
	// syspost.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	syspost.DefaultUpdatedAt = syspostDescUpdatedAt.Default.(func() time.Time)
	// syspost.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	syspost.UpdateDefaultUpdatedAt = syspostDescUpdatedAt.UpdateDefault.(func() time.Time)
	// syspostDescCreateBy is the schema descriptor for createBy field.
	syspostDescCreateBy := syspostFields[7].Descriptor()
	// syspost.DefaultCreateBy holds the default value on creation for the createBy field.
	syspost.DefaultCreateBy = syspostDescCreateBy.Default.(int64)
	// syspostDescUpdateBy is the schema descriptor for updateBy field.
	syspostDescUpdateBy := syspostFields[8].Descriptor()
	// syspost.DefaultUpdateBy holds the default value on creation for the updateBy field.
	syspost.DefaultUpdateBy = syspostDescUpdateBy.Default.(int64)
	// syspostDescTenantId is the schema descriptor for tenantId field.
	syspostDescTenantId := syspostFields[9].Descriptor()
	// syspost.DefaultTenantId holds the default value on creation for the tenantId field.
	syspost.DefaultTenantId = syspostDescTenantId.Default.(int64)
	sysroleFields := schema.SysRole{}.Fields()
	_ = sysroleFields
	// sysroleDescCreatedAt is the schema descriptor for createdAt field.
	sysroleDescCreatedAt := sysroleFields[8].Descriptor()
	// sysrole.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysrole.DefaultCreatedAt = sysroleDescCreatedAt.Default.(func() time.Time)
	// sysroleDescUpdatedAt is the schema descriptor for updatedAt field.
	sysroleDescUpdatedAt := sysroleFields[9].Descriptor()
	// sysrole.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysrole.DefaultUpdatedAt = sysroleDescUpdatedAt.Default.(func() time.Time)
	// sysrole.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysrole.UpdateDefaultUpdatedAt = sysroleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysroleDescCreateBy is the schema descriptor for createBy field.
	sysroleDescCreateBy := sysroleFields[10].Descriptor()
	// sysrole.DefaultCreateBy holds the default value on creation for the createBy field.
	sysrole.DefaultCreateBy = sysroleDescCreateBy.Default.(int64)
	// sysroleDescUpdateBy is the schema descriptor for updateBy field.
	sysroleDescUpdateBy := sysroleFields[11].Descriptor()
	// sysrole.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysrole.DefaultUpdateBy = sysroleDescUpdateBy.Default.(int64)
	// sysroleDescTenantId is the schema descriptor for tenantId field.
	sysroleDescTenantId := sysroleFields[12].Descriptor()
	// sysrole.DefaultTenantId holds the default value on creation for the tenantId field.
	sysrole.DefaultTenantId = sysroleDescTenantId.Default.(int64)
	sysuserFields := schema.SysUser{}.Fields()
	_ = sysuserFields
	// sysuserDescCreatedAt is the schema descriptor for createdAt field.
	sysuserDescCreatedAt := sysuserFields[14].Descriptor()
	// sysuser.DefaultCreatedAt holds the default value on creation for the createdAt field.
	sysuser.DefaultCreatedAt = sysuserDescCreatedAt.Default.(func() time.Time)
	// sysuserDescUpdatedAt is the schema descriptor for updatedAt field.
	sysuserDescUpdatedAt := sysuserFields[15].Descriptor()
	// sysuser.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	sysuser.DefaultUpdatedAt = sysuserDescUpdatedAt.Default.(func() time.Time)
	// sysuser.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	sysuser.UpdateDefaultUpdatedAt = sysuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysuserDescCreateBy is the schema descriptor for createBy field.
	sysuserDescCreateBy := sysuserFields[16].Descriptor()
	// sysuser.DefaultCreateBy holds the default value on creation for the createBy field.
	sysuser.DefaultCreateBy = sysuserDescCreateBy.Default.(int64)
	// sysuserDescUpdateBy is the schema descriptor for updateBy field.
	sysuserDescUpdateBy := sysuserFields[17].Descriptor()
	// sysuser.DefaultUpdateBy holds the default value on creation for the updateBy field.
	sysuser.DefaultUpdateBy = sysuserDescUpdateBy.Default.(int64)
	// sysuserDescTenantId is the schema descriptor for tenantId field.
	sysuserDescTenantId := sysuserFields[18].Descriptor()
	// sysuser.DefaultTenantId holds the default value on creation for the tenantId field.
	sysuser.DefaultTenantId = sysuserDescTenantId.Default.(int64)
}
