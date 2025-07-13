package svc

import (
	"LingXi/app/like/cmd/rpc/likecenter"
	"LingXi/app/post/cmd/rpc/internal/config"
	"LingXi/app/post/model"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
	PostModel   model.PostModel
	//PostImageModel model.PostImageModel
	LikeCenterRpc likecenter.Likecenter
	UserCenterRpc usercenter.UserCenter
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
		Config:        c,
		RedisClient:   rdb,
		PostModel:     model.NewPostModel(db, c.Cache),
		UserCenterRpc: usercenter.NewUserCenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		//PostImageModel: model.NewPostImageModel(db, c.Cache),
		LikeCenterRpc: likecenter.NewLikecenter(zrpc.MustNewClient(c.LikecenterRpcConf)),
	}
}
