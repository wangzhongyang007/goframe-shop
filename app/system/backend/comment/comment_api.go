package comment

import (
	"github.com/gogf/gf/net/ghttp"
	"shop/library/response"
)

var Comment = commentApi{}

type commentApi struct{}

func (*commentApi) Add(r *ghttp.Request) {
	var req *AddCommentReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Add(r, req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

//func (*commentApi) Update(r *ghttp.Request) {
//	var req *UpdateCommentReq
//	if err := r.Parse(&req); err != nil {
//		response.ParamErr(r, err)
//	}
//
//	if res, err := service.Update(r, req); err != nil {
//		response.Code(r, err)
//	} else {
//		response.SuccessWithData(r, res)
//	}
//}

func (*commentApi) Delete(r *ghttp.Request) {
	var req *DeleteReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Delete(r.Context(), req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*commentApi) List(r *ghttp.Request) {
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
//func (*commentApi) Category(r *ghttp.Request) {
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

//func (*commentApi) List(r *ghttp.Request) {
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

func (*commentApi) Detail(r *ghttp.Request) {
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
