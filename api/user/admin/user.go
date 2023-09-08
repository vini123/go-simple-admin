package admin

import (
	model "goSimpleAdmin/internal/model/admin"

	"github.com/gogf/gf/v2/frame/g"
)

// 请求数字验证码
type CaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"CaptchaService" summary:"build captcha"`
}

type CaptchaRes struct {
	model.CaptchaRes
}

// 登录
type SignInReq struct {
	g.Meta `path:"/user/sign-in" method:"post" tags:"UserService" summary:"user sign in"`
	model.SignInReq
}

type SignInRes struct {
	model.SignInRes
}

// 注册（只需要数字验证码就可以）
type SignUpReq struct {
	g.Meta `path:"/user/sing-up" method:"post" tags:"UserService" summary:"user sign up"`
	model.SignUpReq
}

type SignUpRes struct {
	model.SignUpRes
}

// 用户信息
type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"UserService" summary:"get userinfo"`
	model.UserInfoReq
}

type UserInfoRes struct {
	model.UserInfoRes
}
