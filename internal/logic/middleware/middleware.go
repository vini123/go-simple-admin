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
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"localhost:5173"}
	corsOptions.AllowOrigin = "http://localhost:5174"
	corsOptions.AllowCredentials = "false"
	corsOptions.AllowMethods = "POST,GET,OPTIONS,PUT,PATCH,DELETE"
	corsOptions.AllowHeaders = "Accept,Content-Type,Referer,User-Agent,Origin,X-Requested-With,X-XSRF-TOKEN,X-CSRF-TOKEN,Authorization,Time"
	r.Response.CORS(corsOptions)

	// r.Response.CORSDefault()

	r.Middleware.Next()
}

func (s *sMiddleware) Resp(r *ghttp.Request) {
	r.Middleware.Next()

	Resp(r)
}
