package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/vipuser/v1"
	"hope/apps/novel/internal/data/ent"
)

type VipUserRepo interface {
	CreateVipUser(context.Context, *v1.VipUserCreateReq) (*ent.VipUser, error)
	DeleteVipUser(context.Context, *v1.VipUserDeleteReq) error
	BatchDeleteVipUser(context.Context, *v1.VipUserBatchDeleteReq) (int, error)
	UpdateVipUser(context.Context, *v1.VipUserUpdateReq) (*ent.VipUser, error)
	GetVipUser(context.Context, *v1.VipUserReq) (*ent.VipUser, error)
	PageVipUser(context.Context, *v1.VipUserPageReq) ([]*ent.VipUser, error)
}

type VipUserUseCase struct {
	repo VipUserRepo
	log  *log.Helper
}

func NewVipUserUseCase(repo VipUserRepo, logger log.Logger) *VipUserUseCase {
	return &VipUserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *VipUserUseCase) Create(ctx context.Context, req *v1.VipUserCreateReq) (*ent.VipUser, error) {
	return uc.repo.CreateVipUser(ctx, req)
}
func (uc *VipUserUseCase) Delete(ctx context.Context, req *v1.VipUserDeleteReq) error {
	return uc.repo.DeleteVipUser(ctx, req)
}
func (uc *VipUserUseCase) BatchDelete(ctx context.Context, req *v1.VipUserBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteVipUser(ctx, req)
}
func (uc *VipUserUseCase) Update(ctx context.Context, req *v1.VipUserUpdateReq) (*ent.VipUser, error) {
	return uc.repo.UpdateVipUser(ctx, req)
}
func (uc *VipUserUseCase) Get(ctx context.Context, req *v1.VipUserReq) (*ent.VipUser, error) {
	return uc.repo.GetVipUser(ctx, req)
}
func (uc *VipUserUseCase) Page(ctx context.Context, req *v1.VipUserPageReq) ([]*ent.VipUser, error) {
	return uc.repo.PageVipUser(ctx, req)
}
