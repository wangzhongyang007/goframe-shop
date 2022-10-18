package goods

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/container/gmap"
	"shop/internal/dao"
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

func (s *goodsService) List(ctx context.Context, req *PageListReq) (res ListGoodsRes, err error) {
	//实例化map
	whereCondition := gmap.New()
	//很好的理解了map是引用类型的特点 在这个函数中为查询条件赋值
	packListCondition(req, whereCondition)
	//map是引用类型，在packListCondition函数中已经做了赋值操作，不需要在接收返回值
	count, err := dao.GoodsInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.GoodsInfo.Ctx(ctx).Where(whereCondition).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func packListCondition(req *PageListReq, whereCondition *gmap.Map) {
	//使用map支持set的特性 避免在声明的时候赋值，那么写需要做的判断太复杂了。
	if req.Keyword != "" {
		whereCondition.Set(dao.GoodsInfo.Columns.DetailInfo+" like ", "%"+req.Keyword+"%")
	}
	if req.Name != "" {
		whereCondition.Set(dao.GoodsInfo.Columns.Name+" like ", "%"+req.Name+"%")
	}
	if req.Brand != "" {
		whereCondition.Set(dao.GoodsInfo.Columns.Brand+" like ", "%"+req.Brand+"%")
	}
}

func (s *goodsService) Detail(ctx context.Context, req *DetailReq) (res ListGoodsSql, err error) {
	err = dao.GoodsInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
