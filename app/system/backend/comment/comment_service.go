package comment

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"shop/app/dao"
	"shop/app/middleware"
)

var service = new(commentService)

type commentService struct {
}

func (s *commentService) Add(r *ghttp.Request, req *AddCommentReq) (res sql.Result, err error) {
	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	res, err = dao.CommentInfo.Ctx(r.GetCtx()).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

//func (s *commentService) Update(r *ghttp.Request, req *UpdateCommentReq) (res sql.Result, err error) {
//	req.UserId = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
//	res, err = dao.CommentInfo.Ctx(r.GetCtx()).WherePri(req.Id).Update(req)
//	if err != nil {
//		return nil, err
//	}
//	return
//}

func (s *commentService) Delete(ctx context.Context, req *DeleteReq) (res sql.Result, err error) {
	if req.Id != 0 {
		//根据收藏id删除
		res, err = dao.CommentInfo.Ctx(ctx).WherePri(req.Id).Delete()
	}
	//else {
	//	//根据类型和对象id删除
	//	res, err = dao.CommentInfo.Ctx(ctx).
	//		Where(dao.CommentInfo.Columns.Type, req.Type).
	//		Where(dao.CommentInfo.Columns.ObjectId, req.ObjectId).
	//		Where(dao.CommentInfo.Columns.ParentId, req.ParentId).
	//		Delete()
	//}
	if err != nil {
		return nil, err
	}
	return
}

func (s *commentService) List(r *ghttp.Request, req *PageListReq) (res ListCommentRes, err error) {
	//实例化map
	whereCondition := gmap.New()
	//很好的理解了map是引用类型的特点 在这个函数中为查询条件赋值
	packListCondition(req, whereCondition)
	//map是引用类型，在packListCondition函数中已经做了赋值操作，不需要在接收返回值
	count, err := dao.CommentInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.CommentInfo.Ctx(r.GetCtx()).Where(whereCondition).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func packListCondition(req *PageListReq, whereCondition *gmap.Map) {
	//使用map支持set的特性 避免在声明的时候赋值，那么写需要做的判断太复杂了。
	if req.UserId != 0 {
		whereCondition.Set(dao.CommentInfo.Columns.UserId, req.UserId)
	}
	if req.Type != 0 {
		whereCondition.Set(dao.CommentInfo.Columns.Type, req.Type)
	}
	if req.ObjectId != 0 {
		whereCondition.Set(dao.CommentInfo.Columns.ObjectId, req.ObjectId)
	}
	if req.Keyword != "" {
		whereCondition.Set(dao.CommentInfo.Columns.Content+" like ", "%"+req.Keyword+"%")
	}
}

func (s *commentService) Detail(ctx context.Context, req *DetailReq) (res ListCommentSql, err error) {
	err = dao.CommentInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
