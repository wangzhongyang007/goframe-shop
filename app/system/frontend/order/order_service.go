package order

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
	"shop/app/shared"
)

var service = new(orderService)

type orderService struct {
}

func (s *orderService) Add(r *ghttp.Request, req *AddOrderReq) (res sql.Result, err error) {
	req.OrderInfo.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	req.OrderInfo.Number = shared.GetOrderNum()

	tx, err := g.DB().Begin()
	if err != nil {
		return nil, errors.New("启动事务失败")
	}

	//defer方法最后执行 如果有报错则回滚 如果没有报错，则提交事务
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	//生成主订单
	lastInsertId, err := dao.OrderInfo.Ctx(r.GetCtx()).TX(tx).InsertAndGetId(req.OrderInfo)
	if err != nil {
		return nil, err
	}
	//生成商品订单
	for _, info := range req.OrderGoodsInfos {
		info.OrderId = gconv.Int(lastInsertId)
		_, err := dao.OrderGoodsInfo.Ctx(r.GetCtx()).TX(tx).Insert(info)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (s *orderService) Update(r *ghttp.Request, req *UpdateOrderReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.OrderInfo.Ctx(r.GetCtx()).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *orderService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.OrderInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *orderService) List(r *ghttp.Request, req *PageListReq) (res ListOrderRes, err error) {
	whereCondition := g.Map{}
	if req.Status == 0 {
		whereCondition = g.Map{
			dao.OrderInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
		}
	} else {
		whereCondition = g.Map{
			dao.OrderInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
			dao.OrderInfo.Columns.Status: req.Status,
		}
	}
	count, err := dao.OrderInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.OrderInfo.Ctx(r.GetCtx()).With(OrderGoodsInfo{}).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func (s *orderService) Detail(ctx context.Context, req *DetailReq) (res ListOrderSql, err error) {
	err = dao.OrderInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
