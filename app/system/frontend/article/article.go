package article

type AddArticleReq struct {
	UserId  int    `json:"user_id,omitempty"`
	Title   string `json:"title" v:"required#标题必传"`
	Desc    string `json:"desc" v:"required#摘要描述必传"`
	PicUrl  string `json:"pic_url" v:"required#封面图必传"`
	Detail  string `json:"detail" v:"required#详情必传"`
	IsAdmin int    `json:"is_admin" d:"0"` //管理后台发布设置为1
}

type UpdateArticleReq struct {
	Id int `json:"id"`
	AddArticleReq
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

type ListArticleRes struct {
	Count int               `json:"count"`
	List  []*ListArticleSql `json:"list"`
}

type ListArticleSql struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	PicUrl  string `json:"pic_url"`
	Detail  string `json:"detail"`
	IsAdmin int    `json:"is_admin"` //管理后台发布设置为1
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
