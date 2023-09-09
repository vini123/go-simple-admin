package user

import (
	"context"

	"goSimpleAdmin/api/user/admin"
	model "goSimpleAdmin/internal/model/admin"
	"goSimpleAdmin/internal/service"
)

func (c *ControllerAdmin) SignIn(ctx context.Context, req *admin.SignInReq) (res *admin.SignInRes, err error) {
	data, err := service.User().SignIn(ctx, model.SignInReq{
		Account:  req.Account,
		Password: req.Password,
	})

	if err != nil {
		return
	}

	res = new(admin.SignInRes)
	res.SignInRes = data
	return
}
