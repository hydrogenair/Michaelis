package svc

import (
	"LingXi/app/chat/cmd/rpc/chatcenter"
	"LingXi/app/connection/cmd/api/internal/config"
	"LingXi/app/connection/cmd/rpc/connectioncenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	ConnectionCenter connectioncenter.Connectioncenter
	ChatCenter       chatcenter.Chatcenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		ConnectionCenter: connectioncenter.NewConnectioncenter(zrpc.MustNewClient(c.ConnectionRpcConf)),
		ChatCenter:       chatcenter.NewChatcenter(zrpc.MustNewClient(c.ChatRpcConf)),
	}
}
