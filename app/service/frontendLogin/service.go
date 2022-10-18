package frontendLogin

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
	"net/http"
	"shop/app/dao"
	"shop/app/middleware"
	"shop/app/model"
	"shop/library"
)

var service = frontendLoginService{}

type frontendLoginService struct {
}

/*
	退出登录
*/
func (a *frontendLoginService) Logout(r *ghttp.Request, req *LogoutReq) (err error) {
	//token := r.Cookie.Get(_cookie_sso)
	//r.Cookie.SetCookie(_cookie_sso, "", "stbz.top", "/", -1)
	r.Response.RedirectTo(req.Redirect, http.StatusFound)
	return
}

func (s *frontendLoginService) Login(ctx context.Context, req *LoginReq) (tokenInfo *middleware.TokenInfo, err error) {
	//验证账号密码是否正确
	userInfo := model.UserInfo{}
	err = dao.UserInfo.Ctx(ctx).Where("name", req.Name).Scan(&userInfo)
	if err != nil {
		return nil, err
	}
	if library.EncryptPassword(req.PassWord, userInfo.UserSalt) != userInfo.Password {
		return nil, gerror.New("账号或者密码不正确")
	}
	tokenInfo = &middleware.TokenInfo{
		Id:     userInfo.Id,
		Name:   userInfo.Name,
		Avatar: userInfo.Avatar,
		Sex:    userInfo.Sex,
		Status: userInfo.Status,
		Sign:   userInfo.Sign,
	}
	return
}

//注册
func (s *frontendLoginService) Register(ctx context.Context, req *RegisterReq) (err error) {
	//查询用户名是否存在
	count, err := dao.UserInfo.Ctx(ctx).Where("name", req.Name).Count()
	if err != nil || count > 0 {
		return gerror.New("用户名已存在，请换个用户名注册账号吧")
	}

	UserSalt := grand.S(10)
	req.PassWord = library.EncryptPassword(req.PassWord, UserSalt)
	req.UserSalt = UserSalt
	//添加新用户
	_, err = dao.UserInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return err
	}
	return
}

//重置密码
func (s *frontendLoginService) UpdatePasswordReq(r *ghttp.Request, req *UpdatePasswordReq) (err error) {
	//验证密保问题
	userInfo := model.UserInfo{}
	err = dao.UserInfo.Ctx(r.GetCtx()).WherePri(r.GetCtxVar(middleware.CtxAccountId)).Scan(&userInfo)
	if err != nil {
		g.Dump(err)
		return err
	}

	if userInfo.SecretAnswer != req.SecretAnswer {
		return errors.New("密保问题不正确")
	}

	UserSalt := grand.S(10)
	req.PassWord = library.EncryptPassword(req.PassWord, UserSalt)
	req.UserSalt = UserSalt
	//重置用户密码
	_, err = dao.UserInfo.Ctx(r.GetCtx()).WherePri(r.GetCtxVar(middleware.CtxAccountId)).Update(req)
	if err != nil {
		return err
	}
	return
}
