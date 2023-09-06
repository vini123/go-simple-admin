package user

import (
	"context"

	"goSimpleAdmin/api/user/admin"
	"goSimpleAdmin/internal/model"
	"goSimpleAdmin/internal/service"
)

func (c *ControllerAdmin) SignIn(ctx context.Context, req *admin.SignInReq) (res *admin.SignInRes, err error) {
	service.User().SignIn(ctx, model.SignInInput{
		Passport: req.Passport,
		Password: req.Passport,
	})
	return
}
