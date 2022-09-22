package userCoupon

import "github.com/gogf/gf/util/gmeta"

type AddCouponReq struct {
	CouponId int `json:"coupon_id" v:"required#优惠券Id必填"`
	Status   int `json:"status,omitempty"`
	UserId   int `json:"user_id,omitempty"`
}

type UpdateCouponReq struct {
	Id         int    `json:"id"`
	Name       string `json:"name" v:"required#名称必填"`
	Price      int    `json:"price" v:"required#优惠券面值必填"`
	GoodsId    string `json:"goods_id"`
	CategoryId string `json:"category_id"`
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Page  int `json:"page" v:"required#请合理输入页数"`
	Limit int `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
}

type ListCouponRes struct {
	Count int              `json:"count"`
	List  []*ListCouponSql `json:"list"`
}

type ListCouponSql struct {
	gmeta.Meta `orm:"table:user_coupon_info"`
	Id         int         `json:"id"`
	UserId     int         `json:"user_id"`
	CouponId   string      `json:"coupon_id"`
	CouponInfo *CouponInfo `orm:"with:id=coupon_id" json:"coupon_info"`
	TimeCommon
}

type CouponInfo struct {
	gmeta.Meta `orm:"table:coupon_info"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	GoodsId    int    `json:"goods_id"`
	CategoryId int    `json:"category_id"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
