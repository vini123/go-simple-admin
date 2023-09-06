package cmd

import (
	"context"
	"goSimpleAdmin/internal/controller/user"

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
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/api/admin", func(group *ghttp.RouterGroup) {

				group.Bind(
					user.NewAdmin(),
				)
			})
			s.Run()
			return nil
		},
	}
)
