package article

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"shop/library/response"
)

var Article = rotationApi{}

type rotationApi struct{}

func (*rotationApi) Add(r *ghttp.Request) {
	var req *AddArticleReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}
	if res, err := service.Add(r, req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*rotationApi) Update(r *ghttp.Request) {
	var req *UpdateArticleReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Update(r, req); err != nil {
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

	if res, err := service.Delete(r, req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}
}

func (*rotationApi) Detail(r *ghttp.Request) {
	var req *DetailReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	if res, err := service.Detail(r, req); err != nil {
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

func (*rotationApi) MyList(r *ghttp.Request) {
	var req *PageListReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}

	res, err := service.MyList(r, req)
	if err != nil {
		response.FailureWithData(r, 0, err, "")
	}
	response.SuccessWithData(r, res)
}
