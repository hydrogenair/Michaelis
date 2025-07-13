package svc

import (
	"LingXi/app/chat/cmd/api/internal/config"
	"LingXi/app/chat/cmd/rpc/chatcenter"
	"LingXi/common/chat"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	Hub           *chat.Hub
	ChatCenterRpc chatcenter.Chatcenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	hub := chat.NewHub()

	return &ServiceContext{
		Config:        c,
		Hub:           hub,
		ChatCenterRpc: chatcenter.NewChatcenter(zrpc.MustNewClient(c.ChatCenterRpcConf)),
	}
}
