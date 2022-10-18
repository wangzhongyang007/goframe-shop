package cart

import "github.com/gogf/gf/v2/util/gmeta"

type AddCartReq struct {
	GoodsId int `json:"goods_id" v:"required#商品id必传"`
	Count   int `json:"count" v:"required#商品数量必传"`
	UserId  int `json:"user_id,omitempty"`
}

type UpdateCartReq struct {
	Id      int `json:"id"`
	GoodsId int `json:"goods_id" v:"required#商品id必传"`
	Count   int `json:"count" v:"required#商品数量必传"`
	UserId  int `json:"user_id,omitempty"`
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

type ListCartRes struct {
	Count int            `json:"count"`
	List  []*ListCartSql `json:"list"`
}

type ListCartSql struct {
	gmeta.Meta `orm:"table:cart_info"`
	Id         int `json:"id"`
	UserId     int `json:"user_id"`
	GoodsId    int `json:"goods_id"`
	Count      int `json:"count"`
	//以下字段要根据查询goods表关联取值
	//with:当前属性对应表关联字段=当前结构体对应数据表关联字段
	GoodsInfo *GoodsInfo `orm:"with:id=goods_id" json:"goods_info"`
	TimeCommon
}

type GoodsInfo struct {
	gmeta.Meta       `orm:"table:goods_info"`
	Id               string `json:"id"`
	PicUrl           string `json:"pic_url"`
	Name             string `json:"name"`
	Price            int    `json:"price"`
	Level1CategoryId int    `json:"level1_category_id"`
	Level2CategoryId int    `json:"level2_category_id"`
	Level3CategoryId int    `json:"level3_category_id"`
	Brand            string `json:"brand"`
	Stock            int    `json:"stock"`
	Sale             int    `json:"sale"`
	Tags             string `json:"tags"`
	DetailInfo       string `json:"detail_info"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DetailReq struct {
	Id int `json:"id"`
}
