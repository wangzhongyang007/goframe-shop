package rotation

type AddRotationReq struct {
	PicUrl string `json:"pic_url" v:"required#图片链接必传"`
	Link   string `json:"link" v:"required#跳转链接必传"`
	Sort   string `json:"sort" d:"1"`
}

type UpdateRotationReq struct {
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

type ListRotationRes struct {
	Count int                `json:"count"`
	List  []*ListRotationSql `json:"list"`
}

type ListRotationSql struct {
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
