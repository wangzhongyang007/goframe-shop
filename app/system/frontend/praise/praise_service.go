package praise

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
)

var service = new(praiseService)

type praiseService struct {
}

func (s *praiseService) Add(r *ghttp.Request, req *AddPraiseReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.PraiseInfo.Ctx(r.GetCtx()).Replace(req)
	if err != nil {
		return nil, err
	}
	return
}

//func (s *praiseService) Update(r *ghttp.Request, req *UpdatePraiseReq) (res sql.Result, err error) {
//	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
//	res, err = dao.PraiseInfo.Ctx(r.GetCtx()).WherePri(req.Id).Update(req)
//	if err != nil {
//		return nil, err
//	}
//	return
//}

func (s *praiseService) Delete(ctx context.Context, req *DeleteReq) (res sql.Result, err error) {
	if req.Id != 0 {
		//根据收藏id删除
		res, err = dao.PraiseInfo.Ctx(ctx).WherePri(req.Id).Delete()
	} else {
		//根据类型和对象id删除
		res, err = dao.PraiseInfo.Ctx(ctx).
			//Where(dao.PraiseInfo.Columns.Type, req.Type).
			Where(dao.PraiseInfo.Columns.ObjectId, req.ObjectId).
			Delete()
	}
	if err != nil {
		return nil, err
	}
	return
}

func (s *praiseService) List(r *ghttp.Request, req *PageListReq) (res ListPraiseRes, err error) {
	whereCondition := g.Map{}
	if req.Type == 0 {
		whereCondition = g.Map{
			dao.PraiseInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
		}
	} else {
		whereCondition = g.Map{
			dao.PraiseInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
			//dao.PraiseInfo.Columns.Type:   req.Type,
		}
	}

	count, err := dao.PraiseInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.PraiseInfo.Ctx(r.GetCtx()).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func (s *praiseService) Detail(ctx context.Context, req *DetailReq) (res ListPraiseSql, err error) {
	err = dao.PraiseInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
