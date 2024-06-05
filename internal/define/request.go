package define

import "time"

const (
	XForwardFor = "X-Forwarded-For"
	XRealIp     = "X-Real-IP"
	CtxKeyIp    = "ip"
)

const (
	PerIpExpireTime = time.Second
)
