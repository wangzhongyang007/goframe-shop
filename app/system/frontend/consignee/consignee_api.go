package consignee

import (
	"github.com/gogf/gf/net/ghttp"
	"shop/library/response"
)

var Consignee = consigneeApi{}

type consigneeApi struct{}

func (*consigneeApi) Add(r *ghttp.Request) {
	var req *AddConsigneeReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Add(r, req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*consigneeApi) Update(r *ghttp.Request) {
	var req *UpdateConsigneeReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Update(r, req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*consigneeApi) Delete(r *ghttp.Request) {
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

func (*consigneeApi) List(r *ghttp.Request) {
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
//func (*consigneeApi) Category(r *ghttp.Request) {
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

//func (*consigneeApi) List(r *ghttp.Request) {
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

func (*consigneeApi) Detail(r *ghttp.Request) {
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
