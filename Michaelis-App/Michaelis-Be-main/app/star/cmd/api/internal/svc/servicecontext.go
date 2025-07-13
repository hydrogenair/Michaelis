package svc

import (
	"LingXi/app/star/cmd/api/internal/config"
	"LingXi/app/star/cmd/rpc/starcenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	StarCenter starcenter.Starcenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		StarCenter: starcenter.NewStarcenter(zrpc.MustNewClient(c.StarcenterRpcConf)),
	}
}
