package response

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// JsonResponse 数据返回通用JSON数据结构
type JsonResponse struct {
	//ID       string      `json:"id"`                 //
	Code     int         `json:"code"`               // 错误码((1:成功, 0:失败, >1:错误码))
	Message  string      `json:"message"`            // 提示信息
	Data     interface{} `json:"data,omitempty"`     // 返回数据(业务接口定义具体数据结构)
	Redirect string      `json:"redirect,omitempty"` // 引导客户端跳转到指定路由
}

// JsonExit 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	codeMsgs := GetCodeMsg(code, msg)

	response := &JsonResponse{
		//ID:      r.GetCtxVar("RequestId").String(),
		Code:    code,
		Message: codeMsgs,
		Data:    responseData,
	}
	r.SetParam("apiReturnRes", response)
	_ = r.Response.WriteJsonExit(response)
}

func ApiResponse(r *ghttp.Request, res interface{}) {
	Map := gjson.New(res).Map()
	Map["id"] = r.GetCtxVar("request_id")
	r.Response.WriteJson(Map)
	r.Exit()
}

func Success(r *ghttp.Request) {

	res := dataReturn(r, 1)
	_ = r.Response.WriteJsonExit(res)
}

func Failure(r *ghttp.Request, code int, err error) {

	res := dataReturn(r, code, err.Error())
	_ = r.Response.WriteJsonExit(res)
}

func FailureCode(r *ghttp.Request, code int) {
	res := dataReturn(r, code)
	_ = r.Response.WriteJsonExit(res)
}

func dataReturn(r *ghttp.Request, code int, req ...interface{}) *JsonResponse {
	var msg string
	var data interface{}
	if len(req) > 0 {
		msg = gconv.String(req[0])
	}
	if len(req) > 1 {
		data = req[1]
	}
	//msg = GetCodeMsg(code, msg)
	if code != 1 && !gconv.Bool(r.GetCtxVar("api_code")) {
		code = 0
	}
	response := &JsonResponse{
		//ID:      r.GetCtxVar("RequestId").String(),
		Code:    code,
		Message: msg,
		Data:    data,
	}
	r.SetParam("apiReturnRes", response)
	return response
}

func SuccessWithData(r *ghttp.Request, data interface{}) {
	res := dataReturn(r, 1, "ok", data)
	_ = r.Response.WriteJsonExit(res)
}

func FailureWithData(r *ghttp.Request, code int, err error, data interface{}) {
	res := dataReturn(r, code, err.Error(), data)
	_ = r.Response.WriteJsonExit(res)
}

//ParamErr 参数错误

func ParamErr(r *ghttp.Request, err error) {
	res := dataReturn(r, 3000, gerror.Current(err))
	_ = r.Response.WriteJsonExit(res)
}

//NotOnline 应用未上线
func NotOnline(r *ghttp.Request) {
	res := dataReturn(r, 3002)
	_ = r.Response.WriteJsonExit(res)
}

func BackErr(r *ghttp.Request, err error) {
	res := dataReturn(r, 10002, err.Error())
	_ = r.Response.WriteJsonExit(res)
}

func Number(r *ghttp.Request) {
	res := dataReturn(r, 10004)
	_ = r.Response.WriteJsonExit(res)
}

func Code(r *ghttp.Request, err error) {
	//res := dataReturn(r, gerror.Code(err).Code(), gerror.Code(err).Message(), gerror.Code(err).Detail())
	res := dataReturn(r, gerror.Code(err).Code(), err.Error(), gerror.Code(err).Detail())
	_ = r.Response.WriteJsonExit(res)
}

//Fatal 致命错误
func Fatal(r *ghttp.Request) {
	res := dataReturn(r, 500)
	_ = r.Response.WriteJsonExit(res)
}

//Repeat 重复请求
func Repeat(r *ghttp.Request) {
	res := dataReturn(r, 3001)
	_ = r.Response.WriteJsonExit(res)
}

//NotFound 接口不存在
func NotFound(r *ghttp.Request) {
	res := dataReturn(r, 404)
	_ = r.Response.WriteJsonExit(res)
}

//Sign 签名错误
func Sign(r *ghttp.Request, err error) {
	res := dataReturn(r, 99999, err.Error())
	_ = r.Response.WriteJsonExit(res)
}

//Auth 认证失败
func Auth(r *ghttp.Request) {
	res := dataReturn(r, 999, "请登录")
	_ = r.Response.WriteJsonExit(res)
}

//Auth 认证失败 被冻结拉黑
func AuthBlack(r *ghttp.Request) {
	res := dataReturn(r, 888, "您的账号被冻结拉黑，请联系管理员")
	_ = r.Response.WriteJsonExit(res)
}

//Auth 认证失败
func NoPower(r *ghttp.Request) {
	response := &JsonResponse{
		//ID:      r.GetCtxVar("RequestId").String(),
		Code:    5002,
		Message: "没有权限",
	}
	r.SetParam("apiReturnRes", response)
	_ = r.Response.WriteJsonExit(response)
}

//子账号登陆open
func SubAccountLoginOpenStatus(r *ghttp.Request) {
	response := &JsonResponse{
		//ID:      r.GetCtxVar("RequestId").String(),
		Code:    5002,
		Message: "账户未审核",
	}
	r.SetParam("apiReturnRes", response)
	_ = r.Response.WriteJsonExit(response)
}
