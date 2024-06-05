package middleware

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"mmChat/internal/define"
	"net/http"
)

type RateLimitingControlMiddleware struct {
}

func NewRateLimitingControlMiddleware() *RateLimitingControlMiddleware {
	return &RateLimitingControlMiddleware{}
}

func (m *RateLimitingControlMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.Context().Value(define.CtxKeyIp).(string)
		_, err := ipCache.Take(ip, func() (any, error) {
			ipCache.SetWithExpire(ip, "", define.PerIpExpireTime)
			return "", errors.New("new")
		})
		if err == nil {
			// 说明不是刚加的 请求过快 驳回
			xhttp.JsonBaseResponseCtx(r.Context(), w, define.ErrorRequestTooQuick)
			return
		}
		next(w, r)
	}
}
