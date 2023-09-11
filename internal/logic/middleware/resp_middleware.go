package middleware

import (
	"goSimpleAdmin/internal/pkg/response"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Resp(r *ghttp.Request) {
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		httpStatus = http.StatusForbidden
		err        = r.GetError()
		res        = r.GetHandlerResponse()
	)

	status := make(map[int]int)
	status[http.StatusUnauthorized] = http.StatusUnauthorized
	status[http.StatusForbidden] = http.StatusForbidden
	status[http.StatusTooManyRequests] = http.StatusTooManyRequests
	status[http.StatusInternalServerError] = http.StatusInternalServerError

	// err 存在，错误码暂时都为 403（403,500 不处理应该都在里边，后边要分开）
	// err 存在，res 单独定义
	// err 不存在，如果有错误，用户自己应传递 status 和 message
	if err != nil {
		res = g.Map{
			"message": err.Error(),
		}
	} else {
		if r.Response.Status == 0 || r.Response.Status == http.StatusOK {
			httpStatus = http.StatusOK
		} else {
			if _, ok := status[r.Response.Status]; ok {
				httpStatus = status[r.Response.Status]
			}
		}
	}
	response.Json(r, httpStatus, res)
}
