// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CasbinRulesColumns holds the columns for the "casbin_rules" table.
	CasbinRulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "p_type", Type: field.TypeString, Nullable: true},
		{Name: "v0", Type: field.TypeString, Nullable: true},
		{Name: "v1", Type: field.TypeString, Nullable: true},
		{Name: "v2", Type: field.TypeString, Nullable: true},
		{Name: "v3", Type: field.TypeString, Nullable: true},
		{Name: "v4", Type: field.TypeString, Nullable: true},
		{Name: "v5", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
	}
	// CasbinRulesTable holds the schema information for the "casbin_rules" table.
	CasbinRulesTable = &schema.Table{
		Name:       "casbin_rules",
		Columns:    CasbinRulesColumns,
		PrimaryKey: []*schema.Column{CasbinRulesColumns[0]},
	}
	// SysApisColumns holds the columns for the "sys_apis" table.
	SysApisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "handle", Type: field.TypeString, Nullable: true},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "path", Type: field.TypeString, Nullable: true},
		{Name: "action", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
	}
	// SysApisTable holds the schema information for the "sys_apis" table.
	SysApisTable = &schema.Table{
		Name:       "sys_apis",
		Columns:    SysApisColumns,
		PrimaryKey: []*schema.Column{SysApisColumns[0]},
	}
	// SysColumnsColumns holds the columns for the "sys_columns" table.
	SysColumnsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "column_id", Type: field.TypeInt32},
		{Name: "column_name", Type: field.TypeString, Nullable: true},
		{Name: "column_comment", Type: field.TypeString, Nullable: true},
		{Name: "column_type", Type: field.TypeString, Nullable: true},
		{Name: "go_type", Type: field.TypeString, Nullable: true},
		{Name: "go_field", Type: field.TypeString, Nullable: true},
		{Name: "json_field", Type: field.TypeString, Nullable: true},
		{Name: "is_pk", Type: field.TypeString, Nullable: true},
		{Name: "is_increment", Type: field.TypeString, Nullable: true},
		{Name: "is_required", Type: field.TypeString, Nullable: true},
		{Name: "is_insert", Type: field.TypeString, Nullable: true},
		{Name: "is_edit", Type: field.TypeString, Nullable: true},
		{Name: "is_list", Type: field.TypeString, Nullable: true},
		{Name: "is_query", Type: field.TypeString, Nullable: true},
		{Name: "query_type", Type: field.TypeString, Nullable: true},
		{Name: "html_type", Type: field.TypeString, Nullable: true},
		{Name: "dict_type", Type: field.TypeString, Nullable: true},
		{Name: "sort", Type: field.TypeInt32, Nullable: true},
		{Name: "list", Type: field.TypeString, Nullable: true},
		{Name: "pk", Type: field.TypeBool, Nullable: true},
		{Name: "required", Type: field.TypeBool, Nullable: true},
		{Name: "super_column", Type: field.TypeBool, Nullable: true},
		{Name: "usable_column", Type: field.TypeBool, Nullable: true},
		{Name: "increment", Type: field.TypeBool, Nullable: true},
		{Name: "insert", Type: field.TypeBool, Nullable: true},
		{Name: "edit", Type: field.TypeBool, Nullable: true},
		{Name: "query", Type: field.TypeBool, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "fk_label_name", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
	}
	// SysColumnsTable holds the schema information for the "sys_columns" table.
	SysColumnsTable = &schema.Table{
		Name:       "sys_columns",
		Columns:    SysColumnsColumns,
		PrimaryKey: []*schema.Column{SysColumnsColumns[0]},
	}
	// SysConfigsColumns holds the columns for the "sys_configs" table.
	SysConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "config_name", Type: field.TypeString, Nullable: true},
		{Name: "config_key", Type: field.TypeString, Nullable: true},
		{Name: "config_value", Type: field.TypeString, Nullable: true},
		{Name: "config_type", Type: field.TypeString, Nullable: true},
		{Name: "is_frontend", Type: field.TypeInt32, Nullable: true},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"U", "E"}, Default: "U"},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
	}
	// SysConfigsTable holds the schema information for the "sys_configs" table.
	SysConfigsTable = &schema.Table{
		Name:       "sys_configs",
		Columns:    SysConfigsColumns,
		PrimaryKey: []*schema.Column{SysConfigsColumns[0]},
	}
	// SysDeptsColumns holds the columns for the "sys_depts" table.
	SysDeptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "dept_path", Type: field.TypeString, Nullable: true},
		{Name: "dept_name", Type: field.TypeString, Nullable: true},
		{Name: "sort", Type: field.TypeInt32, Nullable: true},
		{Name: "leader", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt32, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
		{Name: "sys_dept_childes", Type: field.TypeInt64, Nullable: true},
	}
	// SysDeptsTable holds the schema information for the "sys_depts" table.
	SysDeptsTable = &schema.Table{
		Name:       "sys_depts",
		Columns:    SysDeptsColumns,
		PrimaryKey: []*schema.Column{SysDeptsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_depts_sys_depts_childes",
				Columns:    []*schema.Column{SysDeptsColumns[13]},
				RefColumns: []*schema.Column{SysDeptsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysDictDataColumns holds the columns for the "sys_dict_data" table.
	SysDictDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "dict_sort", Type: field.TypeInt32, Nullable: true},
		{Name: "dict_label", Type: field.TypeString, Nullable: true},
		{Name: "dict_value", Type: field.TypeString, Nullable: true},
		{Name: "is_default", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt32, Nullable: true},
		{Name: "default", Type: field.TypeString, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
		{Name: "sys_dict_type_data_list", Type: field.TypeInt64, Nullable: true},
	}
	// SysDictDataTable holds the schema information for the "sys_dict_data" table.
	SysDictDataTable = &schema.Table{
		Name:       "sys_dict_data",
		Columns:    SysDictDataColumns,
		PrimaryKey: []*schema.Column{SysDictDataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_dict_data_sys_dict_types_dataList",
				Columns:    []*schema.Column{SysDictDataColumns[13]},
				RefColumns: []*schema.Column{SysDictTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysDictTypesColumns holds the columns for the "sys_dict_types" table.
	SysDictTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "dict_name", Type: field.TypeString, Nullable: true},
		{Name: "dict_type", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt32, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
	}
	// SysDictTypesTable holds the schema information for the "sys_dict_types" table.
	SysDictTypesTable = &schema.Table{
		Name:       "sys_dict_types",
		Columns:    SysDictTypesColumns,
		PrimaryKey: []*schema.Column{SysDictTypesColumns[0]},
	}
	// SysJobsColumns holds the columns for the "sys_jobs" table.
	SysJobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "job_name", Type: field.TypeString, Nullable: true},
		{Name: "job_group", Type: field.TypeString, Nullable: true},
		{Name: "job_type", Type: field.TypeInt32, Nullable: true},
		{Name: "cron_expression", Type: field.TypeString, Nullable: true},
		{Name: "invoke_target", Type: field.TypeString, Nullable: true},
		{Name: "args", Type: field.TypeString, Nullable: true},
		{Name: "exec_policy", Type: field.TypeInt32, Nullable: true},
		{Name: "concurrent", Type: field.TypeInt32, Nullable: true},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"Pause", "Run", "Stop"}, Default: "Stop"},
		{Name: "entry_id", Type: field.TypeInt32, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
	}
	// SysJobsTable holds the schema information for the "sys_jobs" table.
	SysJobsTable = &schema.Table{
		Name:       "sys_jobs",
		Columns:    SysJobsColumns,
		PrimaryKey: []*schema.Column{SysJobsColumns[0]},
	}
	// SysJobLogsColumns holds the columns for the "sys_job_logs" table.
	SysJobLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "job_id", Type: field.TypeInt32, Nullable: true},
		{Name: "job_name", Type: field.TypeString, Nullable: true},
		{Name: "entry_id", Type: field.TypeInt32, Nullable: true},
		{Name: "status", Type: field.TypeBool, Nullable: true},
		{Name: "duration", Type: field.TypeInt64, Nullable: true},
		{Name: "info", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
		{Name: "sys_job_logs", Type: field.TypeInt64, Nullable: true},
	}
	// SysJobLogsTable holds the schema information for the "sys_job_logs" table.
	SysJobLogsTable = &schema.Table{
		Name:       "sys_job_logs",
		Columns:    SysJobLogsColumns,
		PrimaryKey: []*schema.Column{SysJobLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_job_logs_sys_jobs_logs",
				Columns:    []*schema.Column{SysJobLogsColumns[12]},
				RefColumns: []*schema.Column{SysJobsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysLoginLogsColumns holds the columns for the "sys_login_logs" table.
	SysLoginLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "status", Type: field.TypeString, Nullable: true},
		{Name: "ipaddr", Type: field.TypeString, Nullable: true},
		{Name: "login_location", Type: field.TypeString, Nullable: true},
		{Name: "browser", Type: field.TypeString, Nullable: true},
		{Name: "os", Type: field.TypeString, Nullable: true},
		{Name: "platform", Type: field.TypeString, Nullable: true},
		{Name: "login_time", Type: field.TypeTime, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "msg", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
		{Name: "sys_user_login_logs", Type: field.TypeInt64, Nullable: true},
	}
	// SysLoginLogsTable holds the schema information for the "sys_login_logs" table.
	SysLoginLogsTable = &schema.Table{
		Name:       "sys_login_logs",
		Columns:    SysLoginLogsColumns,
		PrimaryKey: []*schema.Column{SysLoginLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_login_logs_sys_users_loginLogs",
				Columns:    []*schema.Column{SysLoginLogsColumns[15]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "menu_name", Type: field.TypeString, Nullable: true},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "icon", Type: field.TypeString, Nullable: true},
		{Name: "path", Type: field.TypeString, Nullable: true},
		{Name: "paths", Type: field.TypeString, Nullable: true},
		{Name: "menu_type", Type: field.TypeString, Nullable: true},
		{Name: "action", Type: field.TypeString, Nullable: true},
		{Name: "permission", Type: field.TypeString, Nullable: true},
		{Name: "no_cache", Type: field.TypeBool, Nullable: true},
		{Name: "breadcrumb", Type: field.TypeString, Nullable: true},
		{Name: "component", Type: field.TypeString, Nullable: true},
		{Name: "sort", Type: field.TypeInt32, Nullable: true},
		{Name: "visible", Type: field.TypeBool, Nullable: true},
		{Name: "is_frame", Type: field.TypeBool, Nullable: true},
		{Name: "sys_api", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
		{Name: "sys_menu_childes", Type: field.TypeInt64, Nullable: true},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:       "sys_menus",
		Columns:    SysMenusColumns,
		PrimaryKey: []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menus_sys_menus_childes",
				Columns:    []*schema.Column{SysMenusColumns[21]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysOperaLogsColumns holds the columns for the "sys_opera_logs" table.
	SysOperaLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64, Nullable: true},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "request_id", Type: field.TypeString, Nullable: true},
		{Name: "business_type", Type: field.TypeString, Nullable: true},
		{Name: "business_types", Type: field.TypeString, Nullable: true},
		{Name: "method", Type: field.TypeString, Nullable: true},
		{Name: "request_method", Type: field.TypeString, Nullable: true},
		{Name: "operator_type", Type: field.TypeString, Nullable: true},
		{Name: "oper_name", Type: field.TypeString, Nullable: true},
		{Name: "dept_name", Type: field.TypeString, Nullable: true},
		{Name: "oper_url", Type: field.TypeString, Nullable: true},
		{Name: "oper_ip", Type: field.TypeString, Nullable: true},
		{Name: "browser", Type: field.TypeString, Nullable: true},
		{Name: "os", Type: field.TypeString, Nullable: true},
		{Name: "platform", Type: field.TypeString, Nullable: true},
		{Name: "oper_location", Type: field.TypeString, Nullable: true},
		{Name: "oper_param", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeString, Nullable: true},
		{Name: "oper_time", Type: field.TypeTime, Nullable: true},
		{Name: "json_result", Type: field.TypeString, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "latency_time", Type: field.TypeString, Nullable: true},
		{Name: "user_agent", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
		{Name: "sys_user_opera_logs", Type: field.TypeInt64, Nullable: true},
	}
	// SysOperaLogsTable holds the schema information for the "sys_opera_logs" table.
	SysOperaLogsTable = &schema.Table{
		Name:       "sys_opera_logs",
		Columns:    SysOperaLogsColumns,
		PrimaryKey: []*schema.Column{SysOperaLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_opera_logs_sys_users_operaLogs",
				Columns:    []*schema.Column{SysOperaLogsColumns[29]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysPostsColumns holds the columns for the "sys_posts" table.
	SysPostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "post_name", Type: field.TypeString, Nullable: true},
		{Name: "post_code", Type: field.TypeString, Nullable: true},
		{Name: "sort", Type: field.TypeInt32, Nullable: true},
		{Name: "status", Type: field.TypeInt32, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
	}
	// SysPostsTable holds the schema information for the "sys_posts" table.
	SysPostsTable = &schema.Table{
		Name:       "sys_posts",
		Columns:    SysPostsColumns,
		PrimaryKey: []*schema.Column{SysPostsColumns[0]},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "role_name", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeString, Nullable: true},
		{Name: "role_key", Type: field.TypeString, Nullable: true},
		{Name: "role_sort", Type: field.TypeInt32, Nullable: true},
		{Name: "flag", Type: field.TypeString, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "admin", Type: field.TypeBool, Nullable: true},
		{Name: "data_scope", Type: field.TypeString, Nullable: true},
		{Name: "sys_dept", Type: field.TypeString, Nullable: true},
		{Name: "sys_menu", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:       "sys_roles",
		Columns:    SysRolesColumns,
		PrimaryKey: []*schema.Column{SysRolesColumns[0]},
	}
	// SysTablesColumns holds the columns for the "sys_tables" table.
	SysTablesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "table_name", Type: field.TypeString, Nullable: true},
		{Name: "table_comment", Type: field.TypeString, Nullable: true},
		{Name: "class_name", Type: field.TypeString, Nullable: true},
		{Name: "tpl_category", Type: field.TypeString, Nullable: true},
		{Name: "package_name", Type: field.TypeString, Nullable: true},
		{Name: "module_name", Type: field.TypeString, Nullable: true},
		{Name: "module_front_name", Type: field.TypeString, Nullable: true},
		{Name: "business_name", Type: field.TypeString, Nullable: true},
		{Name: "function_name", Type: field.TypeString, Nullable: true},
		{Name: "function_author", Type: field.TypeString, Nullable: true},
		{Name: "pk_column", Type: field.TypeString, Nullable: true},
		{Name: "pk_go_field", Type: field.TypeString, Nullable: true},
		{Name: "pk_json_field", Type: field.TypeString, Nullable: true},
		{Name: "options", Type: field.TypeString, Nullable: true},
		{Name: "tree_code", Type: field.TypeString, Nullable: true},
		{Name: "tree_parent_code", Type: field.TypeString, Nullable: true},
		{Name: "tree_name", Type: field.TypeString, Nullable: true},
		{Name: "tree", Type: field.TypeBool, Nullable: true},
		{Name: "crud", Type: field.TypeBool, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "is_data_scope", Type: field.TypeInt32, Nullable: true},
		{Name: "is_actions", Type: field.TypeInt32, Nullable: true},
		{Name: "is_auth", Type: field.TypeInt32, Nullable: true},
		{Name: "is_logical_delete", Type: field.TypeString, Nullable: true},
		{Name: "logical_delete", Type: field.TypeBool, Nullable: true},
		{Name: "logical_delete_column", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
	}
	// SysTablesTable holds the schema information for the "sys_tables" table.
	SysTablesTable = &schema.Table{
		Name:       "sys_tables",
		Columns:    SysTablesColumns,
		PrimaryKey: []*schema.Column{SysTablesColumns[0]},
	}
	// SysUsersColumns holds the columns for the "sys_users" table.
	SysUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "username", Type: field.TypeString, Nullable: true},
		{Name: "nick_name", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "role_id", Type: field.TypeInt32, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "sex", Type: field.TypeInt32, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeString, Nullable: true},
		{Name: "ext_info", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "create_by", Type: field.TypeInt64, Default: 0},
		{Name: "update_by", Type: field.TypeInt64, Default: 0},
		{Name: "tenant_id", Type: field.TypeInt64, Default: 0},
		{Name: "sys_dept_users", Type: field.TypeInt64, Nullable: true},
		{Name: "sys_post_users", Type: field.TypeInt64, Nullable: true},
	}
	// SysUsersTable holds the schema information for the "sys_users" table.
	SysUsersTable = &schema.Table{
		Name:       "sys_users",
		Columns:    SysUsersColumns,
		PrimaryKey: []*schema.Column{SysUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_users_sys_depts_users",
				Columns:    []*schema.Column{SysUsersColumns[16]},
				RefColumns: []*schema.Column{SysDeptsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "sys_users_sys_posts_users",
				Columns:    []*schema.Column{SysUsersColumns[17]},
				RefColumns: []*schema.Column{SysPostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysRoleMenusColumns holds the columns for the "sys_role_menus" table.
	SysRoleMenusColumns = []*schema.Column{
		{Name: "sys_role_id", Type: field.TypeInt},
		{Name: "sys_menu_id", Type: field.TypeInt},
	}
	// SysRoleMenusTable holds the schema information for the "sys_role_menus" table.
	SysRoleMenusTable = &schema.Table{
		Name:       "sys_role_menus",
		Columns:    SysRoleMenusColumns,
		PrimaryKey: []*schema.Column{SysRoleMenusColumns[0], SysRoleMenusColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_role_menus_sys_role_id",
				Columns:    []*schema.Column{SysRoleMenusColumns[0]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "sys_role_menus_sys_menu_id",
				Columns:    []*schema.Column{SysRoleMenusColumns[1]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CasbinRulesTable,
		SysApisTable,
		SysColumnsTable,
		SysConfigsTable,
		SysDeptsTable,
		SysDictDataTable,
		SysDictTypesTable,
		SysJobsTable,
		SysJobLogsTable,
		SysLoginLogsTable,
		SysMenusTable,
		SysOperaLogsTable,
		SysPostsTable,
		SysRolesTable,
		SysTablesTable,
		SysUsersTable,
		SysRoleMenusTable,
	}
)

func init() {
	SysDeptsTable.ForeignKeys[0].RefTable = SysDeptsTable
	SysDictDataTable.ForeignKeys[0].RefTable = SysDictTypesTable
	SysJobLogsTable.ForeignKeys[0].RefTable = SysJobsTable
	SysLoginLogsTable.ForeignKeys[0].RefTable = SysUsersTable
	SysMenusTable.ForeignKeys[0].RefTable = SysMenusTable
	SysOperaLogsTable.ForeignKeys[0].RefTable = SysUsersTable
	SysUsersTable.ForeignKeys[0].RefTable = SysDeptsTable
	SysUsersTable.ForeignKeys[1].RefTable = SysPostsTable
	SysRoleMenusTable.ForeignKeys[0].RefTable = SysRolesTable
	SysRoleMenusTable.ForeignKeys[1].RefTable = SysMenusTable
}