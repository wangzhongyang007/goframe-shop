package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/middleware"
	"shop/utility"
)

var service = loginService{}

type loginService struct {
}

/*
	退出登录
*/
func (a *loginService) Logout(r *ghttp.Request, req *LogoutReq) (err error) {
	//token := r.Cookie.Get(_cookie_sso)
	//r.Cookie.SetCookie(_cookie_sso, "", "stbz.top", "/", -1)
	r.Response.RedirectTo(req.Redirect, http.StatusFound)
	return
}

func (s *loginService) Login(ctx context.Context, req *LoginReq) (tokenInfo *middleware.TokenInfo, err error) {
	//验证账号密码是否正确
	adminInfo := model.AdminInfo{}
	err = dao.AdminInfo.Ctx(ctx).Where("name", req.Name).Scan(&adminInfo)
	if err != nil {
		return nil, err
	}
	if utility.EncryptPassword(req.PassWord, adminInfo.UserSalt) != adminInfo.Password {
		return nil, gerror.New("账号或者密码不正确")
	}
	tokenInfo = &middleware.TokenInfo{
		Id:      adminInfo.Id,
		Name:    adminInfo.Name,
		RoleIds: adminInfo.RoleIds,
		IsAdmin: adminInfo.IsAdmin,
	}
	return
}
