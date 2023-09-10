package model

// 验证规则 https://goframe.org/pages/viewpage.action?pageId=1114367

import (
	"goSimpleAdmin/internal/model/entity"
	"time"
)

type User struct {
	entity.Users
	UserExtends *entity.UserExtends      `orm:"with:user_id=id" json:"user_extends"`
	Roles       []map[string]interface{} `json:"roles,omitempty"`
	Permissions []Permission             `json:"permissions,omitempty"`
}

// 用户权限
type Permission struct {
	entity.Permissions
	Permission []Permission `json:"permission,omitempty"`
}

// 图片验证码
type CaptchaReq struct{}

type CaptchaRes struct {
	Id     string `json:"id"`
	Base64 string `json:"base64"`
}

// 登录
type SignInReq struct {
	Account  string `v:"required|length:2,64#请输入账号|账号长度为 {min} 到 {max} 位" json:"account"`
	Password string `v:"required|length:6,16#请输入密码|密码长度为 {min} 到 {max} 位" json:"password"`
}

type SignInRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// 注册
type SignUpReq struct {
	Nickname   string `v:"required|length:2,12#请输入昵称|昵称长度为 {min} 到 {max} 位" json:"nickname"`
	Account    string `v:"required|length:2,64#请输入账号|账号长度为 {min} 到 {max} 位" json:"account"`
	Password   string `v:"required|length:6,16#请输入密码|密码长度为 {min} 到 {max} 位" json:"password"`
	VerifyKey  string `v:"required|size:20#缺少验证KEY|验证KEY长度不正确" json:"verify_key"`
	VerifyCode string `v:"required|size:4#请输入验证码|验证码长度为 4 位" json:"verify_code"`
}

type SignUpRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// 获取用户信息
type UserInfoReq struct{}

type UserInfoRes struct {
	User
}
