package svc

import (
	"LingXi/app/comment/cmd/rpc/internal/config"
	"LingXi/app/comment/model"
	"LingXi/app/images/cmd/rpc/imagecenter"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	model2 "LingXi/app/usercenter/model"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Client
	CommentModel model.CommentsModel
	UserModel    model2.UserModel
	//ImagesModel   model3.ImagesModel
	UserCenterRpc usercenter.UserCenter
	ImageCenter   imagecenter.ImageCenter
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
		CommentModel: model.NewCommentsModel(db, c.Cache),
		UserModel:    model2.NewUserModel(db, c.Cache),

		UserCenterRpc: usercenter.NewUserCenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		ImageCenter:   imagecenter.NewImageCenter(zrpc.MustNewClient(c.ImagecenterRpcConf)),
	}
}
