package svc

import (
	"LingXi/app/images/cmd/api/internal/config"
	"LingXi/app/images/cmd/rpc/imagecenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	ImageCenter imagecenter.ImageCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		ImageCenter: imagecenter.NewImageCenter(zrpc.MustNewClient(c.ImagecenterRpcConf)),
	}
}
