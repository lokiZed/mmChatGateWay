package define

import (
	"github.com/zeromicro/x/errors"
)

var (
	ErrorInvalidRequest  = errors.New(ResponseCodePanic, "非法请求")
	ErrorRequestTooQuick = errors.New(ResponseCodePanic, "请求过快")
	ErrorTokenExpired    = errors.New(ResponseCodePanic, "登陆过期")
)
