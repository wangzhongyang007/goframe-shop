package refund

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
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
	whereCondition := packCondition(req)
	count, err := dao.RefundInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.RefundInfo.Ctx(r.GetCtx()).Where(whereCondition).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func packCondition(req *PageListReq) (whereCondition g.Map) {
	if req.UserId != 0 && req.Status != 0 {
		whereCondition = g.Map{
			dao.RefundInfo.Columns.UserId: req.UserId,
			dao.RefundInfo.Columns.Status: req.Status,
		}
	} else if req.UserId != 0 {
		whereCondition = g.Map{
			dao.RefundInfo.Columns.UserId: req.UserId,
		}
	} else if req.Status != 0 {
		whereCondition = g.Map{
			dao.RefundInfo.Columns.Status: req.Status,
		}
	}
	return
}

//同类商品推荐
//func (s *refundService) Category(ctx context.Context, req *CategoryPageListReq) (res ListRefundRes, err error) {
//	//获取商品的分类
//	currentRefund := model.RefundInfo{}
//	err = dao.RefundInfo.Ctx(ctx).WherePri(req.Id).Scan(&currentRefund)
//	if err != nil {
//		return ListRefundRes{}, err
//	}
//
//	whereLevelCondition := g.Map{
//		"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{currentRefund.Level1CategoryId, currentRefund.Level2CategoryId, currentRefund.Level3CategoryId},
//	}
//	whereIdCondition := g.Map{
//		"id!=": req.Id,
//	}
//	count, err := dao.RefundInfo.Ctx(ctx).Where(whereIdCondition).Where(whereLevelCondition).Count()
//	if err != nil {
//		return
//	}
//	res.Count = count
//	err = dao.RefundInfo.Ctx(ctx).Where(whereIdCondition).Where(whereLevelCondition).Page(req.Page, req.Limit).Scan(&res.List)
//	if err != nil {
//		return
//	}
//	return
//}

//func (s *refundService) List(ctx context.Context, req *PageListReq) (res ListRefundRes, err error) {
//	whereCondition := g.Map{}
//	if req.Keyword != "" && req.CategoryId != 0 {
//		whereCondition = g.Map{
//			"name like": "%" + req.Keyword + "%",
//			"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{req.CategoryId, req.CategoryId, req.CategoryId},
//		}
//	} else if req.Keyword != "" {
//		whereCondition = g.Map{
//			"name like": "%" + req.Keyword + "%",
//		}
//	} else if req.CategoryId != 0 {
//		whereCondition = g.Map{
//			"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{req.CategoryId, req.CategoryId, req.CategoryId},
//		}
//	} else {
//		whereCondition = g.Map{}
//	}
//
//	//获取数量
//	count, err := dao.RefundInfo.Ctx(ctx).
//		Where(whereCondition).
//		Count()
//	if err != nil {
//		return
//	}
//	res.Count = count
//
//	//获取值
//	//排序规则
//	sortCondition := packSort(req)
//	err = dao.RefundInfo.Ctx(ctx).
//		Where(whereCondition).
//		Page(req.Page, req.Limit).
//		Refund(sortCondition).
//		Scan(&res.List)
//	if err != nil {
//		return
//	}
//	return
//}

//封装排序方法
//func packSort(req *SearchPageListReq) (sortCondition string) {
//	//排序规则
//	sortCondition = dao.RefundInfo.Columns.CreatedAt + " ASC" //id升序
//	if req.Sort == "recent" {                               //最近上架
//		sortCondition = dao.RefundInfo.Columns.CreatedAt + " DESC" //创建时间倒序
//	} else if req.Sort == "sale" {
//		sortCondition = dao.RefundInfo.Columns.Sale + " DESC" //销量倒序
//	} else if req.Sort == "price_up" {
//		sortCondition = dao.RefundInfo.Columns.Price + " ASC" //价格升序
//	} else if req.Sort == "price_down" {
//		sortCondition = dao.RefundInfo.Columns.Price + " DESC" //价格降序
//	}
//	return
//}

func (s *refundService) Detail(ctx context.Context, req *DetailReq) (res ListRefundSql, err error) {
	err = dao.RefundInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
