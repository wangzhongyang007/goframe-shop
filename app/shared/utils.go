package shared

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"math/rand"
	"time"
)

func GetOrderNum() (number string) {
	rand.Seed(time.Now().UnixNano())
	number = gconv.String(time.Now().UnixNano()) + gconv.String(rand.Intn(1000))
	return
}

func GetRefundNum() (number string) {
	rand.Seed(time.Now().UnixNano())
	number = "refund" + gconv.String(time.Now().UnixNano()) + gconv.String(rand.Intn(1000))
	return
}

//生成最近一周的日期
func GetRecent7Date() (dates []string) {
	gt := gtime.New(time.Now())
	dates = []string{
		gt.Format("Y-m-d"),
		gt.Add(-gtime.D).Format("Y-m-d"),
		gt.Add(-gtime.D * 2).Format("Y-m-d"),
		gt.Add(-gtime.D * 3).Format("Y-m-d"),
		gt.Add(-gtime.D * 4).Format("Y-m-d"),
		gt.Add(-gtime.D * 5).Format("Y-m-d"),
		gt.Add(-gtime.D * 6).Format("Y-m-d"),
	}
	return
}

//获取一周前的日期
func GetBefore7Date() (date string) {
	gt := gtime.New(time.Now())
	date = gt.Add(-gtime.D * 6).Format("Y-m-d")
	return
}
