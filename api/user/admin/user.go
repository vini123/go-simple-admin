package admin

import "github.com/gogf/gf/v2/frame/g"

// 请求数字验证码
type CaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"CaptchaService" summary:"build captcha"`
}

type CaptchaRes struct {
	Id string
}

// 登录
type SignInReq struct {
	g.Meta   `path:"/user/sign-in" method:"post" tags:"UserService" summary:"user sign in"`
	Passport string `v:"required|length:2,64#请输入账号|账号长度为 {min} 到 {max} 位"`
	Password string `v:"required|length:6,16#请输入密码|密码长度为 {min} 到 {max} 位"`
}

type SignInRes struct{}

// 注册（只需要数字验证码就可以）
type SignUpReq struct {
	g.Meta     `path:"/user/sing-up" method:"post" tags:"UserService" summary:"user sign up"`
	Nickname   string `v:"required|length:2,12#请输入昵称#昵称长度为 {min} 到 {max} 位"`
	Passport   string `v:"required|length:2,64#账号长度为 {min} 到 {max} 位"`
	Password   string `v:"required|length:6,16#请输入密码|密码长度为 {min} 到 {max} 位"`
	VerifyKey  string `v:"required|length:12#缺少验证KEY|验证KEY长度为 12 位"`
	VerifyCode string `v:"required|length:4#请输入验证码|验证码长度为 4 位"`
}

type SignUpRes struct{}
