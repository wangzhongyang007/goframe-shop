package rotation

import (
	"context"
	"database/sql"
	"shop/app/dao"
)

var service = new(rotationService)

type rotationService struct {
}

func (s *rotationService) Add(ctx context.Context, req *AddRotationReq) (res sql.Result, err error) {
	res, err = dao.RotationInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) Update(ctx context.Context, req *UpdateRotationReq) (res sql.Result, err error) {
	res, err = dao.RotationInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.RotationInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) List(ctx context.Context, req *PageListReq) (res ListRotationRes, err error) {
	count, err := dao.RotationInfo.Ctx(ctx).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.RotationInfo.Ctx(ctx).Page(req.Page, req.Limit).OrderDesc(dao.RotationInfo.Columns.Sort).Scan(&res.List)
	if err != nil {
		return
	}
	return
}
