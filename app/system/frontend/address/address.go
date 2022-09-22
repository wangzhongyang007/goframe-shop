package address

type AddAddressReq struct {
	Pid  int    `json:"pid" v:"required#父级id必传"`
	Name string `json:"name" v:"required#名称必传"`
}

type UpdateAddressReq struct {
	Id   int    `json:"id"`
	Pid  int    `json:"pid,omitempty"`
	Name string `json:"name,omitempty"`
}

type SoftDeleteReq struct {
	Id int `json:"id"`
}

type PageListReq struct {
	Pid   int `json:"pid" v:"required#父id必传"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
	//Page     int `json:"page" v:"required#请合理输入页数"`
	//Limit    int `json:"limit" v:"limit@required|max:100#请合理输入条数|条数最大为100"`
}

type ListAddressRes struct {
	Count int               `json:"count"`
	List  []*ListAddressSql `json:"list"`
}

type ListAddressSql struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pid  int    `json:"pid"`
	TimeCommon
}

type TimeCommon struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
