package admin

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"shop/library/response"
)

var Admin = rotationApi{}

type rotationApi struct{}

func (*rotationApi) Add(r *ghttp.Request) {
	var req *AddAdminReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Add(r.Context(), req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*rotationApi) Update(r *ghttp.Request) {
	var req *UpdateAdminReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Update(r.Context(), req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

//修改我的密码
func (*rotationApi) UpdateMyPassword(r *ghttp.Request) {
	var req *UpdateMyPasswordReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.UpdateMyPassword(r, req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*rotationApi) Delete(r *ghttp.Request) {
	var req *SoftDeleteReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Delete(r.Context(), req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*rotationApi) List(r *ghttp.Request) {
	var req *PageListReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	res, err := service.List(r.Context(), req)
	if err != nil {
		response.FailureWithData(r, 0, err, "")
	}
	response.SuccessWithData(r, res)
}
