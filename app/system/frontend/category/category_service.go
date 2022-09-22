package category

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/frame/g"
	"shop/app/dao"
)

var service = new(categoryService)

type categoryService struct {
}

func (s *categoryService) Add(ctx context.Context, req *AddCategoryReq) (res sql.Result, err error) {
	res, err = dao.CategoryInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *categoryService) Update(ctx context.Context, req *UpdateCategoryReq) (res sql.Result, err error) {
	res, err = dao.CategoryInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *categoryService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.CategoryInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *categoryService) List(ctx context.Context, req *PageListReq) (res ListCategoryRes, err error) {
	whereCondition := g.Map{
		dao.CategoryInfo.Columns.ParentId: req.ParentId,
	}
	count, err := dao.CategoryInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.CategoryInfo.Ctx(ctx).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}
