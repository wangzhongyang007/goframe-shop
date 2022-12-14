// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// OrderGoodsInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type OrderGoodsInfoDao struct {
	gmvc.M                        // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB                // DB is the raw underlying database management object.
	Table   string                // Table is the table name of the DAO.
	Columns orderGoodsInfoColumns // Columns contains all the columns of Table that for convenient usage.
}

// OrderGoodsInfoColumns defines and stores column names for table order_goods_info.
type orderGoodsInfoColumns struct {
	Id          string // 商品维度的订单表
	OrderId     string // 关联的主订单表
	GoodsId     string // 商品id
	Count       string // 商品数量
	PayType     string // 支付方式 1微信 2支付宝 3云闪付
	Remark      string // 备注
	Status      string // 订单状态 0待支付 1已支付 3已确认收货
	Price       string // 订单金额 单位分
	CouponPrice string // 优惠券金额 单位分
	ActualPrice string // 实际支付金额 单位分
	PayAt       string // 支付时间
	CreatedAt   string //
	UpdatedAt   string //
}

func NewOrderGoodsInfoDao() *OrderGoodsInfoDao {
	return &OrderGoodsInfoDao{
		M:     g.DB("default").Model("order_goods_info").Safe(),
		DB:    g.DB("default"),
		Table: "order_goods_info",
		Columns: orderGoodsInfoColumns{
			Id:          "id",
			OrderId:     "order_id",
			GoodsId:     "goods_id",
			Count:       "count",
			PayType:     "pay_type",
			Remark:      "remark",
			Status:      "status",
			Price:       "price",
			CouponPrice: "coupon_price",
			ActualPrice: "actual_price",
			PayAt:       "pay_at",
			CreatedAt:   "created_at",
			UpdatedAt:   "updated_at",
		},
	}
}
