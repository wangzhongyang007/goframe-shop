package permission

import (
	"github.com/gogf/gf/net/ghttp"
	"shop/library/response"
)

var Permission = permissionApi{}

type permissionApi struct{}

func (*permissionApi) Add(r *ghttp.Request) {
	var req *AddPermissionReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}
	if res, err := service.Add(r.Context(), req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*permissionApi) Update(r *ghttp.Request) {
	var req *UpdatePermissionReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Update(r.Context(), req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*permissionApi) Delete(r *ghttp.Request) {
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

func (*permissionApi) List(r *ghttp.Request) {
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
