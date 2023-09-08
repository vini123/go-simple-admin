package captcha

import (
	"context"
	"goSimpleAdmin/internal/consts"
	"sync"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

var captcha *Captcha

type CaptchaConfig struct {
	Width      int     `json:"width"`
	Height     int     `json:"Height"`
	Length     int     `json:"length"`
	Maxskew    float64 `json:"maxskew"`
	Dotcount   int     `json:"dotcount"`
	ExpireTime float32 `json:"expire_time"`
}

func NewCaptcha(ctx context.Context) *Captcha {
	once.Do(func() {
		captcha = &Captcha{}

		store := RedisStore{
			Ctx:       ctx,
			KeyPrefix: consts.REDIS_PREFIX + ":captcha:",
		}

		// 配置文件
		captchaConfig := &CaptchaConfig{}
		err := g.Cfg("captcha").MustGet(ctx, "captcha").Scan(captchaConfig)
		if err != nil {
			gerror.Wrap(err, "captcha 配置错误")
			return
		}

		driver := base64Captcha.NewDriverDigit(
			captchaConfig.Height,
			captchaConfig.Width,
			captchaConfig.Length,
			captchaConfig.Maxskew,
			captchaConfig.Dotcount,
		)

		captcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return captcha
}

// 生成图片验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// 验证验证码
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {
	return c.Base64Captcha.Verify(id, answer, false)
}
