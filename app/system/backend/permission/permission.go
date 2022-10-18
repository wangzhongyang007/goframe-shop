package permission

type AddPermissionReq struct {
	Name string `json:"name" v:"required#用户名必传"`
	Path string `json:"path" v:"required#路径"`
}

type UpdatePermissionReq struct {
	Id int `json:"id"`
	AddPermissionReq
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Page    int    `json:"page" v:"required#请合理输入页数"`
	Limit   int    `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
	Keyword string `json:"keyword"`
}

type ListPermissionRes struct {
	Count int                  `json:"count"`
	List  []*ListPermissionSql `json:"list"`
}

type ListPermissionSql struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
