package main

import (
	"flag"
	"fmt"

	"github.com/wushengyouya/coin_exchange/ucenter-api/internal/config"
	"github.com/wushengyouya/coin_exchange/ucenter-api/internal/handler"
	"github.com/wushengyouya/coin_exchange/ucenter-api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()
	logx.MustSetup(logx.LogConf{Stat: false, Encoding: "plain"})
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	// 路由注册
	r := handler.NewRouters(server)
	handler.RegisterHandlers(r, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
