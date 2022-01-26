package data

import (
	"context"
	"errors"
	mapset "github.com/deckarep/golang-set"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	v1 "hope/api/admin/auth/v1"
	"hope/apps/admin/internal/biz"
	"hope/apps/admin/internal/data/ent"
	"hope/apps/admin/internal/data/ent/sysmenu"
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
	reply := &v1.LoginReply{Code: 200, Message: "登陆成功", Result: &v1.LoginReply_UserInfo{
		UserId:   user.ID,
		Username: user.Username,
		RealName: user.NickName,
		Avatar:   user.Avatar,
		Desc:     user.Desc,
		HomePath: user.HomePath,
	}}
	rs := make([]*v1.LoginReply_Role, 0)
	for _, role := range roles {
		rs = append(rs, &v1.LoginReply_Role{
			RoleName: role.RoleName,
			Value:    role.RoleKey,
		})
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
	reply.Result.Token = token
	reply.Result.Roles = rs
	return reply, nil

}

// Logout 退出
func (r *authRepo) Logout(ctx context.Context, req *v1.LogOutReq) error {
	return nil
}

func (r *authRepo) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (*v1.LoginReply, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	user, err := r.data.db.SysUser.Query().Where(sysuser.ID(claims.UserId)).Only(ctx)
	if err != nil {
		return nil, err
	}
	roles, err := user.QueryRole().All(ctx)
	if err != nil {
		return nil, err
	}
	reply := &v1.LoginReply{Code: 200, Message: "登陆成功", Result: &v1.LoginReply_UserInfo{
		UserId:   user.ID,
		Username: user.Username,
		RealName: user.NickName,
		Avatar:   user.Avatar,
		Desc:     user.Desc,
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
	return reply, err
}

func (r *authRepo) GetPermCode(ctx context.Context, req *v1.GetPermReq) (*v1.GetPermReply, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	user, err := r.data.db.SysUser.Query().Where(sysuser.ID(claims.UserId)).Only(ctx)
	if err != nil {
		return nil, err
	}
	set := mapset.NewSet()
	if user.ID == 1 {
		//超级管理员
		set.Add("*")
	} else {
		roles, err := user.QueryRole().All(ctx)
		if err != nil {
			return nil, err
		}
		for _, role := range roles {
			menus, err := role.QueryMenus().All(ctx)
			if err != nil {
				continue
			}
			for _, menu := range menus {
				set.Add(menu.Permission)
			}
		}
	}
	reply := &v1.GetPermReply{
		Code:    200,
		Message: "成功",
		Result:  cast.ToStringSlice(set.ToSlice()),
	}
	return reply, nil
}

func (r *authRepo) GetMenuList(ctx context.Context, req *v1.GetMenuReq) (*v1.GetMenuReply, error) {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	user, err := r.data.db.SysUser.Query().Where(sysuser.ID(claims.UserId)).Only(ctx)
	if err != nil {
		return nil, err
	}

	var menus []*ent.SysMenu
	if user.ID == 1 {
		//超级管理员
		menus, err = r.data.db.SysMenu.Query().Where(sysmenu.StateEQ(sysmenu.StateU), sysmenu.MenuTypeIn("D", "M")).All(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		set := mapset.NewSet()
		roles, err := user.QueryRole().All(ctx)
		if err != nil {
			return nil, err
		}
		for _, role := range roles {
			sysMenus, err := role.QueryMenus().All(ctx)
			if err != nil {
				continue
			}
			for _, menu := range sysMenus {
				set.Add(menu)
			}
		}
		slice := set.ToSlice()
		for _, menu := range slice {
			sysMenu, ok := menu.(*ent.SysMenu)
			if ok {
				menus = append(menus, sysMenu)
			}
		}
	}
	var retMenus []*v1.GetMenuReply_Menu
	for _, menu := range menus {
		retMenus = append(retMenus, &v1.GetMenuReply_Menu{
			Id:        menu.ID,
			ParentId:  menu.ParentId,
			Path:      menu.Path,
			Name:      menu.Name,
			Component: menu.Component,
			Redirect:  menu.Redirect,
			Meta: &v1.GetMenuReply_Menu_Meta{
				Title:              menu.Title,
				Icon:               menu.Icon,
				IgnoreKeepAlive:    menu.IgnoreKeepAlive,
				HideBreadcrumb:     menu.HideBreadcrumb,
				HideChildrenInMenu: menu.HideChildrenInMenu,
				Sort:               menu.Sort,
				HideMenu:           menu.HideMenu,
				FrameSrc:           menu.FrameSrc,
			},
		})

	}
	reply := &v1.GetMenuReply{
		Code:    200,
		Message: "成功",
		Result:  retMenus,
	}
	return reply, nil
}

/*
func MenuTree(menuList []*ent.SysMenu) {
	//转成IDMap
	idMap := make(map[int64]*ent.SysMenu)
	//根据父ID分组
	pIdMap := make(map[int64][]*ent.SysMenu)
	for _, menu := range menuList {
		idMap[menu.ID]=menu
		pid := menu.ParentId
		childs,ok := pIdMap[pid]
		if !ok {
			childs=make([]*ent.SysMenu,0)
		}
		childs = append(childs, menu)
	}
}*/
