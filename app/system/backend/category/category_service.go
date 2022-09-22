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
	count, err := dao.CategoryInfo.Ctx(ctx).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.CategoryInfo.Ctx(ctx).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

//分级列表
func (s *categoryService) LevelList(ctx context.Context, req *PageListReq) (res LevelListCategoryRes, err error) {
	//处理一级分类
	whereCondition := g.Map{
		dao.CategoryInfo.Columns.Level: 1,
	}
	count, err := dao.CategoryInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.CategoryInfo.Ctx(ctx).Where(whereCondition).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}

	//处理二级分类
	for _, categorySql := range res.List {
		whereCondition = g.Map{
			dao.CategoryInfo.Columns.ParentId: categorySql.Id,
			dao.CategoryInfo.Columns.Level:    2,
		}
		err = dao.CategoryInfo.Ctx(ctx).Where(whereCondition).OrderDesc("id").Scan(&categorySql.Items)
		//处理三级分类
		for _, level2Category := range categorySql.Items {
			whereCondition = g.Map{
				dao.CategoryInfo.Columns.ParentId: level2Category.Id,
				dao.CategoryInfo.Columns.Level:    3,
			}
			err = dao.CategoryInfo.Ctx(ctx).Where(whereCondition).OrderDesc("id").Scan(&level2Category.Items)
		}
		if err != nil {
			return
		}
	}
	return
}
