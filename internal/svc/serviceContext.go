package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"mmChat/internal/config"
	"mmChat/internal/middleware"
)

type ServiceContext struct {
	Config              config.Config
	JwtExpireCheck      rest.Middleware
	IpCheck             rest.Middleware
	RateLimitingControl rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		JwtExpireCheck:      middleware.NewJwtExpireCheckMiddleware().Handle,
		IpCheck:             middleware.NewIpCheckMiddleware().Handle,
		RateLimitingControl: middleware.NewRateLimitingControlMiddleware().Handle,
	}
}
