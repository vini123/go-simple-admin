package user

import (
	"context"
	"goSimpleAdmin/internal/dao"
	model "goSimpleAdmin/internal/model/admin"
	"goSimpleAdmin/internal/model/do"
	"goSimpleAdmin/internal/pkg/captcha"
	"goSimpleAdmin/internal/pkg/hash"
	"goSimpleAdmin/internal/pkg/jwt"
	"goSimpleAdmin/internal/pkg/utils"
	"goSimpleAdmin/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) SignIn(ctx context.Context, in model.SignInReq) (res model.SignInRes, err error) {
	var user *model.User
	err = dao.Users.Ctx(ctx).Where(do.Users{
		Account: in.Account,
	}).WhereOr(do.Users{
		Email: in.Account,
	}).WhereOr(do.Users{
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

	record, err := dao.Users.Ctx(ctx).Where(do.Users{
		Account: in.Account,
	}).WhereOr(do.Users{
		Phone: in.Account,
	}).WhereOr(do.Users{
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

	data := do.Users{
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

	lastInsertId, err := dao.Users.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		err = gerror.New("请稍后再试")
		return
	}

	var user *model.User
	err = dao.Users.Ctx(ctx).Where(do.Users{
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
	var user *model.User

	id := jwt.NewJwt().GetIdentity(ctx)

	// 用户的基本信息
	err = dao.Users.Ctx(ctx).WithAll().Where(do.Users{
		Id: id,
	}).Scan(&user)
	if err != nil {
		return
	}
	res.User = *user

	// 所有角色
	sql := `SELECT r.id, r.name, r.title FROM roles as r 
			LEFT JOIN model_has_roles as mr ON r.id = mr.role_id
			WHERE mr.model_id = ` + gconv.String(user.Id)
	result, err := g.DB().Query(ctx, sql)
	if err != nil {
		return
	}
	res.Roles = result.List()

	// 用户拥有的权限
	sql = `SELECT DISTINCT p.id,p.guard_name,p.name,p.icon, p.title, p.path,p.parent_id,p.hidden,p.always_show,p.component,p.link,p.iframe,p.redirect,p.order FROM permissions AS p
			LEFT JOIN role_has_permissions AS rp ON rp.permission_id = p.id
			LEFT JOIN model_has_roles AS mr ON mr.role_id = rp.role_id
			LEFT JOIN users AS u ON u.id = mr.model_id
			WHERE u.id =` + gconv.String(user.Id)

	result, err = g.DB().Query(ctx, sql)
	if err != nil {
		return
	}

	var ps []model.Permission
	err = result.Structs(&ps)
	if err != nil {
		return
	}
	res.Permissions = transPermissions(ps, 0)

	return res, err
}

// 将权限递归处理一下
func transPermissions(ps []model.Permission, parent_id uint64) []model.Permission {
	temp := []model.Permission{}
	for _, item := range ps {
		if item.ParentId == parent_id {
			item.Permission = transPermissions(ps, item.Id)
			temp = append(temp, item)
		}
	}
	return temp
}
