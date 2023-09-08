package user

import (
	"context"

	"goSimpleAdmin/api/user/admin"
	model "goSimpleAdmin/internal/model/admin"
	"goSimpleAdmin/internal/service"
)

func (c *ControllerAdmin) SignUp(ctx context.Context, req *admin.SignUpReq) (res *admin.SignUpRes, err error) {
	service.User().SignUp(ctx, model.SignUpReq{
		Nickname:   req.Nickname,
		Passport:   req.Passport,
		Password:   req.Password,
		VerifyKey:  req.VerifyKey,
		VerifyCode: req.VerifyCode,
	})
	return
}
