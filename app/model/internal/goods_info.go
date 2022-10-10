// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// GoodsInfo is the golang structure for table goods_info.
type GoodsInfo struct {
	Id               int         `orm:"id,primary"         json:"id"`               //
	PicUrl           string      `orm:"pic_url"            json:"picUrl"`           // 图片
	Name             string      `orm:"name"               json:"name"`             // 商品名称
	Price            int         `orm:"price"              json:"price"`            // 价格 单位分
	Level1CategoryId int         `orm:"level1_category_id" json:"level1CategoryId"` // 1级分类id
	Level2CategoryId int         `orm:"level2_category_id" json:"level2CategoryId"` // 2级分类id
	Level3CategoryId int         `orm:"level3_category_id" json:"level3CategoryId"` // 3级分类id
	Brand            string      `orm:"brand"              json:"brand"`            // 品牌
	CouponId         int         `orm:"coupon_id"          json:"couponId"`         // 优惠券id
	Stock            int         `orm:"stock"              json:"stock"`            // 库存
	Sale             int         `orm:"sale"               json:"sale"`             // 销量
	Tags             string      `orm:"tags"               json:"tags"`             // 标签
	DetailInfo       string      `orm:"detail_info"        json:"detailInfo"`       // 商品详情
	CreatedAt        *gtime.Time `orm:"created_at"         json:"createdAt"`        //
	UpdatedAt        *gtime.Time `orm:"updated_at"         json:"updatedAt"`        //
	DeletedAt        *gtime.Time `orm:"deleted_at"         json:"deletedAt"`        //
}