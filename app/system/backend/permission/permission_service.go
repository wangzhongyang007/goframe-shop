package permission

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"shop/app/dao"
)

var service = new(permissionService)

type permissionService struct {
}

func (s *permissionService) Add(ctx context.Context, req *AddPermissionReq) (res sql.Result, err error) {
	res, err = dao.PermissionInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *permissionService) Update(ctx context.Context, req *UpdatePermissionReq) (res sql.Result, err error) {
	res, err = dao.PermissionInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *permissionService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.PermissionInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *permissionService) List(ctx context.Context, req *PageListReq) (res ListPermissionRes, err error) {
	whereCondition := g.Map{}
	if req.Keyword != "" {
		whereCondition = g.Map{
			dao.PermissionInfo.Columns.Name + " like ": "%" + req.Keyword + "%",
		}
	}
	count, err := dao.PermissionInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.PermissionInfo.Ctx(ctx).Where(whereCondition).Page(req.Page, req.Limit).OrderDesc("id").Scan(&res.List)
	if err != nil {
		return
	}
	return
}
