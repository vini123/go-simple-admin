package user

import (
	"context"

	"goSimpleAdmin/api/user/admin"
	"goSimpleAdmin/internal/model"
	"goSimpleAdmin/internal/service"
)

func (c *ControllerAdmin) Captcha(ctx context.Context, req *admin.CaptchaReq) (res *admin.CaptchaRes, err error) {
	service.User().Captcha(ctx, model.CaptchaInput{})
	return
}
