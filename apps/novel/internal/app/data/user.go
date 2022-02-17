package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	v1 "hope/api/app/user/v1"
	"hope/apps/novel/internal/app/biz"
	"hope/apps/novel/internal/data/ent"
	"hope/apps/novel/internal/data/ent/socialuser"
	"hope/pkg/auth"
	"time"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r userRepo) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	var user *ent.SocialUser
	var err error
	loginType := req.LoginType
	switch loginType {
	case v1.LoginType_Register:
		//注册

	case v1.LoginType_Normal:
		user, err = r.data.db.SocialUser.Query().Where(socialuser.UserName(req.UserName)).First(ctx)
		if err != nil {
			return nil, err
		}
		ok, err := auth.ValidatePassword(user.Password, req.Password)
		if err != nil || !ok {
			return nil, errors.New("invalid username or password")
		}
	case v1.LoginType_Phone:
		user, err = r.data.db.SocialUser.Query().Where(socialuser.Phone(req.UserName)).First(ctx)
		if err != nil {
			return nil, err
		}
	case v1.LoginType_WeChat,
		v1.LoginType_Alipay,
		v1.LoginType_Google,
		v1.LoginType_Facebook,
		v1.LoginType_Line:
		user, err = r.data.db.SocialUser.Query().Where(socialuser.Openid(req.OpenId)).First(ctx)
		if err != nil {
			return nil, err
		}
	}
	claims := &auth.Claims{
		UserId:   user.ID,
		Avatar:   user.Avatar,
		TenantId: user.TenantId,
	}
	claims.ExpiresAt = time.Now().Add(72 * time.Hour).UnixNano()
	claims.Id = cast.ToString(user.ID)
	claims.Subject = user.UserName
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(auth.JwtKey))
	if err != nil {
		return nil, err
	}
	reply := &v1.LoginReply{Code: 200, Message: "登陆成功", Token: token}
	return reply, nil
}

func (r userRepo) Info(ctx context.Context, req *v1.InfoReq) (*v1.InfoReply, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	userId := claims.UserId
	user, err := r.data.db.SocialUser.Query().Where(socialuser.ID(userId)).First(ctx)
	//TODO implement me
	return &v1.InfoReply{
		Code:    200,
		Message: "成功",
		Result: &v1.UserInfo{
			Id:       user.ID,
			UserName: user.UserName,
			Gender:   user.Language,
			Summary:  user.Remark,
			HeadImg:  user.Avatar,
		},
	}, nil
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
