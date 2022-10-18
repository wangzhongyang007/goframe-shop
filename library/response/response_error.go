package response

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/library/logs"
)

type errResponse struct {
}

var Err = errResponse{}

//Service 服务异常
func (errResponse) Service() error {
	return gerror.NewCode(gcode.New(1000, "", ""))
}

func (errResponse) Common(code int, params ...interface{}) error {
	var msg string
	var res interface{}
	if len(params) > 0 {
		msg = gconv.String(params[0])
	}
	if len(params) > 1 {
		res = params[1]
	}
	return gerror.NewCode(gcode.New(code, msg, res))
}

func (errResponse) Param(err error) error {
	return gerror.NewCode(gcode.New(3000, gerror.Current(err).Error(), ""))
}

//Service 服务异常
func (errResponse) ServiceErr(err error) error {
	//记录服务异常日志
	logs.CheckErr(err, "ServiceErr")
	return gerror.NewCode(gcode.New(1000, "", ""))
}
