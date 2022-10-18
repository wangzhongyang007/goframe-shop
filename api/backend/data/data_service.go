package data

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"shop/internal/dao"
	"shop/internal/shared"
	"shop/utility"
	"time"
)

var service = new(dataService)

type dataService struct {
}

func (s *dataService) HeadCard(ctx context.Context) (res HeadCardRes, err error) {
	res.TodayOrderCount = TodayOrderCount(ctx)
	res.DAU = utility.RandInt(200)
	res.ConversionRate = utility.RandInt(80)
	return
}

func (s *dataService) ECharts(ctx context.Context) (res EChartsRes, err error) {
	res.OrderTotal = OrderTotal(ctx)
	res.SalePriceTotal = SalePriceTotalRecentDays(ctx)
	//todo
	res.ConsumptionPerPerson = OrderTotal(ctx)
	res.NewOrder = OrderTotal(ctx) //新增订单和今日订单一致
	//res.ConsumptionPerPerson = ConsumptionPerPerson(ctx)
	//res.NewOrder = TodayOrderCount(ctx) //新增订单和今日订单一致
	return
}

//今日订单数量
func TodayOrderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns.CreatedAt, gtime.New(time.Now()).StartOfDay(), gtime.New(time.Now()).EndOfDay()).
		Count("id")
	if err != nil {
		return 0
	}
	return
}

//select date_format(created_at, '%Y-%m-%d') today, count(*) as cnt from order_info group by today
/**
gf官方示例：
// SELECT COUNT(*) total,age FROM `user` GROUP BY age
db.Model("user").Fields("COUNT(*) total,age").Group("age").All()
*/
func OrderTotal(ctx context.Context) (counts []int) {
	counts = []int{0, 0, 0, 0, 0, 0, 0}
	recent7Dates := shared.GetRecent7Date()
	TodayTotals := []TodayTotal{}
	//只取最近7天
	err := dao.OrderInfo.Ctx(ctx).Where(dao.OrderInfo.Columns.CreatedAt+" >= ", shared.GetBefore7Date()).Fields("count(*) total,date_format(created_at, '%Y-%m-%d') today").Group("today").Scan(&TodayTotals)
	fmt.Printf("result:%v", TodayTotals)
	for i, date := range recent7Dates {
		for _, todayTotal := range TodayTotals {
			if date == todayTotal.Today {
				counts[i] = todayTotal.Total
			}
		}
	}
	if err != nil {
		return counts
	}
	return
}

func SalePriceTotalRecentDays(ctx context.Context) (totals []int) {
	totals = []int{0, 0, 0, 0, 0, 0, 0}
	recent7Dates := shared.GetRecent7Date()
	TodayTotals := []TodayTotal{}
	//只取最近7天
	err := dao.OrderInfo.Ctx(ctx).Where(dao.OrderInfo.Columns.CreatedAt+" >= ", shared.GetBefore7Date()).Fields("sum(actual_price) total,date_format(created_at, '%Y-%m-%d') today").Group("today").Scan(&TodayTotals)
	fmt.Printf("result:%v", TodayTotals)
	for i, date := range recent7Dates {
		for _, todayTotal := range TodayTotals {
			if date == todayTotal.Today {
				totals[i] = todayTotal.Total
			}
		}
	}
	if err != nil {
		return totals
	}
	return
}

func SalePriceTotal(ctx context.Context) (total int) {
	sum, err := dao.OrderInfo.Ctx(ctx).Sum(dao.OrderInfo.Columns.ActualPrice)
	if err != nil {
		return 0
	}
	total = int(sum)
	return
}

func userTotal(ctx context.Context) (total int) {
	total, err := dao.UserInfo.Ctx(ctx).Count("id")
	if err != nil {
		return 0
	}
	return
}

//人均消费 todo 算每天的人均消费
func ConsumptionPerPerson(ctx context.Context) (per int) {
	priceTotal := SalePriceTotal(ctx)
	userTotal := userTotal(ctx)
	per = priceTotal / userTotal
	return
}
