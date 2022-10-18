package goodsOptions

type AddGoodsOptionsReq struct {
	GoodsId int    `json:"goods_id" v:"required#商品id必传"`
	PicUrl  string `json:"pic_url" v:"required#图片链接必传"`
	Name    string `json:"name" v:"required#商品名称必传"`
	Price   int    `json:"price" v:"required#商品价格必传"`
	Stock   int    `json:"stock" v:"required#库存必传"`
}

type UpdateGoodsOptionsReq struct {
	Id      int    `json:"id"`
	GoodsId int    `json:"goods_id" v:"required#商品id必传"`
	PicUrl  string `json:"pic_url" v:"required#图片链接必传"`
	Name    string `json:"name" v:"required#商品名称必传"`
	Price   int    `json:"price" v:"required#商品价格必传"`
	Stock   int    `json:"stock" v:"required#库存必传"`
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type DetailReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Page  int `json:"page" v:"required#请合理输入页数"`
	Limit int `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
}

type ListGoodsOptionsRes struct {
	Count int                    `json:"count"`
	List  []*ListGoodsOptionsSql `json:"list"`
}

type DetailGoodsOptionsRes struct {
	ListGoodsOptionsSql
}

type ListGoodsOptionsSql struct {
	Id      int    `json:"id"`
	GoodsId int    `json:"goods_id"`
	PicUrl  string `json:"pic_url"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Stock   int    `json:"stock"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
