package user

import (
	"context"
	"goSimpleAdmin/internal/dao"
	model "goSimpleAdmin/internal/model/admin"
	"goSimpleAdmin/internal/model/do"
	"goSimpleAdmin/internal/model/entity"
	"goSimpleAdmin/internal/pkg/captcha"
	"goSimpleAdmin/internal/pkg/hash"
	"goSimpleAdmin/internal/pkg/jwt"
	"goSimpleAdmin/internal/pkg/utils"
	"goSimpleAdmin/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
	err = dao.User.Ctx(ctx).Where(do.User{
		Account: in.Account,
	}).WhereOr(do.User{
		Email: in.Account,
	}).WhereOr(do.User{
		Phone: in.Account,
	}).Scan(&user)

	if err != nil {
		return res, err
	}

	if user == nil {
		return res, gerror.New("账号不存在")
	}

	check := hash.BcryptCheck(in.Password, user.Password)
	if !check {
		return res, gerror.New("账号或密码错误")
	}

	payloadData := g.Map{
		"id":       user.Id,
		"nickname": user.Nickname,
	}

	res.Token, res.Expire, err = jwt.NewJwt().TokenGenerator(payloadData)

	return res, err
}

func (s *sUser) SignUp(ctx context.Context, in model.SignUpReq) (res model.SignUpRes, err error) {
	verifyCode := in.VerifyCode
	verifyKey := in.VerifyKey
	match := captcha.NewCaptcha(ctx).VerifyCaptcha(verifyKey, verifyCode)

	if !match {
		err = gerror.New("验证码不正确或已过期")
		return
	}

	record, err := dao.User.Ctx(ctx).Where(do.User{
		Account: in.Account,
	}).WhereOr(do.User{
		Phone: in.Account,
	}).WhereOr(do.User{
		Email: in.Account,
	}).One()

	if err != nil {
		err = gerror.New("请稍后再试")
		return
	}

	if !record.IsEmpty() {
		err = gerror.New("账号已经存在")
		return
	}

	data := do.User{
		Account:  in.Account,
		Password: in.Password,
		Nickname: in.Nickname,
	}

	if utils.IsPhone(in.Account) {
		data.Phone = in.Account
	}

	if utils.IsEmail(in.Account) {
		data.Email = in.Account
	}

	lastInsertId, err := dao.User.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		err = gerror.New("请稍后再试")
		return
	}

	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Id: lastInsertId,
	}).Scan(&user)

	if err != nil {
		err = gerror.New("请稍后再试")
		return
	}

	if user == nil {
		err = gerror.New("请稍后再试")
		return
	}

	payloadData := g.Map{
		"id":       user.Id,
		"nickname": user.Nickname,
	}

	res.Token, res.Expire, err = jwt.NewJwt().TokenGenerator(payloadData)

	return res, err
}

func (s *sUser) Captcha(ctx context.Context, in model.CaptchaReq) (res model.CaptchaRes, err error) {
	id, b64s, err := captcha.NewCaptcha(ctx).GenerateCaptcha()

	if err != nil {
		return res, err
	}

	res.Id = id
	res.Base64 = b64s
	return
}

func (s *sUser) UserInfo(ctx context.Context, in model.UserInfoReq) (res model.UserInfoRes, err error) {
	var user *entity.User

	id := jwt.NewJwt().GetIdentity(ctx)

	err = dao.User.Ctx(ctx).Where(do.User{
		Id: id,
	}).Scan(&user)

	if err != nil {
		return res, err
	}

	res.User = *user
	return res, err
}
