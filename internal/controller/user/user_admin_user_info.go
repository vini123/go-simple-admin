package user

import (
	"context"

	"goSimpleAdmin/api/user/admin"
	model "goSimpleAdmin/internal/model/admin"
	"goSimpleAdmin/internal/service"
)

func (c *ControllerAdmin) UserInfo(ctx context.Context, req *admin.UserInfoReq) (res *admin.UserInfoRes, err error) {

	data, err := service.User().UserInfo(ctx, model.UserInfoReq{})

	if err != nil {
		return
	}

	res = new(admin.UserInfoRes)

	res.UserInfoRes = data

	return res, err
}
