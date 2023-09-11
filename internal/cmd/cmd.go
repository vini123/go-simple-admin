package cmd

import (
	"context"
	"goSimpleAdmin/internal/controller/user"
	"goSimpleAdmin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(service.Middleware().Resp)

			s.Group("/api/admin", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
				)

				group.ALLMap(g.Map{
					"/captcha":      user.NewAdmin().Captcha,
					"/user/sign-in": user.NewAdmin().SignIn,
					"/user/sign-up": user.NewAdmin().SignUp,
				})

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)

					group.ALLMap(g.Map{
						"/user/info": user.NewAdmin().UserInfo,
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
