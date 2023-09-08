package user

import (
	"context"

	"goSimpleAdmin/api/user/admin"
	model "goSimpleAdmin/internal/model/admin"
	"goSimpleAdmin/internal/service"
)

func (c *ControllerAdmin) Captcha(ctx context.Context, req *admin.CaptchaReq) (res *admin.CaptchaRes, err error) {
	data, err := service.User().Captcha(ctx, model.CaptchaReq{})

	if err != nil {
		return
	}

	res = new(admin.CaptchaRes)
	res.CaptchaRes = data
	return
}
