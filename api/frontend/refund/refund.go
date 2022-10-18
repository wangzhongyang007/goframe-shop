package refund

import "github.com/gogf/gf/v2/util/gmeta"

type AddRefundReq struct {
	Number  string `json:"number,omitempty" description:"订单编号"`
	UserId  int    `json:"user_id,omitempty"`
	OrderId int    `json:"order_id,omitempty"`
	GoodsId int    `json:"goods_id,omitempty"`
	Reason  string `json:"reason" description:"原因"`
	Status  int    `json:"status" description:"状态 0待处理 1同意退款 2拒绝退款"`
}

type UpdateRefundReq struct {
	Id      int `json:"id"`
	GoodsId int `json:"goods_id" v:"required#商品id必传"`
	Count   int `json:"count" v:"required#商品数量必传"`
	UserId  int `json:"user_id,omitempty"`
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Status int `json:"status,omitempty"`
	Page   int `json:"page" v:"required#请合理输入页数"`
	Limit  int `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
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

type ListRefundRes struct {
	Count int              `json:"count"`
	List  []*ListRefundSql `json:"list"`
}

type ListRefundSql struct {
	gmeta.Meta `orm:"table:refund_info"`
	Id         int       `json:"id"`
	Number     string    `json:"number" description:"订单编号"`
	UserId     int       `json:"user_id"`
	OrderId    int       `json:"order_id"`
	GoodsId    int       `json:"goods_id"`
	GoodsInfo  GoodsInfo `orm:"with:id=goods_id" json:"goods_info"`
	Reason     string    `json:"reason" description:"原因"`
	Status     int       `json:"status" description:"状态 1待处理 2同意退款 3拒绝退款"`
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
