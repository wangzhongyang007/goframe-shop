package consignee

import (
	"context"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
	"shop/app/dao"
)

var service = new(consigneeService)

type consigneeService struct {
}

func (s *consigneeService) List(r *ghttp.Request, req *PageListReq) (res ListConsigneeRes, err error) {
	//实例化map
	whereCondition := gmap.New()
	//很好的理解了map是引用类型的特点 在这个函数中为查询条件赋值
	packListCondition(req, whereCondition)
	count, err := dao.ConsigneeInfo.Ctx(r.GetCtx()).Where(whereCondition).Count()
	if err != nil {
		return
	}
	res.Count = count
	err = dao.ConsigneeInfo.Ctx(r.GetCtx()).Where(whereCondition).Page(req.Page, req.Limit).Scan(&res.List)
	if err != nil {
		return
	}
	return
}

func packListCondition(req *PageListReq, whereCondition *gmap.Map) {
	//使用map支持set的特性 避免在声明的时候赋值，那么写需要做的判断太复杂了。
	if req.Name != "" {
		whereCondition.Set(dao.ConsigneeInfo.Columns.Name+" like ", "%"+req.Name+"%")
	}
	if req.Phone != "" {
		whereCondition.Set(dao.ConsigneeInfo.Columns.Phone+" like ", "%"+req.Phone+"%")
	}
}

func (s *consigneeService) Detail(ctx context.Context, req *DetailReq) (res ListConsigneeSql, err error) {
	err = dao.ConsigneeInfo.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return
	}
	return
}
