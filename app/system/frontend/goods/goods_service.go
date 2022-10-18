package goods

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"shop/app/dao"
	"shop/app/model"
)

var service = new(goodsService)

type goodsService struct {
}

func (s *goodsService) Add(ctx context.Context, req *AddGoodsReq) (res sql.Result, err error) {
	res, err = dao.GoodsInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *goodsService) Update(ctx context.Context, req *UpdateGoodsReq) (res sql.Result, err error) {
	res, err = dao.GoodsInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *goodsService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.GoodsInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

//func (s *goodsService) List(ctx context.Context, req *PageListReq) (res ListGoodsRes, err error) {
//	count, err := dao.GoodsInfo.Ctx(ctx).Count()
//	if err != nil {
//		return
//	}
//	res.Count = count
//	err = dao.GoodsInfo.Ctx(ctx).Page(req.Page, req.Limit).Scan(&res.List)
//	if err != nil {
//		return
//	}
//	return
//}

//同类商品推荐
func (s *goodsService) Category(ctx context.Context, req *CategoryPageListReq) (res ListGoodsRes, err error) {
	//获取商品的分类
	currentGoods := model.GoodsInfo{}
	err = dao.GoodsInfo.Ctx(ctx).WherePri(req.Id).Scan(&currentGoods)
	if err != nil {
		return ListGoodsRes{}, err
	}

	whereLevelCondition := g.Map{
		"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{currentGoods.Level1CategoryId, currentGoods.Level2CategoryId, currentGoods.Level3CategoryId},
	}
	whereIdCondition := g.Map{
		"id!=": req.Id,
	}
	count, err := dao.GoodsInfo.Ctx(ctx).Where(whereIdCondition).Where(whereLevelCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.GoodsInfo.Ctx(ctx).Where(whereIdCondition).Where(whereLevelCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func packCondition(req *SearchPageListReq, whereCondition *gmap.Map) {
	if req.Keyword != "" {
		whereCondition.Set(dao.GoodsInfo.Columns.Name+" like ", "%"+req.Keyword+"%")
	}
	if req.CategoryId != 0 {
		whereCondition.Set("level1_category_id =? OR level2_category_id =? OR level3_category_id =? ", g.Slice{req.CategoryId, req.CategoryId, req.CategoryId})
	}
}

func (s *goodsService) List(ctx context.Context, req *SearchPageListReq) (res ListGoodsRes, err error) {
	whereCondition := gmap.New()
	packCondition(req, whereCondition)
	//这是优化之前的代码
	//if req.Keyword != "" && req.CategoryId != 0 {
	//	whereCondition = g.Map{
	//		"name like": "%" + req.Keyword + "%",
	//		"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{req.CategoryId, req.CategoryId, req.CategoryId},
	//	}
	//} else if req.Keyword != "" {
	//	whereCondition = g.Map{
	//		"name like": "%" + req.Keyword + "%",
	//	}
	//} else if req.CategoryId != 0 {
	//	whereCondition = g.Map{
	//		"level1_category_id =? OR level2_category_id =? OR level3_category_id =? ": g.Slice{req.CategoryId, req.CategoryId, req.CategoryId},
	//	}
	//} else {
	//	whereCondition = g.Map{}
	//}

	//获取数量
	count, err := dao.GoodsInfo.Ctx(ctx).
		Where(whereCondition).
		Count()
	if err != nil {
		return
	}
	res.Count = count

	//获取值
	//排序规则
	sortCondition := packSort(req)
	err = dao.GoodsInfo.Ctx(ctx).
		Where(whereCondition).
		Page(req.Page, req.Limit).
		Order(sortCondition).
		Scan(&res.List)
	if err != nil {
		return
	}
	return
}

//封装排序方法
func packSort(req *SearchPageListReq) (sortCondition string) {
	//排序规则
	sortCondition = dao.GoodsInfo.Columns.CreatedAt + " ASC" //id升序
	if req.Sort == "recent" {                                //最近上架
		sortCondition = dao.GoodsInfo.Columns.CreatedAt + " DESC" //创建时间倒序
	} else if req.Sort == "sale" {
		sortCondition = dao.GoodsInfo.Columns.Sale + " DESC" //销量倒序
	} else if req.Sort == "price_up" {
		sortCondition = dao.GoodsInfo.Columns.Price + " ASC" //价格升序
	} else if req.Sort == "price_down" {
		sortCondition = dao.GoodsInfo.Columns.Price + " DESC" //价格降序
	}
	return
}

func (s *goodsService) Detail(ctx context.Context, req *DetailReq) (res ListGoodsSql, err error) {
	err = dao.GoodsInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
