package middleware

import (
	"context"
	xhttp "github.com/zeromicro/x/http"
	"mmChat/internal/define"
	"net/http"
	"strings"
)

type IpCheckMiddleware struct {
}

func NewIpCheckMiddleware() *IpCheckMiddleware {
	return &IpCheckMiddleware{}
}

func (m *IpCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 在请求头中检查 X-Forwarded-For 或 X-Real-IP 字段
		xForwardIp := r.Header.Get(define.XForwardFor)
		if xForwardIp != "" {
			xForwardIp = strings.Split(xForwardIp, ",")[0]
		}
		xRealIp := r.Header.Get(define.XRealIp)
		if (xForwardIp == "" && xRealIp == "") || (xForwardIp != "" && xRealIp != "" && xForwardIp != xRealIp) {
			xhttp.JsonBaseResponseCtx(r.Context(), w, define.ErrorInvalidRequest)
			return
		}
		finalRealIp := xForwardIp
		if finalRealIp == "" {
			finalRealIp = xRealIp
		}
		// ip设置到ctx 后续中间件使用
		r = r.WithContext(context.WithValue(r.Context(), define.CtxKeyIp, finalRealIp))

		next(w, r)
	}
}
