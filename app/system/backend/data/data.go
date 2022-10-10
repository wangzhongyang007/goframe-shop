package data

type AddDataReq struct {
	PicUrl string `json:"pic_url" v:"required#图片链接必传"`
	Link   string `json:"link" v:"required#跳转链接必传"`
	Sort   string `json:"sort" d:"1"`
}

type UpdateDataReq struct {
	Id     int    `json:"id"`
	PicUrl string `json:"pic_url,omitempty" v:"required#图片链接必传"`
	Link   string `json:"link,omitempty" v:"required#跳转链接必传"`
	Sort   string `json:"sort,omitempty" d:"1"`
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Page  int `json:"page" v:"required#请合理输入页数"`
	Limit int `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
}

type UploadImgReq struct {
	File string `json:"file"` //要上传文件的本地路径
}

type ListDataRes struct {
	Count int            `json:"count"`
	List  []*ListDataSql `json:"list"`
}

type HeadCardRes struct {
	TodayOrderCount int `json:"today_order_count"`
	DAU             int `json:"dau"`
	ConversionRate  int `json:"conversion_rate" description:"转化率"`
}

type EChartsRes struct {
	OrderTotal           []int `json:"order_total"`
	SalePriceTotal       []int `json:"sale_price_total"`
	ConsumptionPerPerson []int `json:"consumption_per_person" description:"人均消费"`
	NewOrder             []int `json:"new_order" description:"新增订单"`
}

type ListDataSql struct {
	Id     int    `json:"id"`
	PicUrl string `json:"pic_url"`
	Link   string `json:"link"`
	Sort   string `json:"sort"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TodayTotal struct {
	Today string `json:"today"`
	Total int    `json:"total"`
}
