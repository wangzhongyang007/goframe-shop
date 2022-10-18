package category

type AddCategoryReq struct {
	ParentId int    `json:"parent_id" v:"required#父级id必传"`
	Name     string `json:"name" v:"required#名称必传"`
	PicUrl   string `json:"pic_url"`
	Sort     int    `json:"sort" d:"1"`
	Level    int    `json:"level" d:"1"`
}

type UpdateCategoryReq struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent_id,omitempty"`
	PicUrl   string `json:"pic_url,omitempty"`
	Name     string `json:"name,omitempty"`
	Sort     int    `json:"sort,omitempty"`
	Level    int    `json:"level,omitempty"`
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Page  int `json:"page" v:"required#请合理输入页数"`
	Limit int `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
}

type LevelListCategoryRes struct {
	Count int                     `json:"count"`
	List  []*LevelListCategorySql `json:"list"`
}

type LevelListCategorySql struct {
	Id     int                     `json:"id"`
	Name   string                  `json:"name"`
	PicUrl string                  `json:"pic_url"`
	Sort   int                     `json:"sort"`
	Level  int                     `json:"level"`
	Items  []*LevelListCategorySql `json:"items"`
	TimeCommon
}

type ListCategoryRes struct {
	Count int                `json:"count"`
	List  []*ListCategorySql `json:"list"`
}

type ListCategorySql struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	PicUrl string `json:"pic_url"`
	Sort   int    `json:"sort"`
	Level  int    `json:"level"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
