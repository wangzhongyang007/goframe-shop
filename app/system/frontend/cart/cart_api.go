package cart

import (
	"github.com/gogf/gf/net/ghttp"
	"shop/library/response"
)

var Cart = cartApi{}

type cartApi struct{}

func (*cartApi) Add(r *ghttp.Request) {
	var req *AddCartReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Add(r, req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*cartApi) Update(r *ghttp.Request) {
	var req *UpdateCartReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Update(r, req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*cartApi) Delete(r *ghttp.Request) {
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

func (*cartApi) List(r *ghttp.Request) {
	var req *PageListReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	res, err := service.List(r, req)
	if err != nil {
		response.FailureWithData(r, 0, err, "")
	}
	response.SuccessWithData(r, res)
}

//同类商品推荐
//func (*cartApi) Category(r *ghttp.Request) {
//	var req *CategoryPageListReq
//	if err := r.Parse(&req); err != nil {
//		response.ParamErr(r, err)
//	}
//
//	res, err := service.Category(r.Context(), req)
//	if err != nil {
//		response.FailureWithData(r, 0, err, "")
//	}
//	response.SuccessWithData(r, res)
//}

//func (*cartApi) List(r *ghttp.Request) {
//	var req *PageListReq
//	if err := r.Parse(&req); err != nil {
//		response.ParamErr(r, err)
//	}
//
//	res, err := service.List(r.Context(), req)
//	if err != nil {
//		response.FailureWithData(r, 0, err, "")
//	}
//	response.SuccessWithData(r, res)
//}

func (*cartApi) Detail(r *ghttp.Request) {
	var req *DetailReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}
	res, err := service.Detail(r.Context(), req)
	if err != nil {
		response.FailureWithData(r, 0, err, "")
	}
	response.SuccessWithData(r, res)
}
