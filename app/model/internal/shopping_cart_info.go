// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ShoppingCartInfo is the golang structure for table shopping_cart_info.
type ShoppingCartInfo struct {
	Id        int         `orm:"id,primary" json:"id"`        // 购物车表
	UserId    int         `orm:"user_id"    json:"userId"`    //
	GoodsId   int         `orm:"goods_id"   json:"goodsId"`   //
	Count     int         `orm:"count"      json:"count"`     // 商品数量
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
}
