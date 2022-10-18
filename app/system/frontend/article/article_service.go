package article

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
	"shop/app/model"
	"time"
)

const ArticleDetailCacheKey = "ArticleDetailCacheKey_"

var service = new(rotationService)

type rotationService struct {
}

func (s *rotationService) Add(r *ghttp.Request, req *AddArticleReq) (res sql.Result, err error) {
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
	//更新缓存
	cacheKey := ArticleDetailCacheKey + gconv.String(req.Id)
	res, err = dao.ArticleInfo.Ctx(ctx).Cache(-1, cacheKey).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) Delete(r *ghttp.Request, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.ArticleInfo.Ctx(r.GetCtx()).
		WherePri(req.Id).
		Where(dao.ArticleInfo.Columns.UserId, r.GetCtxVar(middleware.CtxAccountId)).
		Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) Detail(r *ghttp.Request, req *DetailReq) (res model.ArticleInfo, err error) {
	//查询时优先查询缓存
	cacheKey := ArticleDetailCacheKey + gconv.String(req.Id)
	err = dao.ArticleInfo.Ctx(r.GetCtx()).Cache(time.Hour, cacheKey).WherePri(req.Id).Scan(&res)
	if err != nil {
		return res, err
	}
	return
}

func (s *rotationService) List(ctx context.Context, req *PageListReq) (res ListArticleRes, err error) {
	count, err := dao.ArticleInfo.Ctx(ctx).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.ArticleInfo.Ctx(ctx).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func (s *rotationService) MyList(r *ghttp.Request, req *PageListReq) (res ListArticleRes, err error) {
	ctx := r.GetCtx()
	whereCondition := g.Map{
		dao.ArticleInfo.Columns.IsAdmin: 0,
		dao.ArticleInfo.Columns.UserId:  r.GetCtxVar(middleware.CtxAccountId),
	}
	count, err := dao.ArticleInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.ArticleInfo.Ctx(ctx).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}
