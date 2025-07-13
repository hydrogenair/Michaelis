package svc

import (
	"LingXi/app/tweet/cmd/rpc/internal/config"
	"LingXi/app/tweet/model"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
	TweetModel  model.TweetModel
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
		Config:      c,
		RedisClient: rdb,
		TweetModel:  model.NewTweetModel(db, c.Cache),
	}
}
