// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// RefundInfo is the golang structure for table refund_info.
type RefundInfo struct {
	Id        int         `orm:"id,primary" json:"id"`        // 售后退款表
	Number    string      `orm:"number"     json:"number"`    // 售后订单号
	OrderId   int         `orm:"order_id"   json:"orderId"`   // 订单id
	GoodsId   int         `orm:"goods_id"   json:"goodsId"`   // 要售后的商品id
	Reason    string      `orm:"reason"     json:"reason"`    // 退款原因
	Status    int         `orm:"status"     json:"status"`    // 状态 1待处理 2同意退款 3拒绝退款
	UserId    int         `orm:"user_id"    json:"userId"`    // 用户id
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` //
}
