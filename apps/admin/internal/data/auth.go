package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	v1 "hope/api/admin/auth/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/data/ent/sysuser"
	"hope/pkg/auth"
	"time"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(
	data *Data,
	logger log.Logger,
) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Login 登陆
func (r *authRepo) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	username := req.Username
	user, err := r.data.db.SysUser.Query().Where(sysuser.Username(username)).First(ctx)
	if err != nil {
		return nil, err
	}
	ok, err := auth.ValidatePassword(user.Password, req.Password)
	if err != nil || !ok {
		return nil, errors.New("invalid username or password")
	}
	roles, err := user.QueryRole().All(ctx)
	if err != nil {
		return nil, err
	}
	claims := &auth.Claims{
		UserId:   user.ID,
		Avatar:   user.Avatar,
		TenantId: 1,
	}
	claims.ExpiresAt = time.Now().Add(72 * time.Hour).UnixNano()
	claims.Id = cast.ToString(user.ID)
	claims.Subject = user.Username
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(auth.JwtKey))
	if err != nil {
		return nil, err
	}
	reply := &v1.LoginReply{Code: 200, Message: "登陆成功", Result: &v1.LoginReply_UserInfo{
		UserId:   user.ID,
		Username: user.Username,
		RealName: user.NickName,
		Avatar:   user.Avatar,
		Desc:     user.Desc,
		Token:    token,
		HomePath: user.HomePath,
	}}
	rs := make([]*v1.LoginReply_Role, 0)
	for _, role := range roles {
		rs = append(rs, &v1.LoginReply_Role{
			RoleName: role.RoleName,
			Value:    role.RoleKey,
		})
	}
	reply.Result.Roles = rs
	return reply, nil

}

// Logout 退出
func (r *authRepo) Logout(ctx context.Context, req *v1.LogOutReq) error {
	return nil
}
