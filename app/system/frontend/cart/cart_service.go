package cart

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
)

var service = new(cartService)

type cartService struct {
}

func (s *cartService) Add(r *ghttp.Request, req *AddCartReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.CartInfo.Ctx(r.GetCtx()).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *cartService) Update(r *ghttp.Request, req *UpdateCartReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.CartInfo.Ctx(r.GetCtx()).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *cartService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.CartInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *cartService) List(r *ghttp.Request, req *PageListReq) (res ListCartRes, err error) {
	whereCondition := g.Map{
		dao.CartInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
	}
	count, err := dao.CartInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.CartInfo.Ctx(r.GetCtx()).With(GoodsInfo{}).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func (s *cartService) Detail(ctx context.Context, req *DetailReq) (res ListCartSql, err error) {
	err = dao.CartInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
