package goodsOptions

import (
	"context"
	"database/sql"
	"shop/internal/dao"
)

var service = new(goodsOptionsService)

type goodsOptionsService struct {
}

func (s *goodsOptionsService) Add(ctx context.Context, req *AddGoodsOptionsReq) (res sql.Result, err error) {
	res, err = dao.GoodsOptionsInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *goodsOptionsService) Update(ctx context.Context, req *UpdateGoodsOptionsReq) (res sql.Result, err error) {
	res, err = dao.GoodsOptionsInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *goodsOptionsService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.GoodsOptionsInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *goodsOptionsService) List(ctx context.Context, req *PageListReq) (res ListGoodsOptionsRes, err error) {
	count, err := dao.GoodsOptionsInfo.Ctx(ctx).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.GoodsOptionsInfo.Ctx(ctx).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func (s *goodsOptionsService) Detail(ctx context.Context, req *DetailReq) (res ListGoodsOptionsSql, err error) {
	err = dao.GoodsOptionsInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
