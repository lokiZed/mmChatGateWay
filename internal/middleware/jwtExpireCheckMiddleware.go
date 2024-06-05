package middleware

import (
	"encoding/json"
	xhttp "github.com/zeromicro/x/http"
	"mmChat/internal/define"
	"net/http"
	"time"
)

type JwtExpireCheckMiddleware struct {
}

func NewJwtExpireCheckMiddleware() *JwtExpireCheckMiddleware {
	return &JwtExpireCheckMiddleware{}
}

func (m *JwtExpireCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expireAtStr := r.Context().Value(define.JwtExpireAt).(json.Number)
		expireAt, _ := expireAtStr.Int64()
		if time.Now().Unix() > expireAt {
			// token过期
			xhttp.JsonBaseResponseCtx(r.Context(), w, define.ErrorTokenExpired)
			return
		}
		next(w, r)
	}
}
