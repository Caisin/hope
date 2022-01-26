package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "hope/api/novel/socialuser/v1"
	"hope/apps/novel/internal/biz"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"hope/pkg/auth"
	"hope/pkg/util/str"

	"hope/pkg/pagin"
	"time"
)

type socialUserRepo struct {
	data *Data
	log  *log.Helper
}

// NewSocialUserRepo .
func NewSocialUserRepo(data *Data, logger log.Logger) biz.SocialUserRepo {
	return &socialUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateSocialUser 创建
func (r *socialUserRepo) CreateSocialUser(ctx context.Context, req *v1.SocialUserCreateReq) (*ent.SocialUser, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return r.data.db.SocialUser.Create().
		SetChId(req.ChId).
		SetUnionid(req.Unionid).
		SetToken(req.Token).
		SetOpenid(req.Openid).
		SetRoutineOpenid(req.RoutineOpenid).
		SetUserName(req.UserName).
		SetNickName(req.NickName).
		SetBirthday(req.Birthday.AsTime()).
		SetPhone(req.Phone).
		SetEmail(req.Email).
		SetPassword(req.Password).
		SetAvatar(req.Avatar).
		SetSex(req.Sex).
		SetRegion(req.Region).
		SetCity(req.City).
		SetLanguage(req.Language).
		SetProvince(req.Province).
		SetCountry(req.Country).
		SetSignature(req.Signature).
		SetRemark(req.Remark).
		SetGroupid(req.Groupid).
		SetTagidList(req.TagidList).
		SetSubscribe(req.Subscribe).
		SetSubscribeTime(req.SubscribeTime).
		SetSessionKey(req.SessionKey).
		SetUserType(req.UserType).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetCreateBy(claims.UserId).
		SetTenantId(claims.TenantId).
		Save(ctx)

}

// DeleteSocialUser 删除
func (r *socialUserRepo) DeleteSocialUser(ctx context.Context, req *v1.SocialUserDeleteReq) error {
	return r.data.db.SocialUser.DeleteOneID(req.Id).Exec(ctx)
}

// BatchDeleteSocialUser 批量删除
func (r *socialUserRepo) BatchDeleteSocialUser(ctx context.Context, req *v1.SocialUserBatchDeleteReq) (int, error) {
	return r.data.db.SocialUser.Delete().Where(socialuser.IDIn(req.Ids...)).Exec(ctx)
}

// UpdateSocialUser 更新
func (r *socialUserRepo) UpdateSocialUser(ctx context.Context, req *v1.SocialUserUpdateReq) (*ent.SocialUser, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	return r.data.db.SocialUser.UpdateOneID(req.Id).
		SetChId(req.ChId).
		SetUnionid(req.Unionid).
		SetToken(req.Token).
		SetOpenid(req.Openid).
		SetRoutineOpenid(req.RoutineOpenid).
		SetUserName(req.UserName).
		SetNickName(req.NickName).
		SetBirthday(req.Birthday.AsTime()).
		SetPhone(req.Phone).
		SetEmail(req.Email).
		SetPassword(req.Password).
		SetAvatar(req.Avatar).
		SetSex(req.Sex).
		SetRegion(req.Region).
		SetCity(req.City).
		SetLanguage(req.Language).
		SetProvince(req.Province).
		SetCountry(req.Country).
		SetSignature(req.Signature).
		SetRemark(req.Remark).
		SetGroupid(req.Groupid).
		SetTagidList(req.TagidList).
		SetSubscribe(req.Subscribe).
		SetSubscribeTime(req.SubscribeTime).
		SetSessionKey(req.SessionKey).
		SetUserType(req.UserType).
		SetUpdateBy(claims.UserId).
		Save(ctx)
}

// GetSocialUser 根据Id查询
func (r *socialUserRepo) GetSocialUser(ctx context.Context, req *v1.SocialUserReq) (*ent.SocialUser, error) {
	return r.data.db.SocialUser.Get(ctx, req.Id)
}

// PageSocialUser 分页查询
func (r *socialUserRepo) PageSocialUser(ctx context.Context, req *v1.SocialUserPageReq) ([]*ent.SocialUser, error) {
	p := req.Pagin
	if p == nil {
		req.Pagin = &pagin.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := r.data.db.SocialUser.
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
	query.Limit(int(p.GetPageSize())).
		Offset(int(p.GetOffSet()))
	if p.NeedOrder() {
		if p.IsDesc() {
			query.Order(ent.Desc(p.GetField()))
		} else {
			query.Order(ent.Asc(p.GetField()))
		}
	}
	return query.All(ctx)
}

// genCondition 构造查询条件
func (r *socialUserRepo) genCondition(req *v1.SocialUserReq) []predicate.SocialUser {
	if req == nil {
		return nil
	}
	list := make([]predicate.SocialUser, 0)
	if req.Id > 0 {
		list = append(list, socialuser.ID(req.Id))
	}
	if req.ChId > 0 {
		list = append(list, socialuser.ChId(req.ChId))
	}
	if str.IsBlank(req.Unionid) {
		list = append(list, socialuser.UnionidContains(req.Unionid))
	}
	if str.IsBlank(req.Token) {
		list = append(list, socialuser.TokenContains(req.Token))
	}
	if str.IsBlank(req.Openid) {
		list = append(list, socialuser.OpenidContains(req.Openid))
	}
	if str.IsBlank(req.RoutineOpenid) {
		list = append(list, socialuser.RoutineOpenidContains(req.RoutineOpenid))
	}
	if str.IsBlank(req.UserName) {
		list = append(list, socialuser.UserNameContains(req.UserName))
	}
	if str.IsBlank(req.NickName) {
		list = append(list, socialuser.NickNameContains(req.NickName))
	}
	if req.Birthday.IsValid() && !req.Birthday.AsTime().IsZero() {
		list = append(list, socialuser.BirthdayGTE(req.Birthday.AsTime()))
	}
	if str.IsBlank(req.Phone) {
		list = append(list, socialuser.PhoneContains(req.Phone))
	}
	if str.IsBlank(req.Email) {
		list = append(list, socialuser.EmailContains(req.Email))
	}
	if str.IsBlank(req.Password) {
		list = append(list, socialuser.PasswordContains(req.Password))
	}
	if str.IsBlank(req.Avatar) {
		list = append(list, socialuser.AvatarContains(req.Avatar))
	}
	if req.Sex > 0 {
		list = append(list, socialuser.Sex(req.Sex))
	}
	if str.IsBlank(req.Region) {
		list = append(list, socialuser.RegionContains(req.Region))
	}
	if str.IsBlank(req.City) {
		list = append(list, socialuser.CityContains(req.City))
	}
	if str.IsBlank(req.Language) {
		list = append(list, socialuser.LanguageContains(req.Language))
	}
	if str.IsBlank(req.Province) {
		list = append(list, socialuser.ProvinceContains(req.Province))
	}
	if str.IsBlank(req.Country) {
		list = append(list, socialuser.CountryContains(req.Country))
	}
	if str.IsBlank(req.Signature) {
		list = append(list, socialuser.SignatureContains(req.Signature))
	}
	if str.IsBlank(req.Remark) {
		list = append(list, socialuser.RemarkContains(req.Remark))
	}
	if req.Groupid > 0 {
		list = append(list, socialuser.Groupid(req.Groupid))
	}
	if str.IsBlank(req.TagidList) {
		list = append(list, socialuser.TagidListContains(req.TagidList))
	}
	if req.Subscribe > 0 {
		list = append(list, socialuser.Subscribe(req.Subscribe))
	}
	if req.SubscribeTime > 0 {
		list = append(list, socialuser.SubscribeTime(req.SubscribeTime))
	}
	if str.IsBlank(req.SessionKey) {
		list = append(list, socialuser.SessionKeyContains(req.SessionKey))
	}
	if str.IsBlank(req.UserType) {
		list = append(list, socialuser.UserTypeContains(req.UserType))
	}
	if req.CreatedAt.IsValid() && !req.CreatedAt.AsTime().IsZero() {
		list = append(list, socialuser.CreatedAtGTE(req.CreatedAt.AsTime()))
	}
	if req.UpdatedAt.IsValid() && !req.UpdatedAt.AsTime().IsZero() {
		list = append(list, socialuser.UpdatedAtGTE(req.UpdatedAt.AsTime()))
	}
	if req.CreateBy > 0 {
		list = append(list, socialuser.CreateBy(req.CreateBy))
	}
	if req.UpdateBy > 0 {
		list = append(list, socialuser.UpdateBy(req.UpdateBy))
	}
	if req.TenantId > 0 {
		list = append(list, socialuser.TenantId(req.TenantId))
	}

	return list
}
