package user

import (
	"context"
	"goSimpleAdmin/internal/model"
	"goSimpleAdmin/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) SignIn(ctx context.Context, in model.SignInInput) (err error) {
	return nil
}

func (s *sUser) SignUp(ctx context.Context, in model.SignUpInput) (err error) {
	return nil
}

func (s *sUser) Captcha(ctx context.Context, in model.CaptchaInput) (err error) {
	return nil
}
