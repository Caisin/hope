package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"

	auth "hope/api/admin/auth/v1"
	casbinrule "hope/api/admin/casbinrule/v1"
	sysapi "hope/api/admin/sysapi/v1"
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
	sysuser "hope/api/admin/sysuser/v1"
	"hope/apps/admin/internal/service"
)

func RegisterHTTPServer(

	casbinRuleService *service.CasbinRuleService,
	sysApiService *service.SysApiService,
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
	sysUserService *service.SysUserService,
	authService *service.AuthService,
) []func(*http.Server) {
	list := make([]func(*http.Server), 0)

	list = append(list, func(srv *http.Server) {
		casbinrule.RegisterCasbinRuleHTTPServer(srv, casbinRuleService)
	})
	list = append(list, func(srv *http.Server) {
		sysapi.RegisterSysApiHTTPServer(srv, sysApiService)
	})
	list = append(list, func(srv *http.Server) {
		sysconfig.RegisterSysConfigHTTPServer(srv, sysConfigService)
	})
	list = append(list, func(srv *http.Server) {
		sysdept.RegisterSysDeptHTTPServer(srv, sysDeptService)
	})
	list = append(list, func(srv *http.Server) {
		sysdictdata.RegisterSysDictDataHTTPServer(srv, sysDictDataService)
	})
	list = append(list, func(srv *http.Server) {
		sysdicttype.RegisterSysDictTypeHTTPServer(srv, sysDictTypeService)
	})
	list = append(list, func(srv *http.Server) {
		sysjob.RegisterSysJobHTTPServer(srv, sysJobService)
	})
	list = append(list, func(srv *http.Server) {
		sysjoblog.RegisterSysJobLogHTTPServer(srv, sysJobLogService)
	})
	list = append(list, func(srv *http.Server) {
		sysloginlog.RegisterSysLoginLogHTTPServer(srv, sysLoginLogService)
	})
	list = append(list, func(srv *http.Server) {
		sysmenu.RegisterSysMenuHTTPServer(srv, sysMenuService)
	})
	list = append(list, func(srv *http.Server) {
		sysoperalog.RegisterSysOperaLogHTTPServer(srv, sysOperaLogService)
	})
	list = append(list, func(srv *http.Server) {
		syspost.RegisterSysPostHTTPServer(srv, sysPostService)
	})
	list = append(list, func(srv *http.Server) {
		sysrole.RegisterSysRoleHTTPServer(srv, sysRoleService)
	})
	list = append(list, func(srv *http.Server) {
		sysuser.RegisterSysUserHTTPServer(srv, sysUserService)
	})
	list = append(list, func(srv *http.Server) {
		auth.RegisterAuthHTTPServer(srv, authService)
	})
	return list
}
