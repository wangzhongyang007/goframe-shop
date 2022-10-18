package middleware

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/library/response"
)

const (
	CtxAccountId      = "account_id"       //token获取
	CtxAccountName    = "account_name"     //token获取
	CtxAccountAvatar  = "account_avatar"   //token获取
	CtxAccountSex     = "account_sex"      //token获取
	CtxAccountStatus  = "account_status"   //token获取
	CtxAccountSign    = "account_sign"     //token获取
	CtxAccountIsAdmin = "account_is_admin" //token获取
	CtxAccountRoleIds = "account_role_ids" //token获取
)

type TokenInfo struct {
	Id      int
	Name    string
	Avatar  string
	Sex     int
	Status  int
	Sign    string
	RoleIds string
	IsAdmin int
}

var GToken *gtoken.GfToken

var MiddlewareGToken = tokenMiddleware{}

type tokenMiddleware struct{}

func (s *tokenMiddleware) GetToken(r *ghttp.Request) {
	var tokenInfo TokenInfo
	token := GToken.GetTokenData(r)
	err := gconv.Struct(token.GetString("data"), &tokenInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//账号被冻结拉黑
	if tokenInfo.Status == 2 {
		response.AuthBlack(r)
		return
	}
	r.SetCtxVar(CtxAccountId, tokenInfo.Id)
	r.SetCtxVar(CtxAccountName, tokenInfo.Name)
	r.SetCtxVar(CtxAccountAvatar, tokenInfo.Avatar)
	r.SetCtxVar(CtxAccountSex, tokenInfo.Sex)
	r.SetCtxVar(CtxAccountStatus, tokenInfo.Status)
	r.SetCtxVar(CtxAccountSign, tokenInfo.Sign)
	r.SetCtxVar(CtxAccountRoleIds, tokenInfo.RoleIds)
	r.SetCtxVar(CtxAccountIsAdmin, tokenInfo.Sign)
	r.Middleware.Next()
}
