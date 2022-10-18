package order

type AddOrderReq struct {
	Number           string `json:"number,omitempty" description:"订单编号"`
	UserId           int    `json:"user_id,omitempty"`
	PayType          int    `json:"pay_type" description:"支付方式 1微信 2支付宝 3云闪付"`
	Remark           string `json:"remark" description:"备注"`
	PayAt            string `json:"pay_at,omitempty" description:"支付时间"`
	Status           int    `json:"status" description:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价"`
	Price            int    `json:"price" description:"订单金额 单位分"`
	CouponPrice      int    `json:"coupon_price" description:"优惠券金额 单位分"`
	ActualPrice      int    `json:"actual_price" description:"实际支付金额 单位分"`
	ConsigneeName    string `json:"consignee_name" description:"收货人姓名"`
	ConsigneePhone   string `json:"consignee_phone" description:"收货人手机号"`
	ConsigneeAddress string `json:"consignee_address" description:"收货人地址"`
}

type UpdateOrderReq struct {
	Id      int `json:"id"`
	GoodsId int `json:"goods_id" v:"required#商品id必传"`
	Count   int `json:"count" v:"required#商品数量必传"`
	UserId  int `json:"user_id,omitempty"`
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	UserId  int    `json:"user_id,omitempty"`
	Status  int    `json:"status,omitempty"`
	Date    string `json:"date,omitempty" description:"日期"`
	Keyword string `json:"keyword,omitempty" description:"模糊搜索地址"`
	Page    int    `json:"page" v:"required#请合理输入页数"`
	Limit   int    `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
	Number  string `json:"number"` //模糊查询
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

type ListOrderRes struct {
	Count int             `json:"count"`
	List  []*ListOrderSql `json:"list"`
}

type ListOrderSql struct {
	Id               int    `json:"id"`
	Number           string `json:"number,omitempty" description:"订单编号"`
	UserId           int    `json:"user_id,omitempty"`
	PayType          int    `json:"pay_type" description:"支付方式 1微信 2支付宝 3云闪付"`
	Remark           string `json:"remark" description:"备注"`
	PayAt            string `json:"pay_at,omitempty" description:"支付时间"`
	Status           int    `json:"status" description:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价"`
	Price            int    `json:"price" description:"订单金额 单位分"`
	CouponPrice      int    `json:"coupon_price" description:"优惠券金额 单位分"`
	ActualPrice      int    `json:"actual_price" description:"实际支付金额 单位分"`
	ConsigneeName    string `json:"consignee_name" description:"收货人姓名"`
	ConsigneePhone   string `json:"consignee_phone" description:"收货人手机号"`
	ConsigneeAddress string `json:"consignee_address" description:"收货人地址"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DetailReq struct {
	Id int `json:"id"`
}
