// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package user

import (
	"context"
	
	"goSimpleAdmin/api/user/admin"
)

type IUserAdmin interface {
	Captcha(ctx context.Context, req *admin.CaptchaReq) (res *admin.CaptchaRes, err error)
	SignIn(ctx context.Context, req *admin.SignInReq) (res *admin.SignInRes, err error)
	SignUp(ctx context.Context, req *admin.SignUpReq) (res *admin.SignUpRes, err error)
	UserInfo(ctx context.Context, req *admin.UserInfoReq) (res *admin.UserInfoRes, err error)
}


