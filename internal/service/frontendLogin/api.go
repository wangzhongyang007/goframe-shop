package frontendLogin

import (
	"errors"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/middleware"
	"shop/utility/response"
	"strconv"
	"time"
)

var FrontendLogin = new(frontendLogin)

type frontendLogin struct {
}

//注册
func (s *frontendLogin) Register(r *ghttp.Request) {
	var data *RegisterReq
	if err := r.Parse(&data); err != nil {
		response.ParamErr(r, err)
	}
	err := service.Register(r.Context(), data)
	if err != nil {
		response.JsonExit(r, 0, "注册失败")
	} else {
		response.SuccessWithData(r, nil)
	}
}

//重置密码
func (s *frontendLogin) UpdatePassword(r *ghttp.Request) {
	var data *UpdatePasswordReq
	if err := r.Parse(&data); err != nil {
		response.ParamErr(r, err)
	}
	err := service.UpdatePasswordReq(r, data)
	if err != nil {
		response.JsonExit(r, 0, "重置密码失败")
	} else {
		response.SuccessWithData(r, nil)
	}
}

func (s *frontendLogin) Login(r *ghttp.Request) (string, interface{}) {
	var data *LoginReq
	if err := r.Parse(&data); err != nil {
		response.ParamErr(r, err)
	}
	token, err := service.Login(r.Context(), data)
	if err != nil {
		response.JsonExit(r, 3, err.Error())
		return "", nil
	}
	//前端token加前缀 和后端管理员登录做区分
	return "frontend" + gconv.String(token.Id), token
}

func (a *frontendLogin) LoginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		err := r.Response.WriteJson(respData)
		if err != nil {
			return
		}
		return
	} else {
		respData.Code = 1
		data := &LoginRes{
			Type:     "Bearer",
			Token:    respData.GetString("token"),
			ExpireIn: 10 * 24 * 60 * 60, //单位秒,
		}
		response.SuccessWithData(r, data)
	}
	return
}

func (a *frontendLogin) AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {

	if respData.Success() {
		r.Middleware.Next()
	} else {
		var params map[string]interface{}
		params = r.GetMap()
		no := gconv.String(gtime.TimestampMilli())

		glog.Info(fmt.Sprintf("[AUTH_%s][url:%s][params:%s][data:%s]",
			no, r.URL.Path, params, respData.Json()))
		response.JsonExit(r, 999, "请求错误或登录超时")
		return

	}
	return
}

//todo 退出登录有问题 没起作用
func (a *frontendLogin) Logout(r *ghttp.Request, respData gtoken.Resp) {
	cacheKey := middleware.GToken.CacheKey + gconv.String(r.GetCtxVar(middleware.CtxAccountId))
	g.Dump(cacheKey)
	var err error
	switch middleware.GToken.CacheMode {
	case 1:
		_, err2 := gcache.Remove(cacheKey)
		if err2 != nil {
			glog.Error("[GToken]cache remove error", err2)
			return
		}
	case 2:
		_, err = g.Redis().Do("DEL", cacheKey)
		if err != nil {
			glog.Error("[GToken]cache remove error", err)
			return
		}
	default:
		err = errors.New("cache model error")
	}

	if err != nil {
		response.JsonExit(r, 0, "缓存异常，请检查缓存")
	} else {
		response.SuccessWithData(r, nil)
	}
}

func GetSecondsToTomorrow() string {
	nowTime := time.Now()
	// 当天秒级时间戳
	nowTimeStamp := nowTime.Unix()

	nowTimeStr := nowTime.Format("2006-01-02")

	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t2, _ := time.ParseInLocation("2006-01-02", nowTimeStr, time.Local)
	// 第二天零点时间戳
	towTimeStamp := t2.AddDate(0, 0, 1).Unix()
	return strconv.FormatInt(towTimeStamp-nowTimeStamp, 10)
}

//封装清空缓存接口
func delGTokenCache(cacheKey string) (err error) {
	switch middleware.GToken.CacheMode {
	case 1:
		_, err = gcache.Remove(cacheKey)
	case 2:

		_, err = g.Redis().Do("DEL", cacheKey)
		if err != nil {
			glog.Error("[GToken]cache remove error", err)
			return
		}
	default:
		err = errors.New("cache model error")
	}
	return
}
