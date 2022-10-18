package role

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"shop/app/dao"
)

var service = new(roleService)

type roleService struct {
}

func (s *roleService) Add(ctx context.Context, req *AddRoleReq) (res sql.Result, err error) {
	res, err = dao.RoleInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *roleService) AddPermission(ctx context.Context, req *RolePermissionReq) (res sql.Result, err error) {
	res, err = dao.RolePermissionInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *roleService) DeletePermission(ctx context.Context, req *RolePermissionReq) (res sql.Result, err error) {
	whereCondition := g.Map{
		dao.RolePermissionInfo.Columns.RoleId:       req.RoleId,
		dao.RolePermissionInfo.Columns.PermissionId: req.PermissionId,
	}
	res, err = dao.RolePermissionInfo.Ctx(ctx).Where(whereCondition).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *roleService) Update(ctx context.Context, req *UpdateRoleReq) (res sql.Result, err error) {
	res, err = dao.RoleInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *roleService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.RoleInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *roleService) List(ctx context.Context, req *PageListReq) (res ListRoleRes, err error) {
	whereCondition := g.Map{}
	if req.Keyword != "" {
		whereCondition = g.Map{
			dao.RoleInfo.Columns.Name + " like ": "%" + req.Keyword + "%",
		}
	}
	count, err := dao.RoleInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.RoleInfo.Ctx(ctx).Where(whereCondition).Page(req.Page, req.Limit).OrderDesc("id").Scan(&res.List)
	if err != nil {
		return
	}
	return
}
