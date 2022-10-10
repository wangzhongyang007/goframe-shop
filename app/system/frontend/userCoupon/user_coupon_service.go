package userCoupon

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
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

//func (s *couponService) Update(r *ghttp.Request, req *UpdateCouponReq) (res sql.Result, err error) {
//	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
//	res, err = dao.UserCouponInfo.Ctx(r.GetCtx()).WherePri(req.Id).Update(req)
//	if err != nil {
//		return nil, err
//	}
//	return
//}

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

//同类商品推荐
//func (s *couponService) Category(ctx context.Context, req *CategoryPageListReq) (res ListCouponRes, err error) {
//	//获取商品的分类
//	currentCoupon := model.UserCouponInfo{}
//	err = dao.UserCouponInfo.Ctx(ctx).WherePri(req.Id).Scan(&currentCoupon)
//	if err != nil {
//		return ListCouponRes{}, err
//	}
//
//	whereLevelCondition := g.Map{
//		"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{currentCoupon.Level1CategoryId, currentCoupon.Level2CategoryId, currentCoupon.Level3CategoryId},
//	}
//	whereIdCondition := g.Map{
//		"id!=": req.Id,
//	}
//	count, err := dao.UserCouponInfo.Ctx(ctx).Where(whereIdCondition).Where(whereLevelCondition).Count()
//	if err != nil {
//		return
//	}
//	res.Count = count
//	err = dao.UserCouponInfo.Ctx(ctx).Where(whereIdCondition).Where(whereLevelCondition).Page(req.Page, req.Limit).Scan(&res.List)
//	if err != nil {
//		return
//	}
//	return
//}

//func (s *couponService) List(ctx context.Context, req *PageListReq) (res ListCouponRes, err error) {
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
//	count, err := dao.UserCouponInfo.Ctx(ctx).
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
//	err = dao.UserCouponInfo.Ctx(ctx).
//		Where(whereCondition).
//		Page(req.Page, req.Limit).
//		Order(sortCondition).
//		Scan(&res.List)
//	if err != nil {
//		return
//	}
//	return
//}

//封装排序方法
//func packSort(req *SearchPageListReq) (sortCondition string) {
//	//排序规则
//	sortCondition = dao.UserCouponInfo.Columns.CreatedAt + " ASC" //id升序
//	if req.Sort == "recent" {                               //最近上架
//		sortCondition = dao.UserCouponInfo.Columns.CreatedAt + " DESC" //创建时间倒序
//	} else if req.Sort == "sale" {
//		sortCondition = dao.UserCouponInfo.Columns.Sale + " DESC" //销量倒序
//	} else if req.Sort == "price_up" {
//		sortCondition = dao.UserCouponInfo.Columns.Price + " ASC" //价格升序
//	} else if req.Sort == "price_down" {
//		sortCondition = dao.UserCouponInfo.Columns.Price + " DESC" //价格降序
//	}
//	return
//}

//func (s *couponService) Detail(ctx context.Context, req *DetailReq) (res ListCouponSql, err error) {
//	err = dao.UserCouponInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
//	if err != nil {
//		return
//	}
//	return
//}
