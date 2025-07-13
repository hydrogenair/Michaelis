package svc

import (
	"LingXi/app/star/cmd/rpc/internal/config"
	"LingXi/app/star/model"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config        config.Config
	RedisClient   *redis.Client
	StarsModel    model.StarsModel
	UserStarModel model.UserStarModel
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
		RedisClient:   rdb,
		Config:        c,
		StarsModel:    model.NewStarsModel(db, c.Cache),
		UserStarModel: model.NewUserStarModel(db, c.Cache),
	}
}
