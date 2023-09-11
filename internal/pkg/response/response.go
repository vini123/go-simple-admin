package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// 主要处理 200 401 403 429 500 这几个状态
func Json(r *ghttp.Request, status int, res interface{}) {
	r.Response.WriteHeader(status)
	r.Response.WriteJson(res)
	r.ExitAll()
}
