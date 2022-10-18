package refund

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
	UserId int `json:"user_id,omitempty"`
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
	Id      int    `json:"id"`
	Number  string `json:"number,omitempty" description:"订单编号"`
	UserId  int    `json:"user_id,omitempty"`
	OrderId int    `json:"order_id,omitempty"`
	GoodsId int    `json:"goods_id,omitempty"`
	Reason  string `json:"reason" description:"原因"`
	Status  int    `json:"status" description:"状态 1待处理 2同意退款 3拒绝退款"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DetailReq struct {
	Id int `json:"id"`
}
