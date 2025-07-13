package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	Email struct {
		Host     string
		Port     int
		UserName string
		Password string
	}
	VerifyCode struct {
		CodeLength int
		ExpireTime int
	}
	APP struct {
		Name string
	}
	Mysql gormc.Mysql
	Cache cache.CacheConf
}
