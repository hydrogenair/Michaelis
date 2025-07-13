package svc

import (
	"LingXi/app/tweet/cmd/api/internal/config"
	"LingXi/app/tweet/cmd/rpc/tweetcenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	TweetCenterRpc tweetcenter.TweetCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TweetCenterRpc: tweetcenter.NewTweetCenter(zrpc.MustNewClient(c.TweetcenterRpcConf)),
	}
}
