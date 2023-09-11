package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"goSimpleAdmin/internal/pkg/jwt"
	"goSimpleAdmin/internal/service"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) Auth(r *ghttp.Request) {
	jwt.NewJwt().MiddlewareFunc()(r)

	r.Middleware.Next()
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddleware) Resp(r *ghttp.Request) {
	r.Middleware.Next()

	Resp(r)
}
