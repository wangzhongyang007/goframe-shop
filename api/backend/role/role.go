package role

type AddRoleReq struct {
	Name string `json:"name" v:"required#用户名必传"`
	Desc string `json:"desc" v:"required#描述"`
}

type RolePermissionReq struct {
	RoleId       int `json:"role_id" v:"required#角色id必传"`
	PermissionId int `json:"permission_id" v:"required#权限id必传"`
}

type UpdateRoleReq struct {
	Id int `json:"id"`
	AddRoleReq
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}
  
type PageListReq struct {
	Page    int    `json:"page" v:"required#请合理输入页数"`
	Limit   int    `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
	Keyword string `json:"keyword"`
}

type ListRoleRes struct {
	Count int            `json:"count"`
	List  []*ListRoleSql `json:"list"`
}

type ListRoleSql struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
