package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"

	casbinrule "hope/api/admin/casbinrule/v1"
	sysapi "hope/api/admin/sysapi/v1"
	syscolumns "hope/api/admin/syscolumns/v1"
	sysconfig "hope/api/admin/sysconfig/v1"
	sysdept "hope/api/admin/sysdept/v1"
	sysdictdata "hope/api/admin/sysdictdata/v1"
	sysdicttype "hope/api/admin/sysdicttype/v1"
	sysjob "hope/api/admin/sysjob/v1"
	sysjoblog "hope/api/admin/sysjoblog/v1"
	sysloginlog "hope/api/admin/sysloginlog/v1"
	sysmenu "hope/api/admin/sysmenu/v1"
	sysoperalog "hope/api/admin/sysoperalog/v1"
	syspost "hope/api/admin/syspost/v1"
	sysrole "hope/api/admin/sysrole/v1"
	systables "hope/api/admin/systables/v1"
	sysuser "hope/api/admin/sysuser/v1"
	"hope/apps/admin/internal/service"
)

func RegisterGRPCServer(

	casbinRuleService *service.CasbinRuleService,
	sysApiService *service.SysApiService,
	sysColumnsService *service.SysColumnsService,
	sysConfigService *service.SysConfigService,
	sysDeptService *service.SysDeptService,
	sysDictDataService *service.SysDictDataService,
	sysDictTypeService *service.SysDictTypeService,
	sysJobService *service.SysJobService,
	sysJobLogService *service.SysJobLogService,
	sysLoginLogService *service.SysLoginLogService,
	sysMenuService *service.SysMenuService,
	sysOperaLogService *service.SysOperaLogService,
	sysPostService *service.SysPostService,
	sysRoleService *service.SysRoleService,
	sysTablesService *service.SysTablesService,
	sysUserService *service.SysUserService,
) []func(*grpc.Server) {
	list := make([]func(*grpc.Server), 0)

	list = append(list, func(srv *grpc.Server) {
		casbinrule.RegisterCasbinRuleServer(srv, casbinRuleService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysapi.RegisterSysApiServer(srv, sysApiService)
	})
	list = append(list, func(srv *grpc.Server) {
		syscolumns.RegisterSysColumnsServer(srv, sysColumnsService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysconfig.RegisterSysConfigServer(srv, sysConfigService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysdept.RegisterSysDeptServer(srv, sysDeptService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysdictdata.RegisterSysDictDataServer(srv, sysDictDataService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysdicttype.RegisterSysDictTypeServer(srv, sysDictTypeService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysjob.RegisterSysJobServer(srv, sysJobService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysjoblog.RegisterSysJobLogServer(srv, sysJobLogService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysloginlog.RegisterSysLoginLogServer(srv, sysLoginLogService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysmenu.RegisterSysMenuServer(srv, sysMenuService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysoperalog.RegisterSysOperaLogServer(srv, sysOperaLogService)
	})
	list = append(list, func(srv *grpc.Server) {
		syspost.RegisterSysPostServer(srv, sysPostService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysrole.RegisterSysRoleServer(srv, sysRoleService)
	})
	list = append(list, func(srv *grpc.Server) {
		systables.RegisterSysTablesServer(srv, sysTablesService)
	})
	list = append(list, func(srv *grpc.Server) {
		sysuser.RegisterSysUserServer(srv, sysUserService)
	})
	return list
}
