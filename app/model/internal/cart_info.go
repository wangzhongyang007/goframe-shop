// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// CartInfo is the golang structure for table cart_info.
type CartInfo struct {
	Id        int         `orm:"id,primary" json:"id"`        // 购物车表
	UserId    int         `orm:"user_id"    json:"userId"`    //
	GoodsId   int         `orm:"goods_id"   json:"goodsId"`   //
	Count     int         `orm:"count"      json:"count"`     // 商品数量
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` //
}
