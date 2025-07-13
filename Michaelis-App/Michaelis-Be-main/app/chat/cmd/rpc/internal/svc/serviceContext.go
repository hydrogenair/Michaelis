package svc

import (
	"LingXi/app/chat/cmd/rpc/internal/config"
	"LingXi/app/chat/model"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Client
	MessageModel model.MessageModel
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
		Config:       c,
		RedisClient:  rdb,
		MessageModel: model.NewMessageModel(db, c.Cache),
	}
}
