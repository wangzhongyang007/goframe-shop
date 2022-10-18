package admin

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"shop/middleware"
	"shop/internal/dao"
	"shop/utility"
)

var service = new(rotationService)

type rotationService struct {
}

func (s *rotationService) Add(ctx context.Context, req *AddAdminReq) (res sql.Result, err error) {
	UserSalt := grand.S(10)
	req.Password = utility.EncryptPassword(req.Password, UserSalt)
	req.UserSalt = UserSalt
	res, err = dao.AdminInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) Update(ctx context.Context, req *UpdateAdminReq) (res sql.Result, err error) {
	if req.Password != "" {
		UserSalt := grand.S(10)
		req.Password = utility.EncryptPassword(req.Password, UserSalt)
		req.UserSalt = UserSalt
	}
	res, err = dao.AdminInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) UpdateMyPassword(r *ghttp.Request, req *UpdateMyPasswordReq) (res sql.Result, err error) {
	if req.Password != "" {
		UserSalt := grand.S(10)
		req.Password = utility.EncryptPassword(req.Password, UserSalt)
		req.UserSalt = UserSalt
	}
	//获得当前登录用户
	req.Id = gconv.Int(r.GetCtxVar(middleware.CtxAccountId))
	ctx := r.GetCtx()
	res, err = dao.AdminInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.AdminInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) List(ctx context.Context, req *PageListReq) (res ListAdminRes, err error) {
	whereCondition := g.Map{}
	if req.Keyword != "" {
		whereCondition = g.Map{
			dao.AdminInfo.Columns.Name + " like ": "%" + req.Keyword + "%",
		}
	}
	count, err := dao.AdminInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.AdminInfo.Ctx(ctx).Where(whereCondition).Page(req.Page, req.Limit).OrderDesc("id").Scan(&res.List)
	if err != nil {
		return
	}
	return
}
