package comment

type AddCommentReq struct {
	UserId   int    `json:"user_id,omitempty"`
	Type     int    `json:"type" d:"2" description:"评论类型：1商品 2文章"`
	ObjectId int    `json:"object_id" v:"required#收藏对象id必传"`
	Content  string `json:"content" v:"required#评论内容必传"`
	ParentId int    `json:"parent_id" description:"父级评论id"`
}

//type UpdateCommentReq struct {
//	Id        int    `json:"id"`
//	UserId    int    `json:"user_id,omitempty"`
//	IsDefault int    `json:"is_default" v:"required#是否默认地址必传"`
//	Name      string `json:"name" v:"required#收货人姓名必传"`
//	Phone     string `json:"phone" v:"required#收货人手机号必传"`
//	Province  string `json:"province" v:"required#省份必传"`
//	City      string `json:"city" v:"required#城市必传"`
//	Town      string `json:"town" v:"required#县区必传"`
//	Street    string `json:"street" v:"required#乡镇街道必传"`
//	Detail    string `json:"detail" v:"required#详细地址必传"`
//}

type DeleteReq struct {
	Id int `json:"id"`
	//Type     int `json:"type"`
	//ObjectId int `json:"object_id"`
	//ParentId int `json:"parent_id"`
}

type PageListReq struct {
	UserId   int    `json:"user_id,omitempty"`
	Type     int    `json:"type,omitempty"`
	ObjectId int    `json:"object_id,omitempty"`
	Keyword  string `json:"keyword"`
	Page     int    `json:"page" v:"required#请合理输入页数"`
	Limit    int    `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
}

//type SearchPageListReq struct {
//	Keyword    string `json:"keyword"`
//	CategoryId int    `json:"category_id"`
//	Page       int    `json:"page" v:"required#请合理输入页数"`
//	Limit      int    `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
//	Sort       string `json:"sort"` //取值排序规则 recent意为最新上架
//}

//同类商品推荐
//type CategoryPageListReq struct {
//	Id    int `json:"id"`
//	Page  int `json:"page" v:"required#请合理输入页数"`
//	Limit int `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
//}

type ListCommentRes struct {
	Count int               `json:"count"`
	List  []*ListCommentSql `json:"list"`
}

type ListCommentSql struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Type     int    `json:"type" d:"2" description:"评论类型：1商品 2文章"`
	ObjectId int    `json:"object_id" v:"required#收藏对象id必传"`
	Content  string `json:"content" v:"required#评论内容必传"`
	ParentId int    `json:"parent_id" description:"父级评论id"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DetailReq struct {
	Id int `json:"id"`
}
