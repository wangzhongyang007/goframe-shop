package user

type AddUserReq struct {
	Name       string `json:"name" v:"required#用户名必传"`
	Password   string `json:"password" v:"required#密码必传"`
	Permission string `json:"permission" v:"required#权限必传"`
	UserSalt   string `json:"user_salt,omitempty"`
}

type UpdateUserReq struct {
	Id     int `json:"id" v:"required#用户名Id必传"`
	Status int `json:"status" v:"required#用户名状态必传：2拉黑冻结 1正常"`
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Sex     int    `json:"sex"`
	Name    string `json:"name"`
	Page    int    `json:"page" v:"required#请合理输入页数"`
	Limit   int    `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
	Keyword string `json:"keyword"`
}

type ListUserRes struct {
	Count int            `json:"count"`
	List  []*ListUserSql `json:"list"`
}

type ListUserSql struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    int    `json:"sex"`
	Status int    `json:"status"`
	Sign   string `json:"sign"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
