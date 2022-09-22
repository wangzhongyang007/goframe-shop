package admin

type AddAdminReq struct {
	Name       string `json:"name" v:"required#用户名必传"`
	Password   string `json:"password" v:"required#密码必传"`
	Permission string `json:"permission" v:"required#权限必传"`
	UserSalt   string `json:"user_salt,omitempty"`
}

type UpdateAdminReq struct {
	Id int `json:"id"`
	AddAdminReq
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Page    int    `json:"page" v:"required#请合理输入页数"`
	Limit   int    `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
	Keyword string `json:"keyword"`
}

type ListAdminRes struct {
	Count int             `json:"count"`
	List  []*ListAdminSql `json:"list"`
}

type ListAdminSql struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
