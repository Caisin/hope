package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/{{.model}}/{{.pkg}}/v1"
	"hope/apps/{{.model}}/internal/data/ent"
)

type {{.name}}Repo interface {
	Create{{.name}}(context.Context, *v1.{{.name}}CreateReq) (*ent.{{.name}}, error)
	Delete{{.name}}(context.Context, *v1.{{.name}}DeleteReq) error
	BatchDelete{{.name}}(context.Context, *v1.{{.name}}BatchDeleteReq) (int, error)
	Update{{.name}}(context.Context, *v1.{{.name}}UpdateReq) (*ent.{{.name}}, error)
	Get{{.name}}(context.Context, *v1.{{.name}}Req) (*ent.{{.name}}, error)
	Page{{.name}}(context.Context, *v1.{{.name}}PageReq) ([]*ent.{{.name}}, error)
}

type {{.name}}UseCase struct {
	repo {{.name}}Repo
	log  *log.Helper
}

func New{{.name}}UseCase(repo {{.name}}Repo, logger log.Logger) *{{.name}}UseCase {
	return &{{.name}}UseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *{{.name}}UseCase) Create(ctx context.Context, req *v1.{{.name}}CreateReq) (*ent.{{.name}}, error) {
	return uc.repo.Create{{.name}}(ctx, req)
}
func (uc *{{.name}}UseCase) Delete(ctx context.Context, req *v1.{{.name}}DeleteReq) error {
	return uc.repo.Delete{{.name}}(ctx, req)
}
func (uc *{{.name}}UseCase) BatchDelete(ctx context.Context, req *v1.{{.name}}BatchDeleteReq) (int, error) {
	return uc.repo.BatchDelete{{.name}}(ctx, req)
}
func (uc *{{.name}}UseCase) Update(ctx context.Context, req *v1.{{.name}}UpdateReq) (*ent.{{.name}}, error) {
	return uc.repo.Update{{.name}}(ctx, req)
}
func (uc *{{.name}}UseCase) Get(ctx context.Context, req *v1.{{.name}}Req) (*ent.{{.name}}, error) {
	return uc.repo.Get{{.name}}(ctx, req)
}
func (uc *{{.name}}UseCase) Page(ctx context.Context, req *v1.{{.name}}PageReq) ([]*ent.{{.name}}, error) {
	return uc.repo.Page{{.name}}(ctx, req)
}
