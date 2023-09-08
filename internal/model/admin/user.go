package model

// 图片验证码
type CaptchaReq struct{}

type CaptchaRes struct {
	Id     string `json:"id"`
	Base64 string `json:"base64"`
}

// 登录
type SignInReq struct {
	Passport string `v:"required|length:2,64#请输入账号|账号长度为 {min} 到 {max} 位" json:"passport"`
	Password string `v:"required|length:6,16#请输入密码|密码长度为 {min} 到 {max} 位" json:"password"`
}

type SignInRes struct {
	Token string `json:"token"`
}

// 注册
type SignUpReq struct {
	Nickname   string `v:"required|length:2,12#请输入昵称#昵称长度为 {min} 到 {max} 位" json:"nickname"`
	Passport   string `v:"required|length:2,64#账号长度为 {min} 到 {max} 位" json:"passport"`
	Password   string `v:"required|length:6,16#请输入密码|密码长度为 {min} 到 {max} 位" json:"password"`
	VerifyKey  string `v:"required|length:12#缺少验证KEY|验证KEY长度为 12 位" json:"verify_key"`
	VerifyCode string `v:"required|length:4#请输入验证码|验证码长度为 4 位" json:"verify_code"`
}

type SignUpRes struct {
	Token string `json:"token"`
}
