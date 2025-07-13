package svc

import (
	"LingXi/app/connection/cmd/rpc/internal/config"
	"LingXi/app/connection/model"
	"LingXi/app/images/cmd/rpc/imagecenter"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	RedisClient     *redis.Client
	ConnectionModel model.ConnectionModel
	UserCenterRpc   usercenter.UserCenter
	ImageCenter     imagecenter.ImageCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: "",
		DB:       0,
	})
	db, err := gormc.ConnectMysql(c.Mysql)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:          c,
		RedisClient:     rdb,
		ConnectionModel: model.NewConnectionModel(db, c.Cache),
		UserCenterRpc:   usercenter.NewUserCenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		ImageCenter:     imagecenter.NewImageCenter(zrpc.MustNewClient(c.ImagecenterRpcConf)),
	}
}
