package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/{{.model}}/{{.pkg}}/v1"
	"hope/apps/{{.model}}/internal/biz"
	"hope/apps/{{.model}}/internal/convert"
	"hope/apps/{{.model}}/internal/data/ent"
	"hope/apps/{{.model}}/internal/data/ent/{{.pkg}}"
	"hope/apps/{{.model}}/internal/data/ent/predicate"
	"hope/pkg/util/str"
	"time"
)

type {{.llName}}Repo struct {
	data *Data
	log  *log.Helper
}

// New{{.name}}Repo .
func New{{.name}}Repo(data *Data, logger log.Logger) biz.{{.name}}Repo {
	return &{{.llName}}Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create{{.name}} 创建
func (r *{{.llName}}Repo) Create{{.name}}(ctx context.Context, req *v1.{{.name}}CreateReq) (*ent.{{.name}}, error) {
	now := time.Now()
	return r.data.db.{{.name}}.Create().
{{ genCreateSetFields .fields }}	SetCreatedAt(now).
	SetUpdatedAt(now).
	Save(ctx)

}

// Delete{{.name}} 删除
func (r *{{.llName}}Repo) Delete{{.name}}(ctx context.Context, req *v1.{{.name}}DeleteReq) error {
	return r.data.db.{{.name}}.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDelete{{.name}} 批量删除
func (r *{{.llName}}Repo) BatchDelete{{.name}}(ctx context.Context, req *v1.{{.name}}BatchDeleteReq) (int, error) {
	return r.data.db.{{.name}}.Delete().Where({{.pkg}}.IDIn(req.Ids...)).Exec(ctx)
}

// Update{{.name}} 更新
func (r *{{.llName}}Repo) Update{{.name}}(ctx context.Context, req *v1.{{.name}}UpdateReq) (*ent.{{.name}}, error) {
	return r.data.db.{{.name}}.UpdateOne(convert.{{.name}}UpdateReq2Data(req)).Save(ctx)
}

// Get{{.name}} 根据Id查询
func (r *{{.llName}}Repo) Get{{.name}}(ctx context.Context, req *v1.{{.name}}Req) (*ent.{{.name}}, error) {
	return r.data.db.{{.name}}.Get(ctx, req.Id)
}

// Page{{.name}} 分页查询
func (r *{{.llName}}Repo) Page{{.name}}(ctx context.Context, req *v1.{{.name}}PageReq) ([]*ent.{{.name}}, error) {
	pagin := req.Pagin
	query := r.data.db.{{.name}}.
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
	query.Limit(int(pagin.GetPage())).
		Offset(int(pagin.GetOffSet()))
	if pagin.NeedOrder() {
		if pagin.IsDesc() {
			query.Order(ent.Desc(pagin.GetField()))
		} else {
			query.Order(ent.Asc(pagin.GetField()))
		}
	}
	return query.All(ctx)
}

// genCondition 构造查询条件
func (r *{{.llName}}Repo) genCondition(req *v1.{{.name}}Req) []predicate.{{.name}} {
	list := make([]predicate.{{.name}}, 0)
	if req.Id > 0 {
		list = append(list, {{.pkg}}.ID(req.Id))
	}
	{{genCondition .pkg .fields}}
	return list
}
