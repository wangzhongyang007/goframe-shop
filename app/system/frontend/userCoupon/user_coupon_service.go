package userCoupon

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
)

var service = new(couponService)

type couponService struct {
}

func (s *couponService) Add(r *ghttp.Request, req *AddCouponReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.UserCouponInfo.Ctx(r.GetCtx()).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *couponService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.UserCouponInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *couponService) List(r *ghttp.Request, req *PageListReq) (res ListCouponRes, err error) {
	whereCondition := g.Map{
		dao.UserCouponInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
	}
	count, err := dao.UserCouponInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.UserCouponInfo.Ctx(r.GetCtx()).With(CouponInfo{}).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}
