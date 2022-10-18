package consignee

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
)

var service = new(consigneeService)

type consigneeService struct {
}

func (s *consigneeService) Add(r *ghttp.Request, req *AddConsigneeReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	//req.Number = shared.GetConsigneeNum()
	res, err = dao.ConsigneeInfo.Ctx(r.GetCtx()).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *consigneeService) Update(r *ghttp.Request, req *UpdateConsigneeReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.ConsigneeInfo.Ctx(r.GetCtx()).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *consigneeService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.ConsigneeInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *consigneeService) List(r *ghttp.Request, req *PageListReq) (res ListConsigneeRes, err error) {
	whereCondition := g.Map{
		dao.ConsigneeInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
	}

	count, err := dao.ConsigneeInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.ConsigneeInfo.Ctx(r.GetCtx()).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

//同类商品推荐
//func (s *consigneeService) Category(ctx context.Context, req *CategoryPageListReq) (res ListConsigneeRes, err error) {
//	//获取商品的分类
//	currentConsignee := model.ConsigneeInfo{}
//	err = dao.ConsigneeInfo.Ctx(ctx).WherePri(req.Id).Scan(&currentConsignee)
//	if err != nil {
//		return ListConsigneeRes{}, err
//	}
//
//	whereLevelCondition := g.Map{
//		"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{currentConsignee.Level1CategoryId, currentConsignee.Level2CategoryId, currentConsignee.Level3CategoryId},
//	}
//	whereIdCondition := g.Map{
//		"id!=": req.Id,
//	}
//	count, err := dao.ConsigneeInfo.Ctx(ctx).Where(whereIdCondition).Where(whereLevelCondition).Count()
//	if err != nil {
//		return
//	}
//	res.Count = count
//	err = dao.ConsigneeInfo.Ctx(ctx).Where(whereIdCondition).Where(whereLevelCondition).Page(req.Page, req.Limit).Scan(&res.List)
//	if err != nil {
//		return
//	}
//	return
//}

//func (s *consigneeService) List(ctx context.Context, req *PageListReq) (res ListConsigneeRes, err error) {
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
//	count, err := dao.ConsigneeInfo.Ctx(ctx).
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
//	err = dao.ConsigneeInfo.Ctx(ctx).
//		Where(whereCondition).
//		Page(req.Page, req.Limit).
//		Consignee(sortCondition).
//		Scan(&res.List)
//	if err != nil {
//		return
//	}
//	return
//}

//封装排序方法
//func packSort(req *SearchPageListReq) (sortCondition string) {
//	//排序规则
//	sortCondition = dao.ConsigneeInfo.Columns.CreatedAt + " ASC" //id升序
//	if req.Sort == "recent" {                               //最近上架
//		sortCondition = dao.ConsigneeInfo.Columns.CreatedAt + " DESC" //创建时间倒序
//	} else if req.Sort == "sale" {
//		sortCondition = dao.ConsigneeInfo.Columns.Sale + " DESC" //销量倒序
//	} else if req.Sort == "price_up" {
//		sortCondition = dao.ConsigneeInfo.Columns.Price + " ASC" //价格升序
//	} else if req.Sort == "price_down" {
//		sortCondition = dao.ConsigneeInfo.Columns.Price + " DESC" //价格降序
//	}
//	return
//}

func (s *consigneeService) Detail(ctx context.Context, req *DetailReq) (res ListConsigneeSql, err error) {
	err = dao.ConsigneeInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
