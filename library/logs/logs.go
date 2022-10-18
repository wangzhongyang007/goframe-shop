package logs

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type RequestLogReq struct {
	Path         string //存储路径
	RequestURI   string //请求URI
	RequestID    string //请求RequestID
	Method       string //请求方法
	Params       string //请求参数
	Response     string //响应参数
	Err          error  //错误信息
	ServerName   string //请求服务名称
	ResponseTime int64  //响应时间 毫秒级
}

const FormatErr = "参数【%v】错误【%v】响应时间【%v】"
const FormatSuc = "参数【%v】响应【%v】响应时间【%v】"

//统一请求日志
func RequestLog(req RequestLogReq) {

	Info(req.Path, "请求ID:【%v】 服务名称: 【%v】 请求路径:【%v】 请求方法: 【%v】 请求参数: 【%v】 响应参数: 【%v】 响应时间:【%v ms】error:【%v】",
		req.RequestID, req.ServerName, req.RequestURI, req.Method, req.Params, req.Response, req.ResponseTime, nil)

	if req.Err != nil {
		Error("%+v", gerror.Wrap(req.Err, req.RequestID))
	}

}

//记录info日志
func Info(path string, format string, v ...interface{}) {
	g.Log().Async().Cat(path).Infof(format, v...)
}

//记录error日志
func Error(format string, v ...interface{}) {
	g.Log().Async().Cat("error").Infof(format, v...)
}

//检查错误
func CheckErr(err error, extra string) bool {
	if err != nil {
		Error("%+v", gerror.Wrap(err, extra))
		return true
	}
	return false
}

func Infof(ctx context.Context, path string, format string, v ...interface{}) {
	g.Log().Ctx(ctx).Cat(path).Infof(format, v...)
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	g.Log().Ctx(ctx).Cat("error").Infof(format, v...)
}
