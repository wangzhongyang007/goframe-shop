package address

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/frame/g"
	"shop/app/dao"
)

var service = new(addressService)

type addressService struct {
}

func (s *addressService) Add(ctx context.Context, req *AddAddressReq) (res sql.Result, err error) {
	res, err = dao.AddressInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *addressService) Update(ctx context.Context, req *UpdateAddressReq) (res sql.Result, err error) {
	res, err = dao.AddressInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *addressService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.AddressInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *addressService) List(ctx context.Context, req *PageListReq) (res ListAddressRes, err error) {
	whereCondition := g.Map{
		dao.AddressInfo.Columns.Pid: req.Pid,
	}
	count, err := dao.AddressInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.AddressInfo.Ctx(ctx).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}
