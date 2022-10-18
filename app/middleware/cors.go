package middleware

import "github.com/gogf/gf/v2/net/ghttp"

var Cors = corsMiddleware{}

type corsMiddleware struct{}

// 允许接口跨域请求
func (s *corsMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
