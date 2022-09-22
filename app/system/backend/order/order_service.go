package order

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
	"shop/app/shared"
)

var service = new(orderService)

type orderService struct {
}

func (s *orderService) Add(r *ghttp.Request, req *AddOrderReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	req.Number = shared.GetOrderNum()
	res, err = dao.OrderInfo.Ctx(r.GetCtx()).Insert(req)
	if err != nil {
		return nil, err
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
	//实例化map
	whereCondition := gmap.New()
	//很好的理解了map是引用类型的特点 在这个函数中为查询条件赋值
	packListCondition(req, whereCondition)
	count, err := dao.OrderInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.OrderInfo.Ctx(r.GetCtx()).Where(whereCondition).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func packListCondition(req *PageListReq, whereCondition *gmap.Map) {
	//使用map支持set的特性 避免在声明的时候赋值，那么写需要做的判断太复杂了。
	if req.UserId != 0 {
		whereCondition.Set(dao.OrderInfo.Columns.UserId, req.UserId)
	}
	if req.Status != 0 {
		whereCondition.Set(dao.OrderInfo.Columns.Status, req.Status)
	}
	if req.Date != "" {
		//日期范围查询
		whereCondition.Set(dao.UserInfo.Columns.CreatedAt+" >=", gtime.New(req.Date).StartOfDay())
		whereCondition.Set(dao.UserInfo.Columns.CreatedAt+" <=", gtime.New(req.Date).EndOfDay())
	}
	if req.Keyword != "" {
		whereCondition.Set(dao.OrderInfo.Columns.ConsigneeAddress+" like ", "%"+req.Keyword+"%")
	}
	if req.Number != "" {
		whereCondition.Set(dao.OrderInfo.Columns.Number+" like ", "%"+req.Number+"%")
	}
}

func (s *orderService) Detail(ctx context.Context, req *DetailReq) (res ListOrderSql, err error) {
	err = dao.OrderInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
