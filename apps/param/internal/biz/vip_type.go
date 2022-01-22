package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/param/viptype/v1"
	"hope/apps/param/internal/data/ent"
)

type VipTypeRepo interface {
	CreateVipType(context.Context, *v1.VipTypeCreateReq) (*ent.VipType, error)
	DeleteVipType(context.Context, *v1.VipTypeDeleteReq) error
	BatchDeleteVipType(context.Context, *v1.VipTypeBatchDeleteReq) (int, error)
	UpdateVipType(context.Context, *v1.VipTypeUpdateReq) (*ent.VipType, error)
	GetVipType(context.Context, *v1.VipTypeReq) (*ent.VipType, error)
	PageVipType(context.Context, *v1.VipTypePageReq) ([]*ent.VipType, error)
}

type VipTypeUseCase struct {
	repo VipTypeRepo
	log  *log.Helper
}

func NewVipTypeUseCase(repo VipTypeRepo, logger log.Logger) *VipTypeUseCase {
	return &VipTypeUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *VipTypeUseCase) Create(ctx context.Context, req *v1.VipTypeCreateReq) (*ent.VipType, error) {
	return uc.repo.CreateVipType(ctx, req)
}
func (uc *VipTypeUseCase) Delete(ctx context.Context, req *v1.VipTypeDeleteReq) error {
	return uc.repo.DeleteVipType(ctx, req)
}
func (uc *VipTypeUseCase) BatchDelete(ctx context.Context, req *v1.VipTypeBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteVipType(ctx, req)
}
func (uc *VipTypeUseCase) Update(ctx context.Context, req *v1.VipTypeUpdateReq) (*ent.VipType, error) {
	return uc.repo.UpdateVipType(ctx, req)
}
func (uc *VipTypeUseCase) Get(ctx context.Context, req *v1.VipTypeReq) (*ent.VipType, error) {
	return uc.repo.GetVipType(ctx, req)
}
func (uc *VipTypeUseCase) Page(ctx context.Context, req *v1.VipTypePageReq) ([]*ent.VipType, error) {
	return uc.repo.PageVipType(ctx, req)
}
