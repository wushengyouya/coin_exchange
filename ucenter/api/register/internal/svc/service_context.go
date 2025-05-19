package svc

import (
	"github.com/wushengyouya/coin_exchange/coin-common/msdb"
	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/config"
	"github.com/wushengyouya/coin_exchange/ucenter/api/register/internal/database"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	Db     *msdb.MsDB
}

func NewServiceContext(c config.Config) *ServiceContext {
	c2 := cache.New(c.CacheRedis, nil, cache.NewStat("mscoin"), nil, func(o *cache.Options) {})
	return &ServiceContext{
		Config: c,
		Cache:  c2,
		Db:     database.ConnMysql(c.Mysql.DataSource),
	}
}
