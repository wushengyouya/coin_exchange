package svc

import "github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
