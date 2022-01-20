package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hope/api/novel/socialuser/v1"
	"hope/apps/novel/internal/data/ent"
)

type SocialUserRepo interface {
	CreateSocialUser(context.Context, *v1.SocialUserCreateReq) (*ent.SocialUser, error)
	DeleteSocialUser(context.Context, *v1.SocialUserDeleteReq) error
	BatchDeleteSocialUser(context.Context, *v1.SocialUserBatchDeleteReq) (int, error)
	UpdateSocialUser(context.Context, *v1.SocialUserUpdateReq) (*ent.SocialUser, error)
	GetSocialUser(context.Context, *v1.SocialUserReq) (*ent.SocialUser, error)
	PageSocialUser(context.Context, *v1.SocialUserPageReq) ([]*ent.SocialUser, error)
}

type SocialUserUseCase struct {
	repo SocialUserRepo
	log  *log.Helper
}

func NewSocialUserUseCase(repo SocialUserRepo, logger log.Logger) *SocialUserUseCase {
	return &SocialUserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SocialUserUseCase) Create(ctx context.Context, req *v1.SocialUserCreateReq) (*ent.SocialUser, error) {
	return uc.repo.CreateSocialUser(ctx, req)
}
func (uc *SocialUserUseCase) Delete(ctx context.Context, req *v1.SocialUserDeleteReq) error {
	return uc.repo.DeleteSocialUser(ctx, req)
}
func (uc *SocialUserUseCase) BatchDelete(ctx context.Context, req *v1.SocialUserBatchDeleteReq) (int, error) {
	return uc.repo.BatchDeleteSocialUser(ctx, req)
}
func (uc *SocialUserUseCase) Update(ctx context.Context, req *v1.SocialUserUpdateReq) (*ent.SocialUser, error) {
	return uc.repo.UpdateSocialUser(ctx, req)
}
func (uc *SocialUserUseCase) Get(ctx context.Context, req *v1.SocialUserReq) (*ent.SocialUser, error) {
	return uc.repo.GetSocialUser(ctx, req)
}
func (uc *SocialUserUseCase) Page(ctx context.Context, req *v1.SocialUserPageReq) ([]*ent.SocialUser, error) {
	return uc.repo.PageSocialUser(ctx, req)
}
