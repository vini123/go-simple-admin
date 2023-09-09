package user

import (
	"context"

	"goSimpleAdmin/api/user/admin"
	model "goSimpleAdmin/internal/model/admin"
	"goSimpleAdmin/internal/service"
)

func (c *ControllerAdmin) SignUp(ctx context.Context, req *admin.SignUpReq) (res *admin.SignUpRes, err error) {
	data, err := service.User().SignUp(ctx, model.SignUpReq{
		Nickname:   req.Nickname,
		Account:    req.Account,
		Password:   req.Password,
		VerifyKey:  req.VerifyKey,
		VerifyCode: req.VerifyCode,
	})

	if err != nil {
		return
	}

	res = new(admin.SignUpRes)
	res.SignUpRes = data
	return res, err
}
