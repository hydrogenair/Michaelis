package svc

import (
	"LingXi/app/usercenter/cmd/api/internal/config"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserCenterRpc usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserCenterRpc: usercenter.NewUserCenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
