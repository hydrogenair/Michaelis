package svc

import (
	"LingXi/app/images/cmd/rpc/internal/config"
	"LingXi/app/images/model"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
	ImageModel  model.ImagesModel
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
		RedisClient: rdb,
		Config:      c,
		ImageModel:  model.NewImagesModel(db, c.Cache),
	}
}
