// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	model "goSimpleAdmin/internal/model/admin"
)

type (
	IUser interface {
		SignIn(ctx context.Context, in model.SignInReq) (res model.SignInRes, err error)
		SignUp(ctx context.Context, in model.SignUpReq) (res model.SignUpRes, err error)
		Captcha(ctx context.Context, in model.CaptchaReq) (res model.CaptchaRes, err error)
		UserInfo(ctx context.Context, in model.UserInfoReq) (res model.UserInfoRes, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
