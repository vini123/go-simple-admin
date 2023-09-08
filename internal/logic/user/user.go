package user

import (
	"context"
	"goSimpleAdmin/internal/dao"
	model "goSimpleAdmin/internal/model/admin"
	"goSimpleAdmin/internal/model/do"
	"goSimpleAdmin/internal/model/entity"
	"goSimpleAdmin/internal/pkg/captcha"
	"goSimpleAdmin/internal/service"

	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) SignIn(ctx context.Context, in model.SignInReq) (res model.SignInRes, err error) {
	var user *entity.User

	password := gsha1.Encrypt(in.Password)

	err = dao.User.Ctx(ctx).Where(do.User{
		Passport: in.Passport,
		Password: password,
	}).WhereOr(do.User{
		Email:    in.Passport,
		Passport: password,
	}).WhereOr(do.User{
		Phone:    in.Passport,
		Passport: password,
	}).Scan(&user)

	if err != nil {
		return res, err
	}

	if user == nil {
		return res, gerror.New(password)
	}

	err = gerror.New("我的天呀")

	return res, err
}

func (s *sUser) SignUp(ctx context.Context, in model.SignUpReq) (res model.SignUpRes, err error) {

	err = gerror.New("我的地呀")

	return res, err
}

func (s *sUser) Captcha(ctx context.Context, in model.CaptchaReq) (res model.CaptchaRes, err error) {
	id, b64s, err := captcha.NewCaptcha(ctx).GenerateCaptcha()

	if err != nil {
		return
	}

	res.Id = id
	res.Base64 = b64s
	return
}
