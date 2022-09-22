package data

import (
	"github.com/gogf/gf/net/ghttp"
	"shop/library/response"
)

var Data = dataApi{}

type dataApi struct{}

//头部卡片
func (*dataApi) HeadCard(r *ghttp.Request) {
	res, err := service.HeadCard(r.Context())
	if err != nil {
		response.FailureWithData(r, 0, err, "")
	}
	response.SuccessWithData(r, res)
}

//ECharts
func (*dataApi) ECharts(r *ghttp.Request) {
	res, err := service.ECharts(r.Context())
	if err != nil {
		response.FailureWithData(r, 0, err, "")
	}
	response.SuccessWithData(r, res)
}
