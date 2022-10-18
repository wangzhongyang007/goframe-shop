package article

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
	"shop/app/model"
)

var service = new(rotationService)

type rotationService struct {
}

func (s *rotationService) Add(r *ghttp.Request, req *AddArticleReq) (res sql.Result, err error) {
	//获得当前登录用户
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.ArticleInfo.Ctx(r.GetCtx()).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) Update(r *ghttp.Request, req *UpdateArticleReq) (res sql.Result, err error) {
	ctx := r.GetCtx()
	//校验是不是自己的文章
	articleInfo := model.ArticleInfo{}
	err = dao.ArticleInfo.Ctx(ctx).WherePri(req.Id).Scan(&articleInfo)
	if err != nil {
		return nil, err
	}
	if articleInfo.UserId != gconv.Int(r.GetCtxVar(middleware.CtxAccountId)) {
		return nil, errors.New("这不是您的文章，不允许修改")
	}

	res, err = dao.ArticleInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.ArticleInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) List(ctx context.Context, req *PageListReq) (res ListArticleRes, err error) {
	//实例化map
	whereCondition := gmap.New()
	//很好的理解了map是引用类型的特点 在这个函数中为查询条件赋值
	packListCondition(req, whereCondition)
	count, err := dao.ArticleInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.ArticleInfo.Ctx(ctx).Where(whereCondition).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func packListCondition(req *PageListReq, whereCondition *gmap.Map) {
	//使用map支持set的特性 避免在声明的时候赋值，那么写需要做的判断太复杂了。
	if req.UserId != 0 {
		whereCondition.Set(dao.ArticleInfo.Columns.UserId, req.UserId)
	}
	if req.IsAdmin != 0 {
		whereCondition.Set(dao.ArticleInfo.Columns.IsAdmin, req.IsAdmin)
	}
	if req.Keyword != "" {
		whereCondition.Set(dao.ArticleInfo.Columns.Title+" like ", "%"+req.Keyword+"%")
	}
}
