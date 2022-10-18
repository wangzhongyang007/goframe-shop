package goods

type AddGoodsReq struct {
	CouponId    int    `json:"coupon_id"`
	PicUrl      string `json:"pic_url" v:"required#图片链接必传"`
	Name        string `json:"name" v:"required#商品名称必传"`
	Price       int    `json:"price" v:"required#商品价格必传"`
	CategoryIds string `json:"category_ids" v:"required#商品分类必传"`
	Brand       string `json:"brand" v:"required#商品品牌必传"`
	Stock       int    `json:"stock" v:"required#库存必传"`
	Tags        string `json:"tags" v:"required#标签必传"`
	DetailInfo  string `json:"detail_info" v:"required#详情必传"`
}

type UpdateGoodsReq struct {
	Id          int    `json:"id"`
	CouponId    int    `json:"coupon_id"`
	PicUrl      string `json:"pic_url" v:"required#图片链接必传"`
	Name        string `json:"name" v:"required#商品名称必传"`
	Price       int    `json:"price" v:"required#商品价格必传"`
	CategoryIds string `json:"category_ids" v:"required#商品分类必传"`
	Brand       string `json:"brand" v:"required#商品品牌必传"`
	Stock       int    `json:"stock" v:"required#库存必传"`
	Tags        string `json:"tags" v:"required#标签必传"`
	DetailInfo  string `json:"detail_info" v:"required#详情必传"`
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

type ListGoodsRes struct {
	Count int             `json:"count"`
	List  []*ListGoodsSql `json:"list"`
}

type ListGoodsSql struct {
	Id          int    `json:"id"`
	CouponId    int    `json:"coupon_id"`
	PicUrl      string `json:"pic_url"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	CategoryIds string `json:"category_ids"`
	Brand       string `json:"brand"`
	Stock       int    `json:"stock"`
	Sale        int    `json:"sale"`
	Tags        string `json:"tags"`
	DetailInfo  string `json:"detail_info"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DetailReq struct {
	Id int `json:"id"`
}
