package collection

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
)

var service = new(collectionService)

type collectionService struct {
}

func (s *collectionService) Add(r *ghttp.Request, req *AddCollectionReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.CollectionInfo.Ctx(r.GetCtx()).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *collectionService) Delete(ctx context.Context, req *DeleteReq) (res sql.Result, err error) {
	if req.Id != 0 {
		//根据收藏id删除
		res, err = dao.CollectionInfo.Ctx(ctx).WherePri(req.Id).Delete()
	} else {
		//根据类型和对象id删除
		res, err = dao.CollectionInfo.Ctx(ctx).
			Where(dao.CollectionInfo.Columns.Type, req.Type).
			Where(dao.CollectionInfo.Columns.ObjectId, req.ObjectId).
			Delete()
	}
	if err != nil {
		return nil, err
	}
	return
}

func (s *collectionService) List(r *ghttp.Request, req *PageListReq) (res ListCollectionRes, err error) {
	whereCondition := g.Map{}
	if req.Type == 0 {
		whereCondition = g.Map{
			dao.CollectionInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
		}
	} else {
		whereCondition = g.Map{
			dao.CollectionInfo.Columns.UserId: r.GetCtxVar(middleware.CtxAccountId),
			dao.CollectionInfo.Columns.Type:   req.Type,
		}
	}

	count, err := dao.CollectionInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.CollectionInfo.Ctx(r.GetCtx()).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func (s *collectionService) Detail(ctx context.Context, req *DetailReq) (res ListCollectionSql, err error) {
	err = dao.CollectionInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
