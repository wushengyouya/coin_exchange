package svc

import (
	"github.com/wushengyouya/coin_exchange/grpc_common/ucclient"
	"github.com/wushengyouya/coin_exchange/ucenter-api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UCRegisterRpc ucclient.Register
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UCRegisterRpc: ucclient.NewRegister(zrpc.MustNewClient(c.UCenterRpcConf)),
	}
}
