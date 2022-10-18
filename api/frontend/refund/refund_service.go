package refund

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/internal/dao"
	"shop/middleware"
	"shop/app/shared"
)

var service = new(refundService)

type refundService struct {
}

func (s *refundService) Add(r *ghttp.Request, req *AddRefundReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	req.Number = shared.GetRefundNum()
	res, err = dao.RefundInfo.Ctx(r.GetCtx()).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *refundService) Update(r *ghttp.Request, req *UpdateRefundReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.RefundInfo.Ctx(r.GetCtx()).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *refundService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.RefundInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *refundService) List(r *ghttp.Request, req *PageListReq) (res ListRefundRes, err error) {
	whereCondition := g.Map{}
	if req.Status == 0 {
		whereCondition = g.Map{
			dao.RefundInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
		}
	} else {
		whereCondition = g.Map{
			dao.RefundInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
			dao.RefundInfo.Columns.Status: req.Status,
		}
	}
	count, err := dao.RefundInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.RefundInfo.Ctx(r.GetCtx()).With(GoodsInfo{}).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func (s *refundService) Detail(ctx context.Context, req *DetailReq) (res ListRefundSql, err error) {
	err = dao.RefundInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
