package coupon

type AddCouponReq struct {
	Name       string `json:"name" v:"required#优惠券名称必传"`
	Price      int    `json:"price" v:"required#优惠券金额必传 单位分"`
	GoodsId    int    `json:"goods_id,omitempty"`
	CategoryId int    `json:"category_id,omitempty"`
}

type UpdateCouponReq struct {
	Id         int    `json:"id"`
	Name       string `json:"name" v:"required#优惠券名称必传"`
	Price      int    `json:"price" v:"required#优惠券金额必传 单位分"`
	GoodsId    int    `json:"goods_id,omitempty"`
	CategoryId int    `json:"category_id,omitempty"`
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Page  int `json:"page" v:"required#请合理输入页数"`
	Limit int `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
}

type SearchPageListReq struct {
	Keyword    string `json:"keyword"`
	CategoryId int    `json:"category_id"`
	Page       int    `json:"page" v:"required#请合理输入页数"`
	Limit      int    `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
	Sort       string `json:"sort"` //取值排序规则 recent意为最新上架
}

//同类商品推荐
type CategoryPageListReq struct {
	Id    int `json:"id"`
	Page  int `json:"page" v:"required#请合理输入页数"`
	Limit int `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
}

type ListCouponRes struct {
	Count int              `json:"count"`
	List  []*ListCouponSql `json:"list"`
}

type ListCouponSql struct {
	Id         int    `json:"id"`
	Name       string `json:"name" v:"required#优惠券名称必传"`
	Price      int    `json:"price" v:"required#优惠券金额必传 单位分"`
	GoodsId    int    `json:"goods_id,omitempty"`
	CategoryId int    `json:"category_id,omitempty"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DetailReq struct {
	Id int `json:"id"`
}
