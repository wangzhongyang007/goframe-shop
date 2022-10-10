package login

type LoginReq struct {
	Name     string `json:"name"`
	PassWord string `json:"password" v:"required-if:type,0|password#password必须传递|密码限定在6-18位之间"`
}

type AccessTokenReq struct {
	AppKey    string `json:"app_key"`
	SecretKey string `json:"secret_key"`
}

type AccessTokenRes struct {
	AccessToken string `json:"access_token"` //获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   //凭证有效时间，单位：秒
	//以下for 刷新token 这版不做
	//Time int `json:"time"`
	//RefreshToken        string `json:"refresh_token"`
	//RefreshTokenExpires int    `json:"refresh_token_expires"` //refresh_token过期时间戳
}

type AccessTokenInvalidReq struct {
	AccessToken string `json:"access_token"`
}

type GetRedirectReq struct {
	Redirect string `p:"redirect" v:"required#redirect必须传递"`
}

type LoginQRCodeRes struct {
	QrCodeUrl string `json:"qr_code_url"`
	SceneStr  string `json:"scene_str"`
}

type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
}

type CheckTicketReq struct {
	Ticket string `p:"ticket" v:"required#ticket必须传递"`
}
type CheckTicketRes struct {
	Token string `p:"token"`
}

type LogoutReq struct {
	Redirect string `p:"redirect" v:"required#redirect必须传递"`
}
