package coupon

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/net/ghttp"
	"shop/internal/dao"
)

var service = new(couponService)

type couponService struct {
}

func (s *couponService) Add(r *ghttp.Request, req *AddCouponReq) (res sql.Result, err error) {
	res, err = dao.CouponInfo.Ctx(r.GetCtx()).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *couponService) Update(ctx context.Context, req *UpdateCouponReq) (res sql.Result, err error) {
	res, err = dao.CouponInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *couponService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.CouponInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *couponService) List(r *ghttp.Request, req *PageListReq) (res ListCouponRes, err error) {
	count, err := dao.CouponInfo.Ctx(r.GetCtx()).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.CouponInfo.Ctx(r.GetCtx()).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}
