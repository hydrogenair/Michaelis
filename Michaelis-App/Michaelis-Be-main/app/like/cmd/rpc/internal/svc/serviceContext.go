package svc

import (
	"LingXi/app/like/cmd/rpc/internal/config"
	"LingXi/app/like/model"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
	LikeModel   model.LikeModel
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
		LikeModel:   model.NewLikeModel(db, c.Cache),
	}
}
