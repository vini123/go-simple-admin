// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goSimpleAdmin/internal/model"
)

type (
	IUser interface {
		SignIn(ctx context.Context, in model.SignInInput) (err error)
		SignUp(ctx context.Context, in model.SignUpInput) (err error)
		Captcha(ctx context.Context, in model.CaptchaInput) (err error)
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
