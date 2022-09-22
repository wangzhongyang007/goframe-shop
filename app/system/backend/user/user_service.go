package user

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/grand"
	"shop/app/dao"
	"shop/library"
)

var service = new(rotationService)

type rotationService struct {
}

func (s *rotationService) Add(ctx context.Context, req *AddUserReq) (res sql.Result, err error) {
	UserSalt := grand.S(10)
	req.Password = library.EncryptPassword(req.Password, UserSalt)
	req.UserSalt = UserSalt
	res, err = dao.UserInfo.Ctx(ctx).Insert(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) Update(ctx context.Context, req *UpdateUserReq) (res sql.Result, err error) {
	res, err = dao.UserInfo.Ctx(ctx).WherePri(req.Id).Update(req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *rotationService) Delete(ctx context.Context, req *SoftDeleteReq) (res sql.Result, err error) {
	res, err = dao.UserInfo.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, err
	}
	return
}

//订单列表
func (s *rotationService) OrderList(ctx context.Context, req *PageListReq) (res ListUserRes, err error) {
	whereCondition := packListCondition
	count, err := dao.UserInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.UserInfo.Ctx(ctx).Where(whereCondition).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func (s *rotationService) List(ctx context.Context, req *PageListReq) (res ListUserRes, err error) {
	//实例化map
	whereCondition := gmap.New()
	//很好的理解了map是引用类型的特点 在这个函数中为查询条件赋值
	packListCondition(req, whereCondition)
	//map是引用类型，在packListCondition函数中已经做了赋值操作，不需要在接收返回值
	count, err := dao.UserInfo.Ctx(ctx).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.UserInfo.Ctx(ctx).Where(whereCondition).OrderDesc("id").Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func packListCondition(req *PageListReq, whereCondition *gmap.Map) {
	//使用map支持set的特性 避免在声明的时候赋值，那么写需要做的判断太复杂了。
	if req.Sex != 0 {
		whereCondition.Set(dao.UserInfo.Columns.Sex, req.Sex)
	}
	if req.Name != "" {
		whereCondition.Set(dao.UserInfo.Columns.Name, req.Name)
	}
	if req.Keyword != "" {
		whereCondition.Set(dao.UserInfo.Columns.Name+" like ", "%"+req.Keyword+"%")
	}
}
